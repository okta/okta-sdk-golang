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
)

// checks if the ProfileMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileMapping{}

// ProfileMapping The profile mapping object describes a mapping between an Okta user's and an app user's properties using [JSON Schema Draft 4](https://datatracker.ietf.org/doc/html/draft-zyp-json-schema-04).  > **Note:** Same type source/target mappings aren't supported by this API. Profile mappings must be between Okta and an app.
type ProfileMapping struct {
	// Unique identifier for a profile mapping
	Id                   *string                `json:"id,omitempty"`
	Properties           ProfileMappingProperty `json:"properties,omitempty"`
	Source               *ProfileMappingSource  `json:"source,omitempty"`
	Target               *ProfileMappingTarget  `json:"target,omitempty"`
	Links                *LinksSelf             `json:"_links,omitempty"`
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProfileMapping) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProfileMapping) SetId(v string) {
	o.Id = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProfileMapping) GetProperties() ProfileMappingProperty {
	if o == nil || IsNil(o.Properties) {
		var ret ProfileMappingProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetPropertiesOk() (ProfileMappingProperty, bool) {
	if o == nil || IsNil(o.Properties) {
		return ProfileMappingProperty{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProfileMapping) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given ProfileMappingProperty and assigns it to the Properties field.
func (o *ProfileMapping) SetProperties(v ProfileMappingProperty) {
	o.Properties = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ProfileMapping) GetSource() ProfileMappingSource {
	if o == nil || IsNil(o.Source) {
		var ret ProfileMappingSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetSourceOk() (*ProfileMappingSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ProfileMapping) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
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
	if o == nil || IsNil(o.Target) {
		var ret ProfileMappingTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetTargetOk() (*ProfileMappingTarget, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ProfileMapping) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
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
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMapping) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProfileMapping) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *ProfileMapping) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o ProfileMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProfileMapping) UnmarshalJSON(data []byte) (err error) {
	varProfileMapping := _ProfileMapping{}

	err = json.Unmarshal(data, &varProfileMapping)

	if err != nil {
		return err
	}

	*o = ProfileMapping(varProfileMapping)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "source")
		delete(additionalProperties, "target")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
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
