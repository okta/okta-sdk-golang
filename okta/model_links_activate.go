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

// checks if the LinksActivate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksActivate{}

// LinksActivate struct for LinksActivate
type LinksActivate struct {
	// Activates an enrolled factor. See [Activate a factor](/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/activateFactor).
	Activate             *HrefObject `json:"activate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksActivate LinksActivate

// NewLinksActivate instantiates a new LinksActivate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksActivate() *LinksActivate {
	this := LinksActivate{}
	return &this
}

// NewLinksActivateWithDefaults instantiates a new LinksActivate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksActivateWithDefaults() *LinksActivate {
	this := LinksActivate{}
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *LinksActivate) GetActivate() HrefObject {
	if o == nil || IsNil(o.Activate) {
		var ret HrefObject
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksActivate) GetActivateOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *LinksActivate) HasActivate() bool {
	if o != nil && !IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObject and assigns it to the Activate field.
func (o *LinksActivate) SetActivate(v HrefObject) {
	o.Activate = &v
}

func (o LinksActivate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksActivate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksActivate) UnmarshalJSON(data []byte) (err error) {
	varLinksActivate := _LinksActivate{}

	err = json.Unmarshal(data, &varLinksActivate)

	if err != nil {
		return err
	}

	*o = LinksActivate(varLinksActivate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksActivate struct {
	value *LinksActivate
	isSet bool
}

func (v NullableLinksActivate) Get() *LinksActivate {
	return v.value
}

func (v *NullableLinksActivate) Set(val *LinksActivate) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksActivate) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksActivate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksActivate(val *LinksActivate) *NullableLinksActivate {
	return &NullableLinksActivate{value: val, isSet: true}
}

func (v NullableLinksActivate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksActivate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
