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

// checks if the LinksGovernanceResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksGovernanceResources{}

// LinksGovernanceResources Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the resources using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification.
type LinksGovernanceResources struct {
	Resources            *HrefObjectGovernanceResourcesLink `json:"resources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksGovernanceResources LinksGovernanceResources

// NewLinksGovernanceResources instantiates a new LinksGovernanceResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksGovernanceResources() *LinksGovernanceResources {
	this := LinksGovernanceResources{}
	return &this
}

// NewLinksGovernanceResourcesWithDefaults instantiates a new LinksGovernanceResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksGovernanceResourcesWithDefaults() *LinksGovernanceResources {
	this := LinksGovernanceResources{}
	return &this
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *LinksGovernanceResources) GetResources() HrefObjectGovernanceResourcesLink {
	if o == nil || IsNil(o.Resources) {
		var ret HrefObjectGovernanceResourcesLink
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksGovernanceResources) GetResourcesOk() (*HrefObjectGovernanceResourcesLink, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *LinksGovernanceResources) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given HrefObjectGovernanceResourcesLink and assigns it to the Resources field.
func (o *LinksGovernanceResources) SetResources(v HrefObjectGovernanceResourcesLink) {
	o.Resources = &v
}

func (o LinksGovernanceResources) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksGovernanceResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksGovernanceResources) UnmarshalJSON(data []byte) (err error) {
	varLinksGovernanceResources := _LinksGovernanceResources{}

	err = json.Unmarshal(data, &varLinksGovernanceResources)

	if err != nil {
		return err
	}

	*o = LinksGovernanceResources(varLinksGovernanceResources)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksGovernanceResources struct {
	value *LinksGovernanceResources
	isSet bool
}

func (v NullableLinksGovernanceResources) Get() *LinksGovernanceResources {
	return v.value
}

func (v *NullableLinksGovernanceResources) Set(val *LinksGovernanceResources) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksGovernanceResources) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksGovernanceResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksGovernanceResources(val *LinksGovernanceResources) *NullableLinksGovernanceResources {
	return &NullableLinksGovernanceResources{value: val, isSet: true}
}

func (v NullableLinksGovernanceResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksGovernanceResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
