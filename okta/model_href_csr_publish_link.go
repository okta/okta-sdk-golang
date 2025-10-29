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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the HrefCsrPublishLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HrefCsrPublishLink{}

// HrefCsrPublishLink Link to publish CSR
type HrefCsrPublishLink struct {
	Hints *CsrPublishHrefHints `json:"hints,omitempty"`
	// Link URI
	Href                 string `json:"href"`
	AdditionalProperties map[string]interface{}
}

type _HrefCsrPublishLink HrefCsrPublishLink

// NewHrefCsrPublishLink instantiates a new HrefCsrPublishLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHrefCsrPublishLink(href string) *HrefCsrPublishLink {
	this := HrefCsrPublishLink{}
	this.Href = href
	return &this
}

// NewHrefCsrPublishLinkWithDefaults instantiates a new HrefCsrPublishLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrefCsrPublishLinkWithDefaults() *HrefCsrPublishLink {
	this := HrefCsrPublishLink{}
	return &this
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *HrefCsrPublishLink) GetHints() CsrPublishHrefHints {
	if o == nil || IsNil(o.Hints) {
		var ret CsrPublishHrefHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefCsrPublishLink) GetHintsOk() (*CsrPublishHrefHints, bool) {
	if o == nil || IsNil(o.Hints) {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *HrefCsrPublishLink) HasHints() bool {
	if o != nil && !IsNil(o.Hints) {
		return true
	}

	return false
}

// SetHints gets a reference to the given CsrPublishHrefHints and assigns it to the Hints field.
func (o *HrefCsrPublishLink) SetHints(v CsrPublishHrefHints) {
	o.Hints = &v
}

// GetHref returns the Href field value
func (o *HrefCsrPublishLink) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *HrefCsrPublishLink) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *HrefCsrPublishLink) SetHref(v string) {
	o.Href = v
}

func (o HrefCsrPublishLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HrefCsrPublishLink) ToMap() (map[string]interface{}, error) {
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

func (o *HrefCsrPublishLink) UnmarshalJSON(data []byte) (err error) {
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

	varHrefCsrPublishLink := _HrefCsrPublishLink{}

	err = json.Unmarshal(data, &varHrefCsrPublishLink)

	if err != nil {
		return err
	}

	*o = HrefCsrPublishLink(varHrefCsrPublishLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hints")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHrefCsrPublishLink struct {
	value *HrefCsrPublishLink
	isSet bool
}

func (v NullableHrefCsrPublishLink) Get() *HrefCsrPublishLink {
	return v.value
}

func (v *NullableHrefCsrPublishLink) Set(val *HrefCsrPublishLink) {
	v.value = val
	v.isSet = true
}

func (v NullableHrefCsrPublishLink) IsSet() bool {
	return v.isSet
}

func (v *NullableHrefCsrPublishLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHrefCsrPublishLink(val *HrefCsrPublishLink) *NullableHrefCsrPublishLink {
	return &NullableHrefCsrPublishLink{value: val, isSet: true}
}

func (v NullableHrefCsrPublishLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHrefCsrPublishLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
