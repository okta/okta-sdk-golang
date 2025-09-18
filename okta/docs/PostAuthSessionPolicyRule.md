# PostAuthSessionPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**PostAuthSessionPolicyRuleAllOfActions**](PostAuthSessionPolicyRuleAllOfActions.md) |  | [optional] 
**Conditions** | Pointer to [**PostAuthSessionPolicyRuleAllOfConditions**](PostAuthSessionPolicyRuleAllOfConditions.md) |  | [optional] 

## Methods

### NewPostAuthSessionPolicyRule

`func NewPostAuthSessionPolicyRule() *PostAuthSessionPolicyRule`

NewPostAuthSessionPolicyRule instantiates a new PostAuthSessionPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAuthSessionPolicyRuleWithDefaults

`func NewPostAuthSessionPolicyRuleWithDefaults() *PostAuthSessionPolicyRule`

NewPostAuthSessionPolicyRuleWithDefaults instantiates a new PostAuthSessionPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *PostAuthSessionPolicyRule) GetActions() PostAuthSessionPolicyRuleAllOfActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PostAuthSessionPolicyRule) GetActionsOk() (*PostAuthSessionPolicyRuleAllOfActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PostAuthSessionPolicyRule) SetActions(v PostAuthSessionPolicyRuleAllOfActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PostAuthSessionPolicyRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *PostAuthSessionPolicyRule) GetConditions() PostAuthSessionPolicyRuleAllOfConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PostAuthSessionPolicyRule) GetConditionsOk() (*PostAuthSessionPolicyRuleAllOfConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PostAuthSessionPolicyRule) SetConditions(v PostAuthSessionPolicyRuleAllOfConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PostAuthSessionPolicyRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


