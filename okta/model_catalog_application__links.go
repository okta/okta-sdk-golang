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

// checks if the CatalogApplicationLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogApplicationLinks{}

// CatalogApplicationLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type CatalogApplicationLinks struct {
	// List of app logo resources
	Logo                 []HrefObjectLogoLink `json:"logo,omitempty"`
	Self                 *HrefObjectSelfLink  `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogApplicationLinks CatalogApplicationLinks

// NewCatalogApplicationLinks instantiates a new CatalogApplicationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogApplicationLinks() *CatalogApplicationLinks {
	this := CatalogApplicationLinks{}
	return &this
}

// NewCatalogApplicationLinksWithDefaults instantiates a new CatalogApplicationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogApplicationLinksWithDefaults() *CatalogApplicationLinks {
	this := CatalogApplicationLinks{}
	return &this
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CatalogApplicationLinks) GetLogo() []HrefObjectLogoLink {
	if o == nil || IsNil(o.Logo) {
		var ret []HrefObjectLogoLink
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplicationLinks) GetLogoOk() ([]HrefObjectLogoLink, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CatalogApplicationLinks) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []HrefObjectLogoLink and assigns it to the Logo field.
func (o *CatalogApplicationLinks) SetLogo(v []HrefObjectLogoLink) {
	o.Logo = v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *CatalogApplicationLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplicationLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *CatalogApplicationLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *CatalogApplicationLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

func (o CatalogApplicationLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogApplicationLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatalogApplicationLinks) UnmarshalJSON(data []byte) (err error) {
	varCatalogApplicationLinks := _CatalogApplicationLinks{}

	err = json.Unmarshal(data, &varCatalogApplicationLinks)

	if err != nil {
		return err
	}

	*o = CatalogApplicationLinks(varCatalogApplicationLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "logo")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCatalogApplicationLinks struct {
	value *CatalogApplicationLinks
	isSet bool
}

func (v NullableCatalogApplicationLinks) Get() *CatalogApplicationLinks {
	return v.value
}

func (v *NullableCatalogApplicationLinks) Set(val *CatalogApplicationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogApplicationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogApplicationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogApplicationLinks(val *CatalogApplicationLinks) *NullableCatalogApplicationLinks {
	return &NullableCatalogApplicationLinks{value: val, isSet: true}
}

func (v NullableCatalogApplicationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogApplicationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
