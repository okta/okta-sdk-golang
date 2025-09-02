/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// ProfileEnrollmentPolicyRuleAction struct for ProfileEnrollmentPolicyRuleAction
type ProfileEnrollmentPolicyRuleAction struct {
	// Indicates if the user profile is granted access  > **Note:** You can't set the `access` property to `DENY` after you create the policy
	Access *string `json:"access,omitempty"`
	ActivationRequirements *ProfileEnrollmentPolicyRuleActivationRequirement `json:"activationRequirements,omitempty"`
	// A list of attributes to identify an end user. Can be used across Okta sign-in, unlock, and recovery flows.
	AllowedIdentifiers []string `json:"allowedIdentifiers,omitempty"`
	// Additional authenticator fields that can be used on the first page of user registration. Valid values only includes `'password'`.
	EnrollAuthenticatorTypes []string `json:"enrollAuthenticatorTypes,omitempty"`
	// (Optional) The `id` of at most one registration inline hook
	PreRegistrationInlineHooks []PreRegistrationInlineHook `json:"preRegistrationInlineHooks,omitempty"`
	// A list of attributes to prompt the user for during registration or progressive profiling. Where defined on the user schema, these attributes are persisted in the user profile. You can also add non-schema attributes, which aren't persisted to the user's profile, but are included in requests to the registration inline hook. A maximum of 10 profile properties is supported.
	ProfileAttributes []ProfileEnrollmentPolicyRuleProfileAttribute `json:"profileAttributes,omitempty"`
	// Progressive profile enrollment helps evaluate the user profile policy at every user login. Users can be prompted to provide input for newly required attributes.
	ProgressiveProfilingAction *string `json:"progressiveProfilingAction,omitempty"`
	// (Optional, max 1 entry) The `id` of a group that this user should be added to
	TargetGroupIds []string `json:"targetGroupIds,omitempty"`
	// Value created by the backend. If present, all policy updates must include this attribute/value.
	UiSchemaId *string `json:"uiSchemaId,omitempty"`
	// Which action should be taken if this user is new
	UnknownUserAction *string `json:"unknownUserAction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileEnrollmentPolicyRuleAction ProfileEnrollmentPolicyRuleAction

// NewProfileEnrollmentPolicyRuleAction instantiates a new ProfileEnrollmentPolicyRuleAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileEnrollmentPolicyRuleAction() *ProfileEnrollmentPolicyRuleAction {
	this := ProfileEnrollmentPolicyRuleAction{}
	return &this
}

// NewProfileEnrollmentPolicyRuleActionWithDefaults instantiates a new ProfileEnrollmentPolicyRuleAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileEnrollmentPolicyRuleActionWithDefaults() *ProfileEnrollmentPolicyRuleAction {
	this := ProfileEnrollmentPolicyRuleAction{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetAccess() string {
	if o == nil || o.Access == nil {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetAccessOk() (*string, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *ProfileEnrollmentPolicyRuleAction) SetAccess(v string) {
	o.Access = &v
}

// GetActivationRequirements returns the ActivationRequirements field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetActivationRequirements() ProfileEnrollmentPolicyRuleActivationRequirement {
	if o == nil || o.ActivationRequirements == nil {
		var ret ProfileEnrollmentPolicyRuleActivationRequirement
		return ret
	}
	return *o.ActivationRequirements
}

// GetActivationRequirementsOk returns a tuple with the ActivationRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetActivationRequirementsOk() (*ProfileEnrollmentPolicyRuleActivationRequirement, bool) {
	if o == nil || o.ActivationRequirements == nil {
		return nil, false
	}
	return o.ActivationRequirements, true
}

// HasActivationRequirements returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasActivationRequirements() bool {
	if o != nil && o.ActivationRequirements != nil {
		return true
	}

	return false
}

// SetActivationRequirements gets a reference to the given ProfileEnrollmentPolicyRuleActivationRequirement and assigns it to the ActivationRequirements field.
func (o *ProfileEnrollmentPolicyRuleAction) SetActivationRequirements(v ProfileEnrollmentPolicyRuleActivationRequirement) {
	o.ActivationRequirements = &v
}

// GetAllowedIdentifiers returns the AllowedIdentifiers field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetAllowedIdentifiers() []string {
	if o == nil || o.AllowedIdentifiers == nil {
		var ret []string
		return ret
	}
	return o.AllowedIdentifiers
}

// GetAllowedIdentifiersOk returns a tuple with the AllowedIdentifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetAllowedIdentifiersOk() ([]string, bool) {
	if o == nil || o.AllowedIdentifiers == nil {
		return nil, false
	}
	return o.AllowedIdentifiers, true
}

// HasAllowedIdentifiers returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasAllowedIdentifiers() bool {
	if o != nil && o.AllowedIdentifiers != nil {
		return true
	}

	return false
}

// SetAllowedIdentifiers gets a reference to the given []string and assigns it to the AllowedIdentifiers field.
func (o *ProfileEnrollmentPolicyRuleAction) SetAllowedIdentifiers(v []string) {
	o.AllowedIdentifiers = v
}

// GetEnrollAuthenticatorTypes returns the EnrollAuthenticatorTypes field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetEnrollAuthenticatorTypes() []string {
	if o == nil || o.EnrollAuthenticatorTypes == nil {
		var ret []string
		return ret
	}
	return o.EnrollAuthenticatorTypes
}

// GetEnrollAuthenticatorTypesOk returns a tuple with the EnrollAuthenticatorTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetEnrollAuthenticatorTypesOk() ([]string, bool) {
	if o == nil || o.EnrollAuthenticatorTypes == nil {
		return nil, false
	}
	return o.EnrollAuthenticatorTypes, true
}

// HasEnrollAuthenticatorTypes returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasEnrollAuthenticatorTypes() bool {
	if o != nil && o.EnrollAuthenticatorTypes != nil {
		return true
	}

	return false
}

// SetEnrollAuthenticatorTypes gets a reference to the given []string and assigns it to the EnrollAuthenticatorTypes field.
func (o *ProfileEnrollmentPolicyRuleAction) SetEnrollAuthenticatorTypes(v []string) {
	o.EnrollAuthenticatorTypes = v
}

// GetPreRegistrationInlineHooks returns the PreRegistrationInlineHooks field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetPreRegistrationInlineHooks() []PreRegistrationInlineHook {
	if o == nil || o.PreRegistrationInlineHooks == nil {
		var ret []PreRegistrationInlineHook
		return ret
	}
	return o.PreRegistrationInlineHooks
}

// GetPreRegistrationInlineHooksOk returns a tuple with the PreRegistrationInlineHooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetPreRegistrationInlineHooksOk() ([]PreRegistrationInlineHook, bool) {
	if o == nil || o.PreRegistrationInlineHooks == nil {
		return nil, false
	}
	return o.PreRegistrationInlineHooks, true
}

// HasPreRegistrationInlineHooks returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasPreRegistrationInlineHooks() bool {
	if o != nil && o.PreRegistrationInlineHooks != nil {
		return true
	}

	return false
}

// SetPreRegistrationInlineHooks gets a reference to the given []PreRegistrationInlineHook and assigns it to the PreRegistrationInlineHooks field.
func (o *ProfileEnrollmentPolicyRuleAction) SetPreRegistrationInlineHooks(v []PreRegistrationInlineHook) {
	o.PreRegistrationInlineHooks = v
}

// GetProfileAttributes returns the ProfileAttributes field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetProfileAttributes() []ProfileEnrollmentPolicyRuleProfileAttribute {
	if o == nil || o.ProfileAttributes == nil {
		var ret []ProfileEnrollmentPolicyRuleProfileAttribute
		return ret
	}
	return o.ProfileAttributes
}

// GetProfileAttributesOk returns a tuple with the ProfileAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetProfileAttributesOk() ([]ProfileEnrollmentPolicyRuleProfileAttribute, bool) {
	if o == nil || o.ProfileAttributes == nil {
		return nil, false
	}
	return o.ProfileAttributes, true
}

// HasProfileAttributes returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasProfileAttributes() bool {
	if o != nil && o.ProfileAttributes != nil {
		return true
	}

	return false
}

// SetProfileAttributes gets a reference to the given []ProfileEnrollmentPolicyRuleProfileAttribute and assigns it to the ProfileAttributes field.
func (o *ProfileEnrollmentPolicyRuleAction) SetProfileAttributes(v []ProfileEnrollmentPolicyRuleProfileAttribute) {
	o.ProfileAttributes = v
}

// GetProgressiveProfilingAction returns the ProgressiveProfilingAction field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetProgressiveProfilingAction() string {
	if o == nil || o.ProgressiveProfilingAction == nil {
		var ret string
		return ret
	}
	return *o.ProgressiveProfilingAction
}

// GetProgressiveProfilingActionOk returns a tuple with the ProgressiveProfilingAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetProgressiveProfilingActionOk() (*string, bool) {
	if o == nil || o.ProgressiveProfilingAction == nil {
		return nil, false
	}
	return o.ProgressiveProfilingAction, true
}

// HasProgressiveProfilingAction returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasProgressiveProfilingAction() bool {
	if o != nil && o.ProgressiveProfilingAction != nil {
		return true
	}

	return false
}

// SetProgressiveProfilingAction gets a reference to the given string and assigns it to the ProgressiveProfilingAction field.
func (o *ProfileEnrollmentPolicyRuleAction) SetProgressiveProfilingAction(v string) {
	o.ProgressiveProfilingAction = &v
}

// GetTargetGroupIds returns the TargetGroupIds field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetTargetGroupIds() []string {
	if o == nil || o.TargetGroupIds == nil {
		var ret []string
		return ret
	}
	return o.TargetGroupIds
}

// GetTargetGroupIdsOk returns a tuple with the TargetGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetTargetGroupIdsOk() ([]string, bool) {
	if o == nil || o.TargetGroupIds == nil {
		return nil, false
	}
	return o.TargetGroupIds, true
}

// HasTargetGroupIds returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasTargetGroupIds() bool {
	if o != nil && o.TargetGroupIds != nil {
		return true
	}

	return false
}

// SetTargetGroupIds gets a reference to the given []string and assigns it to the TargetGroupIds field.
func (o *ProfileEnrollmentPolicyRuleAction) SetTargetGroupIds(v []string) {
	o.TargetGroupIds = v
}

// GetUiSchemaId returns the UiSchemaId field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetUiSchemaId() string {
	if o == nil || o.UiSchemaId == nil {
		var ret string
		return ret
	}
	return *o.UiSchemaId
}

// GetUiSchemaIdOk returns a tuple with the UiSchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetUiSchemaIdOk() (*string, bool) {
	if o == nil || o.UiSchemaId == nil {
		return nil, false
	}
	return o.UiSchemaId, true
}

// HasUiSchemaId returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasUiSchemaId() bool {
	if o != nil && o.UiSchemaId != nil {
		return true
	}

	return false
}

// SetUiSchemaId gets a reference to the given string and assigns it to the UiSchemaId field.
func (o *ProfileEnrollmentPolicyRuleAction) SetUiSchemaId(v string) {
	o.UiSchemaId = &v
}

// GetUnknownUserAction returns the UnknownUserAction field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleAction) GetUnknownUserAction() string {
	if o == nil || o.UnknownUserAction == nil {
		var ret string
		return ret
	}
	return *o.UnknownUserAction
}

// GetUnknownUserActionOk returns a tuple with the UnknownUserAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleAction) GetUnknownUserActionOk() (*string, bool) {
	if o == nil || o.UnknownUserAction == nil {
		return nil, false
	}
	return o.UnknownUserAction, true
}

// HasUnknownUserAction returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleAction) HasUnknownUserAction() bool {
	if o != nil && o.UnknownUserAction != nil {
		return true
	}

	return false
}

// SetUnknownUserAction gets a reference to the given string and assigns it to the UnknownUserAction field.
func (o *ProfileEnrollmentPolicyRuleAction) SetUnknownUserAction(v string) {
	o.UnknownUserAction = &v
}

func (o ProfileEnrollmentPolicyRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.ActivationRequirements != nil {
		toSerialize["activationRequirements"] = o.ActivationRequirements
	}
	if o.AllowedIdentifiers != nil {
		toSerialize["allowedIdentifiers"] = o.AllowedIdentifiers
	}
	if o.EnrollAuthenticatorTypes != nil {
		toSerialize["enrollAuthenticatorTypes"] = o.EnrollAuthenticatorTypes
	}
	if o.PreRegistrationInlineHooks != nil {
		toSerialize["preRegistrationInlineHooks"] = o.PreRegistrationInlineHooks
	}
	if o.ProfileAttributes != nil {
		toSerialize["profileAttributes"] = o.ProfileAttributes
	}
	if o.ProgressiveProfilingAction != nil {
		toSerialize["progressiveProfilingAction"] = o.ProgressiveProfilingAction
	}
	if o.TargetGroupIds != nil {
		toSerialize["targetGroupIds"] = o.TargetGroupIds
	}
	if o.UiSchemaId != nil {
		toSerialize["uiSchemaId"] = o.UiSchemaId
	}
	if o.UnknownUserAction != nil {
		toSerialize["unknownUserAction"] = o.UnknownUserAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProfileEnrollmentPolicyRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	varProfileEnrollmentPolicyRuleAction := _ProfileEnrollmentPolicyRuleAction{}

	err = json.Unmarshal(bytes, &varProfileEnrollmentPolicyRuleAction)
	if err == nil {
		*o = ProfileEnrollmentPolicyRuleAction(varProfileEnrollmentPolicyRuleAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "access")
		delete(additionalProperties, "activationRequirements")
		delete(additionalProperties, "allowedIdentifiers")
		delete(additionalProperties, "enrollAuthenticatorTypes")
		delete(additionalProperties, "preRegistrationInlineHooks")
		delete(additionalProperties, "profileAttributes")
		delete(additionalProperties, "progressiveProfilingAction")
		delete(additionalProperties, "targetGroupIds")
		delete(additionalProperties, "uiSchemaId")
		delete(additionalProperties, "unknownUserAction")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProfileEnrollmentPolicyRuleAction struct {
	value *ProfileEnrollmentPolicyRuleAction
	isSet bool
}

func (v NullableProfileEnrollmentPolicyRuleAction) Get() *ProfileEnrollmentPolicyRuleAction {
	return v.value
}

func (v *NullableProfileEnrollmentPolicyRuleAction) Set(val *ProfileEnrollmentPolicyRuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileEnrollmentPolicyRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileEnrollmentPolicyRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileEnrollmentPolicyRuleAction(val *ProfileEnrollmentPolicyRuleAction) *NullableProfileEnrollmentPolicyRuleAction {
	return &NullableProfileEnrollmentPolicyRuleAction{value: val, isSet: true}
}

func (v NullableProfileEnrollmentPolicyRuleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileEnrollmentPolicyRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

