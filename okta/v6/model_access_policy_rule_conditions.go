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

// checks if the AccessPolicyRuleConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPolicyRuleConditions{}

// AccessPolicyRuleConditions struct for AccessPolicyRuleConditions
type AccessPolicyRuleConditions struct {
	Device               *DeviceAccessPolicyRuleCondition `json:"device,omitempty"`
	ElCondition          *AccessPolicyRuleCustomCondition `json:"elCondition,omitempty"`
	Network              *PolicyNetworkCondition          `json:"network,omitempty"`
	People               *PolicyPeopleCondition           `json:"people,omitempty"`
	Platform             *PlatformPolicyRuleCondition     `json:"platform,omitempty"`
	RiskScore            *RiskScorePolicyRuleCondition    `json:"riskScore,omitempty"`
	UserType             *UserTypeCondition               `json:"userType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyRuleConditions AccessPolicyRuleConditions

// NewAccessPolicyRuleConditions instantiates a new AccessPolicyRuleConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyRuleConditions() *AccessPolicyRuleConditions {
	this := AccessPolicyRuleConditions{}
	return &this
}

// NewAccessPolicyRuleConditionsWithDefaults instantiates a new AccessPolicyRuleConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyRuleConditionsWithDefaults() *AccessPolicyRuleConditions {
	this := AccessPolicyRuleConditions{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetDevice() DeviceAccessPolicyRuleCondition {
	if o == nil || IsNil(o.Device) {
		var ret DeviceAccessPolicyRuleCondition
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetDeviceOk() (*DeviceAccessPolicyRuleCondition, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given DeviceAccessPolicyRuleCondition and assigns it to the Device field.
func (o *AccessPolicyRuleConditions) SetDevice(v DeviceAccessPolicyRuleCondition) {
	o.Device = &v
}

// GetElCondition returns the ElCondition field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetElCondition() AccessPolicyRuleCustomCondition {
	if o == nil || IsNil(o.ElCondition) {
		var ret AccessPolicyRuleCustomCondition
		return ret
	}
	return *o.ElCondition
}

// GetElConditionOk returns a tuple with the ElCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetElConditionOk() (*AccessPolicyRuleCustomCondition, bool) {
	if o == nil || IsNil(o.ElCondition) {
		return nil, false
	}
	return o.ElCondition, true
}

// HasElCondition returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasElCondition() bool {
	if o != nil && !IsNil(o.ElCondition) {
		return true
	}

	return false
}

// SetElCondition gets a reference to the given AccessPolicyRuleCustomCondition and assigns it to the ElCondition field.
func (o *AccessPolicyRuleConditions) SetElCondition(v AccessPolicyRuleCustomCondition) {
	o.ElCondition = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetNetwork() PolicyNetworkCondition {
	if o == nil || IsNil(o.Network) {
		var ret PolicyNetworkCondition
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetNetworkOk() (*PolicyNetworkCondition, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given PolicyNetworkCondition and assigns it to the Network field.
func (o *AccessPolicyRuleConditions) SetNetwork(v PolicyNetworkCondition) {
	o.Network = &v
}

// GetPeople returns the People field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetPeople() PolicyPeopleCondition {
	if o == nil || IsNil(o.People) {
		var ret PolicyPeopleCondition
		return ret
	}
	return *o.People
}

// GetPeopleOk returns a tuple with the People field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetPeopleOk() (*PolicyPeopleCondition, bool) {
	if o == nil || IsNil(o.People) {
		return nil, false
	}
	return o.People, true
}

// HasPeople returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasPeople() bool {
	if o != nil && !IsNil(o.People) {
		return true
	}

	return false
}

// SetPeople gets a reference to the given PolicyPeopleCondition and assigns it to the People field.
func (o *AccessPolicyRuleConditions) SetPeople(v PolicyPeopleCondition) {
	o.People = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetPlatform() PlatformPolicyRuleCondition {
	if o == nil || IsNil(o.Platform) {
		var ret PlatformPolicyRuleCondition
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetPlatformOk() (*PlatformPolicyRuleCondition, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given PlatformPolicyRuleCondition and assigns it to the Platform field.
func (o *AccessPolicyRuleConditions) SetPlatform(v PlatformPolicyRuleCondition) {
	o.Platform = &v
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetRiskScore() RiskScorePolicyRuleCondition {
	if o == nil || IsNil(o.RiskScore) {
		var ret RiskScorePolicyRuleCondition
		return ret
	}
	return *o.RiskScore
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetRiskScoreOk() (*RiskScorePolicyRuleCondition, bool) {
	if o == nil || IsNil(o.RiskScore) {
		return nil, false
	}
	return o.RiskScore, true
}

// HasRiskScore returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasRiskScore() bool {
	if o != nil && !IsNil(o.RiskScore) {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given RiskScorePolicyRuleCondition and assigns it to the RiskScore field.
func (o *AccessPolicyRuleConditions) SetRiskScore(v RiskScorePolicyRuleCondition) {
	o.RiskScore = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *AccessPolicyRuleConditions) GetUserType() UserTypeCondition {
	if o == nil || IsNil(o.UserType) {
		var ret UserTypeCondition
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyRuleConditions) GetUserTypeOk() (*UserTypeCondition, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *AccessPolicyRuleConditions) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given UserTypeCondition and assigns it to the UserType field.
func (o *AccessPolicyRuleConditions) SetUserType(v UserTypeCondition) {
	o.UserType = &v
}

func (o AccessPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessPolicyRuleConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.ElCondition) {
		toSerialize["elCondition"] = o.ElCondition
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.People) {
		toSerialize["people"] = o.People
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.RiskScore) {
		toSerialize["riskScore"] = o.RiskScore
	}
	if !IsNil(o.UserType) {
		toSerialize["userType"] = o.UserType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessPolicyRuleConditions) UnmarshalJSON(data []byte) (err error) {
	varAccessPolicyRuleConditions := _AccessPolicyRuleConditions{}

	err = json.Unmarshal(data, &varAccessPolicyRuleConditions)

	if err != nil {
		return err
	}

	*o = AccessPolicyRuleConditions(varAccessPolicyRuleConditions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "elCondition")
		delete(additionalProperties, "network")
		delete(additionalProperties, "people")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "riskScore")
		delete(additionalProperties, "userType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessPolicyRuleConditions struct {
	value *AccessPolicyRuleConditions
	isSet bool
}

func (v NullableAccessPolicyRuleConditions) Get() *AccessPolicyRuleConditions {
	return v.value
}

func (v *NullableAccessPolicyRuleConditions) Set(val *AccessPolicyRuleConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyRuleConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyRuleConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyRuleConditions(val *AccessPolicyRuleConditions) *NullableAccessPolicyRuleConditions {
	return &NullableAccessPolicyRuleConditions{value: val, isSet: true}
}

func (v NullableAccessPolicyRuleConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyRuleConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
