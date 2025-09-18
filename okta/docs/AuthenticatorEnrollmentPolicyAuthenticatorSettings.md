# AuthenticatorEnrollmentPolicyAuthenticatorSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**NullableAuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints**](AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints.md) |  | [optional] 
**Enroll** | Pointer to [**AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll**](AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll.md) |  | [optional] 
**Key** | Pointer to **string** | A label that identifies the authenticator | [optional] 

## Methods

### NewAuthenticatorEnrollmentPolicyAuthenticatorSettings

`func NewAuthenticatorEnrollmentPolicyAuthenticatorSettings() *AuthenticatorEnrollmentPolicyAuthenticatorSettings`

NewAuthenticatorEnrollmentPolicyAuthenticatorSettings instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsWithDefaults

`func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsWithDefaults() *AuthenticatorEnrollmentPolicyAuthenticatorSettings`

NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsWithDefaults instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetConstraints() AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetConstraintsOk() (*AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) SetConstraints(v AuthenticatorEnrollmentPolicyAuthenticatorSettingsConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### SetConstraintsNil

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) SetConstraintsNil(b bool)`

 SetConstraintsNil sets the value for Constraints to be an explicit nil

### UnsetConstraints
`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) UnsetConstraints()`

UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
### GetEnroll

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetEnroll() AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetEnrollOk() (*AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) SetEnroll(v AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll)`

SetEnroll sets Enroll field to given value.

### HasEnroll

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) HasEnroll() bool`

HasEnroll returns a boolean if a field has been set.

### GetKey

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettings) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


