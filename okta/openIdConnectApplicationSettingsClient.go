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

type OpenIdConnectApplicationSettingsClient struct {
	ApplicationType string `json:"application_type,omitempty"`
	ClientUri string `json:"client_uri,omitempty"`
	ConsentMethod string `json:"consent_method,omitempty"`
	GrantTypes []string `json:"grant_types,omitempty"`
	LogoUri string `json:"logo_uri,omitempty"`
	PolicyUri string `json:"policy_uri,omitempty"`
	RedirectUris []string `json:"redirect_uris,omitempty"`
	ResponseTypes []string `json:"response_types,omitempty"`
	TosUri string `json:"tos_uri,omitempty"`
}

func (m *OpenIdConnectApplicationSettingsClient) WithApplicationType(v string) *OpenIdConnectApplicationSettingsClient {
	m.ApplicationType = v
	return m
}

func (m *OpenIdConnectApplicationSettingsClient) WithClientUri(v string) *OpenIdConnectApplicationSettingsClient {
	m.ClientUri = v
	return m
}

func (m *OpenIdConnectApplicationSettingsClient) WithConsentMethod(v string) *OpenIdConnectApplicationSettingsClient {
	m.ConsentMethod = v
	return m
}

func (m *OpenIdConnectApplicationSettingsClient) WithGrantTypes(v []string) *OpenIdConnectApplicationSettingsClient {
	m.GrantTypes = v
	return m
}

func (m *OpenIdConnectApplicationSettingsClient) WithLogoUri(v string) *OpenIdConnectApplicationSettingsClient {
	m.LogoUri = v
	return m
}

func (m *OpenIdConnectApplicationSettingsClient) WithPolicyUri(v string) *OpenIdConnectApplicationSettingsClient {
	m.PolicyUri = v
	return m
}

func (m *OpenIdConnectApplicationSettingsClient) WithRedirectUris(v []string) *OpenIdConnectApplicationSettingsClient {
	m.RedirectUris = v
	return m
}

func (m *OpenIdConnectApplicationSettingsClient) WithResponseTypes(v []string) *OpenIdConnectApplicationSettingsClient {
	m.ResponseTypes = v
	return m
}

func (m *OpenIdConnectApplicationSettingsClient) WithTosUri(v string) *OpenIdConnectApplicationSettingsClient {
	m.TosUri = v
	return m
}


