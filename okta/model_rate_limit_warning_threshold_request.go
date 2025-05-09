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

// RateLimitWarningThresholdRequest 
type RateLimitWarningThresholdRequest struct {
	// The threshold value (percentage) of a rate limit that, when exceeded, triggers a warning notification. By default, this value is 90 for Workforce orgs and 60 for CIAM orgs.
	WarningThreshold int32 `json:"warningThreshold"`
	AdditionalProperties map[string]interface{}
}

type _RateLimitWarningThresholdRequest RateLimitWarningThresholdRequest

// NewRateLimitWarningThresholdRequest instantiates a new RateLimitWarningThresholdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitWarningThresholdRequest(warningThreshold int32) *RateLimitWarningThresholdRequest {
	this := RateLimitWarningThresholdRequest{}
	this.WarningThreshold = warningThreshold
	return &this
}

// NewRateLimitWarningThresholdRequestWithDefaults instantiates a new RateLimitWarningThresholdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitWarningThresholdRequestWithDefaults() *RateLimitWarningThresholdRequest {
	this := RateLimitWarningThresholdRequest{}
	return &this
}

// GetWarningThreshold returns the WarningThreshold field value
func (o *RateLimitWarningThresholdRequest) GetWarningThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value
// and a boolean to check if the value has been set.
func (o *RateLimitWarningThresholdRequest) GetWarningThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarningThreshold, true
}

// SetWarningThreshold sets field value
func (o *RateLimitWarningThresholdRequest) SetWarningThreshold(v int32) {
	o.WarningThreshold = v
}

func (o RateLimitWarningThresholdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["warningThreshold"] = o.WarningThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RateLimitWarningThresholdRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRateLimitWarningThresholdRequest := _RateLimitWarningThresholdRequest{}

	err = json.Unmarshal(bytes, &varRateLimitWarningThresholdRequest)
	if err == nil {
		*o = RateLimitWarningThresholdRequest(varRateLimitWarningThresholdRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "warningThreshold")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRateLimitWarningThresholdRequest struct {
	value *RateLimitWarningThresholdRequest
	isSet bool
}

func (v NullableRateLimitWarningThresholdRequest) Get() *RateLimitWarningThresholdRequest {
	return v.value
}

func (v *NullableRateLimitWarningThresholdRequest) Set(val *RateLimitWarningThresholdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitWarningThresholdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitWarningThresholdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitWarningThresholdRequest(val *RateLimitWarningThresholdRequest) *NullableRateLimitWarningThresholdRequest {
	return &NullableRateLimitWarningThresholdRequest{value: val, isSet: true}
}

func (v NullableRateLimitWarningThresholdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitWarningThresholdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

