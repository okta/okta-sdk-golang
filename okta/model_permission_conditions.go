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

// checks if the PermissionConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionConditions{}

// PermissionConditions Conditions for further restricting a permission. See [Permission conditions](https://help.okta.com/okta_help.htm?type=oie&id=ext-permission-conditions).
type PermissionConditions struct {
	// Exclude attributes with specific values for the permission
	Exclude map[string]map[string]interface{} `json:"exclude,omitempty"`
	// Include attributes with specific values for the permission
	Include              map[string]map[string]interface{} `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PermissionConditions PermissionConditions

// NewPermissionConditions instantiates a new PermissionConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionConditions() *PermissionConditions {
	this := PermissionConditions{}
	return &this
}

// NewPermissionConditionsWithDefaults instantiates a new PermissionConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionConditionsWithDefaults() *PermissionConditions {
	this := PermissionConditions{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PermissionConditions) GetExclude() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PermissionConditions) GetExcludeOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Exclude) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *PermissionConditions) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given map[string]map[string]interface{} and assigns it to the Exclude field.
func (o *PermissionConditions) SetExclude(v map[string]map[string]interface{}) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PermissionConditions) GetInclude() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PermissionConditions) GetIncludeOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Include) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *PermissionConditions) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given map[string]map[string]interface{} and assigns it to the Include field.
func (o *PermissionConditions) SetInclude(v map[string]map[string]interface{}) {
	o.Include = v
}

func (o PermissionConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PermissionConditions) UnmarshalJSON(data []byte) (err error) {
	varPermissionConditions := _PermissionConditions{}

	err = json.Unmarshal(data, &varPermissionConditions)

	if err != nil {
		return err
	}

	*o = PermissionConditions(varPermissionConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePermissionConditions struct {
	value *PermissionConditions
	isSet bool
}

func (v NullablePermissionConditions) Get() *PermissionConditions {
	return v.value
}

func (v *NullablePermissionConditions) Set(val *PermissionConditions) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionConditions) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionConditions(val *PermissionConditions) *NullablePermissionConditions {
	return &NullablePermissionConditions{value: val, isSet: true}
}

func (v NullablePermissionConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
