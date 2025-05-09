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

// OAuthMetadata struct for OAuthMetadata
type OAuthMetadata struct {
	// URL of the authorization server's authorization endpoint.
	AuthorizationEndpoint *string `json:"authorization_endpoint,omitempty"`
	// <div class=\"x-lifecycle-container\"><x-lifecycle class=\"lea\"></x-lifecycle> <x-lifecycle class=\"oie\"></x-lifecycle></div>A list of signing algorithms that this authorization server supports for signed requests.
	BackchannelAuthenticationRequestSigningAlgValuesSupported []string `json:"backchannel_authentication_request_signing_alg_values_supported,omitempty"`
	// <div class=\"x-lifecycle-container\"><x-lifecycle class=\"lea\"></x-lifecycle> <x-lifecycle class=\"oie\"></x-lifecycle></div>The delivery modes that this authorization server supports for Client-Initiated Backchannel Authentication.
	BackchannelTokenDeliveryModesSupported []string `json:"backchannel_token_delivery_modes_supported,omitempty"`
	// A list of the claims supported by this authorization server.
	ClaimsSupported []string `json:"claims_supported,omitempty"`
	// A list of PKCE code challenge methods supported by this authorization server.
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported,omitempty"`
	DeviceAuthorizationEndpoint *string `json:"device_authorization_endpoint,omitempty"`
	// A list of signing algorithms supported by this authorization server for Demonstrating Proof-of-Possession (DPoP) JWTs.
	DpopSigningAlgValuesSupported []string `json:"dpop_signing_alg_values_supported,omitempty"`
	// URL of the authorization server's logout endpoint.
	EndSessionEndpoint *string `json:"end_session_endpoint,omitempty"`
	// A list of the grant type values that this authorization server supports.
	GrantTypesSupported []string `json:"grant_types_supported,omitempty"`
	// URL of the authorization server's introspection endpoint.
	IntrospectionEndpoint *string `json:"introspection_endpoint,omitempty"`
	// A list of client authentication methods supported by this introspection endpoint.
	IntrospectionEndpointAuthMethodsSupported []string `json:"introspection_endpoint_auth_methods_supported,omitempty"`
	// The authorization server's issuer identifier. In the context of this document, this is your authorization server's base URL. This becomes the `iss` claim in an access token.
	Issuer *string `json:"issuer,omitempty"`
	// URL of the authorization server's JSON Web Key Set document.
	JwksUri *string `json:"jwks_uri,omitempty"`
	PushedAuthorizationRequestEndpoint *string `json:"pushed_authorization_request_endpoint,omitempty"`
	// URL of the authorization server's JSON Web Key Set document.
	RegistrationEndpoint *string `json:"registration_endpoint,omitempty"`
	// A list of signing algorithms that this authorization server supports for signed requests.
	RequestObjectSigningAlgValuesSupported []string `json:"request_object_signing_alg_values_supported,omitempty"`
	// Indicates if Request Parameters are supported by this authorization server.
	RequestParameterSupported *bool `json:"request_parameter_supported,omitempty"`
	// A list of the `response_mode` values that this authorization server supports. More information here.
	ResponseModesSupported []string `json:"response_modes_supported,omitempty"`
	// A list of the `response_type` values that this authorization server supports. Can be a combination of `code`, `token`, and `id_token`.
	ResponseTypesSupported []string `json:"response_types_supported,omitempty"`
	// URL of the authorization server's revocation endpoint.
	RevocationEndpoint *string `json:"revocation_endpoint,omitempty"`
	// A list of client authentication methods supported by this revocation endpoint.
	RevocationEndpointAuthMethodsSupported []string `json:"revocation_endpoint_auth_methods_supported,omitempty"`
	// A list of the scope values that this authorization server supports.
	ScopesSupported []string `json:"scopes_supported,omitempty"`
	// A list of the Subject Identifier types that this authorization server supports. Valid types include `pairwise` and `public`, but only `public` is currently supported. See the [Subject Identifier Types](https://openid.net/specs/openid-connect-core-1_0.html#SubjectIDTypes) section in the OpenID Connect specification.
	SubjectTypesSupported []string `json:"subject_types_supported,omitempty"`
	// URL of the authorization server's token endpoint.
	TokenEndpoint *string `json:"token_endpoint,omitempty"`
	// A list of client authentication methods supported by this token endpoint.
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuthMetadata OAuthMetadata

// NewOAuthMetadata instantiates a new OAuthMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthMetadata() *OAuthMetadata {
	this := OAuthMetadata{}
	return &this
}

