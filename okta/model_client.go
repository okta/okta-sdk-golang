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

// Client struct for Client
type Client struct {
	// The type of client application. Default value: `web`.
	ApplicationType *string `json:"application_type,omitempty"`
	// Unique key for the client application. The `client_id` is immutable. When you create a client Application, you can't specify the `client_id` because Okta uses the application ID for the `client_id`.
	ClientId *string `json:"client_id,omitempty"`
	// Time at which the `client_id` was issued (measured in unix seconds)
	ClientIdIssuedAt *int32 `json:"client_id_issued_at,omitempty"`
	// Human-readable string name of the client application
	ClientName *string `json:"client_name,omitempty"`
	// OAuth 2.0 client secret string (used for confidential clients). The `client_secret` is shown only on the response of the creation or update of a client Application (and only if the `token_endpoint_auth_method` is one that requires a client secret). You can't specify the `client_secret`. If the `token_endpoint_auth_method` requires one, Okta generates a random `client_secret` for the client Application.
	ClientSecret NullableString `json:"client_secret,omitempty"`
	// Time at which the `client_secret` expires or 0 if it doesn't expire (measured in unix seconds)
	ClientSecretExpiresAt NullableInt32 `json:"client_secret_expires_at,omitempty"`
	// Include user session details
	FrontchannelLogoutSessionRequired *bool `json:"frontchannel_logout_session_required,omitempty"`
	// URL where Okta sends the logout request
	FrontchannelLogoutUri NullableString `json:"frontchannel_logout_uri,omitempty"`
	// Array of OAuth 2.0 grant type strings. Default value: `[authorization_code]`
	GrantTypes []string `json:"grant_types,omitempty"`
	// URL that a third party can use to initiate a login by the client
	InitiateLoginUri *string `json:"initiate_login_uri,omitempty"`
	// URL string that references a [JSON Web Key Set](https://tools.ietf.org/html/rfc7517#section-5) for validating JWTs presented to Okta
	JwksUri *string `json:"jwks_uri,omitempty"`
	// URL string that references a logo for the client consent dialog (not the sign-in dialog)
	LogoUri NullableString `json:"logo_uri,omitempty"`
	// URL string of a web page providing the client's policy document
	PolicyUri NullableString `json:"policy_uri,omitempty"`
	// Array of redirection URI strings for use for relying party initiated logouts
	PostLogoutRedirectUris []string `json:"post_logout_redirect_uris,omitempty"`
	// Array of redirection URI strings for use in redirect-based flows. All redirect URIs must be absolute URIs and must not include a fragment component. At least one redirect URI and response type is required for all client types, with the following exceptions: If the client uses the Resource Owner Password flow (if `grant_type` contains the value password) or the Client Credentials flow (if `grant_type` contains the value `client_credentials`), then no redirect URI or response type is necessary. In these cases, you can pass either null or an empty array for these attributes.
	RedirectUris []string `json:"redirect_uris,omitempty"`
	// The type of [JSON Web Key Set](https://tools.ietf.org/html/rfc7517#section-5) algorithm that must be used for signing request objects
	RequestObjectSigningAlg []string `json:"request_object_signing_alg,omitempty"`
	// Array of OAuth 2.0 response type strings. Default value: `[code]`
	ResponseTypes []string `json:"response_types,omitempty"`
	// Requested authentication method for OAuth 2.0 endpoints.
	TokenEndpointAuthMethod *string `json:"token_endpoint_auth_method,omitempty"`
	// URL string of a web page providing the client's terms of service document
	TosUri NullableString `json:"tos_uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Client Client

// NewClient instantiates a new Client object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClient() *Client {
	this := Client{}
	return &this
}

// NewClientWithDefaults instantiates a new Client object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientWithDefaults() *Client {
	this := Client{}
	return &this
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *Client) GetApplicationType() string {
	if o == nil || o.ApplicationType == nil {
		var ret string
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetApplicationTypeOk() (*string, bool) {
	if o == nil || o.ApplicationType == nil {
		return nil, false
	}
	return o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *Client) HasApplicationType() bool {
	if o != nil && o.ApplicationType != nil {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given string and assigns it to the ApplicationType field.
func (o *Client) SetApplicationType(v string) {
	o.ApplicationType = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Client) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Client) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Client) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientIdIssuedAt returns the ClientIdIssuedAt field value if set, zero value otherwise.
func (o *Client) GetClientIdIssuedAt() int32 {
	if o == nil || o.ClientIdIssuedAt == nil {
		var ret int32
		return ret
	}
	return *o.ClientIdIssuedAt
}

// GetClientIdIssuedAtOk returns a tuple with the ClientIdIssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientIdIssuedAtOk() (*int32, bool) {
	if o == nil || o.ClientIdIssuedAt == nil {
		return nil, false
	}
	return o.ClientIdIssuedAt, true
}

// HasClientIdIssuedAt returns a boolean if a field has been set.
func (o *Client) HasClientIdIssuedAt() bool {
	if o != nil && o.ClientIdIssuedAt != nil {
		return true
	}

	return false
}

// SetClientIdIssuedAt gets a reference to the given int32 and assigns it to the ClientIdIssuedAt field.
func (o *Client) SetClientIdIssuedAt(v int32) {
	o.ClientIdIssuedAt = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *Client) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *Client) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *Client) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Client) GetClientSecret() string {
	if o == nil || o.ClientSecret.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Client) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *Client) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *Client) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}
// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *Client) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *Client) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetClientSecretExpiresAt returns the ClientSecretExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Client) GetClientSecretExpiresAt() int32 {
	if o == nil || o.ClientSecretExpiresAt.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ClientSecretExpiresAt.Get()
}

// GetClientSecretExpiresAtOk returns a tuple with the ClientSecretExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Client) GetClientSecretExpiresAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecretExpiresAt.Get(), o.ClientSecretExpiresAt.IsSet()
}

// HasClientSecretExpiresAt returns a boolean if a field has been set.
func (o *Client) HasClientSecretExpiresAt() bool {
	if o != nil && o.ClientSecretExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetClientSecretExpiresAt gets a reference to the given NullableInt32 and assigns it to the ClientSecretExpiresAt field.
func (o *Client) SetClientSecretExpiresAt(v int32) {
	o.ClientSecretExpiresAt.Set(&v)
}
// SetClientSecretExpiresAtNil sets the value for ClientSecretExpiresAt to be an explicit nil
func (o *Client) SetClientSecretExpiresAtNil() {
	o.ClientSecretExpiresAt.Set(nil)
}

// UnsetClientSecretExpiresAt ensures that no value is present for ClientSecretExpiresAt, not even an explicit nil
func (o *Client) UnsetClientSecretExpiresAt() {
	o.ClientSecretExpiresAt.Unset()
}

// GetFrontchannelLogoutSessionRequired returns the FrontchannelLogoutSessionRequired field value if set, zero value otherwise.
func (o *Client) GetFrontchannelLogoutSessionRequired() bool {
	if o == nil || o.FrontchannelLogoutSessionRequired == nil {
		var ret bool
		return ret
	}
	return *o.FrontchannelLogoutSessionRequired
}

// GetFrontchannelLogoutSessionRequiredOk returns a tuple with the FrontchannelLogoutSessionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetFrontchannelLogoutSessionRequiredOk() (*bool, bool) {
	if o == nil || o.FrontchannelLogoutSessionRequired == nil {
		return nil, false
	}
	return o.FrontchannelLogoutSessionRequired, true
}

// HasFrontchannelLogoutSessionRequired returns a boolean if a field has been set.
func (o *Client) HasFrontchannelLogoutSessionRequired() bool {
	if o != nil && o.FrontchannelLogoutSessionRequired != nil {
		return true
	}

	return false
}

// SetFrontchannelLogoutSessionRequired gets a reference to the given bool and assigns it to the FrontchannelLogoutSessionRequired field.
func (o *Client) SetFrontchannelLogoutSessionRequired(v bool) {
	o.FrontchannelLogoutSessionRequired = &v
}

// GetFrontchannelLogoutUri returns the FrontchannelLogoutUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Client) GetFrontchannelLogoutUri() string {
	if o == nil || o.FrontchannelLogoutUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.FrontchannelLogoutUri.Get()
}

// GetFrontchannelLogoutUriOk returns a tuple with the FrontchannelLogoutUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Client) GetFrontchannelLogoutUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrontchannelLogoutUri.Get(), o.FrontchannelLogoutUri.IsSet()
}

// HasFrontchannelLogoutUri returns a boolean if a field has been set.
func (o *Client) HasFrontchannelLogoutUri() bool {
	if o != nil && o.FrontchannelLogoutUri.IsSet() {
		return true
	}

	return false
}

// SetFrontchannelLogoutUri gets a reference to the given NullableString and assigns it to the FrontchannelLogoutUri field.
func (o *Client) SetFrontchannelLogoutUri(v string) {
	o.FrontchannelLogoutUri.Set(&v)
}
// SetFrontchannelLogoutUriNil sets the value for FrontchannelLogoutUri to be an explicit nil
func (o *Client) SetFrontchannelLogoutUriNil() {
	o.FrontchannelLogoutUri.Set(nil)
}

// UnsetFrontchannelLogoutUri ensures that no value is present for FrontchannelLogoutUri, not even an explicit nil
func (o *Client) UnsetFrontchannelLogoutUri() {
	o.FrontchannelLogoutUri.Unset()
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *Client) GetGrantTypes() []string {
	if o == nil || o.GrantTypes == nil {
		var ret []string
		return ret
	}
	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetGrantTypesOk() ([]string, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *Client) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *Client) SetGrantTypes(v []string) {
	o.GrantTypes = v
}

// GetInitiateLoginUri returns the InitiateLoginUri field value if set, zero value otherwise.
func (o *Client) GetInitiateLoginUri() string {
	if o == nil || o.InitiateLoginUri == nil {
		var ret string
		return ret
	}
	return *o.InitiateLoginUri
}

// GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetInitiateLoginUriOk() (*string, bool) {
	if o == nil || o.InitiateLoginUri == nil {
		return nil, false
	}
	return o.InitiateLoginUri, true
}

// HasInitiateLoginUri returns a boolean if a field has been set.
func (o *Client) HasInitiateLoginUri() bool {
	if o != nil && o.InitiateLoginUri != nil {
		return true
	}

	return false
}

// SetInitiateLoginUri gets a reference to the given string and assigns it to the InitiateLoginUri field.
func (o *Client) SetInitiateLoginUri(v string) {
	o.InitiateLoginUri = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *Client) GetJwksUri() string {
	if o == nil || o.JwksUri == nil {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetJwksUriOk() (*string, bool) {
	if o == nil || o.JwksUri == nil {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *Client) HasJwksUri() bool {
	if o != nil && o.JwksUri != nil {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *Client) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Client) GetLogoUri() string {
	if o == nil || o.LogoUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.LogoUri.Get()
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Client) GetLogoUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUri.Get(), o.LogoUri.IsSet()
}

// HasLogoUri returns a boolean if a field has been set.
func (o *Client) HasLogoUri() bool {
	if o != nil && o.LogoUri.IsSet() {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given NullableString and assigns it to the LogoUri field.
func (o *Client) SetLogoUri(v string) {
	o.LogoUri.Set(&v)
}
// SetLogoUriNil sets the value for LogoUri to be an explicit nil
func (o *Client) SetLogoUriNil() {
	o.LogoUri.Set(nil)
}

// UnsetLogoUri ensures that no value is present for LogoUri, not even an explicit nil
func (o *Client) UnsetLogoUri() {
	o.LogoUri.Unset()
}

// GetPolicyUri returns the PolicyUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Client) GetPolicyUri() string {
	if o == nil || o.PolicyUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.PolicyUri.Get()
}

// GetPolicyUriOk returns a tuple with the PolicyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Client) GetPolicyUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyUri.Get(), o.PolicyUri.IsSet()
}

// HasPolicyUri returns a boolean if a field has been set.
func (o *Client) HasPolicyUri() bool {
	if o != nil && o.PolicyUri.IsSet() {
		return true
	}

	return false
}

// SetPolicyUri gets a reference to the given NullableString and assigns it to the PolicyUri field.
func (o *Client) SetPolicyUri(v string) {
	o.PolicyUri.Set(&v)
}
// SetPolicyUriNil sets the value for PolicyUri to be an explicit nil
func (o *Client) SetPolicyUriNil() {
	o.PolicyUri.Set(nil)
}

// UnsetPolicyUri ensures that no value is present for PolicyUri, not even an explicit nil
func (o *Client) UnsetPolicyUri() {
	o.PolicyUri.Unset()
}

// GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field value if set, zero value otherwise.
func (o *Client) GetPostLogoutRedirectUris() []string {
	if o == nil || o.PostLogoutRedirectUris == nil {
		var ret []string
		return ret
	}
	return o.PostLogoutRedirectUris
}

// GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetPostLogoutRedirectUrisOk() ([]string, bool) {
	if o == nil || o.PostLogoutRedirectUris == nil {
		return nil, false
	}
	return o.PostLogoutRedirectUris, true
}

// HasPostLogoutRedirectUris returns a boolean if a field has been set.
func (o *Client) HasPostLogoutRedirectUris() bool {
	if o != nil && o.PostLogoutRedirectUris != nil {
		return true
	}

	return false
}

// SetPostLogoutRedirectUris gets a reference to the given []string and assigns it to the PostLogoutRedirectUris field.
func (o *Client) SetPostLogoutRedirectUris(v []string) {
	o.PostLogoutRedirectUris = v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *Client) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *Client) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *Client) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetRequestObjectSigningAlg returns the RequestObjectSigningAlg field value if set, zero value otherwise.
func (o *Client) GetRequestObjectSigningAlg() []string {
	if o == nil || o.RequestObjectSigningAlg == nil {
		var ret []string
		return ret
	}
	return o.RequestObjectSigningAlg
}

// GetRequestObjectSigningAlgOk returns a tuple with the RequestObjectSigningAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetRequestObjectSigningAlgOk() ([]string, bool) {
	if o == nil || o.RequestObjectSigningAlg == nil {
		return nil, false
	}
	return o.RequestObjectSigningAlg, true
}

// HasRequestObjectSigningAlg returns a boolean if a field has been set.
func (o *Client) HasRequestObjectSigningAlg() bool {
	if o != nil && o.RequestObjectSigningAlg != nil {
		return true
	}

	return false
}

// SetRequestObjectSigningAlg gets a reference to the given []string and assigns it to the RequestObjectSigningAlg field.
func (o *Client) SetRequestObjectSigningAlg(v []string) {
	o.RequestObjectSigningAlg = v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *Client) GetResponseTypes() []string {
	if o == nil || o.ResponseTypes == nil {
		var ret []string
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetResponseTypesOk() ([]string, bool) {
	if o == nil || o.ResponseTypes == nil {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *Client) HasResponseTypes() bool {
	if o != nil && o.ResponseTypes != nil {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []string and assigns it to the ResponseTypes field.
func (o *Client) SetResponseTypes(v []string) {
	o.ResponseTypes = v
}

// GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field value if set, zero value otherwise.
func (o *Client) GetTokenEndpointAuthMethod() string {
	if o == nil || o.TokenEndpointAuthMethod == nil {
		var ret string
		return ret
	}
	return *o.TokenEndpointAuthMethod
}

// GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetTokenEndpointAuthMethodOk() (*string, bool) {
	if o == nil || o.TokenEndpointAuthMethod == nil {
		return nil, false
	}
	return o.TokenEndpointAuthMethod, true
}

// HasTokenEndpointAuthMethod returns a boolean if a field has been set.
func (o *Client) HasTokenEndpointAuthMethod() bool {
	if o != nil && o.TokenEndpointAuthMethod != nil {
		return true
	}

	return false
}

// SetTokenEndpointAuthMethod gets a reference to the given string and assigns it to the TokenEndpointAuthMethod field.
func (o *Client) SetTokenEndpointAuthMethod(v string) {
	o.TokenEndpointAuthMethod = &v
}

// GetTosUri returns the TosUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Client) GetTosUri() string {
	if o == nil || o.TosUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.TosUri.Get()
}

// GetTosUriOk returns a tuple with the TosUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Client) GetTosUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TosUri.Get(), o.TosUri.IsSet()
}

// HasTosUri returns a boolean if a field has been set.
func (o *Client) HasTosUri() bool {
	if o != nil && o.TosUri.IsSet() {
		return true
	}

	return false
}

// SetTosUri gets a reference to the given NullableString and assigns it to the TosUri field.
func (o *Client) SetTosUri(v string) {
	o.TosUri.Set(&v)
}
// SetTosUriNil sets the value for TosUri to be an explicit nil
func (o *Client) SetTosUriNil() {
	o.TosUri.Set(nil)
}

// UnsetTosUri ensures that no value is present for TosUri, not even an explicit nil
func (o *Client) UnsetTosUri() {
	o.TosUri.Unset()
}

func (o Client) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationType != nil {
		toSerialize["application_type"] = o.ApplicationType
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.ClientIdIssuedAt != nil {
		toSerialize["client_id_issued_at"] = o.ClientIdIssuedAt
	}
	if o.ClientName != nil {
		toSerialize["client_name"] = o.ClientName
	}
	if o.ClientSecret.IsSet() {
		toSerialize["client_secret"] = o.ClientSecret.Get()
	}
	if o.ClientSecretExpiresAt.IsSet() {
		toSerialize["client_secret_expires_at"] = o.ClientSecretExpiresAt.Get()
	}
	if o.FrontchannelLogoutSessionRequired != nil {
		toSerialize["frontchannel_logout_session_required"] = o.FrontchannelLogoutSessionRequired
	}
	if o.FrontchannelLogoutUri.IsSet() {
		toSerialize["frontchannel_logout_uri"] = o.FrontchannelLogoutUri.Get()
	}
	if o.GrantTypes != nil {
		toSerialize["grant_types"] = o.GrantTypes
	}
	if o.InitiateLoginUri != nil {
		toSerialize["initiate_login_uri"] = o.InitiateLoginUri
	}
	if o.JwksUri != nil {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if o.LogoUri.IsSet() {
		toSerialize["logo_uri"] = o.LogoUri.Get()
	}
	if o.PolicyUri.IsSet() {
		toSerialize["policy_uri"] = o.PolicyUri.Get()
	}
	if o.PostLogoutRedirectUris != nil {
		toSerialize["post_logout_redirect_uris"] = o.PostLogoutRedirectUris
	}
	if o.RedirectUris != nil {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.RequestObjectSigningAlg != nil {
		toSerialize["request_object_signing_alg"] = o.RequestObjectSigningAlg
	}
	if o.ResponseTypes != nil {
		toSerialize["response_types"] = o.ResponseTypes
	}
	if o.TokenEndpointAuthMethod != nil {
		toSerialize["token_endpoint_auth_method"] = o.TokenEndpointAuthMethod
	}
	if o.TosUri.IsSet() {
		toSerialize["tos_uri"] = o.TosUri.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Client) UnmarshalJSON(bytes []byte) (err error) {
	varClient := _Client{}

	err = json.Unmarshal(bytes, &varClient)
	if err == nil {
		*o = Client(varClient)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "application_type")
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "client_id_issued_at")
		delete(additionalProperties, "client_name")
		delete(additionalProperties, "client_secret")
		delete(additionalProperties, "client_secret_expires_at")
		delete(additionalProperties, "frontchannel_logout_session_required")
		delete(additionalProperties, "frontchannel_logout_uri")
		delete(additionalProperties, "grant_types")
		delete(additionalProperties, "initiate_login_uri")
		delete(additionalProperties, "jwks_uri")
		delete(additionalProperties, "logo_uri")
		delete(additionalProperties, "policy_uri")
		delete(additionalProperties, "post_logout_redirect_uris")
		delete(additionalProperties, "redirect_uris")
		delete(additionalProperties, "request_object_signing_alg")
		delete(additionalProperties, "response_types")
		delete(additionalProperties, "token_endpoint_auth_method")
		delete(additionalProperties, "tos_uri")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableClient struct {
	value *Client
	isSet bool
}

func (v NullableClient) Get() *Client {
	return v.value
}

func (v *NullableClient) Set(val *Client) {
	v.value = val
	v.isSet = true
}

func (v NullableClient) IsSet() bool {
	return v.isSet
}

func (v *NullableClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClient(val *Client) *NullableClient {
	return &NullableClient{value: val, isSet: true}
}

func (v NullableClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

