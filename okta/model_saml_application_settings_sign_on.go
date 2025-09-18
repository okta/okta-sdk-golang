/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the SamlApplicationSettingsSignOn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlApplicationSettingsSignOn{}

// SamlApplicationSettingsSignOn SAML 2.0 sign-on attributes. > **Note:** Set either `destinationOverride` or `ssoAcsUrl` to configure any other SAML 2.0 attributes in this section.
type SamlApplicationSettingsSignOn struct {
	// An array of ACS endpoints. You can configure a maximum of 100 endpoints.
	AcsEndpoints []AcsEndpoint `json:"acsEndpoints,omitempty"`
	// Determines whether the app allows you to configure multiple ACS URIs
	AllowMultipleAcsEndpoints bool                     `json:"allowMultipleAcsEndpoints"`
	AssertionEncryption       *SamlAssertionEncryption `json:"assertionEncryption,omitempty"`
	// Determines whether the SAML assertion is digitally signed
	AssertionSigned bool `json:"assertionSigned"`
	// A list of custom attribute statements for the app's SAML assertion. See [SAML 2.0 Technical Overview](https://docs.oasis-open.org/security/saml/Post2.0/sstc-saml-tech-overview-2.0-cd-02.html).  There are two types of attribute statements: | Type | Description | | ---- | ----------- | | EXPRESSION | Generic attribute statement that can be dynamic and supports [Okta Expression Language](https://developer.okta.com/docs/reference/okta-expression-language/) | | GROUP | Group attribute statement |
	AttributeStatements []SamlAttributeStatement `json:"attributeStatements,omitempty"`
	// The entity ID of the SP. Use the entity ID value exactly as provided by the SP.
	Audience string `json:"audience"`
	// Audience override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	AudienceOverride *string `json:"audienceOverride,omitempty"`
	// Identifies the SAML authentication context class for the assertion's authentication statement
	AuthnContextClassRef string `json:"authnContextClassRef"`
	// The list of dynamic attribute statements for the SAML assertion inherited from app metadata (apps from the OIN) during app creation.  There are two types of attribute statements: `EXPRESSION` and `GROUP`.
	ConfiguredAttributeStatements []SamlAttributeStatement `json:"configuredAttributeStatements,omitempty"`
	// Identifies a specific application resource in an IdP-initiated SSO scenario
	DefaultRelayState *string `json:"defaultRelayState,omitempty"`
	// Identifies the location inside the SAML assertion where the SAML response should be sent
	Destination string `json:"destination"`
	// Destination override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	DestinationOverride *string `json:"destinationOverride,omitempty"`
	// Determines the digest algorithm used to digitally sign the SAML assertion and response
	DigestAlgorithm string `json:"digestAlgorithm"`
	// Set to `true` to prompt users for their credentials when a SAML request has the `ForceAuthn` attribute set to `true`
	HonorForceAuthn bool `json:"honorForceAuthn"`
	// SAML Issuer ID
	IdpIssuer string `json:"idpIssuer"`
	// Associates the app with SAML inline hooks. See [the SAML assertion inline hook reference](https://developer.okta.com/docs/reference/saml-hook/).
	InlineHooks    []SignOnInlineHook `json:"inlineHooks,omitempty"`
	ParticipateSlo *SloParticipate    `json:"participateSlo,omitempty"`
	// The location where the app may present the SAML assertion
	Recipient string `json:"recipient"`
	// Recipient override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	RecipientOverride *string `json:"recipientOverride,omitempty"`
	// Determines whether the SAML request is expected to be compressed
	RequestCompressed bool `json:"requestCompressed"`
	// Determines whether the SAML authentication response message is digitally signed by the IdP > **Note:** Either (or both) `responseSigned` or `assertionSigned` must be `TRUE`.
	ResponseSigned bool `json:"responseSigned"`
	// Determines the SAML app session lifetimes with Okta
	SamlAssertionLifetimeSeconds *int32 `json:"samlAssertionLifetimeSeconds,omitempty"`
	// Determines the signing algorithm used to digitally sign the SAML assertion and response
	SignatureAlgorithm string             `json:"signatureAlgorithm"`
	Slo                *SingleLogout      `json:"slo,omitempty"`
	SpCertificate      *SamlSpCertificate `json:"spCertificate,omitempty"`
	// The issuer ID for the Service Provider. This property appears when SLO is enabled.
	SpIssuer *string `json:"spIssuer,omitempty"`
	// Single Sign-On Assertion Consumer Service (ACS) URL
	SsoAcsUrl string `json:"ssoAcsUrl"`
	// Assertion Consumer Service (ACS) URL override for CASB configuration. See [CASB config guide](https://help.okta.com/en-us/Content/Topics/Apps/CASB-config-guide.htm).
	SsoAcsUrlOverride *string `json:"ssoAcsUrlOverride,omitempty"`
	// Identifies the SAML processing rules. Supported values:
	SubjectNameIdFormat string `json:"subjectNameIdFormat"`
	// Template for app user's username when a user is assigned to the app
	SubjectNameIdTemplate string `json:"subjectNameIdTemplate"`
	AdditionalProperties  map[string]interface{}
}

