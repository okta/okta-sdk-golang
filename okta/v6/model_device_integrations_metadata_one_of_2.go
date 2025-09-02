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

// DeviceIntegrationsMetadataOneOf2 struct for DeviceIntegrationsMetadataOneOf2
type DeviceIntegrationsMetadataOneOf2 struct {
	Type string `json:"type"`
	IdpId string `json:"idpId"`
	AdditionalProperties map[string]interface{}
}

type _DeviceIntegrationsMetadataOneOf2 DeviceIntegrationsMetadataOneOf2

// NewDeviceIntegrationsMetadataOneOf2 instantiates a new DeviceIntegrationsMetadataOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceIntegrationsMetadataOneOf2(type_ string, idpId string) *DeviceIntegrationsMetadataOneOf2 {
	this := DeviceIntegrationsMetadataOneOf2{}
	this.Type = type_
	this.IdpId = idpId
	return &this
}

// NewDeviceIntegrationsMetadataOneOf2WithDefaults instantiates a new DeviceIntegrationsMetadataOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIntegrationsMetadataOneOf2WithDefaults() *DeviceIntegrationsMetadataOneOf2 {
	this := DeviceIntegrationsMetadataOneOf2{}
	return &this
}

// GetType returns the Type field value
func (o *DeviceIntegrationsMetadataOneOf2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeviceIntegrationsMetadataOneOf2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeviceIntegrationsMetadataOneOf2) SetType(v string) {
	o.Type = v
}

// GetIdpId returns the IdpId field value
func (o *DeviceIntegrationsMetadataOneOf2) GetIdpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpId
}

// GetIdpIdOk returns a tuple with the IdpId field value
// and a boolean to check if the value has been set.
func (o *DeviceIntegrationsMetadataOneOf2) GetIdpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpId, true
}

// SetIdpId sets field value
func (o *DeviceIntegrationsMetadataOneOf2) SetIdpId(v string) {
	o.IdpId = v
}

func (o DeviceIntegrationsMetadataOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["idpId"] = o.IdpId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceIntegrationsMetadataOneOf2) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceIntegrationsMetadataOneOf2 := _DeviceIntegrationsMetadataOneOf2{}

	err = json.Unmarshal(bytes, &varDeviceIntegrationsMetadataOneOf2)
	if err == nil {
		*o = DeviceIntegrationsMetadataOneOf2(varDeviceIntegrationsMetadataOneOf2)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "idpId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceIntegrationsMetadataOneOf2 struct {
	value *DeviceIntegrationsMetadataOneOf2
	isSet bool
}

func (v NullableDeviceIntegrationsMetadataOneOf2) Get() *DeviceIntegrationsMetadataOneOf2 {
	return v.value
}

func (v *NullableDeviceIntegrationsMetadataOneOf2) Set(val *DeviceIntegrationsMetadataOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIntegrationsMetadataOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIntegrationsMetadataOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIntegrationsMetadataOneOf2(val *DeviceIntegrationsMetadataOneOf2) *NullableDeviceIntegrationsMetadataOneOf2 {
	return &NullableDeviceIntegrationsMetadataOneOf2{value: val, isSet: true}
}

func (v NullableDeviceIntegrationsMetadataOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIntegrationsMetadataOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

