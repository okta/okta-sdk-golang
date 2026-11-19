# AuthorizationServerPolicyRuleXaaIdpCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpIds** | Pointer to **[]string** | IDs of the IdPs that are allowed. It&#39;s required when &#x60;provider&#x60; is set as &#x60;SPECIFIC_IDP&#x60;. | [optional] 
**Provider** | Pointer to **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Specifies which IdPs that the rule applies to for XAA | [optional] 

## Methods

### NewAuthorizationServerPolicyRuleXaaIdpCondition

`func NewAuthorizationServerPolicyRuleXaaIdpCondition() *AuthorizationServerPolicyRuleXaaIdpCondition`

NewAuthorizationServerPolicyRuleXaaIdpCondition instantiates a new AuthorizationServerPolicyRuleXaaIdpCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyRuleXaaIdpConditionWithDefaults

`func NewAuthorizationServerPolicyRuleXaaIdpConditionWithDefaults() *AuthorizationServerPolicyRuleXaaIdpCondition`

NewAuthorizationServerPolicyRuleXaaIdpConditionWithDefaults instantiates a new AuthorizationServerPolicyRuleXaaIdpCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpIds

`func (o *AuthorizationServerPolicyRuleXaaIdpCondition) GetIdpIds() []string`

GetIdpIds returns the IdpIds field if non-nil, zero value otherwise.

### GetIdpIdsOk

`func (o *AuthorizationServerPolicyRuleXaaIdpCondition) GetIdpIdsOk() (*[]string, bool)`

GetIdpIdsOk returns a tuple with the IdpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpIds

`func (o *AuthorizationServerPolicyRuleXaaIdpCondition) SetIdpIds(v []string)`

SetIdpIds sets IdpIds field to given value.

### HasIdpIds

`func (o *AuthorizationServerPolicyRuleXaaIdpCondition) HasIdpIds() bool`

HasIdpIds returns a boolean if a field has been set.

### GetProvider

`func (o *AuthorizationServerPolicyRuleXaaIdpCondition) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AuthorizationServerPolicyRuleXaaIdpCondition) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AuthorizationServerPolicyRuleXaaIdpCondition) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AuthorizationServerPolicyRuleXaaIdpCondition) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


