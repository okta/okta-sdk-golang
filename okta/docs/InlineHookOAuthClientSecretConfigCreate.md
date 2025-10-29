# InlineHookOAuthClientSecretConfigCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | Pointer to **string** | A private value provided by the service used to authenticate the identity of the app to the service | [optional] 
**Method** | Pointer to **string** | The method of the Okta inline hook request. Only accepts &#x60;POST&#x60;. | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** | A publicly exposed string provided by the service that&#39;s used to identify the OAuth app and build authorization URLs | [optional] 
**Scope** | Pointer to **string** | Include the scopes that allow you to perform the actions on the hook endpoint that you want to access | [optional] 
**TokenUrl** | Pointer to **string** | The URI where inline hooks can exchange an authorization code for access and refresh tokens | [optional] 
**Headers** | Pointer to [**[]InlineHookChannelConfigHeaders**](InlineHookChannelConfigHeaders.md) | An optional list of key/value pairs for headers that you can send with the request to the external service | [optional] 
**Uri** | Pointer to **string** | The external service endpoint that executes the inline hook handler. It must begin with &#x60;https://&#x60; and be reachable by Okta. No white space is allowed in the URI. | [optional] 

## Methods

### NewInlineHookOAuthClientSecretConfigCreate

`func NewInlineHookOAuthClientSecretConfigCreate() *InlineHookOAuthClientSecretConfigCreate`

NewInlineHookOAuthClientSecretConfigCreate instantiates a new InlineHookOAuthClientSecretConfigCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookOAuthClientSecretConfigCreateWithDefaults

`func NewInlineHookOAuthClientSecretConfigCreateWithDefaults() *InlineHookOAuthClientSecretConfigCreate`

NewInlineHookOAuthClientSecretConfigCreateWithDefaults instantiates a new InlineHookOAuthClientSecretConfigCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *InlineHookOAuthClientSecretConfigCreate) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *InlineHookOAuthClientSecretConfigCreate) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *InlineHookOAuthClientSecretConfigCreate) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *InlineHookOAuthClientSecretConfigCreate) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetMethod

`func (o *InlineHookOAuthClientSecretConfigCreate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InlineHookOAuthClientSecretConfigCreate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InlineHookOAuthClientSecretConfigCreate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InlineHookOAuthClientSecretConfigCreate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetAuthType

`func (o *InlineHookOAuthClientSecretConfigCreate) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *InlineHookOAuthClientSecretConfigCreate) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *InlineHookOAuthClientSecretConfigCreate) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *InlineHookOAuthClientSecretConfigCreate) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetClientId

`func (o *InlineHookOAuthClientSecretConfigCreate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InlineHookOAuthClientSecretConfigCreate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InlineHookOAuthClientSecretConfigCreate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InlineHookOAuthClientSecretConfigCreate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetScope

`func (o *InlineHookOAuthClientSecretConfigCreate) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineHookOAuthClientSecretConfigCreate) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineHookOAuthClientSecretConfigCreate) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineHookOAuthClientSecretConfigCreate) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTokenUrl

`func (o *InlineHookOAuthClientSecretConfigCreate) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *InlineHookOAuthClientSecretConfigCreate) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *InlineHookOAuthClientSecretConfigCreate) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *InlineHookOAuthClientSecretConfigCreate) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### GetHeaders

`func (o *InlineHookOAuthClientSecretConfigCreate) GetHeaders() []InlineHookChannelConfigHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *InlineHookOAuthClientSecretConfigCreate) GetHeadersOk() (*[]InlineHookChannelConfigHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *InlineHookOAuthClientSecretConfigCreate) SetHeaders(v []InlineHookChannelConfigHeaders)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *InlineHookOAuthClientSecretConfigCreate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUri

`func (o *InlineHookOAuthClientSecretConfigCreate) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *InlineHookOAuthClientSecretConfigCreate) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *InlineHookOAuthClientSecretConfigCreate) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *InlineHookOAuthClientSecretConfigCreate) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


