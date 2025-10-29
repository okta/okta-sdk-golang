# OpenIdConnectApplicationSettingsClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationType** | Pointer to **string** | The type of client app Specific &#x60;grant_types&#x60; are valid for each &#x60;application_type&#x60;. See [Create a Client Application](/openapi/okta-oauth/oauth/tag/Client/#tag/Client/operation/createClient). | [optional] 
**BackchannelAuthenticationRequestSigningAlg** | Pointer to **string** | The signing algorithm for Client-Initiated Backchannel Authentication (CIBA) signed requests using JWT. If this value isn&#39;t set and a JWT-signed request is sent, the request fails. &gt; **Note:** This property appears for clients with &#x60;urn:openid:params:grant-type:ciba&#x60; defined as one of the &#x60;grant_types&#x60;.  | [optional] 
**BackchannelCustomAuthenticatorId** | Pointer to **string** | The ID of the custom authenticator that authenticates the user &gt; **Note:** This property appears for clients with &#x60;urn:openid:params:grant-type:ciba&#x60; defined as one of the &#x60;grant_types&#x60;.  | [optional] 
**BackchannelTokenDeliveryMode** | Pointer to **string** | The delivery mode for Client-Initiated Backchannel Authentication (CIBA).  Currently, only &#x60;poll&#x60; is supported. &gt; **Note:** This property appears for clients with &#x60;urn:openid:params:grant-type:ciba&#x60; defined as one of the &#x60;grant_types&#x60;.  | [optional] 
**ClientUri** | Pointer to **string** | URL string of a web page providing information about the client | [optional] 
**ConsentMethod** | Pointer to **string** | Indicates whether user consent is required or implicit. A consent dialog appears for the end user depending on the values of three elements:  * [prompt](/openapi/okta-oauth/oauth/tag/OrgAS/#tag/OrgAS/operation/authorize!in&#x3D;query&amp;path&#x3D;prompt&amp;t&#x3D;request): A query parameter that is used in requests to &#x60;/authorize&#x60; * &#x60;consent_method&#x60; (this property) * [consent](/openapi/okta-management/management/tag/AuthorizationServerScopes/#tag/AuthorizationServerScopes/operation/createOAuth2Scope!path&#x3D;consent&amp;t&#x3D;request): A [Scope](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/AuthorizationServerScopes/) property that allows you to enable or disable user consent for an individual scope  | &#x60;prompt&#x60; | &#x60;consent_method&#x60; | &#x60;consent&#x60; | Result | ---------- | ----------- | ---------- | ----------- | | CONSENT | TRUSTED or REQUIRED | REQUIRED | Prompted | | CONSENT | TRUSTED or REQUIRED | FLEXIBLE | Prompted | | CONSENT | TRUSTED | IMPLICIT | Not prompted | | NONE | TRUSTED | FLEXIBLE, IMPLICIT, or REQUIRED | Not prompted | | NONE | REQUIRED | FLEXIBLE or REQUIRED | Prompted | | NONE | REQUIRED | IMPLICIT | Not prompted |  &gt; **Notes:** &gt; * If you request a scope that requires consent while using the &#x60;client_credentials&#x60; flow, an error is returned because the flow doesn&#39;t support user consent. &gt; * If the &#x60;prompt&#x60; value is set to &#x60;NONE&#x60;, but the &#x60;consent_method&#x60; and the consent values are set to &#x60;REQUIRED&#x60;, then an error occurs. &gt; * When a scope is requested during a Client Credentials grant flow and &#x60;consent&#x60; is set to &#x60;FLEXIBLE&#x60;, the scope is granted in the access token with no consent prompt. This occurs because there is no user involved in a two-legged OAuth 2.0 [Client Credentials](https://developer.okta.com/docs/guides/implement-grant-type/clientcreds/main/) grant flow.  | [optional] [default to "TRUSTED"]
**DpopBoundAccessTokens** | Pointer to **bool** | Indicates that the client application uses Demonstrating Proof-of-Possession (DPoP) for token requests. If &#x60;true&#x60;, the authorization server rejects token requests from this client that don&#39;t contain the DPoP header. &gt; **Note:** If &#x60;dpop_bound_access_tokens&#x60; is true, then &#x60;client_credentials&#x60; and &#x60;implicit&#x60; aren&#39;t allowed in &#x60;grant_types&#x60;.  | [optional] [default to false]
**FrontchannelLogoutSessionRequired** | Pointer to **bool** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Determines whether Okta sends &#x60;sid&#x60; and &#x60;iss&#x60; in the logout request | [optional] 
**FrontchannelLogoutUri** | Pointer to **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;URL where Okta sends the logout request | [optional] 
**GrantTypes** | **[]string** |  | 
**IdTokenEncryptedResponseAlg** | Pointer to **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;The algorithm for encrypting access tokens issued by this authorization server. If this is requested, the response is signed, and then encrypted. The result is a nested JWT. The default, if omitted, is that no encryption is performed. | [optional] 
**IdpInitiatedLogin** | Pointer to [**OpenIdConnectApplicationIdpInitiatedLogin**](OpenIdConnectApplicationIdpInitiatedLogin.md) |  | [optional] 
**InitiateLoginUri** | Pointer to **string** | URL string that a third party can use to initiate the sign-in flow by the client | [optional] 
**IssuerMode** | Pointer to **string** | Indicates whether the Okta authorization server uses the original Okta org domain URL or a custom domain URL as the issuer of the ID token for this client | [optional] 
**Jwks** | Pointer to [**OpenIdConnectApplicationSettingsClientKeys**](OpenIdConnectApplicationSettingsClientKeys.md) |  | [optional] 
**JwksUri** | Pointer to **string** | URL string that references a JSON Web Key Set for validating JWTs presented to Okta or for encrypting ID tokens minted by Okta for the client | [optional] 
**LogoUri** | Pointer to **string** | The URL string that references a logo for the client. This logo appears on the client tile in the End-User Dashboard. It also appears on the client consent dialog during the client consent flow. | [optional] 
**Network** | Pointer to [**OpenIdConnectApplicationNetwork**](OpenIdConnectApplicationNetwork.md) |  | [optional] 
**ParticipateSlo** | Pointer to **bool** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Allows the app to participate in front-channel Single Logout  &gt; **Note:** You can only enable &#x60;participate_slo&#x60; for &#x60;web&#x60; and &#x60;browser&#x60; application types (&#x60;application_type&#x60;).  | [optional] 
**PolicyUri** | Pointer to **string** | URL string of a web page providing the client&#39;s policy document | [optional] 
**PostLogoutRedirectUris** | Pointer to **[]string** | Array of redirection URI strings for relying party-initiated logouts | [optional] 
**RedirectUris** | Pointer to **[]string** | Array of redirection URI strings for use in redirect-based flows. &gt; **Note:** At least one &#x60;redirect_uris&#x60; and &#x60;response_types&#x60; are required for all client types, with exceptions: if the client uses the [Resource Owner Password ](https://tools.ietf.org/html/rfc6749#section-4.3)flow (&#x60;grant_types&#x60; contains &#x60;password&#x60;) or [Client Credentials](https://tools.ietf.org/html/rfc6749#section-4.4)flow (&#x60;grant_types&#x60; contains &#x60;client_credentials&#x60;), then no &#x60;redirect_uris&#x60; or &#x60;response_types&#x60; is necessary. In these cases, you can pass either null or an empty array for these attributes. | [optional] 
**RefreshToken** | Pointer to [**OpenIdConnectApplicationSettingsRefreshToken**](OpenIdConnectApplicationSettingsRefreshToken.md) |  | [optional] 
**RequestObjectSigningAlg** | Pointer to **string** | The type of JSON Web Key Set (JWKS) algorithm that must be used for signing request objects | [optional] 
**ResponseTypes** | Pointer to **[]string** | Array of OAuth 2.0 response type strings | [optional] 
**SectorIdentifierUri** | Pointer to **string** | The sector identifier used for pairwise &#x60;subject_type&#x60;. See [OIDC Pairwise Identifier Algorithm](https://openid.net/specs/openid-connect-messages-1_0-20.html#idtype.pairwise.alg) | [optional] 
**SubjectType** | Pointer to **string** | Type of the subject | [optional] 
**TosUri** | Pointer to **string** | URL string of a web page providing the client&#39;s terms of service document | [optional] 
**WildcardRedirect** | Pointer to **string** | Indicates if the client is allowed to use wildcard matching of &#x60;redirect_uris&#x60; | [optional] 

## Methods

### NewOpenIdConnectApplicationSettingsClient

`func NewOpenIdConnectApplicationSettingsClient(grantTypes []string, ) *OpenIdConnectApplicationSettingsClient`

NewOpenIdConnectApplicationSettingsClient instantiates a new OpenIdConnectApplicationSettingsClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectApplicationSettingsClientWithDefaults

`func NewOpenIdConnectApplicationSettingsClientWithDefaults() *OpenIdConnectApplicationSettingsClient`

NewOpenIdConnectApplicationSettingsClientWithDefaults instantiates a new OpenIdConnectApplicationSettingsClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationType

`func (o *OpenIdConnectApplicationSettingsClient) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *OpenIdConnectApplicationSettingsClient) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *OpenIdConnectApplicationSettingsClient) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.

### HasApplicationType

`func (o *OpenIdConnectApplicationSettingsClient) HasApplicationType() bool`

HasApplicationType returns a boolean if a field has been set.

### GetBackchannelAuthenticationRequestSigningAlg

`func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelAuthenticationRequestSigningAlg() string`

GetBackchannelAuthenticationRequestSigningAlg returns the BackchannelAuthenticationRequestSigningAlg field if non-nil, zero value otherwise.

### GetBackchannelAuthenticationRequestSigningAlgOk

`func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelAuthenticationRequestSigningAlgOk() (*string, bool)`

GetBackchannelAuthenticationRequestSigningAlgOk returns a tuple with the BackchannelAuthenticationRequestSigningAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelAuthenticationRequestSigningAlg

`func (o *OpenIdConnectApplicationSettingsClient) SetBackchannelAuthenticationRequestSigningAlg(v string)`

SetBackchannelAuthenticationRequestSigningAlg sets BackchannelAuthenticationRequestSigningAlg field to given value.

### HasBackchannelAuthenticationRequestSigningAlg

`func (o *OpenIdConnectApplicationSettingsClient) HasBackchannelAuthenticationRequestSigningAlg() bool`

HasBackchannelAuthenticationRequestSigningAlg returns a boolean if a field has been set.

### GetBackchannelCustomAuthenticatorId

`func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelCustomAuthenticatorId() string`

GetBackchannelCustomAuthenticatorId returns the BackchannelCustomAuthenticatorId field if non-nil, zero value otherwise.

### GetBackchannelCustomAuthenticatorIdOk

`func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelCustomAuthenticatorIdOk() (*string, bool)`

GetBackchannelCustomAuthenticatorIdOk returns a tuple with the BackchannelCustomAuthenticatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelCustomAuthenticatorId

`func (o *OpenIdConnectApplicationSettingsClient) SetBackchannelCustomAuthenticatorId(v string)`

SetBackchannelCustomAuthenticatorId sets BackchannelCustomAuthenticatorId field to given value.

### HasBackchannelCustomAuthenticatorId

`func (o *OpenIdConnectApplicationSettingsClient) HasBackchannelCustomAuthenticatorId() bool`

HasBackchannelCustomAuthenticatorId returns a boolean if a field has been set.

### GetBackchannelTokenDeliveryMode

`func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelTokenDeliveryMode() string`

GetBackchannelTokenDeliveryMode returns the BackchannelTokenDeliveryMode field if non-nil, zero value otherwise.

### GetBackchannelTokenDeliveryModeOk

`func (o *OpenIdConnectApplicationSettingsClient) GetBackchannelTokenDeliveryModeOk() (*string, bool)`

GetBackchannelTokenDeliveryModeOk returns a tuple with the BackchannelTokenDeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelTokenDeliveryMode

`func (o *OpenIdConnectApplicationSettingsClient) SetBackchannelTokenDeliveryMode(v string)`

SetBackchannelTokenDeliveryMode sets BackchannelTokenDeliveryMode field to given value.

### HasBackchannelTokenDeliveryMode

`func (o *OpenIdConnectApplicationSettingsClient) HasBackchannelTokenDeliveryMode() bool`

HasBackchannelTokenDeliveryMode returns a boolean if a field has been set.

### GetClientUri

`func (o *OpenIdConnectApplicationSettingsClient) GetClientUri() string`

GetClientUri returns the ClientUri field if non-nil, zero value otherwise.

### GetClientUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetClientUriOk() (*string, bool)`

GetClientUriOk returns a tuple with the ClientUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUri

`func (o *OpenIdConnectApplicationSettingsClient) SetClientUri(v string)`

SetClientUri sets ClientUri field to given value.

### HasClientUri

`func (o *OpenIdConnectApplicationSettingsClient) HasClientUri() bool`

HasClientUri returns a boolean if a field has been set.

### GetConsentMethod

`func (o *OpenIdConnectApplicationSettingsClient) GetConsentMethod() string`

GetConsentMethod returns the ConsentMethod field if non-nil, zero value otherwise.

### GetConsentMethodOk

`func (o *OpenIdConnectApplicationSettingsClient) GetConsentMethodOk() (*string, bool)`

GetConsentMethodOk returns a tuple with the ConsentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentMethod

`func (o *OpenIdConnectApplicationSettingsClient) SetConsentMethod(v string)`

SetConsentMethod sets ConsentMethod field to given value.

### HasConsentMethod

`func (o *OpenIdConnectApplicationSettingsClient) HasConsentMethod() bool`

HasConsentMethod returns a boolean if a field has been set.

### GetDpopBoundAccessTokens

`func (o *OpenIdConnectApplicationSettingsClient) GetDpopBoundAccessTokens() bool`

GetDpopBoundAccessTokens returns the DpopBoundAccessTokens field if non-nil, zero value otherwise.

### GetDpopBoundAccessTokensOk

`func (o *OpenIdConnectApplicationSettingsClient) GetDpopBoundAccessTokensOk() (*bool, bool)`

GetDpopBoundAccessTokensOk returns a tuple with the DpopBoundAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopBoundAccessTokens

`func (o *OpenIdConnectApplicationSettingsClient) SetDpopBoundAccessTokens(v bool)`

SetDpopBoundAccessTokens sets DpopBoundAccessTokens field to given value.

### HasDpopBoundAccessTokens

`func (o *OpenIdConnectApplicationSettingsClient) HasDpopBoundAccessTokens() bool`

HasDpopBoundAccessTokens returns a boolean if a field has been set.

### GetFrontchannelLogoutSessionRequired

`func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutSessionRequired() bool`

GetFrontchannelLogoutSessionRequired returns the FrontchannelLogoutSessionRequired field if non-nil, zero value otherwise.

### GetFrontchannelLogoutSessionRequiredOk

`func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutSessionRequiredOk() (*bool, bool)`

GetFrontchannelLogoutSessionRequiredOk returns a tuple with the FrontchannelLogoutSessionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontchannelLogoutSessionRequired

`func (o *OpenIdConnectApplicationSettingsClient) SetFrontchannelLogoutSessionRequired(v bool)`

SetFrontchannelLogoutSessionRequired sets FrontchannelLogoutSessionRequired field to given value.

### HasFrontchannelLogoutSessionRequired

`func (o *OpenIdConnectApplicationSettingsClient) HasFrontchannelLogoutSessionRequired() bool`

HasFrontchannelLogoutSessionRequired returns a boolean if a field has been set.

### GetFrontchannelLogoutUri

`func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutUri() string`

GetFrontchannelLogoutUri returns the FrontchannelLogoutUri field if non-nil, zero value otherwise.

### GetFrontchannelLogoutUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetFrontchannelLogoutUriOk() (*string, bool)`

GetFrontchannelLogoutUriOk returns a tuple with the FrontchannelLogoutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontchannelLogoutUri

`func (o *OpenIdConnectApplicationSettingsClient) SetFrontchannelLogoutUri(v string)`

SetFrontchannelLogoutUri sets FrontchannelLogoutUri field to given value.

### HasFrontchannelLogoutUri

`func (o *OpenIdConnectApplicationSettingsClient) HasFrontchannelLogoutUri() bool`

HasFrontchannelLogoutUri returns a boolean if a field has been set.

### GetGrantTypes

`func (o *OpenIdConnectApplicationSettingsClient) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *OpenIdConnectApplicationSettingsClient) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *OpenIdConnectApplicationSettingsClient) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.


### GetIdTokenEncryptedResponseAlg

`func (o *OpenIdConnectApplicationSettingsClient) GetIdTokenEncryptedResponseAlg() string`

GetIdTokenEncryptedResponseAlg returns the IdTokenEncryptedResponseAlg field if non-nil, zero value otherwise.

### GetIdTokenEncryptedResponseAlgOk

`func (o *OpenIdConnectApplicationSettingsClient) GetIdTokenEncryptedResponseAlgOk() (*string, bool)`

GetIdTokenEncryptedResponseAlgOk returns a tuple with the IdTokenEncryptedResponseAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenEncryptedResponseAlg

`func (o *OpenIdConnectApplicationSettingsClient) SetIdTokenEncryptedResponseAlg(v string)`

SetIdTokenEncryptedResponseAlg sets IdTokenEncryptedResponseAlg field to given value.

### HasIdTokenEncryptedResponseAlg

`func (o *OpenIdConnectApplicationSettingsClient) HasIdTokenEncryptedResponseAlg() bool`

HasIdTokenEncryptedResponseAlg returns a boolean if a field has been set.

### GetIdpInitiatedLogin

`func (o *OpenIdConnectApplicationSettingsClient) GetIdpInitiatedLogin() OpenIdConnectApplicationIdpInitiatedLogin`

GetIdpInitiatedLogin returns the IdpInitiatedLogin field if non-nil, zero value otherwise.

### GetIdpInitiatedLoginOk

`func (o *OpenIdConnectApplicationSettingsClient) GetIdpInitiatedLoginOk() (*OpenIdConnectApplicationIdpInitiatedLogin, bool)`

GetIdpInitiatedLoginOk returns a tuple with the IdpInitiatedLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpInitiatedLogin

`func (o *OpenIdConnectApplicationSettingsClient) SetIdpInitiatedLogin(v OpenIdConnectApplicationIdpInitiatedLogin)`

SetIdpInitiatedLogin sets IdpInitiatedLogin field to given value.

### HasIdpInitiatedLogin

`func (o *OpenIdConnectApplicationSettingsClient) HasIdpInitiatedLogin() bool`

HasIdpInitiatedLogin returns a boolean if a field has been set.

### GetInitiateLoginUri

`func (o *OpenIdConnectApplicationSettingsClient) GetInitiateLoginUri() string`

GetInitiateLoginUri returns the InitiateLoginUri field if non-nil, zero value otherwise.

### GetInitiateLoginUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetInitiateLoginUriOk() (*string, bool)`

GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateLoginUri

`func (o *OpenIdConnectApplicationSettingsClient) SetInitiateLoginUri(v string)`

SetInitiateLoginUri sets InitiateLoginUri field to given value.

### HasInitiateLoginUri

`func (o *OpenIdConnectApplicationSettingsClient) HasInitiateLoginUri() bool`

HasInitiateLoginUri returns a boolean if a field has been set.

### GetIssuerMode

`func (o *OpenIdConnectApplicationSettingsClient) GetIssuerMode() string`

GetIssuerMode returns the IssuerMode field if non-nil, zero value otherwise.

### GetIssuerModeOk

`func (o *OpenIdConnectApplicationSettingsClient) GetIssuerModeOk() (*string, bool)`

GetIssuerModeOk returns a tuple with the IssuerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMode

`func (o *OpenIdConnectApplicationSettingsClient) SetIssuerMode(v string)`

SetIssuerMode sets IssuerMode field to given value.

### HasIssuerMode

`func (o *OpenIdConnectApplicationSettingsClient) HasIssuerMode() bool`

HasIssuerMode returns a boolean if a field has been set.

### GetJwks

`func (o *OpenIdConnectApplicationSettingsClient) GetJwks() OpenIdConnectApplicationSettingsClientKeys`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *OpenIdConnectApplicationSettingsClient) GetJwksOk() (*OpenIdConnectApplicationSettingsClientKeys, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *OpenIdConnectApplicationSettingsClient) SetJwks(v OpenIdConnectApplicationSettingsClientKeys)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *OpenIdConnectApplicationSettingsClient) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetJwksUri

