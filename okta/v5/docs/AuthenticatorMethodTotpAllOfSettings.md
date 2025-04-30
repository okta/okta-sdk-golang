# AuthenticatorMethodTotpAllOfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeIntervalInSeconds** | Pointer to **int32** | Time interval for TOTP in seconds | [optional] 
**Encoding** | Pointer to **string** | The shared secret encoding | [optional] 
**Algorithm** | Pointer to **string** | HMAC algorithm | [optional] 
**PassCodeLength** | Pointer to **int32** | Number of digits in an OTP value | [optional] 

## Methods

### NewAuthenticatorMethodTotpAllOfSettings

`func NewAuthenticatorMethodTotpAllOfSettings() *AuthenticatorMethodTotpAllOfSettings`

NewAuthenticatorMethodTotpAllOfSettings instantiates a new AuthenticatorMethodTotpAllOfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorMethodTotpAllOfSettingsWithDefaults

`func NewAuthenticatorMethodTotpAllOfSettingsWithDefaults() *AuthenticatorMethodTotpAllOfSettings`

NewAuthenticatorMethodTotpAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodTotpAllOfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeIntervalInSeconds

`func (o *AuthenticatorMethodTotpAllOfSettings) GetTimeIntervalInSeconds() int32`

GetTimeIntervalInSeconds returns the TimeIntervalInSeconds field if non-nil, zero value otherwise.

### GetTimeIntervalInSecondsOk

`func (o *AuthenticatorMethodTotpAllOfSettings) GetTimeIntervalInSecondsOk() (*int32, bool)`

GetTimeIntervalInSecondsOk returns a tuple with the TimeIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalInSeconds

`func (o *AuthenticatorMethodTotpAllOfSettings) SetTimeIntervalInSeconds(v int32)`

SetTimeIntervalInSeconds sets TimeIntervalInSeconds field to given value.

### HasTimeIntervalInSeconds

`func (o *AuthenticatorMethodTotpAllOfSettings) HasTimeIntervalInSeconds() bool`

HasTimeIntervalInSeconds returns a boolean if a field has been set.

### GetEncoding

`func (o *AuthenticatorMethodTotpAllOfSettings) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *AuthenticatorMethodTotpAllOfSettings) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *AuthenticatorMethodTotpAllOfSettings) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *AuthenticatorMethodTotpAllOfSettings) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetAlgorithm

`func (o *AuthenticatorMethodTotpAllOfSettings) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *AuthenticatorMethodTotpAllOfSettings) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *AuthenticatorMethodTotpAllOfSettings) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *AuthenticatorMethodTotpAllOfSettings) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetPassCodeLength

`func (o *AuthenticatorMethodTotpAllOfSettings) GetPassCodeLength() int32`

GetPassCodeLength returns the PassCodeLength field if non-nil, zero value otherwise.

### GetPassCodeLengthOk

`func (o *AuthenticatorMethodTotpAllOfSettings) GetPassCodeLengthOk() (*int32, bool)`

GetPassCodeLengthOk returns a tuple with the PassCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCodeLength

`func (o *AuthenticatorMethodTotpAllOfSettings) SetPassCodeLength(v int32)`

SetPassCodeLength sets PassCodeLength field to given value.

### HasPassCodeLength

`func (o *AuthenticatorMethodTotpAllOfSettings) HasPassCodeLength() bool`

HasPassCodeLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


