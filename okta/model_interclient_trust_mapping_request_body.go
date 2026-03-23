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
)

// checks if the InterclientTrustMappingRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterclientTrustMappingRequestBody{}

// InterclientTrustMappingRequestBody struct for InterclientTrustMappingRequestBody
type InterclientTrustMappingRequestBody struct {
	// App ID of the allowed app
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterclientTrustMappingRequestBody InterclientTrustMappingRequestBody

// NewInterclientTrustMappingRequestBody instantiates a new InterclientTrustMappingRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterclientTrustMappingRequestBody() *InterclientTrustMappingRequestBody {
	this := InterclientTrustMappingRequestBody{}
	return &this
}

// NewInterclientTrustMappingRequestBodyWithDefaults instantiates a new InterclientTrustMappingRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterclientTrustMappingRequestBodyWithDefaults() *InterclientTrustMappingRequestBody {
	this := InterclientTrustMappingRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InterclientTrustMappingRequestBody) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterclientTrustMappingRequestBody) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InterclientTrustMappingRequestBody) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InterclientTrustMappingRequestBody) SetId(v string) {
	o.Id = &v
}

func (o InterclientTrustMappingRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterclientTrustMappingRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterclientTrustMappingRequestBody) UnmarshalJSON(data []byte) (err error) {
	varInterclientTrustMappingRequestBody := _InterclientTrustMappingRequestBody{}

	err = json.Unmarshal(data, &varInterclientTrustMappingRequestBody)

	if err != nil {
		return err
	}

	*o = InterclientTrustMappingRequestBody(varInterclientTrustMappingRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterclientTrustMappingRequestBody struct {
	value *InterclientTrustMappingRequestBody
	isSet bool
}

func (v NullableInterclientTrustMappingRequestBody) Get() *InterclientTrustMappingRequestBody {
	return v.value
}

func (v *NullableInterclientTrustMappingRequestBody) Set(val *InterclientTrustMappingRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableInterclientTrustMappingRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableInterclientTrustMappingRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterclientTrustMappingRequestBody(val *InterclientTrustMappingRequestBody) *NullableInterclientTrustMappingRequestBody {
	return &NullableInterclientTrustMappingRequestBody{value: val, isSet: true}
}

func (v NullableInterclientTrustMappingRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterclientTrustMappingRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
