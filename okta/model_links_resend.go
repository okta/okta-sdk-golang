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

// LinksResend struct for LinksResend
type LinksResend struct {
	Resend *LinksResendResend `json:"resend,omitempty"`
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
func (o *LinksResend) GetResend() LinksResendResend {
	if o == nil || o.Resend == nil {
		var ret LinksResendResend
		return ret
	}
	return *o.Resend
}

// GetResendOk returns a tuple with the Resend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksResend) GetResendOk() (*LinksResendResend, bool) {
	if o == nil || o.Resend == nil {
		return nil, false
	}
	return o.Resend, true
}

// HasResend returns a boolean if a field has been set.
func (o *LinksResend) HasResend() bool {
	if o != nil && o.Resend != nil {
		return true
	}

	return false
}

// SetResend gets a reference to the given LinksResendResend and assigns it to the Resend field.
func (o *LinksResend) SetResend(v LinksResendResend) {
	o.Resend = &v
}

func (o LinksResend) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Resend != nil {
		toSerialize["resend"] = o.Resend
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksResend) UnmarshalJSON(bytes []byte) (err error) {
	varLinksResend := _LinksResend{}

	err = json.Unmarshal(bytes, &varLinksResend)
	if err == nil {
		*o = LinksResend(varLinksResend)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resend")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

