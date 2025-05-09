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

// RateLimitWarningThresholdResponse 
type RateLimitWarningThresholdResponse struct {
	// The threshold value (percentage) of a rate limit that, when exceeded, triggers a warning notification. By default, this value is 90 for Workforce orgs and 60 for CIAM orgs.
	WarningThreshold *int32 `json:"warningThreshold,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RateLimitWarningThresholdResponse RateLimitWarningThresholdResponse

// NewRateLimitWarningThresholdResponse instantiates a new RateLimitWarningThresholdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitWarningThresholdResponse() *RateLimitWarningThresholdResponse {
	this := RateLimitWarningThresholdResponse{}
	return &this
}

// NewRateLimitWarningThresholdResponseWithDefaults instantiates a new RateLimitWarningThresholdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitWarningThresholdResponseWithDefaults() *RateLimitWarningThresholdResponse {
	this := RateLimitWarningThresholdResponse{}
	return &this
}

// GetWarningThreshold returns the WarningThreshold field value if set, zero value otherwise.
func (o *RateLimitWarningThresholdResponse) GetWarningThreshold() int32 {
	if o == nil || o.WarningThreshold == nil {
		var ret int32
		return ret
	}
	return *o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitWarningThresholdResponse) GetWarningThresholdOk() (*int32, bool) {
	if o == nil || o.WarningThreshold == nil {
		return nil, false
	}
	return o.WarningThreshold, true
}

// HasWarningThreshold returns a boolean if a field has been set.
func (o *RateLimitWarningThresholdResponse) HasWarningThreshold() bool {
	if o != nil && o.WarningThreshold != nil {
		return true
	}

	return false
}

// SetWarningThreshold gets a reference to the given int32 and assigns it to the WarningThreshold field.
func (o *RateLimitWarningThresholdResponse) SetWarningThreshold(v int32) {
	o.WarningThreshold = &v
}

func (o RateLimitWarningThresholdResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WarningThreshold != nil {
		toSerialize["warningThreshold"] = o.WarningThreshold
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RateLimitWarningThresholdResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRateLimitWarningThresholdResponse := _RateLimitWarningThresholdResponse{}

	err = json.Unmarshal(bytes, &varRateLimitWarningThresholdResponse)
	if err == nil {
		*o = RateLimitWarningThresholdResponse(varRateLimitWarningThresholdResponse)
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

type NullableRateLimitWarningThresholdResponse struct {
	value *RateLimitWarningThresholdResponse
	isSet bool
}

func (v NullableRateLimitWarningThresholdResponse) Get() *RateLimitWarningThresholdResponse {
	return v.value
}

func (v *NullableRateLimitWarningThresholdResponse) Set(val *RateLimitWarningThresholdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitWarningThresholdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitWarningThresholdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitWarningThresholdResponse(val *RateLimitWarningThresholdResponse) *NullableRateLimitWarningThresholdResponse {
	return &NullableRateLimitWarningThresholdResponse{value: val, isSet: true}
}

func (v NullableRateLimitWarningThresholdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitWarningThresholdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

