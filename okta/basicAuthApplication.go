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

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
)

type BasicAuthApplication struct {
	Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
	Name string `json:"name,omitempty"`
	Settings *BasicApplicationSettings `json:"settings,omitempty"`
}

func (m *BasicAuthApplication) WithCredentials(v *SchemeApplicationCredentials) *BasicAuthApplication {
	m.Credentials = v
	return m
}

func (m *BasicAuthApplication) WithName(v string) *BasicAuthApplication {
	m.Name = v
	return m
}

func (m *BasicAuthApplication) WithSettings(v *BasicApplicationSettings) *BasicAuthApplication {
	m.Settings = v
	return m
}


