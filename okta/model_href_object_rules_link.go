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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the HrefObjectRulesLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HrefObjectRulesLink{}

// HrefObjectRulesLink Link to the rules resource
type HrefObjectRulesLink struct {
	Hints *HrefHints `json:"hints,omitempty"`
	// Link URI
	Href string `json:"href"`
	// Link name
	Name *string `json:"name,omitempty"`
	// Indicates whether the link object's `href` property is a URI template.
	Templated *bool `json:"templated,omitempty"`
	// The media type of the link. If omitted, it is implicitly `application/json`.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HrefObjectRulesLink HrefObjectRulesLink

// NewHrefObjectRulesLink instantiates a new HrefObjectRulesLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHrefObjectRulesLink(href string) *HrefObjectRulesLink {
	this := HrefObjectRulesLink{}
	this.Href = href
	return &this
}

// NewHrefObjectRulesLinkWithDefaults instantiates a new HrefObjectRulesLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrefObjectRulesLinkWithDefaults() *HrefObjectRulesLink {
	this := HrefObjectRulesLink{}
	return &this
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *HrefObjectRulesLink) GetHints() HrefHints {
	if o == nil || IsNil(o.Hints) {
		var ret HrefHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefObjectRulesLink) GetHintsOk() (*HrefHints, bool) {
	if o == nil || IsNil(o.Hints) {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *HrefObjectRulesLink) HasHints() bool {
	if o != nil && !IsNil(o.Hints) {
		return true
	}

	return false
}

// SetHints gets a reference to the given HrefHints and assigns it to the Hints field.
func (o *HrefObjectRulesLink) SetHints(v HrefHints) {
	o.Hints = &v
}

// GetHref returns the Href field value
func (o *HrefObjectRulesLink) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *HrefObjectRulesLink) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *HrefObjectRulesLink) SetHref(v string) {
	o.Href = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HrefObjectRulesLink) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefObjectRulesLink) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HrefObjectRulesLink) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HrefObjectRulesLink) SetName(v string) {
	o.Name = &v
}

// GetTemplated returns the Templated field value if set, zero value otherwise.
func (o *HrefObjectRulesLink) GetTemplated() bool {
	if o == nil || IsNil(o.Templated) {
		var ret bool
		return ret
	}
	return *o.Templated
}

// GetTemplatedOk returns a tuple with the Templated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefObjectRulesLink) GetTemplatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Templated) {
		return nil, false
	}
	return o.Templated, true
}

// HasTemplated returns a boolean if a field has been set.
func (o *HrefObjectRulesLink) HasTemplated() bool {
	if o != nil && !IsNil(o.Templated) {
		return true
	}

	return false
}

// SetTemplated gets a reference to the given bool and assigns it to the Templated field.
func (o *HrefObjectRulesLink) SetTemplated(v bool) {
	o.Templated = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *HrefObjectRulesLink) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefObjectRulesLink) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *HrefObjectRulesLink) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *HrefObjectRulesLink) SetType(v string) {
	o.Type = &v
}

func (o HrefObjectRulesLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HrefObjectRulesLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hints) {
		toSerialize["hints"] = o.Hints
	}
	toSerialize["href"] = o.Href
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Templated) {
		toSerialize["templated"] = o.Templated
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HrefObjectRulesLink) UnmarshalJSON(data []byte) (err error) {
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

	varHrefObjectRulesLink := _HrefObjectRulesLink{}

	err = json.Unmarshal(data, &varHrefObjectRulesLink)

	if err != nil {
		return err
	}

	*o = HrefObjectRulesLink(varHrefObjectRulesLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hints")
		delete(additionalProperties, "href")
		delete(additionalProperties, "name")
		delete(additionalProperties, "templated")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHrefObjectRulesLink struct {
	value *HrefObjectRulesLink
	isSet bool
}

func (v NullableHrefObjectRulesLink) Get() *HrefObjectRulesLink {
	return v.value
}

func (v *NullableHrefObjectRulesLink) Set(val *HrefObjectRulesLink) {
	v.value = val
	v.isSet = true
}

func (v NullableHrefObjectRulesLink) IsSet() bool {
	return v.isSet
}

func (v *NullableHrefObjectRulesLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHrefObjectRulesLink(val *HrefObjectRulesLink) *NullableHrefObjectRulesLink {
	return &NullableHrefObjectRulesLink{value: val, isSet: true}
}

func (v NullableHrefObjectRulesLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHrefObjectRulesLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
