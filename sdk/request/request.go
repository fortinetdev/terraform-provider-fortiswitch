package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-fortiswitch/sdk/config"
)

// Request describes the request to Fortiswitch service
type Request struct {
	Config       config.Config
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
}

// Cookies contains the CSRF Token and Cookie by login using username and password
type Cookies struct {
	CSRFToken string
	Cookie    string
}

// New creates reqeust object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the gobal plugin client.
func New(c config.Config, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	var h *http.Request

	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}

	r := &Request{
		Config:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

// Build Request header
// Build Request Sign/Login Info
// Send request data to Fortiswitch.
// If errors are encountered, it returns the error.
func (r *Request) Send() error {
	return r.SendRequest(3)
}

// Send request data to Fortiswitch.
// If errors are encountered, it returns the error.
func (r *Request) SendRequest(retries int) error {
	var err error
	var cookies *Cookies
	cookies, err = r.LoginSession()
	if err != nil {
		log.Printf("[ERROR] Failed to login: %v", err)
		return err
	}
	r.HTTPRequest.Header.Set("x-csrftoken", cookies.CSRFToken)
	r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)

	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	r.HTTPRequest.Header.Set("accept", "application/json")

	u := r.buildURL()

	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		return err
	}

	retry := 0
	for {
		//Send
		log.Default().Printf("[INFO] sending api req: %v", r.HTTPRequest)
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	errLogout := r.LogoutSession(cookies)
	if errLogout != nil {
		log.Printf("[WARNING] Issue occurs when logout session: %v", errLogout)
	}

	return err
}

func filterapikey(v string) string {
	re, _ := regexp.Compile("access_token=.*?\"")
	res := re.ReplaceAllString(v, "access_token=***************\"")

	return res
}

func (r *Request) buildURL() string {
	u := "https://"
	u += r.Config.FwTarget
	u += r.Path

	return u
}

// SendWithSpecialParams sends request data to FortiSwitch with special URL paramaters.
// If errors are encountered, it returns the error.
func (r *Request) SendWithSpecialParams(s string) error {
	var err error
	var cookies *Cookies

	cookies, err = r.LoginSession()
	if err != nil {
		log.Printf("[ERROR] Failed to login: %v", err)
		return err
	}
	r.HTTPRequest.Header.Set("x-csrftoken", cookies.CSRFToken)
	r.HTTPRequest.Header.Set("Cookie", cookies.Cookie)

	r.HTTPRequest.Header.Set("Content-Type", "application/json")
	r.HTTPRequest.Header.Set("accept", "application/json")

	u := r.buildURL()

	if s != "" {
		u += "&"
		u += s
	}

	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > 15 {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Duration(1) * time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}
	return err
}

// Login Fortiswitch using username and password in session mode, and return Cookies.
// If errors are encountered, it returns the error.
func (r *Request) LoginSession() (*Cookies, error) {
	var err error

	data := "username="
	data += r.Config.Auth.Username
	data += "&secretkey="
	data += r.Config.Auth.Password
	data += "&ajax=1"

	bodyBytes := bytes.NewBufferString(data)

	req, _ := http.NewRequest("POST", "", bodyBytes)
	req.Header.Set("Content-Type", "application/json")
	u := "https://"
	u += r.Config.FwTarget
	u += "/logincheck"
	req.URL, err = url.Parse(u)
	if err != nil {
		err = fmt.Errorf("Could not parse URL: %s", err)
		return nil, err
	}

	rsp, err := r.Config.HTTPCon.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "x509: ") {
			err = fmt.Errorf("HTTP request error: %v", err)
			return nil, err
		}
	}

	if rsp == nil {
		err = fmt.Errorf("Host is unreachable. HTTP response is nil.")
		return nil, err
	}

	if rsp.Header == nil {
		err = fmt.Errorf("HTTP response header is nil.")
		return nil, err
	}

	csrfToken := ""
	cookie := ""
	if setCookie, ok := rsp.Header["Set-Cookie"]; ok {
		for _, item := range setCookie {
			reg := regexp.MustCompile(`ccsrftoken="(.*)".*`)
			match := reg.FindStringSubmatch(item)
			if match != nil {
				csrfToken = match[1]
			} else if strings.HasPrefix(item, "APSCOOKIE_") {
				cookie = item
			}
		}
		if csrfToken == "" {
			err = fmt.Errorf("Could not get CSRF Token from HTTP request response.")
			return nil, err
		} else if csrfToken == "0%260" {
			err = fmt.Errorf("Username or password not valid.")
			return nil, err
		}
		if cookie == "" {
			err = fmt.Errorf("Could not get cookie from HTTP request response.")
			return nil, err
		}
	} else {
		err = fmt.Errorf("Logincheck response do not contains Set-Cookie.")
		return nil, err
	}

	ck := &Cookies{
		CSRFToken: csrfToken,
		Cookie:    cookie,
	}

	return ck, nil
}

// Logout current session.
// If errors are encountered, it returns the error.
func (r *Request) LogoutSession(cookies *Cookies) error {
	var err error

	if cookies.CSRFToken == "" {
		log.Printf("[INFO] CSRF Token is empty, no need to logout.")
		return nil
	}

	req, _ := http.NewRequest("POST", "", nil)
	req.Header.Set("x-csrftoken", cookies.CSRFToken)
	req.Header.Set("Cookie", cookies.Cookie)
	req.Header.Set("Content-Type", "application/json")
	u := "https://"
	u += r.Config.FwTarget
	u += "/logout"
	req.URL, err = url.Parse(u)
	if err != nil {
		err = fmt.Errorf("Could not parse URL: %s", err)
		return err
	}

	_, err = r.Config.HTTPCon.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "x509: ") {
			err = fmt.Errorf("HTTP request error: %v", err)
			return err
		}
	}

	return nil
}
