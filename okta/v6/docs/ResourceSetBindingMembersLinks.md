# ResourceSetBindingMembersLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**HrefObjectSelfLink**](HrefObjectSelfLink.md) |  | [optional] 
**Next** | Pointer to [**HrefObject**](HrefObject.md) | Link to the next list of binding members for the specified role and resource set | [optional] 
**Binding** | Pointer to [**HrefObjectBindingLink**](HrefObjectBindingLink.md) |  | [optional] 

## Methods

### NewResourceSetBindingMembersLinks

`func NewResourceSetBindingMembersLinks() *ResourceSetBindingMembersLinks`

NewResourceSetBindingMembersLinks instantiates a new ResourceSetBindingMembersLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetBindingMembersLinksWithDefaults

`func NewResourceSetBindingMembersLinksWithDefaults() *ResourceSetBindingMembersLinks`

NewResourceSetBindingMembersLinksWithDefaults instantiates a new ResourceSetBindingMembersLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ResourceSetBindingMembersLinks) GetSelf() HrefObjectSelfLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ResourceSetBindingMembersLinks) GetSelfOk() (*HrefObjectSelfLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ResourceSetBindingMembersLinks) SetSelf(v HrefObjectSelfLink)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ResourceSetBindingMembersLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetNext

`func (o *ResourceSetBindingMembersLinks) GetNext() HrefObject`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ResourceSetBindingMembersLinks) GetNextOk() (*HrefObject, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ResourceSetBindingMembersLinks) SetNext(v HrefObject)`

SetNext sets Next field to given value.

### HasNext

`func (o *ResourceSetBindingMembersLinks) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetBinding

`func (o *ResourceSetBindingMembersLinks) GetBinding() HrefObjectBindingLink`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *ResourceSetBindingMembersLinks) GetBindingOk() (*HrefObjectBindingLink, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *ResourceSetBindingMembersLinks) SetBinding(v HrefObjectBindingLink)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *ResourceSetBindingMembersLinks) HasBinding() bool`

HasBinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


