# RiskScorePolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | The level to match | [optional] 
**MinRiskLevel** | Pointer to **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;The minimum risk level to match. Only used in a Session Violation Detection (&#x60;SESSION_VIOLATION_DETECTION&#x60;) policy rule. | [optional] 

## Methods

### NewRiskScorePolicyRuleCondition

`func NewRiskScorePolicyRuleCondition() *RiskScorePolicyRuleCondition`

NewRiskScorePolicyRuleCondition instantiates a new RiskScorePolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskScorePolicyRuleConditionWithDefaults

`func NewRiskScorePolicyRuleConditionWithDefaults() *RiskScorePolicyRuleCondition`

NewRiskScorePolicyRuleConditionWithDefaults instantiates a new RiskScorePolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *RiskScorePolicyRuleCondition) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RiskScorePolicyRuleCondition) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RiskScorePolicyRuleCondition) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RiskScorePolicyRuleCondition) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMinRiskLevel

`func (o *RiskScorePolicyRuleCondition) GetMinRiskLevel() string`

GetMinRiskLevel returns the MinRiskLevel field if non-nil, zero value otherwise.

### GetMinRiskLevelOk

`func (o *RiskScorePolicyRuleCondition) GetMinRiskLevelOk() (*string, bool)`

GetMinRiskLevelOk returns a tuple with the MinRiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRiskLevel

`func (o *RiskScorePolicyRuleCondition) SetMinRiskLevel(v string)`

SetMinRiskLevel sets MinRiskLevel field to given value.

### HasMinRiskLevel

`func (o *RiskScorePolicyRuleCondition) HasMinRiskLevel() bool`

HasMinRiskLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