// NewOAuthMetadataWithDefaults instantiates a new OAuthMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthMetadataWithDefaults() *OAuthMetadata {
	this := OAuthMetadata{}
	return &this
}

// GetAuthorizationEndpoint returns the AuthorizationEndpoint field value if set, zero value otherwise.
func (o *OAuthMetadata) GetAuthorizationEndpoint() string {
	if o == nil || o.AuthorizationEndpoint == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationEndpoint
}

// GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetAuthorizationEndpointOk() (*string, bool) {
	if o == nil || o.AuthorizationEndpoint == nil {
		return nil, false
	}
	return o.AuthorizationEndpoint, true
}

// HasAuthorizationEndpoint returns a boolean if a field has been set.
func (o *OAuthMetadata) HasAuthorizationEndpoint() bool {
	if o != nil && o.AuthorizationEndpoint != nil {
		return true
	}

	return false
}

// SetAuthorizationEndpoint gets a reference to the given string and assigns it to the AuthorizationEndpoint field.
func (o *OAuthMetadata) SetAuthorizationEndpoint(v string) {
	o.AuthorizationEndpoint = &v
}

// GetBackchannelAuthenticationRequestSigningAlgValuesSupported returns the BackchannelAuthenticationRequestSigningAlgValuesSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetBackchannelAuthenticationRequestSigningAlgValuesSupported() []string {
	if o == nil || o.BackchannelAuthenticationRequestSigningAlgValuesSupported == nil {
		var ret []string
		return ret
	}
	return o.BackchannelAuthenticationRequestSigningAlgValuesSupported
}

// GetBackchannelAuthenticationRequestSigningAlgValuesSupportedOk returns a tuple with the BackchannelAuthenticationRequestSigningAlgValuesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetBackchannelAuthenticationRequestSigningAlgValuesSupportedOk() ([]string, bool) {
	if o == nil || o.BackchannelAuthenticationRequestSigningAlgValuesSupported == nil {
		return nil, false
	}
	return o.BackchannelAuthenticationRequestSigningAlgValuesSupported, true
}

// HasBackchannelAuthenticationRequestSigningAlgValuesSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasBackchannelAuthenticationRequestSigningAlgValuesSupported() bool {
	if o != nil && o.BackchannelAuthenticationRequestSigningAlgValuesSupported != nil {
		return true
	}

	return false
}

// SetBackchannelAuthenticationRequestSigningAlgValuesSupported gets a reference to the given []string and assigns it to the BackchannelAuthenticationRequestSigningAlgValuesSupported field.
func (o *OAuthMetadata) SetBackchannelAuthenticationRequestSigningAlgValuesSupported(v []string) {
	o.BackchannelAuthenticationRequestSigningAlgValuesSupported = v
}

// GetBackchannelTokenDeliveryModesSupported returns the BackchannelTokenDeliveryModesSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetBackchannelTokenDeliveryModesSupported() []string {
	if o == nil || o.BackchannelTokenDeliveryModesSupported == nil {
		var ret []string
		return ret
	}
	return o.BackchannelTokenDeliveryModesSupported
}

// GetBackchannelTokenDeliveryModesSupportedOk returns a tuple with the BackchannelTokenDeliveryModesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetBackchannelTokenDeliveryModesSupportedOk() ([]string, bool) {
	if o == nil || o.BackchannelTokenDeliveryModesSupported == nil {
		return nil, false
	}
	return o.BackchannelTokenDeliveryModesSupported, true
}

// HasBackchannelTokenDeliveryModesSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasBackchannelTokenDeliveryModesSupported() bool {
	if o != nil && o.BackchannelTokenDeliveryModesSupported != nil {
		return true
	}

	return false
}

// SetBackchannelTokenDeliveryModesSupported gets a reference to the given []string and assigns it to the BackchannelTokenDeliveryModesSupported field.
func (o *OAuthMetadata) SetBackchannelTokenDeliveryModesSupported(v []string) {
	o.BackchannelTokenDeliveryModesSupported = v
}

// GetClaimsSupported returns the ClaimsSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetClaimsSupported() []string {
	if o == nil || o.ClaimsSupported == nil {
		var ret []string
		return ret
	}
	return o.ClaimsSupported
}

