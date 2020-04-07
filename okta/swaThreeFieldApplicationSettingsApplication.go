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

// Code generated by okta-sdk-generator, DO NOT EDIT.

package okta

import ()

type SwaThreeFieldApplicationSettingsApplication struct {
	ButtonSelector     string `json:"buttonSelector,omitempty"`
	ExtraFieldSelector string `json:"extraFieldSelector,omitempty"`
	ExtraFieldValue    string `json:"extraFieldValue,omitempty"`
	LoginUrlRegex      string `json:"loginUrlRegex,omitempty"`
	PasswordSelector   string `json:"passwordSelector,omitempty"`
	TargetUrl          string `json:"targetUrl,omitempty"`
	UserNameSelector   string `json:"userNameSelector,omitempty"`
}

func NewSwaThreeFieldApplicationSettingsApplication() *SwaThreeFieldApplicationSettingsApplication {
	return &SwaThreeFieldApplicationSettingsApplication{}
}

func (a *SwaThreeFieldApplicationSettingsApplication) IsApplicationInstance() bool {
	return true
}
