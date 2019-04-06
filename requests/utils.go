package requests

import url2 "net/url"

func MakeURL(url string, params map[string]interface{}) (string, error) {
	paramData := url2.Values{}
	urlData, err := url2.Parse(url)
	if err != nil {
		return "", err
	}
	for key, value := range params {
		paramData.Set(key, value.(string))
	}
	urlData.RawQuery = paramData.Encode()
	urlPath := urlData.String()
	return urlPath, nil
}
