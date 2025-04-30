# CreateIamRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the role | 
**Label** | **string** | Unique label for the role | 
**Permissions** | **[]string** | Array of permissions that the role will grant. See [Permissions](/openapi/okta-management/guides/roles/#permission). | 

## Methods

### NewCreateIamRoleRequest

`func NewCreateIamRoleRequest(description string, label string, permissions []string, ) *CreateIamRoleRequest`

NewCreateIamRoleRequest instantiates a new CreateIamRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIamRoleRequestWithDefaults

`func NewCreateIamRoleRequestWithDefaults() *CreateIamRoleRequest`

NewCreateIamRoleRequestWithDefaults instantiates a new CreateIamRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateIamRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIamRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIamRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLabel

`func (o *CreateIamRoleRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateIamRoleRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateIamRoleRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPermissions

`func (o *CreateIamRoleRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateIamRoleRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateIamRoleRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


