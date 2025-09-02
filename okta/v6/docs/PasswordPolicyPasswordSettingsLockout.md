# PasswordPolicyPasswordSettingsLockout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUnlockMinutes** | Pointer to **int32** | Specifies the time interval (in minutes) a locked account remains locked before it is automatically unlocked: &#x60;0&#x60; indicates no limit | [optional] [default to 0]
**MaxAttempts** | Pointer to **int32** | Specifies the number of times Users can attempt to sign in to their accounts with an invalid password before their accounts are locked: &#x60;0&#x60; indicates no limit | [optional] [default to 10]
**ShowLockoutFailures** | Pointer to **bool** | Indicates if the User should be informed when their account is locked | [optional] [default to false]
**UserLockoutNotificationChannels** | Pointer to **[]string** | How the user is notified when their account becomes locked. The only acceptable values are &#x60;[]&#x60; and &#x60;[&#39;EMAIL&#39;]&#x60;. | [optional] [default to []]

## Methods

### NewPasswordPolicyPasswordSettingsLockout

`func NewPasswordPolicyPasswordSettingsLockout() *PasswordPolicyPasswordSettingsLockout`

NewPasswordPolicyPasswordSettingsLockout instantiates a new PasswordPolicyPasswordSettingsLockout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyPasswordSettingsLockoutWithDefaults

`func NewPasswordPolicyPasswordSettingsLockoutWithDefaults() *PasswordPolicyPasswordSettingsLockout`

NewPasswordPolicyPasswordSettingsLockoutWithDefaults instantiates a new PasswordPolicyPasswordSettingsLockout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUnlockMinutes

`func (o *PasswordPolicyPasswordSettingsLockout) GetAutoUnlockMinutes() int32`

GetAutoUnlockMinutes returns the AutoUnlockMinutes field if non-nil, zero value otherwise.

### GetAutoUnlockMinutesOk

`func (o *PasswordPolicyPasswordSettingsLockout) GetAutoUnlockMinutesOk() (*int32, bool)`

GetAutoUnlockMinutesOk returns a tuple with the AutoUnlockMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUnlockMinutes

`func (o *PasswordPolicyPasswordSettingsLockout) SetAutoUnlockMinutes(v int32)`

SetAutoUnlockMinutes sets AutoUnlockMinutes field to given value.

### HasAutoUnlockMinutes

`func (o *PasswordPolicyPasswordSettingsLockout) HasAutoUnlockMinutes() bool`

HasAutoUnlockMinutes returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *PasswordPolicyPasswordSettingsLockout) GetMaxAttempts() int32`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *PasswordPolicyPasswordSettingsLockout) GetMaxAttemptsOk() (*int32, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *PasswordPolicyPasswordSettingsLockout) SetMaxAttempts(v int32)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *PasswordPolicyPasswordSettingsLockout) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetShowLockoutFailures

`func (o *PasswordPolicyPasswordSettingsLockout) GetShowLockoutFailures() bool`

GetShowLockoutFailures returns the ShowLockoutFailures field if non-nil, zero value otherwise.

### GetShowLockoutFailuresOk

`func (o *PasswordPolicyPasswordSettingsLockout) GetShowLockoutFailuresOk() (*bool, bool)`

GetShowLockoutFailuresOk returns a tuple with the ShowLockoutFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLockoutFailures

`func (o *PasswordPolicyPasswordSettingsLockout) SetShowLockoutFailures(v bool)`

SetShowLockoutFailures sets ShowLockoutFailures field to given value.

### HasShowLockoutFailures

`func (o *PasswordPolicyPasswordSettingsLockout) HasShowLockoutFailures() bool`

HasShowLockoutFailures returns a boolean if a field has been set.

### GetUserLockoutNotificationChannels

`func (o *PasswordPolicyPasswordSettingsLockout) GetUserLockoutNotificationChannels() []string`

GetUserLockoutNotificationChannels returns the UserLockoutNotificationChannels field if non-nil, zero value otherwise.

### GetUserLockoutNotificationChannelsOk

`func (o *PasswordPolicyPasswordSettingsLockout) GetUserLockoutNotificationChannelsOk() (*[]string, bool)`

GetUserLockoutNotificationChannelsOk returns a tuple with the UserLockoutNotificationChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLockoutNotificationChannels

`func (o *PasswordPolicyPasswordSettingsLockout) SetUserLockoutNotificationChannels(v []string)`

SetUserLockoutNotificationChannels sets UserLockoutNotificationChannels field to given value.

### HasUserLockoutNotificationChannels

`func (o *PasswordPolicyPasswordSettingsLockout) HasUserLockoutNotificationChannels() bool`

HasUserLockoutNotificationChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


