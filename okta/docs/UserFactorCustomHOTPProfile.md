# UserFactorCustomHOTPProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecret** | Pointer to **string** | Unique secret key used to generate the OTP | [optional] 

## Methods

### NewUserFactorCustomHOTPProfile

`func NewUserFactorCustomHOTPProfile() *UserFactorCustomHOTPProfile`

NewUserFactorCustomHOTPProfile instantiates a new UserFactorCustomHOTPProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorCustomHOTPProfileWithDefaults

`func NewUserFactorCustomHOTPProfileWithDefaults() *UserFactorCustomHOTPProfile`

NewUserFactorCustomHOTPProfileWithDefaults instantiates a new UserFactorCustomHOTPProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSecret

`func (o *UserFactorCustomHOTPProfile) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *UserFactorCustomHOTPProfile) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *UserFactorCustomHOTPProfile) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *UserFactorCustomHOTPProfile) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


