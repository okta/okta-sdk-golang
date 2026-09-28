# PreRegistrationErrorCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorSummary** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** | An Okta error code specific to this entry&#39;s failure | [optional] 
**Index** | Pointer to **int32** | Zero-based index of the failing entry in the request array. Useful for disambiguation when the same &#x60;serialNumber&#x60; appears more than once. | [optional] 
**Platform** | Pointer to **string** | Device platform for pre-registration | [optional] 
**SerialNumber** | Pointer to **string** | Serial number of the failing entry | [optional] 

## Methods

### NewPreRegistrationErrorCause

`func NewPreRegistrationErrorCause() *PreRegistrationErrorCause`

NewPreRegistrationErrorCause instantiates a new PreRegistrationErrorCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreRegistrationErrorCauseWithDefaults

`func NewPreRegistrationErrorCauseWithDefaults() *PreRegistrationErrorCause`

NewPreRegistrationErrorCauseWithDefaults instantiates a new PreRegistrationErrorCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorSummary

`func (o *PreRegistrationErrorCause) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *PreRegistrationErrorCause) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *PreRegistrationErrorCause) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *PreRegistrationErrorCause) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.

### GetErrorCode

`func (o *PreRegistrationErrorCause) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PreRegistrationErrorCause) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PreRegistrationErrorCause) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *PreRegistrationErrorCause) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetIndex

`func (o *PreRegistrationErrorCause) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PreRegistrationErrorCause) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PreRegistrationErrorCause) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *PreRegistrationErrorCause) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetPlatform

`func (o *PreRegistrationErrorCause) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PreRegistrationErrorCause) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PreRegistrationErrorCause) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PreRegistrationErrorCause) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PreRegistrationErrorCause) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PreRegistrationErrorCause) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PreRegistrationErrorCause) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PreRegistrationErrorCause) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


