# ErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The error code | [optional] 
**Message** | Pointer to **string** | The error message | [optional] 

## Methods

### NewErrorDetails

`func NewErrorDetails() *ErrorDetails`

NewErrorDetails instantiates a new ErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailsWithDefaults

`func NewErrorDetailsWithDefaults() *ErrorDetails`

NewErrorDetailsWithDefaults instantiates a new ErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorDetails) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorDetails) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorDetails) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorDetails) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetails) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorDetails) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


