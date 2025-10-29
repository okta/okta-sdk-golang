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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the OpenIdConnectApplicationIdpInitiatedLogin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIdConnectApplicationIdpInitiatedLogin{}

// OpenIdConnectApplicationIdpInitiatedLogin The type of IdP-initiated sign-in flow that the client supports
type OpenIdConnectApplicationIdpInitiatedLogin struct {
	// The scopes to use for the request when `mode` is `OKTA`
	DefaultScope []string `json:"default_scope,omitempty"`
	// The mode to use for the IdP-initiated sign-in flow. For `OKTA` or `SPEC` modes, the client must have an `initiate_login_uri` registered. > **Note:** For web and SPA apps, if the mode is `SPEC` or `OKTA`, you must set `grant_types` to `authorization_code`, `implicit`, or `interaction_code`.
	Mode                 string `json:"mode"`
	AdditionalProperties map[string]interface{}
}

type _OpenIdConnectApplicationIdpInitiatedLogin OpenIdConnectApplicationIdpInitiatedLogin

// NewOpenIdConnectApplicationIdpInitiatedLogin instantiates a new OpenIdConnectApplicationIdpInitiatedLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectApplicationIdpInitiatedLogin(mode string) *OpenIdConnectApplicationIdpInitiatedLogin {
	this := OpenIdConnectApplicationIdpInitiatedLogin{}
	this.Mode = mode
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
	if o == nil || IsNil(o.DefaultScope) {
		var ret []string
		return ret
	}
	return o.DefaultScope
}

// GetDefaultScopeOk returns a tuple with the DefaultScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetDefaultScopeOk() ([]string, bool) {
	if o == nil || IsNil(o.DefaultScope) {
		return nil, false
	}
	return o.DefaultScope, true
}

// HasDefaultScope returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) HasDefaultScope() bool {
	if o != nil && !IsNil(o.DefaultScope) {
		return true
	}

	return false
}

// SetDefaultScope gets a reference to the given []string and assigns it to the DefaultScope field.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) SetDefaultScope(v []string) {
	o.DefaultScope = v
}

// GetMode returns the Mode field value
func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationIdpInitiatedLogin) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *OpenIdConnectApplicationIdpInitiatedLogin) SetMode(v string) {
	o.Mode = v
}

func (o OpenIdConnectApplicationIdpInitiatedLogin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIdConnectApplicationIdpInitiatedLogin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultScope) {
		toSerialize["default_scope"] = o.DefaultScope
	}
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenIdConnectApplicationIdpInitiatedLogin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mode",
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

	varOpenIdConnectApplicationIdpInitiatedLogin := _OpenIdConnectApplicationIdpInitiatedLogin{}

	err = json.Unmarshal(data, &varOpenIdConnectApplicationIdpInitiatedLogin)

	if err != nil {
		return err
	}

	*o = OpenIdConnectApplicationIdpInitiatedLogin(varOpenIdConnectApplicationIdpInitiatedLogin)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "default_scope")
		delete(additionalProperties, "mode")
		o.AdditionalProperties = additionalProperties
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
