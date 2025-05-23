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

// SlackApplicationSettingsApplication Slack app instance properties
type SlackApplicationSettingsApplication struct {
	// The Slack app domain name
	Domain string `json:"domain"`
	// The `User.Email` attribute value
	UserEmailValue *string `json:"userEmailValue,omitempty"`
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
	if o == nil || o.UserEmailValue == nil {
		var ret string
		return ret
	}
	return *o.UserEmailValue
}

// GetUserEmailValueOk returns a tuple with the UserEmailValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackApplicationSettingsApplication) GetUserEmailValueOk() (*string, bool) {
	if o == nil || o.UserEmailValue == nil {
		return nil, false
	}
	return o.UserEmailValue, true
}

// HasUserEmailValue returns a boolean if a field has been set.
func (o *SlackApplicationSettingsApplication) HasUserEmailValue() bool {
	if o != nil && o.UserEmailValue != nil {
		return true
	}

	return false
}

// SetUserEmailValue gets a reference to the given string and assigns it to the UserEmailValue field.
func (o *SlackApplicationSettingsApplication) SetUserEmailValue(v string) {
	o.UserEmailValue = &v
}

func (o SlackApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if o.UserEmailValue != nil {
		toSerialize["userEmailValue"] = o.UserEmailValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SlackApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varSlackApplicationSettingsApplication := _SlackApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varSlackApplicationSettingsApplication)
	if err == nil {
		*o = SlackApplicationSettingsApplication(varSlackApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "userEmailValue")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

