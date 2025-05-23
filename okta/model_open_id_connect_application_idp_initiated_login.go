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

// OpenIdConnectApplicationIdpInitiatedLogin struct for OpenIdConnectApplicationIdpInitiatedLogin
type OpenIdConnectApplicationIdpInitiatedLogin struct {
	DefaultScope []string `json:"default_scope,omitempty"`
	Mode *string `json:"mode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenIdConnectApplicationIdpInitiatedLogin OpenIdConnectApplicationIdpInitiatedLogin

// NewOpenIdConnectApplicationIdpInitiatedLogin instantiates a new OpenIdConnectApplicationIdpInitiatedLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectApplicationIdpInitiatedLogin() *OpenIdConnectApplicationIdpInitiatedLogin {
	this := OpenIdConnectApplicationIdpInitiatedLogin{}
	return &this
}

// NewOpenIdConnectApplicationIdpInitiatedLoginWithDefaults instantiates a new OpenIdConnectApplicationIdpInitiatedLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIdConnectApplicationIdpInitiatedLoginWithDefaults() *OpenIdConnectApplicationIdpInitiatedLogin {
	this := OpenIdConnectApplicationIdpInitiatedLogin{}
	return &this
}

// GetDefaultScope returns the DefaultScope field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetDefaultScope() []string {
	if o == nil || o.DefaultScope == nil {
		var ret []string
		return ret
	}
	return o.DefaultScope
}

// GetDefaultScopeOk returns a tuple with the DefaultScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetDefaultScopeOk() ([]string, bool) {
	if o == nil || o.DefaultScope == nil {
		return nil, false
	}
	return o.DefaultScope, true
}

// HasDefaultScope returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) HasDefaultScope() bool {
	if o != nil && o.DefaultScope != nil {
		return true
	}

	return false
}

// SetDefaultScope gets a reference to the given []string and assigns it to the DefaultScope field.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) SetDefaultScope(v []string) {
	o.DefaultScope = v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) SetMode(v string) {
	o.Mode = &v
}

func (o OpenIdConnectApplicationIdpInitiatedLogin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultScope != nil {
		toSerialize["default_scope"] = o.DefaultScope
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenIdConnectApplicationIdpInitiatedLogin) UnmarshalJSON(bytes []byte) (err error) {
	varOpenIdConnectApplicationIdpInitiatedLogin := _OpenIdConnectApplicationIdpInitiatedLogin{}

	err = json.Unmarshal(bytes, &varOpenIdConnectApplicationIdpInitiatedLogin)
	if err == nil {
		*o = OpenIdConnectApplicationIdpInitiatedLogin(varOpenIdConnectApplicationIdpInitiatedLogin)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "default_scope")
		delete(additionalProperties, "mode")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOpenIdConnectApplicationIdpInitiatedLogin struct {
	value *OpenIdConnectApplicationIdpInitiatedLogin
	isSet bool
}

func (v NullableOpenIdConnectApplicationIdpInitiatedLogin) Get() *OpenIdConnectApplicationIdpInitiatedLogin {
	return v.value
}

func (v *NullableOpenIdConnectApplicationIdpInitiatedLogin) Set(val *OpenIdConnectApplicationIdpInitiatedLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectApplicationIdpInitiatedLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectApplicationIdpInitiatedLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectApplicationIdpInitiatedLogin(val *OpenIdConnectApplicationIdpInitiatedLogin) *NullableOpenIdConnectApplicationIdpInitiatedLogin {
	return &NullableOpenIdConnectApplicationIdpInitiatedLogin{value: val, isSet: true}
}

func (v NullableOpenIdConnectApplicationIdpInitiatedLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectApplicationIdpInitiatedLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

