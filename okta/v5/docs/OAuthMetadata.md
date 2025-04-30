# OAuthMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationEndpoint** | Pointer to **string** | URL of the authorization server&#39;s authorization endpoint. | [optional] 
**BackchannelAuthenticationRequestSigningAlgValuesSupported** | Pointer to **[]string** | &lt;div class&#x3D;\&quot;x-lifecycle-container\&quot;&gt;&lt;x-lifecycle class&#x3D;\&quot;lea\&quot;&gt;&lt;/x-lifecycle&gt; &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/div&gt;A list of signing algorithms that this authorization server supports for signed requests. | [optional] 
**BackchannelTokenDeliveryModesSupported** | Pointer to **[]string** | &lt;div class&#x3D;\&quot;x-lifecycle-container\&quot;&gt;&lt;x-lifecycle class&#x3D;\&quot;lea\&quot;&gt;&lt;/x-lifecycle&gt; &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/div&gt;The delivery modes that this authorization server supports for Client-Initiated Backchannel Authentication. | [optional] 
**ClaimsSupported** | Pointer to **[]string** | A list of the claims supported by this authorization server. | [optional] 
**CodeChallengeMethodsSupported** | Pointer to **[]string** | A list of PKCE code challenge methods supported by this authorization server. | [optional] 
**DeviceAuthorizationEndpoint** | Pointer to **string** |  | [optional] 
**DpopSigningAlgValuesSupported** | Pointer to **[]string** | A list of signing algorithms supported by this authorization server for Demonstrating Proof-of-Possession (DPoP) JWTs. | [optional] 
**EndSessionEndpoint** | Pointer to **string** | URL of the authorization server&#39;s logout endpoint. | [optional] 
**GrantTypesSupported** | Pointer to **[]string** | A list of the grant type values that this authorization server supports. | [optional] 
**IntrospectionEndpoint** | Pointer to **string** | URL of the authorization server&#39;s introspection endpoint. | [optional] 
**IntrospectionEndpointAuthMethodsSupported** | Pointer to **[]string** | A list of client authentication methods supported by this introspection endpoint. | [optional] 
**Issuer** | Pointer to **string** | The authorization server&#39;s issuer identifier. In the context of this document, this is your authorization server&#39;s base URL. This becomes the &#x60;iss&#x60; claim in an access token. | [optional] 
**JwksUri** | Pointer to **string** | URL of the authorization server&#39;s JSON Web Key Set document. | [optional] 
**PushedAuthorizationRequestEndpoint** | Pointer to **string** |  | [optional] 
**RegistrationEndpoint** | Pointer to **string** | URL of the authorization server&#39;s JSON Web Key Set document. | [optional] 
**RequestObjectSigningAlgValuesSupported** | Pointer to **[]string** | A list of signing algorithms that this authorization server supports for signed requests. | [optional] 
**RequestParameterSupported** | Pointer to **bool** | Indicates if Request Parameters are supported by this authorization server. | [optional] 
**ResponseModesSupported** | Pointer to **[]string** | A list of the &#x60;response_mode&#x60; values that this authorization server supports. More information here. | [optional] 
**ResponseTypesSupported** | Pointer to **[]string** | A list of the &#x60;response_type&#x60; values that this authorization server supports. Can be a combination of &#x60;code&#x60;, &#x60;token&#x60;, and &#x60;id_token&#x60;. | [optional] 
**RevocationEndpoint** | Pointer to **string** | URL of the authorization server&#39;s revocation endpoint. | [optional] 
**RevocationEndpointAuthMethodsSupported** | Pointer to **[]string** | A list of client authentication methods supported by this revocation endpoint. | [optional] 
**ScopesSupported** | Pointer to **[]string** | A list of the scope values that this authorization server supports. | [optional] 
**SubjectTypesSupported** | Pointer to **[]string** | A list of the Subject Identifier types that this authorization server supports. Valid types include &#x60;pairwise&#x60; and &#x60;public&#x60;, but only &#x60;public&#x60; is currently supported. See the [Subject Identifier Types](https://openid.net/specs/openid-connect-core-1_0.html#SubjectIDTypes) section in the OpenID Connect specification. | [optional] 
**TokenEndpoint** | Pointer to **string** | URL of the authorization server&#39;s token endpoint. | [optional] 
**TokenEndpointAuthMethodsSupported** | Pointer to **[]string** | A list of client authentication methods supported by this token endpoint. | [optional] 

## Methods

### NewOAuthMetadata

`func NewOAuthMetadata() *OAuthMetadata`

NewOAuthMetadata instantiates a new OAuthMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthMetadataWithDefaults

