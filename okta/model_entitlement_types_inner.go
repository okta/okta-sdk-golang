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

// checks if the EntitlementTypesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementTypesInner{}

// EntitlementTypesInner struct for EntitlementTypesInner
type EntitlementTypesInner struct {
	// The entitlement type name
	Name string `json:"name"`
	// Description of the entitlement type
	Description *string `json:"description,omitempty"`
	// URL of the entitlement type endpoint
	Endpoint             string                          `json:"endpoint"`
	Attributes           EntitlementTypesInnerAttributes `json:"attributes"`
	Mappings             EntitlementTypesInnerMappings   `json:"mappings"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementTypesInner EntitlementTypesInner

// NewEntitlementTypesInner instantiates a new EntitlementTypesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementTypesInner(name string, endpoint string, attributes EntitlementTypesInnerAttributes, mappings EntitlementTypesInnerMappings) *EntitlementTypesInner {
	this := EntitlementTypesInner{}
	this.Name = name
	this.Endpoint = endpoint
	this.Attributes = attributes
	this.Mappings = mappings
	return &this
}

// NewEntitlementTypesInnerWithDefaults instantiates a new EntitlementTypesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementTypesInnerWithDefaults() *EntitlementTypesInner {
	this := EntitlementTypesInner{}
	return &this
}

// GetName returns the Name field value
func (o *EntitlementTypesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementTypesInner) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementTypesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementTypesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementTypesInner) SetDescription(v string) {
	o.Description = &v
}

// GetEndpoint returns the Endpoint field value
func (o *EntitlementTypesInner) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInner) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *EntitlementTypesInner) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetAttributes returns the Attributes field value
func (o *EntitlementTypesInner) GetAttributes() EntitlementTypesInnerAttributes {
	if o == nil {
		var ret EntitlementTypesInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInner) GetAttributesOk() (*EntitlementTypesInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *EntitlementTypesInner) SetAttributes(v EntitlementTypesInnerAttributes) {
	o.Attributes = v
}

// GetMappings returns the Mappings field value
func (o *EntitlementTypesInner) GetMappings() EntitlementTypesInnerMappings {
	if o == nil {
		var ret EntitlementTypesInnerMappings
		return ret
	}

	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *EntitlementTypesInner) GetMappingsOk() (*EntitlementTypesInnerMappings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// SetMappings sets field value
func (o *EntitlementTypesInner) SetMappings(v EntitlementTypesInnerMappings) {
	o.Mappings = v
}

func (o EntitlementTypesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementTypesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["attributes"] = o.Attributes
	toSerialize["mappings"] = o.Mappings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementTypesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"endpoint",
		"attributes",
		"mappings",
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

	varEntitlementTypesInner := _EntitlementTypesInner{}

	err = json.Unmarshal(data, &varEntitlementTypesInner)

	if err != nil {
		return err
	}

	*o = EntitlementTypesInner(varEntitlementTypesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "mappings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementTypesInner struct {
	value *EntitlementTypesInner
	isSet bool
}

func (v NullableEntitlementTypesInner) Get() *EntitlementTypesInner {
	return v.value
}

func (v *NullableEntitlementTypesInner) Set(val *EntitlementTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementTypesInner(val *EntitlementTypesInner) *NullableEntitlementTypesInner {
	return &NullableEntitlementTypesInner{value: val, isSet: true}
}

func (v NullableEntitlementTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
