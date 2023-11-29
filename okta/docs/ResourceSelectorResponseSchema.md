# ResourceSelectorResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the Resource Selector | [optional] 
**Id** | Pointer to **string** | Unique key for the Resource Selector | [optional] 
**Name** | Pointer to **string** | Name of the Resource Selector | [optional] 
**Orn** | Pointer to **string** | An Okta resource name | [optional] 
**Links** | Pointer to [**ResourceSelectorResponseSchemaLinks**](ResourceSelectorResponseSchemaLinks.md) |  | [optional] 

## Methods

### NewResourceSelectorResponseSchema

`func NewResourceSelectorResponseSchema() *ResourceSelectorResponseSchema`

NewResourceSelectorResponseSchema instantiates a new ResourceSelectorResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSelectorResponseSchemaWithDefaults

`func NewResourceSelectorResponseSchemaWithDefaults() *ResourceSelectorResponseSchema`

NewResourceSelectorResponseSchemaWithDefaults instantiates a new ResourceSelectorResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ResourceSelectorResponseSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceSelectorResponseSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceSelectorResponseSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceSelectorResponseSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ResourceSelectorResponseSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceSelectorResponseSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceSelectorResponseSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceSelectorResponseSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceSelectorResponseSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceSelectorResponseSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceSelectorResponseSchema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceSelectorResponseSchema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrn

`func (o *ResourceSelectorResponseSchema) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ResourceSelectorResponseSchema) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ResourceSelectorResponseSchema) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *ResourceSelectorResponseSchema) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSelectorResponseSchema) GetLinks() ResourceSelectorResponseSchemaLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSelectorResponseSchema) GetLinksOk() (*ResourceSelectorResponseSchemaLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSelectorResponseSchema) SetLinks(v ResourceSelectorResponseSchemaLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSelectorResponseSchema) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


