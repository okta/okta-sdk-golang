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
	"fmt"
)

// checks if the TestInfoScimTestConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestInfoScimTestConfiguration{}

// TestInfoScimTestConfiguration SCIM test details
type TestInfoScimTestConfiguration struct {
	// The Runscope URL to your SCIM server specification test results. See [Test your SCIM API](https://developer.okta.com/docs/guides/build-provisioning-integration/test-scim-api/).
	SpecTestResults string `json:"specTestResults"`
	// The Runscope URL to your Okta SCIM CRUD test results. See [Test your Okta SCIM integration](https://developer.okta.com/docs/guides/scim-provisioning-integration-test/main/).
	CrudTestResults      string `json:"crudTestResults"`
	AdditionalProperties map[string]interface{}
}

type _TestInfoScimTestConfiguration TestInfoScimTestConfiguration

// NewTestInfoScimTestConfiguration instantiates a new TestInfoScimTestConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestInfoScimTestConfiguration(specTestResults string, crudTestResults string) *TestInfoScimTestConfiguration {
	this := TestInfoScimTestConfiguration{}
	this.SpecTestResults = specTestResults
	this.CrudTestResults = crudTestResults
	return &this
}

// NewTestInfoScimTestConfigurationWithDefaults instantiates a new TestInfoScimTestConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestInfoScimTestConfigurationWithDefaults() *TestInfoScimTestConfiguration {
	this := TestInfoScimTestConfiguration{}
	return &this
}

// GetSpecTestResults returns the SpecTestResults field value
func (o *TestInfoScimTestConfiguration) GetSpecTestResults() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecTestResults
}

// GetSpecTestResultsOk returns a tuple with the SpecTestResults field value
// and a boolean to check if the value has been set.
func (o *TestInfoScimTestConfiguration) GetSpecTestResultsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecTestResults, true
}

// SetSpecTestResults sets field value
func (o *TestInfoScimTestConfiguration) SetSpecTestResults(v string) {
	o.SpecTestResults = v
}

// GetCrudTestResults returns the CrudTestResults field value
func (o *TestInfoScimTestConfiguration) GetCrudTestResults() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CrudTestResults
}

// GetCrudTestResultsOk returns a tuple with the CrudTestResults field value
// and a boolean to check if the value has been set.
func (o *TestInfoScimTestConfiguration) GetCrudTestResultsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CrudTestResults, true
}

// SetCrudTestResults sets field value
func (o *TestInfoScimTestConfiguration) SetCrudTestResults(v string) {
	o.CrudTestResults = v
}

func (o TestInfoScimTestConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestInfoScimTestConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["specTestResults"] = o.SpecTestResults
	toSerialize["crudTestResults"] = o.CrudTestResults

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestInfoScimTestConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"specTestResults",
		"crudTestResults",
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

	varTestInfoScimTestConfiguration := _TestInfoScimTestConfiguration{}

	err = json.Unmarshal(data, &varTestInfoScimTestConfiguration)

	if err != nil {
		return err
	}

	*o = TestInfoScimTestConfiguration(varTestInfoScimTestConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "specTestResults")
		delete(additionalProperties, "crudTestResults")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestInfoScimTestConfiguration struct {
	value *TestInfoScimTestConfiguration
	isSet bool
}

func (v NullableTestInfoScimTestConfiguration) Get() *TestInfoScimTestConfiguration {
	return v.value
}

func (v *NullableTestInfoScimTestConfiguration) Set(val *TestInfoScimTestConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInfoScimTestConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInfoScimTestConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInfoScimTestConfiguration(val *TestInfoScimTestConfiguration) *NullableTestInfoScimTestConfiguration {
	return &NullableTestInfoScimTestConfiguration{value: val, isSet: true}
}

func (v NullableTestInfoScimTestConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInfoScimTestConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
