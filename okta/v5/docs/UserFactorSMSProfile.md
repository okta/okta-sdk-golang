# UserFactorSMSProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** | Phone number of the Factor. You should format phone numbers to use the [E.164 standard](https://www.itu.int/rec/T-REC-E.164/). | [optional] 

## Methods

### NewUserFactorSMSProfile

`func NewUserFactorSMSProfile() *UserFactorSMSProfile`

NewUserFactorSMSProfile instantiates a new UserFactorSMSProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorSMSProfileWithDefaults

`func NewUserFactorSMSProfileWithDefaults() *UserFactorSMSProfile`

NewUserFactorSMSProfileWithDefaults instantiates a new UserFactorSMSProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *UserFactorSMSProfile) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserFactorSMSProfile) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserFactorSMSProfile) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserFactorSMSProfile) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


