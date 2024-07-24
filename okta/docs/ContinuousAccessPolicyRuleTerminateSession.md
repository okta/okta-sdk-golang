# ContinuousAccessPolicyRuleTerminateSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action to take when Continuous Access evaluation detects a failure. | [optional] 
**Slo** | Pointer to [**ContinuousAccessPolicyRuleTerminateSessionSlo**](ContinuousAccessPolicyRuleTerminateSessionSlo.md) |  | [optional] 

## Methods

### NewContinuousAccessPolicyRuleTerminateSession

`func NewContinuousAccessPolicyRuleTerminateSession() *ContinuousAccessPolicyRuleTerminateSession`

NewContinuousAccessPolicyRuleTerminateSession instantiates a new ContinuousAccessPolicyRuleTerminateSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuousAccessPolicyRuleTerminateSessionWithDefaults

`func NewContinuousAccessPolicyRuleTerminateSessionWithDefaults() *ContinuousAccessPolicyRuleTerminateSession`

NewContinuousAccessPolicyRuleTerminateSessionWithDefaults instantiates a new ContinuousAccessPolicyRuleTerminateSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ContinuousAccessPolicyRuleTerminateSession) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ContinuousAccessPolicyRuleTerminateSession) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ContinuousAccessPolicyRuleTerminateSession) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ContinuousAccessPolicyRuleTerminateSession) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSlo

`func (o *ContinuousAccessPolicyRuleTerminateSession) GetSlo() ContinuousAccessPolicyRuleTerminateSessionSlo`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *ContinuousAccessPolicyRuleTerminateSession) GetSloOk() (*ContinuousAccessPolicyRuleTerminateSessionSlo, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *ContinuousAccessPolicyRuleTerminateSession) SetSlo(v ContinuousAccessPolicyRuleTerminateSessionSlo)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *ContinuousAccessPolicyRuleTerminateSession) HasSlo() bool`

HasSlo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


