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

// ApplicationFeature The Feature object is used to configure application feature settings. 
type ApplicationFeature struct {
	// Description of the feature
	Description *string `json:"description,omitempty"`
	// Identifying name of the feature  | Value | Description   | | --------- | ------------- | | USER_PROVISIONING  | Represents the **To App** provisioning feature setting in the Admin Console | | INBOUND_PROVISIONING | Represents the **To Okta** provisioning feature setting in the Admin Console | 
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	Links *ApplicationFeatureLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationFeature ApplicationFeature

// NewApplicationFeature instantiates a new ApplicationFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFeature() *ApplicationFeature {
	this := ApplicationFeature{}
	return &this
}

// NewApplicationFeatureWithDefaults instantiates a new ApplicationFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFeatureWithDefaults() *ApplicationFeature {
	this := ApplicationFeature{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationFeature) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFeature) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationFeature) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationFeature) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationFeature) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFeature) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationFeature) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationFeature) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplicationFeature) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFeature) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplicationFeature) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplicationFeature) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplicationFeature) GetLinks() ApplicationFeatureLinks {
	if o == nil || o.Links == nil {
		var ret ApplicationFeatureLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFeature) GetLinksOk() (*ApplicationFeatureLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplicationFeature) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ApplicationFeatureLinks and assigns it to the Links field.
func (o *ApplicationFeature) SetLinks(v ApplicationFeatureLinks) {
	o.Links = &v
}

func (o ApplicationFeature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationFeature) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationFeature := _ApplicationFeature{}

	err = json.Unmarshal(bytes, &varApplicationFeature)
	if err == nil {
		*o = ApplicationFeature(varApplicationFeature)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationFeature struct {
	value *ApplicationFeature
	isSet bool
}

func (v NullableApplicationFeature) Get() *ApplicationFeature {
	return v.value
}

func (v *NullableApplicationFeature) Set(val *ApplicationFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFeature(val *ApplicationFeature) *NullableApplicationFeature {
	return &NullableApplicationFeature{value: val, isSet: true}
}

func (v NullableApplicationFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

