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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// checks if the UserProvisioningApplicationFeature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProvisioningApplicationFeature{}

// UserProvisioningApplicationFeature struct for UserProvisioningApplicationFeature
type UserProvisioningApplicationFeature struct {
	ApplicationFeature
	Capabilities         *CapabilitiesObject `json:"capabilities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserProvisioningApplicationFeature UserProvisioningApplicationFeature

// NewUserProvisioningApplicationFeature instantiates a new UserProvisioningApplicationFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProvisioningApplicationFeature() *UserProvisioningApplicationFeature {
	this := UserProvisioningApplicationFeature{}
	return &this
}

// NewUserProvisioningApplicationFeatureWithDefaults instantiates a new UserProvisioningApplicationFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProvisioningApplicationFeatureWithDefaults() *UserProvisioningApplicationFeature {
	this := UserProvisioningApplicationFeature{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *UserProvisioningApplicationFeature) GetCapabilities() CapabilitiesObject {
	if o == nil || IsNil(o.Capabilities) {
		var ret CapabilitiesObject
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProvisioningApplicationFeature) GetCapabilitiesOk() (*CapabilitiesObject, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *UserProvisioningApplicationFeature) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given CapabilitiesObject and assigns it to the Capabilities field.
func (o *UserProvisioningApplicationFeature) SetCapabilities(v CapabilitiesObject) {
	o.Capabilities = &v
}

func (o UserProvisioningApplicationFeature) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProvisioningApplicationFeature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedApplicationFeature, errApplicationFeature := json.Marshal(o.ApplicationFeature)
	if errApplicationFeature != nil {
		return map[string]interface{}{}, errApplicationFeature
	}
	errApplicationFeature = json.Unmarshal([]byte(serializedApplicationFeature), &toSerialize)
	if errApplicationFeature != nil {
		return map[string]interface{}{}, errApplicationFeature
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserProvisioningApplicationFeature) UnmarshalJSON(data []byte) (err error) {
	type UserProvisioningApplicationFeatureWithoutEmbeddedStruct struct {
		Capabilities *CapabilitiesObject `json:"capabilities,omitempty"`
	}

	varUserProvisioningApplicationFeatureWithoutEmbeddedStruct := UserProvisioningApplicationFeatureWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varUserProvisioningApplicationFeatureWithoutEmbeddedStruct)
	if err == nil {
		varUserProvisioningApplicationFeature := _UserProvisioningApplicationFeature{}
		varUserProvisioningApplicationFeature.Capabilities = varUserProvisioningApplicationFeatureWithoutEmbeddedStruct.Capabilities
		*o = UserProvisioningApplicationFeature(varUserProvisioningApplicationFeature)
	} else {
		return err
	}

	varUserProvisioningApplicationFeature := _UserProvisioningApplicationFeature{}

	err = json.Unmarshal(data, &varUserProvisioningApplicationFeature)
	if err == nil {
		o.ApplicationFeature = varUserProvisioningApplicationFeature.ApplicationFeature
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilities")

		// remove fields from embedded structs
		reflectApplicationFeature := reflect.ValueOf(o.ApplicationFeature)
		for i := 0; i < reflectApplicationFeature.Type().NumField(); i++ {
			t := reflectApplicationFeature.Type().Field(i)

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

type NullableUserProvisioningApplicationFeature struct {
	value *UserProvisioningApplicationFeature
	isSet bool
}

func (v NullableUserProvisioningApplicationFeature) Get() *UserProvisioningApplicationFeature {
	return v.value
}

func (v *NullableUserProvisioningApplicationFeature) Set(val *UserProvisioningApplicationFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProvisioningApplicationFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProvisioningApplicationFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProvisioningApplicationFeature(val *UserProvisioningApplicationFeature) *NullableUserProvisioningApplicationFeature {
	return &NullableUserProvisioningApplicationFeature{value: val, isSet: true}
}

func (v NullableUserProvisioningApplicationFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProvisioningApplicationFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
