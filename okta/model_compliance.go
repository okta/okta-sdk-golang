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

// Compliance struct for Compliance
type Compliance struct {
	Fips *string `json:"fips,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Compliance Compliance

// NewCompliance instantiates a new Compliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompliance() *Compliance {
	this := Compliance{}
	return &this
}

// NewComplianceWithDefaults instantiates a new Compliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComplianceWithDefaults() *Compliance {
	this := Compliance{}
	return &this
}

// GetFips returns the Fips field value if set, zero value otherwise.
func (o *Compliance) GetFips() string {
	if o == nil || o.Fips == nil {
		var ret string
		return ret
	}
	return *o.Fips
}

// GetFipsOk returns a tuple with the Fips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Compliance) GetFipsOk() (*string, bool) {
	if o == nil || o.Fips == nil {
		return nil, false
	}
	return o.Fips, true
}

// HasFips returns a boolean if a field has been set.
func (o *Compliance) HasFips() bool {
	if o != nil && o.Fips != nil {
		return true
	}

	return false
}

// SetFips gets a reference to the given string and assigns it to the Fips field.
func (o *Compliance) SetFips(v string) {
	o.Fips = &v
}

func (o Compliance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fips != nil {
		toSerialize["fips"] = o.Fips
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Compliance) UnmarshalJSON(bytes []byte) (err error) {
	varCompliance := _Compliance{}

	err = json.Unmarshal(bytes, &varCompliance)
	if err == nil {
		*o = Compliance(varCompliance)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "fips")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCompliance struct {
	value *Compliance
	isSet bool
}

func (v NullableCompliance) Get() *Compliance {
	return v.value
}

func (v *NullableCompliance) Set(val *Compliance) {
	v.value = val
	v.isSet = true
}

func (v NullableCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompliance(val *Compliance) *NullableCompliance {
	return &NullableCompliance{value: val, isSet: true}
}

func (v NullableCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

