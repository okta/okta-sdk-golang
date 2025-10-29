# AgentJsonWebKeyResponseBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Timestamp of when the AI agent JSON Web Key was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of the AI agent JSON Web Key | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp of when the AI agent JSON Web Key was last updated | [optional] [readonly] 
**Links** | Pointer to [**AgentSecretLinks**](AgentSecretLinks.md) |  | [optional] 

## Methods

### NewAgentJsonWebKeyResponseBase

`func NewAgentJsonWebKeyResponseBase() *AgentJsonWebKeyResponseBase`

NewAgentJsonWebKeyResponseBase instantiates a new AgentJsonWebKeyResponseBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentJsonWebKeyResponseBaseWithDefaults

`func NewAgentJsonWebKeyResponseBaseWithDefaults() *AgentJsonWebKeyResponseBase`

NewAgentJsonWebKeyResponseBaseWithDefaults instantiates a new AgentJsonWebKeyResponseBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AgentJsonWebKeyResponseBase) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AgentJsonWebKeyResponseBase) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AgentJsonWebKeyResponseBase) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AgentJsonWebKeyResponseBase) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AgentJsonWebKeyResponseBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentJsonWebKeyResponseBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentJsonWebKeyResponseBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentJsonWebKeyResponseBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AgentJsonWebKeyResponseBase) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AgentJsonWebKeyResponseBase) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AgentJsonWebKeyResponseBase) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AgentJsonWebKeyResponseBase) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *AgentJsonWebKeyResponseBase) GetLinks() AgentSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AgentJsonWebKeyResponseBase) GetLinksOk() (*AgentSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AgentJsonWebKeyResponseBase) SetLinks(v AgentSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AgentJsonWebKeyResponseBase) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


