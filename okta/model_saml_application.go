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

// SamlApplication struct for SamlApplication
type SamlApplication struct {
	Application
	Credentials *ApplicationCredentials `json:"credentials,omitempty"`
	// A unique key is generated for the custom app instance when you use SAML_2_0 `signOnMode`.
	Name *string `json:"name,omitempty"`
	Settings *SamlApplicationSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlApplication SamlApplication

// NewSamlApplication instantiates a new SamlApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlApplication(label string, signOnMode string) *SamlApplication {
	this := SamlApplication{}
	this.Label = label
	this.SignOnMode = signOnMode
	return &this
}

// NewSamlApplicationWithDefaults instantiates a new SamlApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlApplicationWithDefaults() *SamlApplication {
	this := SamlApplication{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *SamlApplication) GetCredentials() ApplicationCredentials {
	if o == nil || o.Credentials == nil {
		var ret ApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplication) GetCredentialsOk() (*ApplicationCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *SamlApplication) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given ApplicationCredentials and assigns it to the Credentials field.
func (o *SamlApplication) SetCredentials(v ApplicationCredentials) {
	o.Credentials = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SamlApplication) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplication) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SamlApplication) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SamlApplication) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SamlApplication) GetSettings() SamlApplicationSettings {
	if o == nil || o.Settings == nil {
		var ret SamlApplicationSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplication) GetSettingsOk() (*SamlApplicationSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SamlApplication) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SamlApplicationSettings and assigns it to the Settings field.
func (o *SamlApplication) SetSettings(v SamlApplicationSettings) {
	o.Settings = &v
}

func (o SamlApplication) MarshalJSON() ([]byte, error) {
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

func (o *SamlApplication) UnmarshalJSON(bytes []byte) (err error) {
	type SamlApplicationWithoutEmbeddedStruct struct {
		Credentials *ApplicationCredentials `json:"credentials,omitempty"`
		// A unique key is generated for the custom app instance when you use SAML_2_0 `signOnMode`.
		Name *string `json:"name,omitempty"`
		Settings *SamlApplicationSettings `json:"settings,omitempty"`
	}

	varSamlApplicationWithoutEmbeddedStruct := SamlApplicationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSamlApplicationWithoutEmbeddedStruct)
	if err == nil {
		varSamlApplication := _SamlApplication{}
		varSamlApplication.Credentials = varSamlApplicationWithoutEmbeddedStruct.Credentials
		varSamlApplication.Name = varSamlApplicationWithoutEmbeddedStruct.Name
		varSamlApplication.Settings = varSamlApplicationWithoutEmbeddedStruct.Settings
		*o = SamlApplication(varSamlApplication)
	} else {
		return err
	}

	varSamlApplication := _SamlApplication{}

	err = json.Unmarshal(bytes, &varSamlApplication)
	if err == nil {
		o.Application = varSamlApplication.Application
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

type NullableSamlApplication struct {
	value *SamlApplication
	isSet bool
}

func (v NullableSamlApplication) Get() *SamlApplication {
	return v.value
}

func (v *NullableSamlApplication) Set(val *SamlApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlApplication(val *SamlApplication) *NullableSamlApplication {
	return &NullableSamlApplication{value: val, isSet: true}
}

func (v NullableSamlApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

