# ContextPolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Migrated** | Pointer to **bool** |  | [optional] 
**Platform** | Pointer to [**DevicePolicyRuleConditionPlatform**](DevicePolicyRuleConditionPlatform.md) |  | [optional] 
**Rooted** | Pointer to **bool** |  | [optional] 
**TrustLevel** | Pointer to [**DevicePolicyTrustLevel**](DevicePolicyTrustLevel.md) |  | [optional] 
**Expression** | Pointer to **string** |  | [optional] 

## Methods

### NewContextPolicyRuleCondition

`func NewContextPolicyRuleCondition() *ContextPolicyRuleCondition`

NewContextPolicyRuleCondition instantiates a new ContextPolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextPolicyRuleConditionWithDefaults

`func NewContextPolicyRuleConditionWithDefaults() *ContextPolicyRuleCondition`

NewContextPolicyRuleConditionWithDefaults instantiates a new ContextPolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrated

`func (o *ContextPolicyRuleCondition) GetMigrated() bool`

GetMigrated returns the Migrated field if non-nil, zero value otherwise.

### GetMigratedOk

`func (o *ContextPolicyRuleCondition) GetMigratedOk() (*bool, bool)`

GetMigratedOk returns a tuple with the Migrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrated

`func (o *ContextPolicyRuleCondition) SetMigrated(v bool)`

SetMigrated sets Migrated field to given value.

### HasMigrated

`func (o *ContextPolicyRuleCondition) HasMigrated() bool`

HasMigrated returns a boolean if a field has been set.

### GetPlatform

`func (o *ContextPolicyRuleCondition) GetPlatform() DevicePolicyRuleConditionPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ContextPolicyRuleCondition) GetPlatformOk() (*DevicePolicyRuleConditionPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ContextPolicyRuleCondition) SetPlatform(v DevicePolicyRuleConditionPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ContextPolicyRuleCondition) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRooted

`func (o *ContextPolicyRuleCondition) GetRooted() bool`

GetRooted returns the Rooted field if non-nil, zero value otherwise.

### GetRootedOk

`func (o *ContextPolicyRuleCondition) GetRootedOk() (*bool, bool)`

GetRootedOk returns a tuple with the Rooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooted

`func (o *ContextPolicyRuleCondition) SetRooted(v bool)`

SetRooted sets Rooted field to given value.

### HasRooted

`func (o *ContextPolicyRuleCondition) HasRooted() bool`

HasRooted returns a boolean if a field has been set.

### GetTrustLevel

`func (o *ContextPolicyRuleCondition) GetTrustLevel() DevicePolicyTrustLevel`

GetTrustLevel returns the TrustLevel field if non-nil, zero value otherwise.

### GetTrustLevelOk

`func (o *ContextPolicyRuleCondition) GetTrustLevelOk() (*DevicePolicyTrustLevel, bool)`

GetTrustLevelOk returns a tuple with the TrustLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustLevel

`func (o *ContextPolicyRuleCondition) SetTrustLevel(v DevicePolicyTrustLevel)`

SetTrustLevel sets TrustLevel field to given value.

### HasTrustLevel

`func (o *ContextPolicyRuleCondition) HasTrustLevel() bool`

HasTrustLevel returns a boolean if a field has been set.

### GetExpression

`func (o *ContextPolicyRuleCondition) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ContextPolicyRuleCondition) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ContextPolicyRuleCondition) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ContextPolicyRuleCondition) HasExpression() bool`

HasExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


