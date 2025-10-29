# EntityRiskPolicyRuleConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityRisk** | Pointer to [**EntityRiskScorePolicyRuleCondition**](EntityRiskScorePolicyRuleCondition.md) |  | [optional] 
**People** | Pointer to [**PolicyPeopleCondition**](PolicyPeopleCondition.md) |  | [optional] 
**RiskDetectionTypes** | Pointer to [**RiskDetectionTypesPolicyRuleCondition**](RiskDetectionTypesPolicyRuleCondition.md) |  | [optional] 

## Methods

### NewEntityRiskPolicyRuleConditions

`func NewEntityRiskPolicyRuleConditions() *EntityRiskPolicyRuleConditions`

NewEntityRiskPolicyRuleConditions instantiates a new EntityRiskPolicyRuleConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityRiskPolicyRuleConditionsWithDefaults

`func NewEntityRiskPolicyRuleConditionsWithDefaults() *EntityRiskPolicyRuleConditions`

NewEntityRiskPolicyRuleConditionsWithDefaults instantiates a new EntityRiskPolicyRuleConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityRisk

`func (o *EntityRiskPolicyRuleConditions) GetEntityRisk() EntityRiskScorePolicyRuleCondition`

GetEntityRisk returns the EntityRisk field if non-nil, zero value otherwise.

### GetEntityRiskOk

`func (o *EntityRiskPolicyRuleConditions) GetEntityRiskOk() (*EntityRiskScorePolicyRuleCondition, bool)`

GetEntityRiskOk returns a tuple with the EntityRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityRisk

`func (o *EntityRiskPolicyRuleConditions) SetEntityRisk(v EntityRiskScorePolicyRuleCondition)`

SetEntityRisk sets EntityRisk field to given value.

### HasEntityRisk

`func (o *EntityRiskPolicyRuleConditions) HasEntityRisk() bool`

HasEntityRisk returns a boolean if a field has been set.

### GetPeople

`func (o *EntityRiskPolicyRuleConditions) GetPeople() PolicyPeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *EntityRiskPolicyRuleConditions) GetPeopleOk() (*PolicyPeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *EntityRiskPolicyRuleConditions) SetPeople(v PolicyPeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *EntityRiskPolicyRuleConditions) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetRiskDetectionTypes

`func (o *EntityRiskPolicyRuleConditions) GetRiskDetectionTypes() RiskDetectionTypesPolicyRuleCondition`

GetRiskDetectionTypes returns the RiskDetectionTypes field if non-nil, zero value otherwise.

### GetRiskDetectionTypesOk

`func (o *EntityRiskPolicyRuleConditions) GetRiskDetectionTypesOk() (*RiskDetectionTypesPolicyRuleCondition, bool)`

GetRiskDetectionTypesOk returns a tuple with the RiskDetectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskDetectionTypes

`func (o *EntityRiskPolicyRuleConditions) SetRiskDetectionTypes(v RiskDetectionTypesPolicyRuleCondition)`

SetRiskDetectionTypes sets RiskDetectionTypes field to given value.

### HasRiskDetectionTypes

`func (o *EntityRiskPolicyRuleConditions) HasRiskDetectionTypes() bool`

HasRiskDetectionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


