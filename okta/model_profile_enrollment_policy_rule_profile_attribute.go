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

// ProfileEnrollmentPolicyRuleProfileAttribute struct for ProfileEnrollmentPolicyRuleProfileAttribute
type ProfileEnrollmentPolicyRuleProfileAttribute struct {
	Label *string `json:"label,omitempty"`
	Name *string `json:"name,omitempty"`
	Required *bool `json:"required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileEnrollmentPolicyRuleProfileAttribute ProfileEnrollmentPolicyRuleProfileAttribute

// NewProfileEnrollmentPolicyRuleProfileAttribute instantiates a new ProfileEnrollmentPolicyRuleProfileAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileEnrollmentPolicyRuleProfileAttribute() *ProfileEnrollmentPolicyRuleProfileAttribute {
	this := ProfileEnrollmentPolicyRuleProfileAttribute{}
	return &this
}

// NewProfileEnrollmentPolicyRuleProfileAttributeWithDefaults instantiates a new ProfileEnrollmentPolicyRuleProfileAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileEnrollmentPolicyRuleProfileAttributeWithDefaults() *ProfileEnrollmentPolicyRuleProfileAttribute {
	this := ProfileEnrollmentPolicyRuleProfileAttribute{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) HasLabel() bool {
	if o != nil && o.Label != nil {
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
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) HasName() bool {
	if o != nil && o.Name != nil {
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
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ProfileEnrollmentPolicyRuleProfileAttribute) SetRequired(v bool) {
	o.Required = &v
}

func (o ProfileEnrollmentPolicyRuleProfileAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProfileEnrollmentPolicyRuleProfileAttribute) UnmarshalJSON(bytes []byte) (err error) {
	varProfileEnrollmentPolicyRuleProfileAttribute := _ProfileEnrollmentPolicyRuleProfileAttribute{}

	err = json.Unmarshal(bytes, &varProfileEnrollmentPolicyRuleProfileAttribute)
	if err == nil {
		*o = ProfileEnrollmentPolicyRuleProfileAttribute(varProfileEnrollmentPolicyRuleProfileAttribute)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

