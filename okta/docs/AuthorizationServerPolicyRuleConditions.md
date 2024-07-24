# AuthorizationServerPolicyRuleConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantTypes** | Pointer to [**GrantTypePolicyRuleCondition**](GrantTypePolicyRuleCondition.md) |  | [optional] 
**People** | Pointer to [**AuthorizationServerPolicyPeopleCondition**](AuthorizationServerPolicyPeopleCondition.md) |  | [optional] 
**Scopes** | Pointer to [**OAuth2ScopesMediationPolicyRuleCondition**](OAuth2ScopesMediationPolicyRuleCondition.md) |  | [optional] 

## Methods

### NewAuthorizationServerPolicyRuleConditions

`func NewAuthorizationServerPolicyRuleConditions() *AuthorizationServerPolicyRuleConditions`

NewAuthorizationServerPolicyRuleConditions instantiates a new AuthorizationServerPolicyRuleConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyRuleConditionsWithDefaults

`func NewAuthorizationServerPolicyRuleConditionsWithDefaults() *AuthorizationServerPolicyRuleConditions`

NewAuthorizationServerPolicyRuleConditionsWithDefaults instantiates a new AuthorizationServerPolicyRuleConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantTypes

`func (o *AuthorizationServerPolicyRuleConditions) GetGrantTypes() GrantTypePolicyRuleCondition`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *AuthorizationServerPolicyRuleConditions) GetGrantTypesOk() (*GrantTypePolicyRuleCondition, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *AuthorizationServerPolicyRuleConditions) SetGrantTypes(v GrantTypePolicyRuleCondition)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *AuthorizationServerPolicyRuleConditions) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetPeople

`func (o *AuthorizationServerPolicyRuleConditions) GetPeople() AuthorizationServerPolicyPeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *AuthorizationServerPolicyRuleConditions) GetPeopleOk() (*AuthorizationServerPolicyPeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *AuthorizationServerPolicyRuleConditions) SetPeople(v AuthorizationServerPolicyPeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *AuthorizationServerPolicyRuleConditions) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetScopes

`func (o *AuthorizationServerPolicyRuleConditions) GetScopes() OAuth2ScopesMediationPolicyRuleCondition`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizationServerPolicyRuleConditions) GetScopesOk() (*OAuth2ScopesMediationPolicyRuleCondition, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizationServerPolicyRuleConditions) SetScopes(v OAuth2ScopesMediationPolicyRuleCondition)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizationServerPolicyRuleConditions) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


