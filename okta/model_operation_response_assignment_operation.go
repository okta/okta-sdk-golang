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

// OperationResponseAssignmentOperation struct for OperationResponseAssignmentOperation
type OperationResponseAssignmentOperation struct {
	Configuration *OperationResponseAssignmentOperationConfiguration `json:"configuration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OperationResponseAssignmentOperation OperationResponseAssignmentOperation

// NewOperationResponseAssignmentOperation instantiates a new OperationResponseAssignmentOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationResponseAssignmentOperation() *OperationResponseAssignmentOperation {
	this := OperationResponseAssignmentOperation{}
	return &this
}

// NewOperationResponseAssignmentOperationWithDefaults instantiates a new OperationResponseAssignmentOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationResponseAssignmentOperationWithDefaults() *OperationResponseAssignmentOperation {
	this := OperationResponseAssignmentOperation{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *OperationResponseAssignmentOperation) GetConfiguration() OperationResponseAssignmentOperationConfiguration {
	if o == nil || o.Configuration == nil {
		var ret OperationResponseAssignmentOperationConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponseAssignmentOperation) GetConfigurationOk() (*OperationResponseAssignmentOperationConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *OperationResponseAssignmentOperation) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given OperationResponseAssignmentOperationConfiguration and assigns it to the Configuration field.
func (o *OperationResponseAssignmentOperation) SetConfiguration(v OperationResponseAssignmentOperationConfiguration) {
	o.Configuration = &v
}

func (o OperationResponseAssignmentOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OperationResponseAssignmentOperation) UnmarshalJSON(bytes []byte) (err error) {
	varOperationResponseAssignmentOperation := _OperationResponseAssignmentOperation{}

	err = json.Unmarshal(bytes, &varOperationResponseAssignmentOperation)
	if err == nil {
		*o = OperationResponseAssignmentOperation(varOperationResponseAssignmentOperation)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "configuration")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOperationResponseAssignmentOperation struct {
	value *OperationResponseAssignmentOperation
	isSet bool
}

func (v NullableOperationResponseAssignmentOperation) Get() *OperationResponseAssignmentOperation {
	return v.value
}

func (v *NullableOperationResponseAssignmentOperation) Set(val *OperationResponseAssignmentOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationResponseAssignmentOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationResponseAssignmentOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationResponseAssignmentOperation(val *OperationResponseAssignmentOperation) *NullableOperationResponseAssignmentOperation {
	return &NullableOperationResponseAssignmentOperation{value: val, isSet: true}
}

func (v NullableOperationResponseAssignmentOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationResponseAssignmentOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

