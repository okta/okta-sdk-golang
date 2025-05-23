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

// SimulatePolicyBody The request body required for a simulate policy operation.
type SimulatePolicyBody struct {
	// The application instance ID for a simulate operation
	AppInstance string `json:"appInstance"`
	PolicyContext *PolicyContext `json:"policyContext,omitempty"`
	// Supported policy types for a simulate operation. The default value, `null`, returns all types.
	PolicyTypes []string `json:"policyTypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimulatePolicyBody SimulatePolicyBody

// NewSimulatePolicyBody instantiates a new SimulatePolicyBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimulatePolicyBody(appInstance string) *SimulatePolicyBody {
	this := SimulatePolicyBody{}
	this.AppInstance = appInstance
	return &this
}

// NewSimulatePolicyBodyWithDefaults instantiates a new SimulatePolicyBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulatePolicyBodyWithDefaults() *SimulatePolicyBody {
	this := SimulatePolicyBody{}
	return &this
}

// GetAppInstance returns the AppInstance field value
func (o *SimulatePolicyBody) GetAppInstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppInstance
}

// GetAppInstanceOk returns a tuple with the AppInstance field value
// and a boolean to check if the value has been set.
func (o *SimulatePolicyBody) GetAppInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppInstance, true
}

// SetAppInstance sets field value
func (o *SimulatePolicyBody) SetAppInstance(v string) {
	o.AppInstance = v
}

// GetPolicyContext returns the PolicyContext field value if set, zero value otherwise.
func (o *SimulatePolicyBody) GetPolicyContext() PolicyContext {
	if o == nil || o.PolicyContext == nil {
		var ret PolicyContext
		return ret
	}
	return *o.PolicyContext
}

// GetPolicyContextOk returns a tuple with the PolicyContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePolicyBody) GetPolicyContextOk() (*PolicyContext, bool) {
	if o == nil || o.PolicyContext == nil {
		return nil, false
	}
	return o.PolicyContext, true
}

// HasPolicyContext returns a boolean if a field has been set.
func (o *SimulatePolicyBody) HasPolicyContext() bool {
	if o != nil && o.PolicyContext != nil {
		return true
	}

	return false
}

// SetPolicyContext gets a reference to the given PolicyContext and assigns it to the PolicyContext field.
func (o *SimulatePolicyBody) SetPolicyContext(v PolicyContext) {
	o.PolicyContext = &v
}

// GetPolicyTypes returns the PolicyTypes field value if set, zero value otherwise.
func (o *SimulatePolicyBody) GetPolicyTypes() []string {
	if o == nil || o.PolicyTypes == nil {
		var ret []string
		return ret
	}
	return o.PolicyTypes
}

// GetPolicyTypesOk returns a tuple with the PolicyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimulatePolicyBody) GetPolicyTypesOk() ([]string, bool) {
	if o == nil || o.PolicyTypes == nil {
		return nil, false
	}
	return o.PolicyTypes, true
}

// HasPolicyTypes returns a boolean if a field has been set.
func (o *SimulatePolicyBody) HasPolicyTypes() bool {
	if o != nil && o.PolicyTypes != nil {
		return true
	}

	return false
}

// SetPolicyTypes gets a reference to the given []string and assigns it to the PolicyTypes field.
func (o *SimulatePolicyBody) SetPolicyTypes(v []string) {
	o.PolicyTypes = v
}

func (o SimulatePolicyBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["appInstance"] = o.AppInstance
	}
	if o.PolicyContext != nil {
		toSerialize["policyContext"] = o.PolicyContext
	}
	if o.PolicyTypes != nil {
		toSerialize["policyTypes"] = o.PolicyTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SimulatePolicyBody) UnmarshalJSON(bytes []byte) (err error) {
	varSimulatePolicyBody := _SimulatePolicyBody{}

	err = json.Unmarshal(bytes, &varSimulatePolicyBody)
	if err == nil {
		*o = SimulatePolicyBody(varSimulatePolicyBody)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appInstance")
		delete(additionalProperties, "policyContext")
		delete(additionalProperties, "policyTypes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSimulatePolicyBody struct {
	value *SimulatePolicyBody
	isSet bool
}

func (v NullableSimulatePolicyBody) Get() *SimulatePolicyBody {
	return v.value
}

func (v *NullableSimulatePolicyBody) Set(val *SimulatePolicyBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSimulatePolicyBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSimulatePolicyBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimulatePolicyBody(val *SimulatePolicyBody) *NullableSimulatePolicyBody {
	return &NullableSimulatePolicyBody{value: val, isSet: true}
}

func (v NullableSimulatePolicyBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimulatePolicyBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

