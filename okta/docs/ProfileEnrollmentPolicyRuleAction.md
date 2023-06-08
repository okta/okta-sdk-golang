# ProfileEnrollmentPolicyRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** |  | [optional] 
**ActivationRequirements** | Pointer to [**ProfileEnrollmentPolicyRuleActivationRequirement**](ProfileEnrollmentPolicyRuleActivationRequirement.md) |  | [optional] 
**PreRegistrationInlineHooks** | Pointer to [**[]PreRegistrationInlineHook**](PreRegistrationInlineHook.md) |  | [optional] 
**ProfileAttributes** | Pointer to [**[]ProfileEnrollmentPolicyRuleProfileAttribute**](ProfileEnrollmentPolicyRuleProfileAttribute.md) |  | [optional] 
**TargetGroupIds** | Pointer to **[]string** |  | [optional] 
**UnknownUserAction** | Pointer to **string** |  | [optional] 

## Methods

### NewProfileEnrollmentPolicyRuleAction

`func NewProfileEnrollmentPolicyRuleAction() *ProfileEnrollmentPolicyRuleAction`

NewProfileEnrollmentPolicyRuleAction instantiates a new ProfileEnrollmentPolicyRuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileEnrollmentPolicyRuleActionWithDefaults

`func NewProfileEnrollmentPolicyRuleActionWithDefaults() *ProfileEnrollmentPolicyRuleAction`

NewProfileEnrollmentPolicyRuleActionWithDefaults instantiates a new ProfileEnrollmentPolicyRuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ProfileEnrollmentPolicyRuleAction) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ProfileEnrollmentPolicyRuleAction) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ProfileEnrollmentPolicyRuleAction) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetActivationRequirements

`func (o *ProfileEnrollmentPolicyRuleAction) GetActivationRequirements() ProfileEnrollmentPolicyRuleActivationRequirement`

GetActivationRequirements returns the ActivationRequirements field if non-nil, zero value otherwise.

### GetActivationRequirementsOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetActivationRequirementsOk() (*ProfileEnrollmentPolicyRuleActivationRequirement, bool)`

GetActivationRequirementsOk returns a tuple with the ActivationRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationRequirements

`func (o *ProfileEnrollmentPolicyRuleAction) SetActivationRequirements(v ProfileEnrollmentPolicyRuleActivationRequirement)`

SetActivationRequirements sets ActivationRequirements field to given value.

### HasActivationRequirements

`func (o *ProfileEnrollmentPolicyRuleAction) HasActivationRequirements() bool`

HasActivationRequirements returns a boolean if a field has been set.

### GetPreRegistrationInlineHooks

`func (o *ProfileEnrollmentPolicyRuleAction) GetPreRegistrationInlineHooks() []PreRegistrationInlineHook`

GetPreRegistrationInlineHooks returns the PreRegistrationInlineHooks field if non-nil, zero value otherwise.

### GetPreRegistrationInlineHooksOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetPreRegistrationInlineHooksOk() (*[]PreRegistrationInlineHook, bool)`

GetPreRegistrationInlineHooksOk returns a tuple with the PreRegistrationInlineHooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRegistrationInlineHooks

`func (o *ProfileEnrollmentPolicyRuleAction) SetPreRegistrationInlineHooks(v []PreRegistrationInlineHook)`

SetPreRegistrationInlineHooks sets PreRegistrationInlineHooks field to given value.

### HasPreRegistrationInlineHooks

`func (o *ProfileEnrollmentPolicyRuleAction) HasPreRegistrationInlineHooks() bool`

HasPreRegistrationInlineHooks returns a boolean if a field has been set.

### GetProfileAttributes

`func (o *ProfileEnrollmentPolicyRuleAction) GetProfileAttributes() []ProfileEnrollmentPolicyRuleProfileAttribute`

GetProfileAttributes returns the ProfileAttributes field if non-nil, zero value otherwise.

### GetProfileAttributesOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetProfileAttributesOk() (*[]ProfileEnrollmentPolicyRuleProfileAttribute, bool)`

GetProfileAttributesOk returns a tuple with the ProfileAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAttributes

`func (o *ProfileEnrollmentPolicyRuleAction) SetProfileAttributes(v []ProfileEnrollmentPolicyRuleProfileAttribute)`

SetProfileAttributes sets ProfileAttributes field to given value.

### HasProfileAttributes

`func (o *ProfileEnrollmentPolicyRuleAction) HasProfileAttributes() bool`

HasProfileAttributes returns a boolean if a field has been set.

### GetTargetGroupIds

`func (o *ProfileEnrollmentPolicyRuleAction) GetTargetGroupIds() []string`

GetTargetGroupIds returns the TargetGroupIds field if non-nil, zero value otherwise.

### GetTargetGroupIdsOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetTargetGroupIdsOk() (*[]string, bool)`

GetTargetGroupIdsOk returns a tuple with the TargetGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupIds

`func (o *ProfileEnrollmentPolicyRuleAction) SetTargetGroupIds(v []string)`

SetTargetGroupIds sets TargetGroupIds field to given value.

### HasTargetGroupIds

`func (o *ProfileEnrollmentPolicyRuleAction) HasTargetGroupIds() bool`

HasTargetGroupIds returns a boolean if a field has been set.

### GetUnknownUserAction

`func (o *ProfileEnrollmentPolicyRuleAction) GetUnknownUserAction() string`

GetUnknownUserAction returns the UnknownUserAction field if non-nil, zero value otherwise.

### GetUnknownUserActionOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetUnknownUserActionOk() (*string, bool)`

GetUnknownUserActionOk returns a tuple with the UnknownUserAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUserAction

`func (o *ProfileEnrollmentPolicyRuleAction) SetUnknownUserAction(v string)`

SetUnknownUserAction sets UnknownUserAction field to given value.

### HasUnknownUserAction

`func (o *ProfileEnrollmentPolicyRuleAction) HasUnknownUserAction() bool`

HasUnknownUserAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