type _SamlApplicationSettingsSignOn SamlApplicationSettingsSignOn

// NewSamlApplicationSettingsSignOn instantiates a new SamlApplicationSettingsSignOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlApplicationSettingsSignOn(allowMultipleAcsEndpoints bool, assertionSigned bool, audience string, authnContextClassRef string, destination string, digestAlgorithm string, honorForceAuthn bool, idpIssuer string, recipient string, requestCompressed bool, responseSigned bool, signatureAlgorithm string, ssoAcsUrl string, subjectNameIdFormat string, subjectNameIdTemplate string) *SamlApplicationSettingsSignOn {
	this := SamlApplicationSettingsSignOn{}
	this.AllowMultipleAcsEndpoints = allowMultipleAcsEndpoints
	this.AssertionSigned = assertionSigned
	this.Audience = audience
	this.AuthnContextClassRef = authnContextClassRef
	this.Destination = destination
	this.DigestAlgorithm = digestAlgorithm
	this.HonorForceAuthn = honorForceAuthn
	this.IdpIssuer = idpIssuer
	this.Recipient = recipient
	this.RequestCompressed = requestCompressed
	this.ResponseSigned = responseSigned
	this.SignatureAlgorithm = signatureAlgorithm
	this.SsoAcsUrl = ssoAcsUrl
	this.SubjectNameIdFormat = subjectNameIdFormat
	this.SubjectNameIdTemplate = subjectNameIdTemplate
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
	if o == nil || IsNil(o.AcsEndpoints) {
		var ret []AcsEndpoint
		return ret
	}
	return o.AcsEndpoints
}

// GetAcsEndpointsOk returns a tuple with the AcsEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAcsEndpointsOk() ([]AcsEndpoint, bool) {
	if o == nil || IsNil(o.AcsEndpoints) {
		return nil, false
	}
	return o.AcsEndpoints, true
}

// HasAcsEndpoints returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAcsEndpoints() bool {
	if o != nil && !IsNil(o.AcsEndpoints) {
		return true
	}

	return false
}

// SetAcsEndpoints gets a reference to the given []AcsEndpoint and assigns it to the AcsEndpoints field.
func (o *SamlApplicationSettingsSignOn) SetAcsEndpoints(v []AcsEndpoint) {
	o.AcsEndpoints = v
}

// GetAllowMultipleAcsEndpoints returns the AllowMultipleAcsEndpoints field value
func (o *SamlApplicationSettingsSignOn) GetAllowMultipleAcsEndpoints() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowMultipleAcsEndpoints
}

// GetAllowMultipleAcsEndpointsOk returns a tuple with the AllowMultipleAcsEndpoints field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAllowMultipleAcsEndpointsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowMultipleAcsEndpoints, true
}

// SetAllowMultipleAcsEndpoints sets field value
func (o *SamlApplicationSettingsSignOn) SetAllowMultipleAcsEndpoints(v bool) {
	o.AllowMultipleAcsEndpoints = v
}

// GetAssertionEncryption returns the AssertionEncryption field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAssertionEncryption() SamlAssertionEncryption {
	if o == nil || IsNil(o.AssertionEncryption) {
		var ret SamlAssertionEncryption
		return ret
	}
	return *o.AssertionEncryption
}

