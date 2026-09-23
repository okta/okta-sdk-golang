/*
Okta Admin Management API

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

// checks if the CimdRequirementEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CimdRequirementEntry{}

// CimdRequirementEntry A single requirement entry specifying the type and value of a metadata constraint applied to matching CIMD clients
type CimdRequirementEntry struct {
	// The type of the CIMD requirement
	RequirementType string `json:"requirementType"`
	// An [Okta Expression Language](https://developer.okta.com/docs/reference/okta-expression-language/) expression that Okta evaluates against the client's fetched CIMD. The expression must evaluate to a boolean. When the expression returns `false`, Okta rejects the client's authorization request.
	RequirementValue     string `json:"requirementValue"`
	AdditionalProperties map[string]interface{}
}

type _CimdRequirementEntry CimdRequirementEntry

// NewCimdRequirementEntry instantiates a new CimdRequirementEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCimdRequirementEntry(requirementType string, requirementValue string) *CimdRequirementEntry {
	this := CimdRequirementEntry{}
	this.RequirementType = requirementType
	this.RequirementValue = requirementValue
	return &this
}

// NewCimdRequirementEntryWithDefaults instantiates a new CimdRequirementEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCimdRequirementEntryWithDefaults() *CimdRequirementEntry {
	this := CimdRequirementEntry{}
	return &this
}

// GetRequirementType returns the RequirementType field value
func (o *CimdRequirementEntry) GetRequirementType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequirementType
}

// GetRequirementTypeOk returns a tuple with the RequirementType field value
// and a boolean to check if the value has been set.
func (o *CimdRequirementEntry) GetRequirementTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequirementType, true
}

// SetRequirementType sets field value
func (o *CimdRequirementEntry) SetRequirementType(v string) {
	o.RequirementType = v
}

// GetRequirementValue returns the RequirementValue field value
func (o *CimdRequirementEntry) GetRequirementValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequirementValue
}

// GetRequirementValueOk returns a tuple with the RequirementValue field value
// and a boolean to check if the value has been set.
func (o *CimdRequirementEntry) GetRequirementValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequirementValue, true
}

// SetRequirementValue sets field value
func (o *CimdRequirementEntry) SetRequirementValue(v string) {
	o.RequirementValue = v
}

func (o CimdRequirementEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CimdRequirementEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requirementType"] = o.RequirementType
	toSerialize["requirementValue"] = o.RequirementValue

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CimdRequirementEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requirementType",
		"requirementValue",
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

	varCimdRequirementEntry := _CimdRequirementEntry{}

	err = json.Unmarshal(data, &varCimdRequirementEntry)

	if err != nil {
		return err
	}

	*o = CimdRequirementEntry(varCimdRequirementEntry)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requirementType")
		delete(additionalProperties, "requirementValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCimdRequirementEntry struct {
	value *CimdRequirementEntry
	isSet bool
}

func (v NullableCimdRequirementEntry) Get() *CimdRequirementEntry {
	return v.value
}

func (v *NullableCimdRequirementEntry) Set(val *CimdRequirementEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableCimdRequirementEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableCimdRequirementEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCimdRequirementEntry(val *CimdRequirementEntry) *NullableCimdRequirementEntry {
	return &NullableCimdRequirementEntry{value: val, isSet: true}
}

func (v NullableCimdRequirementEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCimdRequirementEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
