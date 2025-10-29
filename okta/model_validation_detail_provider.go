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
)

// checks if the ValidationDetailProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationDetailProvider{}

// ValidationDetailProvider Action provider validation details
type ValidationDetailProvider struct {
	// The unique identifier of the action flow in the provider system
	ExternalId string `json:"externalId"`
	// Type of action provider
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ValidationDetailProvider ValidationDetailProvider

// NewValidationDetailProvider instantiates a new ValidationDetailProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationDetailProvider(externalId string, type_ string) *ValidationDetailProvider {
	this := ValidationDetailProvider{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewValidationDetailProviderWithDefaults instantiates a new ValidationDetailProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationDetailProviderWithDefaults() *ValidationDetailProvider {
	this := ValidationDetailProvider{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *ValidationDetailProvider) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *ValidationDetailProvider) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *ValidationDetailProvider) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *ValidationDetailProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ValidationDetailProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ValidationDetailProvider) SetType(v string) {
	o.Type = v
}

func (o ValidationDetailProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationDetailProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalId"] = o.ExternalId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidationDetailProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varValidationDetailProvider := _ValidationDetailProvider{}

	err = json.Unmarshal(data, &varValidationDetailProvider)

	if err != nil {
		return err
	}

	*o = ValidationDetailProvider(varValidationDetailProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidationDetailProvider struct {
	value *ValidationDetailProvider
	isSet bool
}

func (v NullableValidationDetailProvider) Get() *ValidationDetailProvider {
	return v.value
}

func (v *NullableValidationDetailProvider) Set(val *ValidationDetailProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationDetailProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationDetailProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationDetailProvider(val *ValidationDetailProvider) *NullableValidationDetailProvider {
	return &NullableValidationDetailProvider{value: val, isSet: true}
}

func (v NullableValidationDetailProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationDetailProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
