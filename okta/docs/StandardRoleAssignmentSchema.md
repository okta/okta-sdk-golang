# StandardRoleAssignmentSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Specify a [standard admin role](/openapi/okta-management/guides/roles/#standard-roles), an [IAM-based standard role](/openapi/okta-management/guides/roles/#iam-based-standard-roles), or &#x60;CUSTOM&#x60; for a custom role type: | 

## Methods

### NewStandardRoleAssignmentSchema

`func NewStandardRoleAssignmentSchema(type_ string, ) *StandardRoleAssignmentSchema`

NewStandardRoleAssignmentSchema instantiates a new StandardRoleAssignmentSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardRoleAssignmentSchemaWithDefaults

`func NewStandardRoleAssignmentSchemaWithDefaults() *StandardRoleAssignmentSchema`

NewStandardRoleAssignmentSchemaWithDefaults instantiates a new StandardRoleAssignmentSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StandardRoleAssignmentSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StandardRoleAssignmentSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StandardRoleAssignmentSchema) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


