# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **NullableString** | Human-readable failure reason. &#x60;null&#x60; unless &#x60;status&#x60; is &#x60;FAILED&#x60;. | [optional] 
**Status** | **string** | Status of the execution. &#x60;FAILED&#x60; is an execution-level failure (for example, a worker crash), not the outcome of an individual task. | 

## Methods

### NewExecution

`func NewExecution(status string, ) *Execution`

NewExecution instantiates a new Execution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionWithDefaults

`func NewExecutionWithDefaults() *Execution`

NewExecutionWithDefaults instantiates a new Execution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Execution) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Execution) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Execution) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Execution) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *Execution) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Execution) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetStatus

`func (o *Execution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Execution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Execution) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


