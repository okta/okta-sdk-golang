# UserFactorTokenHOTPProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedSecret** | Pointer to **string** | Unique secret key used to generate the OTP | [optional] 

## Methods

### NewUserFactorTokenHOTPProfile

`func NewUserFactorTokenHOTPProfile() *UserFactorTokenHOTPProfile`

NewUserFactorTokenHOTPProfile instantiates a new UserFactorTokenHOTPProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorTokenHOTPProfileWithDefaults

`func NewUserFactorTokenHOTPProfileWithDefaults() *UserFactorTokenHOTPProfile`

NewUserFactorTokenHOTPProfileWithDefaults instantiates a new UserFactorTokenHOTPProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedSecret

`func (o *UserFactorTokenHOTPProfile) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *UserFactorTokenHOTPProfile) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *UserFactorTokenHOTPProfile) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *UserFactorTokenHOTPProfile) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


