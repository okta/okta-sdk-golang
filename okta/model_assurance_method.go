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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// checks if the AssuranceMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssuranceMethod{}

// AssuranceMethod struct for AssuranceMethod
type AssuranceMethod struct {
	VerificationMethod
	// Specifies constraints for the authenticator. Constraints are logically evaluated such that only one constraint object needs to be satisfied. But, within a constraint object, each constraint property must be satisfied.
	Constraints []AccessPolicyConstraints `json:"constraints,omitempty"`
	FactorMode  *string                   `json:"factorMode,omitempty"`
	// The inactivity duration after which the user must re-authenticate. Use the ISO 8601 period format (for example, PT2H).
	InactivityPeriod *string `json:"inactivityPeriod,omitempty"`
	// The duration after which the user must re-authenticate, regardless of user activity. Keep in mind that the re-authentication intervals for constraints take precedent over this value. Use the ISO 8601 period format for recurring time intervals (for example, PT2H, PT0S, PT43800H, and so on).
	ReauthenticateIn     *string `json:"reauthenticateIn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssuranceMethod AssuranceMethod

// NewAssuranceMethod instantiates a new AssuranceMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceMethod() *AssuranceMethod {
	this := AssuranceMethod{}
	return &this
}

// NewAssuranceMethodWithDefaults instantiates a new AssuranceMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceMethodWithDefaults() *AssuranceMethod {
	this := AssuranceMethod{}
	return &this
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *AssuranceMethod) GetConstraints() []AccessPolicyConstraints {
	if o == nil || IsNil(o.Constraints) {
		var ret []AccessPolicyConstraints
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceMethod) GetConstraintsOk() ([]AccessPolicyConstraints, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *AssuranceMethod) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []AccessPolicyConstraints and assigns it to the Constraints field.
func (o *AssuranceMethod) SetConstraints(v []AccessPolicyConstraints) {
	o.Constraints = v
}

// GetFactorMode returns the FactorMode field value if set, zero value otherwise.
func (o *AssuranceMethod) GetFactorMode() string {
	if o == nil || IsNil(o.FactorMode) {
		var ret string
		return ret
	}
	return *o.FactorMode
}

// GetFactorModeOk returns a tuple with the FactorMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceMethod) GetFactorModeOk() (*string, bool) {
	if o == nil || IsNil(o.FactorMode) {
		return nil, false
	}
	return o.FactorMode, true
}

// HasFactorMode returns a boolean if a field has been set.
func (o *AssuranceMethod) HasFactorMode() bool {
	if o != nil && !IsNil(o.FactorMode) {
		return true
	}

	return false
}

// SetFactorMode gets a reference to the given string and assigns it to the FactorMode field.
func (o *AssuranceMethod) SetFactorMode(v string) {
	o.FactorMode = &v
}

// GetInactivityPeriod returns the InactivityPeriod field value if set, zero value otherwise.
func (o *AssuranceMethod) GetInactivityPeriod() string {
	if o == nil || IsNil(o.InactivityPeriod) {
		var ret string
		return ret
	}
	return *o.InactivityPeriod
}

// GetInactivityPeriodOk returns a tuple with the InactivityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceMethod) GetInactivityPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.InactivityPeriod) {
		return nil, false
	}
	return o.InactivityPeriod, true
}

// HasInactivityPeriod returns a boolean if a field has been set.
func (o *AssuranceMethod) HasInactivityPeriod() bool {
	if o != nil && !IsNil(o.InactivityPeriod) {
		return true
	}

	return false
}

// SetInactivityPeriod gets a reference to the given string and assigns it to the InactivityPeriod field.
func (o *AssuranceMethod) SetInactivityPeriod(v string) {
	o.InactivityPeriod = &v
}

// GetReauthenticateIn returns the ReauthenticateIn field value if set, zero value otherwise.
func (o *AssuranceMethod) GetReauthenticateIn() string {
	if o == nil || IsNil(o.ReauthenticateIn) {
		var ret string
		return ret
	}
	return *o.ReauthenticateIn
}

// GetReauthenticateInOk returns a tuple with the ReauthenticateIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceMethod) GetReauthenticateInOk() (*string, bool) {
	if o == nil || IsNil(o.ReauthenticateIn) {
		return nil, false
	}
	return o.ReauthenticateIn, true
}

// HasReauthenticateIn returns a boolean if a field has been set.
func (o *AssuranceMethod) HasReauthenticateIn() bool {
	if o != nil && !IsNil(o.ReauthenticateIn) {
		return true
	}

	return false
}

// SetReauthenticateIn gets a reference to the given string and assigns it to the ReauthenticateIn field.
func (o *AssuranceMethod) SetReauthenticateIn(v string) {
	o.ReauthenticateIn = &v
}

func (o AssuranceMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssuranceMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVerificationMethod, errVerificationMethod := json.Marshal(o.VerificationMethod)
	if errVerificationMethod != nil {
		return map[string]interface{}{}, errVerificationMethod
	}
	errVerificationMethod = json.Unmarshal([]byte(serializedVerificationMethod), &toSerialize)
	if errVerificationMethod != nil {
		return map[string]interface{}{}, errVerificationMethod
	}
	if !IsNil(o.Constraints) {
		toSerialize["constraints"] = o.Constraints
	}
	if !IsNil(o.FactorMode) {
		toSerialize["factorMode"] = o.FactorMode
	}
	if !IsNil(o.InactivityPeriod) {
		toSerialize["inactivityPeriod"] = o.InactivityPeriod
	}
	if !IsNil(o.ReauthenticateIn) {
		toSerialize["reauthenticateIn"] = o.ReauthenticateIn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssuranceMethod) UnmarshalJSON(data []byte) (err error) {
	type AssuranceMethodWithoutEmbeddedStruct struct {
		// Specifies constraints for the authenticator. Constraints are logically evaluated such that only one constraint object needs to be satisfied. But, within a constraint object, each constraint property must be satisfied.
		Constraints []AccessPolicyConstraints `json:"constraints,omitempty"`
		FactorMode  *string                   `json:"factorMode,omitempty"`
		// The inactivity duration after which the user must re-authenticate. Use the ISO 8601 period format (for example, PT2H).
		InactivityPeriod *string `json:"inactivityPeriod,omitempty"`
		// The duration after which the user must re-authenticate, regardless of user activity. Keep in mind that the re-authentication intervals for constraints take precedent over this value. Use the ISO 8601 period format for recurring time intervals (for example, PT2H, PT0S, PT43800H, and so on).
		ReauthenticateIn *string `json:"reauthenticateIn,omitempty"`
	}

	varAssuranceMethodWithoutEmbeddedStruct := AssuranceMethodWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varAssuranceMethodWithoutEmbeddedStruct)
	if err == nil {
		varAssuranceMethod := _AssuranceMethod{}
		varAssuranceMethod.Constraints = varAssuranceMethodWithoutEmbeddedStruct.Constraints
		varAssuranceMethod.FactorMode = varAssuranceMethodWithoutEmbeddedStruct.FactorMode
		varAssuranceMethod.InactivityPeriod = varAssuranceMethodWithoutEmbeddedStruct.InactivityPeriod
		varAssuranceMethod.ReauthenticateIn = varAssuranceMethodWithoutEmbeddedStruct.ReauthenticateIn
		*o = AssuranceMethod(varAssuranceMethod)
	} else {
		return err
	}

	varAssuranceMethod := _AssuranceMethod{}

	err = json.Unmarshal(data, &varAssuranceMethod)
	if err == nil {
		o.VerificationMethod = varAssuranceMethod.VerificationMethod
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "factorMode")
		delete(additionalProperties, "inactivityPeriod")
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

type NullableAssuranceMethod struct {
	value *AssuranceMethod
	isSet bool
}

func (v NullableAssuranceMethod) Get() *AssuranceMethod {
	return v.value
}

func (v *NullableAssuranceMethod) Set(val *AssuranceMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceMethod(val *AssuranceMethod) *NullableAssuranceMethod {
	return &NullableAssuranceMethod{value: val, isSet: true}
}

func (v NullableAssuranceMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
