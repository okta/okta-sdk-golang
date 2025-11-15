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

// checks if the UISchemaObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UISchemaObject{}

// UISchemaObject Properties of the UI schema
type UISchemaObject struct {
	// Specifies the button label for the `Submit` button at the bottom of the enrollment form
	ButtonLabel *string     `json:"buttonLabel,omitempty"`
	Elements    []UIElement `json:"elements,omitempty"`
	// Specifies the label at the top of the enrollment form under the logo
	Label *string `json:"label,omitempty"`
	// Specifies the type of layout
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UISchemaObject UISchemaObject

// NewUISchemaObject instantiates a new UISchemaObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUISchemaObject() *UISchemaObject {
	this := UISchemaObject{}
	var buttonLabel string = "Submit"
	this.ButtonLabel = &buttonLabel
	var label string = "Sign in"
	this.Label = &label
	return &this
}

// NewUISchemaObjectWithDefaults instantiates a new UISchemaObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUISchemaObjectWithDefaults() *UISchemaObject {
	this := UISchemaObject{}
	var buttonLabel string = "Submit"
	this.ButtonLabel = &buttonLabel
	var label string = "Sign in"
	this.Label = &label
	return &this
}

// GetButtonLabel returns the ButtonLabel field value if set, zero value otherwise.
func (o *UISchemaObject) GetButtonLabel() string {
	if o == nil || IsNil(o.ButtonLabel) {
		var ret string
		return ret
	}
	return *o.ButtonLabel
}

// GetButtonLabelOk returns a tuple with the ButtonLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UISchemaObject) GetButtonLabelOk() (*string, bool) {
	if o == nil || IsNil(o.ButtonLabel) {
		return nil, false
	}
	return o.ButtonLabel, true
}

// HasButtonLabel returns a boolean if a field has been set.
func (o *UISchemaObject) HasButtonLabel() bool {
	if o != nil && !IsNil(o.ButtonLabel) {
		return true
	}

	return false
}

// SetButtonLabel gets a reference to the given string and assigns it to the ButtonLabel field.
func (o *UISchemaObject) SetButtonLabel(v string) {
	o.ButtonLabel = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *UISchemaObject) GetElements() []UIElement {
	if o == nil || IsNil(o.Elements) {
		var ret []UIElement
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UISchemaObject) GetElementsOk() ([]UIElement, bool) {
	if o == nil || IsNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *UISchemaObject) HasElements() bool {
	if o != nil && !IsNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []UIElement and assigns it to the Elements field.
func (o *UISchemaObject) SetElements(v []UIElement) {
	o.Elements = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UISchemaObject) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UISchemaObject) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UISchemaObject) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *UISchemaObject) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UISchemaObject) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UISchemaObject) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UISchemaObject) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UISchemaObject) SetType(v string) {
	o.Type = &v
}

func (o UISchemaObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UISchemaObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ButtonLabel) {
		toSerialize["buttonLabel"] = o.ButtonLabel
	}
	if !IsNil(o.Elements) {
		toSerialize["elements"] = o.Elements
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UISchemaObject) UnmarshalJSON(data []byte) (err error) {
	varUISchemaObject := _UISchemaObject{}

	err = json.Unmarshal(data, &varUISchemaObject)

	if err != nil {
		return err
	}

	*o = UISchemaObject(varUISchemaObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "buttonLabel")
		delete(additionalProperties, "elements")
		delete(additionalProperties, "label")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUISchemaObject struct {
	value *UISchemaObject
	isSet bool
}

func (v NullableUISchemaObject) Get() *UISchemaObject {
	return v.value
}

func (v *NullableUISchemaObject) Set(val *UISchemaObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUISchemaObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUISchemaObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUISchemaObject(val *UISchemaObject) *NullableUISchemaObject {
	return &NullableUISchemaObject{value: val, isSet: true}
}

func (v NullableUISchemaObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUISchemaObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
