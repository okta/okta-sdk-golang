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

// checks if the OpenIdConnectApplicationSettingsClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIdConnectApplicationSettingsClient{}

// OpenIdConnectApplicationSettingsClient struct for OpenIdConnectApplicationSettingsClient
type OpenIdConnectApplicationSettingsClient struct {
	// The type of client app Specific `grant_types` are valid for each `application_type`. See [Create a Client Application](/openapi/okta-oauth/oauth/tag/Client/#tag/Client/operation/createClient).
	ApplicationType *string `json:"application_type,omitempty"`
	// The signing algorithm for Client-Initiated Backchannel Authentication (CIBA) signed requests using JWT. If this value isn't set and a JWT-signed request is sent, the request fails. > **Note:** This property appears for clients with `urn:openid:params:grant-type:ciba` defined as one of the `grant_types`.
	BackchannelAuthenticationRequestSigningAlg *string `json:"backchannel_authentication_request_signing_alg,omitempty"`
	// The ID of the custom authenticator that authenticates the user > **Note:** This property appears for clients with `urn:openid:params:grant-type:ciba` defined as one of the `grant_types`.
	BackchannelCustomAuthenticatorId *string `json:"backchannel_custom_authenticator_id,omitempty"`
	// The delivery mode for Client-Initiated Backchannel Authentication (CIBA).  Currently, only `poll` is supported. > **Note:** This property appears for clients with `urn:openid:params:grant-type:ciba` defined as one of the `grant_types`.
	BackchannelTokenDeliveryMode *string `json:"backchannel_token_delivery_mode,omitempty"`
	// URL string of a web page providing information about the client
	ClientUri *string `json:"client_uri,omitempty"`
	// Indicates whether user consent is required or implicit. A consent dialog appears for the end user depending on the values of three elements:  * [prompt](/openapi/okta-oauth/oauth/tag/OrgAS/#tag/OrgAS/operation/authorize!in=query&path=prompt&t=request): A query parameter that is used in requests to `/authorize` * `consent_method` (this property) * [consent](/openapi/okta-management/management/tag/AuthorizationServerScopes/#tag/AuthorizationServerScopes/operation/createOAuth2Scope!path=consent&t=request): A [Scope](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/AuthorizationServerScopes/) property that allows you to enable or disable user consent for an individual scope  | `prompt` | `consent_method` | `consent` | Result | ---------- | ----------- | ---------- | ----------- | | CONSENT | TRUSTED or REQUIRED | REQUIRED | Prompted | | CONSENT | TRUSTED or REQUIRED | FLEXIBLE | Prompted | | CONSENT | TRUSTED | IMPLICIT | Not prompted | | NONE | TRUSTED | FLEXIBLE, IMPLICIT, or REQUIRED | Not prompted | | NONE | REQUIRED | FLEXIBLE or REQUIRED | Prompted | | NONE | REQUIRED | IMPLICIT | Not prompted |  > **Notes:** > * If you request a scope that requires consent while using the `client_credentials` flow, an error is returned because the flow doesn't support user consent. > * If the `prompt` value is set to `NONE`, but the `consent_method` and the consent values are set to `REQUIRED`, then an error occurs. > * When a scope is requested during a Client Credentials grant flow and `consent` is set to `FLEXIBLE`, the scope is granted in the access token with no consent prompt. This occurs because there is no user involved in a two-legged OAuth 2.0 [Client Credentials](https://developer.okta.com/docs/guides/implement-grant-type/clientcreds/main/) grant flow.
	ConsentMethod *string `json:"consent_method,omitempty"`
	// Indicates that the client application uses Demonstrating Proof-of-Possession (DPoP) for token requests. If `true`, the authorization server rejects token requests from this client that don't contain the DPoP header. > **Note:** If `dpop_bound_access_tokens` is true, then `client_credentials` and `implicit` aren't allowed in `grant_types`.
	DpopBoundAccessTokens *bool `json:"dpop_bound_access_tokens,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle> <x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Determines whether Okta sends `sid` and `iss` in the logout request
	FrontchannelLogoutSessionRequired *bool `json:"frontchannel_logout_session_required,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle> <x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>URL where Okta sends the logout request
	FrontchannelLogoutUri *string  `json:"frontchannel_logout_uri,omitempty"`
	GrantTypes            []string `json:"grant_types"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>The algorithm for encrypting access tokens issued by this authorization server. If this is requested, the response is signed, and then encrypted. The result is a nested JWT. The default, if omitted, is that no encryption is performed.
	IdTokenEncryptedResponseAlg *string                                    `json:"id_token_encrypted_response_alg,omitempty"`
	IdpInitiatedLogin           *OpenIdConnectApplicationIdpInitiatedLogin `json:"idp_initiated_login,omitempty"`
	// URL string that a third party can use to initiate the sign-in flow by the client
	InitiateLoginUri *string `json:"initiate_login_uri,omitempty"`
	// Indicates whether the Okta authorization server uses the original Okta org domain URL or a custom domain URL as the issuer of the ID token for this client
	IssuerMode *string                                     `json:"issuer_mode,omitempty"`
	Jwks       *OpenIdConnectApplicationSettingsClientKeys `json:"jwks,omitempty"`
	// URL string that references a JSON Web Key Set for validating JWTs presented to Okta or for encrypting ID tokens minted by Okta for the client
	JwksUri *string `json:"jwks_uri,omitempty"`
	// The URL string that references a logo for the client. This logo appears on the client tile in the End-User Dashboard. It also appears on the client consent dialog during the client consent flow.
	LogoUri *string                          `json:"logo_uri,omitempty"`
	Network *OpenIdConnectApplicationNetwork `json:"network,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle> <x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Allows the app to participate in front-channel Single Logout  > **Note:** You can only enable `participate_slo` for `web` and `browser` application types (`application_type`).
	ParticipateSlo *bool `json:"participate_slo,omitempty"`
	// URL string of a web page providing the client's policy document
	PolicyUri *string `json:"policy_uri,omitempty"`
	// Array of redirection URI strings for relying party-initiated logouts
	PostLogoutRedirectUris []string `json:"post_logout_redirect_uris,omitempty"`
	// Array of redirection URI strings for use in redirect-based flows. > **Note:** At least one `redirect_uris` and `response_types` are required for all client types, with exceptions: if the client uses the [Resource Owner Password ](https://tools.ietf.org/html/rfc6749#section-4.3)flow (`grant_types` contains `password`) or [Client Credentials](https://tools.ietf.org/html/rfc6749#section-4.4)flow (`grant_types` contains `client_credentials`), then no `redirect_uris` or `response_types` is necessary. In these cases, you can pass either null or an empty array for these attributes.
	RedirectUris []string                                      `json:"redirect_uris,omitempty"`
	RefreshToken *OpenIdConnectApplicationSettingsRefreshToken `json:"refresh_token,omitempty"`
	// The type of JSON Web Key Set (JWKS) algorithm that must be used for signing request objects
	RequestObjectSigningAlg *string `json:"request_object_signing_alg,omitempty"`
	// Array of OAuth 2.0 response type strings
	ResponseTypes []string `json:"response_types,omitempty"`
	// The sector identifier used for pairwise `subject_type`. See [OIDC Pairwise Identifier Algorithm](https://openid.net/specs/openid-connect-messages-1_0-20.html#idtype.pairwise.alg)
	SectorIdentifierUri *string `json:"sector_identifier_uri,omitempty"`
	// Type of the subject
	SubjectType *string `json:"subject_type,omitempty"`
	// URL string of a web page providing the client's terms of service document
	TosUri *string `json:"tos_uri,omitempty"`
	// Indicates if the client is allowed to use wildcard matching of `redirect_uris`
	WildcardRedirect     *string `json:"wildcard_redirect,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenIdConnectApplicationSettingsClient OpenIdConnectApplicationSettingsClient

// NewOpenIdConnectApplicationSettingsClient instantiates a new OpenIdConnectApplicationSettingsClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectApplicationSettingsClient(grantTypes []string) *OpenIdConnectApplicationSettingsClient {
	this := OpenIdConnectApplicationSettingsClient{}
	var consentMethod string = "TRUSTED"
	this.ConsentMethod = &consentMethod
	var dpopBoundAccessTokens bool = false
	this.DpopBoundAccessTokens = &dpopBoundAccessTokens
	this.GrantTypes = grantTypes
	return &this
}

// NewOpenIdConnectApplicationSettingsClientWithDefaults instantiates a new OpenIdConnectApplicationSettingsClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIdConnectApplicationSettingsClientWithDefaults() *OpenIdConnectApplicationSettingsClient {
	this := OpenIdConnectApplicationSettingsClient{}
	var consentMethod string = "TRUSTED"
	this.ConsentMethod = &consentMethod
	var dpopBoundAccessTokens bool = false
	this.DpopBoundAccessTokens = &dpopBoundAccessTokens
	return &this
}

// GetApplicationType returns the ApplicationType field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetApplicationType() string {
	if o == nil || IsNil(o.ApplicationType) {
		var ret string
		return ret
	}
	return *o.ApplicationType
}

// GetApplicationTypeOk returns a tuple with the ApplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetApplicationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationType) {
		return nil, false
	}
	return o.ApplicationType, true
}

// HasApplicationType returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasApplicationType() bool {
	if o != nil && !IsNil(o.ApplicationType) {
		return true
	}

	return false
}

// SetApplicationType gets a reference to the given string and assigns it to the ApplicationType field.
func (o *OpenIdConnectApplicationSettingsClient) SetApplicationType(v string) {
	o.ApplicationType = &v
}

// GetBackchannelAuthenticationRequestSigningAlg returns the BackchannelAuthenticationRequestSigningAlg field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelAuthenticationRequestSigningAlg() string {
	if o == nil || IsNil(o.BackchannelAuthenticationRequestSigningAlg) {
		var ret string
		return ret
	}
	return *o.BackchannelAuthenticationRequestSigningAlg
}

// GetBackchannelAuthenticationRequestSigningAlgOk returns a tuple with the BackchannelAuthenticationRequestSigningAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelAuthenticationRequestSigningAlgOk() (*string, bool) {
	if o == nil || IsNil(o.BackchannelAuthenticationRequestSigningAlg) {
		return nil, false
	}
	return o.BackchannelAuthenticationRequestSigningAlg, true
}

// HasBackchannelAuthenticationRequestSigningAlg returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasBackchannelAuthenticationRequestSigningAlg() bool {
	if o != nil && !IsNil(o.BackchannelAuthenticationRequestSigningAlg) {
		return true
	}

	return false
}

// SetBackchannelAuthenticationRequestSigningAlg gets a reference to the given string and assigns it to the BackchannelAuthenticationRequestSigningAlg field.
func (o *OpenIdConnectApplicationSettingsClient) SetBackchannelAuthenticationRequestSigningAlg(v string) {
	o.BackchannelAuthenticationRequestSigningAlg = &v
}

// GetBackchannelCustomAuthenticatorId returns the BackchannelCustomAuthenticatorId field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelCustomAuthenticatorId() string {
	if o == nil || IsNil(o.BackchannelCustomAuthenticatorId) {
		var ret string
		return ret
	}
	return *o.BackchannelCustomAuthenticatorId
}

// GetBackchannelCustomAuthenticatorIdOk returns a tuple with the BackchannelCustomAuthenticatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelCustomAuthenticatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.BackchannelCustomAuthenticatorId) {
		return nil, false
	}
	return o.BackchannelCustomAuthenticatorId, true
}

// HasBackchannelCustomAuthenticatorId returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasBackchannelCustomAuthenticatorId() bool {
	if o != nil && !IsNil(o.BackchannelCustomAuthenticatorId) {
		return true
	}

	return false
}

// SetBackchannelCustomAuthenticatorId gets a reference to the given string and assigns it to the BackchannelCustomAuthenticatorId field.
func (o *OpenIdConnectApplicationSettingsClient) SetBackchannelCustomAuthenticatorId(v string) {
	o.BackchannelCustomAuthenticatorId = &v
}

// GetBackchannelTokenDeliveryMode returns the BackchannelTokenDeliveryMode field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelTokenDeliveryMode() string {
	if o == nil || IsNil(o.BackchannelTokenDeliveryMode) {
		var ret string
		return ret
	}
	return *o.BackchannelTokenDeliveryMode
}

// GetBackchannelTokenDeliveryModeOk returns a tuple with the BackchannelTokenDeliveryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelTokenDeliveryModeOk() (*string, bool) {
	if o == nil || IsNil(o.BackchannelTokenDeliveryMode) {
		return nil, false
	}
	return o.BackchannelTokenDeliveryMode, true
}

// HasBackchannelTokenDeliveryMode returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasBackchannelTokenDeliveryMode() bool {
	if o != nil && !IsNil(o.BackchannelTokenDeliveryMode) {
		return true
	}

	return false
}

// SetBackchannelTokenDeliveryMode gets a reference to the given string and assigns it to the BackchannelTokenDeliveryMode field.
func (o *OpenIdConnectApplicationSettingsClient) SetBackchannelTokenDeliveryMode(v string) {
	o.BackchannelTokenDeliveryMode = &v
}

// GetClientUri returns the ClientUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetClientUri() string {
	if o == nil || IsNil(o.ClientUri) {
		var ret string
		return ret
	}
	return *o.ClientUri
}

// GetClientUriOk returns a tuple with the ClientUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetClientUriOk() (*string, bool) {
	if o == nil || IsNil(o.ClientUri) {
		return nil, false
	}
	return o.ClientUri, true
}

// HasClientUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasClientUri() bool {
	if o != nil && !IsNil(o.ClientUri) {
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
	if o == nil || IsNil(o.ConsentMethod) {
		var ret string
		return ret
	}
	return *o.ConsentMethod
}

// GetConsentMethodOk returns a tuple with the ConsentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetConsentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.ConsentMethod) {
		return nil, false
	}
	return o.ConsentMethod, true
}

// HasConsentMethod returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasConsentMethod() bool {
	if o != nil && !IsNil(o.ConsentMethod) {
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
	if o == nil || IsNil(o.DpopBoundAccessTokens) {
		var ret bool
		return ret
	}
	return *o.DpopBoundAccessTokens
}

// GetDpopBoundAccessTokensOk returns a tuple with the DpopBoundAccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetDpopBoundAccessTokensOk() (*bool, bool) {
	if o == nil || IsNil(o.DpopBoundAccessTokens) {
		return nil, false
	}
	return o.DpopBoundAccessTokens, true
}

// HasDpopBoundAccessTokens returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasDpopBoundAccessTokens() bool {
	if o != nil && !IsNil(o.DpopBoundAccessTokens) {
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
	if o == nil || IsNil(o.FrontchannelLogoutSessionRequired) {
		var ret bool
		return ret
	}
	return *o.FrontchannelLogoutSessionRequired
}

// GetFrontchannelLogoutSessionRequiredOk returns a tuple with the FrontchannelLogoutSessionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutSessionRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.FrontchannelLogoutSessionRequired) {
		return nil, false
	}
	return o.FrontchannelLogoutSessionRequired, true
}

// HasFrontchannelLogoutSessionRequired returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasFrontchannelLogoutSessionRequired() bool {
	if o != nil && !IsNil(o.FrontchannelLogoutSessionRequired) {
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
	if o == nil || IsNil(o.FrontchannelLogoutUri) {
		var ret string
		return ret
	}
	return *o.FrontchannelLogoutUri
}

// GetFrontchannelLogoutUriOk returns a tuple with the FrontchannelLogoutUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutUriOk() (*string, bool) {
	if o == nil || IsNil(o.FrontchannelLogoutUri) {
		return nil, false
	}
	return o.FrontchannelLogoutUri, true
}

// HasFrontchannelLogoutUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasFrontchannelLogoutUri() bool {
	if o != nil && !IsNil(o.FrontchannelLogoutUri) {
		return true
	}

	return false
}

// SetFrontchannelLogoutUri gets a reference to the given string and assigns it to the FrontchannelLogoutUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetFrontchannelLogoutUri(v string) {
	o.FrontchannelLogoutUri = &v
}

// GetGrantTypes returns the GrantTypes field value
func (o *OpenIdConnectApplicationSettingsClient) GetGrantTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetGrantTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantTypes, true
}

// SetGrantTypes sets field value
func (o *OpenIdConnectApplicationSettingsClient) SetGrantTypes(v []string) {
	o.GrantTypes = v
}

// GetIdTokenEncryptedResponseAlg returns the IdTokenEncryptedResponseAlg field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetIdTokenEncryptedResponseAlg() string {
	if o == nil || IsNil(o.IdTokenEncryptedResponseAlg) {
		var ret string
		return ret
	}
	return *o.IdTokenEncryptedResponseAlg
}

// GetIdTokenEncryptedResponseAlgOk returns a tuple with the IdTokenEncryptedResponseAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetIdTokenEncryptedResponseAlgOk() (*string, bool) {
	if o == nil || IsNil(o.IdTokenEncryptedResponseAlg) {
		return nil, false
	}
	return o.IdTokenEncryptedResponseAlg, true
}

// HasIdTokenEncryptedResponseAlg returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasIdTokenEncryptedResponseAlg() bool {
	if o != nil && !IsNil(o.IdTokenEncryptedResponseAlg) {
		return true
	}

	return false
}

// SetIdTokenEncryptedResponseAlg gets a reference to the given string and assigns it to the IdTokenEncryptedResponseAlg field.
func (o *OpenIdConnectApplicationSettingsClient) SetIdTokenEncryptedResponseAlg(v string) {
	o.IdTokenEncryptedResponseAlg = &v
}

// GetIdpInitiatedLogin returns the IdpInitiatedLogin field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetIdpInitiatedLogin() OpenIdConnectApplicationIdpInitiatedLogin {
	if o == nil || IsNil(o.IdpInitiatedLogin) {
		var ret OpenIdConnectApplicationIdpInitiatedLogin
		return ret
	}
	return *o.IdpInitiatedLogin
}

// GetIdpInitiatedLoginOk returns a tuple with the IdpInitiatedLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetIdpInitiatedLoginOk() (*OpenIdConnectApplicationIdpInitiatedLogin, bool) {
	if o == nil || IsNil(o.IdpInitiatedLogin) {
		return nil, false
	}
	return o.IdpInitiatedLogin, true
}

// HasIdpInitiatedLogin returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasIdpInitiatedLogin() bool {
	if o != nil && !IsNil(o.IdpInitiatedLogin) {
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
	if o == nil || IsNil(o.InitiateLoginUri) {
		var ret string
		return ret
	}
	return *o.InitiateLoginUri
}

// GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetInitiateLoginUriOk() (*string, bool) {
	if o == nil || IsNil(o.InitiateLoginUri) {
		return nil, false
	}
	return o.InitiateLoginUri, true
}

// HasInitiateLoginUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasInitiateLoginUri() bool {
	if o != nil && !IsNil(o.InitiateLoginUri) {
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
	if o == nil || IsNil(o.IssuerMode) {
		var ret string
		return ret
	}
	return *o.IssuerMode
}

// GetIssuerModeOk returns a tuple with the IssuerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetIssuerModeOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerMode) {
		return nil, false
	}
	return o.IssuerMode, true
}

// HasIssuerMode returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasIssuerMode() bool {
	if o != nil && !IsNil(o.IssuerMode) {
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
	if o == nil || IsNil(o.Jwks) {
		var ret OpenIdConnectApplicationSettingsClientKeys
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetJwksOk() (*OpenIdConnectApplicationSettingsClientKeys, bool) {
	if o == nil || IsNil(o.Jwks) {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasJwks() bool {
	if o != nil && !IsNil(o.Jwks) {
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
	if o == nil || IsNil(o.JwksUri) {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetJwksUriOk() (*string, bool) {
	if o == nil || IsNil(o.JwksUri) {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasJwksUri() bool {
	if o != nil && !IsNil(o.JwksUri) {
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
	if o == nil || IsNil(o.LogoUri) {
		var ret string
		return ret
	}
	return *o.LogoUri
}

// GetLogoUriOk returns a tuple with the LogoUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetLogoUriOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUri) {
		return nil, false
	}
	return o.LogoUri, true
}

// HasLogoUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasLogoUri() bool {
	if o != nil && !IsNil(o.LogoUri) {
		return true
	}

	return false
}

// SetLogoUri gets a reference to the given string and assigns it to the LogoUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetLogoUri(v string) {
	o.LogoUri = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetNetwork() OpenIdConnectApplicationNetwork {
	if o == nil || IsNil(o.Network) {
		var ret OpenIdConnectApplicationNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetNetworkOk() (*OpenIdConnectApplicationNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OpenIdConnectApplicationNetwork and assigns it to the Network field.
func (o *OpenIdConnectApplicationSettingsClient) SetNetwork(v OpenIdConnectApplicationNetwork) {
	o.Network = &v
}

// GetParticipateSlo returns the ParticipateSlo field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetParticipateSlo() bool {
	if o == nil || IsNil(o.ParticipateSlo) {
		var ret bool
		return ret
	}
	return *o.ParticipateSlo
}

// GetParticipateSloOk returns a tuple with the ParticipateSlo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetParticipateSloOk() (*bool, bool) {
	if o == nil || IsNil(o.ParticipateSlo) {
		return nil, false
	}
	return o.ParticipateSlo, true
}

// HasParticipateSlo returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasParticipateSlo() bool {
	if o != nil && !IsNil(o.ParticipateSlo) {
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
	if o == nil || IsNil(o.PolicyUri) {
		var ret string
		return ret
	}
	return *o.PolicyUri
}

// GetPolicyUriOk returns a tuple with the PolicyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetPolicyUriOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyUri) {
		return nil, false
	}
	return o.PolicyUri, true
}

// HasPolicyUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasPolicyUri() bool {
	if o != nil && !IsNil(o.PolicyUri) {
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
	if o == nil || IsNil(o.PostLogoutRedirectUris) {
		var ret []string
		return ret
	}
	return o.PostLogoutRedirectUris
}

// GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetPostLogoutRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.PostLogoutRedirectUris) {
		return nil, false
	}
	return o.PostLogoutRedirectUris, true
}

// HasPostLogoutRedirectUris returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasPostLogoutRedirectUris() bool {
	if o != nil && !IsNil(o.PostLogoutRedirectUris) {
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
	if o == nil || IsNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasRedirectUris() bool {
	if o != nil && !IsNil(o.RedirectUris) {
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
	if o == nil || IsNil(o.RefreshToken) {
		var ret OpenIdConnectApplicationSettingsRefreshToken
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetRefreshTokenOk() (*OpenIdConnectApplicationSettingsRefreshToken, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given OpenIdConnectApplicationSettingsRefreshToken and assigns it to the RefreshToken field.
func (o *OpenIdConnectApplicationSettingsClient) SetRefreshToken(v OpenIdConnectApplicationSettingsRefreshToken) {
	o.RefreshToken = &v
}

// GetRequestObjectSigningAlg returns the RequestObjectSigningAlg field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetRequestObjectSigningAlg() string {
	if o == nil || IsNil(o.RequestObjectSigningAlg) {
		var ret string
		return ret
	}
	return *o.RequestObjectSigningAlg
}

// GetRequestObjectSigningAlgOk returns a tuple with the RequestObjectSigningAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetRequestObjectSigningAlgOk() (*string, bool) {
	if o == nil || IsNil(o.RequestObjectSigningAlg) {
		return nil, false
	}
	return o.RequestObjectSigningAlg, true
}

// HasRequestObjectSigningAlg returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasRequestObjectSigningAlg() bool {
	if o != nil && !IsNil(o.RequestObjectSigningAlg) {
		return true
	}

	return false
}

// SetRequestObjectSigningAlg gets a reference to the given string and assigns it to the RequestObjectSigningAlg field.
func (o *OpenIdConnectApplicationSettingsClient) SetRequestObjectSigningAlg(v string) {
	o.RequestObjectSigningAlg = &v
}

// GetResponseTypes returns the ResponseTypes field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetResponseTypes() []string {
	if o == nil || IsNil(o.ResponseTypes) {
		var ret []string
		return ret
	}
	return o.ResponseTypes
}

// GetResponseTypesOk returns a tuple with the ResponseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetResponseTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ResponseTypes) {
		return nil, false
	}
	return o.ResponseTypes, true
}

// HasResponseTypes returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasResponseTypes() bool {
	if o != nil && !IsNil(o.ResponseTypes) {
		return true
	}

	return false
}

// SetResponseTypes gets a reference to the given []string and assigns it to the ResponseTypes field.
func (o *OpenIdConnectApplicationSettingsClient) SetResponseTypes(v []string) {
	o.ResponseTypes = v
}

// GetSectorIdentifierUri returns the SectorIdentifierUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetSectorIdentifierUri() string {
	if o == nil || IsNil(o.SectorIdentifierUri) {
		var ret string
		return ret
	}
	return *o.SectorIdentifierUri
}

// GetSectorIdentifierUriOk returns a tuple with the SectorIdentifierUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetSectorIdentifierUriOk() (*string, bool) {
	if o == nil || IsNil(o.SectorIdentifierUri) {
		return nil, false
	}
	return o.SectorIdentifierUri, true
}

// HasSectorIdentifierUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasSectorIdentifierUri() bool {
	if o != nil && !IsNil(o.SectorIdentifierUri) {
		return true
	}

	return false
}

// SetSectorIdentifierUri gets a reference to the given string and assigns it to the SectorIdentifierUri field.
func (o *OpenIdConnectApplicationSettingsClient) SetSectorIdentifierUri(v string) {
	o.SectorIdentifierUri = &v
}

// GetSubjectType returns the SubjectType field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetSubjectType() string {
	if o == nil || IsNil(o.SubjectType) {
		var ret string
		return ret
	}
	return *o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetSubjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectType) {
		return nil, false
	}
	return o.SubjectType, true
}

// HasSubjectType returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasSubjectType() bool {
	if o != nil && !IsNil(o.SubjectType) {
		return true
	}

	return false
}

// SetSubjectType gets a reference to the given string and assigns it to the SubjectType field.
func (o *OpenIdConnectApplicationSettingsClient) SetSubjectType(v string) {
	o.SubjectType = &v
}

// GetTosUri returns the TosUri field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsClient) GetTosUri() string {
	if o == nil || IsNil(o.TosUri) {
		var ret string
		return ret
	}
	return *o.TosUri
}

// GetTosUriOk returns a tuple with the TosUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetTosUriOk() (*string, bool) {
	if o == nil || IsNil(o.TosUri) {
		return nil, false
	}
	return o.TosUri, true
}

// HasTosUri returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasTosUri() bool {
	if o != nil && !IsNil(o.TosUri) {
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
	if o == nil || IsNil(o.WildcardRedirect) {
		var ret string
		return ret
	}
	return *o.WildcardRedirect
}

// GetWildcardRedirectOk returns a tuple with the WildcardRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsClient) GetWildcardRedirectOk() (*string, bool) {
	if o == nil || IsNil(o.WildcardRedirect) {
		return nil, false
	}
	return o.WildcardRedirect, true
}

// HasWildcardRedirect returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsClient) HasWildcardRedirect() bool {
	if o != nil && !IsNil(o.WildcardRedirect) {
		return true
	}

	return false
}

// SetWildcardRedirect gets a reference to the given string and assigns it to the WildcardRedirect field.
func (o *OpenIdConnectApplicationSettingsClient) SetWildcardRedirect(v string) {
	o.WildcardRedirect = &v
}

func (o OpenIdConnectApplicationSettingsClient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIdConnectApplicationSettingsClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationType) {
		toSerialize["application_type"] = o.ApplicationType
	}
	if !IsNil(o.BackchannelAuthenticationRequestSigningAlg) {
		toSerialize["backchannel_authentication_request_signing_alg"] = o.BackchannelAuthenticationRequestSigningAlg
	}
	if !IsNil(o.BackchannelCustomAuthenticatorId) {
		toSerialize["backchannel_custom_authenticator_id"] = o.BackchannelCustomAuthenticatorId
	}
	if !IsNil(o.BackchannelTokenDeliveryMode) {
		toSerialize["backchannel_token_delivery_mode"] = o.BackchannelTokenDeliveryMode
	}
	if !IsNil(o.ClientUri) {
		toSerialize["client_uri"] = o.ClientUri
	}
	if !IsNil(o.ConsentMethod) {
		toSerialize["consent_method"] = o.ConsentMethod
	}
	if !IsNil(o.DpopBoundAccessTokens) {
		toSerialize["dpop_bound_access_tokens"] = o.DpopBoundAccessTokens
	}
	if !IsNil(o.FrontchannelLogoutSessionRequired) {
		toSerialize["frontchannel_logout_session_required"] = o.FrontchannelLogoutSessionRequired
	}
	if !IsNil(o.FrontchannelLogoutUri) {
		toSerialize["frontchannel_logout_uri"] = o.FrontchannelLogoutUri
	}
	toSerialize["grant_types"] = o.GrantTypes
	if !IsNil(o.IdTokenEncryptedResponseAlg) {
		toSerialize["id_token_encrypted_response_alg"] = o.IdTokenEncryptedResponseAlg
	}
	if !IsNil(o.IdpInitiatedLogin) {
		toSerialize["idp_initiated_login"] = o.IdpInitiatedLogin
	}
	if !IsNil(o.InitiateLoginUri) {
		toSerialize["initiate_login_uri"] = o.InitiateLoginUri
	}
	if !IsNil(o.IssuerMode) {
		toSerialize["issuer_mode"] = o.IssuerMode
	}
	if !IsNil(o.Jwks) {
		toSerialize["jwks"] = o.Jwks
	}
	if !IsNil(o.JwksUri) {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if !IsNil(o.LogoUri) {
		toSerialize["logo_uri"] = o.LogoUri
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ParticipateSlo) {
		toSerialize["participate_slo"] = o.ParticipateSlo
	}
	if !IsNil(o.PolicyUri) {
		toSerialize["policy_uri"] = o.PolicyUri
	}
	if !IsNil(o.PostLogoutRedirectUris) {
		toSerialize["post_logout_redirect_uris"] = o.PostLogoutRedirectUris
	}
	if !IsNil(o.RedirectUris) {
		toSerialize["redirect_uris"] = o.RedirectUris
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if !IsNil(o.RequestObjectSigningAlg) {
		toSerialize["request_object_signing_alg"] = o.RequestObjectSigningAlg
	}
	if !IsNil(o.ResponseTypes) {
		toSerialize["response_types"] = o.ResponseTypes
	}
	if !IsNil(o.SectorIdentifierUri) {
		toSerialize["sector_identifier_uri"] = o.SectorIdentifierUri
	}
	if !IsNil(o.SubjectType) {
		toSerialize["subject_type"] = o.SubjectType
	}
	if !IsNil(o.TosUri) {
		toSerialize["tos_uri"] = o.TosUri
	}
	if !IsNil(o.WildcardRedirect) {
		toSerialize["wildcard_redirect"] = o.WildcardRedirect
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenIdConnectApplicationSettingsClient) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grant_types",
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

	varOpenIdConnectApplicationSettingsClient := _OpenIdConnectApplicationSettingsClient{}

	err = json.Unmarshal(data, &varOpenIdConnectApplicationSettingsClient)

	if err != nil {
		return err
	}

	*o = OpenIdConnectApplicationSettingsClient(varOpenIdConnectApplicationSettingsClient)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "application_type")
		delete(additionalProperties, "backchannel_authentication_request_signing_alg")
		delete(additionalProperties, "backchannel_custom_authenticator_id")
		delete(additionalProperties, "backchannel_token_delivery_mode")
		delete(additionalProperties, "client_uri")
		delete(additionalProperties, "consent_method")
		delete(additionalProperties, "dpop_bound_access_tokens")
		delete(additionalProperties, "frontchannel_logout_session_required")
		delete(additionalProperties, "frontchannel_logout_uri")
		delete(additionalProperties, "grant_types")
		delete(additionalProperties, "id_token_encrypted_response_alg")
		delete(additionalProperties, "idp_initiated_login")
		delete(additionalProperties, "initiate_login_uri")
		delete(additionalProperties, "issuer_mode")
		delete(additionalProperties, "jwks")
		delete(additionalProperties, "jwks_uri")
		delete(additionalProperties, "logo_uri")
		delete(additionalProperties, "network")
		delete(additionalProperties, "participate_slo")
		delete(additionalProperties, "policy_uri")
		delete(additionalProperties, "post_logout_redirect_uris")
		delete(additionalProperties, "redirect_uris")
		delete(additionalProperties, "refresh_token")
		delete(additionalProperties, "request_object_signing_alg")
		delete(additionalProperties, "response_types")
		delete(additionalProperties, "sector_identifier_uri")
		delete(additionalProperties, "subject_type")
		delete(additionalProperties, "tos_uri")
		delete(additionalProperties, "wildcard_redirect")
		o.AdditionalProperties = additionalProperties
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
