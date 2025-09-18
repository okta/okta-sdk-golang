# BaseToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to **map[string]interface{}** | Claims included in the token. Consists of name-value pairs for each included claim. For descriptions of the claims that you can include, see the Okta [OpenID Connect and OAuth 2.0 API reference](/openapi/okta-oauth/guides/overview/#claims). | [optional] 
**Token** | Pointer to [**BaseTokenToken**](BaseTokenToken.md) |  | [optional] 

## Methods

### NewBaseToken

`func NewBaseToken() *BaseToken`

NewBaseToken instantiates a new BaseToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTokenWithDefaults

`func NewBaseTokenWithDefaults() *BaseToken`

NewBaseTokenWithDefaults instantiates a new BaseToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *BaseToken) GetClaims() map[string]interface{}`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *BaseToken) GetClaimsOk() (*map[string]interface{}, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *BaseToken) SetClaims(v map[string]interface{})`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *BaseToken) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetToken

`func (o *BaseToken) GetToken() BaseTokenToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *BaseToken) GetTokenOk() (*BaseTokenToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *BaseToken) SetToken(v BaseTokenToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *BaseToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


