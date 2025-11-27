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
	"fmt"
)

// checks if the AIAgentProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIAgentProfile{}

// AIAgentProfile AI agent profile
type AIAgentProfile struct {
	// Description of the AI agent
	Description *string `json:"description,omitempty"`
	// Unique name of the AI agent
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _AIAgentProfile AIAgentProfile

// NewAIAgentProfile instantiates a new AIAgentProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIAgentProfile(name string) *AIAgentProfile {
	this := AIAgentProfile{}
	this.Name = name
	return &this
}

// NewAIAgentProfileWithDefaults instantiates a new AIAgentProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIAgentProfileWithDefaults() *AIAgentProfile {
	this := AIAgentProfile{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AIAgentProfile) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIAgentProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AIAgentProfile) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AIAgentProfile) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *AIAgentProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AIAgentProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AIAgentProfile) SetName(v string) {
	o.Name = v
}

func (o AIAgentProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIAgentProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIAgentProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAIAgentProfile := _AIAgentProfile{}

	err = json.Unmarshal(data, &varAIAgentProfile)

	if err != nil {
		return err
	}

	*o = AIAgentProfile(varAIAgentProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIAgentProfile struct {
	value *AIAgentProfile
	isSet bool
}

func (v NullableAIAgentProfile) Get() *AIAgentProfile {
	return v.value
}

func (v *NullableAIAgentProfile) Set(val *AIAgentProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAIAgentProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAIAgentProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIAgentProfile(val *AIAgentProfile) *NullableAIAgentProfile {
	return &NullableAIAgentProfile{value: val, isSet: true}
}

func (v NullableAIAgentProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIAgentProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
