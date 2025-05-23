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

// SsprStepUpRequirement Defines the secondary authenticators needed for password reset if `required` is true. The following are three valid configurations: * `required`=false * `required`=true with no methods to use any SSO authenticator * `required`=true with `security_question` as the method
type SsprStepUpRequirement struct {
	// Authenticator methods required for secondary authentication step of password recovery. Specify this value only when `required` is true and `security_question` is permitted for the secondary authentication.
	Methods []string `json:"methods,omitempty"`
	Required *bool `json:"required,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SsprStepUpRequirement SsprStepUpRequirement

// NewSsprStepUpRequirement instantiates a new SsprStepUpRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsprStepUpRequirement() *SsprStepUpRequirement {
	this := SsprStepUpRequirement{}
	return &this
}

// NewSsprStepUpRequirementWithDefaults instantiates a new SsprStepUpRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsprStepUpRequirementWithDefaults() *SsprStepUpRequirement {
	this := SsprStepUpRequirement{}
	return &this
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *SsprStepUpRequirement) GetMethods() []string {
	if o == nil || o.Methods == nil {
		var ret []string
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprStepUpRequirement) GetMethodsOk() ([]string, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *SsprStepUpRequirement) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []string and assigns it to the Methods field.
func (o *SsprStepUpRequirement) SetMethods(v []string) {
	o.Methods = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *SsprStepUpRequirement) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprStepUpRequirement) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *SsprStepUpRequirement) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *SsprStepUpRequirement) SetRequired(v bool) {
	o.Required = &v
}

func (o SsprStepUpRequirement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SsprStepUpRequirement) UnmarshalJSON(bytes []byte) (err error) {
	varSsprStepUpRequirement := _SsprStepUpRequirement{}

	err = json.Unmarshal(bytes, &varSsprStepUpRequirement)
	if err == nil {
		*o = SsprStepUpRequirement(varSsprStepUpRequirement)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "methods")
		delete(additionalProperties, "required")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSsprStepUpRequirement struct {
	value *SsprStepUpRequirement
	isSet bool
}

func (v NullableSsprStepUpRequirement) Get() *SsprStepUpRequirement {
	return v.value
}

func (v *NullableSsprStepUpRequirement) Set(val *SsprStepUpRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableSsprStepUpRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableSsprStepUpRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsprStepUpRequirement(val *SsprStepUpRequirement) *NullableSsprStepUpRequirement {
	return &NullableSsprStepUpRequirement{value: val, isSet: true}
}

func (v NullableSsprStepUpRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsprStepUpRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

