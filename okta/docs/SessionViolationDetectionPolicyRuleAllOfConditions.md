# SessionViolationDetectionPolicyRuleAllOfConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**PolicyNetworkCondition**](PolicyNetworkCondition.md) |  | [optional] 
**RiskScore** | Pointer to [**RiskScorePolicyRuleCondition**](RiskScorePolicyRuleCondition.md) |  | [optional] 

## Methods

### NewSessionViolationDetectionPolicyRuleAllOfConditions

`func NewSessionViolationDetectionPolicyRuleAllOfConditions() *SessionViolationDetectionPolicyRuleAllOfConditions`

NewSessionViolationDetectionPolicyRuleAllOfConditions instantiates a new SessionViolationDetectionPolicyRuleAllOfConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionViolationDetectionPolicyRuleAllOfConditionsWithDefaults

`func NewSessionViolationDetectionPolicyRuleAllOfConditionsWithDefaults() *SessionViolationDetectionPolicyRuleAllOfConditions`

NewSessionViolationDetectionPolicyRuleAllOfConditionsWithDefaults instantiates a new SessionViolationDetectionPolicyRuleAllOfConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *SessionViolationDetectionPolicyRuleAllOfConditions) GetNetwork() PolicyNetworkCondition`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SessionViolationDetectionPolicyRuleAllOfConditions) GetNetworkOk() (*PolicyNetworkCondition, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SessionViolationDetectionPolicyRuleAllOfConditions) SetNetwork(v PolicyNetworkCondition)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *SessionViolationDetectionPolicyRuleAllOfConditions) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRiskScore

`func (o *SessionViolationDetectionPolicyRuleAllOfConditions) GetRiskScore() RiskScorePolicyRuleCondition`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *SessionViolationDetectionPolicyRuleAllOfConditions) GetRiskScoreOk() (*RiskScorePolicyRuleCondition, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *SessionViolationDetectionPolicyRuleAllOfConditions) SetRiskScore(v RiskScorePolicyRuleCondition)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *SessionViolationDetectionPolicyRuleAllOfConditions) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


