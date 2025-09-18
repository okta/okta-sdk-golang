# AuthenticatorEnrollmentPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**AuthenticatorEnrollmentPolicyConditions**](AuthenticatorEnrollmentPolicyConditions.md) |  | [optional] 
**Settings** | Pointer to [**AuthenticatorEnrollmentPolicySettings**](AuthenticatorEnrollmentPolicySettings.md) |  | [optional] 

## Methods

### NewAuthenticatorEnrollmentPolicy

`func NewAuthenticatorEnrollmentPolicy() *AuthenticatorEnrollmentPolicy`

NewAuthenticatorEnrollmentPolicy instantiates a new AuthenticatorEnrollmentPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentPolicyWithDefaults

`func NewAuthenticatorEnrollmentPolicyWithDefaults() *AuthenticatorEnrollmentPolicy`

NewAuthenticatorEnrollmentPolicyWithDefaults instantiates a new AuthenticatorEnrollmentPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *AuthenticatorEnrollmentPolicy) GetConditions() AuthenticatorEnrollmentPolicyConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthenticatorEnrollmentPolicy) GetConditionsOk() (*AuthenticatorEnrollmentPolicyConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthenticatorEnrollmentPolicy) SetConditions(v AuthenticatorEnrollmentPolicyConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AuthenticatorEnrollmentPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetSettings

`func (o *AuthenticatorEnrollmentPolicy) GetSettings() AuthenticatorEnrollmentPolicySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AuthenticatorEnrollmentPolicy) GetSettingsOk() (*AuthenticatorEnrollmentPolicySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AuthenticatorEnrollmentPolicy) SetSettings(v AuthenticatorEnrollmentPolicySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AuthenticatorEnrollmentPolicy) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


