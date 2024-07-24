# EntityRiskPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**EntityRiskPolicyRuleAllOfActions**](EntityRiskPolicyRuleAllOfActions.md) |  | [optional] 
**Conditions** | Pointer to [**EntityRiskPolicyRuleAllOfConditions**](EntityRiskPolicyRuleAllOfConditions.md) |  | [optional] 

## Methods

### NewEntityRiskPolicyRule

`func NewEntityRiskPolicyRule() *EntityRiskPolicyRule`

NewEntityRiskPolicyRule instantiates a new EntityRiskPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityRiskPolicyRuleWithDefaults

`func NewEntityRiskPolicyRuleWithDefaults() *EntityRiskPolicyRule`

NewEntityRiskPolicyRuleWithDefaults instantiates a new EntityRiskPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *EntityRiskPolicyRule) GetActions() EntityRiskPolicyRuleAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *EntityRiskPolicyRule) GetActionsOk() (*EntityRiskPolicyRuleAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *EntityRiskPolicyRule) SetActions(v EntityRiskPolicyRuleAllOfActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *EntityRiskPolicyRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *EntityRiskPolicyRule) GetConditions() EntityRiskPolicyRuleAllOfConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EntityRiskPolicyRule) GetConditionsOk() (*EntityRiskPolicyRuleAllOfConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EntityRiskPolicyRule) SetConditions(v EntityRiskPolicyRuleAllOfConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *EntityRiskPolicyRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


