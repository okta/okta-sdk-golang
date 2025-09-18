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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the HrefHintsGuidanceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HrefHintsGuidanceObject{}

// HrefHintsGuidanceObject Describes allowed HTTP verbs and guidance for the `href`
type HrefHintsGuidanceObject struct {
	Allow []string `json:"allow,omitempty"`
	// Specifies the URI to invoke for granting scope consent required to complete the OAuth 2.0 connection
	Guidance             []string `json:"guidance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HrefHintsGuidanceObject HrefHintsGuidanceObject

// NewHrefHintsGuidanceObject instantiates a new HrefHintsGuidanceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHrefHintsGuidanceObject() *HrefHintsGuidanceObject {
	this := HrefHintsGuidanceObject{}
	return &this
}

// NewHrefHintsGuidanceObjectWithDefaults instantiates a new HrefHintsGuidanceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrefHintsGuidanceObjectWithDefaults() *HrefHintsGuidanceObject {
	this := HrefHintsGuidanceObject{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *HrefHintsGuidanceObject) GetAllow() []string {
	if o == nil || IsNil(o.Allow) {
		var ret []string
		return ret
	}
	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefHintsGuidanceObject) GetAllowOk() ([]string, bool) {
	if o == nil || IsNil(o.Allow) {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *HrefHintsGuidanceObject) HasAllow() bool {
	if o != nil && !IsNil(o.Allow) {
		return true
	}

	return false
}

// SetAllow gets a reference to the given []string and assigns it to the Allow field.
func (o *HrefHintsGuidanceObject) SetAllow(v []string) {
	o.Allow = v
}

// GetGuidance returns the Guidance field value if set, zero value otherwise.
func (o *HrefHintsGuidanceObject) GetGuidance() []string {
	if o == nil || IsNil(o.Guidance) {
		var ret []string
		return ret
	}
	return o.Guidance
}

// GetGuidanceOk returns a tuple with the Guidance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefHintsGuidanceObject) GetGuidanceOk() ([]string, bool) {
	if o == nil || IsNil(o.Guidance) {
		return nil, false
	}
	return o.Guidance, true
}

// HasGuidance returns a boolean if a field has been set.
func (o *HrefHintsGuidanceObject) HasGuidance() bool {
	if o != nil && !IsNil(o.Guidance) {
		return true
	}

	return false
}

// SetGuidance gets a reference to the given []string and assigns it to the Guidance field.
func (o *HrefHintsGuidanceObject) SetGuidance(v []string) {
	o.Guidance = v
}

func (o HrefHintsGuidanceObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HrefHintsGuidanceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allow) {
		toSerialize["allow"] = o.Allow
	}
	if !IsNil(o.Guidance) {
		toSerialize["guidance"] = o.Guidance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HrefHintsGuidanceObject) UnmarshalJSON(data []byte) (err error) {
	varHrefHintsGuidanceObject := _HrefHintsGuidanceObject{}

	err = json.Unmarshal(data, &varHrefHintsGuidanceObject)

	if err != nil {
		return err
	}

	*o = HrefHintsGuidanceObject(varHrefHintsGuidanceObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allow")
		delete(additionalProperties, "guidance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHrefHintsGuidanceObject struct {
	value *HrefHintsGuidanceObject
	isSet bool
}

func (v NullableHrefHintsGuidanceObject) Get() *HrefHintsGuidanceObject {
	return v.value
}

func (v *NullableHrefHintsGuidanceObject) Set(val *HrefHintsGuidanceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableHrefHintsGuidanceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableHrefHintsGuidanceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHrefHintsGuidanceObject(val *HrefHintsGuidanceObject) *NullableHrefHintsGuidanceObject {
	return &NullableHrefHintsGuidanceObject{value: val, isSet: true}
}

func (v NullableHrefHintsGuidanceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHrefHintsGuidanceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
