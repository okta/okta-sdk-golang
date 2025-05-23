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

// BundleEntitlementLinks struct for BundleEntitlementLinks
type BundleEntitlementLinks struct {
	Values *HrefObject `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BundleEntitlementLinks BundleEntitlementLinks

// NewBundleEntitlementLinks instantiates a new BundleEntitlementLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleEntitlementLinks() *BundleEntitlementLinks {
	this := BundleEntitlementLinks{}
	return &this
}

// NewBundleEntitlementLinksWithDefaults instantiates a new BundleEntitlementLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleEntitlementLinksWithDefaults() *BundleEntitlementLinks {
	this := BundleEntitlementLinks{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *BundleEntitlementLinks) GetValues() HrefObject {
	if o == nil || o.Values == nil {
		var ret HrefObject
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleEntitlementLinks) GetValuesOk() (*HrefObject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *BundleEntitlementLinks) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given HrefObject and assigns it to the Values field.
func (o *BundleEntitlementLinks) SetValues(v HrefObject) {
	o.Values = &v
}

func (o BundleEntitlementLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleEntitlementLinks) UnmarshalJSON(bytes []byte) (err error) {
	varBundleEntitlementLinks := _BundleEntitlementLinks{}

	err = json.Unmarshal(bytes, &varBundleEntitlementLinks)
	if err == nil {
		*o = BundleEntitlementLinks(varBundleEntitlementLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBundleEntitlementLinks struct {
	value *BundleEntitlementLinks
	isSet bool
}

func (v NullableBundleEntitlementLinks) Get() *BundleEntitlementLinks {
	return v.value
}

func (v *NullableBundleEntitlementLinks) Set(val *BundleEntitlementLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleEntitlementLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleEntitlementLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleEntitlementLinks(val *BundleEntitlementLinks) *NullableBundleEntitlementLinks {
	return &NullableBundleEntitlementLinks{value: val, isSet: true}
}

func (v NullableBundleEntitlementLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleEntitlementLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

