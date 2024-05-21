# PlatformPolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | Pointer to [**[]PlatformConditionEvaluatorPlatform**](PlatformConditionEvaluatorPlatform.md) |  | [optional] 
**Include** | Pointer to [**[]PlatformConditionEvaluatorPlatform**](PlatformConditionEvaluatorPlatform.md) |  | [optional] 

## Methods

### NewPlatformPolicyRuleCondition

`func NewPlatformPolicyRuleCondition() *PlatformPolicyRuleCondition`

NewPlatformPolicyRuleCondition instantiates a new PlatformPolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformPolicyRuleConditionWithDefaults

`func NewPlatformPolicyRuleConditionWithDefaults() *PlatformPolicyRuleCondition`

NewPlatformPolicyRuleConditionWithDefaults instantiates a new PlatformPolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *PlatformPolicyRuleCondition) GetExclude() []PlatformConditionEvaluatorPlatform`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *PlatformPolicyRuleCondition) GetExcludeOk() (*[]PlatformConditionEvaluatorPlatform, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *PlatformPolicyRuleCondition) SetExclude(v []PlatformConditionEvaluatorPlatform)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *PlatformPolicyRuleCondition) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInclude

`func (o *PlatformPolicyRuleCondition) GetInclude() []PlatformConditionEvaluatorPlatform`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *PlatformPolicyRuleCondition) GetIncludeOk() (*[]PlatformConditionEvaluatorPlatform, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *PlatformPolicyRuleCondition) SetInclude(v []PlatformConditionEvaluatorPlatform)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *PlatformPolicyRuleCondition) HasInclude() bool`

HasInclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


