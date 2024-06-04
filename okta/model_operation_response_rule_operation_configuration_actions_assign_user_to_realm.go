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

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// OperationResponseRuleOperationConfigurationActionsAssignUserToRealm struct for OperationResponseRuleOperationConfigurationActionsAssignUserToRealm
type OperationResponseRuleOperationConfigurationActionsAssignUserToRealm struct {
	RealmId              *string `json:"realmId,omitempty"`
	RealmName            *string `json:"realmName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OperationResponseRuleOperationConfigurationActionsAssignUserToRealm OperationResponseRuleOperationConfigurationActionsAssignUserToRealm

// NewOperationResponseRuleOperationConfigurationActionsAssignUserToRealm instantiates a new OperationResponseRuleOperationConfigurationActionsAssignUserToRealm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationResponseRuleOperationConfigurationActionsAssignUserToRealm() *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm {
	this := OperationResponseRuleOperationConfigurationActionsAssignUserToRealm{}
	return &this
}

// NewOperationResponseRuleOperationConfigurationActionsAssignUserToRealmWithDefaults instantiates a new OperationResponseRuleOperationConfigurationActionsAssignUserToRealm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationResponseRuleOperationConfigurationActionsAssignUserToRealmWithDefaults() *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm {
	this := OperationResponseRuleOperationConfigurationActionsAssignUserToRealm{}
	return &this
}

// GetRealmId returns the RealmId field value if set, zero value otherwise.
func (o *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) GetRealmId() string {
	if o == nil || o.RealmId == nil {
		var ret string
		return ret
	}
	return *o.RealmId
}

// GetRealmIdOk returns a tuple with the RealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) GetRealmIdOk() (*string, bool) {
	if o == nil || o.RealmId == nil {
		return nil, false
	}
	return o.RealmId, true
}

// HasRealmId returns a boolean if a field has been set.
func (o *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) HasRealmId() bool {
	if o != nil && o.RealmId != nil {
		return true
	}

	return false
}

// SetRealmId gets a reference to the given string and assigns it to the RealmId field.
func (o *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) SetRealmId(v string) {
	o.RealmId = &v
}

// GetRealmName returns the RealmName field value if set, zero value otherwise.
func (o *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) GetRealmName() string {
	if o == nil || o.RealmName == nil {
		var ret string
		return ret
	}
	return *o.RealmName
}

// GetRealmNameOk returns a tuple with the RealmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) GetRealmNameOk() (*string, bool) {
	if o == nil || o.RealmName == nil {
		return nil, false
	}
	return o.RealmName, true
}

// HasRealmName returns a boolean if a field has been set.
func (o *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) HasRealmName() bool {
	if o != nil && o.RealmName != nil {
		return true
	}

	return false
}

// SetRealmName gets a reference to the given string and assigns it to the RealmName field.
func (o *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) SetRealmName(v string) {
	o.RealmName = &v
}

func (o OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RealmId != nil {
		toSerialize["realmId"] = o.RealmId
	}
	if o.RealmName != nil {
		toSerialize["realmName"] = o.RealmName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) UnmarshalJSON(bytes []byte) (err error) {
	varOperationResponseRuleOperationConfigurationActionsAssignUserToRealm := _OperationResponseRuleOperationConfigurationActionsAssignUserToRealm{}

	err = json.Unmarshal(bytes, &varOperationResponseRuleOperationConfigurationActionsAssignUserToRealm)
	if err == nil {
		*o = OperationResponseRuleOperationConfigurationActionsAssignUserToRealm(varOperationResponseRuleOperationConfigurationActionsAssignUserToRealm)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "realmId")
		delete(additionalProperties, "realmName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm struct {
	value *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm
	isSet bool
}

func (v NullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm) Get() *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm {
	return v.value
}

func (v *NullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm) Set(val *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm(val *OperationResponseRuleOperationConfigurationActionsAssignUserToRealm) *NullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm {
	return &NullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm{value: val, isSet: true}
}

func (v NullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationResponseRuleOperationConfigurationActionsAssignUserToRealm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
