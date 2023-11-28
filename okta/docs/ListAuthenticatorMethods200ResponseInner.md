# ListAuthenticatorMethods200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**LifecycleStatus**](LifecycleStatus.md) |  | [optional] 
**Type** | Pointer to [**AuthenticatorMethodType**](AuthenticatorMethodType.md) |  | [optional] 
**Links** | Pointer to [**LinksSelfAndLifecycle**](LinksSelfAndLifecycle.md) |  | [optional] 
**Settings** | Pointer to [**AuthenticatorMethodWebAuthnAllOfSettings**](AuthenticatorMethodWebAuthnAllOfSettings.md) |  | [optional] 
**VerifiableProperties** | Pointer to [**[]AuthenticatorMethodProperty**](AuthenticatorMethodProperty.md) |  | [optional] 
**AcceptableAdjacentIntervals** | Pointer to **int32** |  | [optional] 
**Algorithm** | Pointer to [**OtpTotpAlgorithm**](OtpTotpAlgorithm.md) |  | [optional] 
**Encoding** | Pointer to [**OtpTotpEncoding**](OtpTotpEncoding.md) |  | [optional] 
**FactorProfileId** | Pointer to **string** |  | [optional] 
**PassCodeLength** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to [**OtpProtocol**](OtpProtocol.md) |  | [optional] 
**TimeIntervalInSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewListAuthenticatorMethods200ResponseInner

`func NewListAuthenticatorMethods200ResponseInner() *ListAuthenticatorMethods200ResponseInner`

NewListAuthenticatorMethods200ResponseInner instantiates a new ListAuthenticatorMethods200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthenticatorMethods200ResponseInnerWithDefaults

`func NewListAuthenticatorMethods200ResponseInnerWithDefaults() *ListAuthenticatorMethods200ResponseInner`

NewListAuthenticatorMethods200ResponseInnerWithDefaults instantiates a new ListAuthenticatorMethods200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListAuthenticatorMethods200ResponseInner) GetStatus() LifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetStatusOk() (*LifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAuthenticatorMethods200ResponseInner) SetStatus(v LifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListAuthenticatorMethods200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ListAuthenticatorMethods200ResponseInner) GetType() AuthenticatorMethodType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetTypeOk() (*AuthenticatorMethodType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListAuthenticatorMethods200ResponseInner) SetType(v AuthenticatorMethodType)`

SetType sets Type field to given value.

### HasType

`func (o *ListAuthenticatorMethods200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ListAuthenticatorMethods200ResponseInner) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListAuthenticatorMethods200ResponseInner) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListAuthenticatorMethods200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSettings

`func (o *ListAuthenticatorMethods200ResponseInner) GetSettings() AuthenticatorMethodWebAuthnAllOfSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetSettingsOk() (*AuthenticatorMethodWebAuthnAllOfSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ListAuthenticatorMethods200ResponseInner) SetSettings(v AuthenticatorMethodWebAuthnAllOfSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ListAuthenticatorMethods200ResponseInner) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetVerifiableProperties

`func (o *ListAuthenticatorMethods200ResponseInner) GetVerifiableProperties() []AuthenticatorMethodProperty`

GetVerifiableProperties returns the VerifiableProperties field if non-nil, zero value otherwise.

### GetVerifiablePropertiesOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetVerifiablePropertiesOk() (*[]AuthenticatorMethodProperty, bool)`

GetVerifiablePropertiesOk returns a tuple with the VerifiableProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiableProperties

`func (o *ListAuthenticatorMethods200ResponseInner) SetVerifiableProperties(v []AuthenticatorMethodProperty)`

SetVerifiableProperties sets VerifiableProperties field to given value.

### HasVerifiableProperties

`func (o *ListAuthenticatorMethods200ResponseInner) HasVerifiableProperties() bool`

HasVerifiableProperties returns a boolean if a field has been set.

### GetAcceptableAdjacentIntervals

`func (o *ListAuthenticatorMethods200ResponseInner) GetAcceptableAdjacentIntervals() int32`

