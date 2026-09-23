/*
Okta Admin Management API

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

// checks if the SelectiveRenewalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectiveRenewalRequest{}

// SelectiveRenewalRequest Which SCEP configurations to migrate onto the renewed certificate, and how. Name at least one configuration, and name each one only once across both lists.
type SelectiveRenewalRequest struct {
	// The configurations to clone against the renewed certificate. Each source configuration stays on the certificate it's already bound to, so both certificates can enroll devices during the migration.
	ParallelConfigs []ParallelScepConfig `json:"parallelConfigs,omitempty"`
	// The IDs of the configurations to move onto the renewed certificate authority. Each keeps its ID and SCEP enrollment URL.
	RolloverConfigIds    []string `json:"rolloverConfigIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelectiveRenewalRequest SelectiveRenewalRequest

// NewSelectiveRenewalRequest instantiates a new SelectiveRenewalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectiveRenewalRequest() *SelectiveRenewalRequest {
	this := SelectiveRenewalRequest{}
	return &this
}

// NewSelectiveRenewalRequestWithDefaults instantiates a new SelectiveRenewalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectiveRenewalRequestWithDefaults() *SelectiveRenewalRequest {
	this := SelectiveRenewalRequest{}
	return &this
}

// GetParallelConfigs returns the ParallelConfigs field value if set, zero value otherwise.
func (o *SelectiveRenewalRequest) GetParallelConfigs() []ParallelScepConfig {
	if o == nil || IsNil(o.ParallelConfigs) {
		var ret []ParallelScepConfig
		return ret
	}
	return o.ParallelConfigs
}

// GetParallelConfigsOk returns a tuple with the ParallelConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveRenewalRequest) GetParallelConfigsOk() ([]ParallelScepConfig, bool) {
	if o == nil || IsNil(o.ParallelConfigs) {
		return nil, false
	}
	return o.ParallelConfigs, true
}

// HasParallelConfigs returns a boolean if a field has been set.
func (o *SelectiveRenewalRequest) HasParallelConfigs() bool {
	if o != nil && !IsNil(o.ParallelConfigs) {
		return true
	}

	return false
}

// SetParallelConfigs gets a reference to the given []ParallelScepConfig and assigns it to the ParallelConfigs field.
func (o *SelectiveRenewalRequest) SetParallelConfigs(v []ParallelScepConfig) {
	o.ParallelConfigs = v
}

// GetRolloverConfigIds returns the RolloverConfigIds field value if set, zero value otherwise.
func (o *SelectiveRenewalRequest) GetRolloverConfigIds() []string {
	if o == nil || IsNil(o.RolloverConfigIds) {
		var ret []string
		return ret
	}
	return o.RolloverConfigIds
}

// GetRolloverConfigIdsOk returns a tuple with the RolloverConfigIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectiveRenewalRequest) GetRolloverConfigIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RolloverConfigIds) {
		return nil, false
	}
	return o.RolloverConfigIds, true
}

// HasRolloverConfigIds returns a boolean if a field has been set.
func (o *SelectiveRenewalRequest) HasRolloverConfigIds() bool {
	if o != nil && !IsNil(o.RolloverConfigIds) {
		return true
	}

	return false
}

// SetRolloverConfigIds gets a reference to the given []string and assigns it to the RolloverConfigIds field.
func (o *SelectiveRenewalRequest) SetRolloverConfigIds(v []string) {
	o.RolloverConfigIds = v
}

func (o SelectiveRenewalRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelectiveRenewalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParallelConfigs) {
		toSerialize["parallelConfigs"] = o.ParallelConfigs
	}
	if !IsNil(o.RolloverConfigIds) {
		toSerialize["rolloverConfigIds"] = o.RolloverConfigIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelectiveRenewalRequest) UnmarshalJSON(data []byte) (err error) {
	varSelectiveRenewalRequest := _SelectiveRenewalRequest{}

	err = json.Unmarshal(data, &varSelectiveRenewalRequest)

	if err != nil {
		return err
	}

	*o = SelectiveRenewalRequest(varSelectiveRenewalRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "parallelConfigs")
		delete(additionalProperties, "rolloverConfigIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelectiveRenewalRequest struct {
	value *SelectiveRenewalRequest
	isSet bool
}

func (v NullableSelectiveRenewalRequest) Get() *SelectiveRenewalRequest {
	return v.value
}

func (v *NullableSelectiveRenewalRequest) Set(val *SelectiveRenewalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectiveRenewalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectiveRenewalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectiveRenewalRequest(val *SelectiveRenewalRequest) *NullableSelectiveRenewalRequest {
	return &NullableSelectiveRenewalRequest{value: val, isSet: true}
}

func (v NullableSelectiveRenewalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectiveRenewalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
