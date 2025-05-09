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

// SpCertificate struct for SpCertificate
type SpCertificate struct {
	X5c []string `json:"x5c,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SpCertificate SpCertificate

// NewSpCertificate instantiates a new SpCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpCertificate() *SpCertificate {
	this := SpCertificate{}
	return &this
}

// NewSpCertificateWithDefaults instantiates a new SpCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpCertificateWithDefaults() *SpCertificate {
	this := SpCertificate{}
	return &this
}

// GetX5c returns the X5c field value if set, zero value otherwise.
func (o *SpCertificate) GetX5c() []string {
	if o == nil || o.X5c == nil {
		var ret []string
		return ret
	}
	return o.X5c
}

// GetX5cOk returns a tuple with the X5c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpCertificate) GetX5cOk() ([]string, bool) {
	if o == nil || o.X5c == nil {
		return nil, false
	}
	return o.X5c, true
}

// HasX5c returns a boolean if a field has been set.
func (o *SpCertificate) HasX5c() bool {
	if o != nil && o.X5c != nil {
		return true
	}

	return false
}

// SetX5c gets a reference to the given []string and assigns it to the X5c field.
func (o *SpCertificate) SetX5c(v []string) {
	o.X5c = v
}

func (o SpCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.X5c != nil {
		toSerialize["x5c"] = o.X5c
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SpCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varSpCertificate := _SpCertificate{}

	err = json.Unmarshal(bytes, &varSpCertificate)
	if err == nil {
		*o = SpCertificate(varSpCertificate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "x5c")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSpCertificate struct {
	value *SpCertificate
	isSet bool
}

func (v NullableSpCertificate) Get() *SpCertificate {
	return v.value
}

func (v *NullableSpCertificate) Set(val *SpCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableSpCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableSpCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpCertificate(val *SpCertificate) *NullableSpCertificate {
	return &NullableSpCertificate{value: val, isSet: true}
}

func (v NullableSpCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

