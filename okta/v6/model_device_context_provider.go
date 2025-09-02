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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// DeviceContextProvider struct for DeviceContextProvider
type DeviceContextProvider struct {
	// Unique identifier for the device context provider
	Id *string `json:"id,omitempty"`
	// Identifies the type of device context provider
	Key string `json:"key"`
	// Whether or not the device context provider is used to identify the user. `IGNORE` prevents the device context provider from being used to authenticate the user. Identification of the device and device context collection happens regardless of this setting.
	UserIdentification *string `json:"userIdentification,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceContextProvider DeviceContextProvider

// NewDeviceContextProvider instantiates a new DeviceContextProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceContextProvider(key string) *DeviceContextProvider {
	this := DeviceContextProvider{}
	this.Key = key
	return &this
}

// NewDeviceContextProviderWithDefaults instantiates a new DeviceContextProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceContextProviderWithDefaults() *DeviceContextProvider {
	this := DeviceContextProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceContextProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceContextProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceContextProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceContextProvider) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value
func (o *DeviceContextProvider) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DeviceContextProvider) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DeviceContextProvider) SetKey(v string) {
	o.Key = v
}

// GetUserIdentification returns the UserIdentification field value if set, zero value otherwise.
func (o *DeviceContextProvider) GetUserIdentification() string {
	if o == nil || o.UserIdentification == nil {
		var ret string
		return ret
	}
	return *o.UserIdentification
}

// GetUserIdentificationOk returns a tuple with the UserIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceContextProvider) GetUserIdentificationOk() (*string, bool) {
	if o == nil || o.UserIdentification == nil {
		return nil, false
	}
	return o.UserIdentification, true
}

// HasUserIdentification returns a boolean if a field has been set.
func (o *DeviceContextProvider) HasUserIdentification() bool {
	if o != nil && o.UserIdentification != nil {
		return true
	}

	return false
}

// SetUserIdentification gets a reference to the given string and assigns it to the UserIdentification field.
func (o *DeviceContextProvider) SetUserIdentification(v string) {
	o.UserIdentification = &v
}

func (o DeviceContextProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.UserIdentification != nil {
		toSerialize["userIdentification"] = o.UserIdentification
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceContextProvider) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceContextProvider := _DeviceContextProvider{}

	err = json.Unmarshal(bytes, &varDeviceContextProvider)
	if err == nil {
		*o = DeviceContextProvider(varDeviceContextProvider)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "userIdentification")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceContextProvider struct {
	value *DeviceContextProvider
	isSet bool
}

func (v NullableDeviceContextProvider) Get() *DeviceContextProvider {
	return v.value
}

func (v *NullableDeviceContextProvider) Set(val *DeviceContextProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceContextProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceContextProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceContextProvider(val *DeviceContextProvider) *NullableDeviceContextProvider {
	return &NullableDeviceContextProvider{value: val, isSet: true}
}

func (v NullableDeviceContextProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceContextProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