// GetAssertionEncryptionOk returns a tuple with the AssertionEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAssertionEncryptionOk() (*SamlAssertionEncryption, bool) {
	if o == nil || IsNil(o.AssertionEncryption) {
		return nil, false
	}
	return o.AssertionEncryption, true
}

// HasAssertionEncryption returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAssertionEncryption() bool {
	if o != nil && !IsNil(o.AssertionEncryption) {
		return true
	}

	return false
}

// SetAssertionEncryption gets a reference to the given SamlAssertionEncryption and assigns it to the AssertionEncryption field.
func (o *SamlApplicationSettingsSignOn) SetAssertionEncryption(v SamlAssertionEncryption) {
	o.AssertionEncryption = &v
}

// GetAssertionSigned returns the AssertionSigned field value
func (o *SamlApplicationSettingsSignOn) GetAssertionSigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AssertionSigned
}

// GetAssertionSignedOk returns a tuple with the AssertionSigned field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAssertionSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssertionSigned, true
}

// SetAssertionSigned sets field value
func (o *SamlApplicationSettingsSignOn) SetAssertionSigned(v bool) {
	o.AssertionSigned = v
}

// GetAttributeStatements returns the AttributeStatements field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAttributeStatements() []SamlAttributeStatement {
	if o == nil || IsNil(o.AttributeStatements) {
		var ret []SamlAttributeStatement
		return ret
	}
	return o.AttributeStatements
}

// GetAttributeStatementsOk returns a tuple with the AttributeStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAttributeStatementsOk() ([]SamlAttributeStatement, bool) {
	if o == nil || IsNil(o.AttributeStatements) {
		return nil, false
	}
	return o.AttributeStatements, true
}

// HasAttributeStatements returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAttributeStatements() bool {
	if o != nil && !IsNil(o.AttributeStatements) {
		return true
	}

	return false
}

// SetAttributeStatements gets a reference to the given []SamlAttributeStatement and assigns it to the AttributeStatements field.
func (o *SamlApplicationSettingsSignOn) SetAttributeStatements(v []SamlAttributeStatement) {
	o.AttributeStatements = v
}

// GetAudience returns the Audience field value
func (o *SamlApplicationSettingsSignOn) GetAudience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *SamlApplicationSettingsSignOn) SetAudience(v string) {
	o.Audience = v
}

// GetAudienceOverride returns the AudienceOverride field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetAudienceOverride() string {
	if o == nil || IsNil(o.AudienceOverride) {
		var ret string
		return ret
	}
	return *o.AudienceOverride
}

// GetAudienceOverrideOk returns a tuple with the AudienceOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAudienceOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.AudienceOverride) {
		return nil, false
	}
	return o.AudienceOverride, true
}

// HasAudienceOverride returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasAudienceOverride() bool {
	if o != nil && !IsNil(o.AudienceOverride) {
		return true
	}

	return false
}

// SetAudienceOverride gets a reference to the given string and assigns it to the AudienceOverride field.
func (o *SamlApplicationSettingsSignOn) SetAudienceOverride(v string) {
	o.AudienceOverride = &v
}

// GetAuthnContextClassRef returns the AuthnContextClassRef field value
func (o *SamlApplicationSettingsSignOn) GetAuthnContextClassRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthnContextClassRef
}

// GetAuthnContextClassRefOk returns a tuple with the AuthnContextClassRef field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetAuthnContextClassRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthnContextClassRef, true
}

// SetAuthnContextClassRef sets field value
func (o *SamlApplicationSettingsSignOn) SetAuthnContextClassRef(v string) {
	o.AuthnContextClassRef = v
}

// GetConfiguredAttributeStatements returns the ConfiguredAttributeStatements field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetConfiguredAttributeStatements() []SamlAttributeStatement {
	if o == nil || IsNil(o.ConfiguredAttributeStatements) {
		var ret []SamlAttributeStatement
		return ret
	}
	return o.ConfiguredAttributeStatements
}

