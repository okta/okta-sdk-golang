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

// ContinuousAccessFailureActionsObject struct for ContinuousAccessFailureActionsObject
type ContinuousAccessFailureActionsObject struct {
	Action *string `json:"action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContinuousAccessFailureActionsObject ContinuousAccessFailureActionsObject

// NewContinuousAccessFailureActionsObject instantiates a new ContinuousAccessFailureActionsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinuousAccessFailureActionsObject() *ContinuousAccessFailureActionsObject {
	this := ContinuousAccessFailureActionsObject{}
	return &this
}

// NewContinuousAccessFailureActionsObjectWithDefaults instantiates a new ContinuousAccessFailureActionsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinuousAccessFailureActionsObjectWithDefaults() *ContinuousAccessFailureActionsObject {
	this := ContinuousAccessFailureActionsObject{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ContinuousAccessFailureActionsObject) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinuousAccessFailureActionsObject) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ContinuousAccessFailureActionsObject) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ContinuousAccessFailureActionsObject) SetAction(v string) {
	o.Action = &v
}

func (o ContinuousAccessFailureActionsObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ContinuousAccessFailureActionsObject) UnmarshalJSON(bytes []byte) (err error) {
	varContinuousAccessFailureActionsObject := _ContinuousAccessFailureActionsObject{}

	err = json.Unmarshal(bytes, &varContinuousAccessFailureActionsObject)
	if err == nil {
		*o = ContinuousAccessFailureActionsObject(varContinuousAccessFailureActionsObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableContinuousAccessFailureActionsObject struct {
	value *ContinuousAccessFailureActionsObject
	isSet bool
}

func (v NullableContinuousAccessFailureActionsObject) Get() *ContinuousAccessFailureActionsObject {
	return v.value
}

func (v *NullableContinuousAccessFailureActionsObject) Set(val *ContinuousAccessFailureActionsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousAccessFailureActionsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousAccessFailureActionsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousAccessFailureActionsObject(val *ContinuousAccessFailureActionsObject) *NullableContinuousAccessFailureActionsObject {
	return &NullableContinuousAccessFailureActionsObject{value: val, isSet: true}
}

func (v NullableContinuousAccessFailureActionsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousAccessFailureActionsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

