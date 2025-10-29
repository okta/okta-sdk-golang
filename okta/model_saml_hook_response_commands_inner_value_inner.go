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

// checks if the SAMLHookResponseCommandsInnerValueInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLHookResponseCommandsInnerValueInner{}

// SAMLHookResponseCommandsInnerValueInner struct for SAMLHookResponseCommandsInnerValueInner
type SAMLHookResponseCommandsInnerValueInner struct {
	// The name of one of the supported ops: `add`:  Add a new claim to the assertion `replace`: Modify any element of the assertion  > **Note:** If a response to the SAML assertion inline hook request isn't received from your external service within three seconds, a timeout occurs. In this scenario, the Okta process flow continues with the original SAML assertion returned.
	Op *string `json:"op,omitempty"`
	// Location, within the assertion, to apply the operation
	Path                 *string                                       `json:"path,omitempty"`
	Value                *SAMLHookResponseCommandsInnerValueInnerValue `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLHookResponseCommandsInnerValueInner SAMLHookResponseCommandsInnerValueInner

// NewSAMLHookResponseCommandsInnerValueInner instantiates a new SAMLHookResponseCommandsInnerValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLHookResponseCommandsInnerValueInner() *SAMLHookResponseCommandsInnerValueInner {
	this := SAMLHookResponseCommandsInnerValueInner{}
	return &this
}

// NewSAMLHookResponseCommandsInnerValueInnerWithDefaults instantiates a new SAMLHookResponseCommandsInnerValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLHookResponseCommandsInnerValueInnerWithDefaults() *SAMLHookResponseCommandsInnerValueInner {
	this := SAMLHookResponseCommandsInnerValueInner{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *SAMLHookResponseCommandsInnerValueInner) GetOp() string {
	if o == nil || IsNil(o.Op) {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLHookResponseCommandsInnerValueInner) GetOpOk() (*string, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *SAMLHookResponseCommandsInnerValueInner) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *SAMLHookResponseCommandsInnerValueInner) SetOp(v string) {
	o.Op = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SAMLHookResponseCommandsInnerValueInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLHookResponseCommandsInnerValueInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SAMLHookResponseCommandsInnerValueInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SAMLHookResponseCommandsInnerValueInner) SetPath(v string) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SAMLHookResponseCommandsInnerValueInner) GetValue() SAMLHookResponseCommandsInnerValueInnerValue {
	if o == nil || IsNil(o.Value) {
		var ret SAMLHookResponseCommandsInnerValueInnerValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLHookResponseCommandsInnerValueInner) GetValueOk() (*SAMLHookResponseCommandsInnerValueInnerValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SAMLHookResponseCommandsInnerValueInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given SAMLHookResponseCommandsInnerValueInnerValue and assigns it to the Value field.
func (o *SAMLHookResponseCommandsInnerValueInner) SetValue(v SAMLHookResponseCommandsInnerValueInnerValue) {
	o.Value = &v
}

func (o SAMLHookResponseCommandsInnerValueInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLHookResponseCommandsInnerValueInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLHookResponseCommandsInnerValueInner) UnmarshalJSON(data []byte) (err error) {
	varSAMLHookResponseCommandsInnerValueInner := _SAMLHookResponseCommandsInnerValueInner{}

	err = json.Unmarshal(data, &varSAMLHookResponseCommandsInnerValueInner)

	if err != nil {
		return err
	}

	*o = SAMLHookResponseCommandsInnerValueInner(varSAMLHookResponseCommandsInnerValueInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLHookResponseCommandsInnerValueInner struct {
	value *SAMLHookResponseCommandsInnerValueInner
	isSet bool
}

func (v NullableSAMLHookResponseCommandsInnerValueInner) Get() *SAMLHookResponseCommandsInnerValueInner {
	return v.value
}

func (v *NullableSAMLHookResponseCommandsInnerValueInner) Set(val *SAMLHookResponseCommandsInnerValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLHookResponseCommandsInnerValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLHookResponseCommandsInnerValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLHookResponseCommandsInnerValueInner(val *SAMLHookResponseCommandsInnerValueInner) *NullableSAMLHookResponseCommandsInnerValueInner {
	return &NullableSAMLHookResponseCommandsInnerValueInner{value: val, isSet: true}
}

func (v NullableSAMLHookResponseCommandsInnerValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLHookResponseCommandsInnerValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
