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
)

// checks if the ProfileEnrollmentPolicyRuleProfileAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileEnrollmentPolicyRuleProfileAttribute{}

// ProfileEnrollmentPolicyRuleProfileAttribute struct for ProfileEnrollmentPolicyRuleProfileAttribute
type ProfileEnrollmentPolicyRuleProfileAttribute struct {
	// A display-friendly label for this property
	Label *string `json:"label,omitempty"`
	// The name of a user profile property. Can be an existing property.
	Name *string `json:"name,omitempty"`
	// (Optional, default `FALSE`) Indicates if this property is required for enrollment
	Required             *bool `json:"required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileEnrollmentPolicyRuleProfileAttribute ProfileEnrollmentPolicyRuleProfileAttribute

// NewProfileEnrollmentPolicyRuleProfileAttribute instantiates a new ProfileEnrollmentPolicyRuleProfileAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileEnrollmentPolicyRuleProfileAttribute() *ProfileEnrollmentPolicyRuleProfileAttribute {
	this := ProfileEnrollmentPolicyRuleProfileAttribute{}
	var required bool = false
	this.Required = &required
	return &this
}

// NewProfileEnrollmentPolicyRuleProfileAttributeWithDefaults instantiates a new ProfileEnrollmentPolicyRuleProfileAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileEnrollmentPolicyRuleProfileAttributeWithDefaults() *ProfileEnrollmentPolicyRuleProfileAttribute {
	this := ProfileEnrollmentPolicyRuleProfileAttribute{}
	var required bool = false
	this.Required = &required
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) SetName(v string) {
	o.Name = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) SetRequired(v bool) {
	o.Required = &v
}

func (o ProfileEnrollmentPolicyRuleProfileAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileEnrollmentPolicyRuleProfileAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProfileEnrollmentPolicyRuleProfileAttribute) UnmarshalJSON(data []byte) (err error) {
	varProfileEnrollmentPolicyRuleProfileAttribute := _ProfileEnrollmentPolicyRuleProfileAttribute{}

	err = json.Unmarshal(data, &varProfileEnrollmentPolicyRuleProfileAttribute)

	if err != nil {
		return err
	}

	*o = ProfileEnrollmentPolicyRuleProfileAttribute(varProfileEnrollmentPolicyRuleProfileAttribute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileEnrollmentPolicyRuleProfileAttribute struct {
	value *ProfileEnrollmentPolicyRuleProfileAttribute
	isSet bool
}

func (v NullableProfileEnrollmentPolicyRuleProfileAttribute) Get() *ProfileEnrollmentPolicyRuleProfileAttribute {
	return v.value
}

func (v *NullableProfileEnrollmentPolicyRuleProfileAttribute) Set(val *ProfileEnrollmentPolicyRuleProfileAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileEnrollmentPolicyRuleProfileAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileEnrollmentPolicyRuleProfileAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileEnrollmentPolicyRuleProfileAttribute(val *ProfileEnrollmentPolicyRuleProfileAttribute) *NullableProfileEnrollmentPolicyRuleProfileAttribute {
	return &NullableProfileEnrollmentPolicyRuleProfileAttribute{value: val, isSet: true}
}

func (v NullableProfileEnrollmentPolicyRuleProfileAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileEnrollmentPolicyRuleProfileAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
