# ResourceSetResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**ResourceConditions**](ResourceConditions.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the resource set resource object was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique ID of the resource set resource object | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when this object was last updated | [optional] [readonly] 
**Orn** | Pointer to **string** | The Okta Resource Name (ORN) of the resource | [optional] 
**Links** | Pointer to [**ResourceSetResourceLinks**](ResourceSetResourceLinks.md) |  | [optional] 

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

### GetConditions

`func (o *ResourceSetResource) GetConditions() ResourceConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ResourceSetResource) GetConditionsOk() (*ResourceConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ResourceSetResource) SetConditions(v ResourceConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ResourceSetResource) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

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

### GetOrn

`func (o *ResourceSetResource) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ResourceSetResource) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ResourceSetResource) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *ResourceSetResource) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSetResource) GetLinks() ResourceSetResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSetResource) GetLinksOk() (*ResourceSetResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSetResource) SetLinks(v ResourceSetResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSetResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


