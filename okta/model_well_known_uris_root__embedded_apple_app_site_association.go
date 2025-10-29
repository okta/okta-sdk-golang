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

// checks if the WellKnownURIsRootEmbeddedAppleAppSiteAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownURIsRootEmbeddedAppleAppSiteAssociation{}

// WellKnownURIsRootEmbeddedAppleAppSiteAssociation Contains the customized well-known URI content and links
type WellKnownURIsRootEmbeddedAppleAppSiteAssociation struct {
	Customized           *WellKnownURIObjectResponse `json:"customized,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownURIsRootEmbeddedAppleAppSiteAssociation WellKnownURIsRootEmbeddedAppleAppSiteAssociation

// NewWellKnownURIsRootEmbeddedAppleAppSiteAssociation instantiates a new WellKnownURIsRootEmbeddedAppleAppSiteAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownURIsRootEmbeddedAppleAppSiteAssociation() *WellKnownURIsRootEmbeddedAppleAppSiteAssociation {
	this := WellKnownURIsRootEmbeddedAppleAppSiteAssociation{}
	return &this
}

// NewWellKnownURIsRootEmbeddedAppleAppSiteAssociationWithDefaults instantiates a new WellKnownURIsRootEmbeddedAppleAppSiteAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownURIsRootEmbeddedAppleAppSiteAssociationWithDefaults() *WellKnownURIsRootEmbeddedAppleAppSiteAssociation {
	this := WellKnownURIsRootEmbeddedAppleAppSiteAssociation{}
	return &this
}

// GetCustomized returns the Customized field value if set, zero value otherwise.
func (o *WellKnownURIsRootEmbeddedAppleAppSiteAssociation) GetCustomized() WellKnownURIObjectResponse {
	if o == nil || IsNil(o.Customized) {
		var ret WellKnownURIObjectResponse
		return ret
	}
	return *o.Customized
}

// GetCustomizedOk returns a tuple with the Customized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRootEmbeddedAppleAppSiteAssociation) GetCustomizedOk() (*WellKnownURIObjectResponse, bool) {
	if o == nil || IsNil(o.Customized) {
		return nil, false
	}
	return o.Customized, true
}

// HasCustomized returns a boolean if a field has been set.
func (o *WellKnownURIsRootEmbeddedAppleAppSiteAssociation) HasCustomized() bool {
	if o != nil && !IsNil(o.Customized) {
		return true
	}

	return false
}

// SetCustomized gets a reference to the given WellKnownURIObjectResponse and assigns it to the Customized field.
func (o *WellKnownURIsRootEmbeddedAppleAppSiteAssociation) SetCustomized(v WellKnownURIObjectResponse) {
	o.Customized = &v
}

func (o WellKnownURIsRootEmbeddedAppleAppSiteAssociation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownURIsRootEmbeddedAppleAppSiteAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customized) {
		toSerialize["customized"] = o.Customized
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WellKnownURIsRootEmbeddedAppleAppSiteAssociation) UnmarshalJSON(data []byte) (err error) {
	varWellKnownURIsRootEmbeddedAppleAppSiteAssociation := _WellKnownURIsRootEmbeddedAppleAppSiteAssociation{}

	err = json.Unmarshal(data, &varWellKnownURIsRootEmbeddedAppleAppSiteAssociation)

	if err != nil {
		return err
	}

	*o = WellKnownURIsRootEmbeddedAppleAppSiteAssociation(varWellKnownURIsRootEmbeddedAppleAppSiteAssociation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customized")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation struct {
	value *WellKnownURIsRootEmbeddedAppleAppSiteAssociation
	isSet bool
}

func (v NullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation) Get() *WellKnownURIsRootEmbeddedAppleAppSiteAssociation {
	return v.value
}

func (v *NullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation) Set(val *WellKnownURIsRootEmbeddedAppleAppSiteAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation(val *WellKnownURIsRootEmbeddedAppleAppSiteAssociation) *NullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation {
	return &NullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation{value: val, isSet: true}
}

func (v NullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownURIsRootEmbeddedAppleAppSiteAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
