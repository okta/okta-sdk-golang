# AccessPolicyRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**AccessPolicyRuleActions**](AccessPolicyRuleActions.md) |  | [optional] 
**Conditions** | Pointer to [**AccessPolicyRuleConditions**](AccessPolicyRuleConditions.md) |  | [optional] 

## Methods

### NewAccessPolicyRuleAllOf

`func NewAccessPolicyRuleAllOf() *AccessPolicyRuleAllOf`

NewAccessPolicyRuleAllOf instantiates a new AccessPolicyRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyRuleAllOfWithDefaults

`func NewAccessPolicyRuleAllOfWithDefaults() *AccessPolicyRuleAllOf`

NewAccessPolicyRuleAllOfWithDefaults instantiates a new AccessPolicyRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AccessPolicyRuleAllOf) GetActions() AccessPolicyRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AccessPolicyRuleAllOf) GetActionsOk() (*AccessPolicyRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AccessPolicyRuleAllOf) SetActions(v AccessPolicyRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AccessPolicyRuleAllOf) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *AccessPolicyRuleAllOf) GetConditions() AccessPolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AccessPolicyRuleAllOf) GetConditionsOk() (*AccessPolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AccessPolicyRuleAllOf) SetConditions(v AccessPolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AccessPolicyRuleAllOf) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


