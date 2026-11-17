# AppAssignmentTaskActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to perform on an app assignment task. Only the actions listed in a task&#39;s &#x60;eligibleActions&#x60; can be performed on it. If you specify &#x60;taskIds&#x60; and a task doesn&#39;t allow the action, the request is rejected with a &#x60;400&#x60; error. If you specify &#x60;appInstanceIds&#x60;, only the tasks that allow the action are acted on. Tasks that no longer allow the action when the request is processed are skipped. | 
**AppInstanceIds** | Pointer to **[]string** | Acts on every currently eligible task across these app instances (maximum 10 by default) | [optional] 
**TaskIds** | Pointer to **[]string** | Explicit IDs of the tasks to act on (maximum 50 by default) | [optional] 

## Methods

### NewAppAssignmentTaskActionRequest

`func NewAppAssignmentTaskActionRequest(action string, ) *AppAssignmentTaskActionRequest`

NewAppAssignmentTaskActionRequest instantiates a new AppAssignmentTaskActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAssignmentTaskActionRequestWithDefaults

`func NewAppAssignmentTaskActionRequestWithDefaults() *AppAssignmentTaskActionRequest`

NewAppAssignmentTaskActionRequestWithDefaults instantiates a new AppAssignmentTaskActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AppAssignmentTaskActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AppAssignmentTaskActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AppAssignmentTaskActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetAppInstanceIds

`func (o *AppAssignmentTaskActionRequest) GetAppInstanceIds() []string`

GetAppInstanceIds returns the AppInstanceIds field if non-nil, zero value otherwise.

### GetAppInstanceIdsOk

`func (o *AppAssignmentTaskActionRequest) GetAppInstanceIdsOk() (*[]string, bool)`

GetAppInstanceIdsOk returns a tuple with the AppInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceIds

`func (o *AppAssignmentTaskActionRequest) SetAppInstanceIds(v []string)`

SetAppInstanceIds sets AppInstanceIds field to given value.

### HasAppInstanceIds

`func (o *AppAssignmentTaskActionRequest) HasAppInstanceIds() bool`

HasAppInstanceIds returns a boolean if a field has been set.

### GetTaskIds

`func (o *AppAssignmentTaskActionRequest) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *AppAssignmentTaskActionRequest) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *AppAssignmentTaskActionRequest) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *AppAssignmentTaskActionRequest) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


