# AuthenticatorProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | The phone number for a &#x60;call&#x60; or &#x60;sms&#x60; authenticator enrollment. | 

## Methods

### NewAuthenticatorProfile

`func NewAuthenticatorProfile(phoneNumber string, ) *AuthenticatorProfile`

NewAuthenticatorProfile instantiates a new AuthenticatorProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorProfileWithDefaults

`func NewAuthenticatorProfileWithDefaults() *AuthenticatorProfile`

NewAuthenticatorProfileWithDefaults instantiates a new AuthenticatorProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *AuthenticatorProfile) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AuthenticatorProfile) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AuthenticatorProfile) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


