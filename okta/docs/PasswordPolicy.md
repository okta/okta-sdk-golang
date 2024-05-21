# PasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**PasswordPolicyConditions**](PasswordPolicyConditions.md) |  | [optional] 
**Settings** | Pointer to [**PasswordPolicySettings**](PasswordPolicySettings.md) |  | [optional] 

## Methods

### NewPasswordPolicy

`func NewPasswordPolicy() *PasswordPolicy`

NewPasswordPolicy instantiates a new PasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyWithDefaults

`func NewPasswordPolicyWithDefaults() *PasswordPolicy`

NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *PasswordPolicy) GetConditions() PasswordPolicyConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PasswordPolicy) GetConditionsOk() (*PasswordPolicyConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PasswordPolicy) SetConditions(v PasswordPolicyConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PasswordPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetSettings

`func (o *PasswordPolicy) GetSettings() PasswordPolicySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PasswordPolicy) GetSettingsOk() (*PasswordPolicySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PasswordPolicy) SetSettings(v PasswordPolicySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PasswordPolicy) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


