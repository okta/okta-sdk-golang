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

// AuthorizationServerPolicy struct for AuthorizationServerPolicy
type AuthorizationServerPolicy struct {
	Conditions *AuthorizationServerPolicyConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicy AuthorizationServerPolicy

// NewAuthorizationServerPolicy instantiates a new AuthorizationServerPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicy() *AuthorizationServerPolicy {
	this := AuthorizationServerPolicy{}
	return &this
}

// NewAuthorizationServerPolicyWithDefaults instantiates a new AuthorizationServerPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyWithDefaults() *AuthorizationServerPolicy {
	this := AuthorizationServerPolicy{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetConditions() AuthorizationServerPolicyConditions {
	if o == nil || o.Conditions == nil {
		var ret AuthorizationServerPolicyConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetConditionsOk() (*AuthorizationServerPolicyConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given AuthorizationServerPolicyConditions and assigns it to the Conditions field.
func (o *AuthorizationServerPolicy) SetConditions(v AuthorizationServerPolicyConditions) {
	o.Conditions = &v
}

func (o AuthorizationServerPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicy) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicy := _AuthorizationServerPolicy{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicy)
	if err == nil {
		*o = AuthorizationServerPolicy(varAuthorizationServerPolicy)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conditions")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicy struct {
	value *AuthorizationServerPolicy
	isSet bool
}

func (v NullableAuthorizationServerPolicy) Get() *AuthorizationServerPolicy {
	return v.value
}

func (v *NullableAuthorizationServerPolicy) Set(val *AuthorizationServerPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicy(val *AuthorizationServerPolicy) *NullableAuthorizationServerPolicy {
	return &NullableAuthorizationServerPolicy{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

