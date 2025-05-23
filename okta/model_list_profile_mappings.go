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

// ListProfileMappings A collection of the profile mappings that include a subset of the profile mapping object's properties. The Profile Mapping object describes a mapping between an Okta User's and an App User's properties using [JSON Schema Draft 4](https://datatracker.ietf.org/doc/html/draft-zyp-json-schema-04).  > **Note:** Same type source/target mappings aren't supported by this API. Profile mappings must either be Okta->App or App->Okta.
type ListProfileMappings struct {
	// Unique identifier for profile mapping
	Id *string `json:"id,omitempty"`
	Source *ProfileMappingSource `json:"source,omitempty"`
	Target *ProfileMappingTarget `json:"target,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListProfileMappings ListProfileMappings

// NewListProfileMappings instantiates a new ListProfileMappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListProfileMappings() *ListProfileMappings {
	this := ListProfileMappings{}
	return &this
}

// NewListProfileMappingsWithDefaults instantiates a new ListProfileMappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListProfileMappingsWithDefaults() *ListProfileMappings {
	this := ListProfileMappings{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListProfileMappings) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProfileMappings) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListProfileMappings) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListProfileMappings) SetId(v string) {
	o.Id = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ListProfileMappings) GetSource() ProfileMappingSource {
	if o == nil || o.Source == nil {
		var ret ProfileMappingSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProfileMappings) GetSourceOk() (*ProfileMappingSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ListProfileMappings) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given ProfileMappingSource and assigns it to the Source field.
func (o *ListProfileMappings) SetSource(v ProfileMappingSource) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ListProfileMappings) GetTarget() ProfileMappingTarget {
	if o == nil || o.Target == nil {
		var ret ProfileMappingTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProfileMappings) GetTargetOk() (*ProfileMappingTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ListProfileMappings) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given ProfileMappingTarget and assigns it to the Target field.
func (o *ListProfileMappings) SetTarget(v ProfileMappingTarget) {
	o.Target = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ListProfileMappings) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListProfileMappings) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ListProfileMappings) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *ListProfileMappings) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o ListProfileMappings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ListProfileMappings) UnmarshalJSON(bytes []byte) (err error) {
	varListProfileMappings := _ListProfileMappings{}

	err = json.Unmarshal(bytes, &varListProfileMappings)
	if err == nil {
		*o = ListProfileMappings(varListProfileMappings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "source")
		delete(additionalProperties, "target")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableListProfileMappings struct {
	value *ListProfileMappings
	isSet bool
}

func (v NullableListProfileMappings) Get() *ListProfileMappings {
	return v.value
}

func (v *NullableListProfileMappings) Set(val *ListProfileMappings) {
	v.value = val
	v.isSet = true
}

func (v NullableListProfileMappings) IsSet() bool {
	return v.isSet
}

func (v *NullableListProfileMappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListProfileMappings(val *ListProfileMappings) *NullableListProfileMappings {
	return &NullableListProfileMappings{value: val, isSet: true}
}

func (v NullableListProfileMappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListProfileMappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