`func NewOAuthMetadataWithDefaults() *OAuthMetadata`

NewOAuthMetadataWithDefaults instantiates a new OAuthMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationEndpoint

`func (o *OAuthMetadata) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *OAuthMetadata) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *OAuthMetadata) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *OAuthMetadata) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetBackchannelAuthenticationRequestSigningAlgValuesSupported

`func (o *OAuthMetadata) GetBackchannelAuthenticationRequestSigningAlgValuesSupported() []string`

GetBackchannelAuthenticationRequestSigningAlgValuesSupported returns the BackchannelAuthenticationRequestSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetBackchannelAuthenticationRequestSigningAlgValuesSupportedOk

`func (o *OAuthMetadata) GetBackchannelAuthenticationRequestSigningAlgValuesSupportedOk() (*[]string, bool)`

GetBackchannelAuthenticationRequestSigningAlgValuesSupportedOk returns a tuple with the BackchannelAuthenticationRequestSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelAuthenticationRequestSigningAlgValuesSupported

`func (o *OAuthMetadata) SetBackchannelAuthenticationRequestSigningAlgValuesSupported(v []string)`

SetBackchannelAuthenticationRequestSigningAlgValuesSupported sets BackchannelAuthenticationRequestSigningAlgValuesSupported field to given value.

### HasBackchannelAuthenticationRequestSigningAlgValuesSupported

`func (o *OAuthMetadata) HasBackchannelAuthenticationRequestSigningAlgValuesSupported() bool`

HasBackchannelAuthenticationRequestSigningAlgValuesSupported returns a boolean if a field has been set.

### GetBackchannelTokenDeliveryModesSupported

`func (o *OAuthMetadata) GetBackchannelTokenDeliveryModesSupported() []string`

GetBackchannelTokenDeliveryModesSupported returns the BackchannelTokenDeliveryModesSupported field if non-nil, zero value otherwise.

### GetBackchannelTokenDeliveryModesSupportedOk

`func (o *OAuthMetadata) GetBackchannelTokenDeliveryModesSupportedOk() (*[]string, bool)`

GetBackchannelTokenDeliveryModesSupportedOk returns a tuple with the BackchannelTokenDeliveryModesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackchannelTokenDeliveryModesSupported

`func (o *OAuthMetadata) SetBackchannelTokenDeliveryModesSupported(v []string)`

SetBackchannelTokenDeliveryModesSupported sets BackchannelTokenDeliveryModesSupported field to given value.

### HasBackchannelTokenDeliveryModesSupported

`func (o *OAuthMetadata) HasBackchannelTokenDeliveryModesSupported() bool`

HasBackchannelTokenDeliveryModesSupported returns a boolean if a field has been set.

### GetClaimsSupported

`func (o *OAuthMetadata) GetClaimsSupported() []string`

GetClaimsSupported returns the ClaimsSupported field if non-nil, zero value otherwise.

### GetClaimsSupportedOk

`func (o *OAuthMetadata) GetClaimsSupportedOk() (*[]string, bool)`

GetClaimsSupportedOk returns a tuple with the ClaimsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsSupported

`func (o *OAuthMetadata) SetClaimsSupported(v []string)`

SetClaimsSupported sets ClaimsSupported field to given value.

### HasClaimsSupported

`func (o *OAuthMetadata) HasClaimsSupported() bool`

HasClaimsSupported returns a boolean if a field has been set.

### GetCodeChallengeMethodsSupported

`func (o *OAuthMetadata) GetCodeChallengeMethodsSupported() []string`

GetCodeChallengeMethodsSupported returns the CodeChallengeMethodsSupported field if non-nil, zero value otherwise.

### GetCodeChallengeMethodsSupportedOk

`func (o *OAuthMetadata) GetCodeChallengeMethodsSupportedOk() (*[]string, bool)`

GetCodeChallengeMethodsSupportedOk returns a tuple with the CodeChallengeMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallengeMethodsSupported

`func (o *OAuthMetadata) SetCodeChallengeMethodsSupported(v []string)`

SetCodeChallengeMethodsSupported sets CodeChallengeMethodsSupported field to given value.

### HasCodeChallengeMethodsSupported

`func (o *OAuthMetadata) HasCodeChallengeMethodsSupported() bool`

HasCodeChallengeMethodsSupported returns a boolean if a field has been set.

### GetDeviceAuthorizationEndpoint

`func (o *OAuthMetadata) GetDeviceAuthorizationEndpoint() string`

GetDeviceAuthorizationEndpoint returns the DeviceAuthorizationEndpoint field if non-nil, zero value otherwise.

### GetDeviceAuthorizationEndpointOk

`func (o *OAuthMetadata) GetDeviceAuthorizationEndpointOk() (*string, bool)`

GetDeviceAuthorizationEndpointOk returns a tuple with the DeviceAuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthorizationEndpoint

`func (o *OAuthMetadata) SetDeviceAuthorizationEndpoint(v string)`

SetDeviceAuthorizationEndpoint sets DeviceAuthorizationEndpoint field to given value.

### HasDeviceAuthorizationEndpoint

`func (o *OAuthMetadata) HasDeviceAuthorizationEndpoint() bool`

HasDeviceAuthorizationEndpoint returns a boolean if a field has been set.

### GetDpopSigningAlgValuesSupported

`func (o *OAuthMetadata) GetDpopSigningAlgValuesSupported() []string`

GetDpopSigningAlgValuesSupported returns the DpopSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetDpopSigningAlgValuesSupportedOk

`func (o *OAuthMetadata) GetDpopSigningAlgValuesSupportedOk() (*[]string, bool)`

GetDpopSigningAlgValuesSupportedOk returns a tuple with the DpopSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpopSigningAlgValuesSupported

`func (o *OAuthMetadata) SetDpopSigningAlgValuesSupported(v []string)`

SetDpopSigningAlgValuesSupported sets DpopSigningAlgValuesSupported field to given value.

### HasDpopSigningAlgValuesSupported

`func (o *OAuthMetadata) HasDpopSigningAlgValuesSupported() bool`

HasDpopSigningAlgValuesSupported returns a boolean if a field has been set.

### GetEndSessionEndpoint

`func (o *OAuthMetadata) GetEndSessionEndpoint() string`

GetEndSessionEndpoint returns the EndSessionEndpoint field if non-nil, zero value otherwise.

### GetEndSessionEndpointOk

`func (o *OAuthMetadata) GetEndSessionEndpointOk() (*string, bool)`

GetEndSessionEndpointOk returns a tuple with the EndSessionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSessionEndpoint

`func (o *OAuthMetadata) SetEndSessionEndpoint(v string)`

SetEndSessionEndpoint sets EndSessionEndpoint field to given value.

### HasEndSessionEndpoint

`func (o *OAuthMetadata) HasEndSessionEndpoint() bool`

HasEndSessionEndpoint returns a boolean if a field has been set.

### GetGrantTypesSupported

`func (o *OAuthMetadata) GetGrantTypesSupported() []string`

GetGrantTypesSupported returns the GrantTypesSupported field if non-nil, zero value otherwise.

### GetGrantTypesSupportedOk

`func (o *OAuthMetadata) GetGrantTypesSupportedOk() (*[]string, bool)`

GetGrantTypesSupportedOk returns a tuple with the GrantTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypesSupported

`func (o *OAuthMetadata) SetGrantTypesSupported(v []string)`

SetGrantTypesSupported sets GrantTypesSupported field to given value.

### HasGrantTypesSupported

`func (o *OAuthMetadata) HasGrantTypesSupported() bool`

HasGrantTypesSupported returns a boolean if a field has been set.

### GetIntrospectionEndpoint

`func (o *OAuthMetadata) GetIntrospectionEndpoint() string`

GetIntrospectionEndpoint returns the IntrospectionEndpoint field if non-nil, zero value otherwise.

### GetIntrospectionEndpointOk

`func (o *OAuthMetadata) GetIntrospectionEndpointOk() (*string, bool)`

GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpoint

`func (o *OAuthMetadata) SetIntrospectionEndpoint(v string)`

SetIntrospectionEndpoint sets IntrospectionEndpoint field to given value.

### HasIntrospectionEndpoint

`func (o *OAuthMetadata) HasIntrospectionEndpoint() bool`

HasIntrospectionEndpoint returns a boolean if a field has been set.

### GetIntrospectionEndpointAuthMethodsSupported

`func (o *OAuthMetadata) GetIntrospectionEndpointAuthMethodsSupported() []string`

GetIntrospectionEndpointAuthMethodsSupported returns the IntrospectionEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetIntrospectionEndpointAuthMethodsSupportedOk

`func (o *OAuthMetadata) GetIntrospectionEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetIntrospectionEndpointAuthMethodsSupportedOk returns a tuple with the IntrospectionEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpointAuthMethodsSupported

