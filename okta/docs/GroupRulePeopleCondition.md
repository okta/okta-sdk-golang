# GroupRulePeopleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**GroupRuleGroupCondition**](GroupRuleGroupCondition.md) |  | [optional] 
**Users** | Pointer to [**GroupRuleUserCondition**](GroupRuleUserCondition.md) |  | [optional] 

## Methods

### NewGroupRulePeopleCondition

`func NewGroupRulePeopleCondition() *GroupRulePeopleCondition`

NewGroupRulePeopleCondition instantiates a new GroupRulePeopleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRulePeopleConditionWithDefaults

`func NewGroupRulePeopleConditionWithDefaults() *GroupRulePeopleCondition`

NewGroupRulePeopleConditionWithDefaults instantiates a new GroupRulePeopleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *GroupRulePeopleCondition) GetGroups() GroupRuleGroupCondition`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupRulePeopleCondition) GetGroupsOk() (*GroupRuleGroupCondition, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupRulePeopleCondition) SetGroups(v GroupRuleGroupCondition)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GroupRulePeopleCondition) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *GroupRulePeopleCondition) GetUsers() GroupRuleUserCondition`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupRulePeopleCondition) GetUsersOk() (*GroupRuleUserCondition, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupRulePeopleCondition) SetUsers(v GroupRuleUserCondition)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupRulePeopleCondition) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


