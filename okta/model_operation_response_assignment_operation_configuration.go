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

// OperationResponseAssignmentOperationConfiguration struct for OperationResponseAssignmentOperationConfiguration
type OperationResponseAssignmentOperationConfiguration struct {
	Actions *OperationResponseAssignmentOperationConfigurationActions `json:"actions,omitempty"`
	Conditions *Conditions `json:"conditions,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OperationResponseAssignmentOperationConfiguration OperationResponseAssignmentOperationConfiguration

// NewOperationResponseAssignmentOperationConfiguration instantiates a new OperationResponseAssignmentOperationConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationResponseAssignmentOperationConfiguration() *OperationResponseAssignmentOperationConfiguration {
	this := OperationResponseAssignmentOperationConfiguration{}
	return &this
}

// NewOperationResponseAssignmentOperationConfigurationWithDefaults instantiates a new OperationResponseAssignmentOperationConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationResponseAssignmentOperationConfigurationWithDefaults() *OperationResponseAssignmentOperationConfiguration {
	this := OperationResponseAssignmentOperationConfiguration{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *OperationResponseAssignmentOperationConfiguration) GetActions() OperationResponseAssignmentOperationConfigurationActions {
	if o == nil || o.Actions == nil {
		var ret OperationResponseAssignmentOperationConfigurationActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseAssignmentOperationConfiguration) GetActionsOk() (*OperationResponseAssignmentOperationConfigurationActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *OperationResponseAssignmentOperationConfiguration) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given OperationResponseAssignmentOperationConfigurationActions and assigns it to the Actions field.
func (o *OperationResponseAssignmentOperationConfiguration) SetActions(v OperationResponseAssignmentOperationConfigurationActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *OperationResponseAssignmentOperationConfiguration) GetConditions() Conditions {
	if o == nil || o.Conditions == nil {
		var ret Conditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseAssignmentOperationConfiguration) GetConditionsOk() (*Conditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *OperationResponseAssignmentOperationConfiguration) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given Conditions and assigns it to the Conditions field.
func (o *OperationResponseAssignmentOperationConfiguration) SetConditions(v Conditions) {
	o.Conditions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OperationResponseAssignmentOperationConfiguration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseAssignmentOperationConfiguration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OperationResponseAssignmentOperationConfiguration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OperationResponseAssignmentOperationConfiguration) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OperationResponseAssignmentOperationConfiguration) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseAssignmentOperationConfiguration) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OperationResponseAssignmentOperationConfiguration) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OperationResponseAssignmentOperationConfiguration) SetName(v string) {
	o.Name = &v
}

func (o OperationResponseAssignmentOperationConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OperationResponseAssignmentOperationConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varOperationResponseAssignmentOperationConfiguration := _OperationResponseAssignmentOperationConfiguration{}

	err = json.Unmarshal(bytes, &varOperationResponseAssignmentOperationConfiguration)
	if err == nil {
		*o = OperationResponseAssignmentOperationConfiguration(varOperationResponseAssignmentOperationConfiguration)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOperationResponseAssignmentOperationConfiguration struct {
	value *OperationResponseAssignmentOperationConfiguration
	isSet bool
}

func (v NullableOperationResponseAssignmentOperationConfiguration) Get() *OperationResponseAssignmentOperationConfiguration {
	return v.value
}

func (v *NullableOperationResponseAssignmentOperationConfiguration) Set(val *OperationResponseAssignmentOperationConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationResponseAssignmentOperationConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationResponseAssignmentOperationConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationResponseAssignmentOperationConfiguration(val *OperationResponseAssignmentOperationConfiguration) *NullableOperationResponseAssignmentOperationConfiguration {
	return &NullableOperationResponseAssignmentOperationConfiguration{value: val, isSet: true}
}

func (v NullableOperationResponseAssignmentOperationConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationResponseAssignmentOperationConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

