# ProfileEnrollmentPolicyRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | Indicates if the user profile is granted access  &gt; **Note:** You can&#39;t set the &#x60;access&#x60; property to &#x60;DENY&#x60; after you create the policy | [optional] 
**ActivationRequirements** | Pointer to [**ProfileEnrollmentPolicyRuleActivationRequirement**](ProfileEnrollmentPolicyRuleActivationRequirement.md) |  | [optional] 
**AllowedIdentifiers** | Pointer to **[]string** | A list of attributes to identify an end user. Can be used across Okta sign-in, unlock, and recovery flows. | [optional] [default to ["login"]]
**EnrollAuthenticatorTypes** | Pointer to **[]string** | Additional authenticator fields that can be used on the first page of user registration. Valid values only includes &#x60;&#39;password&#39;&#x60;. | [optional] 
**PreRegistrationInlineHooks** | Pointer to [**[]PreRegistrationInlineHook**](PreRegistrationInlineHook.md) | (Optional) The &#x60;id&#x60; of at most one registration inline hook | [optional] 
**ProfileAttributes** | Pointer to [**[]ProfileEnrollmentPolicyRuleProfileAttribute**](ProfileEnrollmentPolicyRuleProfileAttribute.md) | A list of attributes to prompt the user for during registration or progressive profiling. Where defined on the user schema, these attributes are persisted in the user profile. You can also add non-schema attributes, which aren&#39;t persisted to the user&#39;s profile, but are included in requests to the registration inline hook. A maximum of 10 profile properties is supported. | [optional] 
**ProgressiveProfilingAction** | Pointer to **string** | Progressive profile enrollment helps evaluate the user profile policy at every user login. Users can be prompted to provide input for newly required attributes. | [optional] 
**TargetGroupIds** | Pointer to **[]string** | (Optional, max 1 entry) The &#x60;id&#x60; of a group that this user should be added to | [optional] 
**UiSchemaId** | Pointer to **string** | Value created by the backend. If present, all policy updates must include this attribute/value. | [optional] 
**UnknownUserAction** | Pointer to **string** | Which action should be taken if this user is new | [optional] 

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

### GetAllowedIdentifiers

`func (o *ProfileEnrollmentPolicyRuleAction) GetAllowedIdentifiers() []string`

GetAllowedIdentifiers returns the AllowedIdentifiers field if non-nil, zero value otherwise.

### GetAllowedIdentifiersOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetAllowedIdentifiersOk() (*[]string, bool)`

GetAllowedIdentifiersOk returns a tuple with the AllowedIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIdentifiers

`func (o *ProfileEnrollmentPolicyRuleAction) SetAllowedIdentifiers(v []string)`

SetAllowedIdentifiers sets AllowedIdentifiers field to given value.

### HasAllowedIdentifiers

`func (o *ProfileEnrollmentPolicyRuleAction) HasAllowedIdentifiers() bool`

HasAllowedIdentifiers returns a boolean if a field has been set.

### GetEnrollAuthenticatorTypes

`func (o *ProfileEnrollmentPolicyRuleAction) GetEnrollAuthenticatorTypes() []string`

GetEnrollAuthenticatorTypes returns the EnrollAuthenticatorTypes field if non-nil, zero value otherwise.

### GetEnrollAuthenticatorTypesOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetEnrollAuthenticatorTypesOk() (*[]string, bool)`

GetEnrollAuthenticatorTypesOk returns a tuple with the EnrollAuthenticatorTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollAuthenticatorTypes

`func (o *ProfileEnrollmentPolicyRuleAction) SetEnrollAuthenticatorTypes(v []string)`

SetEnrollAuthenticatorTypes sets EnrollAuthenticatorTypes field to given value.

### HasEnrollAuthenticatorTypes

`func (o *ProfileEnrollmentPolicyRuleAction) HasEnrollAuthenticatorTypes() bool`

HasEnrollAuthenticatorTypes returns a boolean if a field has been set.

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

### GetProgressiveProfilingAction

`func (o *ProfileEnrollmentPolicyRuleAction) GetProgressiveProfilingAction() string`

GetProgressiveProfilingAction returns the ProgressiveProfilingAction field if non-nil, zero value otherwise.

### GetProgressiveProfilingActionOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetProgressiveProfilingActionOk() (*string, bool)`

GetProgressiveProfilingActionOk returns a tuple with the ProgressiveProfilingAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressiveProfilingAction

`func (o *ProfileEnrollmentPolicyRuleAction) SetProgressiveProfilingAction(v string)`

SetProgressiveProfilingAction sets ProgressiveProfilingAction field to given value.

### HasProgressiveProfilingAction

`func (o *ProfileEnrollmentPolicyRuleAction) HasProgressiveProfilingAction() bool`

HasProgressiveProfilingAction returns a boolean if a field has been set.

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

### GetUiSchemaId

`func (o *ProfileEnrollmentPolicyRuleAction) GetUiSchemaId() string`

GetUiSchemaId returns the UiSchemaId field if non-nil, zero value otherwise.

### GetUiSchemaIdOk

`func (o *ProfileEnrollmentPolicyRuleAction) GetUiSchemaIdOk() (*string, bool)`

GetUiSchemaIdOk returns a tuple with the UiSchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiSchemaId

`func (o *ProfileEnrollmentPolicyRuleAction) SetUiSchemaId(v string)`

SetUiSchemaId sets UiSchemaId field to given value.

### HasUiSchemaId

`func (o *ProfileEnrollmentPolicyRuleAction) HasUiSchemaId() bool`

HasUiSchemaId returns a boolean if a field has been set.

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


