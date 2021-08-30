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
	"net/http"
	"strconv"
	"time"

	"github.com/jarcoal/httpmock"
)

func Mock429Response() *http.Response {
	loc, _ := time.LoadLocation("UTC")
	zulu := time.Now().In(loc)
	header := http.Header{}
	header.Add("X-Okta-Now", strconv.FormatInt(zulu.Unix(), 10))
	header.Add("X-Rate-Limit-Reset", strconv.FormatInt(time.Now().Unix()+1, 10))
	header.Add("X-Okta-Request-id", "a-request-id")
	header.Add("Content-Type", "application/json")
	header.Add("Date", zulu.Format("Mon, 02 Jan 2006 15:04:05 GMT"))

	return &http.Response{
		Status:        strconv.Itoa(429),
		StatusCode:    429,
		Body:          httpmock.NewRespBodyFromString("{}"),
		Header:        header,
		ContentLength: -1,
	}
}

func Mock429ResponseNoResetHeader() *http.Response {
	loc, _ := time.LoadLocation("UTC")
	zulu := time.Now().In(loc)
	header := http.Header{}
	header.Add("X-Okta-Now", strconv.FormatInt(zulu.Unix(), 10))
	header.Add("X-Okta-Request-id", "a-request-id")
	header.Add("Content-Type", "application/json")
	header.Add("Date", zulu.Format("Mon, 02 Jan 2006 15:04:05 GMT"))

	return &http.Response{
		Status:        strconv.Itoa(429),
		StatusCode:    429,
		Body:          httpmock.NewRespBodyFromString("{}"),
		Header:        header,
		ContentLength: -1,
	}
}

func Mock429ResponseNoDateHeader() *http.Response {
	loc, _ := time.LoadLocation("UTC")
	zulu := time.Now().In(loc)
	header := http.Header{}
	header.Add("X-Okta-Now", strconv.FormatInt(zulu.Unix(), 10))
	header.Add("X-Rate-Limit-Reset", strconv.FormatInt(time.Now().Unix()+1, 10))
	header.Add("X-Okta-Request-id", "a-request-id")
	header.Add("Content-Type", "application/json")

	return &http.Response{
		Status:        strconv.Itoa(429),
		StatusCode:    429,
		Body:          httpmock.NewRespBodyFromString("{}"),
		Header:        header,
		ContentLength: -1,
	}
}

func Mock429ResponseMultipleHeaders() *http.Response {
	loc, _ := time.LoadLocation("UTC")
	zulu := time.Now().In(loc)
	header := http.Header{}
	header.Add("X-Okta-Now", strconv.FormatInt(zulu.Unix(), 10))
	header.Add("X-Rate-Limit-Reset", strconv.FormatInt(time.Now().Unix()+20, 10))
	header.Add("X-Rate-Limit-Reset", strconv.FormatInt(time.Now().Unix()+10, 10))
	header.Add("X-Okta-Request-id", "a-request-id")
	header.Add("Content-Type", "application/json")
	header.Add("Date", zulu.Format("Mon, 02 Jan 2006 15:04:05 GMT"))

	return &http.Response{
		Status:        strconv.Itoa(429),
		StatusCode:    429,
		Body:          httpmock.NewRespBodyFromString("{}"),
		Header:        header,
		ContentLength: -1,
	}
}

func MockValidResponse() *http.Response {
	header := http.Header{}
	header.Add("X-Okta-Request-id", "another-request-id")
	header.Add("Content-Type", "application/json")
	header.Add("Date", time.Now().Add(time.Second*10).Format(time.RFC3339))

	return &http.Response{
		Status:        strconv.Itoa(200),
		StatusCode:    200,
		Body:          httpmock.NewRespBodyFromString("[]"),
		Header:        header,
		ContentLength: -1,
	}
}

func MockSessionCreateResponse() *http.Response {
	header := http.Header{}
	header.Add("X-Okta-Request-id", "another-request-id")
	header.Add("Content-Type", "application/json")
	header.Add("Accept", "application/json")
	header.Add("Date", time.Now().Add(time.Second*10).Format(time.RFC3339))

	return &http.Response{
		Status:     strconv.Itoa(200),
		StatusCode: 200,
		Body: httpmock.NewRespBodyFromString(`{
			"id": "101W_juydrDRByB7fUdRyE2JQ",
			"login": "user@example.com",
			"userId": "00ubgaSARVOQDIOXMORI",
			"expiresAt": "2015-08-30T18:41:35.818Z",
			"status": "ACTIVE",
			"lastPasswordVerification": "2015-08-30T18:41:35.818Z",
			"lastFactorVerification": null,
			"amr": [
			  "pwd"
			],
			"idp": {
			  "id": "00oi5cpnylv792IcF0g3",
			  "type": "OKTA"
			},
			"mfaActive": false}`),
		Header:        header,
		ContentLength: -1,
	}
}
