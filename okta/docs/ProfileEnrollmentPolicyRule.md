# ProfileEnrollmentPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**ProfileEnrollmentPolicyRuleActions**](ProfileEnrollmentPolicyRuleActions.md) |  | [optional] 
**Conditions** | Pointer to [**PolicyRuleConditions**](PolicyRuleConditions.md) |  | [optional] 

## Methods

### NewProfileEnrollmentPolicyRule

`func NewProfileEnrollmentPolicyRule() *ProfileEnrollmentPolicyRule`

NewProfileEnrollmentPolicyRule instantiates a new ProfileEnrollmentPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileEnrollmentPolicyRuleWithDefaults

`func NewProfileEnrollmentPolicyRuleWithDefaults() *ProfileEnrollmentPolicyRule`

NewProfileEnrollmentPolicyRuleWithDefaults instantiates a new ProfileEnrollmentPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ProfileEnrollmentPolicyRule) GetActions() ProfileEnrollmentPolicyRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ProfileEnrollmentPolicyRule) GetActionsOk() (*ProfileEnrollmentPolicyRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ProfileEnrollmentPolicyRule) SetActions(v ProfileEnrollmentPolicyRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ProfileEnrollmentPolicyRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *ProfileEnrollmentPolicyRule) GetConditions() PolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ProfileEnrollmentPolicyRule) GetConditionsOk() (*PolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ProfileEnrollmentPolicyRule) SetConditions(v PolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ProfileEnrollmentPolicyRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


