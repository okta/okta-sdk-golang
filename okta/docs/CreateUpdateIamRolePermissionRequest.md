# CreateUpdateIamRolePermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to **map[string]interface{}** | Conditions for further restricting a permission | [optional] 

## Methods

### NewCreateUpdateIamRolePermissionRequest

`func NewCreateUpdateIamRolePermissionRequest() *CreateUpdateIamRolePermissionRequest`

NewCreateUpdateIamRolePermissionRequest instantiates a new CreateUpdateIamRolePermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdateIamRolePermissionRequestWithDefaults

`func NewCreateUpdateIamRolePermissionRequestWithDefaults() *CreateUpdateIamRolePermissionRequest`

NewCreateUpdateIamRolePermissionRequestWithDefaults instantiates a new CreateUpdateIamRolePermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *CreateUpdateIamRolePermissionRequest) GetConditions() map[string]interface{}`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateUpdateIamRolePermissionRequest) GetConditionsOk() (*map[string]interface{}, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateUpdateIamRolePermissionRequest) SetConditions(v map[string]interface{})`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CreateUpdateIamRolePermissionRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *CreateUpdateIamRolePermissionRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *CreateUpdateIamRolePermissionRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


