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

// checks if the FailoverRequestSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailoverRequestSchema{}

// FailoverRequestSchema struct for FailoverRequestSchema
type FailoverRequestSchema struct {
	// List of Okta domains to failover
	Domains              []string `json:"domains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FailoverRequestSchema FailoverRequestSchema

// NewFailoverRequestSchema instantiates a new FailoverRequestSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailoverRequestSchema() *FailoverRequestSchema {
	this := FailoverRequestSchema{}
	return &this
}

// NewFailoverRequestSchemaWithDefaults instantiates a new FailoverRequestSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailoverRequestSchemaWithDefaults() *FailoverRequestSchema {
	this := FailoverRequestSchema{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *FailoverRequestSchema) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FailoverRequestSchema) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *FailoverRequestSchema) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *FailoverRequestSchema) SetDomains(v []string) {
	o.Domains = v
}

func (o FailoverRequestSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailoverRequestSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FailoverRequestSchema) UnmarshalJSON(data []byte) (err error) {
	varFailoverRequestSchema := _FailoverRequestSchema{}

	err = json.Unmarshal(data, &varFailoverRequestSchema)

	if err != nil {
		return err
	}

	*o = FailoverRequestSchema(varFailoverRequestSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFailoverRequestSchema struct {
	value *FailoverRequestSchema
	isSet bool
}

func (v NullableFailoverRequestSchema) Get() *FailoverRequestSchema {
	return v.value
}

func (v *NullableFailoverRequestSchema) Set(val *FailoverRequestSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableFailoverRequestSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableFailoverRequestSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailoverRequestSchema(val *FailoverRequestSchema) *NullableFailoverRequestSchema {
	return &NullableFailoverRequestSchema{value: val, isSet: true}
}

func (v NullableFailoverRequestSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailoverRequestSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
