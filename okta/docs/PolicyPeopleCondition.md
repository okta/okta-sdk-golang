# PolicyPeopleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**GroupCondition**](GroupCondition.md) |  | [optional] 
**Users** | Pointer to [**UserCondition**](UserCondition.md) |  | [optional] 

## Methods

### NewPolicyPeopleCondition

`func NewPolicyPeopleCondition() *PolicyPeopleCondition`

NewPolicyPeopleCondition instantiates a new PolicyPeopleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyPeopleConditionWithDefaults

`func NewPolicyPeopleConditionWithDefaults() *PolicyPeopleCondition`

NewPolicyPeopleConditionWithDefaults instantiates a new PolicyPeopleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *PolicyPeopleCondition) GetGroups() GroupCondition`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PolicyPeopleCondition) GetGroupsOk() (*GroupCondition, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PolicyPeopleCondition) SetGroups(v GroupCondition)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PolicyPeopleCondition) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *PolicyPeopleCondition) GetUsers() UserCondition`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PolicyPeopleCondition) GetUsersOk() (*UserCondition, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PolicyPeopleCondition) SetUsers(v UserCondition)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PolicyPeopleCondition) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


