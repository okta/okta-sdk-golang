/*
Okta Admin Management API

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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the WidgetConfigurationHelpLinksFactorPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetConfigurationHelpLinksFactorPage{}

// WidgetConfigurationHelpLinksFactorPage struct for WidgetConfigurationHelpLinksFactorPage
type WidgetConfigurationHelpLinksFactorPage struct {
	Text   string  `json:"text"`
	Href   string  `json:"href" validate:"regexp=^https?:\\/\\/"`
	Target *string `json:"target,omitempty"`
}

type _WidgetConfigurationHelpLinksFactorPage WidgetConfigurationHelpLinksFactorPage

// NewWidgetConfigurationHelpLinksFactorPage instantiates a new WidgetConfigurationHelpLinksFactorPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetConfigurationHelpLinksFactorPage(text string, href string) *WidgetConfigurationHelpLinksFactorPage {
	this := WidgetConfigurationHelpLinksFactorPage{}
	this.Text = text
	this.Href = href
	return &this
}

// NewWidgetConfigurationHelpLinksFactorPageWithDefaults instantiates a new WidgetConfigurationHelpLinksFactorPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetConfigurationHelpLinksFactorPageWithDefaults() *WidgetConfigurationHelpLinksFactorPage {
	this := WidgetConfigurationHelpLinksFactorPage{}
	return &this
}

// GetText returns the Text field value
func (o *WidgetConfigurationHelpLinksFactorPage) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationHelpLinksFactorPage) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *WidgetConfigurationHelpLinksFactorPage) SetText(v string) {
	o.Text = v
}

// GetHref returns the Href field value
func (o *WidgetConfigurationHelpLinksFactorPage) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationHelpLinksFactorPage) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *WidgetConfigurationHelpLinksFactorPage) SetHref(v string) {
	o.Href = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *WidgetConfigurationHelpLinksFactorPage) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationHelpLinksFactorPage) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *WidgetConfigurationHelpLinksFactorPage) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *WidgetConfigurationHelpLinksFactorPage) SetTarget(v string) {
	o.Target = &v
}

func (o WidgetConfigurationHelpLinksFactorPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetConfigurationHelpLinksFactorPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["text"] = o.Text
	toSerialize["href"] = o.Href
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

func (o *WidgetConfigurationHelpLinksFactorPage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"text",
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

	varWidgetConfigurationHelpLinksFactorPage := _WidgetConfigurationHelpLinksFactorPage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWidgetConfigurationHelpLinksFactorPage)

	if err != nil {
		return err
	}

	*o = WidgetConfigurationHelpLinksFactorPage(varWidgetConfigurationHelpLinksFactorPage)

	return err
}

type NullableWidgetConfigurationHelpLinksFactorPage struct {
	value *WidgetConfigurationHelpLinksFactorPage
	isSet bool
}

func (v NullableWidgetConfigurationHelpLinksFactorPage) Get() *WidgetConfigurationHelpLinksFactorPage {
	return v.value
}

func (v *NullableWidgetConfigurationHelpLinksFactorPage) Set(val *WidgetConfigurationHelpLinksFactorPage) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetConfigurationHelpLinksFactorPage) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetConfigurationHelpLinksFactorPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetConfigurationHelpLinksFactorPage(val *WidgetConfigurationHelpLinksFactorPage) *NullableWidgetConfigurationHelpLinksFactorPage {
	return &NullableWidgetConfigurationHelpLinksFactorPage{value: val, isSet: true}
}

func (v NullableWidgetConfigurationHelpLinksFactorPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetConfigurationHelpLinksFactorPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
