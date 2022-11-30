# PasswordPolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**PasswordPolicyConditions**](PasswordPolicyConditions.md) |  | [optional] 
**Settings** | Pointer to [**PasswordPolicySettings**](PasswordPolicySettings.md) |  | [optional] 

## Methods

### NewPasswordPolicyAllOf

`func NewPasswordPolicyAllOf() *PasswordPolicyAllOf`

NewPasswordPolicyAllOf instantiates a new PasswordPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyAllOfWithDefaults

`func NewPasswordPolicyAllOfWithDefaults() *PasswordPolicyAllOf`

NewPasswordPolicyAllOfWithDefaults instantiates a new PasswordPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *PasswordPolicyAllOf) GetConditions() PasswordPolicyConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PasswordPolicyAllOf) GetConditionsOk() (*PasswordPolicyConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PasswordPolicyAllOf) SetConditions(v PasswordPolicyConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PasswordPolicyAllOf) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetSettings

`func (o *PasswordPolicyAllOf) GetSettings() PasswordPolicySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PasswordPolicyAllOf) GetSettingsOk() (*PasswordPolicySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PasswordPolicyAllOf) SetSettings(v PasswordPolicySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PasswordPolicyAllOf) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


