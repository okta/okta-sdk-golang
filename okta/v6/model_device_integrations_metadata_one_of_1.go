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
	"fmt"
)

// checks if the DeviceIntegrationsMetadataOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceIntegrationsMetadataOneOf1{}

// DeviceIntegrationsMetadataOneOf1 struct for DeviceIntegrationsMetadataOneOf1
type DeviceIntegrationsMetadataOneOf1 struct {
	Type                 string `json:"type"`
	Provider             string `json:"provider"`
	EnrollmentUrl        string `json:"enrollmentUrl"`
	IdpId                string `json:"idpId"`
	AdditionalProperties map[string]interface{}
}

type _DeviceIntegrationsMetadataOneOf1 DeviceIntegrationsMetadataOneOf1

// NewDeviceIntegrationsMetadataOneOf1 instantiates a new DeviceIntegrationsMetadataOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceIntegrationsMetadataOneOf1(type_ string, provider string, enrollmentUrl string, idpId string) *DeviceIntegrationsMetadataOneOf1 {
	this := DeviceIntegrationsMetadataOneOf1{}
	this.Type = type_
	this.Provider = provider
	this.EnrollmentUrl = enrollmentUrl
	this.IdpId = idpId
	return &this
}

// NewDeviceIntegrationsMetadataOneOf1WithDefaults instantiates a new DeviceIntegrationsMetadataOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIntegrationsMetadataOneOf1WithDefaults() *DeviceIntegrationsMetadataOneOf1 {
	this := DeviceIntegrationsMetadataOneOf1{}
	return &this
}

// GetType returns the Type field value
func (o *DeviceIntegrationsMetadataOneOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeviceIntegrationsMetadataOneOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeviceIntegrationsMetadataOneOf1) SetType(v string) {
	o.Type = v
}

// GetProvider returns the Provider field value
func (o *DeviceIntegrationsMetadataOneOf1) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *DeviceIntegrationsMetadataOneOf1) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *DeviceIntegrationsMetadataOneOf1) SetProvider(v string) {
	o.Provider = v
}

// GetEnrollmentUrl returns the EnrollmentUrl field value
func (o *DeviceIntegrationsMetadataOneOf1) GetEnrollmentUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnrollmentUrl
}

// GetEnrollmentUrlOk returns a tuple with the EnrollmentUrl field value
// and a boolean to check if the value has been set.
func (o *DeviceIntegrationsMetadataOneOf1) GetEnrollmentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentUrl, true
}

// SetEnrollmentUrl sets field value
func (o *DeviceIntegrationsMetadataOneOf1) SetEnrollmentUrl(v string) {
	o.EnrollmentUrl = v
}

// GetIdpId returns the IdpId field value
func (o *DeviceIntegrationsMetadataOneOf1) GetIdpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpId
}

// GetIdpIdOk returns a tuple with the IdpId field value
// and a boolean to check if the value has been set.
func (o *DeviceIntegrationsMetadataOneOf1) GetIdpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpId, true
}

// SetIdpId sets field value
func (o *DeviceIntegrationsMetadataOneOf1) SetIdpId(v string) {
	o.IdpId = v
}

func (o DeviceIntegrationsMetadataOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceIntegrationsMetadataOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["provider"] = o.Provider
	toSerialize["enrollmentUrl"] = o.EnrollmentUrl
	toSerialize["idpId"] = o.IdpId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceIntegrationsMetadataOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"provider",
		"enrollmentUrl",
		"idpId",
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

	varDeviceIntegrationsMetadataOneOf1 := _DeviceIntegrationsMetadataOneOf1{}

	err = json.Unmarshal(data, &varDeviceIntegrationsMetadataOneOf1)

	if err != nil {
		return err
	}

	*o = DeviceIntegrationsMetadataOneOf1(varDeviceIntegrationsMetadataOneOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "enrollmentUrl")
		delete(additionalProperties, "idpId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceIntegrationsMetadataOneOf1 struct {
	value *DeviceIntegrationsMetadataOneOf1
	isSet bool
}

func (v NullableDeviceIntegrationsMetadataOneOf1) Get() *DeviceIntegrationsMetadataOneOf1 {
	return v.value
}

func (v *NullableDeviceIntegrationsMetadataOneOf1) Set(val *DeviceIntegrationsMetadataOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIntegrationsMetadataOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIntegrationsMetadataOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIntegrationsMetadataOneOf1(val *DeviceIntegrationsMetadataOneOf1) *NullableDeviceIntegrationsMetadataOneOf1 {
	return &NullableDeviceIntegrationsMetadataOneOf1{value: val, isSet: true}
}

func (v NullableDeviceIntegrationsMetadataOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIntegrationsMetadataOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
