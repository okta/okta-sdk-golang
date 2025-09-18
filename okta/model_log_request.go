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

// checks if the LogRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogRequest{}

// LogRequest The `Request` object describes details that are related to the HTTP request that triggers this event, if available. When the event isn't sourced to an HTTP request, such as an automatic update on the Okta servers, the `Request` object still exists, but the `ipChain` field is empty.
type LogRequest struct {
	// If the incoming request passes through any proxies, the IP addresses of those proxies are stored here in the format of clientIp, proxy1, proxy2, and so on. This field is useful when working with trusted proxies.
	IpChain              []LogIpAddress `json:"ipChain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogRequest LogRequest

// NewLogRequest instantiates a new LogRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogRequest() *LogRequest {
	this := LogRequest{}
	return &this
}

// NewLogRequestWithDefaults instantiates a new LogRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogRequestWithDefaults() *LogRequest {
	this := LogRequest{}
	return &this
}

// GetIpChain returns the IpChain field value if set, zero value otherwise.
func (o *LogRequest) GetIpChain() []LogIpAddress {
	if o == nil || IsNil(o.IpChain) {
		var ret []LogIpAddress
		return ret
	}
	return o.IpChain
}

// GetIpChainOk returns a tuple with the IpChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogRequest) GetIpChainOk() ([]LogIpAddress, bool) {
	if o == nil || IsNil(o.IpChain) {
		return nil, false
	}
	return o.IpChain, true
}

// HasIpChain returns a boolean if a field has been set.
func (o *LogRequest) HasIpChain() bool {
	if o != nil && !IsNil(o.IpChain) {
		return true
	}

	return false
}

// SetIpChain gets a reference to the given []LogIpAddress and assigns it to the IpChain field.
func (o *LogRequest) SetIpChain(v []LogIpAddress) {
	o.IpChain = v
}

func (o LogRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpChain) {
		toSerialize["ipChain"] = o.IpChain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogRequest) UnmarshalJSON(data []byte) (err error) {
	varLogRequest := _LogRequest{}

	err = json.Unmarshal(data, &varLogRequest)

	if err != nil {
		return err
	}

	*o = LogRequest(varLogRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ipChain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogRequest struct {
	value *LogRequest
	isSet bool
}

func (v NullableLogRequest) Get() *LogRequest {
	return v.value
}

func (v *NullableLogRequest) Set(val *LogRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLogRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLogRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogRequest(val *LogRequest) *NullableLogRequest {
	return &NullableLogRequest{value: val, isSet: true}
}

func (v NullableLogRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
