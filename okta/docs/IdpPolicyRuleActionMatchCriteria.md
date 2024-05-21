# IdpPolicyRuleActionMatchCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** | The IdP property that the evaluated string should match to | [optional] 
**ProviderExpression** | Pointer to **string** | You can provide an Okta Expression Language expression with the Login Context that&#39;s evaluated with the IdP. For example, the value &#x60;login.identifier&#x60; refers to the user&#39;s username. If the user is signing in with the username &#x60;john.doe@mycompany.com&#x60;, the expression &#x60;login.identifier.substringAfter(@))&#x60; is evaluated to the domain name of the user, for example: &#x60;mycompany.com&#x60;.  | [optional] 

## Methods

### NewIdpPolicyRuleActionMatchCriteria

`func NewIdpPolicyRuleActionMatchCriteria() *IdpPolicyRuleActionMatchCriteria`

NewIdpPolicyRuleActionMatchCriteria instantiates a new IdpPolicyRuleActionMatchCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPolicyRuleActionMatchCriteriaWithDefaults

`func NewIdpPolicyRuleActionMatchCriteriaWithDefaults() *IdpPolicyRuleActionMatchCriteria`

NewIdpPolicyRuleActionMatchCriteriaWithDefaults instantiates a new IdpPolicyRuleActionMatchCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *IdpPolicyRuleActionMatchCriteria) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *IdpPolicyRuleActionMatchCriteria) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *IdpPolicyRuleActionMatchCriteria) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *IdpPolicyRuleActionMatchCriteria) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetProviderExpression

`func (o *IdpPolicyRuleActionMatchCriteria) GetProviderExpression() string`

GetProviderExpression returns the ProviderExpression field if non-nil, zero value otherwise.

### GetProviderExpressionOk

`func (o *IdpPolicyRuleActionMatchCriteria) GetProviderExpressionOk() (*string, bool)`

GetProviderExpressionOk returns a tuple with the ProviderExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderExpression

`func (o *IdpPolicyRuleActionMatchCriteria) SetProviderExpression(v string)`

SetProviderExpression sets ProviderExpression field to given value.

### HasProviderExpression

`func (o *IdpPolicyRuleActionMatchCriteria) HasProviderExpression() bool`

HasProviderExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


