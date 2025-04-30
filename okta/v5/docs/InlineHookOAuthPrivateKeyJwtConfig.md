# InlineHookOAuthPrivateKeyJwtConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HookKeyId** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**TokenUrl** | Pointer to **string** |  | [optional] 
**AuthScheme** | Pointer to [**InlineHookChannelConfigAuthScheme**](InlineHookChannelConfigAuthScheme.md) |  | [optional] 
**Headers** | Pointer to [**[]InlineHookChannelConfigHeaders**](InlineHookChannelConfigHeaders.md) |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineHookOAuthPrivateKeyJwtConfig

`func NewInlineHookOAuthPrivateKeyJwtConfig() *InlineHookOAuthPrivateKeyJwtConfig`

NewInlineHookOAuthPrivateKeyJwtConfig instantiates a new InlineHookOAuthPrivateKeyJwtConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookOAuthPrivateKeyJwtConfigWithDefaults

`func NewInlineHookOAuthPrivateKeyJwtConfigWithDefaults() *InlineHookOAuthPrivateKeyJwtConfig`

NewInlineHookOAuthPrivateKeyJwtConfigWithDefaults instantiates a new InlineHookOAuthPrivateKeyJwtConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHookKeyId

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetHookKeyId() string`

GetHookKeyId returns the HookKeyId field if non-nil, zero value otherwise.

### GetHookKeyIdOk

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetHookKeyIdOk() (*string, bool)`

GetHookKeyIdOk returns a tuple with the HookKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookKeyId

`func (o *InlineHookOAuthPrivateKeyJwtConfig) SetHookKeyId(v string)`

SetHookKeyId sets HookKeyId field to given value.

### HasHookKeyId

`func (o *InlineHookOAuthPrivateKeyJwtConfig) HasHookKeyId() bool`

HasHookKeyId returns a boolean if a field has been set.

### GetAuthType

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *InlineHookOAuthPrivateKeyJwtConfig) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *InlineHookOAuthPrivateKeyJwtConfig) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineHookOAuthPrivateKeyJwtConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineHookOAuthPrivateKeyJwtConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetScope

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineHookOAuthPrivateKeyJwtConfig) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineHookOAuthPrivateKeyJwtConfig) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTokenUrl

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *InlineHookOAuthPrivateKeyJwtConfig) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *InlineHookOAuthPrivateKeyJwtConfig) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### GetAuthScheme

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetAuthScheme() InlineHookChannelConfigAuthScheme`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetAuthSchemeOk() (*InlineHookChannelConfigAuthScheme, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *InlineHookOAuthPrivateKeyJwtConfig) SetAuthScheme(v InlineHookChannelConfigAuthScheme)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *InlineHookOAuthPrivateKeyJwtConfig) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetHeaders() []InlineHookChannelConfigHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetHeadersOk() (*[]InlineHookChannelConfigHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineHookOAuthPrivateKeyJwtConfig) SetHeaders(v []InlineHookChannelConfigHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineHookOAuthPrivateKeyJwtConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookOAuthPrivateKeyJwtConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookOAuthPrivateKeyJwtConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineHookOAuthPrivateKeyJwtConfig) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineHookOAuthPrivateKeyJwtConfig) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineHookOAuthPrivateKeyJwtConfig) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


