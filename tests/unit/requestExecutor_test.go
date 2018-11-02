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

package unit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/okta/okta-sdk-golang/okta/cache"
)

func doRetryTest(t *testing.T, maxRetries int32, backoffEnabled bool, forceError bool) (*okta.Response, error) {
	backoffCount := 0
	okta.Backoff = func(t time.Duration) {
		backoffCount++
	}
	retries := 0
	var testServer *httptest.Server
	handler := http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if retries < 3 {
			if forceError {
				testServer.CloseClientConnections()
			} else {
				res.WriteHeader(http.StatusTooManyRequests)
			}
		} else {
			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(map[string]string{"success": "yes"})
		}
		retries++
	})
	testServer = httptest.NewServer(handler)
	defer testServer.Close()
	config := &okta.Config{
		MaxRetries:     maxRetries,
		BackoffEnabled: backoffEnabled,
	}
	var body interface{}
	re := okta.NewRequestExecutor(&http.Client{}, cache.NoOpCache{}, config)
	req, err := re.NewRequest("GET", testServer.URL, nil)
	assert.Nil(t, err)
	res, err := re.Do(req, &body)

	if backoffEnabled {
		assert.Equal(t, maxRetries, int32(backoffCount))
	} else {
		assert.Equal(t, 0, backoffCount)
	}

	return res, err
}

func Test_backoff_with_retry(t *testing.T) {
	res, err := doRetryTest(t, 3, true, false)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_retries_without_backoff(t *testing.T) {
	res, err := doRetryTest(t, 3, false, true)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func Test_default_no_backoff_or_retries(t *testing.T) {
	res, err := doRetryTest(t, 0, false, false)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusTooManyRequests, res.StatusCode)
}
