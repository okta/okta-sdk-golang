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

// LinksQrcode struct for LinksQrcode
type LinksQrcode struct {
	Qrcode *LinksQrcodeQrcode `json:"qrcode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksQrcode LinksQrcode

// NewLinksQrcode instantiates a new LinksQrcode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksQrcode() *LinksQrcode {
	this := LinksQrcode{}
	return &this
}

// NewLinksQrcodeWithDefaults instantiates a new LinksQrcode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksQrcodeWithDefaults() *LinksQrcode {
	this := LinksQrcode{}
	return &this
}

// GetQrcode returns the Qrcode field value if set, zero value otherwise.
func (o *LinksQrcode) GetQrcode() LinksQrcodeQrcode {
	if o == nil || o.Qrcode == nil {
		var ret LinksQrcodeQrcode
		return ret
	}
	return *o.Qrcode
}

// GetQrcodeOk returns a tuple with the Qrcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksQrcode) GetQrcodeOk() (*LinksQrcodeQrcode, bool) {
	if o == nil || o.Qrcode == nil {
		return nil, false
	}
	return o.Qrcode, true
}

// HasQrcode returns a boolean if a field has been set.
func (o *LinksQrcode) HasQrcode() bool {
	if o != nil && o.Qrcode != nil {
		return true
	}

	return false
}

// SetQrcode gets a reference to the given LinksQrcodeQrcode and assigns it to the Qrcode field.
func (o *LinksQrcode) SetQrcode(v LinksQrcodeQrcode) {
	o.Qrcode = &v
}

func (o LinksQrcode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Qrcode != nil {
		toSerialize["qrcode"] = o.Qrcode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksQrcode) UnmarshalJSON(bytes []byte) (err error) {
	varLinksQrcode := _LinksQrcode{}

	err = json.Unmarshal(bytes, &varLinksQrcode)
	if err == nil {
		*o = LinksQrcode(varLinksQrcode)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "qrcode")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksQrcode struct {
	value *LinksQrcode
	isSet bool
}

func (v NullableLinksQrcode) Get() *LinksQrcode {
	return v.value
}

func (v *NullableLinksQrcode) Set(val *LinksQrcode) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksQrcode) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksQrcode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksQrcode(val *LinksQrcode) *NullableLinksQrcode {
	return &NullableLinksQrcode{value: val, isSet: true}
}

func (v NullableLinksQrcode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksQrcode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

