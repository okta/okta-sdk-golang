# IamRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the role was created | [optional] [readonly] 
**Description** | **string** | Description of the role | 
**Id** | Pointer to **string** | Unique key for the role | [optional] [readonly] 
**Label** | **string** | Unique label for the role | 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the role was last updated | [optional] [readonly] 
**Permissions** | [**[]RolePermissionType**](RolePermissionType.md) | Array of permissions that the role will grant. See [Permission Types](https://developer.okta.com/docs/concepts/role-assignment/#permission-types). | 
**Links** | Pointer to [**IamRoleLinks**](IamRoleLinks.md) |  | [optional] 

## Methods

### NewIamRole

`func NewIamRole(description string, label string, permissions []RolePermissionType, ) *IamRole`

NewIamRole instantiates a new IamRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRoleWithDefaults

`func NewIamRoleWithDefaults() *IamRole`

NewIamRoleWithDefaults instantiates a new IamRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IamRole) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IamRole) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IamRole) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IamRole) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *IamRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamRole) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *IamRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *IamRole) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *IamRole) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *IamRole) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLastUpdated

`func (o *IamRole) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IamRole) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IamRole) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *IamRole) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPermissions

`func (o *IamRole) GetPermissions() []RolePermissionType`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamRole) GetPermissionsOk() (*[]RolePermissionType, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamRole) SetPermissions(v []RolePermissionType)`

SetPermissions sets Permissions field to given value.


### GetLinks

`func (o *IamRole) GetLinks() IamRoleLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IamRole) GetLinksOk() (*IamRoleLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IamRole) SetLinks(v IamRoleLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IamRole) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


