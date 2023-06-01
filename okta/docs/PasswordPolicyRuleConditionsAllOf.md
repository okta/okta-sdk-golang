# PasswordPolicyRuleConditionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**PolicyNetworkCondition**](PolicyNetworkCondition.md) |  | [optional] 
**People** | Pointer to [**PolicyPeopleCondition**](PolicyPeopleCondition.md) |  | [optional] 

## Methods

### NewPasswordPolicyRuleConditionsAllOf

`func NewPasswordPolicyRuleConditionsAllOf() *PasswordPolicyRuleConditionsAllOf`

NewPasswordPolicyRuleConditionsAllOf instantiates a new PasswordPolicyRuleConditionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyRuleConditionsAllOfWithDefaults

`func NewPasswordPolicyRuleConditionsAllOfWithDefaults() *PasswordPolicyRuleConditionsAllOf`

NewPasswordPolicyRuleConditionsAllOfWithDefaults instantiates a new PasswordPolicyRuleConditionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *PasswordPolicyRuleConditionsAllOf) GetNetwork() PolicyNetworkCondition`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *PasswordPolicyRuleConditionsAllOf) GetNetworkOk() (*PolicyNetworkCondition, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *PasswordPolicyRuleConditionsAllOf) SetNetwork(v PolicyNetworkCondition)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *PasswordPolicyRuleConditionsAllOf) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPeople

`func (o *PasswordPolicyRuleConditionsAllOf) GetPeople() PolicyPeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *PasswordPolicyRuleConditionsAllOf) GetPeopleOk() (*PolicyPeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *PasswordPolicyRuleConditionsAllOf) SetPeople(v PolicyPeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *PasswordPolicyRuleConditionsAllOf) HasPeople() bool`

HasPeople returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


