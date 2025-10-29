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

// checks if the WellKnownOrgMetadataLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownOrgMetadataLinks{}

// WellKnownOrgMetadataLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for this object using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type WellKnownOrgMetadataLinks struct {
	// Link to the custom domain org URL
	Alternate *HrefObject `json:"alternate,omitempty"`
	// Link to the org URL
	Organization         *HrefObject `json:"organization,omitempty"`
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
	if o == nil || IsNil(o.Alternate) {
		var ret HrefObject
		return ret
	}
	return *o.Alternate
}

// GetAlternateOk returns a tuple with the Alternate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadataLinks) GetAlternateOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Alternate) {
		return nil, false
	}
	return o.Alternate, true
}

// HasAlternate returns a boolean if a field has been set.
func (o *WellKnownOrgMetadataLinks) HasAlternate() bool {
	if o != nil && !IsNil(o.Alternate) {
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
	if o == nil || IsNil(o.Organization) {
		var ret HrefObject
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownOrgMetadataLinks) GetOrganizationOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *WellKnownOrgMetadataLinks) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given HrefObject and assigns it to the Organization field.
func (o *WellKnownOrgMetadataLinks) SetOrganization(v HrefObject) {
	o.Organization = &v
}

func (o WellKnownOrgMetadataLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownOrgMetadataLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alternate) {
		toSerialize["alternate"] = o.Alternate
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WellKnownOrgMetadataLinks) UnmarshalJSON(data []byte) (err error) {
	varWellKnownOrgMetadataLinks := _WellKnownOrgMetadataLinks{}

	err = json.Unmarshal(data, &varWellKnownOrgMetadataLinks)

	if err != nil {
		return err
	}

	*o = WellKnownOrgMetadataLinks(varWellKnownOrgMetadataLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alternate")
		delete(additionalProperties, "organization")
		o.AdditionalProperties = additionalProperties
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
