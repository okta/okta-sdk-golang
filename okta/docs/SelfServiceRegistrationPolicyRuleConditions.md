# SelfServiceRegistrationPolicyRuleConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailDomains** | Pointer to [**SelfServiceRegistrationEmailDomainsCondition**](SelfServiceRegistrationEmailDomainsCondition.md) |  | [optional] 
**Network** | Pointer to [**PolicyNetworkCondition**](PolicyNetworkCondition.md) |  | [optional] 

## Methods

### NewSelfServiceRegistrationPolicyRuleConditions

`func NewSelfServiceRegistrationPolicyRuleConditions() *SelfServiceRegistrationPolicyRuleConditions`

NewSelfServiceRegistrationPolicyRuleConditions instantiates a new SelfServiceRegistrationPolicyRuleConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServiceRegistrationPolicyRuleConditionsWithDefaults

`func NewSelfServiceRegistrationPolicyRuleConditionsWithDefaults() *SelfServiceRegistrationPolicyRuleConditions`

NewSelfServiceRegistrationPolicyRuleConditionsWithDefaults instantiates a new SelfServiceRegistrationPolicyRuleConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailDomains

`func (o *SelfServiceRegistrationPolicyRuleConditions) GetEmailDomains() SelfServiceRegistrationEmailDomainsCondition`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *SelfServiceRegistrationPolicyRuleConditions) GetEmailDomainsOk() (*SelfServiceRegistrationEmailDomainsCondition, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *SelfServiceRegistrationPolicyRuleConditions) SetEmailDomains(v SelfServiceRegistrationEmailDomainsCondition)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *SelfServiceRegistrationPolicyRuleConditions) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### GetNetwork

`func (o *SelfServiceRegistrationPolicyRuleConditions) GetNetwork() PolicyNetworkCondition`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *SelfServiceRegistrationPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *SelfServiceRegistrationPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *SelfServiceRegistrationPolicyRuleConditions) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


