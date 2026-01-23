# WorkflowActionProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The unique identifier of the action flow in the provider system | 
**Type** | **string** | Type of action provider | 
**Url** | **string** | The URL to the action flow | 

## Methods

### NewWorkflowActionProvider

`func NewWorkflowActionProvider(externalId string, type_ string, url string, ) *WorkflowActionProvider`

NewWorkflowActionProvider instantiates a new WorkflowActionProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionProviderWithDefaults

`func NewWorkflowActionProviderWithDefaults() *WorkflowActionProvider`

NewWorkflowActionProviderWithDefaults instantiates a new WorkflowActionProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *WorkflowActionProvider) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *WorkflowActionProvider) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *WorkflowActionProvider) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *WorkflowActionProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowActionProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowActionProvider) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *WorkflowActionProvider) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkflowActionProvider) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkflowActionProvider) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


