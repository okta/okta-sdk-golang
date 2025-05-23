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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// ProfileEnrollmentPolicyRuleAction struct for ProfileEnrollmentPolicyRuleAction
type ProfileEnrollmentPolicyRuleAction struct {
	Access *string `json:"access,omitempty"`
	ActivationRequirements *ProfileEnrollmentPolicyRuleActivationRequirement `json:"activationRequirements,omitempty"`
	PreRegistrationInlineHooks []PreRegistrationInlineHook `json:"preRegistrationInlineHooks,omitempty"`
	ProfileAttributes []ProfileEnrollmentPolicyRuleProfileAttribute `json:"profileAttributes,omitempty"`
	ProgressiveProfilingAction *string `json:"progressiveProfilingAction,omitempty"`
	TargetGroupIds []string `json:"targetGroupIds,omitempty"`
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
		delete(additionalProperties, "preRegistrationInlineHooks")
		delete(additionalProperties, "profileAttributes")
		delete(additionalProperties, "progressiveProfilingAction")
		delete(additionalProperties, "targetGroupIds")
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

