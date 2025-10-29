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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the UserFactorActivateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorActivateResponse{}

// UserFactorActivateResponse struct for UserFactorActivateResponse
type UserFactorActivateResponse struct {
	// Type of the factor
	FactorType           *string                          `json:"factorType,omitempty"`
	Links                *UserFactorActivateResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorActivateResponse UserFactorActivateResponse

// NewUserFactorActivateResponse instantiates a new UserFactorActivateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorActivateResponse() *UserFactorActivateResponse {
	this := UserFactorActivateResponse{}
	return &this
}

// NewUserFactorActivateResponseWithDefaults instantiates a new UserFactorActivateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorActivateResponseWithDefaults() *UserFactorActivateResponse {
	this := UserFactorActivateResponse{}
	return &this
}

// GetFactorType returns the FactorType field value if set, zero value otherwise.
func (o *UserFactorActivateResponse) GetFactorType() string {
	if o == nil || IsNil(o.FactorType) {
		var ret string
		return ret
	}
	return *o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateResponse) GetFactorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FactorType) {
		return nil, false
	}
	return o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorActivateResponse) HasFactorType() bool {
	if o != nil && !IsNil(o.FactorType) {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given string and assigns it to the FactorType field.
func (o *UserFactorActivateResponse) SetFactorType(v string) {
	o.FactorType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserFactorActivateResponse) GetLinks() UserFactorActivateResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret UserFactorActivateResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorActivateResponse) GetLinksOk() (*UserFactorActivateResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactorActivateResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserFactorActivateResponseLinks and assigns it to the Links field.
func (o *UserFactorActivateResponse) SetLinks(v UserFactorActivateResponseLinks) {
	o.Links = &v
}

func (o UserFactorActivateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorActivateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FactorType) {
		toSerialize["factorType"] = o.FactorType
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorActivateResponse) UnmarshalJSON(data []byte) (err error) {
	varUserFactorActivateResponse := _UserFactorActivateResponse{}

	err = json.Unmarshal(data, &varUserFactorActivateResponse)

	if err != nil {
		return err
	}

	*o = UserFactorActivateResponse(varUserFactorActivateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "factorType")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorActivateResponse struct {
	value *UserFactorActivateResponse
	isSet bool
}

func (v NullableUserFactorActivateResponse) Get() *UserFactorActivateResponse {
	return v.value
}

func (v *NullableUserFactorActivateResponse) Set(val *UserFactorActivateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorActivateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorActivateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorActivateResponse(val *UserFactorActivateResponse) *NullableUserFactorActivateResponse {
	return &NullableUserFactorActivateResponse{value: val, isSet: true}
}

func (v NullableUserFactorActivateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorActivateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
