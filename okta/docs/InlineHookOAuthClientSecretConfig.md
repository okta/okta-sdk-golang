# InlineHookOAuthClientSecretConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**TokenUrl** | Pointer to **string** |  | [optional] 
**AuthScheme** | Pointer to [**InlineHookChannelConfigAuthScheme**](InlineHookChannelConfigAuthScheme.md) |  | [optional] 
**Headers** | Pointer to [**[]InlineHookChannelConfigHeaders**](InlineHookChannelConfigHeaders.md) |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineHookOAuthClientSecretConfig

`func NewInlineHookOAuthClientSecretConfig() *InlineHookOAuthClientSecretConfig`

NewInlineHookOAuthClientSecretConfig instantiates a new InlineHookOAuthClientSecretConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookOAuthClientSecretConfigWithDefaults

`func NewInlineHookOAuthClientSecretConfigWithDefaults() *InlineHookOAuthClientSecretConfig`

NewInlineHookOAuthClientSecretConfigWithDefaults instantiates a new InlineHookOAuthClientSecretConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *InlineHookOAuthClientSecretConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *InlineHookOAuthClientSecretConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *InlineHookOAuthClientSecretConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *InlineHookOAuthClientSecretConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAuthType

`func (o *InlineHookOAuthClientSecretConfig) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *InlineHookOAuthClientSecretConfig) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *InlineHookOAuthClientSecretConfig) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *InlineHookOAuthClientSecretConfig) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *InlineHookOAuthClientSecretConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineHookOAuthClientSecretConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineHookOAuthClientSecretConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineHookOAuthClientSecretConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetScope

`func (o *InlineHookOAuthClientSecretConfig) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineHookOAuthClientSecretConfig) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineHookOAuthClientSecretConfig) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineHookOAuthClientSecretConfig) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTokenUrl

`func (o *InlineHookOAuthClientSecretConfig) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *InlineHookOAuthClientSecretConfig) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *InlineHookOAuthClientSecretConfig) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *InlineHookOAuthClientSecretConfig) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### GetAuthScheme

`func (o *InlineHookOAuthClientSecretConfig) GetAuthScheme() InlineHookChannelConfigAuthScheme`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *InlineHookOAuthClientSecretConfig) GetAuthSchemeOk() (*InlineHookChannelConfigAuthScheme, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *InlineHookOAuthClientSecretConfig) SetAuthScheme(v InlineHookChannelConfigAuthScheme)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *InlineHookOAuthClientSecretConfig) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineHookOAuthClientSecretConfig) GetHeaders() []InlineHookChannelConfigHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineHookOAuthClientSecretConfig) GetHeadersOk() (*[]InlineHookChannelConfigHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineHookOAuthClientSecretConfig) SetHeaders(v []InlineHookChannelConfigHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineHookOAuthClientSecretConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookOAuthClientSecretConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookOAuthClientSecretConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookOAuthClientSecretConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookOAuthClientSecretConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *InlineHookOAuthClientSecretConfig) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineHookOAuthClientSecretConfig) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineHookOAuthClientSecretConfig) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineHookOAuthClientSecretConfig) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


