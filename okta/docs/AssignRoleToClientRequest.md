# AssignRoleToClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Standard role type | [optional] 
**ResourceSet** | Pointer to **string** | Resource Set ID | [optional] 
**Role** | Pointer to **string** | Custom Role ID | [optional] 

## Methods

### NewAssignRoleToClientRequest

`func NewAssignRoleToClientRequest() *AssignRoleToClientRequest`

NewAssignRoleToClientRequest instantiates a new AssignRoleToClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignRoleToClientRequestWithDefaults

`func NewAssignRoleToClientRequestWithDefaults() *AssignRoleToClientRequest`

NewAssignRoleToClientRequestWithDefaults instantiates a new AssignRoleToClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AssignRoleToClientRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignRoleToClientRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignRoleToClientRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssignRoleToClientRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResourceSet

`func (o *AssignRoleToClientRequest) GetResourceSet() string`

GetResourceSet returns the ResourceSet field if non-nil, zero value otherwise.

### GetResourceSetOk

`func (o *AssignRoleToClientRequest) GetResourceSetOk() (*string, bool)`

GetResourceSetOk returns a tuple with the ResourceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSet

`func (o *AssignRoleToClientRequest) SetResourceSet(v string)`

SetResourceSet sets ResourceSet field to given value.

### HasResourceSet

`func (o *AssignRoleToClientRequest) HasResourceSet() bool`

HasResourceSet returns a boolean if a field has been set.

### GetRole

`func (o *AssignRoleToClientRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AssignRoleToClientRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AssignRoleToClientRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AssignRoleToClientRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


