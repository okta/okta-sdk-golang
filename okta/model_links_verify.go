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

// LinksVerify struct for LinksVerify
type LinksVerify struct {
	Verify *LinksVerifyVerify `json:"verify,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksVerify LinksVerify

// NewLinksVerify instantiates a new LinksVerify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksVerify() *LinksVerify {
	this := LinksVerify{}
	return &this
}

// NewLinksVerifyWithDefaults instantiates a new LinksVerify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksVerifyWithDefaults() *LinksVerify {
	this := LinksVerify{}
	return &this
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *LinksVerify) GetVerify() LinksVerifyVerify {
	if o == nil || o.Verify == nil {
		var ret LinksVerifyVerify
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksVerify) GetVerifyOk() (*LinksVerifyVerify, bool) {
	if o == nil || o.Verify == nil {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *LinksVerify) HasVerify() bool {
	if o != nil && o.Verify != nil {
		return true
	}

	return false
}

// SetVerify gets a reference to the given LinksVerifyVerify and assigns it to the Verify field.
func (o *LinksVerify) SetVerify(v LinksVerifyVerify) {
	o.Verify = &v
}

func (o LinksVerify) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Verify != nil {
		toSerialize["verify"] = o.Verify
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksVerify) UnmarshalJSON(bytes []byte) (err error) {
	varLinksVerify := _LinksVerify{}

	err = json.Unmarshal(bytes, &varLinksVerify)
	if err == nil {
		*o = LinksVerify(varLinksVerify)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "verify")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksVerify struct {
	value *LinksVerify
	isSet bool
}

func (v NullableLinksVerify) Get() *LinksVerify {
	return v.value
}

func (v *NullableLinksVerify) Set(val *LinksVerify) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksVerify) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksVerify(val *LinksVerify) *NullableLinksVerify {
	return &NullableLinksVerify{value: val, isSet: true}
}

func (v NullableLinksVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

