# AccessPolicyRuleCustomCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | **string** | The Okta Expression Language expression to evaluate.   &gt; &lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; **Note:** When the Flexible Okta Verify authenticator configuration feature is enabled, &#x60;okta_verify&#x60; in expressions such as &#x60;accessRequest.authenticator.key &#x3D;&#x3D; &#39;okta_verify&#39;&#x60; is no longer accepted. Use &#x60;okta_verify_totp&#x60;, &#x60;okta_verify_push&#x60;, or &#x60;okta_verify_fastpass&#x60; instead. | 

## Methods

### NewAccessPolicyRuleCustomCondition

`func NewAccessPolicyRuleCustomCondition(condition string, ) *AccessPolicyRuleCustomCondition`

NewAccessPolicyRuleCustomCondition instantiates a new AccessPolicyRuleCustomCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyRuleCustomConditionWithDefaults

`func NewAccessPolicyRuleCustomConditionWithDefaults() *AccessPolicyRuleCustomCondition`

NewAccessPolicyRuleCustomConditionWithDefaults instantiates a new AccessPolicyRuleCustomCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *AccessPolicyRuleCustomCondition) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AccessPolicyRuleCustomCondition) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AccessPolicyRuleCustomCondition) SetCondition(v string)`

SetCondition sets Condition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


