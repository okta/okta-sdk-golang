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

// BundleEntitlementsResponse struct for BundleEntitlementsResponse
type BundleEntitlementsResponse struct {
	Entitlements []BundleEntitlement `json:"entitlements,omitempty"`
	Links NullableBundleEntitlementsResponseLinks `json:"_links,omitempty"`
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
	if o == nil || o.Entitlements == nil {
		var ret []BundleEntitlement
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleEntitlementsResponse) GetEntitlementsOk() ([]BundleEntitlement, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *BundleEntitlementsResponse) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []BundleEntitlement and assigns it to the Entitlements field.
func (o *BundleEntitlementsResponse) SetEntitlements(v []BundleEntitlement) {
	o.Entitlements = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BundleEntitlementsResponse) GetLinks() BundleEntitlementsResponseLinks {
	if o == nil || o.Links.Get() == nil {
		var ret BundleEntitlementsResponseLinks
		return ret
	}
	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BundleEntitlementsResponse) GetLinksOk() (*BundleEntitlementsResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// HasLinks returns a boolean if a field has been set.
func (o *BundleEntitlementsResponse) HasLinks() bool {
	if o != nil && o.Links.IsSet() {
		return true
	}

	return false
}

// SetLinks gets a reference to the given NullableBundleEntitlementsResponseLinks and assigns it to the Links field.
func (o *BundleEntitlementsResponse) SetLinks(v BundleEntitlementsResponseLinks) {
	o.Links.Set(&v)
}
// SetLinksNil sets the value for Links to be an explicit nil
func (o *BundleEntitlementsResponse) SetLinksNil() {
	o.Links.Set(nil)
}

// UnsetLinks ensures that no value is present for Links, not even an explicit nil
func (o *BundleEntitlementsResponse) UnsetLinks() {
	o.Links.Unset()
}

func (o BundleEntitlementsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
	if o.Links.IsSet() {
		toSerialize["_links"] = o.Links.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BundleEntitlementsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBundleEntitlementsResponse := _BundleEntitlementsResponse{}

	err = json.Unmarshal(bytes, &varBundleEntitlementsResponse)
	if err == nil {
		*o = BundleEntitlementsResponse(varBundleEntitlementsResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

