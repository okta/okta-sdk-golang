# SsprSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowRecoveryEmailWithoutEnrollment** | Pointer to **bool** | Allows a user to recover their password with their email address when they have not enrolled the email authenticator | [optional] 

## Methods

### NewSsprSettings

`func NewSsprSettings() *SsprSettings`

NewSsprSettings instantiates a new SsprSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsprSettingsWithDefaults

`func NewSsprSettingsWithDefaults() *SsprSettings`

NewSsprSettingsWithDefaults instantiates a new SsprSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowRecoveryEmailWithoutEnrollment

`func (o *SsprSettings) GetAllowRecoveryEmailWithoutEnrollment() bool`

GetAllowRecoveryEmailWithoutEnrollment returns the AllowRecoveryEmailWithoutEnrollment field if non-nil, zero value otherwise.

### GetAllowRecoveryEmailWithoutEnrollmentOk

`func (o *SsprSettings) GetAllowRecoveryEmailWithoutEnrollmentOk() (*bool, bool)`

GetAllowRecoveryEmailWithoutEnrollmentOk returns a tuple with the AllowRecoveryEmailWithoutEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRecoveryEmailWithoutEnrollment

`func (o *SsprSettings) SetAllowRecoveryEmailWithoutEnrollment(v bool)`

SetAllowRecoveryEmailWithoutEnrollment sets AllowRecoveryEmailWithoutEnrollment field to given value.

### HasAllowRecoveryEmailWithoutEnrollment

`func (o *SsprSettings) HasAllowRecoveryEmailWithoutEnrollment() bool`

HasAllowRecoveryEmailWithoutEnrollment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