// GetConfiguredAttributeStatementsOk returns a tuple with the ConfiguredAttributeStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetConfiguredAttributeStatementsOk() ([]SamlAttributeStatement, bool) {
	if o == nil || IsNil(o.ConfiguredAttributeStatements) {
		return nil, false
	}
	return o.ConfiguredAttributeStatements, true
}

// HasConfiguredAttributeStatements returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasConfiguredAttributeStatements() bool {
	if o != nil && !IsNil(o.ConfiguredAttributeStatements) {
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
	if o == nil || IsNil(o.DefaultRelayState) {
		var ret string
		return ret
	}
	return *o.DefaultRelayState
}

// GetDefaultRelayStateOk returns a tuple with the DefaultRelayState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetDefaultRelayStateOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultRelayState) {
		return nil, false
	}
	return o.DefaultRelayState, true
}

// HasDefaultRelayState returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasDefaultRelayState() bool {
	if o != nil && !IsNil(o.DefaultRelayState) {
		return true
	}

	return false
}

// SetDefaultRelayState gets a reference to the given string and assigns it to the DefaultRelayState field.
func (o *SamlApplicationSettingsSignOn) SetDefaultRelayState(v string) {
	o.DefaultRelayState = &v
}

// GetDestination returns the Destination field value
func (o *SamlApplicationSettingsSignOn) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *SamlApplicationSettingsSignOn) SetDestination(v string) {
	o.Destination = v
}

// GetDestinationOverride returns the DestinationOverride field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetDestinationOverride() string {
	if o == nil || IsNil(o.DestinationOverride) {
		var ret string
		return ret
	}
	return *o.DestinationOverride
}

// GetDestinationOverrideOk returns a tuple with the DestinationOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetDestinationOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationOverride) {
		return nil, false
	}
	return o.DestinationOverride, true
}

// HasDestinationOverride returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasDestinationOverride() bool {
	if o != nil && !IsNil(o.DestinationOverride) {
		return true
	}

	return false
}

// SetDestinationOverride gets a reference to the given string and assigns it to the DestinationOverride field.
func (o *SamlApplicationSettingsSignOn) SetDestinationOverride(v string) {
	o.DestinationOverride = &v
}

// GetDigestAlgorithm returns the DigestAlgorithm field value
func (o *SamlApplicationSettingsSignOn) GetDigestAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetDigestAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DigestAlgorithm, true
}

// SetDigestAlgorithm sets field value
func (o *SamlApplicationSettingsSignOn) SetDigestAlgorithm(v string) {
	o.DigestAlgorithm = v
}

// GetHonorForceAuthn returns the HonorForceAuthn field value
func (o *SamlApplicationSettingsSignOn) GetHonorForceAuthn() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HonorForceAuthn
}

// GetHonorForceAuthnOk returns a tuple with the HonorForceAuthn field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetHonorForceAuthnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HonorForceAuthn, true
}

// SetHonorForceAuthn sets field value
func (o *SamlApplicationSettingsSignOn) SetHonorForceAuthn(v bool) {
	o.HonorForceAuthn = v
}

// GetIdpIssuer returns the IdpIssuer field value
func (o *SamlApplicationSettingsSignOn) GetIdpIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpIssuer
}

// GetIdpIssuerOk returns a tuple with the IdpIssuer field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetIdpIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpIssuer, true
}

// SetIdpIssuer sets field value
func (o *SamlApplicationSettingsSignOn) SetIdpIssuer(v string) {
	o.IdpIssuer = v
}

// GetInlineHooks returns the InlineHooks field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetInlineHooks() []SignOnInlineHook {
	if o == nil || IsNil(o.InlineHooks) {
		var ret []SignOnInlineHook
		return ret
	}
	return o.InlineHooks
}

// GetInlineHooksOk returns a tuple with the InlineHooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetInlineHooksOk() ([]SignOnInlineHook, bool) {
	if o == nil || IsNil(o.InlineHooks) {
		return nil, false
	}
	return o.InlineHooks, true
}

// HasInlineHooks returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasInlineHooks() bool {
	if o != nil && !IsNil(o.InlineHooks) {
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
	if o == nil || IsNil(o.ParticipateSlo) {
		var ret SloParticipate
		return ret
	}
	return *o.ParticipateSlo
}