`func (o *OpenIdConnectApplicationSettingsClient) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *OpenIdConnectApplicationSettingsClient) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *OpenIdConnectApplicationSettingsClient) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetLogoUri

`func (o *OpenIdConnectApplicationSettingsClient) GetLogoUri() string`

GetLogoUri returns the LogoUri field if non-nil, zero value otherwise.

### GetLogoUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetLogoUriOk() (*string, bool)`

GetLogoUriOk returns a tuple with the LogoUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUri

`func (o *OpenIdConnectApplicationSettingsClient) SetLogoUri(v string)`

SetLogoUri sets LogoUri field to given value.

### HasLogoUri

`func (o *OpenIdConnectApplicationSettingsClient) HasLogoUri() bool`

HasLogoUri returns a boolean if a field has been set.

### GetNetwork

`func (o *OpenIdConnectApplicationSettingsClient) GetNetwork() OpenIdConnectApplicationNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *OpenIdConnectApplicationSettingsClient) GetNetworkOk() (*OpenIdConnectApplicationNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *OpenIdConnectApplicationSettingsClient) SetNetwork(v OpenIdConnectApplicationNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *OpenIdConnectApplicationSettingsClient) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetParticipateSlo

`func (o *OpenIdConnectApplicationSettingsClient) GetParticipateSlo() bool`

GetParticipateSlo returns the ParticipateSlo field if non-nil, zero value otherwise.

### GetParticipateSloOk

`func (o *OpenIdConnectApplicationSettingsClient) GetParticipateSloOk() (*bool, bool)`

GetParticipateSloOk returns a tuple with the ParticipateSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipateSlo

`func (o *OpenIdConnectApplicationSettingsClient) SetParticipateSlo(v bool)`

SetParticipateSlo sets ParticipateSlo field to given value.

### HasParticipateSlo

`func (o *OpenIdConnectApplicationSettingsClient) HasParticipateSlo() bool`

HasParticipateSlo returns a boolean if a field has been set.

### GetPolicyUri

`func (o *OpenIdConnectApplicationSettingsClient) GetPolicyUri() string`

GetPolicyUri returns the PolicyUri field if non-nil, zero value otherwise.

### GetPolicyUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetPolicyUriOk() (*string, bool)`

