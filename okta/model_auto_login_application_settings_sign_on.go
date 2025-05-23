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

// AutoLoginApplicationSettingsSignOn struct for AutoLoginApplicationSettingsSignOn
type AutoLoginApplicationSettingsSignOn struct {
	// Primary URL of the sign-in page for this app
	LoginUrl *string `json:"loginUrl,omitempty"`
	// Secondary URL of the sign-in page for this app
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoLoginApplicationSettingsSignOn AutoLoginApplicationSettingsSignOn

// NewAutoLoginApplicationSettingsSignOn instantiates a new AutoLoginApplicationSettingsSignOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoLoginApplicationSettingsSignOn() *AutoLoginApplicationSettingsSignOn {
	this := AutoLoginApplicationSettingsSignOn{}
	return &this
}

// NewAutoLoginApplicationSettingsSignOnWithDefaults instantiates a new AutoLoginApplicationSettingsSignOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoLoginApplicationSettingsSignOnWithDefaults() *AutoLoginApplicationSettingsSignOn {
	this := AutoLoginApplicationSettingsSignOn{}
	return &this
}

// GetLoginUrl returns the LoginUrl field value if set, zero value otherwise.
func (o *AutoLoginApplicationSettingsSignOn) GetLoginUrl() string {
	if o == nil || o.LoginUrl == nil {
		var ret string
		return ret
	}
	return *o.LoginUrl
}

// GetLoginUrlOk returns a tuple with the LoginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplicationSettingsSignOn) GetLoginUrlOk() (*string, bool) {
	if o == nil || o.LoginUrl == nil {
		return nil, false
	}
	return o.LoginUrl, true
}

// HasLoginUrl returns a boolean if a field has been set.
func (o *AutoLoginApplicationSettingsSignOn) HasLoginUrl() bool {
	if o != nil && o.LoginUrl != nil {
		return true
	}

	return false
}

// SetLoginUrl gets a reference to the given string and assigns it to the LoginUrl field.
func (o *AutoLoginApplicationSettingsSignOn) SetLoginUrl(v string) {
	o.LoginUrl = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *AutoLoginApplicationSettingsSignOn) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplicationSettingsSignOn) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *AutoLoginApplicationSettingsSignOn) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *AutoLoginApplicationSettingsSignOn) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

func (o AutoLoginApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LoginUrl != nil {
		toSerialize["loginUrl"] = o.LoginUrl
	}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AutoLoginApplicationSettingsSignOn) UnmarshalJSON(bytes []byte) (err error) {
	varAutoLoginApplicationSettingsSignOn := _AutoLoginApplicationSettingsSignOn{}

	err = json.Unmarshal(bytes, &varAutoLoginApplicationSettingsSignOn)
	if err == nil {
		*o = AutoLoginApplicationSettingsSignOn(varAutoLoginApplicationSettingsSignOn)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "loginUrl")
		delete(additionalProperties, "redirectUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAutoLoginApplicationSettingsSignOn struct {
	value *AutoLoginApplicationSettingsSignOn
	isSet bool
}

func (v NullableAutoLoginApplicationSettingsSignOn) Get() *AutoLoginApplicationSettingsSignOn {
	return v.value
}

func (v *NullableAutoLoginApplicationSettingsSignOn) Set(val *AutoLoginApplicationSettingsSignOn) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoLoginApplicationSettingsSignOn) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoLoginApplicationSettingsSignOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoLoginApplicationSettingsSignOn(val *AutoLoginApplicationSettingsSignOn) *NullableAutoLoginApplicationSettingsSignOn {
	return &NullableAutoLoginApplicationSettingsSignOn{value: val, isSet: true}
}

func (v NullableAutoLoginApplicationSettingsSignOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoLoginApplicationSettingsSignOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

