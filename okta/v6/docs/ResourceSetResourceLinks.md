# ResourceSetResourceLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObject**](HrefObject.md) | The REST API URL of the related resource | [optional] 
**Resource** | Pointer to [**HrefObject**](HrefObject.md) | Link to this resource set resource object (self) | [optional] 
**Groups** | Pointer to [**HrefObject**](HrefObject.md) | If applicable, the REST API URL of the related groups resource | [optional] 
**Users** | Pointer to [**HrefObject**](HrefObject.md) | If applicable, the REST API URL of the related users resource | [optional] 

## Methods

### NewResourceSetResourceLinks

`func NewResourceSetResourceLinks() *ResourceSetResourceLinks`

NewResourceSetResourceLinks instantiates a new ResourceSetResourceLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetResourceLinksWithDefaults

`func NewResourceSetResourceLinksWithDefaults() *ResourceSetResourceLinks`

NewResourceSetResourceLinksWithDefaults instantiates a new ResourceSetResourceLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ResourceSetResourceLinks) GetSelf() HrefObject`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ResourceSetResourceLinks) GetSelfOk() (*HrefObject, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ResourceSetResourceLinks) SetSelf(v HrefObject)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ResourceSetResourceLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetResource

`func (o *ResourceSetResourceLinks) GetResource() HrefObject`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceSetResourceLinks) GetResourceOk() (*HrefObject, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceSetResourceLinks) SetResource(v HrefObject)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceSetResourceLinks) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetGroups

`func (o *ResourceSetResourceLinks) GetGroups() HrefObject`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ResourceSetResourceLinks) GetGroupsOk() (*HrefObject, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ResourceSetResourceLinks) SetGroups(v HrefObject)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ResourceSetResourceLinks) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *ResourceSetResourceLinks) GetUsers() HrefObject`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ResourceSetResourceLinks) GetUsersOk() (*HrefObject, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ResourceSetResourceLinks) SetUsers(v HrefObject)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ResourceSetResourceLinks) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