// GetClaimsSupportedOk returns a tuple with the ClaimsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetClaimsSupportedOk() ([]string, bool) {
	if o == nil || o.ClaimsSupported == nil {
		return nil, false
	}
	return o.ClaimsSupported, true
}

// HasClaimsSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasClaimsSupported() bool {
	if o != nil && o.ClaimsSupported != nil {
		return true
	}

	return false
}

// SetClaimsSupported gets a reference to the given []string and assigns it to the ClaimsSupported field.
func (o *OAuthMetadata) SetClaimsSupported(v []string) {
	o.ClaimsSupported = v
}

// GetCodeChallengeMethodsSupported returns the CodeChallengeMethodsSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetCodeChallengeMethodsSupported() []string {
	if o == nil || o.CodeChallengeMethodsSupported == nil {
		var ret []string
		return ret
	}
	return o.CodeChallengeMethodsSupported
}

// GetCodeChallengeMethodsSupportedOk returns a tuple with the CodeChallengeMethodsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetCodeChallengeMethodsSupportedOk() ([]string, bool) {
	if o == nil || o.CodeChallengeMethodsSupported == nil {
		return nil, false
	}
	return o.CodeChallengeMethodsSupported, true
}

// HasCodeChallengeMethodsSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasCodeChallengeMethodsSupported() bool {
	if o != nil && o.CodeChallengeMethodsSupported != nil {
		return true
	}

	return false
}

// SetCodeChallengeMethodsSupported gets a reference to the given []string and assigns it to the CodeChallengeMethodsSupported field.
func (o *OAuthMetadata) SetCodeChallengeMethodsSupported(v []string) {
	o.CodeChallengeMethodsSupported = v
}

// GetDeviceAuthorizationEndpoint returns the DeviceAuthorizationEndpoint field value if set, zero value otherwise.
func (o *OAuthMetadata) GetDeviceAuthorizationEndpoint() string {
	if o == nil || o.DeviceAuthorizationEndpoint == nil {
		var ret string
		return ret
	}
	return *o.DeviceAuthorizationEndpoint
}

// GetDeviceAuthorizationEndpointOk returns a tuple with the DeviceAuthorizationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetDeviceAuthorizationEndpointOk() (*string, bool) {
	if o == nil || o.DeviceAuthorizationEndpoint == nil {
		return nil, false
	}
	return o.DeviceAuthorizationEndpoint, true
}

// HasDeviceAuthorizationEndpoint returns a boolean if a field has been set.
func (o *OAuthMetadata) HasDeviceAuthorizationEndpoint() bool {
	if o != nil && o.DeviceAuthorizationEndpoint != nil {
		return true
	}

	return false
}

// SetDeviceAuthorizationEndpoint gets a reference to the given string and assigns it to the DeviceAuthorizationEndpoint field.
func (o *OAuthMetadata) SetDeviceAuthorizationEndpoint(v string) {
	o.DeviceAuthorizationEndpoint = &v
}

// GetDpopSigningAlgValuesSupported returns the DpopSigningAlgValuesSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetDpopSigningAlgValuesSupported() []string {
	if o == nil || o.DpopSigningAlgValuesSupported == nil {
		var ret []string
		return ret
	}
	return o.DpopSigningAlgValuesSupported
}

// GetDpopSigningAlgValuesSupportedOk returns a tuple with the DpopSigningAlgValuesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetDpopSigningAlgValuesSupportedOk() ([]string, bool) {
	if o == nil || o.DpopSigningAlgValuesSupported == nil {
		return nil, false
	}
	return o.DpopSigningAlgValuesSupported, true
}

// HasDpopSigningAlgValuesSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasDpopSigningAlgValuesSupported() bool {
	if o != nil && o.DpopSigningAlgValuesSupported != nil {
		return true
	}

	return false
}

// SetDpopSigningAlgValuesSupported gets a reference to the given []string and assigns it to the DpopSigningAlgValuesSupported field.
func (o *OAuthMetadata) SetDpopSigningAlgValuesSupported(v []string) {
	o.DpopSigningAlgValuesSupported = v
}

// GetEndSessionEndpoint returns the EndSessionEndpoint field value if set, zero value otherwise.
func (o *OAuthMetadata) GetEndSessionEndpoint() string {
	if o == nil || o.EndSessionEndpoint == nil {
		var ret string
		return ret
	}
	return *o.EndSessionEndpoint
}

// GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetEndSessionEndpointOk() (*string, bool) {
	if o == nil || o.EndSessionEndpoint == nil {
		return nil, false
	}
	return o.EndSessionEndpoint, true
}

// HasEndSessionEndpoint returns a boolean if a field has been set.
func (o *OAuthMetadata) HasEndSessionEndpoint() bool {
	if o != nil && o.EndSessionEndpoint != nil {
		return true
	}

	return false
}

// SetEndSessionEndpoint gets a reference to the given string and assigns it to the EndSessionEndpoint field.
func (o *OAuthMetadata) SetEndSessionEndpoint(v string) {
	o.EndSessionEndpoint = &v
}

// GetGrantTypesSupported returns the GrantTypesSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetGrantTypesSupported() []string {
	if o == nil || o.GrantTypesSupported == nil {
		var ret []string
		return ret
	}
	return o.GrantTypesSupported
}

// GetGrantTypesSupportedOk returns a tuple with the GrantTypesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetGrantTypesSupportedOk() ([]string, bool) {
	if o == nil || o.GrantTypesSupported == nil {
		return nil, false
	}
	return o.GrantTypesSupported, true
}

// HasGrantTypesSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasGrantTypesSupported() bool {
	if o != nil && o.GrantTypesSupported != nil {
		return true
	}

	return false
}

// SetGrantTypesSupported gets a reference to the given []string and assigns it to the GrantTypesSupported field.
func (o *OAuthMetadata) SetGrantTypesSupported(v []string) {
	o.GrantTypesSupported = v
}

// GetIntrospectionEndpoint returns the IntrospectionEndpoint field value if set, zero value otherwise.
func (o *OAuthMetadata) GetIntrospectionEndpoint() string {
	if o == nil || o.IntrospectionEndpoint == nil {
		var ret string
		return ret
	}
	return *o.IntrospectionEndpoint
}

// GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetIntrospectionEndpointOk() (*string, bool) {
	if o == nil || o.IntrospectionEndpoint == nil {
		return nil, false
	}
	return o.IntrospectionEndpoint, true
}

// HasIntrospectionEndpoint returns a boolean if a field has been set.
func (o *OAuthMetadata) HasIntrospectionEndpoint() bool {
	if o != nil && o.IntrospectionEndpoint != nil {
		return true
	}

	return false
}

// SetIntrospectionEndpoint gets a reference to the given string and assigns it to the IntrospectionEndpoint field.
func (o *OAuthMetadata) SetIntrospectionEndpoint(v string) {
	o.IntrospectionEndpoint = &v
}

// GetIntrospectionEndpointAuthMethodsSupported returns the IntrospectionEndpointAuthMethodsSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetIntrospectionEndpointAuthMethodsSupported() []string {
	if o == nil || o.IntrospectionEndpointAuthMethodsSupported == nil {
		var ret []string
		return ret
	}
	return o.IntrospectionEndpointAuthMethodsSupported
}

// GetIntrospectionEndpointAuthMethodsSupportedOk returns a tuple with the IntrospectionEndpointAuthMethodsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetIntrospectionEndpointAuthMethodsSupportedOk() ([]string, bool) {
	if o == nil || o.IntrospectionEndpointAuthMethodsSupported == nil {
		return nil, false
	}
	return o.IntrospectionEndpointAuthMethodsSupported, true
}

// HasIntrospectionEndpointAuthMethodsSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasIntrospectionEndpointAuthMethodsSupported() bool {
	if o != nil && o.IntrospectionEndpointAuthMethodsSupported != nil {
		return true
	}

	return false
}