GetPolicyUriOk returns a tuple with the PolicyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUri

`func (o *OpenIdConnectApplicationSettingsClient) SetPolicyUri(v string)`

SetPolicyUri sets PolicyUri field to given value.

### HasPolicyUri

`func (o *OpenIdConnectApplicationSettingsClient) HasPolicyUri() bool`

HasPolicyUri returns a boolean if a field has been set.

### GetPostLogoutRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *OpenIdConnectApplicationSettingsClient) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OpenIdConnectApplicationSettingsClient) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *OpenIdConnectApplicationSettingsClient) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshToken

`func (o *OpenIdConnectApplicationSettingsClient) GetRefreshToken() OpenIdConnectApplicationSettingsRefreshToken`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OpenIdConnectApplicationSettingsClient) GetRefreshTokenOk() (*OpenIdConnectApplicationSettingsRefreshToken, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OpenIdConnectApplicationSettingsClient) SetRefreshToken(v OpenIdConnectApplicationSettingsRefreshToken)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OpenIdConnectApplicationSettingsClient) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetRequestObjectSigningAlg

`func (o *OpenIdConnectApplicationSettingsClient) GetRequestObjectSigningAlg() string`

GetRequestObjectSigningAlg returns the RequestObjectSigningAlg field if non-nil, zero value otherwise.