`func (o *OAuthMetadata) SetIntrospectionEndpointAuthMethodsSupported(v []string)`

SetIntrospectionEndpointAuthMethodsSupported sets IntrospectionEndpointAuthMethodsSupported field to given value.

### HasIntrospectionEndpointAuthMethodsSupported

`func (o *OAuthMetadata) HasIntrospectionEndpointAuthMethodsSupported() bool`

HasIntrospectionEndpointAuthMethodsSupported returns a boolean if a field has been set.

### GetIssuer

`func (o *OAuthMetadata) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OAuthMetadata) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OAuthMetadata) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OAuthMetadata) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetJwksUri

`func (o *OAuthMetadata) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *OAuthMetadata) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *OAuthMetadata) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *OAuthMetadata) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetPushedAuthorizationRequestEndpoint

`func (o *OAuthMetadata) GetPushedAuthorizationRequestEndpoint() string`

GetPushedAuthorizationRequestEndpoint returns the PushedAuthorizationRequestEndpoint field if non-nil, zero value otherwise.

### GetPushedAuthorizationRequestEndpointOk

`func (o *OAuthMetadata) GetPushedAuthorizationRequestEndpointOk() (*string, bool)`

GetPushedAuthorizationRequestEndpointOk returns a tuple with the PushedAuthorizationRequestEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushedAuthorizationRequestEndpoint

