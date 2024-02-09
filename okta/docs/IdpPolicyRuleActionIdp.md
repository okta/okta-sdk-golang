# IdpPolicyRuleActionIdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]IdpPolicyRuleActionProvider**](IdpPolicyRuleActionProvider.md) | List of configured Identity Providers that a given Rule can route to. Ability to define multiple providers is a part of the Okta Identity Engine. This allows users to choose a Provider when they sign in. Contact support for information on the Identity Engine. | [optional] 
**IdpSelectionType** | Pointer to **string** |  | [optional] 
**MatchCriteria** | Pointer to [**[]IdpPolicyRuleActionMatchCriteria**](IdpPolicyRuleActionMatchCriteria.md) | Required if &#x60;idpSelectionType&#x60; is set to &#x60;DYNAMIC&#x60; | [optional] 

## Methods

### NewIdpPolicyRuleActionIdp

`func NewIdpPolicyRuleActionIdp() *IdpPolicyRuleActionIdp`

NewIdpPolicyRuleActionIdp instantiates a new IdpPolicyRuleActionIdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpPolicyRuleActionIdpWithDefaults

`func NewIdpPolicyRuleActionIdpWithDefaults() *IdpPolicyRuleActionIdp`

NewIdpPolicyRuleActionIdpWithDefaults instantiates a new IdpPolicyRuleActionIdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *IdpPolicyRuleActionIdp) GetProviders() []IdpPolicyRuleActionProvider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *IdpPolicyRuleActionIdp) GetProvidersOk() (*[]IdpPolicyRuleActionProvider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *IdpPolicyRuleActionIdp) SetProviders(v []IdpPolicyRuleActionProvider)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *IdpPolicyRuleActionIdp) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetIdpSelectionType

`func (o *IdpPolicyRuleActionIdp) GetIdpSelectionType() string`

GetIdpSelectionType returns the IdpSelectionType field if non-nil, zero value otherwise.

### GetIdpSelectionTypeOk

`func (o *IdpPolicyRuleActionIdp) GetIdpSelectionTypeOk() (*string, bool)`

GetIdpSelectionTypeOk returns a tuple with the IdpSelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSelectionType

`func (o *IdpPolicyRuleActionIdp) SetIdpSelectionType(v string)`

SetIdpSelectionType sets IdpSelectionType field to given value.

### HasIdpSelectionType

`func (o *IdpPolicyRuleActionIdp) HasIdpSelectionType() bool`

HasIdpSelectionType returns a boolean if a field has been set.

### GetMatchCriteria

`func (o *IdpPolicyRuleActionIdp) GetMatchCriteria() []IdpPolicyRuleActionMatchCriteria`

GetMatchCriteria returns the MatchCriteria field if non-nil, zero value otherwise.

### GetMatchCriteriaOk

`func (o *IdpPolicyRuleActionIdp) GetMatchCriteriaOk() (*[]IdpPolicyRuleActionMatchCriteria, bool)`

GetMatchCriteriaOk returns a tuple with the MatchCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCriteria

`func (o *IdpPolicyRuleActionIdp) SetMatchCriteria(v []IdpPolicyRuleActionMatchCriteria)`

SetMatchCriteria sets MatchCriteria field to given value.

### HasMatchCriteria

`func (o *IdpPolicyRuleActionIdp) HasMatchCriteria() bool`

HasMatchCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


