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

// PasswordImportRequestData struct for PasswordImportRequestData
type PasswordImportRequestData struct {
	Action *PasswordImportRequestDataAction `json:"action,omitempty"`
	Context *PasswordImportRequestDataContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordImportRequestData PasswordImportRequestData

// NewPasswordImportRequestData instantiates a new PasswordImportRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordImportRequestData() *PasswordImportRequestData {
	this := PasswordImportRequestData{}
	return &this
}

// NewPasswordImportRequestDataWithDefaults instantiates a new PasswordImportRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordImportRequestDataWithDefaults() *PasswordImportRequestData {
	this := PasswordImportRequestData{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PasswordImportRequestData) GetAction() PasswordImportRequestDataAction {
	if o == nil || o.Action == nil {
		var ret PasswordImportRequestDataAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestData) GetActionOk() (*PasswordImportRequestDataAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PasswordImportRequestData) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given PasswordImportRequestDataAction and assigns it to the Action field.
func (o *PasswordImportRequestData) SetAction(v PasswordImportRequestDataAction) {
	o.Action = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PasswordImportRequestData) GetContext() PasswordImportRequestDataContext {
	if o == nil || o.Context == nil {
		var ret PasswordImportRequestDataContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequestData) GetContextOk() (*PasswordImportRequestDataContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PasswordImportRequestData) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given PasswordImportRequestDataContext and assigns it to the Context field.
func (o *PasswordImportRequestData) SetContext(v PasswordImportRequestDataContext) {
	o.Context = &v
}

func (o PasswordImportRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordImportRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordImportRequestData := _PasswordImportRequestData{}

	err = json.Unmarshal(bytes, &varPasswordImportRequestData)
	if err == nil {
		*o = PasswordImportRequestData(varPasswordImportRequestData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordImportRequestData struct {
	value *PasswordImportRequestData
	isSet bool
}

func (v NullablePasswordImportRequestData) Get() *PasswordImportRequestData {
	return v.value
}

func (v *NullablePasswordImportRequestData) Set(val *PasswordImportRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordImportRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordImportRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordImportRequestData(val *PasswordImportRequestData) *NullablePasswordImportRequestData {
	return &NullablePasswordImportRequestData{value: val, isSet: true}
}

func (v NullablePasswordImportRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordImportRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

