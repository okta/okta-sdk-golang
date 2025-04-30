# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | An access token. | [optional] 
**DeviceSecret** | Pointer to **string** | An opaque device secret. This is returned if the &#x60;device_sso&#x60; scope is granted. | [optional] 
**ExpiresIn** | Pointer to **int32** | The expiration time of the access token in seconds. | [optional] 
**IdToken** | Pointer to **string** | An ID token. This is returned if the &#x60;openid&#x60; scope is granted. | [optional] 
**IssuedTokenType** | Pointer to **string** | The type of token for token exchange. | [optional] 
**RefreshToken** | Pointer to **string** | An opaque refresh token. This is returned if the &#x60;offline_access&#x60; scope is granted. | [optional] 
**Scope** | Pointer to **string** | The scopes contained in the access token. | [optional] 
**TokenType** | Pointer to **string** | The token type in a &#x60;/token&#x60; response. The value is generally &#x60;Bearer&#x60; except for a few instances of token exchange. | [optional] 

## Methods

### NewTokenResponse

`func NewTokenResponse() *TokenResponse`

NewTokenResponse instantiates a new TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseWithDefaults

`func NewTokenResponseWithDefaults() *TokenResponse`

NewTokenResponseWithDefaults instantiates a new TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *TokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TokenResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetDeviceSecret

`func (o *TokenResponse) GetDeviceSecret() string`

GetDeviceSecret returns the DeviceSecret field if non-nil, zero value otherwise.

### GetDeviceSecretOk

`func (o *TokenResponse) GetDeviceSecretOk() (*string, bool)`

GetDeviceSecretOk returns a tuple with the DeviceSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSecret

`func (o *TokenResponse) SetDeviceSecret(v string)`

SetDeviceSecret sets DeviceSecret field to given value.

### HasDeviceSecret

`func (o *TokenResponse) HasDeviceSecret() bool`

HasDeviceSecret returns a boolean if a field has been set.

### GetExpiresIn

`func (o *TokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *TokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *TokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *TokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetIdToken

`func (o *TokenResponse) GetIdToken() string`

GetIdToken returns the IdToken field if non-nil, zero value otherwise.

### GetIdTokenOk

`func (o *TokenResponse) GetIdTokenOk() (*string, bool)`

GetIdTokenOk returns a tuple with the IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdToken

`func (o *TokenResponse) SetIdToken(v string)`

SetIdToken sets IdToken field to given value.

### HasIdToken

`func (o *TokenResponse) HasIdToken() bool`

HasIdToken returns a boolean if a field has been set.

### GetIssuedTokenType

`func (o *TokenResponse) GetIssuedTokenType() string`

GetIssuedTokenType returns the IssuedTokenType field if non-nil, zero value otherwise.

### GetIssuedTokenTypeOk

`func (o *TokenResponse) GetIssuedTokenTypeOk() (*string, bool)`

GetIssuedTokenTypeOk returns a tuple with the IssuedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTokenType

`func (o *TokenResponse) SetIssuedTokenType(v string)`

SetIssuedTokenType sets IssuedTokenType field to given value.

### HasIssuedTokenType

`func (o *TokenResponse) HasIssuedTokenType() bool`

HasIssuedTokenType returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetScope

`func (o *TokenResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TokenResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TokenResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TokenResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TokenResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


