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

// AutoLoginApplication struct for AutoLoginApplication
type AutoLoginApplication struct {
	Application
	Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
	// A unique key is generated for the custom SWA app instance when you use AUTO_LOGIN `signOnMode`.
	Name *string `json:"name,omitempty"`
	Settings *AutoLoginApplicationSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoLoginApplication AutoLoginApplication

// NewAutoLoginApplication instantiates a new AutoLoginApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoLoginApplication(label string, signOnMode string) *AutoLoginApplication {
	this := AutoLoginApplication{}
	this.Label = label
	this.SignOnMode = signOnMode
	return &this
}

// NewAutoLoginApplicationWithDefaults instantiates a new AutoLoginApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoLoginApplicationWithDefaults() *AutoLoginApplication {
	this := AutoLoginApplication{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *AutoLoginApplication) GetCredentials() SchemeApplicationCredentials {
	if o == nil || o.Credentials == nil {
		var ret SchemeApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *AutoLoginApplication) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given SchemeApplicationCredentials and assigns it to the Credentials field.
func (o *AutoLoginApplication) SetCredentials(v SchemeApplicationCredentials) {
	o.Credentials = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AutoLoginApplication) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplication) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AutoLoginApplication) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AutoLoginApplication) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *AutoLoginApplication) GetSettings() AutoLoginApplicationSettings {
	if o == nil || o.Settings == nil {
		var ret AutoLoginApplicationSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplication) GetSettingsOk() (*AutoLoginApplicationSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *AutoLoginApplication) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given AutoLoginApplicationSettings and assigns it to the Settings field.
func (o *AutoLoginApplication) SetSettings(v AutoLoginApplicationSettings) {
	o.Settings = &v
}

func (o AutoLoginApplication) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AutoLoginApplication) UnmarshalJSON(bytes []byte) (err error) {
	type AutoLoginApplicationWithoutEmbeddedStruct struct {
		Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
		// A unique key is generated for the custom SWA app instance when you use AUTO_LOGIN `signOnMode`.
		Name *string `json:"name,omitempty"`
		Settings *AutoLoginApplicationSettings `json:"settings,omitempty"`
	}

	varAutoLoginApplicationWithoutEmbeddedStruct := AutoLoginApplicationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAutoLoginApplicationWithoutEmbeddedStruct)
	if err == nil {
		varAutoLoginApplication := _AutoLoginApplication{}
		varAutoLoginApplication.Credentials = varAutoLoginApplicationWithoutEmbeddedStruct.Credentials
		varAutoLoginApplication.Name = varAutoLoginApplicationWithoutEmbeddedStruct.Name
		varAutoLoginApplication.Settings = varAutoLoginApplicationWithoutEmbeddedStruct.Settings
		*o = AutoLoginApplication(varAutoLoginApplication)
	} else {
		return err
	}

	varAutoLoginApplication := _AutoLoginApplication{}

	err = json.Unmarshal(bytes, &varAutoLoginApplication)
	if err == nil {
		o.Application = varAutoLoginApplication.Application
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

type NullableAutoLoginApplication struct {
	value *AutoLoginApplication
	isSet bool
}

func (v NullableAutoLoginApplication) Get() *AutoLoginApplication {
	return v.value
}

func (v *NullableAutoLoginApplication) Set(val *AutoLoginApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoLoginApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoLoginApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoLoginApplication(val *AutoLoginApplication) *NullableAutoLoginApplication {
	return &NullableAutoLoginApplication{value: val, isSet: true}
}

func (v NullableAutoLoginApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoLoginApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

