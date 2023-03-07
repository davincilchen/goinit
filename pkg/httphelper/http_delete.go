package httphelper

import (
	"net/http"
)

func DeleteJSON(url string, requestBody interface{}, retries int) (*http.Response, error) {
	return RequestCUDJSONVersion(http.MethodDelete, url, requestBody, retries)
}
