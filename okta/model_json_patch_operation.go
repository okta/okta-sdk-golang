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

// JsonPatchOperation The update action
type JsonPatchOperation struct {
	// The operation (PATCH action)
	Op *string `json:"op,omitempty"`
	// The resource path of the attribute to update
	Path *string `json:"path,omitempty"`
	// The update operation value
	Value map[string]interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JsonPatchOperation JsonPatchOperation

// NewJsonPatchOperation instantiates a new JsonPatchOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonPatchOperation() *JsonPatchOperation {
	this := JsonPatchOperation{}
	return &this
}

// NewJsonPatchOperationWithDefaults instantiates a new JsonPatchOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonPatchOperationWithDefaults() *JsonPatchOperation {
	this := JsonPatchOperation{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *JsonPatchOperation) GetOp() string {
	if o == nil || o.Op == nil {
		var ret string
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchOperation) GetOpOk() (*string, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *JsonPatchOperation) HasOp() bool {
	if o != nil && o.Op != nil {
		return true
	}

	return false
}

// SetOp gets a reference to the given string and assigns it to the Op field.
func (o *JsonPatchOperation) SetOp(v string) {
	o.Op = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *JsonPatchOperation) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchOperation) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *JsonPatchOperation) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *JsonPatchOperation) SetPath(v string) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *JsonPatchOperation) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchOperation) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *JsonPatchOperation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *JsonPatchOperation) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o JsonPatchOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *JsonPatchOperation) UnmarshalJSON(bytes []byte) (err error) {
	varJsonPatchOperation := _JsonPatchOperation{}

	err = json.Unmarshal(bytes, &varJsonPatchOperation)
	if err == nil {
		*o = JsonPatchOperation(varJsonPatchOperation)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableJsonPatchOperation struct {
	value *JsonPatchOperation
	isSet bool
}

func (v NullableJsonPatchOperation) Get() *JsonPatchOperation {
	return v.value
}

func (v *NullableJsonPatchOperation) Set(val *JsonPatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonPatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonPatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonPatchOperation(val *JsonPatchOperation) *NullableJsonPatchOperation {
	return &NullableJsonPatchOperation{value: val, isSet: true}
}

func (v NullableJsonPatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonPatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

