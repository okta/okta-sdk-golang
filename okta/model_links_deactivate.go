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

// LinksDeactivate struct for LinksDeactivate
type LinksDeactivate struct {
	Deactivate *LinksDeactivateDeactivate `json:"deactivate,omitempty"`
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
func (o *LinksDeactivate) GetDeactivate() LinksDeactivateDeactivate {
	if o == nil || o.Deactivate == nil {
		var ret LinksDeactivateDeactivate
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksDeactivate) GetDeactivateOk() (*LinksDeactivateDeactivate, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *LinksDeactivate) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given LinksDeactivateDeactivate and assigns it to the Deactivate field.
func (o *LinksDeactivate) SetDeactivate(v LinksDeactivateDeactivate) {
	o.Deactivate = &v
}

func (o LinksDeactivate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksDeactivate) UnmarshalJSON(bytes []byte) (err error) {
	varLinksDeactivate := _LinksDeactivate{}

	err = json.Unmarshal(bytes, &varLinksDeactivate)
	if err == nil {
		*o = LinksDeactivate(varLinksDeactivate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "deactivate")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

