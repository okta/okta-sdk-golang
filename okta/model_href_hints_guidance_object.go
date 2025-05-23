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

// HrefHintsGuidanceObject struct for HrefHintsGuidanceObject
type HrefHintsGuidanceObject struct {
	Allow []string `json:"allow,omitempty"`
	// Specifies the URI to invoke for granting scope consent required to complete the OAuth 2.0 connection 
	Guidance []string `json:"guidance,omitempty"`
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
	if o == nil || o.Allow == nil {
		var ret []string
		return ret
	}
	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefHintsGuidanceObject) GetAllowOk() ([]string, bool) {
	if o == nil || o.Allow == nil {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *HrefHintsGuidanceObject) HasAllow() bool {
	if o != nil && o.Allow != nil {
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
	if o == nil || o.Guidance == nil {
		var ret []string
		return ret
	}
	return o.Guidance
}

// GetGuidanceOk returns a tuple with the Guidance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefHintsGuidanceObject) GetGuidanceOk() ([]string, bool) {
	if o == nil || o.Guidance == nil {
		return nil, false
	}
	return o.Guidance, true
}

// HasGuidance returns a boolean if a field has been set.
func (o *HrefHintsGuidanceObject) HasGuidance() bool {
	if o != nil && o.Guidance != nil {
		return true
	}

	return false
}

// SetGuidance gets a reference to the given []string and assigns it to the Guidance field.
func (o *HrefHintsGuidanceObject) SetGuidance(v []string) {
	o.Guidance = v
}

func (o HrefHintsGuidanceObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allow != nil {
		toSerialize["allow"] = o.Allow
	}
	if o.Guidance != nil {
		toSerialize["guidance"] = o.Guidance
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HrefHintsGuidanceObject) UnmarshalJSON(bytes []byte) (err error) {
	varHrefHintsGuidanceObject := _HrefHintsGuidanceObject{}

	err = json.Unmarshal(bytes, &varHrefHintsGuidanceObject)
	if err == nil {
		*o = HrefHintsGuidanceObject(varHrefHintsGuidanceObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allow")
		delete(additionalProperties, "guidance")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

