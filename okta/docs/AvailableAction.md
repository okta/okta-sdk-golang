# AvailableAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Action identifier | 
**Provider** | [**WorkflowAvailableActionProvider**](WorkflowAvailableActionProvider.md) |  | 

## Methods

### NewAvailableAction

`func NewAvailableAction(id string, provider WorkflowAvailableActionProvider, ) *AvailableAction`

NewAvailableAction instantiates a new AvailableAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableActionWithDefaults

`func NewAvailableActionWithDefaults() *AvailableAction`

NewAvailableActionWithDefaults instantiates a new AvailableAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AvailableAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvailableAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvailableAction) SetId(v string)`

SetId sets Id field to given value.


### GetProvider

`func (o *AvailableAction) GetProvider() WorkflowAvailableActionProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AvailableAction) GetProviderOk() (*WorkflowAvailableActionProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AvailableAction) SetProvider(v WorkflowAvailableActionProvider)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


