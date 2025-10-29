# AIAgentOperationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AIAgentOperationResponse**](AIAgentOperationResponse.md) |  | [optional] 
**Links** | Pointer to [**AIAgentOperationListResponseLinks**](AIAgentOperationListResponseLinks.md) |  | [optional] 

## Methods

### NewAIAgentOperationListResponse

`func NewAIAgentOperationListResponse() *AIAgentOperationListResponse`

NewAIAgentOperationListResponse instantiates a new AIAgentOperationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIAgentOperationListResponseWithDefaults

`func NewAIAgentOperationListResponseWithDefaults() *AIAgentOperationListResponse`

NewAIAgentOperationListResponseWithDefaults instantiates a new AIAgentOperationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AIAgentOperationListResponse) GetData() []AIAgentOperationResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AIAgentOperationListResponse) GetDataOk() (*[]AIAgentOperationResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AIAgentOperationListResponse) SetData(v []AIAgentOperationResponse)`

SetData sets Data field to given value.

### HasData

`func (o *AIAgentOperationListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *AIAgentOperationListResponse) GetLinks() AIAgentOperationListResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AIAgentOperationListResponse) GetLinksOk() (*AIAgentOperationListResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AIAgentOperationListResponse) SetLinks(v AIAgentOperationListResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AIAgentOperationListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


