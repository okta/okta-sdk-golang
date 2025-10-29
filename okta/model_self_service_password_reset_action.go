/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SelfServicePasswordResetAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServicePasswordResetAction{}

// SelfServicePasswordResetAction Enables or disables users to reset their own password and defines the authenticators and constraints needed to complete the reset
type SelfServicePasswordResetAction struct {
	Access      *string          `json:"access,omitempty"`
	Requirement *SsprRequirement `json:"requirement,omitempty"`
	// <x-lifecycle class=\"oie\"></x-lifecycle> The type of rule action
	Type                 *string `json:"type,omitempty"`
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
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicePasswordResetAction) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *SelfServicePasswordResetAction) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *SelfServicePasswordResetAction) SetAccess(v string) {
	o.Access = &v
}

// GetRequirement returns the Requirement field value if set, zero value otherwise.
func (o *SelfServicePasswordResetAction) GetRequirement() SsprRequirement {
	if o == nil || IsNil(o.Requirement) {
		var ret SsprRequirement
		return ret
	}
	return *o.Requirement
}

// GetRequirementOk returns a tuple with the Requirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicePasswordResetAction) GetRequirementOk() (*SsprRequirement, bool) {
	if o == nil || IsNil(o.Requirement) {
		return nil, false
	}
	return o.Requirement, true
}

// HasRequirement returns a boolean if a field has been set.
func (o *SelfServicePasswordResetAction) HasRequirement() bool {
	if o != nil && !IsNil(o.Requirement) {
		return true
	}

	return false
}

// SetRequirement gets a reference to the given SsprRequirement and assigns it to the Requirement field.
func (o *SelfServicePasswordResetAction) SetRequirement(v SsprRequirement) {
	o.Requirement = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SelfServicePasswordResetAction) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServicePasswordResetAction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SelfServicePasswordResetAction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SelfServicePasswordResetAction) SetType(v string) {
	o.Type = &v
}

func (o SelfServicePasswordResetAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServicePasswordResetAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.Requirement) {
		toSerialize["requirement"] = o.Requirement
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelfServicePasswordResetAction) UnmarshalJSON(data []byte) (err error) {
	varSelfServicePasswordResetAction := _SelfServicePasswordResetAction{}

	err = json.Unmarshal(data, &varSelfServicePasswordResetAction)

	if err != nil {
		return err
	}

	*o = SelfServicePasswordResetAction(varSelfServicePasswordResetAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access")
		delete(additionalProperties, "requirement")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
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