// GetParticipateSloOk returns a tuple with the ParticipateSlo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetParticipateSloOk() (*SloParticipate, bool) {
	if o == nil || IsNil(o.ParticipateSlo) {
		return nil, false
	}
	return o.ParticipateSlo, true
}

// HasParticipateSlo returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasParticipateSlo() bool {
	if o != nil && !IsNil(o.ParticipateSlo) {
		return true
	}

	return false
}

// SetParticipateSlo gets a reference to the given SloParticipate and assigns it to the ParticipateSlo field.
func (o *SamlApplicationSettingsSignOn) SetParticipateSlo(v SloParticipate) {
	o.ParticipateSlo = &v
}

// GetRecipient returns the Recipient field value
func (o *SamlApplicationSettingsSignOn) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *SamlApplicationSettingsSignOn) SetRecipient(v string) {
	o.Recipient = v
}

// GetRecipientOverride returns the RecipientOverride field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetRecipientOverride() string {
	if o == nil || IsNil(o.RecipientOverride) {
		var ret string
		return ret
	}
	return *o.RecipientOverride
}

// GetRecipientOverrideOk returns a tuple with the RecipientOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetRecipientOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientOverride) {
		return nil, false
	}
	return o.RecipientOverride, true
}

// HasRecipientOverride returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasRecipientOverride() bool {
	if o != nil && !IsNil(o.RecipientOverride) {
		return true
	}

	return false
}

// SetRecipientOverride gets a reference to the given string and assigns it to the RecipientOverride field.
func (o *SamlApplicationSettingsSignOn) SetRecipientOverride(v string) {
	o.RecipientOverride = &v
}

// GetRequestCompressed returns the RequestCompressed field value
func (o *SamlApplicationSettingsSignOn) GetRequestCompressed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequestCompressed
}

// GetRequestCompressedOk returns a tuple with the RequestCompressed field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetRequestCompressedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestCompressed, true
}

// SetRequestCompressed sets field value
func (o *SamlApplicationSettingsSignOn) SetRequestCompressed(v bool) {
	o.RequestCompressed = v
}

// GetResponseSigned returns the ResponseSigned field value
func (o *SamlApplicationSettingsSignOn) GetResponseSigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResponseSigned
}

// GetResponseSignedOk returns a tuple with the ResponseSigned field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetResponseSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseSigned, true
}

// SetResponseSigned sets field value
func (o *SamlApplicationSettingsSignOn) SetResponseSigned(v bool) {
	o.ResponseSigned = v
}

// GetSamlAssertionLifetimeSeconds returns the SamlAssertionLifetimeSeconds field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSamlAssertionLifetimeSeconds() int32 {
	if o == nil || IsNil(o.SamlAssertionLifetimeSeconds) {
		var ret int32
		return ret
	}
	return *o.SamlAssertionLifetimeSeconds
}

// GetSamlAssertionLifetimeSecondsOk returns a tuple with the SamlAssertionLifetimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSamlAssertionLifetimeSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.SamlAssertionLifetimeSeconds) {
		return nil, false
	}
	return o.SamlAssertionLifetimeSeconds, true
}

// HasSamlAssertionLifetimeSeconds returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSamlAssertionLifetimeSeconds() bool {
	if o != nil && !IsNil(o.SamlAssertionLifetimeSeconds) {
		return true
	}

	return false
}

// SetSamlAssertionLifetimeSeconds gets a reference to the given int32 and assigns it to the SamlAssertionLifetimeSeconds field.
func (o *SamlApplicationSettingsSignOn) SetSamlAssertionLifetimeSeconds(v int32) {
	o.SamlAssertionLifetimeSeconds = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value
func (o *SamlApplicationSettingsSignOn) GetSignatureAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureAlgorithm, true
}

// SetSignatureAlgorithm sets field value
func (o *SamlApplicationSettingsSignOn) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = v
}

// GetSlo returns the Slo field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSlo() SingleLogout {
	if o == nil || IsNil(o.Slo) {
		var ret SingleLogout
		return ret
	}
	return *o.Slo
}

// GetSloOk returns a tuple with the Slo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSloOk() (*SingleLogout, bool) {
	if o == nil || IsNil(o.Slo) {
		return nil, false
	}
	return o.Slo, true
}

