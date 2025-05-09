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

// SecurityEventSubject The event subjects
type SecurityEventSubject struct {
	// The device involved with the event
	Device map[string]interface{} `json:"device,omitempty"`
	// The tenant involved with the event
	Tenant map[string]interface{} `json:"tenant,omitempty"`
	// The user involved with the event
	User map[string]interface{} `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventSubject SecurityEventSubject

// NewSecurityEventSubject instantiates a new SecurityEventSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventSubject() *SecurityEventSubject {
	this := SecurityEventSubject{}
	return &this
}

// NewSecurityEventSubjectWithDefaults instantiates a new SecurityEventSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventSubjectWithDefaults() *SecurityEventSubject {
	this := SecurityEventSubject{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *SecurityEventSubject) GetDevice() map[string]interface{} {
	if o == nil || o.Device == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventSubject) GetDeviceOk() (map[string]interface{}, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *SecurityEventSubject) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given map[string]interface{} and assigns it to the Device field.
func (o *SecurityEventSubject) SetDevice(v map[string]interface{}) {
	o.Device = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *SecurityEventSubject) GetTenant() map[string]interface{} {
	if o == nil || o.Tenant == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventSubject) GetTenantOk() (map[string]interface{}, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *SecurityEventSubject) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given map[string]interface{} and assigns it to the Tenant field.
func (o *SecurityEventSubject) SetTenant(v map[string]interface{}) {
	o.Tenant = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SecurityEventSubject) GetUser() map[string]interface{} {
	if o == nil || o.User == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventSubject) GetUserOk() (map[string]interface{}, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SecurityEventSubject) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given map[string]interface{} and assigns it to the User field.
func (o *SecurityEventSubject) SetUser(v map[string]interface{}) {
	o.User = v
}

func (o SecurityEventSubject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventSubject) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventSubject := _SecurityEventSubject{}

	err = json.Unmarshal(bytes, &varSecurityEventSubject)
	if err == nil {
		*o = SecurityEventSubject(varSecurityEventSubject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventSubject struct {
	value *SecurityEventSubject
	isSet bool
}

func (v NullableSecurityEventSubject) Get() *SecurityEventSubject {
	return v.value
}

func (v *NullableSecurityEventSubject) Set(val *SecurityEventSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventSubject(val *SecurityEventSubject) *NullableSecurityEventSubject {
	return &NullableSecurityEventSubject{value: val, isSet: true}
}

func (v NullableSecurityEventSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

