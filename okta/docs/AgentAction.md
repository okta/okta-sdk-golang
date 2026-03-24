# AgentAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the Active Directory or LDAP group to update | 
**Parameters** | [**Parameters**](Parameters.md) |  | 

## Methods

### NewAgentAction

`func NewAgentAction(id string, parameters Parameters, ) *AgentAction`

NewAgentAction instantiates a new AgentAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentActionWithDefaults

`func NewAgentActionWithDefaults() *AgentAction`

NewAgentActionWithDefaults instantiates a new AgentAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentAction) SetId(v string)`

SetId sets Id field to given value.


### GetParameters

`func (o *AgentAction) GetParameters() Parameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AgentAction) GetParametersOk() (*Parameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AgentAction) SetParameters(v Parameters)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


