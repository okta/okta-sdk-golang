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

// checks if the BundleEntitlementsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleEntitlementsResponse{}

// BundleEntitlementsResponse struct for BundleEntitlementsResponse
type BundleEntitlementsResponse struct {
	Entitlements         []BundleEntitlement              `json:"entitlements,omitempty"`
	Links                *BundleEntitlementsResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BundleEntitlementsResponse BundleEntitlementsResponse

// NewBundleEntitlementsResponse instantiates a new BundleEntitlementsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleEntitlementsResponse() *BundleEntitlementsResponse {
	this := BundleEntitlementsResponse{}
	return &this
}

// NewBundleEntitlementsResponseWithDefaults instantiates a new BundleEntitlementsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleEntitlementsResponseWithDefaults() *BundleEntitlementsResponse {
	this := BundleEntitlementsResponse{}
	return &this
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *BundleEntitlementsResponse) GetEntitlements() []BundleEntitlement {
	if o == nil || IsNil(o.Entitlements) {
		var ret []BundleEntitlement
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleEntitlementsResponse) GetEntitlementsOk() ([]BundleEntitlement, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *BundleEntitlementsResponse) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []BundleEntitlement and assigns it to the Entitlements field.
func (o *BundleEntitlementsResponse) SetEntitlements(v []BundleEntitlement) {
	o.Entitlements = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BundleEntitlementsResponse) GetLinks() BundleEntitlementsResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret BundleEntitlementsResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleEntitlementsResponse) GetLinksOk() (*BundleEntitlementsResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BundleEntitlementsResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BundleEntitlementsResponseLinks and assigns it to the Links field.
func (o *BundleEntitlementsResponse) SetLinks(v BundleEntitlementsResponseLinks) {
	o.Links = &v
}

func (o BundleEntitlementsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleEntitlementsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BundleEntitlementsResponse) UnmarshalJSON(data []byte) (err error) {
	varBundleEntitlementsResponse := _BundleEntitlementsResponse{}

	err = json.Unmarshal(data, &varBundleEntitlementsResponse)

	if err != nil {
		return err
	}

	*o = BundleEntitlementsResponse(varBundleEntitlementsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBundleEntitlementsResponse struct {
	value *BundleEntitlementsResponse
	isSet bool
}

func (v NullableBundleEntitlementsResponse) Get() *BundleEntitlementsResponse {
	return v.value
}

func (v *NullableBundleEntitlementsResponse) Set(val *BundleEntitlementsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleEntitlementsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleEntitlementsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleEntitlementsResponse(val *BundleEntitlementsResponse) *NullableBundleEntitlementsResponse {
	return &NullableBundleEntitlementsResponse{value: val, isSet: true}
}

func (v NullableBundleEntitlementsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleEntitlementsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
