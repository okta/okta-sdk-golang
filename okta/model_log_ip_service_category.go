/*
Okta Admin Management API

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
)

// checks if the LogIpServiceCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogIpServiceCategory{}

// LogIpServiceCategory Describes the IP service category associated with an IP.
type LogIpServiceCategory struct {
	// Indicates whether the service is an anonymizer
	IsAnonymous NullableBool `json:"isAnonymous,omitempty"`
	// The name of the associated operator
	Operator NullableString `json:"operator,omitempty"`
	// The type of service provided from this IP address
	Type                 NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogIpServiceCategory LogIpServiceCategory

// NewLogIpServiceCategory instantiates a new LogIpServiceCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogIpServiceCategory() *LogIpServiceCategory {
	this := LogIpServiceCategory{}
	return &this
}

// NewLogIpServiceCategoryWithDefaults instantiates a new LogIpServiceCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogIpServiceCategoryWithDefaults() *LogIpServiceCategory {
	this := LogIpServiceCategory{}
	return &this
}

// GetIsAnonymous returns the IsAnonymous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogIpServiceCategory) GetIsAnonymous() bool {
	if o == nil || IsNil(o.IsAnonymous.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAnonymous.Get()
}

// GetIsAnonymousOk returns a tuple with the IsAnonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogIpServiceCategory) GetIsAnonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAnonymous.Get(), o.IsAnonymous.IsSet()
}

// HasIsAnonymous returns a boolean if a field has been set.
func (o *LogIpServiceCategory) HasIsAnonymous() bool {
	if o != nil && o.IsAnonymous.IsSet() {
		return true
	}

	return false
}

// SetIsAnonymous gets a reference to the given NullableBool and assigns it to the IsAnonymous field.
func (o *LogIpServiceCategory) SetIsAnonymous(v bool) {
	o.IsAnonymous.Set(&v)
}

// SetIsAnonymousNil sets the value for IsAnonymous to be an explicit nil
func (o *LogIpServiceCategory) SetIsAnonymousNil() {
	o.IsAnonymous.Set(nil)
}

// UnsetIsAnonymous ensures that no value is present for IsAnonymous, not even an explicit nil
func (o *LogIpServiceCategory) UnsetIsAnonymous() {
	o.IsAnonymous.Unset()
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogIpServiceCategory) GetOperator() string {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret string
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogIpServiceCategory) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *LogIpServiceCategory) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableString and assigns it to the Operator field.
func (o *LogIpServiceCategory) SetOperator(v string) {
	o.Operator.Set(&v)
}

// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *LogIpServiceCategory) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *LogIpServiceCategory) UnsetOperator() {
	o.Operator.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogIpServiceCategory) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogIpServiceCategory) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *LogIpServiceCategory) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *LogIpServiceCategory) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *LogIpServiceCategory) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *LogIpServiceCategory) UnsetType() {
	o.Type.Unset()
}

func (o LogIpServiceCategory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogIpServiceCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsAnonymous.IsSet() {
		toSerialize["isAnonymous"] = o.IsAnonymous.Get()
	}
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogIpServiceCategory) UnmarshalJSON(data []byte) (err error) {
	varLogIpServiceCategory := _LogIpServiceCategory{}

	err = json.Unmarshal(data, &varLogIpServiceCategory)

	if err != nil {
		return err
	}

	*o = LogIpServiceCategory(varLogIpServiceCategory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isAnonymous")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogIpServiceCategory struct {
	value *LogIpServiceCategory
	isSet bool
}

func (v NullableLogIpServiceCategory) Get() *LogIpServiceCategory {
	return v.value
}

func (v *NullableLogIpServiceCategory) Set(val *LogIpServiceCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableLogIpServiceCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableLogIpServiceCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogIpServiceCategory(val *LogIpServiceCategory) *NullableLogIpServiceCategory {
	return &NullableLogIpServiceCategory{value: val, isSet: true}
}

func (v NullableLogIpServiceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogIpServiceCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
