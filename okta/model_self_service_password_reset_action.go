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

// SelfServicePasswordResetAction struct for SelfServicePasswordResetAction
type SelfServicePasswordResetAction struct {
	Access *string `json:"access,omitempty"`
	// The type of rule action
	Type *string `json:"type,omitempty"`
	Requirement *SsprRequirement `json:"requirement,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServicePasswordResetAction SelfServicePasswordResetAction

// NewSelfServicePasswordResetAction instantiates a new SelfServicePasswordResetAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServicePasswordResetAction() *SelfServicePasswordResetAction {
	this := SelfServicePasswordResetAction{}
	return &this
}

// NewSelfServicePasswordResetActionWithDefaults instantiates a new SelfServicePasswordResetAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServicePasswordResetActionWithDefaults() *SelfServicePasswordResetAction {
	this := SelfServicePasswordResetAction{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *SelfServicePasswordResetAction) GetAccess() string {
	if o == nil || o.Access == nil {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicePasswordResetAction) GetAccessOk() (*string, bool) {
	if o == nil || o.Access == nil {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *SelfServicePasswordResetAction) HasAccess() bool {
	if o != nil && o.Access != nil {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *SelfServicePasswordResetAction) SetAccess(v string) {
	o.Access = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SelfServicePasswordResetAction) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicePasswordResetAction) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SelfServicePasswordResetAction) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SelfServicePasswordResetAction) SetType(v string) {
	o.Type = &v
}

// GetRequirement returns the Requirement field value if set, zero value otherwise.
func (o *SelfServicePasswordResetAction) GetRequirement() SsprRequirement {
	if o == nil || o.Requirement == nil {
		var ret SsprRequirement
		return ret
	}
	return *o.Requirement
}

// GetRequirementOk returns a tuple with the Requirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicePasswordResetAction) GetRequirementOk() (*SsprRequirement, bool) {
	if o == nil || o.Requirement == nil {
		return nil, false
	}
	return o.Requirement, true
}

// HasRequirement returns a boolean if a field has been set.
func (o *SelfServicePasswordResetAction) HasRequirement() bool {
	if o != nil && o.Requirement != nil {
		return true
	}

	return false
}

// SetRequirement gets a reference to the given SsprRequirement and assigns it to the Requirement field.
func (o *SelfServicePasswordResetAction) SetRequirement(v SsprRequirement) {
	o.Requirement = &v
}

func (o SelfServicePasswordResetAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Access != nil {
		toSerialize["access"] = o.Access
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Requirement != nil {
		toSerialize["requirement"] = o.Requirement
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SelfServicePasswordResetAction) UnmarshalJSON(bytes []byte) (err error) {
	varSelfServicePasswordResetAction := _SelfServicePasswordResetAction{}

	err = json.Unmarshal(bytes, &varSelfServicePasswordResetAction)
	if err == nil {
		*o = SelfServicePasswordResetAction(varSelfServicePasswordResetAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "access")
		delete(additionalProperties, "type")
		delete(additionalProperties, "requirement")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSelfServicePasswordResetAction struct {
	value *SelfServicePasswordResetAction
	isSet bool
}

func (v NullableSelfServicePasswordResetAction) Get() *SelfServicePasswordResetAction {
	return v.value
}

func (v *NullableSelfServicePasswordResetAction) Set(val *SelfServicePasswordResetAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServicePasswordResetAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServicePasswordResetAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServicePasswordResetAction(val *SelfServicePasswordResetAction) *NullableSelfServicePasswordResetAction {
	return &NullableSelfServicePasswordResetAction{value: val, isSet: true}
}

func (v NullableSelfServicePasswordResetAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServicePasswordResetAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

