# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCauses** | Pointer to [**[]ErrorErrorCausesInner**](ErrorErrorCausesInner.md) |  | [optional] 
**ErrorCode** | Pointer to **string** | An Okta code for this type of error | [optional] 
**ErrorId** | Pointer to **string** | A unique identifier for this error. This can be used by Okta Support to help with troubleshooting. | [optional] 
**ErrorLink** | Pointer to **string** | An Okta code for this type of error | [optional] 
**ErrorSummary** | Pointer to **string** | A short description of what caused this error. Sometimes this contains dynamically-generated information about your specific error. | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCauses

`func (o *Error) GetErrorCauses() []ErrorErrorCausesInner`

GetErrorCauses returns the ErrorCauses field if non-nil, zero value otherwise.

### GetErrorCausesOk

`func (o *Error) GetErrorCausesOk() (*[]ErrorErrorCausesInner, bool)`

GetErrorCausesOk returns a tuple with the ErrorCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCauses

`func (o *Error) SetErrorCauses(v []ErrorErrorCausesInner)`

SetErrorCauses sets ErrorCauses field to given value.

### HasErrorCauses

`func (o *Error) HasErrorCauses() bool`

HasErrorCauses returns a boolean if a field has been set.

### GetErrorCode

`func (o *Error) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Error) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Error) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Error) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorId

`func (o *Error) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *Error) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *Error) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *Error) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetErrorLink

`func (o *Error) GetErrorLink() string`

GetErrorLink returns the ErrorLink field if non-nil, zero value otherwise.

### GetErrorLinkOk

`func (o *Error) GetErrorLinkOk() (*string, bool)`

GetErrorLinkOk returns a tuple with the ErrorLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLink

`func (o *Error) SetErrorLink(v string)`

SetErrorLink sets ErrorLink field to given value.

### HasErrorLink

`func (o *Error) HasErrorLink() bool`

HasErrorLink returns a boolean if a field has been set.

### GetErrorSummary

`func (o *Error) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *Error) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *Error) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *Error) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


