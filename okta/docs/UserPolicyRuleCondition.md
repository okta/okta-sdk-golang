# UserPolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | Pointer to **[]string** |  | [optional] 
**Inactivity** | Pointer to [**InactivityPolicyRuleCondition**](InactivityPolicyRuleCondition.md) |  | [optional] 
**Include** | Pointer to **[]string** |  | [optional] 
**LifecycleExpiration** | Pointer to [**LifecycleExpirationPolicyRuleCondition**](LifecycleExpirationPolicyRuleCondition.md) |  | [optional] 
**PasswordExpiration** | Pointer to [**PasswordExpirationPolicyRuleCondition**](PasswordExpirationPolicyRuleCondition.md) |  | [optional] 
**UserLifecycleAttribute** | Pointer to [**UserLifecycleAttributePolicyRuleCondition**](UserLifecycleAttributePolicyRuleCondition.md) |  | [optional] 

## Methods

### NewUserPolicyRuleCondition

`func NewUserPolicyRuleCondition() *UserPolicyRuleCondition`

NewUserPolicyRuleCondition instantiates a new UserPolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPolicyRuleConditionWithDefaults

`func NewUserPolicyRuleConditionWithDefaults() *UserPolicyRuleCondition`

NewUserPolicyRuleConditionWithDefaults instantiates a new UserPolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *UserPolicyRuleCondition) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *UserPolicyRuleCondition) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *UserPolicyRuleCondition) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *UserPolicyRuleCondition) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInactivity

`func (o *UserPolicyRuleCondition) GetInactivity() InactivityPolicyRuleCondition`

GetInactivity returns the Inactivity field if non-nil, zero value otherwise.

### GetInactivityOk

`func (o *UserPolicyRuleCondition) GetInactivityOk() (*InactivityPolicyRuleCondition, bool)`

GetInactivityOk returns a tuple with the Inactivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivity

`func (o *UserPolicyRuleCondition) SetInactivity(v InactivityPolicyRuleCondition)`

SetInactivity sets Inactivity field to given value.

### HasInactivity

`func (o *UserPolicyRuleCondition) HasInactivity() bool`

HasInactivity returns a boolean if a field has been set.

### GetInclude

`func (o *UserPolicyRuleCondition) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *UserPolicyRuleCondition) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *UserPolicyRuleCondition) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *UserPolicyRuleCondition) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### GetLifecycleExpiration

`func (o *UserPolicyRuleCondition) GetLifecycleExpiration() LifecycleExpirationPolicyRuleCondition`

GetLifecycleExpiration returns the LifecycleExpiration field if non-nil, zero value otherwise.

### GetLifecycleExpirationOk

`func (o *UserPolicyRuleCondition) GetLifecycleExpirationOk() (*LifecycleExpirationPolicyRuleCondition, bool)`

GetLifecycleExpirationOk returns a tuple with the LifecycleExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleExpiration

`func (o *UserPolicyRuleCondition) SetLifecycleExpiration(v LifecycleExpirationPolicyRuleCondition)`

SetLifecycleExpiration sets LifecycleExpiration field to given value.

### HasLifecycleExpiration

`func (o *UserPolicyRuleCondition) HasLifecycleExpiration() bool`

HasLifecycleExpiration returns a boolean if a field has been set.

### GetPasswordExpiration

`func (o *UserPolicyRuleCondition) GetPasswordExpiration() PasswordExpirationPolicyRuleCondition`

GetPasswordExpiration returns the PasswordExpiration field if non-nil, zero value otherwise.

### GetPasswordExpirationOk

`func (o *UserPolicyRuleCondition) GetPasswordExpirationOk() (*PasswordExpirationPolicyRuleCondition, bool)`

GetPasswordExpirationOk returns a tuple with the PasswordExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiration

`func (o *UserPolicyRuleCondition) SetPasswordExpiration(v PasswordExpirationPolicyRuleCondition)`

SetPasswordExpiration sets PasswordExpiration field to given value.

### HasPasswordExpiration

`func (o *UserPolicyRuleCondition) HasPasswordExpiration() bool`

HasPasswordExpiration returns a boolean if a field has been set.

### GetUserLifecycleAttribute

`func (o *UserPolicyRuleCondition) GetUserLifecycleAttribute() UserLifecycleAttributePolicyRuleCondition`

GetUserLifecycleAttribute returns the UserLifecycleAttribute field if non-nil, zero value otherwise.

### GetUserLifecycleAttributeOk

`func (o *UserPolicyRuleCondition) GetUserLifecycleAttributeOk() (*UserLifecycleAttributePolicyRuleCondition, bool)`

GetUserLifecycleAttributeOk returns a tuple with the UserLifecycleAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLifecycleAttribute

`func (o *UserPolicyRuleCondition) SetUserLifecycleAttribute(v UserLifecycleAttributePolicyRuleCondition)`

SetUserLifecycleAttribute sets UserLifecycleAttribute field to given value.

### HasUserLifecycleAttribute

`func (o *UserPolicyRuleCondition) HasUserLifecycleAttribute() bool`

HasUserLifecycleAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


