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

// Feature Specifies feature release cycle information
type Feature struct {
	// Brief description of the feature and what it provides
	Description *string `json:"description,omitempty"`
	// Unique identifier for this feature
	Id *string `json:"id,omitempty"`
	// Name of the feature
	Name *string `json:"name,omitempty"`
	Stage *FeatureStage `json:"stage,omitempty"`
	// Setting status
	Status *string `json:"status,omitempty"`
	// Type of feature
	Type *string `json:"type,omitempty"`
	Links *FeatureLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Feature Feature

// NewFeature instantiates a new Feature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeature() *Feature {
	this := Feature{}
	return &this
}

// NewFeatureWithDefaults instantiates a new Feature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureWithDefaults() *Feature {
	this := Feature{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Feature) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feature) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Feature) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Feature) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Feature) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feature) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Feature) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Feature) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Feature) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feature) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Feature) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Feature) SetName(v string) {
	o.Name = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *Feature) GetStage() FeatureStage {
	if o == nil || o.Stage == nil {
		var ret FeatureStage
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feature) GetStageOk() (*FeatureStage, bool) {
	if o == nil || o.Stage == nil {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *Feature) HasStage() bool {
	if o != nil && o.Stage != nil {
		return true
	}

	return false
}

// SetStage gets a reference to the given FeatureStage and assigns it to the Stage field.
func (o *Feature) SetStage(v FeatureStage) {
	o.Stage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Feature) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feature) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Feature) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Feature) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Feature) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feature) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Feature) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Feature) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Feature) GetLinks() FeatureLinks {
	if o == nil || o.Links == nil {
		var ret FeatureLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feature) GetLinksOk() (*FeatureLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Feature) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given FeatureLinks and assigns it to the Links field.
func (o *Feature) SetLinks(v FeatureLinks) {
	o.Links = &v
}

func (o Feature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Stage != nil {
		toSerialize["stage"] = o.Stage
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Feature) UnmarshalJSON(bytes []byte) (err error) {
	varFeature := _Feature{}

	err = json.Unmarshal(bytes, &varFeature)
	if err == nil {
		*o = Feature(varFeature)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "stage")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFeature struct {
	value *Feature
	isSet bool
}

func (v NullableFeature) Get() *Feature {
	return v.value
}

func (v *NullableFeature) Set(val *Feature) {
	v.value = val
	v.isSet = true
}

func (v NullableFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeature(val *Feature) *NullableFeature {
	return &NullableFeature{value: val, isSet: true}
}

func (v NullableFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

