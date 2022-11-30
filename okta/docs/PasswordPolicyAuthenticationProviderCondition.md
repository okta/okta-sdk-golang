# PasswordPolicyAuthenticationProviderCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Include** | Pointer to **[]string** |  | [optional] 
**Provider** | Pointer to [**PasswordPolicyAuthenticationProviderType**](PasswordPolicyAuthenticationProviderType.md) |  | [optional] 

## Methods

### NewPasswordPolicyAuthenticationProviderCondition

`func NewPasswordPolicyAuthenticationProviderCondition() *PasswordPolicyAuthenticationProviderCondition`

NewPasswordPolicyAuthenticationProviderCondition instantiates a new PasswordPolicyAuthenticationProviderCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyAuthenticationProviderConditionWithDefaults

`func NewPasswordPolicyAuthenticationProviderConditionWithDefaults() *PasswordPolicyAuthenticationProviderCondition`

NewPasswordPolicyAuthenticationProviderConditionWithDefaults instantiates a new PasswordPolicyAuthenticationProviderCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInclude

`func (o *PasswordPolicyAuthenticationProviderCondition) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *PasswordPolicyAuthenticationProviderCondition) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *PasswordPolicyAuthenticationProviderCondition) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *PasswordPolicyAuthenticationProviderCondition) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetProvider

`func (o *PasswordPolicyAuthenticationProviderCondition) GetProvider() PasswordPolicyAuthenticationProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PasswordPolicyAuthenticationProviderCondition) GetProviderOk() (*PasswordPolicyAuthenticationProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PasswordPolicyAuthenticationProviderCondition) SetProvider(v PasswordPolicyAuthenticationProviderType)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PasswordPolicyAuthenticationProviderCondition) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


