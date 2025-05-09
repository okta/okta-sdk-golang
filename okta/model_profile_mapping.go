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

// ProfileMapping The Profile Mapping object describes a mapping between an Okta User's and an App User's properties using [JSON Schema Draft 4](https://datatracker.ietf.org/doc/html/draft-zyp-json-schema-04).  > **Note:** Same type source/target mappings aren't supported by this API. Profile mappings must either be Okta->App or App->Okta.
type ProfileMapping struct {
	// Unique identifier for a profile mapping
	Id *string `json:"id,omitempty"`
	Properties *map[string]ProfileMappingProperty `json:"properties,omitempty"`
	Source *ProfileMappingSource `json:"source,omitempty"`
	Target *ProfileMappingTarget `json:"target,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileMapping ProfileMapping

// NewProfileMapping instantiates a new ProfileMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileMapping() *ProfileMapping {
	this := ProfileMapping{}
	return &this
}

// NewProfileMappingWithDefaults instantiates a new ProfileMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileMappingWithDefaults() *ProfileMapping {
	this := ProfileMapping{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProfileMapping) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProfileMapping) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProfileMapping) SetId(v string) {
	o.Id = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProfileMapping) GetProperties() map[string]ProfileMappingProperty {
	if o == nil || o.Properties == nil {
		var ret map[string]ProfileMappingProperty
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetPropertiesOk() (*map[string]ProfileMappingProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProfileMapping) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]ProfileMappingProperty and assigns it to the Properties field.
func (o *ProfileMapping) SetProperties(v map[string]ProfileMappingProperty) {
	o.Properties = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProfileMapping) GetSource() ProfileMappingSource {
	if o == nil || o.Source == nil {
		var ret ProfileMappingSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetSourceOk() (*ProfileMappingSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProfileMapping) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given ProfileMappingSource and assigns it to the Source field.
func (o *ProfileMapping) SetSource(v ProfileMappingSource) {
	o.Source = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ProfileMapping) GetTarget() ProfileMappingTarget {
	if o == nil || o.Target == nil {
		var ret ProfileMappingTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetTargetOk() (*ProfileMappingTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ProfileMapping) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given ProfileMappingTarget and assigns it to the Target field.
func (o *ProfileMapping) SetTarget(v ProfileMappingTarget) {
	o.Target = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProfileMapping) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProfileMapping) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *ProfileMapping) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o ProfileMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
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

func (o *ProfileMapping) UnmarshalJSON(bytes []byte) (err error) {
	varProfileMapping := _ProfileMapping{}

	err = json.Unmarshal(bytes, &varProfileMapping)
	if err == nil {
		*o = ProfileMapping(varProfileMapping)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "source")
		delete(additionalProperties, "target")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProfileMapping struct {
	value *ProfileMapping
	isSet bool
}

func (v NullableProfileMapping) Get() *ProfileMapping {
	return v.value
}

func (v *NullableProfileMapping) Set(val *ProfileMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileMapping(val *ProfileMapping) *NullableProfileMapping {
	return &NullableProfileMapping{value: val, isSet: true}
}

func (v NullableProfileMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

