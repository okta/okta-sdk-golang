# AssignRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | | Role type                    | Description                                                 | |------------------------------|-------------------------------------------------------------| | ACCESS_CERTIFICATIONS_ADMIN  | Access Certifications Administrator IAM-based standard role | | ACCESS_REQUESTS_ADMIN        | Access Requests Administrator IAM-based standard role       | | API_ACCESS_MANAGEMENT_ADMIN  | Access Management Administrator standard role               | | APP_ADMIN                    | Application Administrator standard role                     | | CUSTOM                       | Custom admin role                                           | | GROUP_MEMBERSHIP_ADMIN       | Group Membership Administrator standard role                | | HELP_DESK_ADMIN              | Help Desk Administrator standard role                       | | ORG_ADMIN                    | Organizational Administrator standard role                  | | READ_ONLY_ADMIN              | Read-Only Administrator standard role                       | | REPORT_ADMIN                 | Report Administrator standard role                          | | SUPER_ADMIN                  | Super Administrator standard role                           | | USER_ADMIN                   | User Administrator standard role                            | | WORKFLOWS_ADMIN              | Workflows Administrator IAM-based standard role             | | [optional] 

## Methods

### NewAssignRoleRequest

`func NewAssignRoleRequest() *AssignRoleRequest`

NewAssignRoleRequest instantiates a new AssignRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignRoleRequestWithDefaults

`func NewAssignRoleRequestWithDefaults() *AssignRoleRequest`

NewAssignRoleRequestWithDefaults instantiates a new AssignRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AssignRoleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssignRoleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssignRoleRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssignRoleRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


