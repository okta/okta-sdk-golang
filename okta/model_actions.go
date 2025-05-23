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

// Actions struct for Actions
type Actions struct {
	AssignUserToRealm *AssignUserToRealm `json:"assignUserToRealm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Actions Actions

// NewActions instantiates a new Actions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActions() *Actions {
	this := Actions{}
	return &this
}

// NewActionsWithDefaults instantiates a new Actions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsWithDefaults() *Actions {
	this := Actions{}
	return &this
}

// GetAssignUserToRealm returns the AssignUserToRealm field value if set, zero value otherwise.
func (o *Actions) GetAssignUserToRealm() AssignUserToRealm {
	if o == nil || o.AssignUserToRealm == nil {
		var ret AssignUserToRealm
		return ret
	}
	return *o.AssignUserToRealm
}

// GetAssignUserToRealmOk returns a tuple with the AssignUserToRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actions) GetAssignUserToRealmOk() (*AssignUserToRealm, bool) {
	if o == nil || o.AssignUserToRealm == nil {
		return nil, false
	}
	return o.AssignUserToRealm, true
}

// HasAssignUserToRealm returns a boolean if a field has been set.
func (o *Actions) HasAssignUserToRealm() bool {
	if o != nil && o.AssignUserToRealm != nil {
		return true
	}

	return false
}

// SetAssignUserToRealm gets a reference to the given AssignUserToRealm and assigns it to the AssignUserToRealm field.
func (o *Actions) SetAssignUserToRealm(v AssignUserToRealm) {
	o.AssignUserToRealm = &v
}

func (o Actions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignUserToRealm != nil {
		toSerialize["assignUserToRealm"] = o.AssignUserToRealm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Actions) UnmarshalJSON(bytes []byte) (err error) {
	varActions := _Actions{}

	err = json.Unmarshal(bytes, &varActions)
	if err == nil {
		*o = Actions(varActions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "assignUserToRealm")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableActions struct {
	value *Actions
	isSet bool
}

func (v NullableActions) Get() *Actions {
	return v.value
}

func (v *NullableActions) Set(val *Actions) {
	v.value = val
	v.isSet = true
}

func (v NullableActions) IsSet() bool {
	return v.isSet
}

func (v *NullableActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActions(val *Actions) *NullableActions {
	return &NullableActions{value: val, isSet: true}
}

func (v NullableActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

