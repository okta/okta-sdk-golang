# IdentityProviderPolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpIds** | Pointer to **[]string** |  | [optional] 
**Provider** | Pointer to [**IdentityProviderPolicyProvider**](IdentityProviderPolicyProvider.md) |  | [optional] 

## Methods

### NewIdentityProviderPolicyRuleCondition

`func NewIdentityProviderPolicyRuleCondition() *IdentityProviderPolicyRuleCondition`

NewIdentityProviderPolicyRuleCondition instantiates a new IdentityProviderPolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderPolicyRuleConditionWithDefaults

`func NewIdentityProviderPolicyRuleConditionWithDefaults() *IdentityProviderPolicyRuleCondition`

NewIdentityProviderPolicyRuleConditionWithDefaults instantiates a new IdentityProviderPolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpIds

`func (o *IdentityProviderPolicyRuleCondition) GetIdpIds() []string`

GetIdpIds returns the IdpIds field if non-nil, zero value otherwise.

### GetIdpIdsOk

`func (o *IdentityProviderPolicyRuleCondition) GetIdpIdsOk() (*[]string, bool)`

GetIdpIdsOk returns a tuple with the IdpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpIds

`func (o *IdentityProviderPolicyRuleCondition) SetIdpIds(v []string)`

SetIdpIds sets IdpIds field to given value.

### HasIdpIds

`func (o *IdentityProviderPolicyRuleCondition) HasIdpIds() bool`

HasIdpIds returns a boolean if a field has been set.

### GetProvider

`func (o *IdentityProviderPolicyRuleCondition) GetProvider() IdentityProviderPolicyProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IdentityProviderPolicyRuleCondition) GetProviderOk() (*IdentityProviderPolicyProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IdentityProviderPolicyRuleCondition) SetProvider(v IdentityProviderPolicyProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IdentityProviderPolicyRuleCondition) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


