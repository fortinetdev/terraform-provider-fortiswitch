package fortiswitch

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"

	"github.com/terraform-providers/terraform-provider-fortiswitch/sdk/auth"
	forticlient "github.com/terraform-providers/terraform-provider-fortiswitch/sdk/sdkcore"
)

// Config gets the authentication information from the given metadata
type Config struct {
	Hostname        string
	Username        string
	Password        string
	Insecure        *bool
	CABundle        string
	CABundleContent string
	HTTPProxy       string

	PeerAuth   string
	CaCert     string
	ClientCert string
	ClientKey  string
}

// FortiClient contains the basic FortiSwitch SDK connection information to FortiSwitch
// It can be used to as a client of FortiSwitch for the plugin
// Now FortiClient contains two kinds of clients:
// Client is for FortiSwitch
type FortiClient struct {
	Client *forticlient.FortiSDKClient
}

// CreateClient creates a FortiClient Object with the authentication information.
// It returns the FortiClient Object for the use when the plugin is initialized.
func (c *Config) CreateClient() (interface{}, error) {
	var fClient FortiClient

	bFSWExist := bFortiSwitchHostnameExist(c)

	if bFSWExist {
		err := createFortiSwitchClient(&fClient, c)
		if err != nil {
			return nil, fmt.Errorf("Error create fortiswitch client: %v", err)
		}
	}

	return &fClient, nil
}

func bFortiSwitchHostnameExist(c *Config) bool {
	if c.Hostname == "" {
		if os.Getenv("FORTISWITCH_ACCESS_HOSTNAME") == "" {
			return false
		}
	}

	return true
}

func createFortiSwitchClient(fClient *FortiClient, c *Config) error {
	config := &tls.Config{}

	auth := auth.NewAuth(c.Hostname, c.Username, c.Password, c.CABundle, c.CABundleContent, c.PeerAuth, c.CaCert, c.ClientCert, c.ClientKey, c.HTTPProxy)

	if auth.Hostname == "" {
		_, err := auth.GetEnvHostname()
		if err != nil {
			return fmt.Errorf("Error reading Hostname")
		}
	}

	if auth.Username == "" {
		_, err := auth.GetEnvUsername()
		if err != nil {
			return fmt.Errorf("Error reading Username")
		}
	}

	if auth.Password == "" {
		_, err := auth.GetEnvPassword()
		if err != nil {
			return fmt.Errorf("Error reading Password")
		}
	}

	if auth.CABundle == "" {
		auth.GetEnvCABundle()
	}

	if auth.PeerAuth == "" {
		_, err := auth.GetEnvPeerAuth()
		if err != nil {
			return fmt.Errorf("Error reading PeerAuth")
		}
	}
	if auth.CaCert == "" {
		_, err := auth.GetEnvCaCert()
		if err != nil {
			return fmt.Errorf("Error reading CaCert")
		}
	}
	if auth.ClientCert == "" {
		_, err := auth.GetEnvClientCert()
		if err != nil {
			return fmt.Errorf("Error reading ClientCert")
		}
	}
	if auth.ClientKey == "" {
		_, err := auth.GetEnvClientKey()
		if err != nil {
			return fmt.Errorf("Error reading ClientKey")
		}
	}
	if auth.HTTPProxy == "" {
		_, err := auth.GetEnvHTTPProxy()
		if err != nil {
			return fmt.Errorf("Error reading HTTP proxy")
		}
	}

	pool := x509.NewCertPool()

	if auth.CABundle != "" {
		if auth.CABundleContent != "" {
			return fmt.Errorf("\"cabundlefile\" and \"cabundlecontent\" could not exist at the same time! Please only configure one of them. If you are not configure \"cabundlefile\", please check the environment variable \"FORTISWITCH_CA_CABUNDLE\".")
		}

		f, err := os.Open(auth.CABundle)
		if err != nil {
			return fmt.Errorf("Error reading CA Bundle: %v", err)
		}
		defer f.Close()

		caBundle, err := ioutil.ReadAll(f)
		if err != nil {
			return fmt.Errorf("Error reading CA Bundle: %v", err)
		}

		if !pool.AppendCertsFromPEM([]byte(caBundle)) {
			return fmt.Errorf("Error reading CA Bundle")
		}
		config.RootCAs = pool
	} else if auth.CABundleContent != "" {
		if !pool.AppendCertsFromPEM([]byte(auth.CABundleContent)) {
			return fmt.Errorf("Error reading CA Bundle")
		}
		config.RootCAs = pool
	}

	if auth.PeerAuth == "enable" {
		if auth.CaCert != "" {
			caCertFile := auth.CaCert
			caCert, err := ioutil.ReadFile(caCertFile)
			if err != nil {
				return fmt.Errorf("client ioutil.ReadFile couldn't load cacert file: %v", err)
			}

			pool.AppendCertsFromPEM(caCert)
		}

		if auth.ClientCert == "" {
			return fmt.Errorf("User Cert file doesn't exist!")
		}

		if auth.ClientKey == "" {
			return fmt.Errorf("User Key file doesn't exist!")
		}

		clientCert, err := tls.LoadX509KeyPair(auth.ClientCert, auth.ClientKey)
		if err != nil {
			return fmt.Errorf("Client ioutil.ReadFile couldn't load clientCert/clientKey file: %v", err)
		}

		config.Certificates = []tls.Certificate{clientCert}
	}

	if c.Insecure == nil {
		// defaut value for Insecure is false
		b, _ := auth.GetEnvInsecure()
		config.InsecureSkipVerify = b
	} else {
		config.InsecureSkipVerify = *c.Insecure
	}

	if config.InsecureSkipVerify == false && auth.CABundle == "" && auth.CABundleContent == "" {
		return fmt.Errorf("Error getting CA Bundle, CA Bundle should be set when insecure is false")
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	if auth.HTTPProxy != "" {
		var httpProxy *url.URL
		httpProxy, err := url.Parse(auth.HTTPProxy)
		if err != nil {
			return fmt.Errorf("Error parsing HTTP proxy: %w", err)
		}
		tr.Proxy = http.ProxyURL(httpProxy)
	}

	// All users of cookiejar should import "golang.org/x/net/publicsuffix"
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return fmt.Errorf("Error initializing cookiejar %v", err)
	}

	log.Println("Transport config")
	client := &http.Client{
		// setting this makes backend returns no cookie
		Transport: tr,
		Timeout:   time.Second * 250,
		Jar:       jar,
	}

	fc, err := forticlient.NewClient(auth, client)

	if err != nil {
		return fmt.Errorf("connection error: %v", err)
	}

	fClient.Client = fc

	return nil
}
