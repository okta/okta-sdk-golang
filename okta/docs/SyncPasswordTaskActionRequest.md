# SyncPasswordTaskActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to perform on a sync password task. Only the actions listed in a task&#39;s &#x60;eligibleActions&#x60; can be performed on it. If you specify &#x60;taskIds&#x60; and a task doesn&#39;t allow the action, the request is rejected with a &#x60;400&#x60; error. If you specify &#x60;appInstanceIds&#x60;, only the tasks that allow the action are acted on. Tasks that no longer allow the action when the request is processed are skipped. | 
**AppInstanceIds** | Pointer to **[]string** | Acts on every currently eligible task across these app instances (maximum 10 by default) | [optional] 
**TaskIds** | Pointer to **[]string** | Explicit IDs of the tasks to act on (maximum 50 by default) | [optional] 

## Methods

### NewSyncPasswordTaskActionRequest

`func NewSyncPasswordTaskActionRequest(action string, ) *SyncPasswordTaskActionRequest`

NewSyncPasswordTaskActionRequest instantiates a new SyncPasswordTaskActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPasswordTaskActionRequestWithDefaults

`func NewSyncPasswordTaskActionRequestWithDefaults() *SyncPasswordTaskActionRequest`

NewSyncPasswordTaskActionRequestWithDefaults instantiates a new SyncPasswordTaskActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SyncPasswordTaskActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SyncPasswordTaskActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SyncPasswordTaskActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetAppInstanceIds

`func (o *SyncPasswordTaskActionRequest) GetAppInstanceIds() []string`

GetAppInstanceIds returns the AppInstanceIds field if non-nil, zero value otherwise.

### GetAppInstanceIdsOk

`func (o *SyncPasswordTaskActionRequest) GetAppInstanceIdsOk() (*[]string, bool)`

GetAppInstanceIdsOk returns a tuple with the AppInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceIds

`func (o *SyncPasswordTaskActionRequest) SetAppInstanceIds(v []string)`

SetAppInstanceIds sets AppInstanceIds field to given value.

### HasAppInstanceIds

`func (o *SyncPasswordTaskActionRequest) HasAppInstanceIds() bool`

HasAppInstanceIds returns a boolean if a field has been set.

### GetTaskIds

`func (o *SyncPasswordTaskActionRequest) GetTaskIds() []string`

GetTaskIds returns the TaskIds field if non-nil, zero value otherwise.

### GetTaskIdsOk

`func (o *SyncPasswordTaskActionRequest) GetTaskIdsOk() (*[]string, bool)`

GetTaskIdsOk returns a tuple with the TaskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIds

`func (o *SyncPasswordTaskActionRequest) SetTaskIds(v []string)`

SetTaskIds sets TaskIds field to given value.

### HasTaskIds

`func (o *SyncPasswordTaskActionRequest) HasTaskIds() bool`

HasTaskIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


