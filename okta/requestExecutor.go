/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package okta

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/okta/okta-sdk-golang/okta/cache"
)

type RequestExecutor struct {
	httpClient *http.Client
	config     *config
	BaseUrl    *url.URL
	cache      cache.Cache
}

func NewRequestExecutor(httpClient *http.Client, cache cache.Cache, config *config) *RequestExecutor {
	re := RequestExecutor{}
	re.httpClient = httpClient
	re.config = config
	re.cache = cache

	if httpClient == nil {
		tr := &http.Transport{
			IdleConnTimeout: 30 * time.Second,
		}
		re.httpClient = &http.Client{Transport: tr}
	}

	return &re
}

func (re *RequestExecutor) NewRequest(method string, url string, body interface{}) (*http.Request, error) {
	var buff io.ReadWriter
	if body != nil {
		buff = new(bytes.Buffer)
		encoder := json.NewEncoder(buff)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	url = re.config.Okta.Client.OrgUrl + url

	req, err := http.NewRequest(method, url, buff)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "SSWS "+re.config.Okta.Client.Token)
	req.Header.Add("User-Agent", NewUserAgent(re.config).String())
	req.Header.Add("Accept", "application/json")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (re *RequestExecutor) Do(req *http.Request, v interface{}) (*Response, error) {
	retryCount := int32(0)
	requestStart := time.Now().Unix()
	cacheKey := cache.CreateCacheKey(req)
	if req.Method != http.MethodGet {
		re.cache.Delete(cacheKey)
	}
	inCache := re.cache.Has(cacheKey)

	if !inCache {

		for {
			iterationStartTime := time.Now().Unix()
			if (iterationStartTime - requestStart) >= int64(re.config.Okta.Client.RequestTimeout) {
				return nil, errors.New("reached the max request time")
			}

			resp, err := re.httpDo(req)

			if err != nil {
				return nil, err
			}

			if resp.StatusCode >= 200 && resp.StatusCode <= 299 {

				if re.config.Okta.Client.RateLimit.MaxRetries > 0 {

				}

				respBody, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return nil, err
				}

				origResp := ioutil.NopCloser(bytes.NewBuffer(respBody))
				resp.Body = origResp

				if req.Method == http.MethodGet && reflect.TypeOf(v).Kind() != reflect.Slice && (resp.StatusCode >= 200 && resp.StatusCode <= 299) {
					re.cache.Set(cacheKey, resp)
				}

				return buildResponse(resp, &v)
			}

			if retryCount > re.config.Okta.Client.RateLimit.MaxRetries {
				return nil, errors.New("reached the max number of 429 retries")
			}

			retryCount++

			err = re.pauseBeforeRetry(retryCount, resp)
			if err != nil {
				return nil, err
			}
		}

	}

	resp := re.cache.Get(cacheKey)
	return buildResponse(resp, &v)

}

func (re *RequestExecutor) httpDo(request *http.Request) (*http.Response, error) {
	resp, err := re.httpClient.Do(request)
	return resp, err
}

func (re *RequestExecutor) pauseBeforeRetry(retryCount int32, response *http.Response) error {
	if response.StatusCode == http.StatusTooManyRequests {
		resetLimit, _ := strconv.Atoi(response.Header.Get("X-Rate-Limit-Reset"))
		requestDate, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 Z", response.Header.Get("Date"))
		requestDateUnix := requestDate.Unix()

		backoffSeconds := int64(resetLimit) - requestDateUnix + 1
		time.Sleep(time.Duration(backoffSeconds) * time.Second)

		return nil
	}

	return nil
}

type Response struct {
	*http.Response
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func CheckResponseForError(resp *http.Response) error {
	statusCode := resp.StatusCode
	if statusCode >= http.StatusOK && statusCode < http.StatusBadRequest {
		return nil
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	e := new(Error)
	json.Unmarshal(bodyBytes, &e)
	return e

}

func buildResponse(resp *http.Response, v interface{}) (*Response, error) {
	response := newResponse(resp)

	err := CheckResponseForError(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		decodeError := json.NewDecoder(resp.Body).Decode(v)
		if decodeError == io.EOF {
			decodeError = nil
		}
		if decodeError != nil {
			err = decodeError
		}

	}
	return response, err
}
