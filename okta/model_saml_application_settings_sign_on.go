/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// SamlApplicationSettingsSignOn SAML sign-on attributes. > **Note:** Only for SAML 2.0, set either `destinationOverride` or `ssoAcsUrl` to configure any other SAML 2.0 attributes in this section.
type SamlApplicationSettingsSignOn struct {
	AcsEndpoints []AcsEndpoint `json:"acsEndpoints,omitempty"`
	AllowMultipleAcsEndpoints *bool `json:"allowMultipleAcsEndpoints,omitempty"`
	AssertionSigned *bool `json:"assertionSigned,omitempty"`
	AttributeStatements []SamlAttributeStatement `json:"attributeStatements,omitempty"`
	Audience *string `json:"audience,omitempty"`
	// Audience override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	AudienceOverride *string `json:"audienceOverride,omitempty"`
	AuthnContextClassRef *string `json:"authnContextClassRef,omitempty"`
	ConfiguredAttributeStatements []SamlAttributeStatement `json:"configuredAttributeStatements,omitempty"`
	// Identifies a specific application resource in an IdP-initiated SSO scenario
	DefaultRelayState *string `json:"defaultRelayState,omitempty"`
	Destination *string `json:"destination,omitempty"`
	// Destination override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	DestinationOverride *string `json:"destinationOverride,omitempty"`
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty"`
	// Set to `true` to prompt users for their credentials when a SAML request has the `ForceAuthn` attribute set to `true`
	HonorForceAuthn *bool `json:"honorForceAuthn,omitempty"`
	IdpIssuer *string `json:"idpIssuer,omitempty"`
	InlineHooks []SignOnInlineHook `json:"inlineHooks,omitempty"`
	ParticipateSlo *SloParticipate `json:"participateSlo,omitempty"`
	Recipient *string `json:"recipient,omitempty"`
	// Recipient override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	RecipientOverride *string `json:"recipientOverride,omitempty"`
	RequestCompressed *bool `json:"requestCompressed,omitempty"`
	ResponseSigned *bool `json:"responseSigned,omitempty"`
	// For SAML 2.0 only.<br>Determines the SAML app session lifetimes with Okta
	SamlAssertionLifetimeSeconds *int32 `json:"samlAssertionLifetimeSeconds,omitempty"`
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
	Slo *SingleLogout `json:"slo,omitempty"`
	SpCertificate *SpCertificate `json:"spCertificate,omitempty"`
	SpIssuer *string `json:"spIssuer,omitempty"`
	// Single Sign-On Assertion Consumer Service (ACS) URL
	SsoAcsUrl *string `json:"ssoAcsUrl,omitempty"`
	// Assertion Consumer Service (ACS) URL override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	SsoAcsUrlOverride *string `json:"ssoAcsUrlOverride,omitempty"`
	SubjectNameIdFormat *string `json:"subjectNameIdFormat,omitempty"`
	SubjectNameIdTemplate *string `json:"subjectNameIdTemplate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlApplicationSettingsSignOn SamlApplicationSettingsSignOn

// NewSamlApplicationSettingsSignOn instantiates a new SamlApplicationSettingsSignOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlApplicationSettingsSignOn() *SamlApplicationSettingsSignOn {
	this := SamlApplicationSettingsSignOn{}
	return &this
}

// NewSamlApplicationSettingsSignOnWithDefaults instantiates a new SamlApplicationSettingsSignOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlApplicationSettingsSignOnWithDefaults() *SamlApplicationSettingsSignOn {
	this := SamlApplicationSettingsSignOn{}
	return &this
}

// GetAcsEndpoints returns the AcsEndpoints field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAcsEndpoints() []AcsEndpoint {
	if o == nil || o.AcsEndpoints == nil {
		var ret []AcsEndpoint
		return ret
	}
	return o.AcsEndpoints
}

// GetAcsEndpointsOk returns a tuple with the AcsEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAcsEndpointsOk() ([]AcsEndpoint, bool) {
	if o == nil || o.AcsEndpoints == nil {
		return nil, false
	}
	return o.AcsEndpoints, true
}

// HasAcsEndpoints returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAcsEndpoints() bool {
	if o != nil && o.AcsEndpoints != nil {
		return true
	}

	return false
}

// SetAcsEndpoints gets a reference to the given []AcsEndpoint and assigns it to the AcsEndpoints field.
func (o *SamlApplicationSettingsSignOn) SetAcsEndpoints(v []AcsEndpoint) {
	o.AcsEndpoints = v
}

// GetAllowMultipleAcsEndpoints returns the AllowMultipleAcsEndpoints field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAllowMultipleAcsEndpoints() bool {
	if o == nil || o.AllowMultipleAcsEndpoints == nil {
		var ret bool
		return ret
	}
	return *o.AllowMultipleAcsEndpoints
}

// GetAllowMultipleAcsEndpointsOk returns a tuple with the AllowMultipleAcsEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAllowMultipleAcsEndpointsOk() (*bool, bool) {
	if o == nil || o.AllowMultipleAcsEndpoints == nil {
		return nil, false
	}
	return o.AllowMultipleAcsEndpoints, true
}

// HasAllowMultipleAcsEndpoints returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAllowMultipleAcsEndpoints() bool {
	if o != nil && o.AllowMultipleAcsEndpoints != nil {
		return true
	}

	return false
}

// SetAllowMultipleAcsEndpoints gets a reference to the given bool and assigns it to the AllowMultipleAcsEndpoints field.
func (o *SamlApplicationSettingsSignOn) SetAllowMultipleAcsEndpoints(v bool) {
	o.AllowMultipleAcsEndpoints = &v
}

// GetAssertionSigned returns the AssertionSigned field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAssertionSigned() bool {
	if o == nil || o.AssertionSigned == nil {
		var ret bool
		return ret
	}
	return *o.AssertionSigned
}

// GetAssertionSignedOk returns a tuple with the AssertionSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAssertionSignedOk() (*bool, bool) {
	if o == nil || o.AssertionSigned == nil {
		return nil, false
	}
	return o.AssertionSigned, true
}

// HasAssertionSigned returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAssertionSigned() bool {
	if o != nil && o.AssertionSigned != nil {
		return true
	}

	return false
}

// SetAssertionSigned gets a reference to the given bool and assigns it to the AssertionSigned field.
func (o *SamlApplicationSettingsSignOn) SetAssertionSigned(v bool) {
	o.AssertionSigned = &v
}

// GetAttributeStatements returns the AttributeStatements field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAttributeStatements() []SamlAttributeStatement {
	if o == nil || o.AttributeStatements == nil {
		var ret []SamlAttributeStatement
		return ret
	}
	return o.AttributeStatements
}

// GetAttributeStatementsOk returns a tuple with the AttributeStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAttributeStatementsOk() ([]SamlAttributeStatement, bool) {
	if o == nil || o.AttributeStatements == nil {
		return nil, false
	}
	return o.AttributeStatements, true
}

// HasAttributeStatements returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAttributeStatements() bool {
	if o != nil && o.AttributeStatements != nil {
		return true
	}

	return false
}

// SetAttributeStatements gets a reference to the given []SamlAttributeStatement and assigns it to the AttributeStatements field.
func (o *SamlApplicationSettingsSignOn) SetAttributeStatements(v []SamlAttributeStatement) {
	o.AttributeStatements = v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *SamlApplicationSettingsSignOn) SetAudience(v string) {
	o.Audience = &v
}

// GetAudienceOverride returns the AudienceOverride field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAudienceOverride() string {
	if o == nil || o.AudienceOverride == nil {
		var ret string
		return ret
	}
	return *o.AudienceOverride
}

// GetAudienceOverrideOk returns a tuple with the AudienceOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAudienceOverrideOk() (*string, bool) {
	if o == nil || o.AudienceOverride == nil {
		return nil, false
	}
	return o.AudienceOverride, true
}

// HasAudienceOverride returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAudienceOverride() bool {
	if o != nil && o.AudienceOverride != nil {
		return true
	}

	return false
}

// SetAudienceOverride gets a reference to the given string and assigns it to the AudienceOverride field.
func (o *SamlApplicationSettingsSignOn) SetAudienceOverride(v string) {
	o.AudienceOverride = &v
}

// GetAuthnContextClassRef returns the AuthnContextClassRef field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAuthnContextClassRef() string {
	if o == nil || o.AuthnContextClassRef == nil {
		var ret string
		return ret
	}
	return *o.AuthnContextClassRef
}

// GetAuthnContextClassRefOk returns a tuple with the AuthnContextClassRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAuthnContextClassRefOk() (*string, bool) {
	if o == nil || o.AuthnContextClassRef == nil {
		return nil, false
	}
	return o.AuthnContextClassRef, true
}

// HasAuthnContextClassRef returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAuthnContextClassRef() bool {
	if o != nil && o.AuthnContextClassRef != nil {
		return true
	}

	return false
}

// SetAuthnContextClassRef gets a reference to the given string and assigns it to the AuthnContextClassRef field.
func (o *SamlApplicationSettingsSignOn) SetAuthnContextClassRef(v string) {
	o.AuthnContextClassRef = &v
}

// GetConfiguredAttributeStatements returns the ConfiguredAttributeStatements field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetConfiguredAttributeStatements() []SamlAttributeStatement {
	if o == nil || o.ConfiguredAttributeStatements == nil {
		var ret []SamlAttributeStatement
		return ret
	}
	return o.ConfiguredAttributeStatements
}

// GetConfiguredAttributeStatementsOk returns a tuple with the ConfiguredAttributeStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetConfiguredAttributeStatementsOk() ([]SamlAttributeStatement, bool) {
	if o == nil || o.ConfiguredAttributeStatements == nil {
		return nil, false
	}
	return o.ConfiguredAttributeStatements, true
}

// HasConfiguredAttributeStatements returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasConfiguredAttributeStatements() bool {
	if o != nil && o.ConfiguredAttributeStatements != nil {
		return true
	}

	return false
}

// SetConfiguredAttributeStatements gets a reference to the given []SamlAttributeStatement and assigns it to the ConfiguredAttributeStatements field.
func (o *SamlApplicationSettingsSignOn) SetConfiguredAttributeStatements(v []SamlAttributeStatement) {
	o.ConfiguredAttributeStatements = v
}

// GetDefaultRelayState returns the DefaultRelayState field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetDefaultRelayState() string {
	if o == nil || o.DefaultRelayState == nil {
		var ret string
		return ret
	}
	return *o.DefaultRelayState
}

// GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetDefaultRelayStateOk() (*string, bool) {
	if o == nil || o.DefaultRelayState == nil {
		return nil, false
	}
	return o.DefaultRelayState, true
}

// HasDefaultRelayState returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasDefaultRelayState() bool {
	if o != nil && o.DefaultRelayState != nil {
		return true
	}

	return false
}

// SetDefaultRelayState gets a reference to the given string and assigns it to the DefaultRelayState field.
func (o *SamlApplicationSettingsSignOn) SetDefaultRelayState(v string) {
	o.DefaultRelayState = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetDestination() string {
	if o == nil || o.Destination == nil {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetDestinationOk() (*string, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *SamlApplicationSettingsSignOn) SetDestination(v string) {
	o.Destination = &v
}

// GetDestinationOverride returns the DestinationOverride field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetDestinationOverride() string {
	if o == nil || o.DestinationOverride == nil {
		var ret string
		return ret
	}
	return *o.DestinationOverride
}

// GetDestinationOverrideOk returns a tuple with the DestinationOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetDestinationOverrideOk() (*string, bool) {
	if o == nil || o.DestinationOverride == nil {
		return nil, false
	}
	return o.DestinationOverride, true
}

// HasDestinationOverride returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasDestinationOverride() bool {
	if o != nil && o.DestinationOverride != nil {
		return true
	}

	return false
}

// SetDestinationOverride gets a reference to the given string and assigns it to the DestinationOverride field.
func (o *SamlApplicationSettingsSignOn) SetDestinationOverride(v string) {
	o.DestinationOverride = &v
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetDigestAlgorithm() string {
	if o == nil || o.DigestAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetDigestAlgorithmOk() (*string, bool) {
	if o == nil || o.DigestAlgorithm == nil {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasDigestAlgorithm() bool {
	if o != nil && o.DigestAlgorithm != nil {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given string and assigns it to the DigestAlgorithm field.
func (o *SamlApplicationSettingsSignOn) SetDigestAlgorithm(v string) {
	o.DigestAlgorithm = &v
}

// GetHonorForceAuthn returns the HonorForceAuthn field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetHonorForceAuthn() bool {
	if o == nil || o.HonorForceAuthn == nil {
		var ret bool
		return ret
	}
	return *o.HonorForceAuthn
}

// GetHonorForceAuthnOk returns a tuple with the HonorForceAuthn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetHonorForceAuthnOk() (*bool, bool) {
	if o == nil || o.HonorForceAuthn == nil {
		return nil, false
	}
	return o.HonorForceAuthn, true
}

// HasHonorForceAuthn returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasHonorForceAuthn() bool {
	if o != nil && o.HonorForceAuthn != nil {
		return true
	}

	return false
}

// SetHonorForceAuthn gets a reference to the given bool and assigns it to the HonorForceAuthn field.
func (o *SamlApplicationSettingsSignOn) SetHonorForceAuthn(v bool) {
	o.HonorForceAuthn = &v
}

// GetIdpIssuer returns the IdpIssuer field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetIdpIssuer() string {
	if o == nil || o.IdpIssuer == nil {
		var ret string
		return ret
	}
	return *o.IdpIssuer
}

// GetIdpIssuerOk returns a tuple with the IdpIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetIdpIssuerOk() (*string, bool) {
	if o == nil || o.IdpIssuer == nil {
		return nil, false
	}
	return o.IdpIssuer, true
}

// HasIdpIssuer returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasIdpIssuer() bool {
	if o != nil && o.IdpIssuer != nil {
		return true
	}

	return false
}

// SetIdpIssuer gets a reference to the given string and assigns it to the IdpIssuer field.
func (o *SamlApplicationSettingsSignOn) SetIdpIssuer(v string) {
	o.IdpIssuer = &v
}

// GetInlineHooks returns the InlineHooks field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetInlineHooks() []SignOnInlineHook {
	if o == nil || o.InlineHooks == nil {
		var ret []SignOnInlineHook
		return ret
	}
	return o.InlineHooks
}

// GetInlineHooksOk returns a tuple with the InlineHooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetInlineHooksOk() ([]SignOnInlineHook, bool) {
	if o == nil || o.InlineHooks == nil {
		return nil, false
	}
	return o.InlineHooks, true
}

// HasInlineHooks returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasInlineHooks() bool {
	if o != nil && o.InlineHooks != nil {
		return true
	}

	return false
}

// SetInlineHooks gets a reference to the given []SignOnInlineHook and assigns it to the InlineHooks field.
func (o *SamlApplicationSettingsSignOn) SetInlineHooks(v []SignOnInlineHook) {
	o.InlineHooks = v
}

// GetParticipateSlo returns the ParticipateSlo field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetParticipateSlo() SloParticipate {
	if o == nil || o.ParticipateSlo == nil {
		var ret SloParticipate
		return ret
	}
	return *o.ParticipateSlo
}

// GetParticipateSloOk returns a tuple with the ParticipateSlo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetParticipateSloOk() (*SloParticipate, bool) {
	if o == nil || o.ParticipateSlo == nil {
		return nil, false
	}
	return o.ParticipateSlo, true
}

// HasParticipateSlo returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasParticipateSlo() bool {
	if o != nil && o.ParticipateSlo != nil {
		return true
	}

	return false
}

// SetParticipateSlo gets a reference to the given SloParticipate and assigns it to the ParticipateSlo field.
func (o *SamlApplicationSettingsSignOn) SetParticipateSlo(v SloParticipate) {
	o.ParticipateSlo = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetRecipient() string {
	if o == nil || o.Recipient == nil {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetRecipientOk() (*string, bool) {
	if o == nil || o.Recipient == nil {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasRecipient() bool {
	if o != nil && o.Recipient != nil {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *SamlApplicationSettingsSignOn) SetRecipient(v string) {
	o.Recipient = &v
}

// GetRecipientOverride returns the RecipientOverride field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetRecipientOverride() string {
	if o == nil || o.RecipientOverride == nil {
		var ret string
		return ret
	}
	return *o.RecipientOverride
}

// GetRecipientOverrideOk returns a tuple with the RecipientOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetRecipientOverrideOk() (*string, bool) {
	if o == nil || o.RecipientOverride == nil {
		return nil, false
	}
	return o.RecipientOverride, true
}

// HasRecipientOverride returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasRecipientOverride() bool {
	if o != nil && o.RecipientOverride != nil {
		return true
	}

	return false
}

// SetRecipientOverride gets a reference to the given string and assigns it to the RecipientOverride field.
func (o *SamlApplicationSettingsSignOn) SetRecipientOverride(v string) {
	o.RecipientOverride = &v
}

// GetRequestCompressed returns the RequestCompressed field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetRequestCompressed() bool {
	if o == nil || o.RequestCompressed == nil {
		var ret bool
		return ret
	}
	return *o.RequestCompressed
}

// GetRequestCompressedOk returns a tuple with the RequestCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetRequestCompressedOk() (*bool, bool) {
	if o == nil || o.RequestCompressed == nil {
		return nil, false
	}
	return o.RequestCompressed, true
}

// HasRequestCompressed returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasRequestCompressed() bool {
	if o != nil && o.RequestCompressed != nil {
		return true
	}

	return false
}

// SetRequestCompressed gets a reference to the given bool and assigns it to the RequestCompressed field.
func (o *SamlApplicationSettingsSignOn) SetRequestCompressed(v bool) {
	o.RequestCompressed = &v
}

// GetResponseSigned returns the ResponseSigned field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetResponseSigned() bool {
	if o == nil || o.ResponseSigned == nil {
		var ret bool
		return ret
	}
	return *o.ResponseSigned
}

// GetResponseSignedOk returns a tuple with the ResponseSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetResponseSignedOk() (*bool, bool) {
	if o == nil || o.ResponseSigned == nil {
		return nil, false
	}
	return o.ResponseSigned, true
}

// HasResponseSigned returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasResponseSigned() bool {
	if o != nil && o.ResponseSigned != nil {
		return true
	}

	return false
}

// SetResponseSigned gets a reference to the given bool and assigns it to the ResponseSigned field.
func (o *SamlApplicationSettingsSignOn) SetResponseSigned(v bool) {
	o.ResponseSigned = &v
}

// GetSamlAssertionLifetimeSeconds returns the SamlAssertionLifetimeSeconds field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSamlAssertionLifetimeSeconds() int32 {
	if o == nil || o.SamlAssertionLifetimeSeconds == nil {
		var ret int32
		return ret
	}
	return *o.SamlAssertionLifetimeSeconds
}

// GetSamlAssertionLifetimeSecondsOk returns a tuple with the SamlAssertionLifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSamlAssertionLifetimeSecondsOk() (*int32, bool) {
	if o == nil || o.SamlAssertionLifetimeSeconds == nil {
		return nil, false
	}
	return o.SamlAssertionLifetimeSeconds, true
}

// HasSamlAssertionLifetimeSeconds returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSamlAssertionLifetimeSeconds() bool {
	if o != nil && o.SamlAssertionLifetimeSeconds != nil {
		return true
	}

	return false
}

// SetSamlAssertionLifetimeSeconds gets a reference to the given int32 and assigns it to the SamlAssertionLifetimeSeconds field.
func (o *SamlApplicationSettingsSignOn) SetSamlAssertionLifetimeSeconds(v int32) {
	o.SamlAssertionLifetimeSeconds = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSignatureAlgorithm() string {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *SamlApplicationSettingsSignOn) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetSlo returns the Slo field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSlo() SingleLogout {
	if o == nil || o.Slo == nil {
		var ret SingleLogout
		return ret
	}
	return *o.Slo
}

// GetSloOk returns a tuple with the Slo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSloOk() (*SingleLogout, bool) {
	if o == nil || o.Slo == nil {
		return nil, false
	}
	return o.Slo, true
}

// HasSlo returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSlo() bool {
	if o != nil && o.Slo != nil {
		return true
	}

	return false
}

// SetSlo gets a reference to the given SingleLogout and assigns it to the Slo field.
func (o *SamlApplicationSettingsSignOn) SetSlo(v SingleLogout) {
	o.Slo = &v
}

// GetSpCertificate returns the SpCertificate field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSpCertificate() SpCertificate {
	if o == nil || o.SpCertificate == nil {
		var ret SpCertificate
		return ret
	}
	return *o.SpCertificate
}

// GetSpCertificateOk returns a tuple with the SpCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSpCertificateOk() (*SpCertificate, bool) {
	if o == nil || o.SpCertificate == nil {
		return nil, false
	}
	return o.SpCertificate, true
}

// HasSpCertificate returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSpCertificate() bool {
	if o != nil && o.SpCertificate != nil {
		return true
	}

	return false
}

// SetSpCertificate gets a reference to the given SpCertificate and assigns it to the SpCertificate field.
func (o *SamlApplicationSettingsSignOn) SetSpCertificate(v SpCertificate) {
	o.SpCertificate = &v
}

// GetSpIssuer returns the SpIssuer field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSpIssuer() string {
	if o == nil || o.SpIssuer == nil {
		var ret string
		return ret
	}
	return *o.SpIssuer
}

// GetSpIssuerOk returns a tuple with the SpIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSpIssuerOk() (*string, bool) {
	if o == nil || o.SpIssuer == nil {
		return nil, false
	}
	return o.SpIssuer, true
}

// HasSpIssuer returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSpIssuer() bool {
	if o != nil && o.SpIssuer != nil {
		return true
	}

	return false
}

// SetSpIssuer gets a reference to the given string and assigns it to the SpIssuer field.
func (o *SamlApplicationSettingsSignOn) SetSpIssuer(v string) {
	o.SpIssuer = &v
}

// GetSsoAcsUrl returns the SsoAcsUrl field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrl() string {
	if o == nil || o.SsoAcsUrl == nil {
		var ret string
		return ret
	}
	return *o.SsoAcsUrl
}

// GetSsoAcsUrlOk returns a tuple with the SsoAcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrlOk() (*string, bool) {
	if o == nil || o.SsoAcsUrl == nil {
		return nil, false
	}
	return o.SsoAcsUrl, true
}

// HasSsoAcsUrl returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSsoAcsUrl() bool {
	if o != nil && o.SsoAcsUrl != nil {
		return true
	}

	return false
}

// SetSsoAcsUrl gets a reference to the given string and assigns it to the SsoAcsUrl field.
func (o *SamlApplicationSettingsSignOn) SetSsoAcsUrl(v string) {
	o.SsoAcsUrl = &v
}

// GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrlOverride() string {
	if o == nil || o.SsoAcsUrlOverride == nil {
		var ret string
		return ret
	}
	return *o.SsoAcsUrlOverride
}

// GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrlOverrideOk() (*string, bool) {
	if o == nil || o.SsoAcsUrlOverride == nil {
		return nil, false
	}
	return o.SsoAcsUrlOverride, true
}

// HasSsoAcsUrlOverride returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSsoAcsUrlOverride() bool {
	if o != nil && o.SsoAcsUrlOverride != nil {
		return true
	}

	return false
}

// SetSsoAcsUrlOverride gets a reference to the given string and assigns it to the SsoAcsUrlOverride field.
func (o *SamlApplicationSettingsSignOn) SetSsoAcsUrlOverride(v string) {
	o.SsoAcsUrlOverride = &v
}

// GetSubjectNameIdFormat returns the SubjectNameIdFormat field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdFormat() string {
	if o == nil || o.SubjectNameIdFormat == nil {
		var ret string
		return ret
	}
	return *o.SubjectNameIdFormat
}

// GetSubjectNameIdFormatOk returns a tuple with the SubjectNameIdFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdFormatOk() (*string, bool) {
	if o == nil || o.SubjectNameIdFormat == nil {
		return nil, false
	}
	return o.SubjectNameIdFormat, true
}

// HasSubjectNameIdFormat returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSubjectNameIdFormat() bool {
	if o != nil && o.SubjectNameIdFormat != nil {
		return true
	}

	return false
}

// SetSubjectNameIdFormat gets a reference to the given string and assigns it to the SubjectNameIdFormat field.
func (o *SamlApplicationSettingsSignOn) SetSubjectNameIdFormat(v string) {
	o.SubjectNameIdFormat = &v
}

// GetSubjectNameIdTemplate returns the SubjectNameIdTemplate field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdTemplate() string {
	if o == nil || o.SubjectNameIdTemplate == nil {
		var ret string
		return ret
	}
	return *o.SubjectNameIdTemplate
}

// GetSubjectNameIdTemplateOk returns a tuple with the SubjectNameIdTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdTemplateOk() (*string, bool) {
	if o == nil || o.SubjectNameIdTemplate == nil {
		return nil, false
	}
	return o.SubjectNameIdTemplate, true
}

// HasSubjectNameIdTemplate returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSubjectNameIdTemplate() bool {
	if o != nil && o.SubjectNameIdTemplate != nil {
		return true
	}

	return false
}

// SetSubjectNameIdTemplate gets a reference to the given string and assigns it to the SubjectNameIdTemplate field.
func (o *SamlApplicationSettingsSignOn) SetSubjectNameIdTemplate(v string) {
	o.SubjectNameIdTemplate = &v
}

func (o SamlApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcsEndpoints != nil {
		toSerialize["acsEndpoints"] = o.AcsEndpoints
	}
	if o.AllowMultipleAcsEndpoints != nil {
		toSerialize["allowMultipleAcsEndpoints"] = o.AllowMultipleAcsEndpoints
	}
	if o.AssertionSigned != nil {
		toSerialize["assertionSigned"] = o.AssertionSigned
	}
	if o.AttributeStatements != nil {
		toSerialize["attributeStatements"] = o.AttributeStatements
	}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.AudienceOverride != nil {
		toSerialize["audienceOverride"] = o.AudienceOverride
	}
	if o.AuthnContextClassRef != nil {
		toSerialize["authnContextClassRef"] = o.AuthnContextClassRef
	}
	if o.ConfiguredAttributeStatements != nil {
		toSerialize["configuredAttributeStatements"] = o.ConfiguredAttributeStatements
	}
	if o.DefaultRelayState != nil {
		toSerialize["defaultRelayState"] = o.DefaultRelayState
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.DestinationOverride != nil {
		toSerialize["destinationOverride"] = o.DestinationOverride
	}
	if o.DigestAlgorithm != nil {
		toSerialize["digestAlgorithm"] = o.DigestAlgorithm
	}
	if o.HonorForceAuthn != nil {
		toSerialize["honorForceAuthn"] = o.HonorForceAuthn
	}
	if o.IdpIssuer != nil {
		toSerialize["idpIssuer"] = o.IdpIssuer
	}
	if o.InlineHooks != nil {
		toSerialize["inlineHooks"] = o.InlineHooks
	}
	if o.ParticipateSlo != nil {
		toSerialize["participateSlo"] = o.ParticipateSlo
	}
	if o.Recipient != nil {
		toSerialize["recipient"] = o.Recipient
	}
	if o.RecipientOverride != nil {
		toSerialize["recipientOverride"] = o.RecipientOverride
	}
	if o.RequestCompressed != nil {
		toSerialize["requestCompressed"] = o.RequestCompressed
	}
	if o.ResponseSigned != nil {
		toSerialize["responseSigned"] = o.ResponseSigned
	}
	if o.SamlAssertionLifetimeSeconds != nil {
		toSerialize["samlAssertionLifetimeSeconds"] = o.SamlAssertionLifetimeSeconds
	}
	if o.SignatureAlgorithm != nil {
		toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	}
	if o.Slo != nil {
		toSerialize["slo"] = o.Slo
	}
	if o.SpCertificate != nil {
		toSerialize["spCertificate"] = o.SpCertificate
	}
	if o.SpIssuer != nil {
		toSerialize["spIssuer"] = o.SpIssuer
	}
	if o.SsoAcsUrl != nil {
		toSerialize["ssoAcsUrl"] = o.SsoAcsUrl
	}
	if o.SsoAcsUrlOverride != nil {
		toSerialize["ssoAcsUrlOverride"] = o.SsoAcsUrlOverride
	}
	if o.SubjectNameIdFormat != nil {
		toSerialize["subjectNameIdFormat"] = o.SubjectNameIdFormat
	}
	if o.SubjectNameIdTemplate != nil {
		toSerialize["subjectNameIdTemplate"] = o.SubjectNameIdTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlApplicationSettingsSignOn) UnmarshalJSON(bytes []byte) (err error) {
	varSamlApplicationSettingsSignOn := _SamlApplicationSettingsSignOn{}

	err = json.Unmarshal(bytes, &varSamlApplicationSettingsSignOn)
	if err == nil {
		*o = SamlApplicationSettingsSignOn(varSamlApplicationSettingsSignOn)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "acsEndpoints")
		delete(additionalProperties, "allowMultipleAcsEndpoints")
		delete(additionalProperties, "assertionSigned")
		delete(additionalProperties, "attributeStatements")
		delete(additionalProperties, "audience")
		delete(additionalProperties, "audienceOverride")
		delete(additionalProperties, "authnContextClassRef")
		delete(additionalProperties, "configuredAttributeStatements")
		delete(additionalProperties, "defaultRelayState")
		delete(additionalProperties, "destination")
		delete(additionalProperties, "destinationOverride")
		delete(additionalProperties, "digestAlgorithm")
		delete(additionalProperties, "honorForceAuthn")
		delete(additionalProperties, "idpIssuer")
		delete(additionalProperties, "inlineHooks")
		delete(additionalProperties, "participateSlo")
		delete(additionalProperties, "recipient")
		delete(additionalProperties, "recipientOverride")
		delete(additionalProperties, "requestCompressed")
		delete(additionalProperties, "responseSigned")
		delete(additionalProperties, "samlAssertionLifetimeSeconds")
		delete(additionalProperties, "signatureAlgorithm")
		delete(additionalProperties, "slo")
		delete(additionalProperties, "spCertificate")
		delete(additionalProperties, "spIssuer")
		delete(additionalProperties, "ssoAcsUrl")
		delete(additionalProperties, "ssoAcsUrlOverride")
		delete(additionalProperties, "subjectNameIdFormat")
		delete(additionalProperties, "subjectNameIdTemplate")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlApplicationSettingsSignOn struct {
	value *SamlApplicationSettingsSignOn
	isSet bool
}

func (v NullableSamlApplicationSettingsSignOn) Get() *SamlApplicationSettingsSignOn {
	return v.value
}

func (v *NullableSamlApplicationSettingsSignOn) Set(val *SamlApplicationSettingsSignOn) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlApplicationSettingsSignOn) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlApplicationSettingsSignOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlApplicationSettingsSignOn(val *SamlApplicationSettingsSignOn) *NullableSamlApplicationSettingsSignOn {
	return &NullableSamlApplicationSettingsSignOn{value: val, isSet: true}
}

func (v NullableSamlApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlApplicationSettingsSignOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

