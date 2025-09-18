# ResourceSetBindingResponseLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**ResourceSet** | Pointer to [**HrefObjectResourceSetLink**](HrefObjectResourceSetLink.md) |  | [optional] 
**Members** | Pointer to [**HrefObjectMembersLink**](HrefObjectMembersLink.md) |  | [optional] 

## Methods

### NewResourceSetBindingResponseLinks

`func NewResourceSetBindingResponseLinks() *ResourceSetBindingResponseLinks`

NewResourceSetBindingResponseLinks instantiates a new ResourceSetBindingResponseLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetBindingResponseLinksWithDefaults

`func NewResourceSetBindingResponseLinksWithDefaults() *ResourceSetBindingResponseLinks`

NewResourceSetBindingResponseLinksWithDefaults instantiates a new ResourceSetBindingResponseLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ResourceSetBindingResponseLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ResourceSetBindingResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ResourceSetBindingResponseLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ResourceSetBindingResponseLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetResourceSet

`func (o *ResourceSetBindingResponseLinks) GetResourceSet() HrefObjectResourceSetLink`

GetResourceSet returns the ResourceSet field if non-nil, zero value otherwise.

### GetResourceSetOk

`func (o *ResourceSetBindingResponseLinks) GetResourceSetOk() (*HrefObjectResourceSetLink, bool)`

GetResourceSetOk returns a tuple with the ResourceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSet

`func (o *ResourceSetBindingResponseLinks) SetResourceSet(v HrefObjectResourceSetLink)`

SetResourceSet sets ResourceSet field to given value.

### HasResourceSet

`func (o *ResourceSetBindingResponseLinks) HasResourceSet() bool`

HasResourceSet returns a boolean if a field has been set.

### GetMembers

`func (o *ResourceSetBindingResponseLinks) GetMembers() HrefObjectMembersLink`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ResourceSetBindingResponseLinks) GetMembersOk() (*HrefObjectMembersLink, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ResourceSetBindingResponseLinks) SetMembers(v HrefObjectMembersLink)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ResourceSetBindingResponseLinks) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


