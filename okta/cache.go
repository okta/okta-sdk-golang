package okta

import (
	"bytes"
	"io"
	"net/http"
)

type Cache interface {
	Get(key string) *http.Response
	Set(key string, value *http.Response)
	GetString(key string) string
	SetString(key string, value string)
	Delete(key string)
	Clear()
	Has(key string) bool
}

func CreateCacheKey(req *http.Request) string {
	s := req.URL.Scheme + "://" + req.URL.Host + req.URL.RequestURI()
	return s
}

func CopyResponse(resp *http.Response) *http.Response {
	c := *resp
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp
	}

	c.Body = io.NopCloser(bytes.NewBuffer(respBody))

	return &c
}
