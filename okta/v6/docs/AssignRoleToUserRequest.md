# AssignRoleToUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Specify the standard or IAM-based role type. See [standard roles](/openapi/okta-management/guides/roles/#standard-roles). | [optional] 
**ResourceSet** | Pointer to **string** | Resource set ID | [optional] 
**Role** | Pointer to **string** | Custom role ID | [optional] 

## Methods

### NewAssignRoleToUserRequest

`func NewAssignRoleToUserRequest() *AssignRoleToUserRequest`

NewAssignRoleToUserRequest instantiates a new AssignRoleToUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignRoleToUserRequestWithDefaults

`func NewAssignRoleToUserRequestWithDefaults() *AssignRoleToUserRequest`

NewAssignRoleToUserRequestWithDefaults instantiates a new AssignRoleToUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AssignRoleToUserRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignRoleToUserRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignRoleToUserRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssignRoleToUserRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResourceSet

`func (o *AssignRoleToUserRequest) GetResourceSet() string`

GetResourceSet returns the ResourceSet field if non-nil, zero value otherwise.

### GetResourceSetOk

`func (o *AssignRoleToUserRequest) GetResourceSetOk() (*string, bool)`

GetResourceSetOk returns a tuple with the ResourceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSet

`func (o *AssignRoleToUserRequest) SetResourceSet(v string)`

SetResourceSet sets ResourceSet field to given value.

### HasResourceSet

`func (o *AssignRoleToUserRequest) HasResourceSet() bool`

HasResourceSet returns a boolean if a field has been set.

### GetRole

`func (o *AssignRoleToUserRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AssignRoleToUserRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AssignRoleToUserRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AssignRoleToUserRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


