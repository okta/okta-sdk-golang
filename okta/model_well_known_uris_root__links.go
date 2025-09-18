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

// checks if the WellKnownURIsRootLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WellKnownURIsRootLinks{}

// WellKnownURIsRootLinks struct for WellKnownURIsRootLinks
type WellKnownURIsRootLinks struct {
	Self                    *HrefObjectSelfLink `json:"self,omitempty"`
	AppleAppSiteAssociation *HrefObject         `json:"apple-app-site-association,omitempty"`
	AssetlinksJson          *HrefObject         `json:"assetlinks.json,omitempty"`
	Webauthn                *HrefObject         `json:"webauthn,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _WellKnownURIsRootLinks WellKnownURIsRootLinks

// NewWellKnownURIsRootLinks instantiates a new WellKnownURIsRootLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownURIsRootLinks() *WellKnownURIsRootLinks {
	this := WellKnownURIsRootLinks{}
	return &this
}

// NewWellKnownURIsRootLinksWithDefaults instantiates a new WellKnownURIsRootLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownURIsRootLinksWithDefaults() *WellKnownURIsRootLinks {
	this := WellKnownURIsRootLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *WellKnownURIsRootLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRootLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *WellKnownURIsRootLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *WellKnownURIsRootLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetAppleAppSiteAssociation returns the AppleAppSiteAssociation field value if set, zero value otherwise.
func (o *WellKnownURIsRootLinks) GetAppleAppSiteAssociation() HrefObject {
	if o == nil || IsNil(o.AppleAppSiteAssociation) {
		var ret HrefObject
		return ret
	}
	return *o.AppleAppSiteAssociation
}

// GetAppleAppSiteAssociationOk returns a tuple with the AppleAppSiteAssociation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRootLinks) GetAppleAppSiteAssociationOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.AppleAppSiteAssociation) {
		return nil, false
	}
	return o.AppleAppSiteAssociation, true
}

// HasAppleAppSiteAssociation returns a boolean if a field has been set.
func (o *WellKnownURIsRootLinks) HasAppleAppSiteAssociation() bool {
	if o != nil && !IsNil(o.AppleAppSiteAssociation) {
		return true
	}

	return false
}

// SetAppleAppSiteAssociation gets a reference to the given HrefObject and assigns it to the AppleAppSiteAssociation field.
func (o *WellKnownURIsRootLinks) SetAppleAppSiteAssociation(v HrefObject) {
	o.AppleAppSiteAssociation = &v
}

// GetAssetlinksJson returns the AssetlinksJson field value if set, zero value otherwise.
func (o *WellKnownURIsRootLinks) GetAssetlinksJson() HrefObject {
	if o == nil || IsNil(o.AssetlinksJson) {
		var ret HrefObject
		return ret
	}
	return *o.AssetlinksJson
}

// GetAssetlinksJsonOk returns a tuple with the AssetlinksJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRootLinks) GetAssetlinksJsonOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.AssetlinksJson) {
		return nil, false
	}
	return o.AssetlinksJson, true
}

// HasAssetlinksJson returns a boolean if a field has been set.
func (o *WellKnownURIsRootLinks) HasAssetlinksJson() bool {
	if o != nil && !IsNil(o.AssetlinksJson) {
		return true
	}

	return false
}

// SetAssetlinksJson gets a reference to the given HrefObject and assigns it to the AssetlinksJson field.
func (o *WellKnownURIsRootLinks) SetAssetlinksJson(v HrefObject) {
	o.AssetlinksJson = &v
}

// GetWebauthn returns the Webauthn field value if set, zero value otherwise.
func (o *WellKnownURIsRootLinks) GetWebauthn() HrefObject {
	if o == nil || IsNil(o.Webauthn) {
		var ret HrefObject
		return ret
	}
	return *o.Webauthn
}

// GetWebauthnOk returns a tuple with the Webauthn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownURIsRootLinks) GetWebauthnOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Webauthn) {
		return nil, false
	}
	return o.Webauthn, true
}

// HasWebauthn returns a boolean if a field has been set.
func (o *WellKnownURIsRootLinks) HasWebauthn() bool {
	if o != nil && !IsNil(o.Webauthn) {
		return true
	}

	return false
}

// SetWebauthn gets a reference to the given HrefObject and assigns it to the Webauthn field.
func (o *WellKnownURIsRootLinks) SetWebauthn(v HrefObject) {
	o.Webauthn = &v
}

func (o WellKnownURIsRootLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WellKnownURIsRootLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
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

func (o *WellKnownURIsRootLinks) UnmarshalJSON(data []byte) (err error) {
	varWellKnownURIsRootLinks := _WellKnownURIsRootLinks{}

	err = json.Unmarshal(data, &varWellKnownURIsRootLinks)

	if err != nil {
		return err
	}

	*o = WellKnownURIsRootLinks(varWellKnownURIsRootLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "apple-app-site-association")
		delete(additionalProperties, "assetlinks.json")
		delete(additionalProperties, "webauthn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWellKnownURIsRootLinks struct {
	value *WellKnownURIsRootLinks
	isSet bool
}

func (v NullableWellKnownURIsRootLinks) Get() *WellKnownURIsRootLinks {
	return v.value
}

func (v *NullableWellKnownURIsRootLinks) Set(val *WellKnownURIsRootLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownURIsRootLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownURIsRootLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownURIsRootLinks(val *WellKnownURIsRootLinks) *NullableWellKnownURIsRootLinks {
	return &NullableWellKnownURIsRootLinks{value: val, isSet: true}
}

func (v NullableWellKnownURIsRootLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownURIsRootLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