// SetIntrospectionEndpointAuthMethodsSupported gets a reference to the given []string and assigns it to the IntrospectionEndpointAuthMethodsSupported field.
func (o *OAuthMetadata) SetIntrospectionEndpointAuthMethodsSupported(v []string) {
	o.IntrospectionEndpointAuthMethodsSupported = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *OAuthMetadata) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *OAuthMetadata) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *OAuthMetadata) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *OAuthMetadata) GetJwksUri() string {
	if o == nil || o.JwksUri == nil {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetJwksUriOk() (*string, bool) {
	if o == nil || o.JwksUri == nil {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *OAuthMetadata) HasJwksUri() bool {
	if o != nil && o.JwksUri != nil {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *OAuthMetadata) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetPushedAuthorizationRequestEndpoint returns the PushedAuthorizationRequestEndpoint field value if set, zero value otherwise.
func (o *OAuthMetadata) GetPushedAuthorizationRequestEndpoint() string {
	if o == nil || o.PushedAuthorizationRequestEndpoint == nil {
		var ret string
		return ret
	}
	return *o.PushedAuthorizationRequestEndpoint
}

// GetPushedAuthorizationRequestEndpointOk returns a tuple with the PushedAuthorizationRequestEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetPushedAuthorizationRequestEndpointOk() (*string, bool) {
	if o == nil || o.PushedAuthorizationRequestEndpoint == nil {
		return nil, false
	}
	return o.PushedAuthorizationRequestEndpoint, true
}

// HasPushedAuthorizationRequestEndpoint returns a boolean if a field has been set.
func (o *OAuthMetadata) HasPushedAuthorizationRequestEndpoint() bool {
	if o != nil && o.PushedAuthorizationRequestEndpoint != nil {
		return true
	}

	return false
}

// SetPushedAuthorizationRequestEndpoint gets a reference to the given string and assigns it to the PushedAuthorizationRequestEndpoint field.
func (o *OAuthMetadata) SetPushedAuthorizationRequestEndpoint(v string) {
	o.PushedAuthorizationRequestEndpoint = &v
}

// GetRegistrationEndpoint returns the RegistrationEndpoint field value if set, zero value otherwise.
func (o *OAuthMetadata) GetRegistrationEndpoint() string {
	if o == nil || o.RegistrationEndpoint == nil {
		var ret string
		return ret
	}
	return *o.RegistrationEndpoint
}

// GetRegistrationEndpointOk returns a tuple with the RegistrationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetRegistrationEndpointOk() (*string, bool) {
	if o == nil || o.RegistrationEndpoint == nil {
		return nil, false
	}
	return o.RegistrationEndpoint, true
}

// HasRegistrationEndpoint returns a boolean if a field has been set.
func (o *OAuthMetadata) HasRegistrationEndpoint() bool {
	if o != nil && o.RegistrationEndpoint != nil {
		return true
	}

	return false
}

// SetRegistrationEndpoint gets a reference to the given string and assigns it to the RegistrationEndpoint field.
func (o *OAuthMetadata) SetRegistrationEndpoint(v string) {
	o.RegistrationEndpoint = &v
}

// GetRequestObjectSigningAlgValuesSupported returns the RequestObjectSigningAlgValuesSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetRequestObjectSigningAlgValuesSupported() []string {
	if o == nil || o.RequestObjectSigningAlgValuesSupported == nil {
		var ret []string
		return ret
	}
	return o.RequestObjectSigningAlgValuesSupported
}

// GetRequestObjectSigningAlgValuesSupportedOk returns a tuple with the RequestObjectSigningAlgValuesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetRequestObjectSigningAlgValuesSupportedOk() ([]string, bool) {
	if o == nil || o.RequestObjectSigningAlgValuesSupported == nil {
		return nil, false
	}
	return o.RequestObjectSigningAlgValuesSupported, true
}

// HasRequestObjectSigningAlgValuesSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasRequestObjectSigningAlgValuesSupported() bool {
	if o != nil && o.RequestObjectSigningAlgValuesSupported != nil {
		return true
	}

	return false
}

// SetRequestObjectSigningAlgValuesSupported gets a reference to the given []string and assigns it to the RequestObjectSigningAlgValuesSupported field.
func (o *OAuthMetadata) SetRequestObjectSigningAlgValuesSupported(v []string) {
	o.RequestObjectSigningAlgValuesSupported = v
}

// GetRequestParameterSupported returns the RequestParameterSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetRequestParameterSupported() bool {
	if o == nil || o.RequestParameterSupported == nil {
		var ret bool
		return ret
	}
	return *o.RequestParameterSupported
}

// GetRequestParameterSupportedOk returns a tuple with the RequestParameterSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetRequestParameterSupportedOk() (*bool, bool) {
	if o == nil || o.RequestParameterSupported == nil {
		return nil, false
	}
	return o.RequestParameterSupported, true
}

// HasRequestParameterSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasRequestParameterSupported() bool {
	if o != nil && o.RequestParameterSupported != nil {
		return true
	}

	return false
}

// SetRequestParameterSupported gets a reference to the given bool and assigns it to the RequestParameterSupported field.
func (o *OAuthMetadata) SetRequestParameterSupported(v bool) {
	o.RequestParameterSupported = &v
}

