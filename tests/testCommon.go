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

package tests

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jarcoal/httpmock"

	"github.com/okta/okta-sdk-golang/okta"
)

func NewClient(conf ...okta.ConfigSetter) (*okta.Client, error) {
	return okta.NewClient(context.Background(), conf...)
}

func MockResponse(responses ...*http.Response) httpmock.Responder {
	return func(req *http.Request) (*http.Response, error) {
		httpmock.GetTotalCallCount()
		info := httpmock.GetCallCountInfo()
		count := info[req.Method+" "+req.URL.Path]

		if len(responses) >= count {
			return responses[count-1], nil
		}

		return nil, fmt.Errorf("no response found for call %v to %s", count, req.URL.Path)
	}
}
