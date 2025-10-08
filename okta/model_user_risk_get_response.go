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

// checks if the UserRiskGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRiskGetResponse{}

// UserRiskGetResponse struct for UserRiskGetResponse
type UserRiskGetResponse struct {
	// The risk level associated with the user
	RiskLevel            *string                   `json:"riskLevel,omitempty"`
	Links                *UserRiskGetResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserRiskGetResponse UserRiskGetResponse

// NewUserRiskGetResponse instantiates a new UserRiskGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRiskGetResponse() *UserRiskGetResponse {
	this := UserRiskGetResponse{}
	return &this
}

// NewUserRiskGetResponseWithDefaults instantiates a new UserRiskGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRiskGetResponseWithDefaults() *UserRiskGetResponse {
	this := UserRiskGetResponse{}
	return &this
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *UserRiskGetResponse) GetRiskLevel() string {
	if o == nil || IsNil(o.RiskLevel) {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskGetResponse) GetRiskLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RiskLevel) {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *UserRiskGetResponse) HasRiskLevel() bool {
	if o != nil && !IsNil(o.RiskLevel) {
		return true
	}

	return false
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *UserRiskGetResponse) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserRiskGetResponse) GetLinks() UserRiskGetResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret UserRiskGetResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskGetResponse) GetLinksOk() (*UserRiskGetResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserRiskGetResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserRiskGetResponseLinks and assigns it to the Links field.
func (o *UserRiskGetResponse) SetLinks(v UserRiskGetResponseLinks) {
	o.Links = &v
}

func (o UserRiskGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRiskGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RiskLevel) {
		toSerialize["riskLevel"] = o.RiskLevel
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserRiskGetResponse) UnmarshalJSON(data []byte) (err error) {
	varUserRiskGetResponse := _UserRiskGetResponse{}

	err = json.Unmarshal(data, &varUserRiskGetResponse)

	if err != nil {
		return err
	}

	*o = UserRiskGetResponse(varUserRiskGetResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "riskLevel")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserRiskGetResponse struct {
	value *UserRiskGetResponse
	isSet bool
}

func (v NullableUserRiskGetResponse) Get() *UserRiskGetResponse {
	return v.value
}

func (v *NullableUserRiskGetResponse) Set(val *UserRiskGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRiskGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRiskGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRiskGetResponse(val *UserRiskGetResponse) *NullableUserRiskGetResponse {
	return &NullableUserRiskGetResponse{value: val, isSet: true}
}

func (v NullableUserRiskGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRiskGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
