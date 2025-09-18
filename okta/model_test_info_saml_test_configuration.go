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

// checks if the TestInfoSamlTestConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestInfoSamlTestConfiguration{}

// TestInfoSamlTestConfiguration SAML test details
type TestInfoSamlTestConfiguration struct {
	// Indicates if your integration supports IdP-initiated sign-in
	Idp *bool `json:"idp,omitempty"`
	// Indicates if your integration supports SP-initiated sign-in
	Sp *bool `json:"sp,omitempty"`
	// Indicates if your integration supports Just-In-Time (JIT) provisioning
	Jit *bool `json:"jit,omitempty"`
	// URL for SP-initiated sign-in flows (required if `sp = true`)
	SpInitiateUrl string `json:"spInitiateUrl"`
	// Instructions on how to sign in to your app using the SP-initiated flow (required if `sp = true`)
	SpInitiateDescription *string `json:"spInitiateDescription,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _TestInfoSamlTestConfiguration TestInfoSamlTestConfiguration

// NewTestInfoSamlTestConfiguration instantiates a new TestInfoSamlTestConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestInfoSamlTestConfiguration(spInitiateUrl string) *TestInfoSamlTestConfiguration {
	this := TestInfoSamlTestConfiguration{}
	this.SpInitiateUrl = spInitiateUrl
	return &this
}

// NewTestInfoSamlTestConfigurationWithDefaults instantiates a new TestInfoSamlTestConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestInfoSamlTestConfigurationWithDefaults() *TestInfoSamlTestConfiguration {
	this := TestInfoSamlTestConfiguration{}
	return &this
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *TestInfoSamlTestConfiguration) GetIdp() bool {
	if o == nil || IsNil(o.Idp) {
		var ret bool
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfoSamlTestConfiguration) GetIdpOk() (*bool, bool) {
	if o == nil || IsNil(o.Idp) {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *TestInfoSamlTestConfiguration) HasIdp() bool {
	if o != nil && !IsNil(o.Idp) {
		return true
	}

	return false
}

// SetIdp gets a reference to the given bool and assigns it to the Idp field.
func (o *TestInfoSamlTestConfiguration) SetIdp(v bool) {
	o.Idp = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *TestInfoSamlTestConfiguration) GetSp() bool {
	if o == nil || IsNil(o.Sp) {
		var ret bool
		return ret
	}
	return *o.Sp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfoSamlTestConfiguration) GetSpOk() (*bool, bool) {
	if o == nil || IsNil(o.Sp) {
		return nil, false
	}
	return o.Sp, true
}

// HasSp returns a boolean if a field has been set.
func (o *TestInfoSamlTestConfiguration) HasSp() bool {
	if o != nil && !IsNil(o.Sp) {
		return true
	}

	return false
}

// SetSp gets a reference to the given bool and assigns it to the Sp field.
func (o *TestInfoSamlTestConfiguration) SetSp(v bool) {
	o.Sp = &v
}

// GetJit returns the Jit field value if set, zero value otherwise.
func (o *TestInfoSamlTestConfiguration) GetJit() bool {
	if o == nil || IsNil(o.Jit) {
		var ret bool
		return ret
	}
	return *o.Jit
}

// GetJitOk returns a tuple with the Jit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfoSamlTestConfiguration) GetJitOk() (*bool, bool) {
	if o == nil || IsNil(o.Jit) {
		return nil, false
	}
	return o.Jit, true
}

// HasJit returns a boolean if a field has been set.
func (o *TestInfoSamlTestConfiguration) HasJit() bool {
	if o != nil && !IsNil(o.Jit) {
		return true
	}

	return false
}

// SetJit gets a reference to the given bool and assigns it to the Jit field.
func (o *TestInfoSamlTestConfiguration) SetJit(v bool) {
	o.Jit = &v
}

// GetSpInitiateUrl returns the SpInitiateUrl field value
func (o *TestInfoSamlTestConfiguration) GetSpInitiateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpInitiateUrl
}

// GetSpInitiateUrlOk returns a tuple with the SpInitiateUrl field value
// and a boolean to check if the value has been set.
func (o *TestInfoSamlTestConfiguration) GetSpInitiateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpInitiateUrl, true
}

// SetSpInitiateUrl sets field value
func (o *TestInfoSamlTestConfiguration) SetSpInitiateUrl(v string) {
	o.SpInitiateUrl = v
}

// GetSpInitiateDescription returns the SpInitiateDescription field value if set, zero value otherwise.
func (o *TestInfoSamlTestConfiguration) GetSpInitiateDescription() string {
	if o == nil || IsNil(o.SpInitiateDescription) {
		var ret string
		return ret
	}
	return *o.SpInitiateDescription
}

// GetSpInitiateDescriptionOk returns a tuple with the SpInitiateDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfoSamlTestConfiguration) GetSpInitiateDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SpInitiateDescription) {
		return nil, false
	}
	return o.SpInitiateDescription, true
}

// HasSpInitiateDescription returns a boolean if a field has been set.
func (o *TestInfoSamlTestConfiguration) HasSpInitiateDescription() bool {
	if o != nil && !IsNil(o.SpInitiateDescription) {
		return true
	}

	return false
}

// SetSpInitiateDescription gets a reference to the given string and assigns it to the SpInitiateDescription field.
func (o *TestInfoSamlTestConfiguration) SetSpInitiateDescription(v string) {
	o.SpInitiateDescription = &v
}

func (o TestInfoSamlTestConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestInfoSamlTestConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Idp) {
		toSerialize["idp"] = o.Idp
	}
	if !IsNil(o.Sp) {
		toSerialize["sp"] = o.Sp
	}
	if !IsNil(o.Jit) {
		toSerialize["jit"] = o.Jit
	}
	toSerialize["spInitiateUrl"] = o.SpInitiateUrl
	if !IsNil(o.SpInitiateDescription) {
		toSerialize["spInitiateDescription"] = o.SpInitiateDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestInfoSamlTestConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"spInitiateUrl",
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

	varTestInfoSamlTestConfiguration := _TestInfoSamlTestConfiguration{}

	err = json.Unmarshal(data, &varTestInfoSamlTestConfiguration)

	if err != nil {
		return err
	}

	*o = TestInfoSamlTestConfiguration(varTestInfoSamlTestConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "idp")
		delete(additionalProperties, "sp")
		delete(additionalProperties, "jit")
		delete(additionalProperties, "spInitiateUrl")
		delete(additionalProperties, "spInitiateDescription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestInfoSamlTestConfiguration struct {
	value *TestInfoSamlTestConfiguration
	isSet bool
}

func (v NullableTestInfoSamlTestConfiguration) Get() *TestInfoSamlTestConfiguration {
	return v.value
}

func (v *NullableTestInfoSamlTestConfiguration) Set(val *TestInfoSamlTestConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInfoSamlTestConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInfoSamlTestConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInfoSamlTestConfiguration(val *TestInfoSamlTestConfiguration) *NullableTestInfoSamlTestConfiguration {
	return &NullableTestInfoSamlTestConfiguration{value: val, isSet: true}
}

func (v NullableTestInfoSamlTestConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInfoSamlTestConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
