# UserFactorCallProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneExtension** | Pointer to **NullableString** | Extension of the associated &#x60;phoneNumber&#x60; | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number of the Factor. You should format phone numbers to use the [E.164 standard](https://www.itu.int/rec/T-REC-E.164/). | [optional] 

## Methods

### NewUserFactorCallProfile

`func NewUserFactorCallProfile() *UserFactorCallProfile`

NewUserFactorCallProfile instantiates a new UserFactorCallProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorCallProfileWithDefaults

`func NewUserFactorCallProfileWithDefaults() *UserFactorCallProfile`

NewUserFactorCallProfileWithDefaults instantiates a new UserFactorCallProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneExtension

`func (o *UserFactorCallProfile) GetPhoneExtension() string`

GetPhoneExtension returns the PhoneExtension field if non-nil, zero value otherwise.

### GetPhoneExtensionOk

`func (o *UserFactorCallProfile) GetPhoneExtensionOk() (*string, bool)`

GetPhoneExtensionOk returns a tuple with the PhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneExtension

`func (o *UserFactorCallProfile) SetPhoneExtension(v string)`

SetPhoneExtension sets PhoneExtension field to given value.

### HasPhoneExtension

`func (o *UserFactorCallProfile) HasPhoneExtension() bool`

HasPhoneExtension returns a boolean if a field has been set.

### SetPhoneExtensionNil

`func (o *UserFactorCallProfile) SetPhoneExtensionNil(b bool)`

 SetPhoneExtensionNil sets the value for PhoneExtension to be an explicit nil

### UnsetPhoneExtension
`func (o *UserFactorCallProfile) UnsetPhoneExtension()`

UnsetPhoneExtension ensures that no value is present for PhoneExtension, not even an explicit nil
### GetPhoneNumber

`func (o *UserFactorCallProfile) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserFactorCallProfile) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserFactorCallProfile) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserFactorCallProfile) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


