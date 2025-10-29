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
)

// checks if the SAMLPayLoadDataAssertionConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataAssertionConditions{}

// SAMLPayLoadDataAssertionConditions Provides a JSON representation of the `<saml:Conditions>` element of the SAML assertion
type SAMLPayLoadDataAssertionConditions struct {
	// Describes which service providers the assertion is valid for
	AudienceRestriction  []string `json:"audienceRestriction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataAssertionConditions SAMLPayLoadDataAssertionConditions

// NewSAMLPayLoadDataAssertionConditions instantiates a new SAMLPayLoadDataAssertionConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataAssertionConditions() *SAMLPayLoadDataAssertionConditions {
	this := SAMLPayLoadDataAssertionConditions{}
	return &this
}

// NewSAMLPayLoadDataAssertionConditionsWithDefaults instantiates a new SAMLPayLoadDataAssertionConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataAssertionConditionsWithDefaults() *SAMLPayLoadDataAssertionConditions {
	this := SAMLPayLoadDataAssertionConditions{}
	return &this
}

// GetAudienceRestriction returns the AudienceRestriction field value if set, zero value otherwise.
func (o *SAMLPayLoadDataAssertionConditions) GetAudienceRestriction() []string {
	if o == nil || IsNil(o.AudienceRestriction) {
		var ret []string
		return ret
	}
	return o.AudienceRestriction
}

// GetAudienceRestrictionOk returns a tuple with the AudienceRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataAssertionConditions) GetAudienceRestrictionOk() ([]string, bool) {
	if o == nil || IsNil(o.AudienceRestriction) {
		return nil, false
	}
	return o.AudienceRestriction, true
}

// HasAudienceRestriction returns a boolean if a field has been set.
func (o *SAMLPayLoadDataAssertionConditions) HasAudienceRestriction() bool {
	if o != nil && !IsNil(o.AudienceRestriction) {
		return true
	}

	return false
}

// SetAudienceRestriction gets a reference to the given []string and assigns it to the AudienceRestriction field.
func (o *SAMLPayLoadDataAssertionConditions) SetAudienceRestriction(v []string) {
	o.AudienceRestriction = v
}

func (o SAMLPayLoadDataAssertionConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataAssertionConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudienceRestriction) {
		toSerialize["audienceRestriction"] = o.AudienceRestriction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataAssertionConditions) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataAssertionConditions := _SAMLPayLoadDataAssertionConditions{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataAssertionConditions)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataAssertionConditions(varSAMLPayLoadDataAssertionConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "audienceRestriction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataAssertionConditions struct {
	value *SAMLPayLoadDataAssertionConditions
	isSet bool
}

func (v NullableSAMLPayLoadDataAssertionConditions) Get() *SAMLPayLoadDataAssertionConditions {
	return v.value
}

func (v *NullableSAMLPayLoadDataAssertionConditions) Set(val *SAMLPayLoadDataAssertionConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataAssertionConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataAssertionConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataAssertionConditions(val *SAMLPayLoadDataAssertionConditions) *NullableSAMLPayLoadDataAssertionConditions {
	return &NullableSAMLPayLoadDataAssertionConditions{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataAssertionConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataAssertionConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
