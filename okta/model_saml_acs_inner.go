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

// SamlAcsInner struct for SamlAcsInner
type SamlAcsInner struct {
	// Index of ACS URL. You can't reuse the same index in the ACS URL array.
	Index *float32 `json:"index,omitempty"`
	// Assertion Consumer Service (ACS) URL
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlAcsInner SamlAcsInner

// NewSamlAcsInner instantiates a new SamlAcsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlAcsInner() *SamlAcsInner {
	this := SamlAcsInner{}
	return &this
}

// NewSamlAcsInnerWithDefaults instantiates a new SamlAcsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAcsInnerWithDefaults() *SamlAcsInner {
	this := SamlAcsInner{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SamlAcsInner) GetIndex() float32 {
	if o == nil || o.Index == nil {
		var ret float32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAcsInner) GetIndexOk() (*float32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SamlAcsInner) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given float32 and assigns it to the Index field.
func (o *SamlAcsInner) SetIndex(v float32) {
	o.Index = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SamlAcsInner) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAcsInner) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SamlAcsInner) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SamlAcsInner) SetUrl(v string) {
	o.Url = &v
}

func (o SamlAcsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlAcsInner) UnmarshalJSON(bytes []byte) (err error) {
	varSamlAcsInner := _SamlAcsInner{}

	err = json.Unmarshal(bytes, &varSamlAcsInner)
	if err == nil {
		*o = SamlAcsInner(varSamlAcsInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "index")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlAcsInner struct {
	value *SamlAcsInner
	isSet bool
}

func (v NullableSamlAcsInner) Get() *SamlAcsInner {
	return v.value
}

func (v *NullableSamlAcsInner) Set(val *SamlAcsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlAcsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlAcsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlAcsInner(val *SamlAcsInner) *NullableSamlAcsInner {
	return &NullableSamlAcsInner{value: val, isSet: true}
}

func (v NullableSamlAcsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlAcsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

