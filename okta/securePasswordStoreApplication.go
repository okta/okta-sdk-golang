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

type SecurePasswordStoreApplication struct {
	Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
	Name string `json:"name,omitempty"`
	Settings *SecurePasswordStoreApplicationSettings `json:"settings,omitempty"`
}

func (m *SecurePasswordStoreApplication) WithCredentials(v *SchemeApplicationCredentials) *SecurePasswordStoreApplication {
	m.Credentials = v
	return m
}

func (m *SecurePasswordStoreApplication) WithName(v string) *SecurePasswordStoreApplication {
	m.Name = v
	return m
}

func (m *SecurePasswordStoreApplication) WithSettings(v *SecurePasswordStoreApplicationSettings) *SecurePasswordStoreApplication {
	m.Settings = v
	return m
}


