# ResourceSetBindingMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]ResourceSetBindingMember**](ResourceSetBindingMember.md) |  | [optional] 
**Links** | Pointer to [**ResourceSetBindingMembersLinks**](ResourceSetBindingMembersLinks.md) |  | [optional] 

## Methods

### NewResourceSetBindingMembers

`func NewResourceSetBindingMembers() *ResourceSetBindingMembers`

NewResourceSetBindingMembers instantiates a new ResourceSetBindingMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetBindingMembersWithDefaults

`func NewResourceSetBindingMembersWithDefaults() *ResourceSetBindingMembers`

NewResourceSetBindingMembersWithDefaults instantiates a new ResourceSetBindingMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ResourceSetBindingMembers) GetMembers() []ResourceSetBindingMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ResourceSetBindingMembers) GetMembersOk() (*[]ResourceSetBindingMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ResourceSetBindingMembers) SetMembers(v []ResourceSetBindingMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ResourceSetBindingMembers) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSetBindingMembers) GetLinks() ResourceSetBindingMembersLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSetBindingMembers) GetLinksOk() (*ResourceSetBindingMembersLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSetBindingMembers) SetLinks(v ResourceSetBindingMembersLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSetBindingMembers) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


