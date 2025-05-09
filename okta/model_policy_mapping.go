/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// PolicyMapping struct for PolicyMapping
type PolicyMapping struct {
	Id *string `json:"id,omitempty"`
	Links *PolicyMappingLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyMapping PolicyMapping

// NewPolicyMapping instantiates a new PolicyMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyMapping() *PolicyMapping {
	this := PolicyMapping{}
	return &this
}

// NewPolicyMappingWithDefaults instantiates a new PolicyMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyMappingWithDefaults() *PolicyMapping {
	this := PolicyMapping{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PolicyMapping) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMapping) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PolicyMapping) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PolicyMapping) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PolicyMapping) GetLinks() PolicyMappingLinks {
	if o == nil || o.Links == nil {
		var ret PolicyMappingLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMapping) GetLinksOk() (*PolicyMappingLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PolicyMapping) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PolicyMappingLinks and assigns it to the Links field.
func (o *PolicyMapping) SetLinks(v PolicyMappingLinks) {
	o.Links = &v
}

func (o PolicyMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyMapping) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyMapping := _PolicyMapping{}

	err = json.Unmarshal(bytes, &varPolicyMapping)
	if err == nil {
		*o = PolicyMapping(varPolicyMapping)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyMapping struct {
	value *PolicyMapping
	isSet bool
}

func (v NullablePolicyMapping) Get() *PolicyMapping {
	return v.value
}

func (v *NullablePolicyMapping) Set(val *PolicyMapping) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyMapping) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyMapping(val *PolicyMapping) *NullablePolicyMapping {
	return &NullablePolicyMapping{value: val, isSet: true}
}

func (v NullablePolicyMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

