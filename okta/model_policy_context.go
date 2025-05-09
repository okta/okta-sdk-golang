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

// PolicyContext struct for PolicyContext
type PolicyContext struct {
	Device *PolicyContextDevice `json:"device,omitempty"`
	Groups PolicyContextGroups `json:"groups"`
	// The network rule condition, zone, or IP address
	Ip *string `json:"ip,omitempty"`
	Risk *PolicyContextRisk `json:"risk,omitempty"`
	User PolicyContextUser `json:"user"`
	Zones *PolicyContextZones `json:"zones,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyContext PolicyContext

// NewPolicyContext instantiates a new PolicyContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyContext(groups PolicyContextGroups, user PolicyContextUser) *PolicyContext {
	this := PolicyContext{}
	this.Groups = groups
	this.User = user
	return &this
}

// NewPolicyContextWithDefaults instantiates a new PolicyContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyContextWithDefaults() *PolicyContext {
	this := PolicyContext{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *PolicyContext) GetDevice() PolicyContextDevice {
	if o == nil || o.Device == nil {
		var ret PolicyContextDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContext) GetDeviceOk() (*PolicyContextDevice, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *PolicyContext) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given PolicyContextDevice and assigns it to the Device field.
func (o *PolicyContext) SetDevice(v PolicyContextDevice) {
	o.Device = &v
}

// GetGroups returns the Groups field value
func (o *PolicyContext) GetGroups() PolicyContextGroups {
	if o == nil {
		var ret PolicyContextGroups
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *PolicyContext) GetGroupsOk() (*PolicyContextGroups, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value
func (o *PolicyContext) SetGroups(v PolicyContextGroups) {
	o.Groups = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *PolicyContext) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContext) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *PolicyContext) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *PolicyContext) SetIp(v string) {
	o.Ip = &v
}

// GetRisk returns the Risk field value if set, zero value otherwise.
func (o *PolicyContext) GetRisk() PolicyContextRisk {
	if o == nil || o.Risk == nil {
		var ret PolicyContextRisk
		return ret
	}
	return *o.Risk
}

// GetRiskOk returns a tuple with the Risk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContext) GetRiskOk() (*PolicyContextRisk, bool) {
	if o == nil || o.Risk == nil {
		return nil, false
	}
	return o.Risk, true
}

// HasRisk returns a boolean if a field has been set.
func (o *PolicyContext) HasRisk() bool {
	if o != nil && o.Risk != nil {
		return true
	}

	return false
}

// SetRisk gets a reference to the given PolicyContextRisk and assigns it to the Risk field.
func (o *PolicyContext) SetRisk(v PolicyContextRisk) {
	o.Risk = &v
}

// GetUser returns the User field value
func (o *PolicyContext) GetUser() PolicyContextUser {
	if o == nil {
		var ret PolicyContextUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *PolicyContext) GetUserOk() (*PolicyContextUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *PolicyContext) SetUser(v PolicyContextUser) {
	o.User = v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *PolicyContext) GetZones() PolicyContextZones {
	if o == nil || o.Zones == nil {
		var ret PolicyContextZones
		return ret
	}
	return *o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyContext) GetZonesOk() (*PolicyContextZones, bool) {
	if o == nil || o.Zones == nil {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *PolicyContext) HasZones() bool {
	if o != nil && o.Zones != nil {
		return true
	}

	return false
}

// SetZones gets a reference to the given PolicyContextZones and assigns it to the Zones field.
func (o *PolicyContext) SetZones(v PolicyContextZones) {
	o.Zones = &v
}

func (o PolicyContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if true {
		toSerialize["groups"] = o.Groups
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.Risk != nil {
		toSerialize["risk"] = o.Risk
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.Zones != nil {
		toSerialize["zones"] = o.Zones
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyContext) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyContext := _PolicyContext{}

	err = json.Unmarshal(bytes, &varPolicyContext)
	if err == nil {
		*o = PolicyContext(varPolicyContext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "risk")
		delete(additionalProperties, "user")
		delete(additionalProperties, "zones")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyContext struct {
	value *PolicyContext
	isSet bool
}

func (v NullablePolicyContext) Get() *PolicyContext {
	return v.value
}

func (v *NullablePolicyContext) Set(val *PolicyContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyContext(val *PolicyContext) *NullablePolicyContext {
	return &NullablePolicyContext{value: val, isSet: true}
}

func (v NullablePolicyContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

