# AIAgentOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **time.Time** | Timestamp of when the AI agent operation completed | [optional] 
**Created** | **time.Time** | Timestamp of when the AI agent operation was created | 
**ErrorDetails** | Pointer to [**ErrorDetails**](ErrorDetails.md) |  | [optional] 
**Id** | **string** | ID of the AI agent operation | 
**Resource** | Pointer to [**AIAgentResource**](AIAgentResource.md) |  | [optional] 
**Started** | Pointer to **time.Time** | Timestamp of when the AI agent operation started | [optional] 
**Status** | **string** | The status of the AI agent operation | 
**Type** | **string** | The AI agent operation type | 

## Methods

### NewAIAgentOperationResponse

`func NewAIAgentOperationResponse(created time.Time, id string, status string, type_ string, ) *AIAgentOperationResponse`

NewAIAgentOperationResponse instantiates a new AIAgentOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAgentOperationResponseWithDefaults

`func NewAIAgentOperationResponseWithDefaults() *AIAgentOperationResponse`

NewAIAgentOperationResponseWithDefaults instantiates a new AIAgentOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *AIAgentOperationResponse) GetCompleted() time.Time`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *AIAgentOperationResponse) GetCompletedOk() (*time.Time, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *AIAgentOperationResponse) SetCompleted(v time.Time)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *AIAgentOperationResponse) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetCreated

`func (o *AIAgentOperationResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AIAgentOperationResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AIAgentOperationResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetErrorDetails

`func (o *AIAgentOperationResponse) GetErrorDetails() ErrorDetails`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *AIAgentOperationResponse) GetErrorDetailsOk() (*ErrorDetails, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *AIAgentOperationResponse) SetErrorDetails(v ErrorDetails)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *AIAgentOperationResponse) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetId

`func (o *AIAgentOperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIAgentOperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIAgentOperationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetResource

`func (o *AIAgentOperationResponse) GetResource() AIAgentResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AIAgentOperationResponse) GetResourceOk() (*AIAgentResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AIAgentOperationResponse) SetResource(v AIAgentResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AIAgentOperationResponse) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetStarted

`func (o *AIAgentOperationResponse) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *AIAgentOperationResponse) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *AIAgentOperationResponse) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *AIAgentOperationResponse) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *AIAgentOperationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AIAgentOperationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AIAgentOperationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *AIAgentOperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AIAgentOperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AIAgentOperationResponse) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


