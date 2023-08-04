# InlineHookOAuthBasicConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**TokenUrl** | Pointer to **string** |  | [optional] 
**AuthScheme** | Pointer to [**InlineHookChannelConfigAuthScheme**](InlineHookChannelConfigAuthScheme.md) |  | [optional] 
**Headers** | Pointer to [**[]InlineHookChannelConfigHeaders**](InlineHookChannelConfigHeaders.md) |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Uri** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineHookOAuthBasicConfig

`func NewInlineHookOAuthBasicConfig() *InlineHookOAuthBasicConfig`

NewInlineHookOAuthBasicConfig instantiates a new InlineHookOAuthBasicConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookOAuthBasicConfigWithDefaults

`func NewInlineHookOAuthBasicConfigWithDefaults() *InlineHookOAuthBasicConfig`

NewInlineHookOAuthBasicConfigWithDefaults instantiates a new InlineHookOAuthBasicConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *InlineHookOAuthBasicConfig) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *InlineHookOAuthBasicConfig) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *InlineHookOAuthBasicConfig) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *InlineHookOAuthBasicConfig) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *InlineHookOAuthBasicConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineHookOAuthBasicConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineHookOAuthBasicConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineHookOAuthBasicConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetScope

`func (o *InlineHookOAuthBasicConfig) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineHookOAuthBasicConfig) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineHookOAuthBasicConfig) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineHookOAuthBasicConfig) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTokenUrl

`func (o *InlineHookOAuthBasicConfig) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *InlineHookOAuthBasicConfig) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *InlineHookOAuthBasicConfig) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *InlineHookOAuthBasicConfig) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### GetAuthScheme

`func (o *InlineHookOAuthBasicConfig) GetAuthScheme() InlineHookChannelConfigAuthScheme`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *InlineHookOAuthBasicConfig) GetAuthSchemeOk() (*InlineHookChannelConfigAuthScheme, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *InlineHookOAuthBasicConfig) SetAuthScheme(v InlineHookChannelConfigAuthScheme)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *InlineHookOAuthBasicConfig) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineHookOAuthBasicConfig) GetHeaders() []InlineHookChannelConfigHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineHookOAuthBasicConfig) GetHeadersOk() (*[]InlineHookChannelConfigHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineHookOAuthBasicConfig) SetHeaders(v []InlineHookChannelConfigHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineHookOAuthBasicConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookOAuthBasicConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookOAuthBasicConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookOAuthBasicConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookOAuthBasicConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetUri

`func (o *InlineHookOAuthBasicConfig) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineHookOAuthBasicConfig) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineHookOAuthBasicConfig) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineHookOAuthBasicConfig) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


