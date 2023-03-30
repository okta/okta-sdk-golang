# PasswordPolicyPasswordSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | Pointer to [**PasswordPolicyPasswordSettingsAge**](PasswordPolicyPasswordSettingsAge.md) |  | [optional] 
**Complexity** | Pointer to [**PasswordPolicyPasswordSettingsComplexity**](PasswordPolicyPasswordSettingsComplexity.md) |  | [optional] 
**Lockout** | Pointer to [**PasswordPolicyPasswordSettingsLockout**](PasswordPolicyPasswordSettingsLockout.md) |  | [optional] 

## Methods

### NewPasswordPolicyPasswordSettings

`func NewPasswordPolicyPasswordSettings() *PasswordPolicyPasswordSettings`

NewPasswordPolicyPasswordSettings instantiates a new PasswordPolicyPasswordSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyPasswordSettingsWithDefaults

`func NewPasswordPolicyPasswordSettingsWithDefaults() *PasswordPolicyPasswordSettings`

NewPasswordPolicyPasswordSettingsWithDefaults instantiates a new PasswordPolicyPasswordSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *PasswordPolicyPasswordSettings) GetAge() PasswordPolicyPasswordSettingsAge`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *PasswordPolicyPasswordSettings) GetAgeOk() (*PasswordPolicyPasswordSettingsAge, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *PasswordPolicyPasswordSettings) SetAge(v PasswordPolicyPasswordSettingsAge)`

SetAge sets Age field to given value.

### HasAge

`func (o *PasswordPolicyPasswordSettings) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetComplexity

`func (o *PasswordPolicyPasswordSettings) GetComplexity() PasswordPolicyPasswordSettingsComplexity`

GetComplexity returns the Complexity field if non-nil, zero value otherwise.

### GetComplexityOk

`func (o *PasswordPolicyPasswordSettings) GetComplexityOk() (*PasswordPolicyPasswordSettingsComplexity, bool)`

GetComplexityOk returns a tuple with the Complexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexity

`func (o *PasswordPolicyPasswordSettings) SetComplexity(v PasswordPolicyPasswordSettingsComplexity)`

SetComplexity sets Complexity field to given value.

### HasComplexity

`func (o *PasswordPolicyPasswordSettings) HasComplexity() bool`

HasComplexity returns a boolean if a field has been set.

### GetLockout

`func (o *PasswordPolicyPasswordSettings) GetLockout() PasswordPolicyPasswordSettingsLockout`

GetLockout returns the Lockout field if non-nil, zero value otherwise.

### GetLockoutOk

`func (o *PasswordPolicyPasswordSettings) GetLockoutOk() (*PasswordPolicyPasswordSettingsLockout, bool)`

GetLockoutOk returns a tuple with the Lockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockout

`func (o *PasswordPolicyPasswordSettings) SetLockout(v PasswordPolicyPasswordSettingsLockout)`

SetLockout sets Lockout field to given value.

### HasLockout

`func (o *PasswordPolicyPasswordSettings) HasLockout() bool`

HasLockout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