// HasSlo returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSlo() bool {
	if o != nil && !IsNil(o.Slo) {
		return true
	}

	return false
}

// SetSlo gets a reference to the given SingleLogout and assigns it to the Slo field.
func (o *SamlApplicationSettingsSignOn) SetSlo(v SingleLogout) {
	o.Slo = &v
}

// GetSpCertificate returns the SpCertificate field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSpCertificate() SamlSpCertificate {
	if o == nil || IsNil(o.SpCertificate) {
		var ret SamlSpCertificate
		return ret
	}
	return *o.SpCertificate
}

// GetSpCertificateOk returns a tuple with the SpCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSpCertificateOk() (*SamlSpCertificate, bool) {
	if o == nil || IsNil(o.SpCertificate) {
		return nil, false
	}
	return o.SpCertificate, true
}

// HasSpCertificate returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSpCertificate() bool {
	if o != nil && !IsNil(o.SpCertificate) {
		return true
	}

	return false
}

// SetSpCertificate gets a reference to the given SamlSpCertificate and assigns it to the SpCertificate field.
func (o *SamlApplicationSettingsSignOn) SetSpCertificate(v SamlSpCertificate) {
	o.SpCertificate = &v
}

// GetSpIssuer returns the SpIssuer field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSpIssuer() string {
	if o == nil || IsNil(o.SpIssuer) {
		var ret string
		return ret
	}
	return *o.SpIssuer
}

// GetSpIssuerOk returns a tuple with the SpIssuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSpIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.SpIssuer) {
		return nil, false
	}
	return o.SpIssuer, true
}

// HasSpIssuer returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSpIssuer() bool {
	if o != nil && !IsNil(o.SpIssuer) {
		return true
	}

	return false
}

// SetSpIssuer gets a reference to the given string and assigns it to the SpIssuer field.
func (o *SamlApplicationSettingsSignOn) SetSpIssuer(v string) {
	o.SpIssuer = &v
}

// GetSsoAcsUrl returns the SsoAcsUrl field value
func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SsoAcsUrl
}

// GetSsoAcsUrlOk returns a tuple with the SsoAcsUrl field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SsoAcsUrl, true
}

// SetSsoAcsUrl sets field value
func (o *SamlApplicationSettingsSignOn) SetSsoAcsUrl(v string) {
	o.SsoAcsUrl = v
}

// GetSsoAcsUrlOverride returns the SsoAcsUrlOverride field value if set, zero value otherwise.
func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrlOverride() string {
	if o == nil || IsNil(o.SsoAcsUrlOverride) {
		var ret string
		return ret
	}
	return *o.SsoAcsUrlOverride
}

// GetSsoAcsUrlOverrideOk returns a tuple with the SsoAcsUrlOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSsoAcsUrlOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.SsoAcsUrlOverride) {
		return nil, false
	}
	return o.SsoAcsUrlOverride, true
}

// HasSsoAcsUrlOverride returns a boolean if a field has been set.
func (o *SamlApplicationSettingsSignOn) HasSsoAcsUrlOverride() bool {
	if o != nil && !IsNil(o.SsoAcsUrlOverride) {
		return true
	}

	return false
}

// SetSsoAcsUrlOverride gets a reference to the given string and assigns it to the SsoAcsUrlOverride field.
func (o *SamlApplicationSettingsSignOn) SetSsoAcsUrlOverride(v string) {
	o.SsoAcsUrlOverride = &v
}

// GetSubjectNameIdFormat returns the SubjectNameIdFormat field value
func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectNameIdFormat
}

// GetSubjectNameIdFormatOk returns a tuple with the SubjectNameIdFormat field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectNameIdFormat, true
}

// SetSubjectNameIdFormat sets field value
func (o *SamlApplicationSettingsSignOn) SetSubjectNameIdFormat(v string) {
	o.SubjectNameIdFormat = v
}

// GetSubjectNameIdTemplate returns the SubjectNameIdTemplate field value
func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectNameIdTemplate
}

// GetSubjectNameIdTemplateOk returns a tuple with the SubjectNameIdTemplate field value
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettingsSignOn) GetSubjectNameIdTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectNameIdTemplate, true
}

