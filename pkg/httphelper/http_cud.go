package httphelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var clientCUD = client

func reTry(fn func() (*http.Response, error), retries int) (res *http.Response, err error) {
	allError := ""
	for i := 0; i <= retries; i++ {
		res, err = fn()
		if err == nil {
			return res, nil
		}
		err = fmt.Errorf("current try is %d , %s", i+1, err.Error())
		allError = fmt.Sprintf("%s[%s]", allError, err.Error())
	}

	if allError != "" {
		//Error
		fmt.Println(allError)
	}
	return res, err
}

func RequestCUDJSONVersion(method, url string, requestBody interface{}, retries int) (resp *http.Response, err error) {
	now := time.Now()
	reqBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	response, err := RequestCUDJSON(method, url, reqBody, retries)
	if response == nil { //no error if http 405
		NilResponseLog(APICommentNormalV, method, url, string(reqBody), time.Now().Sub(now), err)
	} else {
		ResponseLog(APICommentNormalV, string(reqBody), response, time.Now().Sub(now), err)
	}

	return response, err
}

func RequestCUDJSON(method, url string, requestBody []byte, retries int) (resp *http.Response, err error) {

	resp, err = reTry(
		func() (*http.Response, error) {
			now := time.Now()
			req, theErr := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
			if theErr != nil {
				s := fmt.Sprint(time.Since(now))
				errf := fmt.Errorf("[NewRequest Error], url:%s, method:%s, error:%v, %s, request:%s", url, method, theErr, s, requestBody)
				return nil, errf
			}
			req.Header.Add("Content-Type", "application/json")
			//req.Close = true
			response, theErr := clientCUD.Do(req)

			if theErr != nil {
				//watchout %s
				s := fmt.Sprint(time.Since(now))
				errf := fmt.Errorf("url:%s,  method:%s, error:%v, %s, request:%s , response:%v", url, method, theErr, s, requestBody, response)
				return response, errf
			}
			return response, nil
		}, retries)

	if resp == nil && err == nil {
		err = fmt.Errorf("url:%s,  method:%s,  request:%s , response is nil", url, method, requestBody)
	}
	// if err != nil {
	// 	return resp, err
	// }

	return resp, err
}
