# AIAgentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the AI agent resource | 
**Status** | **string** | The status of the AI agent resource | 
**Type** | **string** | The type of resource | 
**Links** | [**LinksSelf**](LinksSelf.md) |  | 

## Methods

### NewAIAgentResource

`func NewAIAgentResource(id string, status string, type_ string, links LinksSelf, ) *AIAgentResource`

NewAIAgentResource instantiates a new AIAgentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAgentResourceWithDefaults

`func NewAIAgentResourceWithDefaults() *AIAgentResource`

NewAIAgentResourceWithDefaults instantiates a new AIAgentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AIAgentResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIAgentResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIAgentResource) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *AIAgentResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AIAgentResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AIAgentResource) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *AIAgentResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AIAgentResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AIAgentResource) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *AIAgentResource) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AIAgentResource) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AIAgentResource) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