// GetResponseModesSupported returns the ResponseModesSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetResponseModesSupported() []string {
	if o == nil || o.ResponseModesSupported == nil {
		var ret []string
		return ret
	}
	return o.ResponseModesSupported
}

// GetResponseModesSupportedOk returns a tuple with the ResponseModesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetResponseModesSupportedOk() ([]string, bool) {
	if o == nil || o.ResponseModesSupported == nil {
		return nil, false
	}
	return o.ResponseModesSupported, true
}

// HasResponseModesSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasResponseModesSupported() bool {
	if o != nil && o.ResponseModesSupported != nil {
		return true
	}

	return false
}

// SetResponseModesSupported gets a reference to the given []string and assigns it to the ResponseModesSupported field.
func (o *OAuthMetadata) SetResponseModesSupported(v []string) {
	o.ResponseModesSupported = v
}

// GetResponseTypesSupported returns the ResponseTypesSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetResponseTypesSupported() []string {
	if o == nil || o.ResponseTypesSupported == nil {
		var ret []string
		return ret
	}
	return o.ResponseTypesSupported
}

// GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetResponseTypesSupportedOk() ([]string, bool) {
	if o == nil || o.ResponseTypesSupported == nil {
		return nil, false
	}
	return o.ResponseTypesSupported, true
}

// HasResponseTypesSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasResponseTypesSupported() bool {
	if o != nil && o.ResponseTypesSupported != nil {
		return true
	}

	return false
}

// SetResponseTypesSupported gets a reference to the given []string and assigns it to the ResponseTypesSupported field.
func (o *OAuthMetadata) SetResponseTypesSupported(v []string) {
	o.ResponseTypesSupported = v
}

// GetRevocationEndpoint returns the RevocationEndpoint field value if set, zero value otherwise.
func (o *OAuthMetadata) GetRevocationEndpoint() string {
	if o == nil || o.RevocationEndpoint == nil {
		var ret string
		return ret
	}
	return *o.RevocationEndpoint
}

// GetRevocationEndpointOk returns a tuple with the RevocationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetRevocationEndpointOk() (*string, bool) {
	if o == nil || o.RevocationEndpoint == nil {
		return nil, false
	}
	return o.RevocationEndpoint, true
}

// HasRevocationEndpoint returns a boolean if a field has been set.
func (o *OAuthMetadata) HasRevocationEndpoint() bool {
	if o != nil && o.RevocationEndpoint != nil {
		return true
	}

	return false
}

// SetRevocationEndpoint gets a reference to the given string and assigns it to the RevocationEndpoint field.
func (o *OAuthMetadata) SetRevocationEndpoint(v string) {
	o.RevocationEndpoint = &v
}

// GetRevocationEndpointAuthMethodsSupported returns the RevocationEndpointAuthMethodsSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetRevocationEndpointAuthMethodsSupported() []string {
	if o == nil || o.RevocationEndpointAuthMethodsSupported == nil {
		var ret []string
		return ret
	}
	return o.RevocationEndpointAuthMethodsSupported
}

// GetRevocationEndpointAuthMethodsSupportedOk returns a tuple with the RevocationEndpointAuthMethodsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetRevocationEndpointAuthMethodsSupportedOk() ([]string, bool) {
	if o == nil || o.RevocationEndpointAuthMethodsSupported == nil {
		return nil, false
	}
	return o.RevocationEndpointAuthMethodsSupported, true
}

// HasRevocationEndpointAuthMethodsSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasRevocationEndpointAuthMethodsSupported() bool {
	if o != nil && o.RevocationEndpointAuthMethodsSupported != nil {
		return true
	}

	return false
}

// SetRevocationEndpointAuthMethodsSupported gets a reference to the given []string and assigns it to the RevocationEndpointAuthMethodsSupported field.
func (o *OAuthMetadata) SetRevocationEndpointAuthMethodsSupported(v []string) {
	o.RevocationEndpointAuthMethodsSupported = v
}

// GetScopesSupported returns the ScopesSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetScopesSupported() []string {
	if o == nil || o.ScopesSupported == nil {
		var ret []string
		return ret
	}
	return o.ScopesSupported
}

// GetScopesSupportedOk returns a tuple with the ScopesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetScopesSupportedOk() ([]string, bool) {
	if o == nil || o.ScopesSupported == nil {
		return nil, false
	}
	return o.ScopesSupported, true
}

// HasScopesSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasScopesSupported() bool {
	if o != nil && o.ScopesSupported != nil {
		return true
	}

	return false
}

// SetScopesSupported gets a reference to the given []string and assigns it to the ScopesSupported field.
func (o *OAuthMetadata) SetScopesSupported(v []string) {
	o.ScopesSupported = v
}

// GetSubjectTypesSupported returns the SubjectTypesSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetSubjectTypesSupported() []string {
	if o == nil || o.SubjectTypesSupported == nil {
		var ret []string
		return ret
	}
	return o.SubjectTypesSupported
}

// GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetSubjectTypesSupportedOk() ([]string, bool) {
	if o == nil || o.SubjectTypesSupported == nil {
		return nil, false
	}
	return o.SubjectTypesSupported, true
}

// HasSubjectTypesSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasSubjectTypesSupported() bool {
	if o != nil && o.SubjectTypesSupported != nil {
		return true
	}

	return false
}

// SetSubjectTypesSupported gets a reference to the given []string and assigns it to the SubjectTypesSupported field.
func (o *OAuthMetadata) SetSubjectTypesSupported(v []string) {
	o.SubjectTypesSupported = v
}

// GetTokenEndpoint returns the TokenEndpoint field value if set, zero value otherwise.
func (o *OAuthMetadata) GetTokenEndpoint() string {
	if o == nil || o.TokenEndpoint == nil {
		var ret string
		return ret
	}
	return *o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetTokenEndpointOk() (*string, bool) {
	if o == nil || o.TokenEndpoint == nil {
		return nil, false
	}
	return o.TokenEndpoint, true
}

// HasTokenEndpoint returns a boolean if a field has been set.
func (o *OAuthMetadata) HasTokenEndpoint() bool {
	if o != nil && o.TokenEndpoint != nil {
		return true
	}

	return false
}

// SetTokenEndpoint gets a reference to the given string and assigns it to the TokenEndpoint field.
func (o *OAuthMetadata) SetTokenEndpoint(v string) {
	o.TokenEndpoint = &v
}

// GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field value if set, zero value otherwise.
func (o *OAuthMetadata) GetTokenEndpointAuthMethodsSupported() []string {
	if o == nil || o.TokenEndpointAuthMethodsSupported == nil {
		var ret []string
		return ret
	}
	return o.TokenEndpointAuthMethodsSupported
}

// GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthMetadata) GetTokenEndpointAuthMethodsSupportedOk() ([]string, bool) {
	if o == nil || o.TokenEndpointAuthMethodsSupported == nil {
		return nil, false
	}
	return o.TokenEndpointAuthMethodsSupported, true
}

// HasTokenEndpointAuthMethodsSupported returns a boolean if a field has been set.
func (o *OAuthMetadata) HasTokenEndpointAuthMethodsSupported() bool {
	if o != nil && o.TokenEndpointAuthMethodsSupported != nil {
		return true
	}

	return false
}

// SetTokenEndpointAuthMethodsSupported gets a reference to the given []string and assigns it to the TokenEndpointAuthMethodsSupported field.
func (o *OAuthMetadata) SetTokenEndpointAuthMethodsSupported(v []string) {
	o.TokenEndpointAuthMethodsSupported = v
}

