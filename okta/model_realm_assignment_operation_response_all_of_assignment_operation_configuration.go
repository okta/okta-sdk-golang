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

// checks if the RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration{}

// RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration Configuration defintion of the realm
type RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration struct {
	Actions    *RealmAssignmentOperationResponseAllOfAssignmentOperationConfigurationActions `json:"actions,omitempty"`
	Conditions *Conditions                                                                   `json:"conditions,omitempty"`
	// ID of the realm assignment operation
	Id *string `json:"id,omitempty"`
	// Name of the realm assignment operation
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration

// NewRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration instantiates a new RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration() *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration {
	this := RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration{}
	return &this
}

// NewRealmAssignmentOperationResponseAllOfAssignmentOperationConfigurationWithDefaults instantiates a new RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmAssignmentOperationResponseAllOfAssignmentOperationConfigurationWithDefaults() *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration {
	this := RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) GetActions() RealmAssignmentOperationResponseAllOfAssignmentOperationConfigurationActions {
	if o == nil || IsNil(o.Actions) {
		var ret RealmAssignmentOperationResponseAllOfAssignmentOperationConfigurationActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) GetActionsOk() (*RealmAssignmentOperationResponseAllOfAssignmentOperationConfigurationActions, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given RealmAssignmentOperationResponseAllOfAssignmentOperationConfigurationActions and assigns it to the Actions field.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) SetActions(v RealmAssignmentOperationResponseAllOfAssignmentOperationConfigurationActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) GetConditions() Conditions {
	if o == nil || IsNil(o.Conditions) {
		var ret Conditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) GetConditionsOk() (*Conditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given Conditions and assigns it to the Conditions field.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) SetConditions(v Conditions) {
	o.Conditions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) SetName(v string) {
	o.Name = &v
}

func (o RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) UnmarshalJSON(data []byte) (err error) {
	varRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration := _RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration{}

	err = json.Unmarshal(data, &varRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration)

	if err != nil {
		return err
	}

	*o = RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration(varRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration struct {
	value *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration
	isSet bool
}

func (v NullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) Get() *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration {
	return v.value
}

func (v *NullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) Set(val *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration(val *RealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) *NullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration {
	return &NullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration{value: val, isSet: true}
}

func (v NullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmAssignmentOperationResponseAllOfAssignmentOperationConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
