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
	"fmt"
)

// checks if the TestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestInfo{}

// TestInfo Integration Testing Information
type TestInfo struct {
	// An email for Okta to contact your company about your integration. This email isn't shared with customers.
	EscalationSupportContact string                         `json:"escalationSupportContact"`
	OidcTestConfiguration    *TestInfoOidcTestConfiguration `json:"oidcTestConfiguration,omitempty"`
	SamlTestConfiguration    *TestInfoSamlTestConfiguration `json:"samlTestConfiguration,omitempty"`
	ScimTestConfiguration    *TestInfoScimTestConfiguration `json:"scimTestConfiguration,omitempty"`
	TestAccount              *TestInfoTestAccount           `json:"testAccount,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _TestInfo TestInfo

// NewTestInfo instantiates a new TestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestInfo(escalationSupportContact string) *TestInfo {
	this := TestInfo{}
	this.EscalationSupportContact = escalationSupportContact
	return &this
}

// NewTestInfoWithDefaults instantiates a new TestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestInfoWithDefaults() *TestInfo {
	this := TestInfo{}
	return &this
}

// GetEscalationSupportContact returns the EscalationSupportContact field value
func (o *TestInfo) GetEscalationSupportContact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EscalationSupportContact
}

// GetEscalationSupportContactOk returns a tuple with the EscalationSupportContact field value
// and a boolean to check if the value has been set.
func (o *TestInfo) GetEscalationSupportContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EscalationSupportContact, true
}

// SetEscalationSupportContact sets field value
func (o *TestInfo) SetEscalationSupportContact(v string) {
	o.EscalationSupportContact = v
}

// GetOidcTestConfiguration returns the OidcTestConfiguration field value if set, zero value otherwise.
func (o *TestInfo) GetOidcTestConfiguration() TestInfoOidcTestConfiguration {
	if o == nil || IsNil(o.OidcTestConfiguration) {
		var ret TestInfoOidcTestConfiguration
		return ret
	}
	return *o.OidcTestConfiguration
}

// GetOidcTestConfigurationOk returns a tuple with the OidcTestConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfo) GetOidcTestConfigurationOk() (*TestInfoOidcTestConfiguration, bool) {
	if o == nil || IsNil(o.OidcTestConfiguration) {
		return nil, false
	}
	return o.OidcTestConfiguration, true
}

// HasOidcTestConfiguration returns a boolean if a field has been set.
func (o *TestInfo) HasOidcTestConfiguration() bool {
	if o != nil && !IsNil(o.OidcTestConfiguration) {
		return true
	}

	return false
}

// SetOidcTestConfiguration gets a reference to the given TestInfoOidcTestConfiguration and assigns it to the OidcTestConfiguration field.
func (o *TestInfo) SetOidcTestConfiguration(v TestInfoOidcTestConfiguration) {
	o.OidcTestConfiguration = &v
}

// GetSamlTestConfiguration returns the SamlTestConfiguration field value if set, zero value otherwise.
func (o *TestInfo) GetSamlTestConfiguration() TestInfoSamlTestConfiguration {
	if o == nil || IsNil(o.SamlTestConfiguration) {
		var ret TestInfoSamlTestConfiguration
		return ret
	}
	return *o.SamlTestConfiguration
}

// GetSamlTestConfigurationOk returns a tuple with the SamlTestConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfo) GetSamlTestConfigurationOk() (*TestInfoSamlTestConfiguration, bool) {
	if o == nil || IsNil(o.SamlTestConfiguration) {
		return nil, false
	}
	return o.SamlTestConfiguration, true
}

// HasSamlTestConfiguration returns a boolean if a field has been set.
func (o *TestInfo) HasSamlTestConfiguration() bool {
	if o != nil && !IsNil(o.SamlTestConfiguration) {
		return true
	}

	return false
}

// SetSamlTestConfiguration gets a reference to the given TestInfoSamlTestConfiguration and assigns it to the SamlTestConfiguration field.
func (o *TestInfo) SetSamlTestConfiguration(v TestInfoSamlTestConfiguration) {
	o.SamlTestConfiguration = &v
}

// GetScimTestConfiguration returns the ScimTestConfiguration field value if set, zero value otherwise.
func (o *TestInfo) GetScimTestConfiguration() TestInfoScimTestConfiguration {
	if o == nil || IsNil(o.ScimTestConfiguration) {
		var ret TestInfoScimTestConfiguration
		return ret
	}
	return *o.ScimTestConfiguration
}

// GetScimTestConfigurationOk returns a tuple with the ScimTestConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfo) GetScimTestConfigurationOk() (*TestInfoScimTestConfiguration, bool) {
	if o == nil || IsNil(o.ScimTestConfiguration) {
		return nil, false
	}
	return o.ScimTestConfiguration, true
}

// HasScimTestConfiguration returns a boolean if a field has been set.
func (o *TestInfo) HasScimTestConfiguration() bool {
	if o != nil && !IsNil(o.ScimTestConfiguration) {
		return true
	}

	return false
}

// SetScimTestConfiguration gets a reference to the given TestInfoScimTestConfiguration and assigns it to the ScimTestConfiguration field.
func (o *TestInfo) SetScimTestConfiguration(v TestInfoScimTestConfiguration) {
	o.ScimTestConfiguration = &v
}

// GetTestAccount returns the TestAccount field value if set, zero value otherwise.
func (o *TestInfo) GetTestAccount() TestInfoTestAccount {
	if o == nil || IsNil(o.TestAccount) {
		var ret TestInfoTestAccount
		return ret
	}
	return *o.TestAccount
}

// GetTestAccountOk returns a tuple with the TestAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfo) GetTestAccountOk() (*TestInfoTestAccount, bool) {
	if o == nil || IsNil(o.TestAccount) {
		return nil, false
	}
	return o.TestAccount, true
}

// HasTestAccount returns a boolean if a field has been set.
func (o *TestInfo) HasTestAccount() bool {
	if o != nil && !IsNil(o.TestAccount) {
		return true
	}

	return false
}

// SetTestAccount gets a reference to the given TestInfoTestAccount and assigns it to the TestAccount field.
func (o *TestInfo) SetTestAccount(v TestInfoTestAccount) {
	o.TestAccount = &v
}

func (o TestInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["escalationSupportContact"] = o.EscalationSupportContact
	if !IsNil(o.OidcTestConfiguration) {
		toSerialize["oidcTestConfiguration"] = o.OidcTestConfiguration
	}
	if !IsNil(o.SamlTestConfiguration) {
		toSerialize["samlTestConfiguration"] = o.SamlTestConfiguration
	}
	if !IsNil(o.ScimTestConfiguration) {
		toSerialize["scimTestConfiguration"] = o.ScimTestConfiguration
	}
	if !IsNil(o.TestAccount) {
		toSerialize["testAccount"] = o.TestAccount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"escalationSupportContact",
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

	varTestInfo := _TestInfo{}

	err = json.Unmarshal(data, &varTestInfo)

	if err != nil {
		return err
	}

	*o = TestInfo(varTestInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "escalationSupportContact")
		delete(additionalProperties, "oidcTestConfiguration")
		delete(additionalProperties, "samlTestConfiguration")
		delete(additionalProperties, "scimTestConfiguration")
		delete(additionalProperties, "testAccount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestInfo struct {
	value *TestInfo
	isSet bool
}

func (v NullableTestInfo) Get() *TestInfo {
	return v.value
}

func (v *NullableTestInfo) Set(val *TestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInfo(val *TestInfo) *NullableTestInfo {
	return &NullableTestInfo{value: val, isSet: true}
}

func (v NullableTestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
