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

type SamlApplicationSettingsSignOn struct {
	AssertionSigned bool `json:"assertionSigned,omitempty"`
	AttributeStatements []string `json:"attributeStatements,omitempty"`
	Audience string `json:"audience,omitempty"`
	AudienceOverride string `json:"audienceOverride,omitempty"`
	AuthnContextClassRef string `json:"authnContextClassRef,omitempty"`
	DefaultRelayState string `json:"defaultRelayState,omitempty"`
	Destination string `json:"destination,omitempty"`
	DestinationOverride string `json:"destinationOverride,omitempty"`
	DigestAlgorithm string `json:"digestAlgorithm,omitempty"`
	HonorForceAuthn bool `json:"honorForceAuthn,omitempty"`
	IdpIssuer string `json:"idpIssuer,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	RecipientOverride string `json:"recipientOverride,omitempty"`
	RequestCompressed bool `json:"requestCompressed,omitempty"`
	ResponseSigned bool `json:"responseSigned,omitempty"`
	SignatureAlgorithm string `json:"signatureAlgorithm,omitempty"`
	SpIssuer string `json:"spIssuer,omitempty"`
	SsoAcsUrl string `json:"ssoAcsUrl,omitempty"`
	SsoAcsUrlOverride string `json:"ssoAcsUrlOverride,omitempty"`
	SubjectNameIdFormat string `json:"subjectNameIdFormat,omitempty"`
	SubjectNameIdTemplate string `json:"subjectNameIdTemplate,omitempty"`
}

func (m *SamlApplicationSettingsSignOn) WithAssertionSigned(v bool) *SamlApplicationSettingsSignOn {
	m.AssertionSigned = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithAttributeStatements(v []string) *SamlApplicationSettingsSignOn {
	m.AttributeStatements = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithAudience(v string) *SamlApplicationSettingsSignOn {
	m.Audience = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithAudienceOverride(v string) *SamlApplicationSettingsSignOn {
	m.AudienceOverride = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithAuthnContextClassRef(v string) *SamlApplicationSettingsSignOn {
	m.AuthnContextClassRef = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithDefaultRelayState(v string) *SamlApplicationSettingsSignOn {
	m.DefaultRelayState = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithDestination(v string) *SamlApplicationSettingsSignOn {
	m.Destination = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithDestinationOverride(v string) *SamlApplicationSettingsSignOn {
	m.DestinationOverride = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithDigestAlgorithm(v string) *SamlApplicationSettingsSignOn {
	m.DigestAlgorithm = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithHonorForceAuthn(v bool) *SamlApplicationSettingsSignOn {
	m.HonorForceAuthn = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithIdpIssuer(v string) *SamlApplicationSettingsSignOn {
	m.IdpIssuer = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithRecipient(v string) *SamlApplicationSettingsSignOn {
	m.Recipient = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithRecipientOverride(v string) *SamlApplicationSettingsSignOn {
	m.RecipientOverride = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithRequestCompressed(v bool) *SamlApplicationSettingsSignOn {
	m.RequestCompressed = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithResponseSigned(v bool) *SamlApplicationSettingsSignOn {
	m.ResponseSigned = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithSignatureAlgorithm(v string) *SamlApplicationSettingsSignOn {
	m.SignatureAlgorithm = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithSpIssuer(v string) *SamlApplicationSettingsSignOn {
	m.SpIssuer = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithSsoAcsUrl(v string) *SamlApplicationSettingsSignOn {
	m.SsoAcsUrl = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithSsoAcsUrlOverride(v string) *SamlApplicationSettingsSignOn {
	m.SsoAcsUrlOverride = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithSubjectNameIdFormat(v string) *SamlApplicationSettingsSignOn {
	m.SubjectNameIdFormat = v
	return m
}

func (m *SamlApplicationSettingsSignOn) WithSubjectNameIdTemplate(v string) *SamlApplicationSettingsSignOn {
	m.SubjectNameIdTemplate = v
	return m
}


