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
)

// checks if the UIElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UIElement{}

// UIElement Specifies the configuration of an input field on an enrollment form
type UIElement struct {
	// Label name for the UI element
	Label   *string           `json:"label,omitempty"`
	Options *UIElementOptions `json:"options,omitempty"`
	// Specifies the property bound to the input field. It must follow the format `#/properties/PROPERTY_NAME` where `PROPERTY_NAME` is a variable name for an attribute in `profile editor`.
	Scope *string `json:"scope,omitempty"`
	// Specifies the relationship between this input element and `scope`. The `Control` value specifies that this input controls the value represented by `scope`.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UIElement UIElement

// NewUIElement instantiates a new UIElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUIElement() *UIElement {
	this := UIElement{}
	return &this
}

// NewUIElementWithDefaults instantiates a new UIElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUIElementWithDefaults() *UIElement {
	this := UIElement{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UIElement) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UIElement) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UIElement) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *UIElement) SetLabel(v string) {
	o.Label = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *UIElement) GetOptions() UIElementOptions {
	if o == nil || IsNil(o.Options) {
		var ret UIElementOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UIElement) GetOptionsOk() (*UIElementOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *UIElement) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given UIElementOptions and assigns it to the Options field.
func (o *UIElement) SetOptions(v UIElementOptions) {
	o.Options = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *UIElement) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UIElement) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *UIElement) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *UIElement) SetScope(v string) {
	o.Scope = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UIElement) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UIElement) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UIElement) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UIElement) SetType(v string) {
	o.Type = &v
}

func (o UIElement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UIElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UIElement) UnmarshalJSON(data []byte) (err error) {
	varUIElement := _UIElement{}

	err = json.Unmarshal(data, &varUIElement)

	if err != nil {
		return err
	}

	*o = UIElement(varUIElement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "options")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUIElement struct {
	value *UIElement
	isSet bool
}

func (v NullableUIElement) Get() *UIElement {
	return v.value
}

func (v *NullableUIElement) Set(val *UIElement) {
	v.value = val
	v.isSet = true
}

func (v NullableUIElement) IsSet() bool {
	return v.isSet
}

func (v *NullableUIElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUIElement(val *UIElement) *NullableUIElement {
	return &NullableUIElement{value: val, isSet: true}
}

func (v NullableUIElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUIElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
