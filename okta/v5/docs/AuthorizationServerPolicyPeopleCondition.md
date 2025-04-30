# AuthorizationServerPolicyPeopleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**AuthorizationServerPolicyRuleGroupCondition**](AuthorizationServerPolicyRuleGroupCondition.md) |  | [optional] 
**Users** | Pointer to [**AuthorizationServerPolicyRuleUserCondition**](AuthorizationServerPolicyRuleUserCondition.md) |  | [optional] 

## Methods

### NewAuthorizationServerPolicyPeopleCondition

`func NewAuthorizationServerPolicyPeopleCondition() *AuthorizationServerPolicyPeopleCondition`

NewAuthorizationServerPolicyPeopleCondition instantiates a new AuthorizationServerPolicyPeopleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyPeopleConditionWithDefaults

`func NewAuthorizationServerPolicyPeopleConditionWithDefaults() *AuthorizationServerPolicyPeopleCondition`

NewAuthorizationServerPolicyPeopleConditionWithDefaults instantiates a new AuthorizationServerPolicyPeopleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *AuthorizationServerPolicyPeopleCondition) GetGroups() AuthorizationServerPolicyRuleGroupCondition`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AuthorizationServerPolicyPeopleCondition) GetGroupsOk() (*AuthorizationServerPolicyRuleGroupCondition, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AuthorizationServerPolicyPeopleCondition) SetGroups(v AuthorizationServerPolicyRuleGroupCondition)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AuthorizationServerPolicyPeopleCondition) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *AuthorizationServerPolicyPeopleCondition) GetUsers() AuthorizationServerPolicyRuleUserCondition`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AuthorizationServerPolicyPeopleCondition) GetUsersOk() (*AuthorizationServerPolicyRuleUserCondition, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AuthorizationServerPolicyPeopleCondition) SetUsers(v AuthorizationServerPolicyRuleUserCondition)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AuthorizationServerPolicyPeopleCondition) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


