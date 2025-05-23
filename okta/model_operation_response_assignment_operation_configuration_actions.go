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

// OperationResponseAssignmentOperationConfigurationActions struct for OperationResponseAssignmentOperationConfigurationActions
type OperationResponseAssignmentOperationConfigurationActions struct {
	AssignUserToRealm *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm `json:"assignUserToRealm,omitempty"`
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
	if o == nil || o.AssignUserToRealm == nil {
		var ret OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm
		return ret
	}
	return *o.AssignUserToRealm
}

// GetAssignUserToRealmOk returns a tuple with the AssignUserToRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseAssignmentOperationConfigurationActions) GetAssignUserToRealmOk() (*OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm, bool) {
	if o == nil || o.AssignUserToRealm == nil {
		return nil, false
	}
	return o.AssignUserToRealm, true
}

// HasAssignUserToRealm returns a boolean if a field has been set.
func (o *OperationResponseAssignmentOperationConfigurationActions) HasAssignUserToRealm() bool {
	if o != nil && o.AssignUserToRealm != nil {
		return true
	}

	return false
}

// SetAssignUserToRealm gets a reference to the given OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm and assigns it to the AssignUserToRealm field.
func (o *OperationResponseAssignmentOperationConfigurationActions) SetAssignUserToRealm(v OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) {
	o.AssignUserToRealm = &v
}

func (o OperationResponseAssignmentOperationConfigurationActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignUserToRealm != nil {
		toSerialize["assignUserToRealm"] = o.AssignUserToRealm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OperationResponseAssignmentOperationConfigurationActions) UnmarshalJSON(bytes []byte) (err error) {
	varOperationResponseAssignmentOperationConfigurationActions := _OperationResponseAssignmentOperationConfigurationActions{}

	err = json.Unmarshal(bytes, &varOperationResponseAssignmentOperationConfigurationActions)
	if err == nil {
		*o = OperationResponseAssignmentOperationConfigurationActions(varOperationResponseAssignmentOperationConfigurationActions)
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

