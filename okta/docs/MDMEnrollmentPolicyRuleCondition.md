# MDMEnrollmentPolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockNonSafeAndroid** | Pointer to **bool** |  | [optional] 
**Enrollment** | Pointer to [**MDMEnrollmentPolicyEnrollment**](MDMEnrollmentPolicyEnrollment.md) |  | [optional] 

## Methods

### NewMDMEnrollmentPolicyRuleCondition

`func NewMDMEnrollmentPolicyRuleCondition() *MDMEnrollmentPolicyRuleCondition`

NewMDMEnrollmentPolicyRuleCondition instantiates a new MDMEnrollmentPolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMDMEnrollmentPolicyRuleConditionWithDefaults

`func NewMDMEnrollmentPolicyRuleConditionWithDefaults() *MDMEnrollmentPolicyRuleCondition`

NewMDMEnrollmentPolicyRuleConditionWithDefaults instantiates a new MDMEnrollmentPolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockNonSafeAndroid

`func (o *MDMEnrollmentPolicyRuleCondition) GetBlockNonSafeAndroid() bool`

GetBlockNonSafeAndroid returns the BlockNonSafeAndroid field if non-nil, zero value otherwise.

### GetBlockNonSafeAndroidOk

`func (o *MDMEnrollmentPolicyRuleCondition) GetBlockNonSafeAndroidOk() (*bool, bool)`

GetBlockNonSafeAndroidOk returns a tuple with the BlockNonSafeAndroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNonSafeAndroid

`func (o *MDMEnrollmentPolicyRuleCondition) SetBlockNonSafeAndroid(v bool)`

SetBlockNonSafeAndroid sets BlockNonSafeAndroid field to given value.

### HasBlockNonSafeAndroid

`func (o *MDMEnrollmentPolicyRuleCondition) HasBlockNonSafeAndroid() bool`

HasBlockNonSafeAndroid returns a boolean if a field has been set.

### GetEnrollment

`func (o *MDMEnrollmentPolicyRuleCondition) GetEnrollment() MDMEnrollmentPolicyEnrollment`

GetEnrollment returns the Enrollment field if non-nil, zero value otherwise.

### GetEnrollmentOk

`func (o *MDMEnrollmentPolicyRuleCondition) GetEnrollmentOk() (*MDMEnrollmentPolicyEnrollment, bool)`

GetEnrollmentOk returns a tuple with the Enrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollment

`func (o *MDMEnrollmentPolicyRuleCondition) SetEnrollment(v MDMEnrollmentPolicyEnrollment)`

SetEnrollment sets Enrollment field to given value.

### HasEnrollment

`func (o *MDMEnrollmentPolicyRuleCondition) HasEnrollment() bool`

HasEnrollment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


