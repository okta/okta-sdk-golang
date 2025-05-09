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

// SingleLogout struct for SingleLogout
type SingleLogout struct {
	Enabled *bool `json:"enabled,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SingleLogout SingleLogout

// NewSingleLogout instantiates a new SingleLogout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleLogout() *SingleLogout {
	this := SingleLogout{}
	return &this
}

// NewSingleLogoutWithDefaults instantiates a new SingleLogout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleLogoutWithDefaults() *SingleLogout {
	this := SingleLogout{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SingleLogout) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleLogout) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SingleLogout) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SingleLogout) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SingleLogout) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleLogout) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SingleLogout) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SingleLogout) SetIssuer(v string) {
	o.Issuer = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *SingleLogout) GetLogoutUrl() string {
	if o == nil || o.LogoutUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleLogout) GetLogoutUrlOk() (*string, bool) {
	if o == nil || o.LogoutUrl == nil {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *SingleLogout) HasLogoutUrl() bool {
	if o != nil && o.LogoutUrl != nil {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *SingleLogout) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

func (o SingleLogout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.LogoutUrl != nil {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SingleLogout) UnmarshalJSON(bytes []byte) (err error) {
	varSingleLogout := _SingleLogout{}

	err = json.Unmarshal(bytes, &varSingleLogout)
	if err == nil {
		*o = SingleLogout(varSingleLogout)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "logoutUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSingleLogout struct {
	value *SingleLogout
	isSet bool
}

func (v NullableSingleLogout) Get() *SingleLogout {
	return v.value
}

func (v *NullableSingleLogout) Set(val *SingleLogout) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleLogout) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleLogout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleLogout(val *SingleLogout) *NullableSingleLogout {
	return &NullableSingleLogout{value: val, isSet: true}
}

func (v NullableSingleLogout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleLogout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

