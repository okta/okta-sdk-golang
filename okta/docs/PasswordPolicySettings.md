# PasswordPolicySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delegation** | Pointer to [**PasswordPolicyDelegationSettings**](PasswordPolicyDelegationSettings.md) |  | [optional] 
**Password** | Pointer to [**PasswordPolicyPasswordSettings**](PasswordPolicyPasswordSettings.md) |  | [optional] 
**Recovery** | Pointer to [**PasswordPolicyRecoverySettings**](PasswordPolicyRecoverySettings.md) |  | [optional] 

## Methods

### NewPasswordPolicySettings

`func NewPasswordPolicySettings() *PasswordPolicySettings`

NewPasswordPolicySettings instantiates a new PasswordPolicySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicySettingsWithDefaults

`func NewPasswordPolicySettingsWithDefaults() *PasswordPolicySettings`

NewPasswordPolicySettingsWithDefaults instantiates a new PasswordPolicySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegation

`func (o *PasswordPolicySettings) GetDelegation() PasswordPolicyDelegationSettings`

GetDelegation returns the Delegation field if non-nil, zero value otherwise.

### GetDelegationOk

`func (o *PasswordPolicySettings) GetDelegationOk() (*PasswordPolicyDelegationSettings, bool)`

GetDelegationOk returns a tuple with the Delegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegation

`func (o *PasswordPolicySettings) SetDelegation(v PasswordPolicyDelegationSettings)`

SetDelegation sets Delegation field to given value.

### HasDelegation

`func (o *PasswordPolicySettings) HasDelegation() bool`

HasDelegation returns a boolean if a field has been set.

### GetPassword

`func (o *PasswordPolicySettings) GetPassword() PasswordPolicyPasswordSettings`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordPolicySettings) GetPasswordOk() (*PasswordPolicyPasswordSettings, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordPolicySettings) SetPassword(v PasswordPolicyPasswordSettings)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PasswordPolicySettings) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRecovery

`func (o *PasswordPolicySettings) GetRecovery() PasswordPolicyRecoverySettings`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *PasswordPolicySettings) GetRecoveryOk() (*PasswordPolicyRecoverySettings, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *PasswordPolicySettings) SetRecovery(v PasswordPolicyRecoverySettings)`

SetRecovery sets Recovery field to given value.

### HasRecovery

`func (o *PasswordPolicySettings) HasRecovery() bool`

HasRecovery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


