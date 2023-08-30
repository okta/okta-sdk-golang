# AuthenticatorMethodOtpAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptableAdjacentIntervals** | Pointer to **int32** |  | [optional] 
**Algorithm** | Pointer to [**OtpTotpAlgorithm**](OtpTotpAlgorithm.md) |  | [optional] 
**Encoding** | Pointer to [**OtpTotpEncoding**](OtpTotpEncoding.md) |  | [optional] 
**FactorProfileId** | Pointer to **string** |  | [optional] 
**PassCodeLength** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to [**OtpProtocol**](OtpProtocol.md) |  | [optional] 
**TimeIntervalInSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthenticatorMethodOtpAllOf

`func NewAuthenticatorMethodOtpAllOf() *AuthenticatorMethodOtpAllOf`

NewAuthenticatorMethodOtpAllOf instantiates a new AuthenticatorMethodOtpAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorMethodOtpAllOfWithDefaults

`func NewAuthenticatorMethodOtpAllOfWithDefaults() *AuthenticatorMethodOtpAllOf`

NewAuthenticatorMethodOtpAllOfWithDefaults instantiates a new AuthenticatorMethodOtpAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptableAdjacentIntervals

`func (o *AuthenticatorMethodOtpAllOf) GetAcceptableAdjacentIntervals() int32`

GetAcceptableAdjacentIntervals returns the AcceptableAdjacentIntervals field if non-nil, zero value otherwise.

### GetAcceptableAdjacentIntervalsOk

`func (o *AuthenticatorMethodOtpAllOf) GetAcceptableAdjacentIntervalsOk() (*int32, bool)`

GetAcceptableAdjacentIntervalsOk returns a tuple with the AcceptableAdjacentIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableAdjacentIntervals

`func (o *AuthenticatorMethodOtpAllOf) SetAcceptableAdjacentIntervals(v int32)`

SetAcceptableAdjacentIntervals sets AcceptableAdjacentIntervals field to given value.

### HasAcceptableAdjacentIntervals

`func (o *AuthenticatorMethodOtpAllOf) HasAcceptableAdjacentIntervals() bool`

HasAcceptableAdjacentIntervals returns a boolean if a field has been set.

### GetAlgorithm

`func (o *AuthenticatorMethodOtpAllOf) GetAlgorithm() OtpTotpAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *AuthenticatorMethodOtpAllOf) GetAlgorithmOk() (*OtpTotpAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *AuthenticatorMethodOtpAllOf) SetAlgorithm(v OtpTotpAlgorithm)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *AuthenticatorMethodOtpAllOf) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetEncoding

`func (o *AuthenticatorMethodOtpAllOf) GetEncoding() OtpTotpEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *AuthenticatorMethodOtpAllOf) GetEncodingOk() (*OtpTotpEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *AuthenticatorMethodOtpAllOf) SetEncoding(v OtpTotpEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *AuthenticatorMethodOtpAllOf) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetFactorProfileId

`func (o *AuthenticatorMethodOtpAllOf) GetFactorProfileId() string`

GetFactorProfileId returns the FactorProfileId field if non-nil, zero value otherwise.

### GetFactorProfileIdOk

`func (o *AuthenticatorMethodOtpAllOf) GetFactorProfileIdOk() (*string, bool)`

GetFactorProfileIdOk returns a tuple with the FactorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorProfileId

`func (o *AuthenticatorMethodOtpAllOf) SetFactorProfileId(v string)`

SetFactorProfileId sets FactorProfileId field to given value.

### HasFactorProfileId

`func (o *AuthenticatorMethodOtpAllOf) HasFactorProfileId() bool`

HasFactorProfileId returns a boolean if a field has been set.

### GetPassCodeLength

`func (o *AuthenticatorMethodOtpAllOf) GetPassCodeLength() int32`

GetPassCodeLength returns the PassCodeLength field if non-nil, zero value otherwise.

### GetPassCodeLengthOk

`func (o *AuthenticatorMethodOtpAllOf) GetPassCodeLengthOk() (*int32, bool)`

GetPassCodeLengthOk returns a tuple with the PassCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCodeLength

`func (o *AuthenticatorMethodOtpAllOf) SetPassCodeLength(v int32)`

SetPassCodeLength sets PassCodeLength field to given value.

### HasPassCodeLength

`func (o *AuthenticatorMethodOtpAllOf) HasPassCodeLength() bool`

HasPassCodeLength returns a boolean if a field has been set.

### GetProtocol

`func (o *AuthenticatorMethodOtpAllOf) GetProtocol() OtpProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AuthenticatorMethodOtpAllOf) GetProtocolOk() (*OtpProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AuthenticatorMethodOtpAllOf) SetProtocol(v OtpProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AuthenticatorMethodOtpAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeIntervalInSeconds

`func (o *AuthenticatorMethodOtpAllOf) GetTimeIntervalInSeconds() int32`

GetTimeIntervalInSeconds returns the TimeIntervalInSeconds field if non-nil, zero value otherwise.

### GetTimeIntervalInSecondsOk

`func (o *AuthenticatorMethodOtpAllOf) GetTimeIntervalInSecondsOk() (*int32, bool)`

GetTimeIntervalInSecondsOk returns a tuple with the TimeIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalInSeconds

`func (o *AuthenticatorMethodOtpAllOf) SetTimeIntervalInSeconds(v int32)`

SetTimeIntervalInSeconds sets TimeIntervalInSeconds field to given value.

### HasTimeIntervalInSeconds

`func (o *AuthenticatorMethodOtpAllOf) HasTimeIntervalInSeconds() bool`

HasTimeIntervalInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


