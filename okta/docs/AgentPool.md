# AgentPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]Agent**](Agent.md) |  | [optional] 
**DisruptedAgents** | Pointer to **int32** | Number of agents in the pool that are in a disrupted state | [optional] 
**Id** | Pointer to **string** | Agent pool ID | [optional] [readonly] 
**InactiveAgents** | Pointer to **int32** | Number of agents in the pool that are in an inactive state | [optional] 
**Name** | Pointer to **string** | Agent pool name | [optional] 
**OperationalStatus** | Pointer to **string** | Operational status of a given agent | [optional] 
**Type** | Pointer to **string** | Agent types that are being monitored | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

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

### GetDisruptedAgents

`func (o *AgentPool) GetDisruptedAgents() int32`

GetDisruptedAgents returns the DisruptedAgents field if non-nil, zero value otherwise.

### GetDisruptedAgentsOk

`func (o *AgentPool) GetDisruptedAgentsOk() (*int32, bool)`

GetDisruptedAgentsOk returns a tuple with the DisruptedAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptedAgents

`func (o *AgentPool) SetDisruptedAgents(v int32)`

SetDisruptedAgents sets DisruptedAgents field to given value.

### HasDisruptedAgents

`func (o *AgentPool) HasDisruptedAgents() bool`

HasDisruptedAgents returns a boolean if a field has been set.

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

### GetInactiveAgents

`func (o *AgentPool) GetInactiveAgents() int32`

GetInactiveAgents returns the InactiveAgents field if non-nil, zero value otherwise.

### GetInactiveAgentsOk

`func (o *AgentPool) GetInactiveAgentsOk() (*int32, bool)`

GetInactiveAgentsOk returns a tuple with the InactiveAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveAgents

`func (o *AgentPool) SetInactiveAgents(v int32)`

SetInactiveAgents sets InactiveAgents field to given value.

### HasInactiveAgents

`func (o *AgentPool) HasInactiveAgents() bool`

HasInactiveAgents returns a boolean if a field has been set.

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

`func (o *AgentPool) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *AgentPool) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *AgentPool) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *AgentPool) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetType

`func (o *AgentPool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgentPool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgentPool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgentPool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *AgentPool) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentPool) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentPool) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentPool) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


