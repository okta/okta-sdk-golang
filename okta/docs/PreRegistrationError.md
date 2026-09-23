# PreRegistrationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCauses** | Pointer to [**[]PreRegistrationErrorCause**](PreRegistrationErrorCause.md) | Per-entry validation errors. Each cause identifies the failing device by &#x60;serialNumber&#x60;. | [optional] 
**ErrorCode** | Pointer to **string** | An Okta code for this type of error | [optional] 
**ErrorId** | Pointer to **string** | A unique identifier for this error. This can be used by Okta Support to help with troubleshooting. | [optional] 
**ErrorLink** | Pointer to **string** | An Okta code for this type of error | [optional] 
**ErrorSummary** | Pointer to **string** | A short description of what caused this error. Sometimes this contains dynamically-generated information about your specific error. | [optional] 

## Methods

### NewPreRegistrationError

`func NewPreRegistrationError() *PreRegistrationError`

NewPreRegistrationError instantiates a new PreRegistrationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreRegistrationErrorWithDefaults

`func NewPreRegistrationErrorWithDefaults() *PreRegistrationError`

NewPreRegistrationErrorWithDefaults instantiates a new PreRegistrationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCauses

`func (o *PreRegistrationError) GetErrorCauses() []PreRegistrationErrorCause`

GetErrorCauses returns the ErrorCauses field if non-nil, zero value otherwise.

### GetErrorCausesOk

`func (o *PreRegistrationError) GetErrorCausesOk() (*[]PreRegistrationErrorCause, bool)`

GetErrorCausesOk returns a tuple with the ErrorCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCauses

`func (o *PreRegistrationError) SetErrorCauses(v []PreRegistrationErrorCause)`

SetErrorCauses sets ErrorCauses field to given value.

### HasErrorCauses

`func (o *PreRegistrationError) HasErrorCauses() bool`

HasErrorCauses returns a boolean if a field has been set.

### GetErrorCode

`func (o *PreRegistrationError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PreRegistrationError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PreRegistrationError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *PreRegistrationError) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorId

`func (o *PreRegistrationError) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *PreRegistrationError) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *PreRegistrationError) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *PreRegistrationError) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetErrorLink

`func (o *PreRegistrationError) GetErrorLink() string`

GetErrorLink returns the ErrorLink field if non-nil, zero value otherwise.

### GetErrorLinkOk

`func (o *PreRegistrationError) GetErrorLinkOk() (*string, bool)`

GetErrorLinkOk returns a tuple with the ErrorLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLink

`func (o *PreRegistrationError) SetErrorLink(v string)`

SetErrorLink sets ErrorLink field to given value.

### HasErrorLink

`func (o *PreRegistrationError) HasErrorLink() bool`

HasErrorLink returns a boolean if a field has been set.

### GetErrorSummary

`func (o *PreRegistrationError) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *PreRegistrationError) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *PreRegistrationError) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *PreRegistrationError) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


