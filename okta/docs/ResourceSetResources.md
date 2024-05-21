# ResourceSetResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ResourceSetResource**](ResourceSetResource.md) |  | [optional] 
**Links** | Pointer to [**ResourceSetResourcesLinks**](ResourceSetResourcesLinks.md) |  | [optional] 

## Methods

### NewResourceSetResources

`func NewResourceSetResources() *ResourceSetResources`

NewResourceSetResources instantiates a new ResourceSetResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetResourcesWithDefaults

`func NewResourceSetResourcesWithDefaults() *ResourceSetResources`

NewResourceSetResourcesWithDefaults instantiates a new ResourceSetResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResources

`func (o *ResourceSetResources) GetResources() []ResourceSetResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourceSetResources) GetResourcesOk() (*[]ResourceSetResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourceSetResources) SetResources(v []ResourceSetResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ResourceSetResources) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSetResources) GetLinks() ResourceSetResourcesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSetResources) GetLinksOk() (*ResourceSetResourcesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSetResources) SetLinks(v ResourceSetResourcesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSetResources) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


