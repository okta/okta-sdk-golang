# ResourceSetResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the role was created | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the resource set | [optional] 
**Id** | Pointer to **string** | Unique key for the role | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the role was last updated | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewResourceSetResource

`func NewResourceSetResource() *ResourceSetResource`

NewResourceSetResource instantiates a new ResourceSetResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetResourceWithDefaults

`func NewResourceSetResourceWithDefaults() *ResourceSetResource`

NewResourceSetResourceWithDefaults instantiates a new ResourceSetResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ResourceSetResource) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResourceSetResource) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResourceSetResource) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResourceSetResource) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *ResourceSetResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceSetResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceSetResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceSetResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ResourceSetResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceSetResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceSetResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceSetResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ResourceSetResource) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ResourceSetResource) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ResourceSetResource) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ResourceSetResource) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSetResource) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSetResource) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSetResource) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSetResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


