# IamRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]IamRole**](IamRole.md) |  | [optional] 
**Links** | Pointer to [**IamRolesLinks**](IamRolesLinks.md) |  | [optional] 

## Methods

### NewIamRoles

`func NewIamRoles() *IamRoles`

NewIamRoles instantiates a new IamRoles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRolesWithDefaults

`func NewIamRolesWithDefaults() *IamRoles`

NewIamRolesWithDefaults instantiates a new IamRoles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *IamRoles) GetRoles() []IamRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamRoles) GetRolesOk() (*[]IamRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamRoles) SetRoles(v []IamRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamRoles) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetLinks

`func (o *IamRoles) GetLinks() IamRolesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IamRoles) GetLinksOk() (*IamRolesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IamRoles) SetLinks(v IamRolesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IamRoles) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


