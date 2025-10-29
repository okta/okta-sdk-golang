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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the DeviceIntegrationsMetadataOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceIntegrationsMetadataOneOf{}

// DeviceIntegrationsMetadataOneOf struct for DeviceIntegrationsMetadataOneOf
type DeviceIntegrationsMetadataOneOf struct {
	Type                 string `json:"type"`
	ServiceAccountName   string `json:"serviceAccountName"`
	ServiceAccountEmail  string `json:"serviceAccountEmail"`
	AdditionalProperties map[string]interface{}
}

type _DeviceIntegrationsMetadataOneOf DeviceIntegrationsMetadataOneOf

// NewDeviceIntegrationsMetadataOneOf instantiates a new DeviceIntegrationsMetadataOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceIntegrationsMetadataOneOf(type_ string, serviceAccountName string, serviceAccountEmail string) *DeviceIntegrationsMetadataOneOf {
	this := DeviceIntegrationsMetadataOneOf{}
	this.Type = type_
	this.ServiceAccountName = serviceAccountName
	this.ServiceAccountEmail = serviceAccountEmail
	return &this
}

// NewDeviceIntegrationsMetadataOneOfWithDefaults instantiates a new DeviceIntegrationsMetadataOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIntegrationsMetadataOneOfWithDefaults() *DeviceIntegrationsMetadataOneOf {
	this := DeviceIntegrationsMetadataOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *DeviceIntegrationsMetadataOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeviceIntegrationsMetadataOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeviceIntegrationsMetadataOneOf) SetType(v string) {
	o.Type = v
}

// GetServiceAccountName returns the ServiceAccountName field value
func (o *DeviceIntegrationsMetadataOneOf) GetServiceAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAccountName
}

// GetServiceAccountNameOk returns a tuple with the ServiceAccountName field value
// and a boolean to check if the value has been set.
func (o *DeviceIntegrationsMetadataOneOf) GetServiceAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccountName, true
}

// SetServiceAccountName sets field value
func (o *DeviceIntegrationsMetadataOneOf) SetServiceAccountName(v string) {
	o.ServiceAccountName = v
}

// GetServiceAccountEmail returns the ServiceAccountEmail field value
func (o *DeviceIntegrationsMetadataOneOf) GetServiceAccountEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAccountEmail
}

// GetServiceAccountEmailOk returns a tuple with the ServiceAccountEmail field value
// and a boolean to check if the value has been set.
func (o *DeviceIntegrationsMetadataOneOf) GetServiceAccountEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccountEmail, true
}

// SetServiceAccountEmail sets field value
func (o *DeviceIntegrationsMetadataOneOf) SetServiceAccountEmail(v string) {
	o.ServiceAccountEmail = v
}

func (o DeviceIntegrationsMetadataOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceIntegrationsMetadataOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["serviceAccountName"] = o.ServiceAccountName
	toSerialize["serviceAccountEmail"] = o.ServiceAccountEmail

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceIntegrationsMetadataOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"serviceAccountName",
		"serviceAccountEmail",
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

	varDeviceIntegrationsMetadataOneOf := _DeviceIntegrationsMetadataOneOf{}

	err = json.Unmarshal(data, &varDeviceIntegrationsMetadataOneOf)

	if err != nil {
		return err
	}

	*o = DeviceIntegrationsMetadataOneOf(varDeviceIntegrationsMetadataOneOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "serviceAccountName")
		delete(additionalProperties, "serviceAccountEmail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceIntegrationsMetadataOneOf struct {
	value *DeviceIntegrationsMetadataOneOf
	isSet bool
}

func (v NullableDeviceIntegrationsMetadataOneOf) Get() *DeviceIntegrationsMetadataOneOf {
	return v.value
}

func (v *NullableDeviceIntegrationsMetadataOneOf) Set(val *DeviceIntegrationsMetadataOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIntegrationsMetadataOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIntegrationsMetadataOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIntegrationsMetadataOneOf(val *DeviceIntegrationsMetadataOneOf) *NullableDeviceIntegrationsMetadataOneOf {
	return &NullableDeviceIntegrationsMetadataOneOf{value: val, isSet: true}
}

func (v NullableDeviceIntegrationsMetadataOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIntegrationsMetadataOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
