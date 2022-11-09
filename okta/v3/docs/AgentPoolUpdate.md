# AgentPoolUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]Agent**](Agent.md) |  | [optional] 
**AgentType** | Pointer to [**AgentType**](AgentType.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**NotifyAdmin** | Pointer to **bool** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**AutoUpdateSchedule**](AutoUpdateSchedule.md) |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**AgentUpdateJobStatus**](AgentUpdateJobStatus.md) |  | [optional] 
**TargetVersion** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 

## Methods

### NewAgentPoolUpdate

`func NewAgentPoolUpdate() *AgentPoolUpdate`

NewAgentPoolUpdate instantiates a new AgentPoolUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPoolUpdateWithDefaults

`func NewAgentPoolUpdateWithDefaults() *AgentPoolUpdate`

NewAgentPoolUpdateWithDefaults instantiates a new AgentPoolUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *AgentPoolUpdate) GetAgents() []Agent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *AgentPoolUpdate) GetAgentsOk() (*[]Agent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *AgentPoolUpdate) SetAgents(v []Agent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *AgentPoolUpdate) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetAgentType

`func (o *AgentPoolUpdate) GetAgentType() AgentType`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *AgentPoolUpdate) GetAgentTypeOk() (*AgentType, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *AgentPoolUpdate) SetAgentType(v AgentType)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *AgentPoolUpdate) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetEnabled

`func (o *AgentPoolUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AgentPoolUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AgentPoolUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AgentPoolUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *AgentPoolUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentPoolUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentPoolUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentPoolUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AgentPoolUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentPoolUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentPoolUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentPoolUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifyAdmin

`func (o *AgentPoolUpdate) GetNotifyAdmin() bool`

GetNotifyAdmin returns the NotifyAdmin field if non-nil, zero value otherwise.

### GetNotifyAdminOk

`func (o *AgentPoolUpdate) GetNotifyAdminOk() (*bool, bool)`

GetNotifyAdminOk returns a tuple with the NotifyAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyAdmin

`func (o *AgentPoolUpdate) SetNotifyAdmin(v bool)`

SetNotifyAdmin sets NotifyAdmin field to given value.

### HasNotifyAdmin

`func (o *AgentPoolUpdate) HasNotifyAdmin() bool`

HasNotifyAdmin returns a boolean if a field has been set.

### GetReason

`func (o *AgentPoolUpdate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AgentPoolUpdate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AgentPoolUpdate) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AgentPoolUpdate) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSchedule

`func (o *AgentPoolUpdate) GetSchedule() AutoUpdateSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AgentPoolUpdate) GetScheduleOk() (*AutoUpdateSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AgentPoolUpdate) SetSchedule(v AutoUpdateSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AgentPoolUpdate) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSortOrder

`func (o *AgentPoolUpdate) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AgentPoolUpdate) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AgentPoolUpdate) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AgentPoolUpdate) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetStatus

`func (o *AgentPoolUpdate) GetStatus() AgentUpdateJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentPoolUpdate) GetStatusOk() (*AgentUpdateJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentPoolUpdate) SetStatus(v AgentUpdateJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentPoolUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetVersion

`func (o *AgentPoolUpdate) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *AgentPoolUpdate) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *AgentPoolUpdate) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *AgentPoolUpdate) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetLinks

`func (o *AgentPoolUpdate) GetLinks() HrefObject`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentPoolUpdate) GetLinksOk() (*HrefObject, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentPoolUpdate) SetLinks(v HrefObject)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentPoolUpdate) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


