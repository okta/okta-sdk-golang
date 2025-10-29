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

// checks if the AvailableAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableAction{}

// AvailableAction struct for AvailableAction
type AvailableAction struct {
	// Action identifier
	Id                   string                          `json:"id"`
	Provider             WorkflowAvailableActionProvider `json:"provider"`
	AdditionalProperties map[string]interface{}
}

type _AvailableAction AvailableAction

// NewAvailableAction instantiates a new AvailableAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableAction(id string, provider WorkflowAvailableActionProvider) *AvailableAction {
	this := AvailableAction{}
	this.Id = id
	this.Provider = provider
	return &this
}

// NewAvailableActionWithDefaults instantiates a new AvailableAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableActionWithDefaults() *AvailableAction {
	this := AvailableAction{}
	return &this
}

// GetId returns the Id field value
func (o *AvailableAction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AvailableAction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AvailableAction) SetId(v string) {
	o.Id = v
}

// GetProvider returns the Provider field value
func (o *AvailableAction) GetProvider() WorkflowAvailableActionProvider {
	if o == nil {
		var ret WorkflowAvailableActionProvider
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AvailableAction) GetProviderOk() (*WorkflowAvailableActionProvider, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *AvailableAction) SetProvider(v WorkflowAvailableActionProvider) {
	o.Provider = v
}

func (o AvailableAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailableAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["provider"] = o.Provider

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AvailableAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"provider",
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

	varAvailableAction := _AvailableAction{}

	err = json.Unmarshal(data, &varAvailableAction)

	if err != nil {
		return err
	}

	*o = AvailableAction(varAvailableAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "provider")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAvailableAction struct {
	value *AvailableAction
	isSet bool
}

func (v NullableAvailableAction) Get() *AvailableAction {
	return v.value
}

func (v *NullableAvailableAction) Set(val *AvailableAction) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableAction) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableAction(val *AvailableAction) *NullableAvailableAction {
	return &NullableAvailableAction{value: val, isSet: true}
}

func (v NullableAvailableAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
