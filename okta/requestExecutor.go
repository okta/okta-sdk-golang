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
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/okta/okta-sdk-golang/okta/query"
)

type RequestExecutor struct {
	httpClient *http.Client
	config     *Config
}

func NewRequestExecutor(httpClient *http.Client, config *Config) *RequestExecutor {
	re := RequestExecutor{}
	re.httpClient = httpClient
	re.config = config

	if httpClient == nil {
		re.httpClient = &http.Client{}
	}

	return &re
}

func (re *RequestExecutor) Get(url string, qp *query.Params) ([]byte, error) {
	return re.doRequest("GET", url, nil, qp)
}

func (re *RequestExecutor) Post(url string, body io.Reader, qp *query.Params) ([]byte, error) {
	return re.doRequest("POST", url, body, qp)
}

func (re *RequestExecutor) Put(url string, body io.Reader, qp *query.Params) ([]byte, error) {
	return re.doRequest("PUT", url, body, qp)
}

func (re *RequestExecutor) Delete(url string, qp *query.Params) ([]byte, error) {
	return re.doRequest("DELETE", url, nil, qp)
}

func (re *RequestExecutor) doRequest(method string, url string, body io.Reader, qp *query.Params) ([]byte, error) {
	url = re.config.Okta.Client.OrgUrl + url
	if &qp != nil {
		url = url + qp.String()
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if method == "POST" && body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Authorization", "SSWS "+re.config.Okta.Client.Token)
	req.Header.Add("User-Agent", NewUserAgent(re.config).String())
	req.Header.Add("Accept", "application/json")

	resp, err := re.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	e := new(Error)
	json.Unmarshal(bodyBytes, &e)
	if e.ErrorId != "" {
		return nil, e
	}

	return bodyBytes, nil
}
