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

// checks if the UserRiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRiskRequest{}

// UserRiskRequest struct for UserRiskRequest
type UserRiskRequest struct {
	// The risk level associated with the user
	RiskLevel            *string `json:"riskLevel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserRiskRequest UserRiskRequest

// NewUserRiskRequest instantiates a new UserRiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRiskRequest() *UserRiskRequest {
	this := UserRiskRequest{}
	return &this
}

// NewUserRiskRequestWithDefaults instantiates a new UserRiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRiskRequestWithDefaults() *UserRiskRequest {
	this := UserRiskRequest{}
	return &this
}

// GetRiskLevel returns the RiskLevel field value if set, zero value otherwise.
func (o *UserRiskRequest) GetRiskLevel() string {
	if o == nil || IsNil(o.RiskLevel) {
		var ret string
		return ret
	}
	return *o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRiskRequest) GetRiskLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RiskLevel) {
		return nil, false
	}
	return o.RiskLevel, true
}

// HasRiskLevel returns a boolean if a field has been set.
func (o *UserRiskRequest) HasRiskLevel() bool {
	if o != nil && !IsNil(o.RiskLevel) {
		return true
	}

	return false
}

// SetRiskLevel gets a reference to the given string and assigns it to the RiskLevel field.
func (o *UserRiskRequest) SetRiskLevel(v string) {
	o.RiskLevel = &v
}

func (o UserRiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RiskLevel) {
		toSerialize["riskLevel"] = o.RiskLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserRiskRequest) UnmarshalJSON(data []byte) (err error) {
	varUserRiskRequest := _UserRiskRequest{}

	err = json.Unmarshal(data, &varUserRiskRequest)

	if err != nil {
		return err
	}

	*o = UserRiskRequest(varUserRiskRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "riskLevel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserRiskRequest struct {
	value *UserRiskRequest
	isSet bool
}

func (v NullableUserRiskRequest) Get() *UserRiskRequest {
	return v.value
}

func (v *NullableUserRiskRequest) Set(val *UserRiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRiskRequest(val *UserRiskRequest) *NullableUserRiskRequest {
	return &NullableUserRiskRequest{value: val, isSet: true}
}

func (v NullableUserRiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
