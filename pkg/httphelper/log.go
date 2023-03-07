package httphelper

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//NilResponseLog is a
func NilResponseLog(comment, method, url, reqBody string, latency time.Duration, err error) {
	errString := ""
	if err != nil {
		errString = err.Error()
	}

	//Warningf
	fmt.Printf("%s, method:%s, url:%s, reqBody:%s, response:nil, latency:%v,  err:%s",
		comment,
		method,
		url,
		reqBody,
		latency,
		errString)
	fmt.Println()
}

//ResponseLog is a
func ResponseLog(comment, reqBody string, res *http.Response, latency time.Duration, err error) {
	errString := ""
	if err != nil {
		errString = err.Error()
	}
	if res == nil {
		//Errorf
		fmt.Printf("%s, reqBody:%s, responseLog latency:%v, res_body:nil, err:%s",
			comment,
			reqBody,
			latency,
			errString)
		fmt.Println()
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	//Infof
	fmt.Printf("%s, method:%s, url:%s, reqBody:%s, res_statusCode:%3d, latency:%v, res_body:%s, err:%s",
		comment,
		res.Request.Method,
		res.Request.URL.String(),
		reqBody,
		res.StatusCode,
		latency,
		body,
		errString)
	fmt.Println()
}
