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

// checks if the EntitlementTypesInnerMappings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementTypesInnerMappings{}

// EntitlementTypesInnerMappings The property mapping between an Okta entitlement and an app entitlement
type EntitlementTypesInnerMappings struct {
	// The field that maps to the entitlement ID
	Id string `json:"id"`
	// The field that maps to the entitlement display name
	DisplayName string `json:"displayName"`
	// The field that maps to entitlement description
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementTypesInnerMappings EntitlementTypesInnerMappings

// NewEntitlementTypesInnerMappings instantiates a new EntitlementTypesInnerMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementTypesInnerMappings(id string, displayName string) *EntitlementTypesInnerMappings {
	this := EntitlementTypesInnerMappings{}
	this.Id = id
	this.DisplayName = displayName
	return &this
}

// NewEntitlementTypesInnerMappingsWithDefaults instantiates a new EntitlementTypesInnerMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementTypesInnerMappingsWithDefaults() *EntitlementTypesInnerMappings {
	this := EntitlementTypesInnerMappings{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementTypesInnerMappings) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInnerMappings) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementTypesInnerMappings) SetId(v string) {
	o.Id = v
}

// GetDisplayName returns the DisplayName field value
func (o *EntitlementTypesInnerMappings) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInnerMappings) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EntitlementTypesInnerMappings) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementTypesInnerMappings) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInnerMappings) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementTypesInnerMappings) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementTypesInnerMappings) SetDescription(v string) {
	o.Description = &v
}

func (o EntitlementTypesInnerMappings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementTypesInnerMappings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementTypesInnerMappings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"displayName",
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

	varEntitlementTypesInnerMappings := _EntitlementTypesInnerMappings{}

	err = json.Unmarshal(data, &varEntitlementTypesInnerMappings)

	if err != nil {
		return err
	}

	*o = EntitlementTypesInnerMappings(varEntitlementTypesInnerMappings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementTypesInnerMappings struct {
	value *EntitlementTypesInnerMappings
	isSet bool
}

func (v NullableEntitlementTypesInnerMappings) Get() *EntitlementTypesInnerMappings {
	return v.value
}

func (v *NullableEntitlementTypesInnerMappings) Set(val *EntitlementTypesInnerMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementTypesInnerMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementTypesInnerMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementTypesInnerMappings(val *EntitlementTypesInnerMappings) *NullableEntitlementTypesInnerMappings {
	return &NullableEntitlementTypesInnerMappings{value: val, isSet: true}
}

func (v NullableEntitlementTypesInnerMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementTypesInnerMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
