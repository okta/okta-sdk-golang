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

type UserProfile struct {
	Email string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Login string `json:"login,omitempty"`
	MobilePhone string `json:"mobilePhone,omitempty"`
	SecondEmail string `json:"secondEmail,omitempty"`
}

func (m *UserProfile) WithEmail(v string) *UserProfile {
	m.Email = v
	return m
}

func (m *UserProfile) WithFirstName(v string) *UserProfile {
	m.FirstName = v
	return m
}

func (m *UserProfile) WithLastName(v string) *UserProfile {
	m.LastName = v
	return m
}

func (m *UserProfile) WithLogin(v string) *UserProfile {
	m.Login = v
	return m
}

func (m *UserProfile) WithMobilePhone(v string) *UserProfile {
	m.MobilePhone = v
	return m
}

func (m *UserProfile) WithSecondEmail(v string) *UserProfile {
	m.SecondEmail = v
	return m
}


