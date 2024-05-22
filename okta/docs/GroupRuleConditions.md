# GroupRuleConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to [**GroupRuleExpression**](GroupRuleExpression.md) |  | [optional] 
**People** | Pointer to [**GroupRulePeopleCondition**](GroupRulePeopleCondition.md) |  | [optional] 

## Methods

### NewGroupRuleConditions

`func NewGroupRuleConditions() *GroupRuleConditions`

NewGroupRuleConditions instantiates a new GroupRuleConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRuleConditionsWithDefaults

`func NewGroupRuleConditionsWithDefaults() *GroupRuleConditions`

NewGroupRuleConditionsWithDefaults instantiates a new GroupRuleConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *GroupRuleConditions) GetExpression() GroupRuleExpression`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *GroupRuleConditions) GetExpressionOk() (*GroupRuleExpression, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *GroupRuleConditions) SetExpression(v GroupRuleExpression)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *GroupRuleConditions) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetPeople

`func (o *GroupRuleConditions) GetPeople() GroupRulePeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *GroupRuleConditions) GetPeopleOk() (*GroupRulePeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *GroupRuleConditions) SetPeople(v GroupRulePeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *GroupRuleConditions) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