### GetRequestObjectSigningAlgOk

`func (o *OpenIdConnectApplicationSettingsClient) GetRequestObjectSigningAlgOk() (*string, bool)`

GetRequestObjectSigningAlgOk returns a tuple with the RequestObjectSigningAlg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectSigningAlg

`func (o *OpenIdConnectApplicationSettingsClient) SetRequestObjectSigningAlg(v string)`

SetRequestObjectSigningAlg sets RequestObjectSigningAlg field to given value.

### HasRequestObjectSigningAlg

`func (o *OpenIdConnectApplicationSettingsClient) HasRequestObjectSigningAlg() bool`

HasRequestObjectSigningAlg returns a boolean if a field has been set.

### GetResponseTypes

`func (o *OpenIdConnectApplicationSettingsClient) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *OpenIdConnectApplicationSettingsClient) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *OpenIdConnectApplicationSettingsClient) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *OpenIdConnectApplicationSettingsClient) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetSectorIdentifierUri

`func (o *OpenIdConnectApplicationSettingsClient) GetSectorIdentifierUri() string`

GetSectorIdentifierUri returns the SectorIdentifierUri field if non-nil, zero value otherwise.

### GetSectorIdentifierUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetSectorIdentifierUriOk() (*string, bool)`

GetSectorIdentifierUriOk returns a tuple with the SectorIdentifierUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorIdentifierUri

