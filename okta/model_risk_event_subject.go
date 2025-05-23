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

// RiskEventSubject struct for RiskEventSubject
type RiskEventSubject struct {
	// The risk event subject IP address (either an IPv4 or IPv6 address)
	Ip string `json:"ip"`
	// Additional reasons for the risk level of the IP
	Message *string `json:"message,omitempty"`
	// The risk level associated with the IP
	RiskLevel string `json:"riskLevel"`
	AdditionalProperties map[string]interface{}
}

type _RiskEventSubject RiskEventSubject

// NewRiskEventSubject instantiates a new RiskEventSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEventSubject(ip string, riskLevel string) *RiskEventSubject {
	this := RiskEventSubject{}
	this.Ip = ip
	this.RiskLevel = riskLevel
	return &this
}

// NewRiskEventSubjectWithDefaults instantiates a new RiskEventSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEventSubjectWithDefaults() *RiskEventSubject {
	this := RiskEventSubject{}
	return &this
}

// GetIp returns the Ip field value
func (o *RiskEventSubject) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *RiskEventSubject) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *RiskEventSubject) SetIp(v string) {
	o.Ip = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RiskEventSubject) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEventSubject) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RiskEventSubject) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RiskEventSubject) SetMessage(v string) {
	o.Message = &v
}

// GetRiskLevel returns the RiskLevel field value
func (o *RiskEventSubject) GetRiskLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RiskLevel
}

// GetRiskLevelOk returns a tuple with the RiskLevel field value
// and a boolean to check if the value has been set.
func (o *RiskEventSubject) GetRiskLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskLevel, true
}

// SetRiskLevel sets field value
func (o *RiskEventSubject) SetRiskLevel(v string) {
	o.RiskLevel = v
}

func (o RiskEventSubject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ip"] = o.Ip
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["riskLevel"] = o.RiskLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskEventSubject) UnmarshalJSON(bytes []byte) (err error) {
	varRiskEventSubject := _RiskEventSubject{}

	err = json.Unmarshal(bytes, &varRiskEventSubject)
	if err == nil {
		*o = RiskEventSubject(varRiskEventSubject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "ip")
		delete(additionalProperties, "message")
		delete(additionalProperties, "riskLevel")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskEventSubject struct {
	value *RiskEventSubject
	isSet bool
}

func (v NullableRiskEventSubject) Get() *RiskEventSubject {
	return v.value
}

func (v *NullableRiskEventSubject) Set(val *RiskEventSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEventSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEventSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEventSubject(val *RiskEventSubject) *NullableRiskEventSubject {
	return &NullableRiskEventSubject{value: val, isSet: true}
}

func (v NullableRiskEventSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEventSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

