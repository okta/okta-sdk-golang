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

// OpenIdConnectApplication struct for OpenIdConnectApplication
type OpenIdConnectApplication struct {
	Application
	Credentials OAuthApplicationCredentials `json:"credentials"`
	// `oidc_client` is the key name for an OIDC app instance
	Name string `json:"name"`
	Settings OpenIdConnectApplicationSettings `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _OpenIdConnectApplication OpenIdConnectApplication

// NewOpenIdConnectApplication instantiates a new OpenIdConnectApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectApplication(credentials OAuthApplicationCredentials, name string, settings OpenIdConnectApplicationSettings, label string, signOnMode string) *OpenIdConnectApplication {
	this := OpenIdConnectApplication{}
	this.Label = label
	this.SignOnMode = signOnMode
	this.Credentials = credentials
	this.Name = name
	this.Settings = settings
	return &this
}

// NewOpenIdConnectApplicationWithDefaults instantiates a new OpenIdConnectApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIdConnectApplicationWithDefaults() *OpenIdConnectApplication {
	this := OpenIdConnectApplication{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *OpenIdConnectApplication) GetCredentials() OAuthApplicationCredentials {
	if o == nil {
		var ret OAuthApplicationCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplication) GetCredentialsOk() (*OAuthApplicationCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *OpenIdConnectApplication) SetCredentials(v OAuthApplicationCredentials) {
	o.Credentials = v
}

// GetName returns the Name field value
func (o *OpenIdConnectApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpenIdConnectApplication) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value
func (o *OpenIdConnectApplication) GetSettings() OpenIdConnectApplicationSettings {
	if o == nil {
		var ret OpenIdConnectApplicationSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplication) GetSettingsOk() (*OpenIdConnectApplicationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *OpenIdConnectApplication) SetSettings(v OpenIdConnectApplicationSettings) {
	o.Settings = v
}

func (o OpenIdConnectApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedApplication, errApplication := json.Marshal(o.Application)
	if errApplication != nil {
		return []byte{}, errApplication
	}
	errApplication = json.Unmarshal([]byte(serializedApplication), &toSerialize)
	if errApplication != nil {
		return []byte{}, errApplication
	}
	if true {
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

func (o *OpenIdConnectApplication) UnmarshalJSON(bytes []byte) (err error) {
	type OpenIdConnectApplicationWithoutEmbeddedStruct struct {
		Credentials OAuthApplicationCredentials `json:"credentials"`
		// `oidc_client` is the key name for an OIDC app instance
		Name string `json:"name"`
		Settings OpenIdConnectApplicationSettings `json:"settings"`
	}

	varOpenIdConnectApplicationWithoutEmbeddedStruct := OpenIdConnectApplicationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varOpenIdConnectApplicationWithoutEmbeddedStruct)
	if err == nil {
		varOpenIdConnectApplication := _OpenIdConnectApplication{}
		varOpenIdConnectApplication.Credentials = varOpenIdConnectApplicationWithoutEmbeddedStruct.Credentials
		varOpenIdConnectApplication.Name = varOpenIdConnectApplicationWithoutEmbeddedStruct.Name
		varOpenIdConnectApplication.Settings = varOpenIdConnectApplicationWithoutEmbeddedStruct.Settings
		*o = OpenIdConnectApplication(varOpenIdConnectApplication)
	} else {
		return err
	}

	varOpenIdConnectApplication := _OpenIdConnectApplication{}

	err = json.Unmarshal(bytes, &varOpenIdConnectApplication)
	if err == nil {
		o.Application = varOpenIdConnectApplication.Application
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

type NullableOpenIdConnectApplication struct {
	value *OpenIdConnectApplication
	isSet bool
}

func (v NullableOpenIdConnectApplication) Get() *OpenIdConnectApplication {
	return v.value
}

func (v *NullableOpenIdConnectApplication) Set(val *OpenIdConnectApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectApplication(val *OpenIdConnectApplication) *NullableOpenIdConnectApplication {
	return &NullableOpenIdConnectApplication{value: val, isSet: true}
}

func (v NullableOpenIdConnectApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