`func (o *OpenIdConnectApplicationSettingsClient) SetSectorIdentifierUri(v string)`

SetSectorIdentifierUri sets SectorIdentifierUri field to given value.

### HasSectorIdentifierUri

`func (o *OpenIdConnectApplicationSettingsClient) HasSectorIdentifierUri() bool`

HasSectorIdentifierUri returns a boolean if a field has been set.

### GetSubjectType

`func (o *OpenIdConnectApplicationSettingsClient) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *OpenIdConnectApplicationSettingsClient) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *OpenIdConnectApplicationSettingsClient) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.

### HasSubjectType

`func (o *OpenIdConnectApplicationSettingsClient) HasSubjectType() bool`

HasSubjectType returns a boolean if a field has been set.

### GetTosUri

`func (o *OpenIdConnectApplicationSettingsClient) GetTosUri() string`

GetTosUri returns the TosUri field if non-nil, zero value otherwise.

### GetTosUriOk

`func (o *OpenIdConnectApplicationSettingsClient) GetTosUriOk() (*string, bool)`

GetTosUriOk returns a tuple with the TosUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosUri

`func (o *OpenIdConnectApplicationSettingsClient) SetTosUri(v string)`

SetTosUri sets TosUri field to given value.

### HasTosUri

`func (o *OpenIdConnectApplicationSettingsClient) HasTosUri() bool`

HasTosUri returns a boolean if a field has been set.

### GetWildcardRedirect

`func (o *OpenIdConnectApplicationSettingsClient) GetWildcardRedirect() string`

GetWildcardRedirect returns the WildcardRedirect field if non-nil, zero value otherwise.

### GetWildcardRedirectOk

`func (o *OpenIdConnectApplicationSettingsClient) GetWildcardRedirectOk() (*string, bool)`

GetWildcardRedirectOk returns a tuple with the WildcardRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcardRedirect

`func (o *OpenIdConnectApplicationSettingsClient) SetWildcardRedirect(v string)`

SetWildcardRedirect sets WildcardRedirect field to given value.

### HasWildcardRedirect

`func (o *OpenIdConnectApplicationSettingsClient) HasWildcardRedirect() bool`

HasWildcardRedirect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


