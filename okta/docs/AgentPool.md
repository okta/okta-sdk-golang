# AgentPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]Agent**](Agent.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**OperationalStatus** | Pointer to [**OperationalStatus**](OperationalStatus.md) |  | [optional] 
**Type** | Pointer to [**AgentType**](AgentType.md) |  | [optional] 

## Methods

### NewAgentPool

`func NewAgentPool() *AgentPool`

NewAgentPool instantiates a new AgentPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentPoolWithDefaults

`func NewAgentPoolWithDefaults() *AgentPool`

NewAgentPoolWithDefaults instantiates a new AgentPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *AgentPool) GetAgents() []Agent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *AgentPool) GetAgentsOk() (*[]Agent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *AgentPool) SetAgents(v []Agent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *AgentPool) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetId

`func (o *AgentPool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentPool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentPool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AgentPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *AgentPool) GetOperationalStatus() OperationalStatus`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *AgentPool) GetOperationalStatusOk() (*OperationalStatus, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *AgentPool) SetOperationalStatus(v OperationalStatus)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *AgentPool) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetType

`func (o *AgentPool) GetType() AgentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentPool) GetTypeOk() (*AgentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentPool) SetType(v AgentType)`

SetType sets Type field to given value.

### HasType

`func (o *AgentPool) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


