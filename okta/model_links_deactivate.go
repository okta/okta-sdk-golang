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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the LinksDeactivate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksDeactivate{}

// LinksDeactivate struct for LinksDeactivate
type LinksDeactivate struct {
	// Deactivates the factor. See [Unenroll a factor](/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/unenrollFactor).
	Deactivate           *HrefObject `json:"deactivate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksDeactivate LinksDeactivate

// NewLinksDeactivate instantiates a new LinksDeactivate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksDeactivate() *LinksDeactivate {
	this := LinksDeactivate{}
	return &this
}

// NewLinksDeactivateWithDefaults instantiates a new LinksDeactivate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksDeactivateWithDefaults() *LinksDeactivate {
	this := LinksDeactivate{}
	return &this
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *LinksDeactivate) GetDeactivate() HrefObject {
	if o == nil || IsNil(o.Deactivate) {
		var ret HrefObject
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksDeactivate) GetDeactivateOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *LinksDeactivate) HasDeactivate() bool {
	if o != nil && !IsNil(o.Deactivate) {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObject and assigns it to the Deactivate field.
func (o *LinksDeactivate) SetDeactivate(v HrefObject) {
	o.Deactivate = &v
}

func (o LinksDeactivate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksDeactivate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deactivate) {
		toSerialize["deactivate"] = o.Deactivate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksDeactivate) UnmarshalJSON(data []byte) (err error) {
	varLinksDeactivate := _LinksDeactivate{}

	err = json.Unmarshal(data, &varLinksDeactivate)

	if err != nil {
		return err
	}

	*o = LinksDeactivate(varLinksDeactivate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deactivate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksDeactivate struct {
	value *LinksDeactivate
	isSet bool
}

func (v NullableLinksDeactivate) Get() *LinksDeactivate {
	return v.value
}

func (v *NullableLinksDeactivate) Set(val *LinksDeactivate) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksDeactivate) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksDeactivate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksDeactivate(val *LinksDeactivate) *NullableLinksDeactivate {
	return &NullableLinksDeactivate{value: val, isSet: true}
}

func (v NullableLinksDeactivate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksDeactivate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
