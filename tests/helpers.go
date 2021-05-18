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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/v2/okta"
)

func NewClient(ctx context.Context, conf ...okta.ConfigSetter) (context.Context, *okta.Client, error) {
	return okta.NewClient(ctx, conf...)
}

func Assert_response(t *testing.T, response *okta.Response, requestMethod string, requestPath string) {
	require.IsType(t, &okta.Response{}, response, "did not return `*okta.Response` type as second variable")
	assert.Equal(t, requestMethod, response.Response.Request.Method, "did not make a requestMethod request")
	assert.Equal(t, requestPath, response.Response.Request.URL.Path, "path for request was incorrect")
}
