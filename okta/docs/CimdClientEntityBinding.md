# CimdClientEntityBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundEntityId** | **string** | Unique identifier of the entity that&#39;s bound to the CIMD client entity | 
**BoundEntityType** | **string** | The type of resource that&#39;s bound to the CIMD client entity | 
**Created** | Pointer to **time.Time** | Timestamp when the binding was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier for the CIMD client entity binding | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the binding was last updated | [optional] [readonly] 
**Links** | Pointer to [**CimdClientEntityBindingLinks**](CimdClientEntityBindingLinks.md) |  | [optional] 

## Methods

### NewCimdClientEntityBinding

`func NewCimdClientEntityBinding(boundEntityId string, boundEntityType string, ) *CimdClientEntityBinding`

NewCimdClientEntityBinding instantiates a new CimdClientEntityBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCimdClientEntityBindingWithDefaults

`func NewCimdClientEntityBindingWithDefaults() *CimdClientEntityBinding`

NewCimdClientEntityBindingWithDefaults instantiates a new CimdClientEntityBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundEntityId

`func (o *CimdClientEntityBinding) GetBoundEntityId() string`

GetBoundEntityId returns the BoundEntityId field if non-nil, zero value otherwise.

### GetBoundEntityIdOk

`func (o *CimdClientEntityBinding) GetBoundEntityIdOk() (*string, bool)`

GetBoundEntityIdOk returns a tuple with the BoundEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEntityId

`func (o *CimdClientEntityBinding) SetBoundEntityId(v string)`

SetBoundEntityId sets BoundEntityId field to given value.


### GetBoundEntityType

`func (o *CimdClientEntityBinding) GetBoundEntityType() string`

GetBoundEntityType returns the BoundEntityType field if non-nil, zero value otherwise.

### GetBoundEntityTypeOk

`func (o *CimdClientEntityBinding) GetBoundEntityTypeOk() (*string, bool)`

GetBoundEntityTypeOk returns a tuple with the BoundEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundEntityType

`func (o *CimdClientEntityBinding) SetBoundEntityType(v string)`

SetBoundEntityType sets BoundEntityType field to given value.


### GetCreated

`func (o *CimdClientEntityBinding) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CimdClientEntityBinding) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CimdClientEntityBinding) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CimdClientEntityBinding) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *CimdClientEntityBinding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CimdClientEntityBinding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CimdClientEntityBinding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CimdClientEntityBinding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CimdClientEntityBinding) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CimdClientEntityBinding) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CimdClientEntityBinding) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CimdClientEntityBinding) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *CimdClientEntityBinding) GetLinks() CimdClientEntityBindingLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CimdClientEntityBinding) GetLinksOk() (*CimdClientEntityBindingLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CimdClientEntityBinding) SetLinks(v CimdClientEntityBindingLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CimdClientEntityBinding) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


