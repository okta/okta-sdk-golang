# CustomRoleAssignmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceSet** | Pointer to **string** | Resource Set ID | [optional] 
**Role** | Pointer to **string** | Custom Role ID | [optional] 
**Type** | Pointer to **string** | Standard role type | [optional] 

## Methods

### NewCustomRoleAssignmentSchema

`func NewCustomRoleAssignmentSchema() *CustomRoleAssignmentSchema`

NewCustomRoleAssignmentSchema instantiates a new CustomRoleAssignmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRoleAssignmentSchemaWithDefaults

`func NewCustomRoleAssignmentSchemaWithDefaults() *CustomRoleAssignmentSchema`

NewCustomRoleAssignmentSchemaWithDefaults instantiates a new CustomRoleAssignmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceSet

`func (o *CustomRoleAssignmentSchema) GetResourceSet() string`

GetResourceSet returns the ResourceSet field if non-nil, zero value otherwise.

### GetResourceSetOk

`func (o *CustomRoleAssignmentSchema) GetResourceSetOk() (*string, bool)`

GetResourceSetOk returns a tuple with the ResourceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSet

`func (o *CustomRoleAssignmentSchema) SetResourceSet(v string)`

SetResourceSet sets ResourceSet field to given value.

### HasResourceSet

`func (o *CustomRoleAssignmentSchema) HasResourceSet() bool`

HasResourceSet returns a boolean if a field has been set.

### GetRole

`func (o *CustomRoleAssignmentSchema) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CustomRoleAssignmentSchema) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CustomRoleAssignmentSchema) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CustomRoleAssignmentSchema) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetType

`func (o *CustomRoleAssignmentSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomRoleAssignmentSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomRoleAssignmentSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomRoleAssignmentSchema) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


