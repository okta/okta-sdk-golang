# IdpDiscoveryPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**IdpPolicyRuleAction**](IdpPolicyRuleAction.md) |  | [optional] 
**Conditions** | Pointer to [**IdpDiscoveryPolicyRuleCondition**](IdpDiscoveryPolicyRuleCondition.md) |  | [optional] 

## Methods

### NewIdpDiscoveryPolicyRule

`func NewIdpDiscoveryPolicyRule() *IdpDiscoveryPolicyRule`

NewIdpDiscoveryPolicyRule instantiates a new IdpDiscoveryPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpDiscoveryPolicyRuleWithDefaults

`func NewIdpDiscoveryPolicyRuleWithDefaults() *IdpDiscoveryPolicyRule`

NewIdpDiscoveryPolicyRuleWithDefaults instantiates a new IdpDiscoveryPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *IdpDiscoveryPolicyRule) GetActions() IdpPolicyRuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *IdpDiscoveryPolicyRule) GetActionsOk() (*IdpPolicyRuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *IdpDiscoveryPolicyRule) SetActions(v IdpPolicyRuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *IdpDiscoveryPolicyRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *IdpDiscoveryPolicyRule) GetConditions() IdpDiscoveryPolicyRuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *IdpDiscoveryPolicyRule) GetConditionsOk() (*IdpDiscoveryPolicyRuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *IdpDiscoveryPolicyRule) SetConditions(v IdpDiscoveryPolicyRuleCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *IdpDiscoveryPolicyRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


