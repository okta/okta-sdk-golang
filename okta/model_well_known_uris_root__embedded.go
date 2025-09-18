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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the WellKnownURIsRootEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownURIsRootEmbedded{}

// WellKnownURIsRootEmbedded struct for WellKnownURIsRootEmbedded
type WellKnownURIsRootEmbedded struct {
	AppleAppSiteAssociation *WellKnownURIsRootEmbeddedAppleAppSiteAssociation `json:"apple-app-site-association,omitempty"`
	AssetlinksJson          *WellKnownURIsRootEmbeddedAssetlinksJson          `json:"assetlinks.json,omitempty"`
	Webauthn                *WellKnownURIsRootEmbeddedAppleAppSiteAssociation `json:"webauthn,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _WellKnownURIsRootEmbedded WellKnownURIsRootEmbedded

// NewWellKnownURIsRootEmbedded instantiates a new WellKnownURIsRootEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownURIsRootEmbedded() *WellKnownURIsRootEmbedded {
	this := WellKnownURIsRootEmbedded{}
	return &this
}

// NewWellKnownURIsRootEmbeddedWithDefaults instantiates a new WellKnownURIsRootEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownURIsRootEmbeddedWithDefaults() *WellKnownURIsRootEmbedded {
	this := WellKnownURIsRootEmbedded{}
	return &this
}

// GetAppleAppSiteAssociation returns the AppleAppSiteAssociation field value if set, zero value otherwise.
func (o *WellKnownURIsRootEmbedded) GetAppleAppSiteAssociation() WellKnownURIsRootEmbeddedAppleAppSiteAssociation {
	if o == nil || IsNil(o.AppleAppSiteAssociation) {
		var ret WellKnownURIsRootEmbeddedAppleAppSiteAssociation
		return ret
	}
	return *o.AppleAppSiteAssociation
}

// GetAppleAppSiteAssociationOk returns a tuple with the AppleAppSiteAssociation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRootEmbedded) GetAppleAppSiteAssociationOk() (*WellKnownURIsRootEmbeddedAppleAppSiteAssociation, bool) {
	if o == nil || IsNil(o.AppleAppSiteAssociation) {
		return nil, false
	}
	return o.AppleAppSiteAssociation, true
}

// HasAppleAppSiteAssociation returns a boolean if a field has been set.
func (o *WellKnownURIsRootEmbedded) HasAppleAppSiteAssociation() bool {
	if o != nil && !IsNil(o.AppleAppSiteAssociation) {
		return true
	}

	return false
}

// SetAppleAppSiteAssociation gets a reference to the given WellKnownURIsRootEmbeddedAppleAppSiteAssociation and assigns it to the AppleAppSiteAssociation field.
func (o *WellKnownURIsRootEmbedded) SetAppleAppSiteAssociation(v WellKnownURIsRootEmbeddedAppleAppSiteAssociation) {
	o.AppleAppSiteAssociation = &v
}

// GetAssetlinksJson returns the AssetlinksJson field value if set, zero value otherwise.
func (o *WellKnownURIsRootEmbedded) GetAssetlinksJson() WellKnownURIsRootEmbeddedAssetlinksJson {
	if o == nil || IsNil(o.AssetlinksJson) {
		var ret WellKnownURIsRootEmbeddedAssetlinksJson
		return ret
	}
	return *o.AssetlinksJson
}

// GetAssetlinksJsonOk returns a tuple with the AssetlinksJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRootEmbedded) GetAssetlinksJsonOk() (*WellKnownURIsRootEmbeddedAssetlinksJson, bool) {
	if o == nil || IsNil(o.AssetlinksJson) {
		return nil, false
	}
	return o.AssetlinksJson, true
}

// HasAssetlinksJson returns a boolean if a field has been set.
func (o *WellKnownURIsRootEmbedded) HasAssetlinksJson() bool {
	if o != nil && !IsNil(o.AssetlinksJson) {
		return true
	}

	return false
}

// SetAssetlinksJson gets a reference to the given WellKnownURIsRootEmbeddedAssetlinksJson and assigns it to the AssetlinksJson field.
func (o *WellKnownURIsRootEmbedded) SetAssetlinksJson(v WellKnownURIsRootEmbeddedAssetlinksJson) {
	o.AssetlinksJson = &v
}

// GetWebauthn returns the Webauthn field value if set, zero value otherwise.
func (o *WellKnownURIsRootEmbedded) GetWebauthn() WellKnownURIsRootEmbeddedAppleAppSiteAssociation {
	if o == nil || IsNil(o.Webauthn) {
		var ret WellKnownURIsRootEmbeddedAppleAppSiteAssociation
		return ret
	}
	return *o.Webauthn
}

// GetWebauthnOk returns a tuple with the Webauthn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRootEmbedded) GetWebauthnOk() (*WellKnownURIsRootEmbeddedAppleAppSiteAssociation, bool) {
	if o == nil || IsNil(o.Webauthn) {
		return nil, false
	}
	return o.Webauthn, true
}

// HasWebauthn returns a boolean if a field has been set.
func (o *WellKnownURIsRootEmbedded) HasWebauthn() bool {
	if o != nil && !IsNil(o.Webauthn) {
		return true
	}

	return false
}

// SetWebauthn gets a reference to the given WellKnownURIsRootEmbeddedAppleAppSiteAssociation and assigns it to the Webauthn field.
func (o *WellKnownURIsRootEmbedded) SetWebauthn(v WellKnownURIsRootEmbeddedAppleAppSiteAssociation) {
	o.Webauthn = &v
}

func (o WellKnownURIsRootEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownURIsRootEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppleAppSiteAssociation) {
		toSerialize["apple-app-site-association"] = o.AppleAppSiteAssociation
	}
	if !IsNil(o.AssetlinksJson) {
		toSerialize["assetlinks.json"] = o.AssetlinksJson
	}
	if !IsNil(o.Webauthn) {
		toSerialize["webauthn"] = o.Webauthn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WellKnownURIsRootEmbedded) UnmarshalJSON(data []byte) (err error) {
	varWellKnownURIsRootEmbedded := _WellKnownURIsRootEmbedded{}

	err = json.Unmarshal(data, &varWellKnownURIsRootEmbedded)

	if err != nil {
		return err
	}

	*o = WellKnownURIsRootEmbedded(varWellKnownURIsRootEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apple-app-site-association")
		delete(additionalProperties, "assetlinks.json")
		delete(additionalProperties, "webauthn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWellKnownURIsRootEmbedded struct {
	value *WellKnownURIsRootEmbedded
	isSet bool
}

func (v NullableWellKnownURIsRootEmbedded) Get() *WellKnownURIsRootEmbedded {
	return v.value
}

func (v *NullableWellKnownURIsRootEmbedded) Set(val *WellKnownURIsRootEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownURIsRootEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownURIsRootEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownURIsRootEmbedded(val *WellKnownURIsRootEmbedded) *NullableWellKnownURIsRootEmbedded {
	return &NullableWellKnownURIsRootEmbedded{value: val, isSet: true}
}

func (v NullableWellKnownURIsRootEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownURIsRootEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
