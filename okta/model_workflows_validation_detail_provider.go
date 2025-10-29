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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the WorkflowsValidationDetailProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowsValidationDetailProvider{}

// WorkflowsValidationDetailProvider struct for WorkflowsValidationDetailProvider
type WorkflowsValidationDetailProvider struct {
	ValidationDetailProvider
	// Validation error type
	Result               string `json:"result"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowsValidationDetailProvider WorkflowsValidationDetailProvider

// NewWorkflowsValidationDetailProvider instantiates a new WorkflowsValidationDetailProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowsValidationDetailProvider(result string, externalId string, type_ string) *WorkflowsValidationDetailProvider {
	this := WorkflowsValidationDetailProvider{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewWorkflowsValidationDetailProviderWithDefaults instantiates a new WorkflowsValidationDetailProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowsValidationDetailProviderWithDefaults() *WorkflowsValidationDetailProvider {
	this := WorkflowsValidationDetailProvider{}
	return &this
}

// GetResult returns the Result field value
func (o *WorkflowsValidationDetailProvider) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *WorkflowsValidationDetailProvider) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *WorkflowsValidationDetailProvider) SetResult(v string) {
	o.Result = v
}

func (o WorkflowsValidationDetailProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowsValidationDetailProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedValidationDetailProvider, errValidationDetailProvider := json.Marshal(o.ValidationDetailProvider)
	if errValidationDetailProvider != nil {
		return map[string]interface{}{}, errValidationDetailProvider
	}
	errValidationDetailProvider = json.Unmarshal([]byte(serializedValidationDetailProvider), &toSerialize)
	if errValidationDetailProvider != nil {
		return map[string]interface{}{}, errValidationDetailProvider
	}
	toSerialize["result"] = o.Result

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowsValidationDetailProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
		"externalId",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	type WorkflowsValidationDetailProviderWithoutEmbeddedStruct struct {
		// Validation error type
		Result string `json:"result"`
	}

	varWorkflowsValidationDetailProviderWithoutEmbeddedStruct := WorkflowsValidationDetailProviderWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowsValidationDetailProviderWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowsValidationDetailProvider := _WorkflowsValidationDetailProvider{}
		varWorkflowsValidationDetailProvider.Result = varWorkflowsValidationDetailProviderWithoutEmbeddedStruct.Result
		*o = WorkflowsValidationDetailProvider(varWorkflowsValidationDetailProvider)
	} else {
		return err
	}

	varWorkflowsValidationDetailProvider := _WorkflowsValidationDetailProvider{}

	err = json.Unmarshal(data, &varWorkflowsValidationDetailProvider)
	if err == nil {
		o.ValidationDetailProvider = varWorkflowsValidationDetailProvider.ValidationDetailProvider
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")

		// remove fields from embedded structs
		reflectValidationDetailProvider := reflect.ValueOf(o.ValidationDetailProvider)
		for i := 0; i < reflectValidationDetailProvider.Type().NumField(); i++ {
			t := reflectValidationDetailProvider.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowsValidationDetailProvider struct {
	value *WorkflowsValidationDetailProvider
	isSet bool
}

func (v NullableWorkflowsValidationDetailProvider) Get() *WorkflowsValidationDetailProvider {
	return v.value
}

func (v *NullableWorkflowsValidationDetailProvider) Set(val *WorkflowsValidationDetailProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowsValidationDetailProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowsValidationDetailProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowsValidationDetailProvider(val *WorkflowsValidationDetailProvider) *NullableWorkflowsValidationDetailProvider {
	return &NullableWorkflowsValidationDetailProvider{value: val, isSet: true}
}

func (v NullableWorkflowsValidationDetailProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowsValidationDetailProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
