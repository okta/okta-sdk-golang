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
	"reflect"
	"strings"
)

// WsFederationApplication struct for WsFederationApplication
type WsFederationApplication struct {
	Application
	Credentials *ApplicationCredentials `json:"credentials,omitempty"`
	// `template_wsfed` is the key name for a WS-Federated app instance with a SAML 2.0 token
	Name string `json:"name"`
	Settings WsFederationApplicationSettings `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _WsFederationApplication WsFederationApplication

// NewWsFederationApplication instantiates a new WsFederationApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWsFederationApplication(name string, settings WsFederationApplicationSettings, label string, signOnMode string) *WsFederationApplication {
	this := WsFederationApplication{}
	this.Label = label
	this.SignOnMode = signOnMode
	this.Name = name
	this.Settings = settings
	return &this
}

// NewWsFederationApplicationWithDefaults instantiates a new WsFederationApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWsFederationApplicationWithDefaults() *WsFederationApplication {
	this := WsFederationApplication{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *WsFederationApplication) GetCredentials() ApplicationCredentials {
	if o == nil || o.Credentials == nil {
		var ret ApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplication) GetCredentialsOk() (*ApplicationCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *WsFederationApplication) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given ApplicationCredentials and assigns it to the Credentials field.
func (o *WsFederationApplication) SetCredentials(v ApplicationCredentials) {
	o.Credentials = &v
}

// GetName returns the Name field value
func (o *WsFederationApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WsFederationApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WsFederationApplication) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value
func (o *WsFederationApplication) GetSettings() WsFederationApplicationSettings {
	if o == nil {
		var ret WsFederationApplicationSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *WsFederationApplication) GetSettingsOk() (*WsFederationApplicationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *WsFederationApplication) SetSettings(v WsFederationApplicationSettings) {
	o.Settings = v
}

func (o WsFederationApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedApplication, errApplication := json.Marshal(o.Application)
	if errApplication != nil {
		return []byte{}, errApplication
	}
	errApplication = json.Unmarshal([]byte(serializedApplication), &toSerialize)
	if errApplication != nil {
		return []byte{}, errApplication
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WsFederationApplication) UnmarshalJSON(bytes []byte) (err error) {
	type WsFederationApplicationWithoutEmbeddedStruct struct {
		Credentials *ApplicationCredentials `json:"credentials,omitempty"`
		// `template_wsfed` is the key name for a WS-Federated app instance with a SAML 2.0 token
		Name string `json:"name"`
		Settings WsFederationApplicationSettings `json:"settings"`
	}

	varWsFederationApplicationWithoutEmbeddedStruct := WsFederationApplicationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWsFederationApplicationWithoutEmbeddedStruct)
	if err == nil {
		varWsFederationApplication := _WsFederationApplication{}
		varWsFederationApplication.Credentials = varWsFederationApplicationWithoutEmbeddedStruct.Credentials
		varWsFederationApplication.Name = varWsFederationApplicationWithoutEmbeddedStruct.Name
		varWsFederationApplication.Settings = varWsFederationApplicationWithoutEmbeddedStruct.Settings
		*o = WsFederationApplication(varWsFederationApplication)
	} else {
		return err
	}

	varWsFederationApplication := _WsFederationApplication{}

	err = json.Unmarshal(bytes, &varWsFederationApplication)
	if err == nil {
		o.Application = varWsFederationApplication.Application
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "name")
		delete(additionalProperties, "settings")

		// remove fields from embedded structs
		reflectApplication := reflect.ValueOf(o.Application)
		for i := 0; i < reflectApplication.Type().NumField(); i++ {
			t := reflectApplication.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWsFederationApplication struct {
	value *WsFederationApplication
	isSet bool
}

func (v NullableWsFederationApplication) Get() *WsFederationApplication {
	return v.value
}

func (v *NullableWsFederationApplication) Set(val *WsFederationApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableWsFederationApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableWsFederationApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWsFederationApplication(val *WsFederationApplication) *NullableWsFederationApplication {
	return &NullableWsFederationApplication{value: val, isSet: true}
}

func (v NullableWsFederationApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWsFederationApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

