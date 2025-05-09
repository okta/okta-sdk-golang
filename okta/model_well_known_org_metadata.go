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

// WellKnownOrgMetadata struct for WellKnownOrgMetadata
type WellKnownOrgMetadata struct {
	// The unique identifier of the Org
	Id *string `json:"id,omitempty"`
	// The authentication pipeline of the org. `idx` means the org is using the Identity Engine, while `v1` means the org is using the Classic authentication pipeline.
	Pipeline *string `json:"pipeline,omitempty"`
	Settings *WellKnownOrgMetadataSettings `json:"settings,omitempty"`
	Links *WellKnownOrgMetadataLinks `json:"_links,omitempty"`
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
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadata) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WellKnownOrgMetadata) HasId() bool {
	if o != nil && o.Id != nil {
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
	if o == nil || o.Pipeline == nil {
		var ret string
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadata) GetPipelineOk() (*string, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *WellKnownOrgMetadata) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given string and assigns it to the Pipeline field.
func (o *WellKnownOrgMetadata) SetPipeline(v string) {
	o.Pipeline = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *WellKnownOrgMetadata) GetSettings() WellKnownOrgMetadataSettings {
	if o == nil || o.Settings == nil {
		var ret WellKnownOrgMetadataSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadata) GetSettingsOk() (*WellKnownOrgMetadataSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *WellKnownOrgMetadata) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given WellKnownOrgMetadataSettings and assigns it to the Settings field.
func (o *WellKnownOrgMetadata) SetSettings(v WellKnownOrgMetadataSettings) {
	o.Settings = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WellKnownOrgMetadata) GetLinks() WellKnownOrgMetadataLinks {
	if o == nil || o.Links == nil {
		var ret WellKnownOrgMetadataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadata) GetLinksOk() (*WellKnownOrgMetadataLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WellKnownOrgMetadata) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given WellKnownOrgMetadataLinks and assigns it to the Links field.
func (o *WellKnownOrgMetadata) SetLinks(v WellKnownOrgMetadataLinks) {
	o.Links = &v
}

func (o WellKnownOrgMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WellKnownOrgMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varWellKnownOrgMetadata := _WellKnownOrgMetadata{}

	err = json.Unmarshal(bytes, &varWellKnownOrgMetadata)
	if err == nil {
		*o = WellKnownOrgMetadata(varWellKnownOrgMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "pipeline")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