`func (o *OAuthMetadata) SetPushedAuthorizationRequestEndpoint(v string)`

SetPushedAuthorizationRequestEndpoint sets PushedAuthorizationRequestEndpoint field to given value.

### HasPushedAuthorizationRequestEndpoint

`func (o *OAuthMetadata) HasPushedAuthorizationRequestEndpoint() bool`

HasPushedAuthorizationRequestEndpoint returns a boolean if a field has been set.

### GetRegistrationEndpoint

`func (o *OAuthMetadata) GetRegistrationEndpoint() string`

GetRegistrationEndpoint returns the RegistrationEndpoint field if non-nil, zero value otherwise.

### GetRegistrationEndpointOk

`func (o *OAuthMetadata) GetRegistrationEndpointOk() (*string, bool)`

GetRegistrationEndpointOk returns a tuple with the RegistrationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEndpoint

`func (o *OAuthMetadata) SetRegistrationEndpoint(v string)`

SetRegistrationEndpoint sets RegistrationEndpoint field to given value.

### HasRegistrationEndpoint

`func (o *OAuthMetadata) HasRegistrationEndpoint() bool`

HasRegistrationEndpoint returns a boolean if a field has been set.

### GetRequestObjectSigningAlgValuesSupported

`func (o *OAuthMetadata) GetRequestObjectSigningAlgValuesSupported() []string`

GetRequestObjectSigningAlgValuesSupported returns the RequestObjectSigningAlgValuesSupported field if non-nil, zero value otherwise.

### GetRequestObjectSigningAlgValuesSupportedOk

`func (o *OAuthMetadata) GetRequestObjectSigningAlgValuesSupportedOk() (*[]string, bool)`

GetRequestObjectSigningAlgValuesSupportedOk returns a tuple with the RequestObjectSigningAlgValuesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestObjectSigningAlgValuesSupported

`func (o *OAuthMetadata) SetRequestObjectSigningAlgValuesSupported(v []string)`

SetRequestObjectSigningAlgValuesSupported sets RequestObjectSigningAlgValuesSupported field to given value.

### HasRequestObjectSigningAlgValuesSupported

`func (o *OAuthMetadata) HasRequestObjectSigningAlgValuesSupported() bool`

HasRequestObjectSigningAlgValuesSupported returns a boolean if a field has been set.

### GetRequestParameterSupported

`func (o *OAuthMetadata) GetRequestParameterSupported() bool`

GetRequestParameterSupported returns the RequestParameterSupported field if non-nil, zero value otherwise.

### GetRequestParameterSupportedOk

`func (o *OAuthMetadata) GetRequestParameterSupportedOk() (*bool, bool)`

GetRequestParameterSupportedOk returns a tuple with the RequestParameterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParameterSupported

`func (o *OAuthMetadata) SetRequestParameterSupported(v bool)`

SetRequestParameterSupported sets RequestParameterSupported field to given value.

### HasRequestParameterSupported

`func (o *OAuthMetadata) HasRequestParameterSupported() bool`

HasRequestParameterSupported returns a boolean if a field has been set.

### GetResponseModesSupported

`func (o *OAuthMetadata) GetResponseModesSupported() []string`

GetResponseModesSupported returns the ResponseModesSupported field if non-nil, zero value otherwise.

### GetResponseModesSupportedOk

`func (o *OAuthMetadata) GetResponseModesSupportedOk() (*[]string, bool)`

GetResponseModesSupportedOk returns a tuple with the ResponseModesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseModesSupported

`func (o *OAuthMetadata) SetResponseModesSupported(v []string)`

SetResponseModesSupported sets ResponseModesSupported field to given value.

### HasResponseModesSupported

