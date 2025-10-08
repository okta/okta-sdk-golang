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

// checks if the SlackApplicationSettingsApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlackApplicationSettingsApplication{}

// SlackApplicationSettingsApplication Slack app instance properties
type SlackApplicationSettingsApplication struct {
	// The Slack app domain name
	Domain string `json:"domain"`
	// The `User.Email` attribute value
	UserEmailValue       *string `json:"userEmailValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SlackApplicationSettingsApplication SlackApplicationSettingsApplication

// NewSlackApplicationSettingsApplication instantiates a new SlackApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackApplicationSettingsApplication(domain string) *SlackApplicationSettingsApplication {
	this := SlackApplicationSettingsApplication{}
	this.Domain = domain
	return &this
}

// NewSlackApplicationSettingsApplicationWithDefaults instantiates a new SlackApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackApplicationSettingsApplicationWithDefaults() *SlackApplicationSettingsApplication {
	this := SlackApplicationSettingsApplication{}
	return &this
}

// GetDomain returns the Domain field value
func (o *SlackApplicationSettingsApplication) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *SlackApplicationSettingsApplication) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *SlackApplicationSettingsApplication) SetDomain(v string) {
	o.Domain = v
}

// GetUserEmailValue returns the UserEmailValue field value if set, zero value otherwise.
func (o *SlackApplicationSettingsApplication) GetUserEmailValue() string {
	if o == nil || IsNil(o.UserEmailValue) {
		var ret string
		return ret
	}
	return *o.UserEmailValue
}

// GetUserEmailValueOk returns a tuple with the UserEmailValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackApplicationSettingsApplication) GetUserEmailValueOk() (*string, bool) {
	if o == nil || IsNil(o.UserEmailValue) {
		return nil, false
	}
	return o.UserEmailValue, true
}

// HasUserEmailValue returns a boolean if a field has been set.
func (o *SlackApplicationSettingsApplication) HasUserEmailValue() bool {
	if o != nil && !IsNil(o.UserEmailValue) {
		return true
	}

	return false
}

// SetUserEmailValue gets a reference to the given string and assigns it to the UserEmailValue field.
func (o *SlackApplicationSettingsApplication) SetUserEmailValue(v string) {
	o.UserEmailValue = &v
}

func (o SlackApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlackApplicationSettingsApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	if !IsNil(o.UserEmailValue) {
		toSerialize["userEmailValue"] = o.UserEmailValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SlackApplicationSettingsApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
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

	varSlackApplicationSettingsApplication := _SlackApplicationSettingsApplication{}

	err = json.Unmarshal(data, &varSlackApplicationSettingsApplication)

	if err != nil {
		return err
	}

	*o = SlackApplicationSettingsApplication(varSlackApplicationSettingsApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "userEmailValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSlackApplicationSettingsApplication struct {
	value *SlackApplicationSettingsApplication
	isSet bool
}

func (v NullableSlackApplicationSettingsApplication) Get() *SlackApplicationSettingsApplication {
	return v.value
}

func (v *NullableSlackApplicationSettingsApplication) Set(val *SlackApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackApplicationSettingsApplication(val *SlackApplicationSettingsApplication) *NullableSlackApplicationSettingsApplication {
	return &NullableSlackApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableSlackApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
