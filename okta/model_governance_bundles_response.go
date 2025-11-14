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

// checks if the GovernanceBundlesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GovernanceBundlesResponse{}

// GovernanceBundlesResponse struct for GovernanceBundlesResponse
type GovernanceBundlesResponse struct {
	// List of governance bundles
	Bundles              []GovernanceBundle              `json:"bundles,omitempty"`
	Links                *GovernanceBundlesResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GovernanceBundlesResponse GovernanceBundlesResponse

// NewGovernanceBundlesResponse instantiates a new GovernanceBundlesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGovernanceBundlesResponse() *GovernanceBundlesResponse {
	this := GovernanceBundlesResponse{}
	return &this
}

// NewGovernanceBundlesResponseWithDefaults instantiates a new GovernanceBundlesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGovernanceBundlesResponseWithDefaults() *GovernanceBundlesResponse {
	this := GovernanceBundlesResponse{}
	return &this
}

// GetBundles returns the Bundles field value if set, zero value otherwise.
func (o *GovernanceBundlesResponse) GetBundles() []GovernanceBundle {
	if o == nil || IsNil(o.Bundles) {
		var ret []GovernanceBundle
		return ret
	}
	return o.Bundles
}

// GetBundlesOk returns a tuple with the Bundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceBundlesResponse) GetBundlesOk() ([]GovernanceBundle, bool) {
	if o == nil || IsNil(o.Bundles) {
		return nil, false
	}
	return o.Bundles, true
}

// HasBundles returns a boolean if a field has been set.
func (o *GovernanceBundlesResponse) HasBundles() bool {
	if o != nil && !IsNil(o.Bundles) {
		return true
	}

	return false
}

// SetBundles gets a reference to the given []GovernanceBundle and assigns it to the Bundles field.
func (o *GovernanceBundlesResponse) SetBundles(v []GovernanceBundle) {
	o.Bundles = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GovernanceBundlesResponse) GetLinks() GovernanceBundlesResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret GovernanceBundlesResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceBundlesResponse) GetLinksOk() (*GovernanceBundlesResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GovernanceBundlesResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GovernanceBundlesResponseLinks and assigns it to the Links field.
func (o *GovernanceBundlesResponse) SetLinks(v GovernanceBundlesResponseLinks) {
	o.Links = &v
}

func (o GovernanceBundlesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GovernanceBundlesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bundles) {
		toSerialize["bundles"] = o.Bundles
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GovernanceBundlesResponse) UnmarshalJSON(data []byte) (err error) {
	varGovernanceBundlesResponse := _GovernanceBundlesResponse{}

	err = json.Unmarshal(data, &varGovernanceBundlesResponse)

	if err != nil {
		return err
	}

	*o = GovernanceBundlesResponse(varGovernanceBundlesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bundles")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGovernanceBundlesResponse struct {
	value *GovernanceBundlesResponse
	isSet bool
}

func (v NullableGovernanceBundlesResponse) Get() *GovernanceBundlesResponse {
	return v.value
}

func (v *NullableGovernanceBundlesResponse) Set(val *GovernanceBundlesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGovernanceBundlesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGovernanceBundlesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGovernanceBundlesResponse(val *GovernanceBundlesResponse) *NullableGovernanceBundlesResponse {
	return &NullableGovernanceBundlesResponse{value: val, isSet: true}
}

func (v NullableGovernanceBundlesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGovernanceBundlesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
