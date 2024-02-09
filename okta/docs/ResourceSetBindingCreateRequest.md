# ResourceSetBindingCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to **[]string** |  | [optional] 
**Role** | Pointer to **string** | Unique key for the role | [optional] 

## Methods

### NewResourceSetBindingCreateRequest

`func NewResourceSetBindingCreateRequest() *ResourceSetBindingCreateRequest`

NewResourceSetBindingCreateRequest instantiates a new ResourceSetBindingCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSetBindingCreateRequestWithDefaults

`func NewResourceSetBindingCreateRequestWithDefaults() *ResourceSetBindingCreateRequest`

NewResourceSetBindingCreateRequestWithDefaults instantiates a new ResourceSetBindingCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ResourceSetBindingCreateRequest) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ResourceSetBindingCreateRequest) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ResourceSetBindingCreateRequest) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ResourceSetBindingCreateRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetRole

`func (o *ResourceSetBindingCreateRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ResourceSetBindingCreateRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ResourceSetBindingCreateRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ResourceSetBindingCreateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


