# SessionViolationDetectionPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**SessionViolationDetectionPolicyRuleAllOfActions**](SessionViolationDetectionPolicyRuleAllOfActions.md) |  | [optional] 
**Conditions** | Pointer to [**SessionViolationDetectionPolicyRuleAllOfConditions**](SessionViolationDetectionPolicyRuleAllOfConditions.md) |  | [optional] 

## Methods

### NewSessionViolationDetectionPolicyRule

`func NewSessionViolationDetectionPolicyRule() *SessionViolationDetectionPolicyRule`

NewSessionViolationDetectionPolicyRule instantiates a new SessionViolationDetectionPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionViolationDetectionPolicyRuleWithDefaults

`func NewSessionViolationDetectionPolicyRuleWithDefaults() *SessionViolationDetectionPolicyRule`

NewSessionViolationDetectionPolicyRuleWithDefaults instantiates a new SessionViolationDetectionPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *SessionViolationDetectionPolicyRule) GetActions() SessionViolationDetectionPolicyRuleAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *SessionViolationDetectionPolicyRule) GetActionsOk() (*SessionViolationDetectionPolicyRuleAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *SessionViolationDetectionPolicyRule) SetActions(v SessionViolationDetectionPolicyRuleAllOfActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *SessionViolationDetectionPolicyRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *SessionViolationDetectionPolicyRule) GetConditions() SessionViolationDetectionPolicyRuleAllOfConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *SessionViolationDetectionPolicyRule) GetConditionsOk() (*SessionViolationDetectionPolicyRuleAllOfConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *SessionViolationDetectionPolicyRule) SetConditions(v SessionViolationDetectionPolicyRuleAllOfConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *SessionViolationDetectionPolicyRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


