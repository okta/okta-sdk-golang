# RoleTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentType** | Pointer to **string** | The assignment type of how the user receives this target | [optional] [readonly] 
**Expiration** | Pointer to **time.Time** | The expiry time stamp of the associated target. It&#39;s only included in the response if the associated target will expire. | [optional] [readonly] 
**Orn** | Pointer to **string** | The [Okta Resource Name (ORN)](https://support.okta.com/help/s/article/understanding-okta-resource-name-orn) of the app target or group target | [optional] [readonly] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewRoleTarget

`func NewRoleTarget() *RoleTarget`

NewRoleTarget instantiates a new RoleTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleTargetWithDefaults

`func NewRoleTargetWithDefaults() *RoleTarget`

NewRoleTargetWithDefaults instantiates a new RoleTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentType

`func (o *RoleTarget) GetAssignmentType() string`

GetAssignmentType returns the AssignmentType field if non-nil, zero value otherwise.

### GetAssignmentTypeOk

`func (o *RoleTarget) GetAssignmentTypeOk() (*string, bool)`

GetAssignmentTypeOk returns a tuple with the AssignmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentType

`func (o *RoleTarget) SetAssignmentType(v string)`

SetAssignmentType sets AssignmentType field to given value.

### HasAssignmentType

`func (o *RoleTarget) HasAssignmentType() bool`

HasAssignmentType returns a boolean if a field has been set.

### GetExpiration

`func (o *RoleTarget) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *RoleTarget) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *RoleTarget) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *RoleTarget) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetOrn

`func (o *RoleTarget) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *RoleTarget) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *RoleTarget) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *RoleTarget) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetLinks

`func (o *RoleTarget) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleTarget) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleTarget) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoleTarget) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


