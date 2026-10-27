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
	"fmt"
	"time"
)

// checks if the DeviceSessionFactor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceSessionFactor{}

// DeviceSessionFactor struct for DeviceSessionFactor
type DeviceSessionFactor struct {
	// The human-readable authenticator name that's displayed.  There are three places where the display name can be retrieved from, which are evaluated in a priority order. 1. The display name is retrieved from the org's configured authenticator name. 2. The display name is retrieved from the authenticator catalog default. 3. For legacy verifications with no authenticator key, the display name is retrieved from the enrollment's own display name.  The display name is omitted only when all three are absent.
	DisplayName *string `json:"displayName,omitempty"`
	// The authenticator key, `webauthn` or `okta_verify`, for example. This property is omitted for legacy verifications that carry no authenticator key.
	Key *string `json:"key,omitempty"`
	// Timestamp when this factor was last verified in the session
	LastUsedOn time.Time `json:"lastUsedOn"`
	// The authenticator method name, `password`, `signed_nonce`, `push`, or `otp`, for example
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _DeviceSessionFactor DeviceSessionFactor

// NewDeviceSessionFactor instantiates a new DeviceSessionFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceSessionFactor(lastUsedOn time.Time, type_ string) *DeviceSessionFactor {
	this := DeviceSessionFactor{}
	this.LastUsedOn = lastUsedOn
	this.Type = type_
	return &this
}

// NewDeviceSessionFactorWithDefaults instantiates a new DeviceSessionFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceSessionFactorWithDefaults() *DeviceSessionFactor {
	this := DeviceSessionFactor{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DeviceSessionFactor) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSessionFactor) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DeviceSessionFactor) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DeviceSessionFactor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DeviceSessionFactor) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceSessionFactor) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DeviceSessionFactor) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DeviceSessionFactor) SetKey(v string) {
	o.Key = &v
}

// GetLastUsedOn returns the LastUsedOn field value
func (o *DeviceSessionFactor) GetLastUsedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUsedOn
}

// GetLastUsedOnOk returns a tuple with the LastUsedOn field value
// and a boolean to check if the value has been set.
func (o *DeviceSessionFactor) GetLastUsedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUsedOn, true
}

// SetLastUsedOn sets field value
func (o *DeviceSessionFactor) SetLastUsedOn(v time.Time) {
	o.LastUsedOn = v
}

// GetType returns the Type field value
func (o *DeviceSessionFactor) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeviceSessionFactor) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeviceSessionFactor) SetType(v string) {
	o.Type = v
}

func (o DeviceSessionFactor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceSessionFactor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["lastUsedOn"] = o.LastUsedOn
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceSessionFactor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lastUsedOn",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeviceSessionFactor := _DeviceSessionFactor{}

	err = json.Unmarshal(data, &varDeviceSessionFactor)

	if err != nil {
		return err
	}

	*o = DeviceSessionFactor(varDeviceSessionFactor)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "key")
		delete(additionalProperties, "lastUsedOn")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceSessionFactor struct {
	value *DeviceSessionFactor
	isSet bool
}

func (v NullableDeviceSessionFactor) Get() *DeviceSessionFactor {
	return v.value
}

func (v *NullableDeviceSessionFactor) Set(val *DeviceSessionFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceSessionFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceSessionFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceSessionFactor(val *DeviceSessionFactor) *NullableDeviceSessionFactor {
	return &NullableDeviceSessionFactor{value: val, isSet: true}
}

func (v NullableDeviceSessionFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceSessionFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
