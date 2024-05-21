# ResourceSetBindings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]ResourceSetBindingRole**](ResourceSetBindingRole.md) |  | [optional] 
**Links** | Pointer to [**ResourceSetBindingResponseLinks**](ResourceSetBindingResponseLinks.md) |  | [optional] 

## Methods

### NewResourceSetBindings

`func NewResourceSetBindings() *ResourceSetBindings`

NewResourceSetBindings instantiates a new ResourceSetBindings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetBindingsWithDefaults

`func NewResourceSetBindingsWithDefaults() *ResourceSetBindings`

NewResourceSetBindingsWithDefaults instantiates a new ResourceSetBindings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *ResourceSetBindings) GetRoles() []ResourceSetBindingRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ResourceSetBindings) GetRolesOk() (*[]ResourceSetBindingRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ResourceSetBindings) SetRoles(v []ResourceSetBindingRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ResourceSetBindings) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceSetBindings) GetLinks() ResourceSetBindingResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceSetBindings) GetLinksOk() (*ResourceSetBindingResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceSetBindings) SetLinks(v ResourceSetBindingResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ResourceSetBindings) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


