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

// checks if the PatchAIAgentProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchAIAgentProfile{}

// PatchAIAgentProfile Partial update for AI agent profile
type PatchAIAgentProfile struct {
	// Description of the AI agent
	Description NullableString `json:"description,omitempty"`
	// Unique name of the AI agent
	Name                 NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchAIAgentProfile PatchAIAgentProfile

// NewPatchAIAgentProfile instantiates a new PatchAIAgentProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchAIAgentProfile() *PatchAIAgentProfile {
	this := PatchAIAgentProfile{}
	return &this
}

// NewPatchAIAgentProfileWithDefaults instantiates a new PatchAIAgentProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchAIAgentProfileWithDefaults() *PatchAIAgentProfile {
	this := PatchAIAgentProfile{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchAIAgentProfile) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchAIAgentProfile) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchAIAgentProfile) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchAIAgentProfile) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchAIAgentProfile) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchAIAgentProfile) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchAIAgentProfile) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchAIAgentProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PatchAIAgentProfile) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PatchAIAgentProfile) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *PatchAIAgentProfile) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PatchAIAgentProfile) UnsetName() {
	o.Name.Unset()
}

func (o PatchAIAgentProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchAIAgentProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchAIAgentProfile) UnmarshalJSON(data []byte) (err error) {
	varPatchAIAgentProfile := _PatchAIAgentProfile{}

	err = json.Unmarshal(data, &varPatchAIAgentProfile)

	if err != nil {
		return err
	}

	*o = PatchAIAgentProfile(varPatchAIAgentProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchAIAgentProfile struct {
	value *PatchAIAgentProfile
	isSet bool
}

func (v NullablePatchAIAgentProfile) Get() *PatchAIAgentProfile {
	return v.value
}

func (v *NullablePatchAIAgentProfile) Set(val *PatchAIAgentProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchAIAgentProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchAIAgentProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchAIAgentProfile(val *PatchAIAgentProfile) *NullablePatchAIAgentProfile {
	return &NullablePatchAIAgentProfile{value: val, isSet: true}
}

func (v NullablePatchAIAgentProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchAIAgentProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
