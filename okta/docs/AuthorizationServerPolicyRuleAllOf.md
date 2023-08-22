# AuthorizationServerPolicyRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**AuthorizationServerPolicyRuleActions**](AuthorizationServerPolicyRuleActions.md) |  | [optional] 
**Conditions** | Pointer to [**AuthorizationServerPolicyRuleConditions**](AuthorizationServerPolicyRuleConditions.md) |  | [optional] 

## Methods

### NewAuthorizationServerPolicyRuleAllOf

`func NewAuthorizationServerPolicyRuleAllOf() *AuthorizationServerPolicyRuleAllOf`

NewAuthorizationServerPolicyRuleAllOf instantiates a new AuthorizationServerPolicyRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyRuleAllOfWithDefaults

`func NewAuthorizationServerPolicyRuleAllOfWithDefaults() *AuthorizationServerPolicyRuleAllOf`

NewAuthorizationServerPolicyRuleAllOfWithDefaults instantiates a new AuthorizationServerPolicyRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AuthorizationServerPolicyRuleAllOf) GetActions() AuthorizationServerPolicyRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuthorizationServerPolicyRuleAllOf) GetActionsOk() (*AuthorizationServerPolicyRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuthorizationServerPolicyRuleAllOf) SetActions(v AuthorizationServerPolicyRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuthorizationServerPolicyRuleAllOf) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *AuthorizationServerPolicyRuleAllOf) GetConditions() AuthorizationServerPolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthorizationServerPolicyRuleAllOf) GetConditionsOk() (*AuthorizationServerPolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthorizationServerPolicyRuleAllOf) SetConditions(v AuthorizationServerPolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AuthorizationServerPolicyRuleAllOf) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