`func (o *OAuthMetadata) HasResponseModesSupported() bool`

HasResponseModesSupported returns a boolean if a field has been set.

### GetResponseTypesSupported

`func (o *OAuthMetadata) GetResponseTypesSupported() []string`

GetResponseTypesSupported returns the ResponseTypesSupported field if non-nil, zero value otherwise.

### GetResponseTypesSupportedOk

`func (o *OAuthMetadata) GetResponseTypesSupportedOk() (*[]string, bool)`

GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypesSupported

`func (o *OAuthMetadata) SetResponseTypesSupported(v []string)`

SetResponseTypesSupported sets ResponseTypesSupported field to given value.

### HasResponseTypesSupported

`func (o *OAuthMetadata) HasResponseTypesSupported() bool`

HasResponseTypesSupported returns a boolean if a field has been set.

### GetRevocationEndpoint

`func (o *OAuthMetadata) GetRevocationEndpoint() string`

GetRevocationEndpoint returns the RevocationEndpoint field if non-nil, zero value otherwise.

### GetRevocationEndpointOk

`func (o *OAuthMetadata) GetRevocationEndpointOk() (*string, bool)`

GetRevocationEndpointOk returns a tuple with the RevocationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpoint

`func (o *OAuthMetadata) SetRevocationEndpoint(v string)`

SetRevocationEndpoint sets RevocationEndpoint field to given value.

### HasRevocationEndpoint

`func (o *OAuthMetadata) HasRevocationEndpoint() bool`

HasRevocationEndpoint returns a boolean if a field has been set.

### GetRevocationEndpointAuthMethodsSupported

`func (o *OAuthMetadata) GetRevocationEndpointAuthMethodsSupported() []string`

GetRevocationEndpointAuthMethodsSupported returns the RevocationEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetRevocationEndpointAuthMethodsSupportedOk

`func (o *OAuthMetadata) GetRevocationEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetRevocationEndpointAuthMethodsSupportedOk returns a tuple with the RevocationEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpointAuthMethodsSupported

`func (o *OAuthMetadata) SetRevocationEndpointAuthMethodsSupported(v []string)`

SetRevocationEndpointAuthMethodsSupported sets RevocationEndpointAuthMethodsSupported field to given value.

### HasRevocationEndpointAuthMethodsSupported

`func (o *OAuthMetadata) HasRevocationEndpointAuthMethodsSupported() bool`

HasRevocationEndpointAuthMethodsSupported returns a boolean if a field has been set.

### GetScopesSupported

`func (o *OAuthMetadata) GetScopesSupported() []string`

GetScopesSupported returns the ScopesSupported field if non-nil, zero value otherwise.

### GetScopesSupportedOk

`func (o *OAuthMetadata) GetScopesSupportedOk() (*[]string, bool)`

GetScopesSupportedOk returns a tuple with the ScopesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesSupported

`func (o *OAuthMetadata) SetScopesSupported(v []string)`

SetScopesSupported sets ScopesSupported field to given value.

### HasScopesSupported

`func (o *OAuthMetadata) HasScopesSupported() bool`

HasScopesSupported returns a boolean if a field has been set.

### GetSubjectTypesSupported

`func (o *OAuthMetadata) GetSubjectTypesSupported() []string`

GetSubjectTypesSupported returns the SubjectTypesSupported field if non-nil, zero value otherwise.

### GetSubjectTypesSupportedOk

`func (o *OAuthMetadata) GetSubjectTypesSupportedOk() (*[]string, bool)`

GetSubjectTypesSupportedOk returns a tuple with the SubjectTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTypesSupported

`func (o *OAuthMetadata) SetSubjectTypesSupported(v []string)`

SetSubjectTypesSupported sets SubjectTypesSupported field to given value.

### HasSubjectTypesSupported

`func (o *OAuthMetadata) HasSubjectTypesSupported() bool`

HasSubjectTypesSupported returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *OAuthMetadata) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OAuthMetadata) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OAuthMetadata) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *OAuthMetadata) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetTokenEndpointAuthMethodsSupported

`func (o *OAuthMetadata) GetTokenEndpointAuthMethodsSupported() []string`

GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodsSupportedOk

`func (o *OAuthMetadata) GetTokenEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethodsSupported

`func (o *OAuthMetadata) SetTokenEndpointAuthMethodsSupported(v []string)`

SetTokenEndpointAuthMethodsSupported sets TokenEndpointAuthMethodsSupported field to given value.

### HasTokenEndpointAuthMethodsSupported

`func (o *OAuthMetadata) HasTokenEndpointAuthMethodsSupported() bool`

HasTokenEndpointAuthMethodsSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


