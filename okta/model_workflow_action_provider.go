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

// checks if the WorkflowActionProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowActionProvider{}

// WorkflowActionProvider struct for WorkflowActionProvider
type WorkflowActionProvider struct {
	ActionProvider
	AdditionalProperties map[string]interface{}
}

type _WorkflowActionProvider WorkflowActionProvider

// NewWorkflowActionProvider instantiates a new WorkflowActionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowActionProvider(externalId string, type_ string, url string) *WorkflowActionProvider {
	this := WorkflowActionProvider{}
	this.ExternalId = externalId
	this.Type = type_
	this.Url = url
	return &this
}

// NewWorkflowActionProviderWithDefaults instantiates a new WorkflowActionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowActionProviderWithDefaults() *WorkflowActionProvider {
	this := WorkflowActionProvider{}
	return &this
}

func (o WorkflowActionProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowActionProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedActionProvider, errActionProvider := json.Marshal(o.ActionProvider)
	if errActionProvider != nil {
		return map[string]interface{}{}, errActionProvider
	}
	errActionProvider = json.Unmarshal([]byte(serializedActionProvider), &toSerialize)
	if errActionProvider != nil {
		return map[string]interface{}{}, errActionProvider
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowActionProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalId",
		"type",
		"url",
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

	type WorkflowActionProviderWithoutEmbeddedStruct struct {
	}

	varWorkflowActionProviderWithoutEmbeddedStruct := WorkflowActionProviderWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varWorkflowActionProviderWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowActionProvider := _WorkflowActionProvider{}
		*o = WorkflowActionProvider(varWorkflowActionProvider)
	} else {
		return err
	}

	varWorkflowActionProvider := _WorkflowActionProvider{}

	err = json.Unmarshal(data, &varWorkflowActionProvider)
	if err == nil {
		o.ActionProvider = varWorkflowActionProvider.ActionProvider
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectActionProvider := reflect.ValueOf(o.ActionProvider)
		for i := 0; i < reflectActionProvider.Type().NumField(); i++ {
			t := reflectActionProvider.Type().Field(i)

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

type NullableWorkflowActionProvider struct {
	value *WorkflowActionProvider
	isSet bool
}

func (v NullableWorkflowActionProvider) Get() *WorkflowActionProvider {
	return v.value
}

func (v *NullableWorkflowActionProvider) Set(val *WorkflowActionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowActionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowActionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowActionProvider(val *WorkflowActionProvider) *NullableWorkflowActionProvider {
	return &NullableWorkflowActionProvider{value: val, isSet: true}
}

func (v NullableWorkflowActionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowActionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
