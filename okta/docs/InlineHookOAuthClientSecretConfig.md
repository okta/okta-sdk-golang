# InlineHookOAuthClientSecretConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthScheme** | Pointer to **NullableString** | Not applicable. Must be &#x60;null&#x60;. | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** | A publicly exposed string provided by the service that&#39;s used to identify the OAuth app and build authorization URLs | [optional] 
**Scope** | Pointer to **string** | Include the scopes that allow you to perform the actions on the hook endpoint that you want to access | [optional] 
**TokenUrl** | Pointer to **string** | The URI where inline hooks can exchange an authorization code for access and refresh tokens | [optional] 
**Headers** | Pointer to [**[]InlineHookChannelConfigHeaders**](InlineHookChannelConfigHeaders.md) | An optional list of key/value pairs for headers that you can send with the request to the external service | [optional] 
**Method** | Pointer to **string** | The method of the Okta inline hook request | [optional] 
**Uri** | Pointer to **string** | The external service endpoint that executes the inline hook handler. It must begin with &#x60;https://&#x60; and be reachable by Okta. No white space is allowed in the URI. | [optional] 

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

### GetAuthScheme

`func (o *InlineHookOAuthClientSecretConfig) GetAuthScheme() string`

GetAuthScheme returns the AuthScheme field if non-nil, zero value otherwise.

### GetAuthSchemeOk

`func (o *InlineHookOAuthClientSecretConfig) GetAuthSchemeOk() (*string, bool)`

GetAuthSchemeOk returns a tuple with the AuthScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthScheme

`func (o *InlineHookOAuthClientSecretConfig) SetAuthScheme(v string)`

SetAuthScheme sets AuthScheme field to given value.

### HasAuthScheme

`func (o *InlineHookOAuthClientSecretConfig) HasAuthScheme() bool`

HasAuthScheme returns a boolean if a field has been set.

### SetAuthSchemeNil

`func (o *InlineHookOAuthClientSecretConfig) SetAuthSchemeNil(b bool)`

 SetAuthSchemeNil sets the value for AuthScheme to be an explicit nil

### UnsetAuthScheme
`func (o *InlineHookOAuthClientSecretConfig) UnsetAuthScheme()`

UnsetAuthScheme ensures that no value is present for AuthScheme, not even an explicit nil
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


