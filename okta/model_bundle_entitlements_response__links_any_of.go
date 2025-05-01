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

// BundleEntitlementsResponseLinksAnyOf struct for BundleEntitlementsResponseLinksAnyOf
type BundleEntitlementsResponseLinksAnyOf struct {
	Bundle *HrefObject `json:"bundle,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BundleEntitlementsResponseLinksAnyOf BundleEntitlementsResponseLinksAnyOf

// NewBundleEntitlementsResponseLinksAnyOf instantiates a new BundleEntitlementsResponseLinksAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleEntitlementsResponseLinksAnyOf() *BundleEntitlementsResponseLinksAnyOf {
	this := BundleEntitlementsResponseLinksAnyOf{}
	return &this
}

// NewBundleEntitlementsResponseLinksAnyOfWithDefaults instantiates a new BundleEntitlementsResponseLinksAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleEntitlementsResponseLinksAnyOfWithDefaults() *BundleEntitlementsResponseLinksAnyOf {
	this := BundleEntitlementsResponseLinksAnyOf{}
	return &this
}

// GetBundle returns the Bundle field value if set, zero value otherwise.
func (o *BundleEntitlementsResponseLinksAnyOf) GetBundle() HrefObject {
	if o == nil || o.Bundle == nil {
		var ret HrefObject
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleEntitlementsResponseLinksAnyOf) GetBundleOk() (*HrefObject, bool) {
	if o == nil || o.Bundle == nil {
		return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *BundleEntitlementsResponseLinksAnyOf) HasBundle() bool {
	if o != nil && o.Bundle != nil {
		return true
	}

	return false
}

// SetBundle gets a reference to the given HrefObject and assigns it to the Bundle field.
func (o *BundleEntitlementsResponseLinksAnyOf) SetBundle(v HrefObject) {
	o.Bundle = &v
}

func (o BundleEntitlementsResponseLinksAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bundle != nil {
		toSerialize["bundle"] = o.Bundle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleEntitlementsResponseLinksAnyOf) UnmarshalJSON(bytes []byte) (err error) {
	varBundleEntitlementsResponseLinksAnyOf := _BundleEntitlementsResponseLinksAnyOf{}

	err = json.Unmarshal(bytes, &varBundleEntitlementsResponseLinksAnyOf)
	if err == nil {
		*o = BundleEntitlementsResponseLinksAnyOf(varBundleEntitlementsResponseLinksAnyOf)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "bundle")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBundleEntitlementsResponseLinksAnyOf struct {
	value *BundleEntitlementsResponseLinksAnyOf
	isSet bool
}

func (v NullableBundleEntitlementsResponseLinksAnyOf) Get() *BundleEntitlementsResponseLinksAnyOf {
	return v.value
}

func (v *NullableBundleEntitlementsResponseLinksAnyOf) Set(val *BundleEntitlementsResponseLinksAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleEntitlementsResponseLinksAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleEntitlementsResponseLinksAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleEntitlementsResponseLinksAnyOf(val *BundleEntitlementsResponseLinksAnyOf) *NullableBundleEntitlementsResponseLinksAnyOf {
	return &NullableBundleEntitlementsResponseLinksAnyOf{value: val, isSet: true}
}

func (v NullableBundleEntitlementsResponseLinksAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleEntitlementsResponseLinksAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

