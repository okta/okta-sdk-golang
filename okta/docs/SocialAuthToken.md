# SocialAuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TokenAuthScheme** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 

## Methods

### NewSocialAuthToken

`func NewSocialAuthToken() *SocialAuthToken`

NewSocialAuthToken instantiates a new SocialAuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialAuthTokenWithDefaults

`func NewSocialAuthTokenWithDefaults() *SocialAuthToken`

NewSocialAuthTokenWithDefaults instantiates a new SocialAuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *SocialAuthToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SocialAuthToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SocialAuthToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SocialAuthToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *SocialAuthToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialAuthToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialAuthToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialAuthToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScopes

`func (o *SocialAuthToken) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *SocialAuthToken) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *SocialAuthToken) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *SocialAuthToken) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetToken

`func (o *SocialAuthToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SocialAuthToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SocialAuthToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SocialAuthToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenAuthScheme

`func (o *SocialAuthToken) GetTokenAuthScheme() string`

GetTokenAuthScheme returns the TokenAuthScheme field if non-nil, zero value otherwise.

### GetTokenAuthSchemeOk

`func (o *SocialAuthToken) GetTokenAuthSchemeOk() (*string, bool)`

GetTokenAuthSchemeOk returns a tuple with the TokenAuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAuthScheme

`func (o *SocialAuthToken) SetTokenAuthScheme(v string)`

SetTokenAuthScheme sets TokenAuthScheme field to given value.

### HasTokenAuthScheme

`func (o *SocialAuthToken) HasTokenAuthScheme() bool`

HasTokenAuthScheme returns a boolean if a field has been set.

### GetTokenType

`func (o *SocialAuthToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *SocialAuthToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *SocialAuthToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *SocialAuthToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


