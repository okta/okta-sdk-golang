# PasswordPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**PasswordPolicyRuleActions**](PasswordPolicyRuleActions.md) |  | [optional] 
**Conditions** | Pointer to [**PasswordPolicyRuleConditions**](PasswordPolicyRuleConditions.md) |  | [optional] 

## Methods

### NewPasswordPolicyRule

`func NewPasswordPolicyRule() *PasswordPolicyRule`

NewPasswordPolicyRule instantiates a new PasswordPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyRuleWithDefaults

`func NewPasswordPolicyRuleWithDefaults() *PasswordPolicyRule`

NewPasswordPolicyRuleWithDefaults instantiates a new PasswordPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *PasswordPolicyRule) GetActions() PasswordPolicyRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PasswordPolicyRule) GetActionsOk() (*PasswordPolicyRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PasswordPolicyRule) SetActions(v PasswordPolicyRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PasswordPolicyRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *PasswordPolicyRule) GetConditions() PasswordPolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PasswordPolicyRule) GetConditionsOk() (*PasswordPolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PasswordPolicyRule) SetConditions(v PasswordPolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PasswordPolicyRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


