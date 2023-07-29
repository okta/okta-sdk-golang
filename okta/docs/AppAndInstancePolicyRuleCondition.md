# AppAndInstancePolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | Pointer to [**[]AppAndInstanceConditionEvaluatorAppOrInstance**](AppAndInstanceConditionEvaluatorAppOrInstance.md) |  | [optional] 
**Include** | Pointer to [**[]AppAndInstanceConditionEvaluatorAppOrInstance**](AppAndInstanceConditionEvaluatorAppOrInstance.md) |  | [optional] 

## Methods

### NewAppAndInstancePolicyRuleCondition

`func NewAppAndInstancePolicyRuleCondition() *AppAndInstancePolicyRuleCondition`

NewAppAndInstancePolicyRuleCondition instantiates a new AppAndInstancePolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAndInstancePolicyRuleConditionWithDefaults

`func NewAppAndInstancePolicyRuleConditionWithDefaults() *AppAndInstancePolicyRuleCondition`

NewAppAndInstancePolicyRuleConditionWithDefaults instantiates a new AppAndInstancePolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *AppAndInstancePolicyRuleCondition) GetExclude() []AppAndInstanceConditionEvaluatorAppOrInstance`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *AppAndInstancePolicyRuleCondition) GetExcludeOk() (*[]AppAndInstanceConditionEvaluatorAppOrInstance, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *AppAndInstancePolicyRuleCondition) SetExclude(v []AppAndInstanceConditionEvaluatorAppOrInstance)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *AppAndInstancePolicyRuleCondition) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInclude

`func (o *AppAndInstancePolicyRuleCondition) GetInclude() []AppAndInstanceConditionEvaluatorAppOrInstance`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *AppAndInstancePolicyRuleCondition) GetIncludeOk() (*[]AppAndInstanceConditionEvaluatorAppOrInstance, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *AppAndInstancePolicyRuleCondition) SetInclude(v []AppAndInstanceConditionEvaluatorAppOrInstance)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *AppAndInstancePolicyRuleCondition) HasInclude() bool`

HasInclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