// SetSubjectNameIdTemplate sets field value
func (o *SamlApplicationSettingsSignOn) SetSubjectNameIdTemplate(v string) {
	o.SubjectNameIdTemplate = v
}

func (o SamlApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlApplicationSettingsSignOn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcsEndpoints) {
		toSerialize["acsEndpoints"] = o.AcsEndpoints
	}
	toSerialize["allowMultipleAcsEndpoints"] = o.AllowMultipleAcsEndpoints
	if !IsNil(o.AssertionEncryption) {
		toSerialize["assertionEncryption"] = o.AssertionEncryption
	}
	toSerialize["assertionSigned"] = o.AssertionSigned
	if !IsNil(o.AttributeStatements) {
		toSerialize["attributeStatements"] = o.AttributeStatements
	}
	toSerialize["audience"] = o.Audience
	if !IsNil(o.AudienceOverride) {
		toSerialize["audienceOverride"] = o.AudienceOverride
	}
	toSerialize["authnContextClassRef"] = o.AuthnContextClassRef
	if !IsNil(o.ConfiguredAttributeStatements) {
		toSerialize["configuredAttributeStatements"] = o.ConfiguredAttributeStatements
	}
	if !IsNil(o.DefaultRelayState) {
		toSerialize["defaultRelayState"] = o.DefaultRelayState
	}
	toSerialize["destination"] = o.Destination
	if !IsNil(o.DestinationOverride) {
		toSerialize["destinationOverride"] = o.DestinationOverride
	}
	toSerialize["digestAlgorithm"] = o.DigestAlgorithm
	toSerialize["honorForceAuthn"] = o.HonorForceAuthn
	toSerialize["idpIssuer"] = o.IdpIssuer
	if !IsNil(o.InlineHooks) {
		toSerialize["inlineHooks"] = o.InlineHooks
	}
	if !IsNil(o.ParticipateSlo) {
		toSerialize["participateSlo"] = o.ParticipateSlo
	}
	toSerialize["recipient"] = o.Recipient
	if !IsNil(o.RecipientOverride) {
		toSerialize["recipientOverride"] = o.RecipientOverride
	}
	toSerialize["requestCompressed"] = o.RequestCompressed
	toSerialize["responseSigned"] = o.ResponseSigned
	if !IsNil(o.SamlAssertionLifetimeSeconds) {
		toSerialize["samlAssertionLifetimeSeconds"] = o.SamlAssertionLifetimeSeconds
	}
	toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	if !IsNil(o.Slo) {
		toSerialize["slo"] = o.Slo
	}
	if !IsNil(o.SpCertificate) {
		toSerialize["spCertificate"] = o.SpCertificate
	}
	if !IsNil(o.SpIssuer) {
		toSerialize["spIssuer"] = o.SpIssuer
	}
	toSerialize["ssoAcsUrl"] = o.SsoAcsUrl
	if !IsNil(o.SsoAcsUrlOverride) {
		toSerialize["ssoAcsUrlOverride"] = o.SsoAcsUrlOverride
	}
	toSerialize["subjectNameIdFormat"] = o.SubjectNameIdFormat
	toSerialize["subjectNameIdTemplate"] = o.SubjectNameIdTemplate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlApplicationSettingsSignOn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowMultipleAcsEndpoints",
		"assertionSigned",
		"audience",
		"authnContextClassRef",
		"destination",
		"digestAlgorithm",
		"honorForceAuthn",
		"idpIssuer",
		"recipient",
		"requestCompressed",
		"responseSigned",
		"signatureAlgorithm",
		"ssoAcsUrl",
		"subjectNameIdFormat",
		"subjectNameIdTemplate",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSamlApplicationSettingsSignOn := _SamlApplicationSettingsSignOn{}

	err = json.Unmarshal(data, &varSamlApplicationSettingsSignOn)

	if err != nil {
		return err
	}

	*o = SamlApplicationSettingsSignOn(varSamlApplicationSettingsSignOn)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "acsEndpoints")
		delete(additionalProperties, "allowMultipleAcsEndpoints")
		delete(additionalProperties, "assertionEncryption")
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
