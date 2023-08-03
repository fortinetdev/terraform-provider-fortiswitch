package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func createUpdate(c *FortiSDKClient, method string, path string, params *map[string]interface{}, output map[string]interface{}) (err error) {
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest(method, path, nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	log.Printf("[INFO] str body %v", string(body))

	err = fortiAPIErrorFormat(result, string(body))
	log.Printf("[INFO] result %v", result)

	if err == nil {
		if result["mkey"] != nil {
			output["mkey"] = result["mkey"]
		} else if result["results"] != nil {
			log.Printf("[INFO] RESULTS body %v", result["results"])
			output["mkey"] = result["results"].(map[string]interface{})["mkey"]
		}
	}

	return
}

func delete(c *FortiSDKClient, method string, path string) (err error) {

	req := c.NewRequest(method, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}
	log.Printf("FSW-fortiswitch response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

func read(c *FortiSDKClient, method string, path string, bcomplex bool) (mapTmp map[string]interface{}, err error) {
	req := c.NewRequest(method, path, nil, nil)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close() //#

	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}
	log.Printf("FSW-fortiswitch reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if fortiAPIHttpStatus404Checking(result) == true {
		mapTmp = nil
		return
	}

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		if bcomplex {
			mapTmp = result["results"].(map[string]interface{})
		} else {
			mapTmp = (result["results"].([]interface{}))[0].(map[string]interface{})
		}
	}

	return
}
