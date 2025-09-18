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
)

// checks if the PlatformConditionEvaluatorPlatformOperatingSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlatformConditionEvaluatorPlatformOperatingSystem{}

// PlatformConditionEvaluatorPlatformOperatingSystem struct for PlatformConditionEvaluatorPlatformOperatingSystem
type PlatformConditionEvaluatorPlatformOperatingSystem struct {
	Expression           *string                                                   `json:"expression,omitempty"`
	Type                 *string                                                   `json:"type,omitempty"`
	Version              *PlatformConditionEvaluatorPlatformOperatingSystemVersion `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlatformConditionEvaluatorPlatformOperatingSystem PlatformConditionEvaluatorPlatformOperatingSystem

// NewPlatformConditionEvaluatorPlatformOperatingSystem instantiates a new PlatformConditionEvaluatorPlatformOperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformConditionEvaluatorPlatformOperatingSystem() *PlatformConditionEvaluatorPlatformOperatingSystem {
	this := PlatformConditionEvaluatorPlatformOperatingSystem{}
	return &this
}

// NewPlatformConditionEvaluatorPlatformOperatingSystemWithDefaults instantiates a new PlatformConditionEvaluatorPlatformOperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformConditionEvaluatorPlatformOperatingSystemWithDefaults() *PlatformConditionEvaluatorPlatformOperatingSystem {
	this := PlatformConditionEvaluatorPlatformOperatingSystem{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) SetExpression(v string) {
	o.Expression = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) GetVersion() PlatformConditionEvaluatorPlatformOperatingSystemVersion {
	if o == nil || IsNil(o.Version) {
		var ret PlatformConditionEvaluatorPlatformOperatingSystemVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) GetVersionOk() (*PlatformConditionEvaluatorPlatformOperatingSystemVersion, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given PlatformConditionEvaluatorPlatformOperatingSystemVersion and assigns it to the Version field.
func (o *PlatformConditionEvaluatorPlatformOperatingSystem) SetVersion(v PlatformConditionEvaluatorPlatformOperatingSystemVersion) {
	o.Version = &v
}

func (o PlatformConditionEvaluatorPlatformOperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlatformConditionEvaluatorPlatformOperatingSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlatformConditionEvaluatorPlatformOperatingSystem) UnmarshalJSON(data []byte) (err error) {
	varPlatformConditionEvaluatorPlatformOperatingSystem := _PlatformConditionEvaluatorPlatformOperatingSystem{}

	err = json.Unmarshal(data, &varPlatformConditionEvaluatorPlatformOperatingSystem)

	if err != nil {
		return err
	}

	*o = PlatformConditionEvaluatorPlatformOperatingSystem(varPlatformConditionEvaluatorPlatformOperatingSystem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlatformConditionEvaluatorPlatformOperatingSystem struct {
	value *PlatformConditionEvaluatorPlatformOperatingSystem
	isSet bool
}

func (v NullablePlatformConditionEvaluatorPlatformOperatingSystem) Get() *PlatformConditionEvaluatorPlatformOperatingSystem {
	return v.value
}

func (v *NullablePlatformConditionEvaluatorPlatformOperatingSystem) Set(val *PlatformConditionEvaluatorPlatformOperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformConditionEvaluatorPlatformOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformConditionEvaluatorPlatformOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformConditionEvaluatorPlatformOperatingSystem(val *PlatformConditionEvaluatorPlatformOperatingSystem) *NullablePlatformConditionEvaluatorPlatformOperatingSystem {
	return &NullablePlatformConditionEvaluatorPlatformOperatingSystem{value: val, isSet: true}
}

func (v NullablePlatformConditionEvaluatorPlatformOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformConditionEvaluatorPlatformOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
