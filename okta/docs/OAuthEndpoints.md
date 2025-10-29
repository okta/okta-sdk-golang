# OAuthEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to [**OAuthAuthorizationEndpoint**](OAuthAuthorizationEndpoint.md) |  | [optional] 
**Jwks** | Pointer to [**OidcJwksEndpoint**](OidcJwksEndpoint.md) |  | [optional] 
**Slo** | Pointer to [**OidcSloEndpoint**](OidcSloEndpoint.md) |  | [optional] 
**Token** | Pointer to [**OAuthTokenEndpoint**](OAuthTokenEndpoint.md) |  | [optional] 
**UserInfo** | Pointer to [**OidcUserInfoEndpoint**](OidcUserInfoEndpoint.md) |  | [optional] 

## Methods

### NewOAuthEndpoints

`func NewOAuthEndpoints() *OAuthEndpoints`

NewOAuthEndpoints instantiates a new OAuthEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthEndpointsWithDefaults

`func NewOAuthEndpointsWithDefaults() *OAuthEndpoints`

NewOAuthEndpointsWithDefaults instantiates a new OAuthEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *OAuthEndpoints) GetAuthorization() OAuthAuthorizationEndpoint`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *OAuthEndpoints) GetAuthorizationOk() (*OAuthAuthorizationEndpoint, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *OAuthEndpoints) SetAuthorization(v OAuthAuthorizationEndpoint)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *OAuthEndpoints) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetJwks

`func (o *OAuthEndpoints) GetJwks() OidcJwksEndpoint`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *OAuthEndpoints) GetJwksOk() (*OidcJwksEndpoint, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *OAuthEndpoints) SetJwks(v OidcJwksEndpoint)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *OAuthEndpoints) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetSlo

`func (o *OAuthEndpoints) GetSlo() OidcSloEndpoint`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *OAuthEndpoints) GetSloOk() (*OidcSloEndpoint, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *OAuthEndpoints) SetSlo(v OidcSloEndpoint)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *OAuthEndpoints) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetToken

`func (o *OAuthEndpoints) GetToken() OAuthTokenEndpoint`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OAuthEndpoints) GetTokenOk() (*OAuthTokenEndpoint, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OAuthEndpoints) SetToken(v OAuthTokenEndpoint)`

SetToken sets Token field to given value.

### HasToken

`func (o *OAuthEndpoints) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserInfo

`func (o *OAuthEndpoints) GetUserInfo() OidcUserInfoEndpoint`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *OAuthEndpoints) GetUserInfoOk() (*OidcUserInfoEndpoint, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *OAuthEndpoints) SetUserInfo(v OidcUserInfoEndpoint)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *OAuthEndpoints) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


