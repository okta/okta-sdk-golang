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
)

// checks if the PasswordPolicyAuthenticationProviderCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyAuthenticationProviderCondition{}

// PasswordPolicyAuthenticationProviderCondition Specifies an authentication provider that's the source of some or all users
type PasswordPolicyAuthenticationProviderCondition struct {
	Include              []string `json:"include,omitempty"`
	Provider             *string  `json:"provider,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyAuthenticationProviderCondition PasswordPolicyAuthenticationProviderCondition

// NewPasswordPolicyAuthenticationProviderCondition instantiates a new PasswordPolicyAuthenticationProviderCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyAuthenticationProviderCondition() *PasswordPolicyAuthenticationProviderCondition {
	this := PasswordPolicyAuthenticationProviderCondition{}
	return &this
}

// NewPasswordPolicyAuthenticationProviderConditionWithDefaults instantiates a new PasswordPolicyAuthenticationProviderCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyAuthenticationProviderConditionWithDefaults() *PasswordPolicyAuthenticationProviderCondition {
	this := PasswordPolicyAuthenticationProviderCondition{}
	return &this
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *PasswordPolicyAuthenticationProviderCondition) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyAuthenticationProviderCondition) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *PasswordPolicyAuthenticationProviderCondition) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *PasswordPolicyAuthenticationProviderCondition) SetInclude(v []string) {
	o.Include = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *PasswordPolicyAuthenticationProviderCondition) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyAuthenticationProviderCondition) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *PasswordPolicyAuthenticationProviderCondition) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *PasswordPolicyAuthenticationProviderCondition) SetProvider(v string) {
	o.Provider = &v
}

func (o PasswordPolicyAuthenticationProviderCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyAuthenticationProviderCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyAuthenticationProviderCondition) UnmarshalJSON(data []byte) (err error) {
	varPasswordPolicyAuthenticationProviderCondition := _PasswordPolicyAuthenticationProviderCondition{}

	err = json.Unmarshal(data, &varPasswordPolicyAuthenticationProviderCondition)

	if err != nil {
		return err
	}

	*o = PasswordPolicyAuthenticationProviderCondition(varPasswordPolicyAuthenticationProviderCondition)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include")
		delete(additionalProperties, "provider")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyAuthenticationProviderCondition struct {
	value *PasswordPolicyAuthenticationProviderCondition
	isSet bool
}

func (v NullablePasswordPolicyAuthenticationProviderCondition) Get() *PasswordPolicyAuthenticationProviderCondition {
	return v.value
}

func (v *NullablePasswordPolicyAuthenticationProviderCondition) Set(val *PasswordPolicyAuthenticationProviderCondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyAuthenticationProviderCondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyAuthenticationProviderCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyAuthenticationProviderCondition(val *PasswordPolicyAuthenticationProviderCondition) *NullablePasswordPolicyAuthenticationProviderCondition {
	return &NullablePasswordPolicyAuthenticationProviderCondition{value: val, isSet: true}
}

func (v NullablePasswordPolicyAuthenticationProviderCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyAuthenticationProviderCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
