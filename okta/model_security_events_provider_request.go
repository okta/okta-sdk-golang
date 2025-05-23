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

// SecurityEventsProviderRequest The request schema for creating or updating a Security Events Provider. The `settings` must match one of the schemas.
type SecurityEventsProviderRequest struct {
	// The name of the Security Events Provider instance
	Name string `json:"name"`
	Settings SecurityEventsProviderRequestSettings `json:"settings"`
	// The application type of the Security Events Provider
	Type string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventsProviderRequest SecurityEventsProviderRequest

// NewSecurityEventsProviderRequest instantiates a new SecurityEventsProviderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventsProviderRequest(name string, settings SecurityEventsProviderRequestSettings, type_ string) *SecurityEventsProviderRequest {
	this := SecurityEventsProviderRequest{}
	this.Name = name
	this.Settings = settings
	this.Type = type_
	return &this
}

// NewSecurityEventsProviderRequestWithDefaults instantiates a new SecurityEventsProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventsProviderRequestWithDefaults() *SecurityEventsProviderRequest {
	this := SecurityEventsProviderRequest{}
	return &this
}

// GetName returns the Name field value
func (o *SecurityEventsProviderRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityEventsProviderRequest) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value
func (o *SecurityEventsProviderRequest) GetSettings() SecurityEventsProviderRequestSettings {
	if o == nil {
		var ret SecurityEventsProviderRequestSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderRequest) GetSettingsOk() (*SecurityEventsProviderRequestSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *SecurityEventsProviderRequest) SetSettings(v SecurityEventsProviderRequestSettings) {
	o.Settings = v
}

// GetType returns the Type field value
func (o *SecurityEventsProviderRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityEventsProviderRequest) SetType(v string) {
	o.Type = v
}

func (o SecurityEventsProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["settings"] = o.Settings
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventsProviderRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventsProviderRequest := _SecurityEventsProviderRequest{}

	err = json.Unmarshal(bytes, &varSecurityEventsProviderRequest)
	if err == nil {
		*o = SecurityEventsProviderRequest(varSecurityEventsProviderRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventsProviderRequest struct {
	value *SecurityEventsProviderRequest
	isSet bool
}

func (v NullableSecurityEventsProviderRequest) Get() *SecurityEventsProviderRequest {
	return v.value
}

func (v *NullableSecurityEventsProviderRequest) Set(val *SecurityEventsProviderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventsProviderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventsProviderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventsProviderRequest(val *SecurityEventsProviderRequest) *NullableSecurityEventsProviderRequest {
	return &NullableSecurityEventsProviderRequest{value: val, isSet: true}
}

func (v NullableSecurityEventsProviderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventsProviderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

