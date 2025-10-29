# PasswordPolicyPasswordSettingsBreachedProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedWorkflowId** | Pointer to **NullableString** | The &#x60;id&#x60; of the workflow that runs when a breached password is found during a sign-in attempt. | [optional] 
**ExpireAfterDays** | Pointer to **NullableInt32** | Specifies the number of days after a breached password is found during a sign-in attempt that the user&#39;s password should expire. Valid values: 0 through 10. If set to 0, it happens immediately. | [optional] 
**LogoutEnabled** | Pointer to **NullableBool** | (Optional, default is false) If true, you must also specify a value for &#x60;expireAfterDays&#x60;. When enabled, the user&#39;s session(s) are terminated immediately the first time the user&#39;s credentials are detected as part of a breach. | [optional] [default to false]

## Methods

### NewPasswordPolicyPasswordSettingsBreachedProtection

`func NewPasswordPolicyPasswordSettingsBreachedProtection() *PasswordPolicyPasswordSettingsBreachedProtection`

NewPasswordPolicyPasswordSettingsBreachedProtection instantiates a new PasswordPolicyPasswordSettingsBreachedProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyPasswordSettingsBreachedProtectionWithDefaults

`func NewPasswordPolicyPasswordSettingsBreachedProtectionWithDefaults() *PasswordPolicyPasswordSettingsBreachedProtection`

NewPasswordPolicyPasswordSettingsBreachedProtectionWithDefaults instantiates a new PasswordPolicyPasswordSettingsBreachedProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedWorkflowId

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetDelegatedWorkflowId() string`

GetDelegatedWorkflowId returns the DelegatedWorkflowId field if non-nil, zero value otherwise.

### GetDelegatedWorkflowIdOk

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetDelegatedWorkflowIdOk() (*string, bool)`

GetDelegatedWorkflowIdOk returns a tuple with the DelegatedWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedWorkflowId

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetDelegatedWorkflowId(v string)`

SetDelegatedWorkflowId sets DelegatedWorkflowId field to given value.

### HasDelegatedWorkflowId

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) HasDelegatedWorkflowId() bool`

HasDelegatedWorkflowId returns a boolean if a field has been set.

### SetDelegatedWorkflowIdNil

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetDelegatedWorkflowIdNil(b bool)`

 SetDelegatedWorkflowIdNil sets the value for DelegatedWorkflowId to be an explicit nil

### UnsetDelegatedWorkflowId
`func (o *PasswordPolicyPasswordSettingsBreachedProtection) UnsetDelegatedWorkflowId()`

UnsetDelegatedWorkflowId ensures that no value is present for DelegatedWorkflowId, not even an explicit nil
### GetExpireAfterDays

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetExpireAfterDays() int32`

GetExpireAfterDays returns the ExpireAfterDays field if non-nil, zero value otherwise.

### GetExpireAfterDaysOk

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetExpireAfterDaysOk() (*int32, bool)`

GetExpireAfterDaysOk returns a tuple with the ExpireAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfterDays

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetExpireAfterDays(v int32)`

SetExpireAfterDays sets ExpireAfterDays field to given value.

### HasExpireAfterDays

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) HasExpireAfterDays() bool`

HasExpireAfterDays returns a boolean if a field has been set.

### SetExpireAfterDaysNil

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetExpireAfterDaysNil(b bool)`

 SetExpireAfterDaysNil sets the value for ExpireAfterDays to be an explicit nil

### UnsetExpireAfterDays
`func (o *PasswordPolicyPasswordSettingsBreachedProtection) UnsetExpireAfterDays()`

UnsetExpireAfterDays ensures that no value is present for ExpireAfterDays, not even an explicit nil
### GetLogoutEnabled

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetLogoutEnabled() bool`

GetLogoutEnabled returns the LogoutEnabled field if non-nil, zero value otherwise.

### GetLogoutEnabledOk

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) GetLogoutEnabledOk() (*bool, bool)`

GetLogoutEnabledOk returns a tuple with the LogoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutEnabled

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetLogoutEnabled(v bool)`

SetLogoutEnabled sets LogoutEnabled field to given value.

### HasLogoutEnabled

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) HasLogoutEnabled() bool`

HasLogoutEnabled returns a boolean if a field has been set.

### SetLogoutEnabledNil

`func (o *PasswordPolicyPasswordSettingsBreachedProtection) SetLogoutEnabledNil(b bool)`

 SetLogoutEnabledNil sets the value for LogoutEnabled to be an explicit nil

### UnsetLogoutEnabled
`func (o *PasswordPolicyPasswordSettingsBreachedProtection) UnsetLogoutEnabled()`

UnsetLogoutEnabled ensures that no value is present for LogoutEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


