# AuthenticatorMethodOtp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptableAdjacentIntervals** | Pointer to **int32** | The number of acceptable adjacent intervals, also known as the clock drift interval. This setting allows you to build in tolerance for any time difference between the token and the server. For example, with a &#x60;timeIntervalInSeconds&#x60; of 60 seconds and an &#x60;acceptableAdjacentIntervals&#x60; value of 5, Okta accepts passcodes within 300 seconds (60 * 5) before or after the end user enters their code. | [optional] 
**Algorithm** | Pointer to **string** | HMAC algorithm | [optional] 
**Encoding** | Pointer to **string** | The shared secret encoding | [optional] 
**FactorProfileId** | Pointer to **string** | The &#x60;id&#x60; value of the factor profile | [optional] 
**PassCodeLength** | Pointer to **int32** | Number of digits in an OTP value | [optional] 
**Protocol** | Pointer to **string** | The protocol used | [optional] 
**TimeIntervalInSeconds** | Pointer to **int32** | Time interval for TOTP in seconds | [optional] 

## Methods

### NewAuthenticatorMethodOtp

`func NewAuthenticatorMethodOtp() *AuthenticatorMethodOtp`

NewAuthenticatorMethodOtp instantiates a new AuthenticatorMethodOtp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorMethodOtpWithDefaults

`func NewAuthenticatorMethodOtpWithDefaults() *AuthenticatorMethodOtp`

NewAuthenticatorMethodOtpWithDefaults instantiates a new AuthenticatorMethodOtp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptableAdjacentIntervals

`func (o *AuthenticatorMethodOtp) GetAcceptableAdjacentIntervals() int32`

GetAcceptableAdjacentIntervals returns the AcceptableAdjacentIntervals field if non-nil, zero value otherwise.

### GetAcceptableAdjacentIntervalsOk

`func (o *AuthenticatorMethodOtp) GetAcceptableAdjacentIntervalsOk() (*int32, bool)`

GetAcceptableAdjacentIntervalsOk returns a tuple with the AcceptableAdjacentIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableAdjacentIntervals

`func (o *AuthenticatorMethodOtp) SetAcceptableAdjacentIntervals(v int32)`

SetAcceptableAdjacentIntervals sets AcceptableAdjacentIntervals field to given value.

### HasAcceptableAdjacentIntervals

`func (o *AuthenticatorMethodOtp) HasAcceptableAdjacentIntervals() bool`

HasAcceptableAdjacentIntervals returns a boolean if a field has been set.

### GetAlgorithm

`func (o *AuthenticatorMethodOtp) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *AuthenticatorMethodOtp) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *AuthenticatorMethodOtp) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *AuthenticatorMethodOtp) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetEncoding

`func (o *AuthenticatorMethodOtp) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *AuthenticatorMethodOtp) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *AuthenticatorMethodOtp) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *AuthenticatorMethodOtp) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetFactorProfileId

`func (o *AuthenticatorMethodOtp) GetFactorProfileId() string`

GetFactorProfileId returns the FactorProfileId field if non-nil, zero value otherwise.

### GetFactorProfileIdOk

`func (o *AuthenticatorMethodOtp) GetFactorProfileIdOk() (*string, bool)`

GetFactorProfileIdOk returns a tuple with the FactorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorProfileId

`func (o *AuthenticatorMethodOtp) SetFactorProfileId(v string)`

SetFactorProfileId sets FactorProfileId field to given value.

### HasFactorProfileId

`func (o *AuthenticatorMethodOtp) HasFactorProfileId() bool`

HasFactorProfileId returns a boolean if a field has been set.

### GetPassCodeLength

`func (o *AuthenticatorMethodOtp) GetPassCodeLength() int32`

GetPassCodeLength returns the PassCodeLength field if non-nil, zero value otherwise.

### GetPassCodeLengthOk

`func (o *AuthenticatorMethodOtp) GetPassCodeLengthOk() (*int32, bool)`

GetPassCodeLengthOk returns a tuple with the PassCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCodeLength

`func (o *AuthenticatorMethodOtp) SetPassCodeLength(v int32)`

SetPassCodeLength sets PassCodeLength field to given value.

### HasPassCodeLength

`func (o *AuthenticatorMethodOtp) HasPassCodeLength() bool`

HasPassCodeLength returns a boolean if a field has been set.

### GetProtocol

`func (o *AuthenticatorMethodOtp) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AuthenticatorMethodOtp) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AuthenticatorMethodOtp) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AuthenticatorMethodOtp) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeIntervalInSeconds

`func (o *AuthenticatorMethodOtp) GetTimeIntervalInSeconds() int32`

GetTimeIntervalInSeconds returns the TimeIntervalInSeconds field if non-nil, zero value otherwise.

### GetTimeIntervalInSecondsOk

`func (o *AuthenticatorMethodOtp) GetTimeIntervalInSecondsOk() (*int32, bool)`

GetTimeIntervalInSecondsOk returns a tuple with the TimeIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalInSeconds

`func (o *AuthenticatorMethodOtp) SetTimeIntervalInSeconds(v int32)`

SetTimeIntervalInSeconds sets TimeIntervalInSeconds field to given value.

### HasTimeIntervalInSeconds

`func (o *AuthenticatorMethodOtp) HasTimeIntervalInSeconds() bool`

HasTimeIntervalInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


