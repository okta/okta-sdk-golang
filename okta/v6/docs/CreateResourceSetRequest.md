# CreateResourceSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the resource set | 
**Label** | **string** | Unique name for the resource set | 
**Resources** | **[]string** | The endpoint (URL) that references all resource objects included in the resource set. Resources are identified by either an Okta Resource Name (ORN) or by a REST URL format. See [Okta Resource Name](/openapi/okta-management/guides/roles/#okta-resource-name-orn). | 

## Methods

### NewCreateResourceSetRequest

`func NewCreateResourceSetRequest(description string, label string, resources []string, ) *CreateResourceSetRequest`

NewCreateResourceSetRequest instantiates a new CreateResourceSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceSetRequestWithDefaults

`func NewCreateResourceSetRequestWithDefaults() *CreateResourceSetRequest`

NewCreateResourceSetRequestWithDefaults instantiates a new CreateResourceSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateResourceSetRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateResourceSetRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateResourceSetRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLabel

`func (o *CreateResourceSetRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateResourceSetRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateResourceSetRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetResources

`func (o *CreateResourceSetRequest) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CreateResourceSetRequest) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CreateResourceSetRequest) SetResources(v []string)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


