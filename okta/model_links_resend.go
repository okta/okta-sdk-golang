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

// checks if the LinksResend type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksResend{}

// LinksResend struct for LinksResend
type LinksResend struct {
	// Resends the factor enrollment challenge. See [Resend a factor enrollment](/openapi/okta-management/management/tag/UserFactor/#tag/UserFactor/operation/resendEnrollFactor).
	Resend               *HrefObject `json:"resend,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksResend LinksResend

// NewLinksResend instantiates a new LinksResend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksResend() *LinksResend {
	this := LinksResend{}
	return &this
}

// NewLinksResendWithDefaults instantiates a new LinksResend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksResendWithDefaults() *LinksResend {
	this := LinksResend{}
	return &this
}

// GetResend returns the Resend field value if set, zero value otherwise.
func (o *LinksResend) GetResend() HrefObject {
	if o == nil || IsNil(o.Resend) {
		var ret HrefObject
		return ret
	}
	return *o.Resend
}

// GetResendOk returns a tuple with the Resend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksResend) GetResendOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Resend) {
		return nil, false
	}
	return o.Resend, true
}

// HasResend returns a boolean if a field has been set.
func (o *LinksResend) HasResend() bool {
	if o != nil && !IsNil(o.Resend) {
		return true
	}

	return false
}

// SetResend gets a reference to the given HrefObject and assigns it to the Resend field.
func (o *LinksResend) SetResend(v HrefObject) {
	o.Resend = &v
}

func (o LinksResend) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksResend) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resend) {
		toSerialize["resend"] = o.Resend
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksResend) UnmarshalJSON(data []byte) (err error) {
	varLinksResend := _LinksResend{}

	err = json.Unmarshal(data, &varLinksResend)

	if err != nil {
		return err
	}

	*o = LinksResend(varLinksResend)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resend")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksResend struct {
	value *LinksResend
	isSet bool
}

func (v NullableLinksResend) Get() *LinksResend {
	return v.value
}

func (v *NullableLinksResend) Set(val *LinksResend) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksResend) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksResend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksResend(val *LinksResend) *NullableLinksResend {
	return &NullableLinksResend{value: val, isSet: true}
}

func (v NullableLinksResend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksResend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
