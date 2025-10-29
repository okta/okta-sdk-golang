# ProtocolOAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**OAuthCredentials**](OAuthCredentials.md) |  | [optional] 
**Endpoints** | Pointer to [**OAuthEndpoints**](OAuthEndpoints.md) |  | [optional] 
**Scopes** | Pointer to **[]string** | IdP-defined permission bundles to request delegated access from the user. &gt; **Note:** The [identity provider type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path&#x3D;type&amp;t&#x3D;request) table lists the scopes that are supported for each IdP. | [optional] 
**Type** | Pointer to **string** | OAuth 2.0 Authorization Code flow | [optional] 

## Methods

### NewProtocolOAuth

`func NewProtocolOAuth() *ProtocolOAuth`

NewProtocolOAuth instantiates a new ProtocolOAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolOAuthWithDefaults

`func NewProtocolOAuthWithDefaults() *ProtocolOAuth`

NewProtocolOAuthWithDefaults instantiates a new ProtocolOAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ProtocolOAuth) GetCredentials() OAuthCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ProtocolOAuth) GetCredentialsOk() (*OAuthCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ProtocolOAuth) SetCredentials(v OAuthCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ProtocolOAuth) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoints

`func (o *ProtocolOAuth) GetEndpoints() OAuthEndpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ProtocolOAuth) GetEndpointsOk() (*OAuthEndpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ProtocolOAuth) SetEndpoints(v OAuthEndpoints)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ProtocolOAuth) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetScopes

`func (o *ProtocolOAuth) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ProtocolOAuth) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ProtocolOAuth) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ProtocolOAuth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetType

`func (o *ProtocolOAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProtocolOAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProtocolOAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProtocolOAuth) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


