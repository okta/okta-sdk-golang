# AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | Requirements for the user-initiated enrollment | [optional] [default to "NOT_ALLOWED"]
**GracePeriod** | Pointer to [**EnrollmentPolicyAuthenticatorGracePeriod**](EnrollmentPolicyAuthenticatorGracePeriod.md) |  | [optional] 

## Methods

### NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll

`func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll`

NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults

`func NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults() *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll`

NewAuthenticatorEnrollmentPolicyAuthenticatorSettingsEnrollWithDefaults instantiates a new AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetGracePeriod

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetGracePeriod() EnrollmentPolicyAuthenticatorGracePeriod`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) GetGracePeriodOk() (*EnrollmentPolicyAuthenticatorGracePeriod, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) SetGracePeriod(v EnrollmentPolicyAuthenticatorGracePeriod)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *AuthenticatorEnrollmentPolicyAuthenticatorSettingsEnroll) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


