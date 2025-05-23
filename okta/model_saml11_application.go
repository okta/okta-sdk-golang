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

// Saml11Application struct for Saml11Application
type Saml11Application struct {
	Application
	Credentials *ApplicationCredentials `json:"credentials,omitempty"`
	// The key name for the SAML 1.1 app definition. You can't create a custom SAML 1.1 app integration instance. Only existing OIN SAML 1.1 app integrations are supported.
	Name string `json:"name"`
	Settings *Saml11ApplicationSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Saml11Application Saml11Application

// NewSaml11Application instantiates a new Saml11Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaml11Application(name string, label string, signOnMode string) *Saml11Application {
	this := Saml11Application{}
	this.Label = label
	this.SignOnMode = signOnMode
	this.Name = name
	return &this
}

// NewSaml11ApplicationWithDefaults instantiates a new Saml11Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaml11ApplicationWithDefaults() *Saml11Application {
	this := Saml11Application{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *Saml11Application) GetCredentials() ApplicationCredentials {
	if o == nil || o.Credentials == nil {
		var ret ApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11Application) GetCredentialsOk() (*ApplicationCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *Saml11Application) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given ApplicationCredentials and assigns it to the Credentials field.
func (o *Saml11Application) SetCredentials(v ApplicationCredentials) {
	o.Credentials = &v
}

// GetName returns the Name field value
func (o *Saml11Application) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Saml11Application) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Saml11Application) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Saml11Application) GetSettings() Saml11ApplicationSettings {
	if o == nil || o.Settings == nil {
		var ret Saml11ApplicationSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11Application) GetSettingsOk() (*Saml11ApplicationSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Saml11Application) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given Saml11ApplicationSettings and assigns it to the Settings field.
func (o *Saml11Application) SetSettings(v Saml11ApplicationSettings) {
	o.Settings = &v
}

func (o Saml11Application) MarshalJSON() ([]byte, error) {
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
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Saml11Application) UnmarshalJSON(bytes []byte) (err error) {
	type Saml11ApplicationWithoutEmbeddedStruct struct {
		Credentials *ApplicationCredentials `json:"credentials,omitempty"`
		// The key name for the SAML 1.1 app definition. You can't create a custom SAML 1.1 app integration instance. Only existing OIN SAML 1.1 app integrations are supported.
		Name string `json:"name"`
		Settings *Saml11ApplicationSettings `json:"settings,omitempty"`
	}

	varSaml11ApplicationWithoutEmbeddedStruct := Saml11ApplicationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSaml11ApplicationWithoutEmbeddedStruct)
	if err == nil {
		varSaml11Application := _Saml11Application{}
		varSaml11Application.Credentials = varSaml11ApplicationWithoutEmbeddedStruct.Credentials
		varSaml11Application.Name = varSaml11ApplicationWithoutEmbeddedStruct.Name
		varSaml11Application.Settings = varSaml11ApplicationWithoutEmbeddedStruct.Settings
		*o = Saml11Application(varSaml11Application)
	} else {
		return err
	}

	varSaml11Application := _Saml11Application{}

	err = json.Unmarshal(bytes, &varSaml11Application)
	if err == nil {
		o.Application = varSaml11Application.Application
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

type NullableSaml11Application struct {
	value *Saml11Application
	isSet bool
}

func (v NullableSaml11Application) Get() *Saml11Application {
	return v.value
}

func (v *NullableSaml11Application) Set(val *Saml11Application) {
	v.value = val
	v.isSet = true
}

func (v NullableSaml11Application) IsSet() bool {
	return v.isSet
}

func (v *NullableSaml11Application) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaml11Application(val *Saml11Application) *NullableSaml11Application {
	return &NullableSaml11Application{value: val, isSet: true}
}

func (v NullableSaml11Application) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaml11Application) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

