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
	"testing"

	"github.com/okta/okta-sdk-golang/okta"
	"github.com/stretchr/testify/assert"
)

func Test_WithOrgUrl_removes_backslash(t *testing.T) {
	config := okta.NewConfig().WithOrgUrl("http://test.b.com/")
	_ = okta.NewClient(config, nil, nil)
	assert.Equal(t, "http://test.b.com", config.Okta.Client.OrgUrl)
}
func Test_WithOrgUrl_handles_properly_formatted_url(t *testing.T) {
	config := okta.NewConfig().WithOrgUrl("http://test.b.com")
	_ = okta.NewClient(config, nil, nil)
	assert.Equal(t, "http://test.b.com", config.Okta.Client.OrgUrl)
}
func Test_WithOrgUrl_handles_empty_string(t *testing.T) {
	config := okta.NewConfig().WithOrgUrl("")
	_ = okta.NewClient(config, nil, nil)
	assert.Equal(t, "", config.Okta.Client.OrgUrl)
}
