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

type WsFederationApplicationSettingsApplication struct {
	AttributeStatements string `json:"attributeStatements,omitempty"`
	AudienceRestriction string `json:"audienceRestriction,omitempty"`
	AuthnContextClassRef string `json:"authnContextClassRef,omitempty"`
	GroupFilter string `json:"groupFilter,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	GroupValueFormat string `json:"groupValueFormat,omitempty"`
	NameIDFormat string `json:"nameIDFormat,omitempty"`
	Realm string `json:"realm,omitempty"`
	SiteURL string `json:"siteURL,omitempty"`
	UsernameAttribute string `json:"usernameAttribute,omitempty"`
	WReplyOverride bool `json:"wReplyOverride,omitempty"`
	WReplyURL string `json:"wReplyURL,omitempty"`
}

func (m *WsFederationApplicationSettingsApplication) WithAttributeStatements(v string) *WsFederationApplicationSettingsApplication {
	m.AttributeStatements = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithAudienceRestriction(v string) *WsFederationApplicationSettingsApplication {
	m.AudienceRestriction = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithAuthnContextClassRef(v string) *WsFederationApplicationSettingsApplication {
	m.AuthnContextClassRef = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithGroupFilter(v string) *WsFederationApplicationSettingsApplication {
	m.GroupFilter = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithGroupName(v string) *WsFederationApplicationSettingsApplication {
	m.GroupName = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithGroupValueFormat(v string) *WsFederationApplicationSettingsApplication {
	m.GroupValueFormat = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithNameIDFormat(v string) *WsFederationApplicationSettingsApplication {
	m.NameIDFormat = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithRealm(v string) *WsFederationApplicationSettingsApplication {
	m.Realm = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithSiteURL(v string) *WsFederationApplicationSettingsApplication {
	m.SiteURL = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithUsernameAttribute(v string) *WsFederationApplicationSettingsApplication {
	m.UsernameAttribute = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithWReplyOverride(v bool) *WsFederationApplicationSettingsApplication {
	m.WReplyOverride = v
	return m
}

func (m *WsFederationApplicationSettingsApplication) WithWReplyURL(v string) *WsFederationApplicationSettingsApplication {
	m.WReplyURL = v
	return m
}


