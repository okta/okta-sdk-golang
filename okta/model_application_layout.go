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

// ApplicationLayout struct for ApplicationLayout
type ApplicationLayout struct {
	Elements []map[string]interface{} `json:"elements,omitempty"`
	Label *string `json:"label,omitempty"`
	Options map[string]interface{} `json:"options,omitempty"`
	Rule *ApplicationLayoutRule `json:"rule,omitempty"`
	Scope *string `json:"scope,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationLayout ApplicationLayout

// NewApplicationLayout instantiates a new ApplicationLayout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLayout() *ApplicationLayout {
	this := ApplicationLayout{}
	return &this
}

// NewApplicationLayoutWithDefaults instantiates a new ApplicationLayout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLayoutWithDefaults() *ApplicationLayout {
	this := ApplicationLayout{}
	return &this
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *ApplicationLayout) GetElements() []map[string]interface{} {
	if o == nil || o.Elements == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayout) GetElementsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *ApplicationLayout) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []map[string]interface{} and assigns it to the Elements field.
func (o *ApplicationLayout) SetElements(v []map[string]interface{}) {
	o.Elements = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ApplicationLayout) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayout) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ApplicationLayout) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ApplicationLayout) SetLabel(v string) {
	o.Label = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ApplicationLayout) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayout) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApplicationLayout) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *ApplicationLayout) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *ApplicationLayout) GetRule() ApplicationLayoutRule {
	if o == nil || o.Rule == nil {
		var ret ApplicationLayoutRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayout) GetRuleOk() (*ApplicationLayoutRule, bool) {
	if o == nil || o.Rule == nil {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *ApplicationLayout) HasRule() bool {
	if o != nil && o.Rule != nil {
		return true
	}

	return false
}

// SetRule gets a reference to the given ApplicationLayoutRule and assigns it to the Rule field.
func (o *ApplicationLayout) SetRule(v ApplicationLayoutRule) {
	o.Rule = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ApplicationLayout) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayout) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ApplicationLayout) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ApplicationLayout) SetScope(v string) {
	o.Scope = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApplicationLayout) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayout) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationLayout) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApplicationLayout) SetType(v string) {
	o.Type = &v
}

func (o ApplicationLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Rule != nil {
		toSerialize["rule"] = o.Rule
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationLayout) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationLayout := _ApplicationLayout{}

	err = json.Unmarshal(bytes, &varApplicationLayout)
	if err == nil {
		*o = ApplicationLayout(varApplicationLayout)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "elements")
		delete(additionalProperties, "label")
		delete(additionalProperties, "options")
		delete(additionalProperties, "rule")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationLayout struct {
	value *ApplicationLayout
	isSet bool
}

func (v NullableApplicationLayout) Get() *ApplicationLayout {
	return v.value
}

func (v *NullableApplicationLayout) Set(val *ApplicationLayout) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLayout(val *ApplicationLayout) *NullableApplicationLayout {
	return &NullableApplicationLayout{value: val, isSet: true}
}

func (v NullableApplicationLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

