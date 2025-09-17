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
	"fmt"
)

// checks if the HrefCsrSelfLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HrefCsrSelfLink{}

// HrefCsrSelfLink Link to the resource (self)
type HrefCsrSelfLink struct {
	Hints *CsrSelfHrefHints `json:"hints,omitempty"`
	// Link URI
	Href                 string `json:"href"`
	AdditionalProperties map[string]interface{}
}

type _HrefCsrSelfLink HrefCsrSelfLink

// NewHrefCsrSelfLink instantiates a new HrefCsrSelfLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHrefCsrSelfLink(href string) *HrefCsrSelfLink {
	this := HrefCsrSelfLink{}
	this.Href = href
	return &this
}

// NewHrefCsrSelfLinkWithDefaults instantiates a new HrefCsrSelfLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrefCsrSelfLinkWithDefaults() *HrefCsrSelfLink {
	this := HrefCsrSelfLink{}
	return &this
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *HrefCsrSelfLink) GetHints() CsrSelfHrefHints {
	if o == nil || IsNil(o.Hints) {
		var ret CsrSelfHrefHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefCsrSelfLink) GetHintsOk() (*CsrSelfHrefHints, bool) {
	if o == nil || IsNil(o.Hints) {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *HrefCsrSelfLink) HasHints() bool {
	if o != nil && !IsNil(o.Hints) {
		return true
	}

	return false
}

// SetHints gets a reference to the given CsrSelfHrefHints and assigns it to the Hints field.
func (o *HrefCsrSelfLink) SetHints(v CsrSelfHrefHints) {
	o.Hints = &v
}

// GetHref returns the Href field value
func (o *HrefCsrSelfLink) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *HrefCsrSelfLink) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *HrefCsrSelfLink) SetHref(v string) {
	o.Href = v
}

func (o HrefCsrSelfLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HrefCsrSelfLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hints) {
		toSerialize["hints"] = o.Hints
	}
	toSerialize["href"] = o.Href

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HrefCsrSelfLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHrefCsrSelfLink := _HrefCsrSelfLink{}

	err = json.Unmarshal(data, &varHrefCsrSelfLink)

	if err != nil {
		return err
	}

	*o = HrefCsrSelfLink(varHrefCsrSelfLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hints")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHrefCsrSelfLink struct {
	value *HrefCsrSelfLink
	isSet bool
}

func (v NullableHrefCsrSelfLink) Get() *HrefCsrSelfLink {
	return v.value
}

func (v *NullableHrefCsrSelfLink) Set(val *HrefCsrSelfLink) {
	v.value = val
	v.isSet = true
}

func (v NullableHrefCsrSelfLink) IsSet() bool {
	return v.isSet
}

func (v *NullableHrefCsrSelfLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHrefCsrSelfLink(val *HrefCsrSelfLink) *NullableHrefCsrSelfLink {
	return &NullableHrefCsrSelfLink{value: val, isSet: true}
}

func (v NullableHrefCsrSelfLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHrefCsrSelfLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
