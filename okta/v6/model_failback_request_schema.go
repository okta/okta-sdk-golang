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
)

// checks if the FailbackRequestSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailbackRequestSchema{}

// FailbackRequestSchema struct for FailbackRequestSchema
type FailbackRequestSchema struct {
	// List of Okta domains to failback
	Domains              []string `json:"domains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FailbackRequestSchema FailbackRequestSchema

// NewFailbackRequestSchema instantiates a new FailbackRequestSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailbackRequestSchema() *FailbackRequestSchema {
	this := FailbackRequestSchema{}
	return &this
}

// NewFailbackRequestSchemaWithDefaults instantiates a new FailbackRequestSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailbackRequestSchemaWithDefaults() *FailbackRequestSchema {
	this := FailbackRequestSchema{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *FailbackRequestSchema) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailbackRequestSchema) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *FailbackRequestSchema) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *FailbackRequestSchema) SetDomains(v []string) {
	o.Domains = v
}

func (o FailbackRequestSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailbackRequestSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FailbackRequestSchema) UnmarshalJSON(data []byte) (err error) {
	varFailbackRequestSchema := _FailbackRequestSchema{}

	err = json.Unmarshal(data, &varFailbackRequestSchema)

	if err != nil {
		return err
	}

	*o = FailbackRequestSchema(varFailbackRequestSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFailbackRequestSchema struct {
	value *FailbackRequestSchema
	isSet bool
}

func (v NullableFailbackRequestSchema) Get() *FailbackRequestSchema {
	return v.value
}

func (v *NullableFailbackRequestSchema) Set(val *FailbackRequestSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFailbackRequestSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFailbackRequestSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailbackRequestSchema(val *FailbackRequestSchema) *NullableFailbackRequestSchema {
	return &NullableFailbackRequestSchema{value: val, isSet: true}
}

func (v NullableFailbackRequestSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailbackRequestSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
