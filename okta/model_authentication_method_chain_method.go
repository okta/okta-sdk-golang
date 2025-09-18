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
	"reflect"
	"strings"
)

// checks if the AuthenticationMethodChainMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationMethodChainMethod{}

// AuthenticationMethodChainMethod struct for AuthenticationMethodChainMethod
type AuthenticationMethodChainMethod struct {
	VerificationMethod
	// Authentication method chains. Only supports 5 items in the array. Each chain can support maximum 3 steps.
	Chains []AuthenticationMethodChain `json:"chains,omitempty"`
	// Specifies how often the user is prompted for authentication using duration format for the time period. For example, `PT2H30M` for two and a half hours. Don't set this parameter if you're setting the `reauthenticateIn` parameter in `chains`.
	ReauthenticateIn     *string `json:"reauthenticateIn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticationMethodChainMethod AuthenticationMethodChainMethod

// NewAuthenticationMethodChainMethod instantiates a new AuthenticationMethodChainMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationMethodChainMethod() *AuthenticationMethodChainMethod {
	this := AuthenticationMethodChainMethod{}
	return &this
}

// NewAuthenticationMethodChainMethodWithDefaults instantiates a new AuthenticationMethodChainMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationMethodChainMethodWithDefaults() *AuthenticationMethodChainMethod {
	this := AuthenticationMethodChainMethod{}
	return &this
}

// GetChains returns the Chains field value if set, zero value otherwise.
func (o *AuthenticationMethodChainMethod) GetChains() []AuthenticationMethodChain {
	if o == nil || IsNil(o.Chains) {
		var ret []AuthenticationMethodChain
		return ret
	}
	return o.Chains
}

// GetChainsOk returns a tuple with the Chains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodChainMethod) GetChainsOk() ([]AuthenticationMethodChain, bool) {
	if o == nil || IsNil(o.Chains) {
		return nil, false
	}
	return o.Chains, true
}

// HasChains returns a boolean if a field has been set.
func (o *AuthenticationMethodChainMethod) HasChains() bool {
	if o != nil && !IsNil(o.Chains) {
		return true
	}

	return false
}

// SetChains gets a reference to the given []AuthenticationMethodChain and assigns it to the Chains field.
func (o *AuthenticationMethodChainMethod) SetChains(v []AuthenticationMethodChain) {
	o.Chains = v
}

// GetReauthenticateIn returns the ReauthenticateIn field value if set, zero value otherwise.
func (o *AuthenticationMethodChainMethod) GetReauthenticateIn() string {
	if o == nil || IsNil(o.ReauthenticateIn) {
		var ret string
		return ret
	}
	return *o.ReauthenticateIn
}

// GetReauthenticateInOk returns a tuple with the ReauthenticateIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodChainMethod) GetReauthenticateInOk() (*string, bool) {
	if o == nil || IsNil(o.ReauthenticateIn) {
		return nil, false
	}
	return o.ReauthenticateIn, true
}

// HasReauthenticateIn returns a boolean if a field has been set.
func (o *AuthenticationMethodChainMethod) HasReauthenticateIn() bool {
	if o != nil && !IsNil(o.ReauthenticateIn) {
		return true
	}

	return false
}

// SetReauthenticateIn gets a reference to the given string and assigns it to the ReauthenticateIn field.
func (o *AuthenticationMethodChainMethod) SetReauthenticateIn(v string) {
	o.ReauthenticateIn = &v
}

func (o AuthenticationMethodChainMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationMethodChainMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVerificationMethod, errVerificationMethod := json.Marshal(o.VerificationMethod)
	if errVerificationMethod != nil {
		return map[string]interface{}{}, errVerificationMethod
	}
	errVerificationMethod = json.Unmarshal([]byte(serializedVerificationMethod), &toSerialize)
	if errVerificationMethod != nil {
		return map[string]interface{}{}, errVerificationMethod
	}
	if !IsNil(o.Chains) {
		toSerialize["chains"] = o.Chains
	}
	if !IsNil(o.ReauthenticateIn) {
		toSerialize["reauthenticateIn"] = o.ReauthenticateIn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticationMethodChainMethod) UnmarshalJSON(data []byte) (err error) {
	type AuthenticationMethodChainMethodWithoutEmbeddedStruct struct {
		// Authentication method chains. Only supports 5 items in the array. Each chain can support maximum 3 steps.
		Chains []AuthenticationMethodChain `json:"chains,omitempty"`
		// Specifies how often the user is prompted for authentication using duration format for the time period. For example, `PT2H30M` for two and a half hours. Don't set this parameter if you're setting the `reauthenticateIn` parameter in `chains`.
		ReauthenticateIn *string `json:"reauthenticateIn,omitempty"`
	}

	varAuthenticationMethodChainMethodWithoutEmbeddedStruct := AuthenticationMethodChainMethodWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAuthenticationMethodChainMethodWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticationMethodChainMethod := _AuthenticationMethodChainMethod{}
		varAuthenticationMethodChainMethod.Chains = varAuthenticationMethodChainMethodWithoutEmbeddedStruct.Chains
		varAuthenticationMethodChainMethod.ReauthenticateIn = varAuthenticationMethodChainMethodWithoutEmbeddedStruct.ReauthenticateIn
		*o = AuthenticationMethodChainMethod(varAuthenticationMethodChainMethod)
	} else {
		return err
	}

	varAuthenticationMethodChainMethod := _AuthenticationMethodChainMethod{}

	err = json.Unmarshal(data, &varAuthenticationMethodChainMethod)
	if err == nil {
		o.VerificationMethod = varAuthenticationMethodChainMethod.VerificationMethod
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "chains")
		delete(additionalProperties, "reauthenticateIn")

		// remove fields from embedded structs
		reflectVerificationMethod := reflect.ValueOf(o.VerificationMethod)
		for i := 0; i < reflectVerificationMethod.Type().NumField(); i++ {
			t := reflectVerificationMethod.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticationMethodChainMethod struct {
	value *AuthenticationMethodChainMethod
	isSet bool
}

func (v NullableAuthenticationMethodChainMethod) Get() *AuthenticationMethodChainMethod {
	return v.value
}

func (v *NullableAuthenticationMethodChainMethod) Set(val *AuthenticationMethodChainMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationMethodChainMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationMethodChainMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationMethodChainMethod(val *AuthenticationMethodChainMethod) *NullableAuthenticationMethodChainMethod {
	return &NullableAuthenticationMethodChainMethod{value: val, isSet: true}
}

func (v NullableAuthenticationMethodChainMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationMethodChainMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
