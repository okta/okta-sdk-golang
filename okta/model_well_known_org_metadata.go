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

// checks if the WellKnownOrgMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownOrgMetadata{}

// WellKnownOrgMetadata struct for WellKnownOrgMetadata
type WellKnownOrgMetadata struct {
	// Org unique identifier
	Id *string `json:"id,omitempty"`
	// The Okta authentication pipeline of the org
	Pipeline             *string                    `json:"pipeline,omitempty"`
	Links                *WellKnownOrgMetadataLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownOrgMetadata WellKnownOrgMetadata

// NewWellKnownOrgMetadata instantiates a new WellKnownOrgMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownOrgMetadata() *WellKnownOrgMetadata {
	this := WellKnownOrgMetadata{}
	return &this
}

// NewWellKnownOrgMetadataWithDefaults instantiates a new WellKnownOrgMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownOrgMetadataWithDefaults() *WellKnownOrgMetadata {
	this := WellKnownOrgMetadata{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WellKnownOrgMetadata) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadata) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WellKnownOrgMetadata) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WellKnownOrgMetadata) SetId(v string) {
	o.Id = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *WellKnownOrgMetadata) GetPipeline() string {
	if o == nil || IsNil(o.Pipeline) {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadata) GetPipelineOk() (*string, bool) {
	if o == nil || IsNil(o.Pipeline) {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *WellKnownOrgMetadata) HasPipeline() bool {
	if o != nil && !IsNil(o.Pipeline) {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *WellKnownOrgMetadata) SetPipeline(v string) {
	o.Pipeline = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WellKnownOrgMetadata) GetLinks() WellKnownOrgMetadataLinks {
	if o == nil || IsNil(o.Links) {
		var ret WellKnownOrgMetadataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadata) GetLinksOk() (*WellKnownOrgMetadataLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WellKnownOrgMetadata) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WellKnownOrgMetadataLinks and assigns it to the Links field.
func (o *WellKnownOrgMetadata) SetLinks(v WellKnownOrgMetadataLinks) {
	o.Links = &v
}

func (o WellKnownOrgMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownOrgMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Pipeline) {
		toSerialize["pipeline"] = o.Pipeline
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WellKnownOrgMetadata) UnmarshalJSON(data []byte) (err error) {
	varWellKnownOrgMetadata := _WellKnownOrgMetadata{}

	err = json.Unmarshal(data, &varWellKnownOrgMetadata)

	if err != nil {
		return err
	}

	*o = WellKnownOrgMetadata(varWellKnownOrgMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "pipeline")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWellKnownOrgMetadata struct {
	value *WellKnownOrgMetadata
	isSet bool
}

func (v NullableWellKnownOrgMetadata) Get() *WellKnownOrgMetadata {
	return v.value
}

func (v *NullableWellKnownOrgMetadata) Set(val *WellKnownOrgMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownOrgMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownOrgMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownOrgMetadata(val *WellKnownOrgMetadata) *NullableWellKnownOrgMetadata {
	return &NullableWellKnownOrgMetadata{value: val, isSet: true}
}

func (v NullableWellKnownOrgMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownOrgMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
