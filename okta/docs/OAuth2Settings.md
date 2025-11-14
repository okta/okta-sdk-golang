# OAuth2Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizeEndpoint** | **string** | The URL to the authorization server&#39;s authorization endpoint | 
**ClientId** | **string** | The OAuth 2.0 client identifier | 
**ClientSecret** | **string** | The OAuth 2.0 client secret | 
**Scopes** | Pointer to **[]string** | List of OAuth 2.0 scopes | [optional] 
**TokenEndpoint** | **string** | The URL to the authorization server&#39;s token endpoint | 

## Methods

### NewOAuth2Settings

`func NewOAuth2Settings(authorizeEndpoint string, clientId string, clientSecret string, tokenEndpoint string, ) *OAuth2Settings`

NewOAuth2Settings instantiates a new OAuth2Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2SettingsWithDefaults

`func NewOAuth2SettingsWithDefaults() *OAuth2Settings`

NewOAuth2SettingsWithDefaults instantiates a new OAuth2Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizeEndpoint

`func (o *OAuth2Settings) GetAuthorizeEndpoint() string`

GetAuthorizeEndpoint returns the AuthorizeEndpoint field if non-nil, zero value otherwise.

### GetAuthorizeEndpointOk

`func (o *OAuth2Settings) GetAuthorizeEndpointOk() (*string, bool)`

GetAuthorizeEndpointOk returns a tuple with the AuthorizeEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeEndpoint

`func (o *OAuth2Settings) SetAuthorizeEndpoint(v string)`

SetAuthorizeEndpoint sets AuthorizeEndpoint field to given value.


### GetClientId

`func (o *OAuth2Settings) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2Settings) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2Settings) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuth2Settings) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2Settings) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2Settings) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetScopes

`func (o *OAuth2Settings) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAuth2Settings) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAuth2Settings) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAuth2Settings) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *OAuth2Settings) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OAuth2Settings) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OAuth2Settings) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


