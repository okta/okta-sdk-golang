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
)

// checks if the OperationResponseAssignmentOperationConfigurationActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationResponseAssignmentOperationConfigurationActions{}

// OperationResponseAssignmentOperationConfigurationActions Realm assignment action
type OperationResponseAssignmentOperationConfigurationActions struct {
	AssignUserToRealm    *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm `json:"assignUserToRealm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OperationResponseAssignmentOperationConfigurationActions OperationResponseAssignmentOperationConfigurationActions

// NewOperationResponseAssignmentOperationConfigurationActions instantiates a new OperationResponseAssignmentOperationConfigurationActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationResponseAssignmentOperationConfigurationActions() *OperationResponseAssignmentOperationConfigurationActions {
	this := OperationResponseAssignmentOperationConfigurationActions{}
	return &this
}

// NewOperationResponseAssignmentOperationConfigurationActionsWithDefaults instantiates a new OperationResponseAssignmentOperationConfigurationActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationResponseAssignmentOperationConfigurationActionsWithDefaults() *OperationResponseAssignmentOperationConfigurationActions {
	this := OperationResponseAssignmentOperationConfigurationActions{}
	return &this
}

// GetAssignUserToRealm returns the AssignUserToRealm field value if set, zero value otherwise.
func (o *OperationResponseAssignmentOperationConfigurationActions) GetAssignUserToRealm() OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm {
	if o == nil || IsNil(o.AssignUserToRealm) {
		var ret OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm
		return ret
	}
	return *o.AssignUserToRealm
}

// GetAssignUserToRealmOk returns a tuple with the AssignUserToRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseAssignmentOperationConfigurationActions) GetAssignUserToRealmOk() (*OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm, bool) {
	if o == nil || IsNil(o.AssignUserToRealm) {
		return nil, false
	}
	return o.AssignUserToRealm, true
}

// HasAssignUserToRealm returns a boolean if a field has been set.
func (o *OperationResponseAssignmentOperationConfigurationActions) HasAssignUserToRealm() bool {
	if o != nil && !IsNil(o.AssignUserToRealm) {
		return true
	}

	return false
}

// SetAssignUserToRealm gets a reference to the given OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm and assigns it to the AssignUserToRealm field.
func (o *OperationResponseAssignmentOperationConfigurationActions) SetAssignUserToRealm(v OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) {
	o.AssignUserToRealm = &v
}

func (o OperationResponseAssignmentOperationConfigurationActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationResponseAssignmentOperationConfigurationActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignUserToRealm) {
		toSerialize["assignUserToRealm"] = o.AssignUserToRealm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OperationResponseAssignmentOperationConfigurationActions) UnmarshalJSON(data []byte) (err error) {
	varOperationResponseAssignmentOperationConfigurationActions := _OperationResponseAssignmentOperationConfigurationActions{}

	err = json.Unmarshal(data, &varOperationResponseAssignmentOperationConfigurationActions)

	if err != nil {
		return err
	}

	*o = OperationResponseAssignmentOperationConfigurationActions(varOperationResponseAssignmentOperationConfigurationActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assignUserToRealm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOperationResponseAssignmentOperationConfigurationActions struct {
	value *OperationResponseAssignmentOperationConfigurationActions
	isSet bool
}

func (v NullableOperationResponseAssignmentOperationConfigurationActions) Get() *OperationResponseAssignmentOperationConfigurationActions {
	return v.value
}

func (v *NullableOperationResponseAssignmentOperationConfigurationActions) Set(val *OperationResponseAssignmentOperationConfigurationActions) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationResponseAssignmentOperationConfigurationActions) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationResponseAssignmentOperationConfigurationActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationResponseAssignmentOperationConfigurationActions(val *OperationResponseAssignmentOperationConfigurationActions) *NullableOperationResponseAssignmentOperationConfigurationActions {
	return &NullableOperationResponseAssignmentOperationConfigurationActions{value: val, isSet: true}
}

func (v NullableOperationResponseAssignmentOperationConfigurationActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationResponseAssignmentOperationConfigurationActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
