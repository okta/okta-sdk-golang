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
	"reflect"
	"strings"
)

// checks if the IdProofingMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdProofingMethod{}

// IdProofingMethod struct for IdProofingMethod
type IdProofingMethod struct {
	VerificationMethod
	// ID for ID proofing entity
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdProofingMethod IdProofingMethod

// NewIdProofingMethod instantiates a new IdProofingMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdProofingMethod() *IdProofingMethod {
	this := IdProofingMethod{}
	return &this
}

// NewIdProofingMethodWithDefaults instantiates a new IdProofingMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdProofingMethodWithDefaults() *IdProofingMethod {
	this := IdProofingMethod{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdProofingMethod) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdProofingMethod) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdProofingMethod) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdProofingMethod) SetId(v string) {
	o.Id = &v
}

func (o IdProofingMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdProofingMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedVerificationMethod, errVerificationMethod := json.Marshal(o.VerificationMethod)
	if errVerificationMethod != nil {
		return map[string]interface{}{}, errVerificationMethod
	}
	errVerificationMethod = json.Unmarshal([]byte(serializedVerificationMethod), &toSerialize)
	if errVerificationMethod != nil {
		return map[string]interface{}{}, errVerificationMethod
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdProofingMethod) UnmarshalJSON(data []byte) (err error) {
	type IdProofingMethodWithoutEmbeddedStruct struct {
		// ID for ID proofing entity
		Id *string `json:"id,omitempty"`
	}

	varIdProofingMethodWithoutEmbeddedStruct := IdProofingMethodWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIdProofingMethodWithoutEmbeddedStruct)
	if err == nil {
		varIdProofingMethod := _IdProofingMethod{}
		varIdProofingMethod.Id = varIdProofingMethodWithoutEmbeddedStruct.Id
		*o = IdProofingMethod(varIdProofingMethod)
	} else {
		return err
	}

	varIdProofingMethod := _IdProofingMethod{}

	err = json.Unmarshal(data, &varIdProofingMethod)
	if err == nil {
		o.VerificationMethod = varIdProofingMethod.VerificationMethod
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")

		// remove fields from embedded structs
		reflectVerificationMethod := reflect.ValueOf(o.VerificationMethod)
		for i := 0; i < reflectVerificationMethod.Type().NumField(); i++ {
			t := reflectVerificationMethod.Type().Field(i)

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

type NullableIdProofingMethod struct {
	value *IdProofingMethod
	isSet bool
}

func (v NullableIdProofingMethod) Get() *IdProofingMethod {
	return v.value
}

func (v *NullableIdProofingMethod) Set(val *IdProofingMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableIdProofingMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableIdProofingMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdProofingMethod(val *IdProofingMethod) *NullableIdProofingMethod {
	return &NullableIdProofingMethod{value: val, isSet: true}
}

func (v NullableIdProofingMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdProofingMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
