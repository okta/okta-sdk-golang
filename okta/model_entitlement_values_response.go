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

// EntitlementValuesResponse struct for EntitlementValuesResponse
type EntitlementValuesResponse struct {
	EntitlementValues []EntitlementValue `json:"entitlementValues,omitempty"`
	Links NullableBundleEntitlementsResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValuesResponse EntitlementValuesResponse

// NewEntitlementValuesResponse instantiates a new EntitlementValuesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValuesResponse() *EntitlementValuesResponse {
	this := EntitlementValuesResponse{}
	return &this
}

// NewEntitlementValuesResponseWithDefaults instantiates a new EntitlementValuesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValuesResponseWithDefaults() *EntitlementValuesResponse {
	this := EntitlementValuesResponse{}
	return &this
}

// GetEntitlementValues returns the EntitlementValues field value if set, zero value otherwise.
func (o *EntitlementValuesResponse) GetEntitlementValues() []EntitlementValue {
	if o == nil || o.EntitlementValues == nil {
		var ret []EntitlementValue
		return ret
	}
	return o.EntitlementValues
}

// GetEntitlementValuesOk returns a tuple with the EntitlementValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValuesResponse) GetEntitlementValuesOk() ([]EntitlementValue, bool) {
	if o == nil || o.EntitlementValues == nil {
		return nil, false
	}
	return o.EntitlementValues, true
}

// HasEntitlementValues returns a boolean if a field has been set.
func (o *EntitlementValuesResponse) HasEntitlementValues() bool {
	if o != nil && o.EntitlementValues != nil {
		return true
	}

	return false
}

// SetEntitlementValues gets a reference to the given []EntitlementValue and assigns it to the EntitlementValues field.
func (o *EntitlementValuesResponse) SetEntitlementValues(v []EntitlementValue) {
	o.EntitlementValues = v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitlementValuesResponse) GetLinks() BundleEntitlementsResponseLinks {
	if o == nil || o.Links.Get() == nil {
		var ret BundleEntitlementsResponseLinks
		return ret
	}
	return *o.Links.Get()
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitlementValuesResponse) GetLinksOk() (*BundleEntitlementsResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links.Get(), o.Links.IsSet()
}

// HasLinks returns a boolean if a field has been set.
func (o *EntitlementValuesResponse) HasLinks() bool {
	if o != nil && o.Links.IsSet() {
		return true
	}

	return false
}

// SetLinks gets a reference to the given NullableBundleEntitlementsResponseLinks and assigns it to the Links field.
func (o *EntitlementValuesResponse) SetLinks(v BundleEntitlementsResponseLinks) {
	o.Links.Set(&v)
}
// SetLinksNil sets the value for Links to be an explicit nil
func (o *EntitlementValuesResponse) SetLinksNil() {
	o.Links.Set(nil)
}

// UnsetLinks ensures that no value is present for Links, not even an explicit nil
func (o *EntitlementValuesResponse) UnsetLinks() {
	o.Links.Unset()
}

func (o EntitlementValuesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntitlementValues != nil {
		toSerialize["entitlementValues"] = o.EntitlementValues
	}
	if o.Links.IsSet() {
		toSerialize["_links"] = o.Links.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementValuesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementValuesResponse := _EntitlementValuesResponse{}

	err = json.Unmarshal(bytes, &varEntitlementValuesResponse)
	if err == nil {
		*o = EntitlementValuesResponse(varEntitlementValuesResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "entitlementValues")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementValuesResponse struct {
	value *EntitlementValuesResponse
	isSet bool
}

func (v NullableEntitlementValuesResponse) Get() *EntitlementValuesResponse {
	return v.value
}

func (v *NullableEntitlementValuesResponse) Set(val *EntitlementValuesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValuesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValuesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValuesResponse(val *EntitlementValuesResponse) *NullableEntitlementValuesResponse {
	return &NullableEntitlementValuesResponse{value: val, isSet: true}
}

func (v NullableEntitlementValuesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValuesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

