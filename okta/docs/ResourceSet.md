# ResourceSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the role was created | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the resource set | [optional] 
**Id** | Pointer to **string** | Unique key for the role | [optional] [readonly] 
**Label** | Pointer to **string** | Unique label for the resource set | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the role was last updated | [optional] [readonly] 
**Links** | Pointer to [**ResourceSetLinks**](ResourceSetLinks.md) |  | [optional] 

## Methods

### NewResourceSet

`func NewResourceSet() *ResourceSet`

NewResourceSet instantiates a new ResourceSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetWithDefaults

`func NewResourceSetWithDefaults() *ResourceSet`

NewResourceSetWithDefaults instantiates a new ResourceSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ResourceSet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResourceSet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResourceSet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResourceSet) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *ResourceSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ResourceSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ResourceSet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ResourceSet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ResourceSet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ResourceSet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ResourceSet) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ResourceSet) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ResourceSet) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ResourceSet) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSet) GetLinks() ResourceSetLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSet) GetLinksOk() (*ResourceSetLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSet) SetLinks(v ResourceSetLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


