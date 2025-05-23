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

// LinksEnroll struct for LinksEnroll
type LinksEnroll struct {
	Enroll *LinksEnrollEnroll `json:"enroll,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksEnroll LinksEnroll

// NewLinksEnroll instantiates a new LinksEnroll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksEnroll() *LinksEnroll {
	this := LinksEnroll{}
	return &this
}

// NewLinksEnrollWithDefaults instantiates a new LinksEnroll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksEnrollWithDefaults() *LinksEnroll {
	this := LinksEnroll{}
	return &this
}

// GetEnroll returns the Enroll field value if set, zero value otherwise.
func (o *LinksEnroll) GetEnroll() LinksEnrollEnroll {
	if o == nil || o.Enroll == nil {
		var ret LinksEnrollEnroll
		return ret
	}
	return *o.Enroll
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksEnroll) GetEnrollOk() (*LinksEnrollEnroll, bool) {
	if o == nil || o.Enroll == nil {
		return nil, false
	}
	return o.Enroll, true
}

// HasEnroll returns a boolean if a field has been set.
func (o *LinksEnroll) HasEnroll() bool {
	if o != nil && o.Enroll != nil {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given LinksEnrollEnroll and assigns it to the Enroll field.
func (o *LinksEnroll) SetEnroll(v LinksEnrollEnroll) {
	o.Enroll = &v
}

func (o LinksEnroll) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enroll != nil {
		toSerialize["enroll"] = o.Enroll
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksEnroll) UnmarshalJSON(bytes []byte) (err error) {
	varLinksEnroll := _LinksEnroll{}

	err = json.Unmarshal(bytes, &varLinksEnroll)
	if err == nil {
		*o = LinksEnroll(varLinksEnroll)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "enroll")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksEnroll struct {
	value *LinksEnroll
	isSet bool
}

func (v NullableLinksEnroll) Get() *LinksEnroll {
	return v.value
}

func (v *NullableLinksEnroll) Set(val *LinksEnroll) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksEnroll) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksEnroll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksEnroll(val *LinksEnroll) *NullableLinksEnroll {
	return &NullableLinksEnroll{value: val, isSet: true}
}

func (v NullableLinksEnroll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksEnroll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

