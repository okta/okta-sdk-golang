# AuthenticatorEnrollmentPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**AuthenticatorEnrollmentPolicyRuleActions**](AuthenticatorEnrollmentPolicyRuleActions.md) |  | [optional] 
**Conditions** | Pointer to [**AuthenticatorEnrollmentPolicyRuleConditions**](AuthenticatorEnrollmentPolicyRuleConditions.md) |  | [optional] 

## Methods

### NewAuthenticatorEnrollmentPolicyRule

`func NewAuthenticatorEnrollmentPolicyRule() *AuthenticatorEnrollmentPolicyRule`

NewAuthenticatorEnrollmentPolicyRule instantiates a new AuthenticatorEnrollmentPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEnrollmentPolicyRuleWithDefaults

`func NewAuthenticatorEnrollmentPolicyRuleWithDefaults() *AuthenticatorEnrollmentPolicyRule`

NewAuthenticatorEnrollmentPolicyRuleWithDefaults instantiates a new AuthenticatorEnrollmentPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AuthenticatorEnrollmentPolicyRule) GetActions() AuthenticatorEnrollmentPolicyRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuthenticatorEnrollmentPolicyRule) GetActionsOk() (*AuthenticatorEnrollmentPolicyRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuthenticatorEnrollmentPolicyRule) SetActions(v AuthenticatorEnrollmentPolicyRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuthenticatorEnrollmentPolicyRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *AuthenticatorEnrollmentPolicyRule) GetConditions() AuthenticatorEnrollmentPolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthenticatorEnrollmentPolicyRule) GetConditionsOk() (*AuthenticatorEnrollmentPolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthenticatorEnrollmentPolicyRule) SetConditions(v AuthenticatorEnrollmentPolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AuthenticatorEnrollmentPolicyRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


