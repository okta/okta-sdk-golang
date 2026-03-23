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

// checks if the LogUserBehavior type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogUserBehavior{}

// LogUserBehavior The result of the user behavior detection models associated with the event
type LogUserBehavior struct {
	// The unique identifier of the user behavior detection model
	Id NullableString `json:"id,omitempty"`
	// The name of the user behavior detection model [configured by admins](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Behavior/)
	Name NullableString `json:"name,omitempty"`
	// The result of the user behavior analysis
	Result               NullableString `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogUserBehavior LogUserBehavior

// NewLogUserBehavior instantiates a new LogUserBehavior object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogUserBehavior() *LogUserBehavior {
	this := LogUserBehavior{}
	return &this
}

// NewLogUserBehaviorWithDefaults instantiates a new LogUserBehavior object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogUserBehaviorWithDefaults() *LogUserBehavior {
	this := LogUserBehavior{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogUserBehavior) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogUserBehavior) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *LogUserBehavior) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *LogUserBehavior) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *LogUserBehavior) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *LogUserBehavior) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogUserBehavior) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogUserBehavior) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *LogUserBehavior) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *LogUserBehavior) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *LogUserBehavior) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *LogUserBehavior) UnsetName() {
	o.Name.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogUserBehavior) GetResult() string {
	if o == nil || IsNil(o.Result.Get()) {
		var ret string
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogUserBehavior) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *LogUserBehavior) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableString and assigns it to the Result field.
func (o *LogUserBehavior) SetResult(v string) {
	o.Result.Set(&v)
}

// SetResultNil sets the value for Result to be an explicit nil
func (o *LogUserBehavior) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *LogUserBehavior) UnsetResult() {
	o.Result.Unset()
}

func (o LogUserBehavior) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogUserBehavior) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogUserBehavior) UnmarshalJSON(data []byte) (err error) {
	varLogUserBehavior := _LogUserBehavior{}

	err = json.Unmarshal(data, &varLogUserBehavior)

	if err != nil {
		return err
	}

	*o = LogUserBehavior(varLogUserBehavior)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogUserBehavior struct {
	value *LogUserBehavior
	isSet bool
}

func (v NullableLogUserBehavior) Get() *LogUserBehavior {
	return v.value
}

func (v *NullableLogUserBehavior) Set(val *LogUserBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableLogUserBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableLogUserBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogUserBehavior(val *LogUserBehavior) *NullableLogUserBehavior {
	return &NullableLogUserBehavior{value: val, isSet: true}
}

func (v NullableLogUserBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogUserBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
