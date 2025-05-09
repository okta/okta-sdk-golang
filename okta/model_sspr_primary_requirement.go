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

// SsprPrimaryRequirement Defines the authenticators permitted for the initial authentication step of password recovery
type SsprPrimaryRequirement struct {
	// Constraints on the values specified in the `methods` array. Specifying a constraint limits methods to specific authenticator(s). Currently, Google OTP is the only accepted constraint.
	MethodConstraints []AuthenticatorMethodConstraint `json:"methodConstraints,omitempty"`
	// Authenticator methods allowed for the initial authentication step of password recovery. Method `otp` requires a constraint limiting it to a Google authenticator.
	Methods []string `json:"methods,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SsprPrimaryRequirement SsprPrimaryRequirement

// NewSsprPrimaryRequirement instantiates a new SsprPrimaryRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsprPrimaryRequirement() *SsprPrimaryRequirement {
	this := SsprPrimaryRequirement{}
	return &this
}

// NewSsprPrimaryRequirementWithDefaults instantiates a new SsprPrimaryRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsprPrimaryRequirementWithDefaults() *SsprPrimaryRequirement {
	this := SsprPrimaryRequirement{}
	return &this
}

// GetMethodConstraints returns the MethodConstraints field value if set, zero value otherwise.
func (o *SsprPrimaryRequirement) GetMethodConstraints() []AuthenticatorMethodConstraint {
	if o == nil || o.MethodConstraints == nil {
		var ret []AuthenticatorMethodConstraint
		return ret
	}
	return o.MethodConstraints
}

// GetMethodConstraintsOk returns a tuple with the MethodConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprPrimaryRequirement) GetMethodConstraintsOk() ([]AuthenticatorMethodConstraint, bool) {
	if o == nil || o.MethodConstraints == nil {
		return nil, false
	}
	return o.MethodConstraints, true
}

// HasMethodConstraints returns a boolean if a field has been set.
func (o *SsprPrimaryRequirement) HasMethodConstraints() bool {
	if o != nil && o.MethodConstraints != nil {
		return true
	}

	return false
}

// SetMethodConstraints gets a reference to the given []AuthenticatorMethodConstraint and assigns it to the MethodConstraints field.
func (o *SsprPrimaryRequirement) SetMethodConstraints(v []AuthenticatorMethodConstraint) {
	o.MethodConstraints = v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *SsprPrimaryRequirement) GetMethods() []string {
	if o == nil || o.Methods == nil {
		var ret []string
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprPrimaryRequirement) GetMethodsOk() ([]string, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *SsprPrimaryRequirement) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []string and assigns it to the Methods field.
func (o *SsprPrimaryRequirement) SetMethods(v []string) {
	o.Methods = v
}

func (o SsprPrimaryRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MethodConstraints != nil {
		toSerialize["methodConstraints"] = o.MethodConstraints
	}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SsprPrimaryRequirement) UnmarshalJSON(bytes []byte) (err error) {
	varSsprPrimaryRequirement := _SsprPrimaryRequirement{}

	err = json.Unmarshal(bytes, &varSsprPrimaryRequirement)
	if err == nil {
		*o = SsprPrimaryRequirement(varSsprPrimaryRequirement)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "methodConstraints")
		delete(additionalProperties, "methods")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSsprPrimaryRequirement struct {
	value *SsprPrimaryRequirement
	isSet bool
}

func (v NullableSsprPrimaryRequirement) Get() *SsprPrimaryRequirement {
	return v.value
}

func (v *NullableSsprPrimaryRequirement) Set(val *SsprPrimaryRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableSsprPrimaryRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableSsprPrimaryRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsprPrimaryRequirement(val *SsprPrimaryRequirement) *NullableSsprPrimaryRequirement {
	return &NullableSsprPrimaryRequirement{value: val, isSet: true}
}

func (v NullableSsprPrimaryRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsprPrimaryRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

