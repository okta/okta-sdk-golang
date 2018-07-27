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

type SwaThreeFieldApplicationSettingsApplication struct {
	ButtonSelector string `json:"buttonSelector,omitempty"`
	ExtraFieldSelector string `json:"extraFieldSelector,omitempty"`
	ExtraFieldValue string `json:"extraFieldValue,omitempty"`
	LoginUrlRegex string `json:"loginUrlRegex,omitempty"`
	PasswordSelector string `json:"passwordSelector,omitempty"`
	TargetUrl string `json:"targetUrl,omitempty"`
	UserNameSelector string `json:"userNameSelector,omitempty"`
}

func (m *SwaThreeFieldApplicationSettingsApplication) WithButtonSelector(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.ButtonSelector = v
	return m
}

func (m *SwaThreeFieldApplicationSettingsApplication) WithExtraFieldSelector(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.ExtraFieldSelector = v
	return m
}

func (m *SwaThreeFieldApplicationSettingsApplication) WithExtraFieldValue(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.ExtraFieldValue = v
	return m
}

func (m *SwaThreeFieldApplicationSettingsApplication) WithLoginUrlRegex(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.LoginUrlRegex = v
	return m
}

func (m *SwaThreeFieldApplicationSettingsApplication) WithPasswordSelector(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.PasswordSelector = v
	return m
}

func (m *SwaThreeFieldApplicationSettingsApplication) WithTargetUrl(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.TargetUrl = v
	return m
}

func (m *SwaThreeFieldApplicationSettingsApplication) WithUserNameSelector(v string) *SwaThreeFieldApplicationSettingsApplication {
	m.UserNameSelector = v
	return m
}


