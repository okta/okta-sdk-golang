# AuthorizationServerPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**AuthorizationServerPolicyRuleActions**](AuthorizationServerPolicyRuleActions.md) |  | [optional] 
**Conditions** | Pointer to [**AuthorizationServerPolicyRuleConditions**](AuthorizationServerPolicyRuleConditions.md) |  | [optional] 

## Methods

### NewAuthorizationServerPolicyRule

`func NewAuthorizationServerPolicyRule() *AuthorizationServerPolicyRule`

NewAuthorizationServerPolicyRule instantiates a new AuthorizationServerPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyRuleWithDefaults

`func NewAuthorizationServerPolicyRuleWithDefaults() *AuthorizationServerPolicyRule`

NewAuthorizationServerPolicyRuleWithDefaults instantiates a new AuthorizationServerPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AuthorizationServerPolicyRule) GetActions() AuthorizationServerPolicyRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuthorizationServerPolicyRule) GetActionsOk() (*AuthorizationServerPolicyRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuthorizationServerPolicyRule) SetActions(v AuthorizationServerPolicyRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuthorizationServerPolicyRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *AuthorizationServerPolicyRule) GetConditions() AuthorizationServerPolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthorizationServerPolicyRule) GetConditionsOk() (*AuthorizationServerPolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthorizationServerPolicyRule) SetConditions(v AuthorizationServerPolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AuthorizationServerPolicyRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


