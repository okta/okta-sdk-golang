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
	"encoding/json"
)

// checks if the TestInfoApiServiceTestConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestInfoApiServiceTestConfiguration{}

// TestInfoApiServiceTestConfiguration Details required to test the API service integration
type TestInfoApiServiceTestConfiguration struct {
	// Justification for each OAuth 2.0 scope—specifically write scopes—required for the API service integration to function. This information helps the operations team understand the integration and test it effectively.
	ScopeExplanation     *string `json:"scopeExplanation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestInfoApiServiceTestConfiguration TestInfoApiServiceTestConfiguration

// NewTestInfoApiServiceTestConfiguration instantiates a new TestInfoApiServiceTestConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestInfoApiServiceTestConfiguration() *TestInfoApiServiceTestConfiguration {
	this := TestInfoApiServiceTestConfiguration{}
	return &this
}

// NewTestInfoApiServiceTestConfigurationWithDefaults instantiates a new TestInfoApiServiceTestConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestInfoApiServiceTestConfigurationWithDefaults() *TestInfoApiServiceTestConfiguration {
	this := TestInfoApiServiceTestConfiguration{}
	return &this
}

// GetScopeExplanation returns the ScopeExplanation field value if set, zero value otherwise.
func (o *TestInfoApiServiceTestConfiguration) GetScopeExplanation() string {
	if o == nil || IsNil(o.ScopeExplanation) {
		var ret string
		return ret
	}
	return *o.ScopeExplanation
}

// GetScopeExplanationOk returns a tuple with the ScopeExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfoApiServiceTestConfiguration) GetScopeExplanationOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeExplanation) {
		return nil, false
	}
	return o.ScopeExplanation, true
}

// HasScopeExplanation returns a boolean if a field has been set.
func (o *TestInfoApiServiceTestConfiguration) HasScopeExplanation() bool {
	if o != nil && !IsNil(o.ScopeExplanation) {
		return true
	}

	return false
}

// SetScopeExplanation gets a reference to the given string and assigns it to the ScopeExplanation field.
func (o *TestInfoApiServiceTestConfiguration) SetScopeExplanation(v string) {
	o.ScopeExplanation = &v
}

func (o TestInfoApiServiceTestConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestInfoApiServiceTestConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScopeExplanation) {
		toSerialize["scopeExplanation"] = o.ScopeExplanation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestInfoApiServiceTestConfiguration) UnmarshalJSON(data []byte) (err error) {
	varTestInfoApiServiceTestConfiguration := _TestInfoApiServiceTestConfiguration{}

	err = json.Unmarshal(data, &varTestInfoApiServiceTestConfiguration)

	if err != nil {
		return err
	}

	*o = TestInfoApiServiceTestConfiguration(varTestInfoApiServiceTestConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "scopeExplanation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestInfoApiServiceTestConfiguration struct {
	value *TestInfoApiServiceTestConfiguration
	isSet bool
}

func (v NullableTestInfoApiServiceTestConfiguration) Get() *TestInfoApiServiceTestConfiguration {
	return v.value
}

func (v *NullableTestInfoApiServiceTestConfiguration) Set(val *TestInfoApiServiceTestConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInfoApiServiceTestConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInfoApiServiceTestConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInfoApiServiceTestConfiguration(val *TestInfoApiServiceTestConfiguration) *NullableTestInfoApiServiceTestConfiguration {
	return &NullableTestInfoApiServiceTestConfiguration{value: val, isSet: true}
}

func (v NullableTestInfoApiServiceTestConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInfoApiServiceTestConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
