# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsLatestGAedVersion** | Pointer to **bool** |  | [optional] 
**LastConnection** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OperationalStatus** | Pointer to [**OperationalStatus**](OperationalStatus.md) |  | [optional] 
**PoolId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AgentType**](AgentType.md) |  | [optional] 
**UpdateMessage** | Pointer to **string** |  | [optional] 
**UpdateStatus** | Pointer to [**AgentUpdateInstanceStatus**](AgentUpdateInstanceStatus.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**HrefObject**](HrefObject.md) |  | [optional] 

## Methods

### NewAgent

`func NewAgent() *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Agent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Agent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsHidden

`func (o *Agent) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *Agent) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *Agent) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *Agent) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsLatestGAedVersion

`func (o *Agent) GetIsLatestGAedVersion() bool`

GetIsLatestGAedVersion returns the IsLatestGAedVersion field if non-nil, zero value otherwise.

### GetIsLatestGAedVersionOk

`func (o *Agent) GetIsLatestGAedVersionOk() (*bool, bool)`

GetIsLatestGAedVersionOk returns a tuple with the IsLatestGAedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatestGAedVersion

`func (o *Agent) SetIsLatestGAedVersion(v bool)`

SetIsLatestGAedVersion sets IsLatestGAedVersion field to given value.

### HasIsLatestGAedVersion

`func (o *Agent) HasIsLatestGAedVersion() bool`

HasIsLatestGAedVersion returns a boolean if a field has been set.

### GetLastConnection

`func (o *Agent) GetLastConnection() time.Time`

GetLastConnection returns the LastConnection field if non-nil, zero value otherwise.

### GetLastConnectionOk

`func (o *Agent) GetLastConnectionOk() (*time.Time, bool)`

GetLastConnectionOk returns a tuple with the LastConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnection

`func (o *Agent) SetLastConnection(v time.Time)`

SetLastConnection sets LastConnection field to given value.

### HasLastConnection

`func (o *Agent) HasLastConnection() bool`

HasLastConnection returns a boolean if a field has been set.

### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Agent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *Agent) GetOperationalStatus() OperationalStatus`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *Agent) GetOperationalStatusOk() (*OperationalStatus, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *Agent) SetOperationalStatus(v OperationalStatus)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *Agent) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetPoolId

`func (o *Agent) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *Agent) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *Agent) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *Agent) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetType

`func (o *Agent) GetType() AgentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Agent) GetTypeOk() (*AgentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Agent) SetType(v AgentType)`

SetType sets Type field to given value.

### HasType

`func (o *Agent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateMessage

`func (o *Agent) GetUpdateMessage() string`

GetUpdateMessage returns the UpdateMessage field if non-nil, zero value otherwise.

### GetUpdateMessageOk

`func (o *Agent) GetUpdateMessageOk() (*string, bool)`

GetUpdateMessageOk returns a tuple with the UpdateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMessage

`func (o *Agent) SetUpdateMessage(v string)`

SetUpdateMessage sets UpdateMessage field to given value.

### HasUpdateMessage

`func (o *Agent) HasUpdateMessage() bool`

HasUpdateMessage returns a boolean if a field has been set.

### GetUpdateStatus

`func (o *Agent) GetUpdateStatus() AgentUpdateInstanceStatus`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *Agent) GetUpdateStatusOk() (*AgentUpdateInstanceStatus, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *Agent) SetUpdateStatus(v AgentUpdateInstanceStatus)`

SetUpdateStatus sets UpdateStatus field to given value.

### HasUpdateStatus

`func (o *Agent) HasUpdateStatus() bool`

HasUpdateStatus returns a boolean if a field has been set.

### GetVersion

`func (o *Agent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Agent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Agent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Agent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLinks

`func (o *Agent) GetLinks() HrefObject`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Agent) GetLinksOk() (*HrefObject, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Agent) SetLinks(v HrefObject)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Agent) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


