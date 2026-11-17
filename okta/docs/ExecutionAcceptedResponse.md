# ExecutionAcceptedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** | ID used to retrieve the execution status | 
**Status** | **string** | Status of the accepted execution | 

## Methods

### NewExecutionAcceptedResponse

`func NewExecutionAcceptedResponse(executionId string, status string, ) *ExecutionAcceptedResponse`

NewExecutionAcceptedResponse instantiates a new ExecutionAcceptedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionAcceptedResponseWithDefaults

`func NewExecutionAcceptedResponseWithDefaults() *ExecutionAcceptedResponse`

NewExecutionAcceptedResponseWithDefaults instantiates a new ExecutionAcceptedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *ExecutionAcceptedResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ExecutionAcceptedResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ExecutionAcceptedResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetStatus

`func (o *ExecutionAcceptedResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionAcceptedResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionAcceptedResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