func (o OAuthMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationEndpoint != nil {
		toSerialize["authorization_endpoint"] = o.AuthorizationEndpoint
	}
	if o.BackchannelAuthenticationRequestSigningAlgValuesSupported != nil {
		toSerialize["backchannel_authentication_request_signing_alg_values_supported"] = o.BackchannelAuthenticationRequestSigningAlgValuesSupported
	}
	if o.BackchannelTokenDeliveryModesSupported != nil {
		toSerialize["backchannel_token_delivery_modes_supported"] = o.BackchannelTokenDeliveryModesSupported
	}
	if o.ClaimsSupported != nil {
		toSerialize["claims_supported"] = o.ClaimsSupported
	}
	if o.CodeChallengeMethodsSupported != nil {
		toSerialize["code_challenge_methods_supported"] = o.CodeChallengeMethodsSupported
	}
	if o.DeviceAuthorizationEndpoint != nil {
		toSerialize["device_authorization_endpoint"] = o.DeviceAuthorizationEndpoint
	}
	if o.DpopSigningAlgValuesSupported != nil {
		toSerialize["dpop_signing_alg_values_supported"] = o.DpopSigningAlgValuesSupported
	}
	if o.EndSessionEndpoint != nil {
		toSerialize["end_session_endpoint"] = o.EndSessionEndpoint
	}
	if o.GrantTypesSupported != nil {
		toSerialize["grant_types_supported"] = o.GrantTypesSupported
	}
	if o.IntrospectionEndpoint != nil {
		toSerialize["introspection_endpoint"] = o.IntrospectionEndpoint
	}
	if o.IntrospectionEndpointAuthMethodsSupported != nil {
		toSerialize["introspection_endpoint_auth_methods_supported"] = o.IntrospectionEndpointAuthMethodsSupported
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.JwksUri != nil {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if o.PushedAuthorizationRequestEndpoint != nil {
		toSerialize["pushed_authorization_request_endpoint"] = o.PushedAuthorizationRequestEndpoint
	}
	if o.RegistrationEndpoint != nil {
		toSerialize["registration_endpoint"] = o.RegistrationEndpoint
	}
	if o.RequestObjectSigningAlgValuesSupported != nil {
		toSerialize["request_object_signing_alg_values_supported"] = o.RequestObjectSigningAlgValuesSupported
	}
	if o.RequestParameterSupported != nil {
		toSerialize["request_parameter_supported"] = o.RequestParameterSupported
	}
	if o.ResponseModesSupported != nil {
		toSerialize["response_modes_supported"] = o.ResponseModesSupported
	}
	if o.ResponseTypesSupported != nil {
		toSerialize["response_types_supported"] = o.ResponseTypesSupported
	}
	if o.RevocationEndpoint != nil {
		toSerialize["revocation_endpoint"] = o.RevocationEndpoint
	}
	if o.RevocationEndpointAuthMethodsSupported != nil {
		toSerialize["revocation_endpoint_auth_methods_supported"] = o.RevocationEndpointAuthMethodsSupported
	}
	if o.ScopesSupported != nil {
		toSerialize["scopes_supported"] = o.ScopesSupported
	}
	if o.SubjectTypesSupported != nil {
		toSerialize["subject_types_supported"] = o.SubjectTypesSupported
	}
	if o.TokenEndpoint != nil {
		toSerialize["token_endpoint"] = o.TokenEndpoint
	}
	if o.TokenEndpointAuthMethodsSupported != nil {
		toSerialize["token_endpoint_auth_methods_supported"] = o.TokenEndpointAuthMethodsSupported
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuthMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varOAuthMetadata := _OAuthMetadata{}

	err = json.Unmarshal(bytes, &varOAuthMetadata)
	if err == nil {
		*o = OAuthMetadata(varOAuthMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authorization_endpoint")
		delete(additionalProperties, "backchannel_authentication_request_signing_alg_values_supported")
		delete(additionalProperties, "backchannel_token_delivery_modes_supported")
		delete(additionalProperties, "claims_supported")
		delete(additionalProperties, "code_challenge_methods_supported")
		delete(additionalProperties, "device_authorization_endpoint")
		delete(additionalProperties, "dpop_signing_alg_values_supported")
		delete(additionalProperties, "end_session_endpoint")
		delete(additionalProperties, "grant_types_supported")
		delete(additionalProperties, "introspection_endpoint")
		delete(additionalProperties, "introspection_endpoint_auth_methods_supported")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "jwks_uri")
		delete(additionalProperties, "pushed_authorization_request_endpoint")
		delete(additionalProperties, "registration_endpoint")
		delete(additionalProperties, "request_object_signing_alg_values_supported")
		delete(additionalProperties, "request_parameter_supported")
		delete(additionalProperties, "response_modes_supported")
		delete(additionalProperties, "response_types_supported")
		delete(additionalProperties, "revocation_endpoint")
		delete(additionalProperties, "revocation_endpoint_auth_methods_supported")
		delete(additionalProperties, "scopes_supported")
		delete(additionalProperties, "subject_types_supported")
		delete(additionalProperties, "token_endpoint")
		delete(additionalProperties, "token_endpoint_auth_methods_supported")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuthMetadata struct {
	value *OAuthMetadata
	isSet bool
}

func (v NullableOAuthMetadata) Get() *OAuthMetadata {
	return v.value
}

func (v *NullableOAuthMetadata) Set(val *OAuthMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthMetadata(val *OAuthMetadata) *NullableOAuthMetadata {
	return &NullableOAuthMetadata{value: val, isSet: true}
}

func (v NullableOAuthMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

