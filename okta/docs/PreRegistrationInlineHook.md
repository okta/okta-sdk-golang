# PreRegistrationInlineHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineHookId** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** | Specifies the use cases when the configured pre-registration inline hook is invoked. This property is required when &#x60;preRegistrationInlineHooks&#x60; is set. Include both values to invoke the hook during self-service registration and progressive profiling. | [optional] 

## Methods

### NewPreRegistrationInlineHook

`func NewPreRegistrationInlineHook() *PreRegistrationInlineHook`

NewPreRegistrationInlineHook instantiates a new PreRegistrationInlineHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreRegistrationInlineHookWithDefaults

`func NewPreRegistrationInlineHookWithDefaults() *PreRegistrationInlineHook`

NewPreRegistrationInlineHookWithDefaults instantiates a new PreRegistrationInlineHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineHookId

`func (o *PreRegistrationInlineHook) GetInlineHookId() string`

GetInlineHookId returns the InlineHookId field if non-nil, zero value otherwise.

### GetInlineHookIdOk

`func (o *PreRegistrationInlineHook) GetInlineHookIdOk() (*string, bool)`

GetInlineHookIdOk returns a tuple with the InlineHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHookId

`func (o *PreRegistrationInlineHook) SetInlineHookId(v string)`

SetInlineHookId sets InlineHookId field to given value.

### HasInlineHookId

`func (o *PreRegistrationInlineHook) HasInlineHookId() bool`

HasInlineHookId returns a boolean if a field has been set.

### GetScopes

`func (o *PreRegistrationInlineHook) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PreRegistrationInlineHook) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PreRegistrationInlineHook) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PreRegistrationInlineHook) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


