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

// checks if the UserRiskPutResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRiskPutResponse{}

// UserRiskPutResponse struct for UserRiskPutResponse
type UserRiskPutResponse struct {
	// Describes the risk level for the user
	Reason *string `json:"reason,omitempty"`
	// The risk level associated with the user
	RiskLevel            *string                   `json:"riskLevel,omitempty"`
	Links                *UserRiskGetResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserRiskPutResponse UserRiskPutResponse

// NewUserRiskPutResponse instantiates a new UserRiskPutResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRiskPutResponse() *UserRiskPutResponse {
	this := UserRiskPutResponse{}
	return &this
}

// NewUserRiskPutResponseWithDefaults instantiates a new UserRiskPutResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRiskPutResponseWithDefaults() *UserRiskPutResponse {
	this := UserRiskPutResponse{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *UserRiskPutResponse) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskPutResponse) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *UserRiskPutResponse) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *UserRiskPutResponse) SetReason(v string) {
	o.Reason = &v
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *UserRiskPutResponse) GetRiskLevel() string {
	if o == nil || IsNil(o.RiskLevel) {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskPutResponse) GetRiskLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RiskLevel) {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *UserRiskPutResponse) HasRiskLevel() bool {
	if o != nil && !IsNil(o.RiskLevel) {
		return true
	}

	return false
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *UserRiskPutResponse) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserRiskPutResponse) GetLinks() UserRiskGetResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret UserRiskGetResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskPutResponse) GetLinksOk() (*UserRiskGetResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserRiskPutResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserRiskGetResponseLinks and assigns it to the Links field.
func (o *UserRiskPutResponse) SetLinks(v UserRiskGetResponseLinks) {
	o.Links = &v
}

func (o UserRiskPutResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRiskPutResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
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

func (o *UserRiskPutResponse) UnmarshalJSON(data []byte) (err error) {
	varUserRiskPutResponse := _UserRiskPutResponse{}

	err = json.Unmarshal(data, &varUserRiskPutResponse)

	if err != nil {
		return err
	}

	*o = UserRiskPutResponse(varUserRiskPutResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reason")
		delete(additionalProperties, "riskLevel")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserRiskPutResponse struct {
	value *UserRiskPutResponse
	isSet bool
}

func (v NullableUserRiskPutResponse) Get() *UserRiskPutResponse {
	return v.value
}

func (v *NullableUserRiskPutResponse) Set(val *UserRiskPutResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRiskPutResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRiskPutResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRiskPutResponse(val *UserRiskPutResponse) *NullableUserRiskPutResponse {
	return &NullableUserRiskPutResponse{value: val, isSet: true}
}

func (v NullableUserRiskPutResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRiskPutResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
