# ResourceSelectorsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceSelectors** | Pointer to [**[]ResourceSelectorResponseWithoutSelfLinkSchema**](ResourceSelectorResponseWithoutSelfLinkSchema.md) |  | [optional] 
**Links** | Pointer to [**LinksNext**](LinksNext.md) |  | [optional] 

## Methods

### NewResourceSelectorsSchema

`func NewResourceSelectorsSchema() *ResourceSelectorsSchema`

NewResourceSelectorsSchema instantiates a new ResourceSelectorsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSelectorsSchemaWithDefaults

`func NewResourceSelectorsSchemaWithDefaults() *ResourceSelectorsSchema`

NewResourceSelectorsSchemaWithDefaults instantiates a new ResourceSelectorsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceSelectors

`func (o *ResourceSelectorsSchema) GetResourceSelectors() []ResourceSelectorResponseWithoutSelfLinkSchema`

GetResourceSelectors returns the ResourceSelectors field if non-nil, zero value otherwise.

### GetResourceSelectorsOk

`func (o *ResourceSelectorsSchema) GetResourceSelectorsOk() (*[]ResourceSelectorResponseWithoutSelfLinkSchema, bool)`

GetResourceSelectorsOk returns a tuple with the ResourceSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSelectors

`func (o *ResourceSelectorsSchema) SetResourceSelectors(v []ResourceSelectorResponseWithoutSelfLinkSchema)`

SetResourceSelectors sets ResourceSelectors field to given value.

### HasResourceSelectors

`func (o *ResourceSelectorsSchema) HasResourceSelectors() bool`

HasResourceSelectors returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSelectorsSchema) GetLinks() LinksNext`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSelectorsSchema) GetLinksOk() (*LinksNext, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSelectorsSchema) SetLinks(v LinksNext)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSelectorsSchema) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


