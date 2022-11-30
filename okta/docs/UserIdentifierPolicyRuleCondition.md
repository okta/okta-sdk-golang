# UserIdentifierPolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **string** |  | [optional] 
**Patterns** | Pointer to [**[]UserIdentifierConditionEvaluatorPattern**](UserIdentifierConditionEvaluatorPattern.md) |  | [optional] 
**Type** | Pointer to [**UserIdentifierType**](UserIdentifierType.md) |  | [optional] 

## Methods

### NewUserIdentifierPolicyRuleCondition

`func NewUserIdentifierPolicyRuleCondition() *UserIdentifierPolicyRuleCondition`

NewUserIdentifierPolicyRuleCondition instantiates a new UserIdentifierPolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentifierPolicyRuleConditionWithDefaults

`func NewUserIdentifierPolicyRuleConditionWithDefaults() *UserIdentifierPolicyRuleCondition`

NewUserIdentifierPolicyRuleConditionWithDefaults instantiates a new UserIdentifierPolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *UserIdentifierPolicyRuleCondition) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *UserIdentifierPolicyRuleCondition) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *UserIdentifierPolicyRuleCondition) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *UserIdentifierPolicyRuleCondition) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetPatterns

`func (o *UserIdentifierPolicyRuleCondition) GetPatterns() []UserIdentifierConditionEvaluatorPattern`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *UserIdentifierPolicyRuleCondition) GetPatternsOk() (*[]UserIdentifierConditionEvaluatorPattern, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *UserIdentifierPolicyRuleCondition) SetPatterns(v []UserIdentifierConditionEvaluatorPattern)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *UserIdentifierPolicyRuleCondition) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.

### GetType

`func (o *UserIdentifierPolicyRuleCondition) GetType() UserIdentifierType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserIdentifierPolicyRuleCondition) GetTypeOk() (*UserIdentifierType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserIdentifierPolicyRuleCondition) SetType(v UserIdentifierType)`

SetType sets Type field to given value.

### HasType

`func (o *UserIdentifierPolicyRuleCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


