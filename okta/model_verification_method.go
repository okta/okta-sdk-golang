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

// VerificationMethod struct for VerificationMethod
type VerificationMethod struct {
	Constraints []AccessPolicyConstraints `json:"constraints,omitempty"`
	FactorMode *string `json:"factorMode,omitempty"`
	ReauthenticateIn *string `json:"reauthenticateIn,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VerificationMethod VerificationMethod

// NewVerificationMethod instantiates a new VerificationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationMethod() *VerificationMethod {
	this := VerificationMethod{}
	return &this
}

// NewVerificationMethodWithDefaults instantiates a new VerificationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationMethodWithDefaults() *VerificationMethod {
	this := VerificationMethod{}
	return &this
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *VerificationMethod) GetConstraints() []AccessPolicyConstraints {
	if o == nil || o.Constraints == nil {
		var ret []AccessPolicyConstraints
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMethod) GetConstraintsOk() ([]AccessPolicyConstraints, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *VerificationMethod) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []AccessPolicyConstraints and assigns it to the Constraints field.
func (o *VerificationMethod) SetConstraints(v []AccessPolicyConstraints) {
	o.Constraints = v
}

// GetFactorMode returns the FactorMode field value if set, zero value otherwise.
func (o *VerificationMethod) GetFactorMode() string {
	if o == nil || o.FactorMode == nil {
		var ret string
		return ret
	}
	return *o.FactorMode
}

// GetFactorModeOk returns a tuple with the FactorMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMethod) GetFactorModeOk() (*string, bool) {
	if o == nil || o.FactorMode == nil {
		return nil, false
	}
	return o.FactorMode, true
}

// HasFactorMode returns a boolean if a field has been set.
func (o *VerificationMethod) HasFactorMode() bool {
	if o != nil && o.FactorMode != nil {
		return true
	}

	return false
}

// SetFactorMode gets a reference to the given string and assigns it to the FactorMode field.
func (o *VerificationMethod) SetFactorMode(v string) {
	o.FactorMode = &v
}

// GetReauthenticateIn returns the ReauthenticateIn field value if set, zero value otherwise.
func (o *VerificationMethod) GetReauthenticateIn() string {
	if o == nil || o.ReauthenticateIn == nil {
		var ret string
		return ret
	}
	return *o.ReauthenticateIn
}

// GetReauthenticateInOk returns a tuple with the ReauthenticateIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMethod) GetReauthenticateInOk() (*string, bool) {
	if o == nil || o.ReauthenticateIn == nil {
		return nil, false
	}
	return o.ReauthenticateIn, true
}

// HasReauthenticateIn returns a boolean if a field has been set.
func (o *VerificationMethod) HasReauthenticateIn() bool {
	if o != nil && o.ReauthenticateIn != nil {
		return true
	}

	return false
}

// SetReauthenticateIn gets a reference to the given string and assigns it to the ReauthenticateIn field.
func (o *VerificationMethod) SetReauthenticateIn(v string) {
	o.ReauthenticateIn = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VerificationMethod) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationMethod) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VerificationMethod) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VerificationMethod) SetType(v string) {
	o.Type = &v
}

func (o VerificationMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if o.FactorMode != nil {
		toSerialize["factorMode"] = o.FactorMode
	}
	if o.ReauthenticateIn != nil {
		toSerialize["reauthenticateIn"] = o.ReauthenticateIn
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VerificationMethod) UnmarshalJSON(bytes []byte) (err error) {
	varVerificationMethod := _VerificationMethod{}

	err = json.Unmarshal(bytes, &varVerificationMethod)
	if err == nil {
		*o = VerificationMethod(varVerificationMethod)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "factorMode")
		delete(additionalProperties, "reauthenticateIn")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableVerificationMethod struct {
	value *VerificationMethod
	isSet bool
}

func (v NullableVerificationMethod) Get() *VerificationMethod {
	return v.value
}

func (v *NullableVerificationMethod) Set(val *VerificationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationMethod(val *VerificationMethod) *NullableVerificationMethod {
	return &NullableVerificationMethod{value: val, isSet: true}
}

func (v NullableVerificationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

