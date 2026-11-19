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
	"fmt"
)

// checks if the Execution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Execution{}

// Execution struct for Execution
type Execution struct {
	// Human-readable failure reason. `null` unless `status` is `FAILED`.
	Error NullableString `json:"error,omitempty"`
	// Status of the execution. `FAILED` is an execution-level failure (for example, a worker crash), not the outcome of an individual task.
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _Execution Execution

// NewExecution instantiates a new Execution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecution(status string) *Execution {
	this := Execution{}
	this.Status = status
	return &this
}

// NewExecutionWithDefaults instantiates a new Execution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionWithDefaults() *Execution {
	this := Execution{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Execution) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Execution) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *Execution) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *Execution) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *Execution) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *Execution) UnsetError() {
	o.Error.Unset()
}

// GetStatus returns the Status field value
func (o *Execution) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Execution) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Execution) SetStatus(v string) {
	o.Status = v
}

func (o Execution) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Execution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Execution) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExecution := _Execution{}

	err = json.Unmarshal(data, &varExecution)

	if err != nil {
		return err
	}

	*o = Execution(varExecution)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecution struct {
	value *Execution
	isSet bool
}

func (v NullableExecution) Get() *Execution {
	return v.value
}

func (v *NullableExecution) Set(val *Execution) {
	v.value = val
	v.isSet = true
}

func (v NullableExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecution(val *Execution) *NullableExecution {
	return &NullableExecution{value: val, isSet: true}
}

func (v NullableExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
