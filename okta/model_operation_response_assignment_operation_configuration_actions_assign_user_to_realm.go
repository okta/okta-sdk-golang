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

// checks if the OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm{}

// OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm Action that assigns a user to a realm
type OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm struct {
	// ID of the realm
	RealmId              *string `json:"realmId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm

// NewOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm instantiates a new OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm() *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm {
	this := OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm{}
	return &this
}

// NewOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealmWithDefaults instantiates a new OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealmWithDefaults() *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm {
	this := OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm{}
	return &this
}

// GetRealmId returns the RealmId field value if set, zero value otherwise.
func (o *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) GetRealmId() string {
	if o == nil || IsNil(o.RealmId) {
		var ret string
		return ret
	}
	return *o.RealmId
}

// GetRealmIdOk returns a tuple with the RealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) GetRealmIdOk() (*string, bool) {
	if o == nil || IsNil(o.RealmId) {
		return nil, false
	}
	return o.RealmId, true
}

// HasRealmId returns a boolean if a field has been set.
func (o *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) HasRealmId() bool {
	if o != nil && !IsNil(o.RealmId) {
		return true
	}

	return false
}

// SetRealmId gets a reference to the given string and assigns it to the RealmId field.
func (o *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) SetRealmId(v string) {
	o.RealmId = &v
}

func (o OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RealmId) {
		toSerialize["realmId"] = o.RealmId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) UnmarshalJSON(data []byte) (err error) {
	varOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm := _OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm{}

	err = json.Unmarshal(data, &varOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm)

	if err != nil {
		return err
	}

	*o = OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm(varOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "realmId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm struct {
	value *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm
	isSet bool
}

func (v NullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) Get() *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm {
	return v.value
}

func (v *NullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) Set(val *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm(val *OperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) *NullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm {
	return &NullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm{value: val, isSet: true}
}

func (v NullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationResponseAssignmentOperationConfigurationActionsAssignUserToRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
