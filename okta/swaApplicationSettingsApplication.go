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

type SwaApplicationSettingsApplication struct {
	ButtonField string `json:"buttonField,omitempty"`
	LoginUrlRegex string `json:"loginUrlRegex,omitempty"`
	PasswordField string `json:"passwordField,omitempty"`
	Url string `json:"url,omitempty"`
	UsernameField string `json:"usernameField,omitempty"`
}

func (m *SwaApplicationSettingsApplication) WithButtonField(v string) *SwaApplicationSettingsApplication {
	m.ButtonField = v
	return m
}

func (m *SwaApplicationSettingsApplication) WithLoginUrlRegex(v string) *SwaApplicationSettingsApplication {
	m.LoginUrlRegex = v
	return m
}

func (m *SwaApplicationSettingsApplication) WithPasswordField(v string) *SwaApplicationSettingsApplication {
	m.PasswordField = v
	return m
}

func (m *SwaApplicationSettingsApplication) WithUrl(v string) *SwaApplicationSettingsApplication {
	m.Url = v
	return m
}

func (m *SwaApplicationSettingsApplication) WithUsernameField(v string) *SwaApplicationSettingsApplication {
	m.UsernameField = v
	return m
}


