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

// OpenIdConnectApplicationSettingsClient struct for OpenIdConnectApplicationSettingsClient
type OpenIdConnectApplicationSettingsClient struct {
	ApplicationType *string `json:"application_type,omitempty"`
	ClientUri *string `json:"client_uri,omitempty"`
	ConsentMethod *string `json:"consent_method,omitempty"`
	// Indicates that the client application uses Demonstrating Proof-of-Possession (DPoP) for token requests. If `true`, the authorization server rejects token requests from this client that don't contain the DPoP header.
	DpopBoundAccessTokens *bool `json:"dpop_bound_access_tokens,omitempty"`
	// Include user session details.
	FrontchannelLogoutSessionRequired *bool `json:"frontchannel_logout_session_required,omitempty"`
	// URL where Okta sends the logout request.
	FrontchannelLogoutUri *string `json:"frontchannel_logout_uri,omitempty"`
	GrantTypes []string `json:"grant_types,omitempty"`
	IdpInitiatedLogin *OpenIdConnectApplicationIdpInitiatedLogin `json:"idp_initiated_login,omitempty"`
	InitiateLoginUri *string `json:"initiate_login_uri,omitempty"`
	IssuerMode *string `json:"issuer_mode,omitempty"`
	Jwks *OpenIdConnectApplicationSettingsClientKeys `json:"jwks,omitempty"`
	// URL string that references a JSON Web Key Set for validating JWTs presented to Okta.
	JwksUri *string `json:"jwks_uri,omitempty"`
	LogoUri *string `json:"logo_uri,omitempty"`
	// Allows the app to participate in front-channel single logout.
	ParticipateSlo *bool `json:"participate_slo,omitempty"`
	PolicyUri *string `json:"policy_uri,omitempty"`
	PostLogoutRedirectUris []string `json:"post_logout_redirect_uris,omitempty"`
	RedirectUris []string `json:"redirect_uris,omitempty"`
	RefreshToken *OpenIdConnectApplicationSettingsRefreshToken `json:"refresh_token,omitempty"`
	ResponseTypes []string `json:"response_types,omitempty"`
	TosUri *string `json:"tos_uri,omitempty"`
	WildcardRedirect *string `json:"wildcard_redirect,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenIdConnectApplicationSettingsClient OpenIdConnectApplicationSettingsClient

// NewOpenIdConnectApplicationSettingsClient instantiates a new OpenIdConnectApplicationSettingsClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectApplicationSettingsClient() *OpenIdConnectApplicationSettingsClient {
	this := OpenIdConnectApplicationSettingsClient{}
	var dpopBoundAccessTokens bool = false
	this.DpopBoundAccessTokens = &dpopBoundAccessTokens
	return &this
}

// NewOpenIdConnectApplicationSettingsClientWithDefaults instantiates a new OpenIdConnectApplicationSettingsClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIdConnectApplicationSettingsClientWithDefaults() *OpenIdConnectApplicationSettingsClient {
	this := OpenIdConnectApplicationSettingsClient{}
	var dpopBoundAccessTokens bool = false
	this.DpopBoundAccessTokens = &dpopBoundAccessTokens
	return &this
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetApplicationType() string {
	if o == nil || o.ApplicationType == nil {
		var ret string
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetApplicationTypeOk() (*string, bool) {
	if o == nil || o.ApplicationType == nil {
		return nil, false
	}
	return o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasApplicationType() bool {
	if o != nil && o.ApplicationType != nil {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given string and assigns it to the ApplicationType field.
func (o *OpenIdConnectApplicationSettingsClient) SetApplicationType(v string) {
	o.ApplicationType = &v
}

// GetClientUri returns the ClientUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetClientUri() string {
	if o == nil || o.ClientUri == nil {
		var ret string
		return ret
	}
	return *o.ClientUri
}

// GetClientUriOk returns a tuple with the ClientUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetClientUriOk() (*string, bool) {
	if o == nil || o.ClientUri == nil {
		return nil, false
	}
	return o.ClientUri, true
}

// HasClientUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasClientUri() bool {
	if o != nil && o.ClientUri != nil {
		return true
	}

	return false
}

// SetClientUri gets a reference to the given string and assigns it to the ClientUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetClientUri(v string) {
	o.ClientUri = &v
}

// GetConsentMethod returns the ConsentMethod field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetConsentMethod() string {
	if o == nil || o.ConsentMethod == nil {
		var ret string
		return ret
	}
	return *o.ConsentMethod
}

// GetConsentMethodOk returns a tuple with the ConsentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetConsentMethodOk() (*string, bool) {
	if o == nil || o.ConsentMethod == nil {
		return nil, false
	}
	return o.ConsentMethod, true
}

// HasConsentMethod returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasConsentMethod() bool {
	if o != nil && o.ConsentMethod != nil {
		return true
	}

	return false
}

// SetConsentMethod gets a reference to the given string and assigns it to the ConsentMethod field.
func (o *OpenIdConnectApplicationSettingsClient) SetConsentMethod(v string) {
	o.ConsentMethod = &v
}

// GetDpopBoundAccessTokens returns the DpopBoundAccessTokens field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetDpopBoundAccessTokens() bool {
	if o == nil || o.DpopBoundAccessTokens == nil {
		var ret bool
		return ret
	}
	return *o.DpopBoundAccessTokens
}

// GetDpopBoundAccessTokensOk returns a tuple with the DpopBoundAccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetDpopBoundAccessTokensOk() (*bool, bool) {
	if o == nil || o.DpopBoundAccessTokens == nil {
		return nil, false
	}
	return o.DpopBoundAccessTokens, true
}

// HasDpopBoundAccessTokens returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasDpopBoundAccessTokens() bool {
	if o != nil && o.DpopBoundAccessTokens != nil {
		return true
	}

	return false
}

// SetDpopBoundAccessTokens gets a reference to the given bool and assigns it to the DpopBoundAccessTokens field.
func (o *OpenIdConnectApplicationSettingsClient) SetDpopBoundAccessTokens(v bool) {
	o.DpopBoundAccessTokens = &v
}

// GetFrontchannelLogoutSessionRequired returns the FrontchannelLogoutSessionRequired field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutSessionRequired() bool {
	if o == nil || o.FrontchannelLogoutSessionRequired == nil {
		var ret bool
		return ret
	}
	return *o.FrontchannelLogoutSessionRequired
}

// GetFrontchannelLogoutSessionRequiredOk returns a tuple with the FrontchannelLogoutSessionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutSessionRequiredOk() (*bool, bool) {
	if o == nil || o.FrontchannelLogoutSessionRequired == nil {
		return nil, false
	}
	return o.FrontchannelLogoutSessionRequired, true
}

// HasFrontchannelLogoutSessionRequired returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasFrontchannelLogoutSessionRequired() bool {
	if o != nil && o.FrontchannelLogoutSessionRequired != nil {
		return true
	}

	return false
}

// SetFrontchannelLogoutSessionRequired gets a reference to the given bool and assigns it to the FrontchannelLogoutSessionRequired field.
func (o *OpenIdConnectApplicationSettingsClient) SetFrontchannelLogoutSessionRequired(v bool) {
	o.FrontchannelLogoutSessionRequired = &v
}

// GetFrontchannelLogoutUri returns the FrontchannelLogoutUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutUri() string {
	if o == nil || o.FrontchannelLogoutUri == nil {
		var ret string
		return ret
	}
	return *o.FrontchannelLogoutUri
}

// GetFrontchannelLogoutUriOk returns a tuple with the FrontchannelLogoutUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutUriOk() (*string, bool) {
	if o == nil || o.FrontchannelLogoutUri == nil {
		return nil, false
	}
	return o.FrontchannelLogoutUri, true
}

// HasFrontchannelLogoutUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasFrontchannelLogoutUri() bool {
	if o != nil && o.FrontchannelLogoutUri != nil {
		return true
	}

	return false
}

// SetFrontchannelLogoutUri gets a reference to the given string and assigns it to the FrontchannelLogoutUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetFrontchannelLogoutUri(v string) {
	o.FrontchannelLogoutUri = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetGrantTypes() []string {
	if o == nil || o.GrantTypes == nil {
		var ret []string
		return ret
	}
	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetGrantTypesOk() ([]string, bool) {
	if o == nil || o.GrantTypes == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasGrantTypes() bool {
	if o != nil && o.GrantTypes != nil {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *OpenIdConnectApplicationSettingsClient) SetGrantTypes(v []string) {
	o.GrantTypes = v
}

// GetIdpInitiatedLogin returns the IdpInitiatedLogin field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetIdpInitiatedLogin() OpenIdConnectApplicationIdpInitiatedLogin {
	if o == nil || o.IdpInitiatedLogin == nil {
		var ret OpenIdConnectApplicationIdpInitiatedLogin
		return ret
	}
	return *o.IdpInitiatedLogin
}

// GetIdpInitiatedLoginOk returns a tuple with the IdpInitiatedLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetIdpInitiatedLoginOk() (*OpenIdConnectApplicationIdpInitiatedLogin, bool) {
	if o == nil || o.IdpInitiatedLogin == nil {
		return nil, false
	}
	return o.IdpInitiatedLogin, true
}

// HasIdpInitiatedLogin returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasIdpInitiatedLogin() bool {
	if o != nil && o.IdpInitiatedLogin != nil {
		return true
	}

	return false
}

// SetIdpInitiatedLogin gets a reference to the given OpenIdConnectApplicationIdpInitiatedLogin and assigns it to the IdpInitiatedLogin field.
func (o *OpenIdConnectApplicationSettingsClient) SetIdpInitiatedLogin(v OpenIdConnectApplicationIdpInitiatedLogin) {
	o.IdpInitiatedLogin = &v
}

// GetInitiateLoginUri returns the InitiateLoginUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetInitiateLoginUri() string {
	if o == nil || o.InitiateLoginUri == nil {
		var ret string
		return ret
	}
	return *o.InitiateLoginUri
}

// GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetInitiateLoginUriOk() (*string, bool) {
	if o == nil || o.InitiateLoginUri == nil {
		return nil, false
	}
	return o.InitiateLoginUri, true
}

// HasInitiateLoginUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasInitiateLoginUri() bool {
	if o != nil && o.InitiateLoginUri != nil {
		return true
	}

	return false
}

// SetInitiateLoginUri gets a reference to the given string and assigns it to the InitiateLoginUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetInitiateLoginUri(v string) {
	o.InitiateLoginUri = &v
}

// GetIssuerMode returns the IssuerMode field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetIssuerMode() string {
	if o == nil || o.IssuerMode == nil {
		var ret string
		return ret
	}
	return *o.IssuerMode
}

// GetIssuerModeOk returns a tuple with the IssuerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetIssuerModeOk() (*string, bool) {
	if o == nil || o.IssuerMode == nil {
		return nil, false
	}
	return o.IssuerMode, true
}

// HasIssuerMode returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasIssuerMode() bool {
	if o != nil && o.IssuerMode != nil {
		return true
	}

	return false
}

// SetIssuerMode gets a reference to the given string and assigns it to the IssuerMode field.
func (o *OpenIdConnectApplicationSettingsClient) SetIssuerMode(v string) {
	o.IssuerMode = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetJwks() OpenIdConnectApplicationSettingsClientKeys {
	if o == nil || o.Jwks == nil {
		var ret OpenIdConnectApplicationSettingsClientKeys
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetJwksOk() (*OpenIdConnectApplicationSettingsClientKeys, bool) {
	if o == nil || o.Jwks == nil {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasJwks() bool {
	if o != nil && o.Jwks != nil {
		return true
	}

	return false
}

// SetJwks gets a reference to the given OpenIdConnectApplicationSettingsClientKeys and assigns it to the Jwks field.
func (o *OpenIdConnectApplicationSettingsClient) SetJwks(v OpenIdConnectApplicationSettingsClientKeys) {
	o.Jwks = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetJwksUri() string {
	if o == nil || o.JwksUri == nil {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetJwksUriOk() (*string, bool) {
	if o == nil || o.JwksUri == nil {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasJwksUri() bool {
	if o != nil && o.JwksUri != nil {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetLogoUri returns the LogoUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetLogoUri() string {
	if o == nil || o.LogoUri == nil {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetLogoUriOk() (*string, bool) {
	if o == nil || o.LogoUri == nil {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasLogoUri() bool {
	if o != nil && o.LogoUri != nil {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetLogoUri(v string) {
	o.LogoUri = &v
}

// GetParticipateSlo returns the ParticipateSlo field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetParticipateSlo() bool {
	if o == nil || o.ParticipateSlo == nil {
		var ret bool
		return ret
	}
	return *o.ParticipateSlo
}

// GetParticipateSloOk returns a tuple with the ParticipateSlo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetParticipateSloOk() (*bool, bool) {
	if o == nil || o.ParticipateSlo == nil {
		return nil, false
	}
	return o.ParticipateSlo, true
}

// HasParticipateSlo returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasParticipateSlo() bool {
	if o != nil && o.ParticipateSlo != nil {
		return true
	}

	return false
}

// SetParticipateSlo gets a reference to the given bool and assigns it to the ParticipateSlo field.
func (o *OpenIdConnectApplicationSettingsClient) SetParticipateSlo(v bool) {
	o.ParticipateSlo = &v
}

// GetPolicyUri returns the PolicyUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetPolicyUri() string {
	if o == nil || o.PolicyUri == nil {
		var ret string
		return ret
	}
	return *o.PolicyUri
}

// GetPolicyUriOk returns a tuple with the PolicyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetPolicyUriOk() (*string, bool) {
	if o == nil || o.PolicyUri == nil {
		return nil, false
	}
	return o.PolicyUri, true
}

// HasPolicyUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasPolicyUri() bool {
	if o != nil && o.PolicyUri != nil {
		return true
	}

	return false
}

// SetPolicyUri gets a reference to the given string and assigns it to the PolicyUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetPolicyUri(v string) {
	o.PolicyUri = &v
}

// GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetPostLogoutRedirectUris() []string {
	if o == nil || o.PostLogoutRedirectUris == nil {
		var ret []string
		return ret
	}
	return o.PostLogoutRedirectUris
}

// GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetPostLogoutRedirectUrisOk() ([]string, bool) {
	if o == nil || o.PostLogoutRedirectUris == nil {
		return nil, false
	}
	return o.PostLogoutRedirectUris, true
}

// HasPostLogoutRedirectUris returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasPostLogoutRedirectUris() bool {
	if o != nil && o.PostLogoutRedirectUris != nil {
		return true
	}

	return false
}

// SetPostLogoutRedirectUris gets a reference to the given []string and assigns it to the PostLogoutRedirectUris field.
func (o *OpenIdConnectApplicationSettingsClient) SetPostLogoutRedirectUris(v []string) {
	o.PostLogoutRedirectUris = v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetRedirectUris() []string {
	if o == nil || o.RedirectUris == nil {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || o.RedirectUris == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasRedirectUris() bool {
	if o != nil && o.RedirectUris != nil {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *OpenIdConnectApplicationSettingsClient) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetRefreshToken() OpenIdConnectApplicationSettingsRefreshToken {
	if o == nil || o.RefreshToken == nil {
		var ret OpenIdConnectApplicationSettingsRefreshToken
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetRefreshTokenOk() (*OpenIdConnectApplicationSettingsRefreshToken, bool) {
	if o == nil || o.RefreshToken == nil {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasRefreshToken() bool {
	if o != nil && o.RefreshToken != nil {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given OpenIdConnectApplicationSettingsRefreshToken and assigns it to the RefreshToken field.
func (o *OpenIdConnectApplicationSettingsClient) SetRefreshToken(v OpenIdConnectApplicationSettingsRefreshToken) {
	o.RefreshToken = &v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetResponseTypes() []string {
	if o == nil || o.ResponseTypes == nil {
		var ret []string
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetResponseTypesOk() ([]string, bool) {
	if o == nil || o.ResponseTypes == nil {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasResponseTypes() bool {
	if o != nil && o.ResponseTypes != nil {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []string and assigns it to the ResponseTypes field.
func (o *OpenIdConnectApplicationSettingsClient) SetResponseTypes(v []string) {
	o.ResponseTypes = v
}

// GetTosUri returns the TosUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetTosUri() string {
	if o == nil || o.TosUri == nil {
		var ret string
		return ret
	}
	return *o.TosUri
}

// GetTosUriOk returns a tuple with the TosUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetTosUriOk() (*string, bool) {
	if o == nil || o.TosUri == nil {
		return nil, false
	}
	return o.TosUri, true
}

// HasTosUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasTosUri() bool {
	if o != nil && o.TosUri != nil {
		return true
	}

	return false
}

// SetTosUri gets a reference to the given string and assigns it to the TosUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetTosUri(v string) {
	o.TosUri = &v
}

// GetWildcardRedirect returns the WildcardRedirect field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetWildcardRedirect() string {
	if o == nil || o.WildcardRedirect == nil {
		var ret string
		return ret
	}
	return *o.WildcardRedirect
}

// GetWildcardRedirectOk returns a tuple with the WildcardRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetWildcardRedirectOk() (*string, bool) {
	if o == nil || o.WildcardRedirect == nil {
		return nil, false
	}
	return o.WildcardRedirect, true
}

// HasWildcardRedirect returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasWildcardRedirect() bool {
	if o != nil && o.WildcardRedirect != nil {
		return true
	}

	return false
}

// SetWildcardRedirect gets a reference to the given string and assigns it to the WildcardRedirect field.
func (o *OpenIdConnectApplicationSettingsClient) SetWildcardRedirect(v string) {
	o.WildcardRedirect = &v
}

func (o OpenIdConnectApplicationSettingsClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationType != nil {
		toSerialize["application_type"] = o.ApplicationType
	}
	if o.ClientUri != nil {
		toSerialize["client_uri"] = o.ClientUri
	}
	if o.ConsentMethod != nil {
		toSerialize["consent_method"] = o.ConsentMethod
	}
	if o.DpopBoundAccessTokens != nil {
		toSerialize["dpop_bound_access_tokens"] = o.DpopBoundAccessTokens
	}
	if o.FrontchannelLogoutSessionRequired != nil {
		toSerialize["frontchannel_logout_session_required"] = o.FrontchannelLogoutSessionRequired
	}
	if o.FrontchannelLogoutUri != nil {
		toSerialize["frontchannel_logout_uri"] = o.FrontchannelLogoutUri
	}
	if o.GrantTypes != nil {
		toSerialize["grant_types"] = o.GrantTypes
	}
	if o.IdpInitiatedLogin != nil {
		toSerialize["idp_initiated_login"] = o.IdpInitiatedLogin
	}
	if o.InitiateLoginUri != nil {
		toSerialize["initiate_login_uri"] = o.InitiateLoginUri
	}
	if o.IssuerMode != nil {
		toSerialize["issuer_mode"] = o.IssuerMode
	}
	if o.Jwks != nil {
		toSerialize["jwks"] = o.Jwks
	}
	if o.JwksUri != nil {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if o.LogoUri != nil {
		toSerialize["logo_uri"] = o.LogoUri
	}
	if o.ParticipateSlo != nil {
		toSerialize["participate_slo"] = o.ParticipateSlo
	}
	if o.PolicyUri != nil {
		toSerialize["policy_uri"] = o.PolicyUri
	}
	if o.PostLogoutRedirectUris != nil {
		toSerialize["post_logout_redirect_uris"] = o.PostLogoutRedirectUris
	}
	if o.RedirectUris != nil {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if o.RefreshToken != nil {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if o.ResponseTypes != nil {
		toSerialize["response_types"] = o.ResponseTypes
	}
	if o.TosUri != nil {
		toSerialize["tos_uri"] = o.TosUri
	}
	if o.WildcardRedirect != nil {
		toSerialize["wildcard_redirect"] = o.WildcardRedirect
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenIdConnectApplicationSettingsClient) UnmarshalJSON(bytes []byte) (err error) {
	varOpenIdConnectApplicationSettingsClient := _OpenIdConnectApplicationSettingsClient{}

	err = json.Unmarshal(bytes, &varOpenIdConnectApplicationSettingsClient)
	if err == nil {
		*o = OpenIdConnectApplicationSettingsClient(varOpenIdConnectApplicationSettingsClient)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "application_type")
		delete(additionalProperties, "client_uri")
		delete(additionalProperties, "consent_method")
		delete(additionalProperties, "dpop_bound_access_tokens")
		delete(additionalProperties, "frontchannel_logout_session_required")
		delete(additionalProperties, "frontchannel_logout_uri")
		delete(additionalProperties, "grant_types")
		delete(additionalProperties, "idp_initiated_login")
		delete(additionalProperties, "initiate_login_uri")
		delete(additionalProperties, "issuer_mode")
		delete(additionalProperties, "jwks")
		delete(additionalProperties, "jwks_uri")
		delete(additionalProperties, "logo_uri")
		delete(additionalProperties, "participate_slo")
		delete(additionalProperties, "policy_uri")
		delete(additionalProperties, "post_logout_redirect_uris")
		delete(additionalProperties, "redirect_uris")
		delete(additionalProperties, "refresh_token")
		delete(additionalProperties, "response_types")
		delete(additionalProperties, "tos_uri")
		delete(additionalProperties, "wildcard_redirect")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOpenIdConnectApplicationSettingsClient struct {
	value *OpenIdConnectApplicationSettingsClient
	isSet bool
}

func (v NullableOpenIdConnectApplicationSettingsClient) Get() *OpenIdConnectApplicationSettingsClient {
	return v.value
}

func (v *NullableOpenIdConnectApplicationSettingsClient) Set(val *OpenIdConnectApplicationSettingsClient) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectApplicationSettingsClient) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectApplicationSettingsClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectApplicationSettingsClient(val *OpenIdConnectApplicationSettingsClient) *NullableOpenIdConnectApplicationSettingsClient {
	return &NullableOpenIdConnectApplicationSettingsClient{value: val, isSet: true}
}

func (v NullableOpenIdConnectApplicationSettingsClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectApplicationSettingsClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

