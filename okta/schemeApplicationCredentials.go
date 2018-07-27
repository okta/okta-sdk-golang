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

type SchemeApplicationCredentials struct {
	Password *PasswordCredential `json:"password,omitempty"`
	RevealPassword bool `json:"revealPassword,omitempty"`
	Scheme string `json:"scheme,omitempty"`
	Signing *ApplicationCredentialsSigning `json:"signing,omitempty"`
	UserName string `json:"userName,omitempty"`
}

func (m *SchemeApplicationCredentials) WithPassword(v *PasswordCredential) *SchemeApplicationCredentials {
	m.Password = v
	return m
}

func (m *SchemeApplicationCredentials) WithRevealPassword(v bool) *SchemeApplicationCredentials {
	m.RevealPassword = v
	return m
}

func (m *SchemeApplicationCredentials) WithScheme(v string) *SchemeApplicationCredentials {
	m.Scheme = v
	return m
}

func (m *SchemeApplicationCredentials) WithSigning(v *ApplicationCredentialsSigning) *SchemeApplicationCredentials {
	m.Signing = v
	return m
}

func (m *SchemeApplicationCredentials) WithUserName(v string) *SchemeApplicationCredentials {
	m.UserName = v
	return m
}


