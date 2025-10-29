# ValidationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Action identifier | 
**Provider** | [**WorkflowsValidationDetailProvider**](WorkflowsValidationDetailProvider.md) |  | 

## Methods

### NewValidationDetail

`func NewValidationDetail(id string, provider WorkflowsValidationDetailProvider, ) *ValidationDetail`

NewValidationDetail instantiates a new ValidationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationDetailWithDefaults

`func NewValidationDetailWithDefaults() *ValidationDetail`

NewValidationDetailWithDefaults instantiates a new ValidationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ValidationDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ValidationDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ValidationDetail) SetId(v string)`

SetId sets Id field to given value.


### GetProvider

`func (o *ValidationDetail) GetProvider() WorkflowsValidationDetailProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ValidationDetail) GetProviderOk() (*WorkflowsValidationDetailProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ValidationDetail) SetProvider(v WorkflowsValidationDetailProvider)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


