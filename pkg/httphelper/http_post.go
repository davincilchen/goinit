package httphelper

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func PostJSON_NoRetry(url string, requestBody interface{}) (
	resp *http.Response, reqHeader string, err error) {

	j, err := json.Marshal(requestBody)
	if err != nil {
		return nil, reqHeader, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	if err != nil {
		return nil, reqHeader, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		return nil, reqHeader, err
	}
	return resp, reqHeader, err
}

func PostJSON(url string, requestBody interface{}, retries int) (
	res *http.Response, err error) {

	return RequestCUDJSONVersion(http.MethodPost, url, requestBody, retries)
}

func Post(url string) (
	res *http.Response, err error) {

	return HttpDo(http.MethodPost, url)
}
