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

// SubmissionResponseConfigInner struct for SubmissionResponseConfigInner
type SubmissionResponseConfigInner struct {
	// Display name of the variable in the Admin Console
	Label *string `json:"label,omitempty"`
	// Name of the variable
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubmissionResponseConfigInner SubmissionResponseConfigInner

// NewSubmissionResponseConfigInner instantiates a new SubmissionResponseConfigInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionResponseConfigInner() *SubmissionResponseConfigInner {
	this := SubmissionResponseConfigInner{}
	return &this
}

// NewSubmissionResponseConfigInnerWithDefaults instantiates a new SubmissionResponseConfigInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionResponseConfigInnerWithDefaults() *SubmissionResponseConfigInner {
	this := SubmissionResponseConfigInner{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SubmissionResponseConfigInner) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseConfigInner) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SubmissionResponseConfigInner) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SubmissionResponseConfigInner) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubmissionResponseConfigInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseConfigInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubmissionResponseConfigInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubmissionResponseConfigInner) SetName(v string) {
	o.Name = &v
}

func (o SubmissionResponseConfigInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SubmissionResponseConfigInner) UnmarshalJSON(bytes []byte) (err error) {
	varSubmissionResponseConfigInner := _SubmissionResponseConfigInner{}

	err = json.Unmarshal(bytes, &varSubmissionResponseConfigInner)
	if err == nil {
		*o = SubmissionResponseConfigInner(varSubmissionResponseConfigInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSubmissionResponseConfigInner struct {
	value *SubmissionResponseConfigInner
	isSet bool
}

func (v NullableSubmissionResponseConfigInner) Get() *SubmissionResponseConfigInner {
	return v.value
}

func (v *NullableSubmissionResponseConfigInner) Set(val *SubmissionResponseConfigInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionResponseConfigInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionResponseConfigInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionResponseConfigInner(val *SubmissionResponseConfigInner) *NullableSubmissionResponseConfigInner {
	return &NullableSubmissionResponseConfigInner{value: val, isSet: true}
}

func (v NullableSubmissionResponseConfigInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionResponseConfigInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

