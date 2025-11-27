# CustomRoleAssignmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceSet** | **string** | Resource set ID | 
**Role** | **string** | Custom role ID | 
**Type** | **string** | Specify a [standard admin role](/openapi/okta-management/guides/roles/#standard-roles), an [IAM-based standard role](/openapi/okta-management/guides/roles/#iam-based-standard-roles), or &#x60;CUSTOM&#x60; for a custom role type: | 

## Methods

### NewCustomRoleAssignmentSchema

`func NewCustomRoleAssignmentSchema(resourceSet string, role string, type_ string, ) *CustomRoleAssignmentSchema`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


