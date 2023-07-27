# AuthorizationServerPolicyRuleConditionsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**ClientPolicyCondition**](ClientPolicyCondition.md) |  | [optional] 
**GrantTypes** | Pointer to [**GrantTypePolicyRuleCondition**](GrantTypePolicyRuleCondition.md) |  | [optional] 
**People** | Pointer to [**PolicyPeopleCondition**](PolicyPeopleCondition.md) |  | [optional] 
**Scopes** | Pointer to [**OAuth2ScopesMediationPolicyRuleCondition**](OAuth2ScopesMediationPolicyRuleCondition.md) |  | [optional] 

## Methods

### NewAuthorizationServerPolicyRuleConditionsAllOf

`func NewAuthorizationServerPolicyRuleConditionsAllOf() *AuthorizationServerPolicyRuleConditionsAllOf`

NewAuthorizationServerPolicyRuleConditionsAllOf instantiates a new AuthorizationServerPolicyRuleConditionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyRuleConditionsAllOfWithDefaults

`func NewAuthorizationServerPolicyRuleConditionsAllOfWithDefaults() *AuthorizationServerPolicyRuleConditionsAllOf`

NewAuthorizationServerPolicyRuleConditionsAllOfWithDefaults instantiates a new AuthorizationServerPolicyRuleConditionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) GetClients() ClientPolicyCondition`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) GetClientsOk() (*ClientPolicyCondition, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) SetClients(v ClientPolicyCondition)`

SetClients sets Clients field to given value.

### HasClients

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetGrantTypes

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) GetGrantTypes() GrantTypePolicyRuleCondition`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) GetGrantTypesOk() (*GrantTypePolicyRuleCondition, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) SetGrantTypes(v GrantTypePolicyRuleCondition)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetPeople

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) GetPeople() PolicyPeopleCondition`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) GetPeopleOk() (*PolicyPeopleCondition, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) SetPeople(v PolicyPeopleCondition)`

SetPeople sets People field to given value.

### HasPeople

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetScopes

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) GetScopes() OAuth2ScopesMediationPolicyRuleCondition`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) GetScopesOk() (*OAuth2ScopesMediationPolicyRuleCondition, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) SetScopes(v OAuth2ScopesMediationPolicyRuleCondition)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AuthorizationServerPolicyRuleConditionsAllOf) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


