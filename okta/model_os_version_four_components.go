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

// OSVersionFourComponents Current version of the operating system (maximum of four components in the versioning scheme)
type OSVersionFourComponents struct {
	Minimum *string `json:"minimum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OSVersionFourComponents OSVersionFourComponents

// NewOSVersionFourComponents instantiates a new OSVersionFourComponents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSVersionFourComponents() *OSVersionFourComponents {
	this := OSVersionFourComponents{}
	return &this
}

// NewOSVersionFourComponentsWithDefaults instantiates a new OSVersionFourComponents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSVersionFourComponentsWithDefaults() *OSVersionFourComponents {
	this := OSVersionFourComponents{}
	return &this
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *OSVersionFourComponents) GetMinimum() string {
	if o == nil || o.Minimum == nil {
		var ret string
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSVersionFourComponents) GetMinimumOk() (*string, bool) {
	if o == nil || o.Minimum == nil {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *OSVersionFourComponents) HasMinimum() bool {
	if o != nil && o.Minimum != nil {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given string and assigns it to the Minimum field.
func (o *OSVersionFourComponents) SetMinimum(v string) {
	o.Minimum = &v
}

func (o OSVersionFourComponents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Minimum != nil {
		toSerialize["minimum"] = o.Minimum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OSVersionFourComponents) UnmarshalJSON(bytes []byte) (err error) {
	varOSVersionFourComponents := _OSVersionFourComponents{}

	err = json.Unmarshal(bytes, &varOSVersionFourComponents)
	if err == nil {
		*o = OSVersionFourComponents(varOSVersionFourComponents)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "minimum")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOSVersionFourComponents struct {
	value *OSVersionFourComponents
	isSet bool
}

func (v NullableOSVersionFourComponents) Get() *OSVersionFourComponents {
	return v.value
}

func (v *NullableOSVersionFourComponents) Set(val *OSVersionFourComponents) {
	v.value = val
	v.isSet = true
}

func (v NullableOSVersionFourComponents) IsSet() bool {
	return v.isSet
}

func (v *NullableOSVersionFourComponents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSVersionFourComponents(val *OSVersionFourComponents) *NullableOSVersionFourComponents {
	return &NullableOSVersionFourComponents{value: val, isSet: true}
}

func (v NullableOSVersionFourComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSVersionFourComponents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

