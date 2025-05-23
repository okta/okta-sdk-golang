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

// WellKnownOrgMetadataLinks struct for WellKnownOrgMetadataLinks
type WellKnownOrgMetadataLinks struct {
	Alternate *HrefObject `json:"alternate,omitempty"`
	Organization *HrefObject `json:"organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownOrgMetadataLinks WellKnownOrgMetadataLinks

// NewWellKnownOrgMetadataLinks instantiates a new WellKnownOrgMetadataLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownOrgMetadataLinks() *WellKnownOrgMetadataLinks {
	this := WellKnownOrgMetadataLinks{}
	return &this
}

// NewWellKnownOrgMetadataLinksWithDefaults instantiates a new WellKnownOrgMetadataLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownOrgMetadataLinksWithDefaults() *WellKnownOrgMetadataLinks {
	this := WellKnownOrgMetadataLinks{}
	return &this
}

// GetAlternate returns the Alternate field value if set, zero value otherwise.
func (o *WellKnownOrgMetadataLinks) GetAlternate() HrefObject {
	if o == nil || o.Alternate == nil {
		var ret HrefObject
		return ret
	}
	return *o.Alternate
}

// GetAlternateOk returns a tuple with the Alternate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadataLinks) GetAlternateOk() (*HrefObject, bool) {
	if o == nil || o.Alternate == nil {
		return nil, false
	}
	return o.Alternate, true
}

// HasAlternate returns a boolean if a field has been set.
func (o *WellKnownOrgMetadataLinks) HasAlternate() bool {
	if o != nil && o.Alternate != nil {
		return true
	}

	return false
}

// SetAlternate gets a reference to the given HrefObject and assigns it to the Alternate field.
func (o *WellKnownOrgMetadataLinks) SetAlternate(v HrefObject) {
	o.Alternate = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *WellKnownOrgMetadataLinks) GetOrganization() HrefObject {
	if o == nil || o.Organization == nil {
		var ret HrefObject
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadataLinks) GetOrganizationOk() (*HrefObject, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *WellKnownOrgMetadataLinks) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given HrefObject and assigns it to the Organization field.
func (o *WellKnownOrgMetadataLinks) SetOrganization(v HrefObject) {
	o.Organization = &v
}

func (o WellKnownOrgMetadataLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alternate != nil {
		toSerialize["alternate"] = o.Alternate
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WellKnownOrgMetadataLinks) UnmarshalJSON(bytes []byte) (err error) {
	varWellKnownOrgMetadataLinks := _WellKnownOrgMetadataLinks{}

	err = json.Unmarshal(bytes, &varWellKnownOrgMetadataLinks)
	if err == nil {
		*o = WellKnownOrgMetadataLinks(varWellKnownOrgMetadataLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "alternate")
		delete(additionalProperties, "organization")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWellKnownOrgMetadataLinks struct {
	value *WellKnownOrgMetadataLinks
	isSet bool
}

func (v NullableWellKnownOrgMetadataLinks) Get() *WellKnownOrgMetadataLinks {
	return v.value
}

func (v *NullableWellKnownOrgMetadataLinks) Set(val *WellKnownOrgMetadataLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownOrgMetadataLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownOrgMetadataLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownOrgMetadataLinks(val *WellKnownOrgMetadataLinks) *NullableWellKnownOrgMetadataLinks {
	return &NullableWellKnownOrgMetadataLinks{value: val, isSet: true}
}

func (v NullableWellKnownOrgMetadataLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownOrgMetadataLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

