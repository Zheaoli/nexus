package wechat

import (
	"bytes"
	"encoding/json"
	"github.com/Zheaoli/nexus/requests"
	"io/ioutil"
	"net/http"
)

func GetRequest(url string, params map[string]interface{}) (string, error) {
	urlPath, err := requests.MakeURL(url, params)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(urlPath)
	if err != nil {
		return "", err
	}
	message, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := resp.Body.Close(); err != nil {
		return "", err
	}
	return string(message), nil
}

func PostRequest(url string, params, data map[string]interface{}) (string, error) {
	requestData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	urlPath, err := requests.MakeURL(url, params)
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	message, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := resp.Body.Close(); err != nil {
		return "", err
	}
	return string(message), nil
}
