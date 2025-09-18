# TokenPayLoadDataAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to **map[string]interface{}** | Claims included in the token. Consists of name-value pairs for each included claim. For descriptions of the claims that you can include, see the Okta [OpenID Connect and OAuth 2.0 API reference](/openapi/okta-oauth/guides/overview/#claims). | [optional] 
**Token** | Pointer to [**BaseTokenToken**](BaseTokenToken.md) |  | [optional] 
**Scopes** | Pointer to **map[string]interface{}** | The scopes contained in the token. For descriptions of the scopes that you can include, see the Okta [OpenID Connect and OAuth 2.0 API reference](/openapi/okta-oauth/guides/overview/#scopes). | [optional] 

## Methods

### NewTokenPayLoadDataAccess

`func NewTokenPayLoadDataAccess() *TokenPayLoadDataAccess`

NewTokenPayLoadDataAccess instantiates a new TokenPayLoadDataAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenPayLoadDataAccessWithDefaults

`func NewTokenPayLoadDataAccessWithDefaults() *TokenPayLoadDataAccess`

NewTokenPayLoadDataAccessWithDefaults instantiates a new TokenPayLoadDataAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *TokenPayLoadDataAccess) GetClaims() map[string]interface{}`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *TokenPayLoadDataAccess) GetClaimsOk() (*map[string]interface{}, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *TokenPayLoadDataAccess) SetClaims(v map[string]interface{})`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *TokenPayLoadDataAccess) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetToken

`func (o *TokenPayLoadDataAccess) GetToken() BaseTokenToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenPayLoadDataAccess) GetTokenOk() (*BaseTokenToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenPayLoadDataAccess) SetToken(v BaseTokenToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenPayLoadDataAccess) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetScopes

`func (o *TokenPayLoadDataAccess) GetScopes() map[string]interface{}`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenPayLoadDataAccess) GetScopesOk() (*map[string]interface{}, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenPayLoadDataAccess) SetScopes(v map[string]interface{})`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenPayLoadDataAccess) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