GetAcceptableAdjacentIntervals returns the AcceptableAdjacentIntervals field if non-nil, zero value otherwise.

### GetAcceptableAdjacentIntervalsOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetAcceptableAdjacentIntervalsOk() (*int32, bool)`

GetAcceptableAdjacentIntervalsOk returns a tuple with the AcceptableAdjacentIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableAdjacentIntervals

`func (o *ListAuthenticatorMethods200ResponseInner) SetAcceptableAdjacentIntervals(v int32)`

SetAcceptableAdjacentIntervals sets AcceptableAdjacentIntervals field to given value.

### HasAcceptableAdjacentIntervals

`func (o *ListAuthenticatorMethods200ResponseInner) HasAcceptableAdjacentIntervals() bool`

HasAcceptableAdjacentIntervals returns a boolean if a field has been set.

### GetAlgorithm

`func (o *ListAuthenticatorMethods200ResponseInner) GetAlgorithm() OtpTotpAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetAlgorithmOk() (*OtpTotpAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ListAuthenticatorMethods200ResponseInner) SetAlgorithm(v OtpTotpAlgorithm)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *ListAuthenticatorMethods200ResponseInner) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetEncoding

`func (o *ListAuthenticatorMethods200ResponseInner) GetEncoding() OtpTotpEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetEncodingOk() (*OtpTotpEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *ListAuthenticatorMethods200ResponseInner) SetEncoding(v OtpTotpEncoding)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *ListAuthenticatorMethods200ResponseInner) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetFactorProfileId

`func (o *ListAuthenticatorMethods200ResponseInner) GetFactorProfileId() string`

GetFactorProfileId returns the FactorProfileId field if non-nil, zero value otherwise.

### GetFactorProfileIdOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetFactorProfileIdOk() (*string, bool)`

GetFactorProfileIdOk returns a tuple with the FactorProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorProfileId

`func (o *ListAuthenticatorMethods200ResponseInner) SetFactorProfileId(v string)`

SetFactorProfileId sets FactorProfileId field to given value.

### HasFactorProfileId

`func (o *ListAuthenticatorMethods200ResponseInner) HasFactorProfileId() bool`

HasFactorProfileId returns a boolean if a field has been set.

### GetPassCodeLength

`func (o *ListAuthenticatorMethods200ResponseInner) GetPassCodeLength() int32`

GetPassCodeLength returns the PassCodeLength field if non-nil, zero value otherwise.

### GetPassCodeLengthOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetPassCodeLengthOk() (*int32, bool)`

GetPassCodeLengthOk returns a tuple with the PassCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCodeLength

`func (o *ListAuthenticatorMethods200ResponseInner) SetPassCodeLength(v int32)`

SetPassCodeLength sets PassCodeLength field to given value.

### HasPassCodeLength

`func (o *ListAuthenticatorMethods200ResponseInner) HasPassCodeLength() bool`

HasPassCodeLength returns a boolean if a field has been set.

### GetProtocol

`func (o *ListAuthenticatorMethods200ResponseInner) GetProtocol() OtpProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetProtocolOk() (*OtpProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ListAuthenticatorMethods200ResponseInner) SetProtocol(v OtpProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ListAuthenticatorMethods200ResponseInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTimeIntervalInSeconds

`func (o *ListAuthenticatorMethods200ResponseInner) GetTimeIntervalInSeconds() int32`

GetTimeIntervalInSeconds returns the TimeIntervalInSeconds field if non-nil, zero value otherwise.

### GetTimeIntervalInSecondsOk

`func (o *ListAuthenticatorMethods200ResponseInner) GetTimeIntervalInSecondsOk() (*int32, bool)`

GetTimeIntervalInSecondsOk returns a tuple with the TimeIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalInSeconds

`func (o *ListAuthenticatorMethods200ResponseInner) SetTimeIntervalInSeconds(v int32)`

SetTimeIntervalInSeconds sets TimeIntervalInSeconds field to given value.

### HasTimeIntervalInSeconds

`func (o *ListAuthenticatorMethods200ResponseInner) HasTimeIntervalInSeconds() bool`

HasTimeIntervalInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


