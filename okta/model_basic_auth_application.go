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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// BasicAuthApplication struct for BasicAuthApplication
type BasicAuthApplication struct {
	Application
	Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
	// Unique key for the app definition
	Name *string `json:"name,omitempty"`
	Settings *BasicApplicationSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BasicAuthApplication BasicAuthApplication

// NewBasicAuthApplication instantiates a new BasicAuthApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicAuthApplication() *BasicAuthApplication {
	this := BasicAuthApplication{}
	var name string = "template_basic_auth"
	this.Name = &name
	return &this
}

// NewBasicAuthApplicationWithDefaults instantiates a new BasicAuthApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicAuthApplicationWithDefaults() *BasicAuthApplication {
	this := BasicAuthApplication{}
	var name string = "template_basic_auth"
	this.Name = &name
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *BasicAuthApplication) GetCredentials() SchemeApplicationCredentials {
	if o == nil || o.Credentials == nil {
		var ret SchemeApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *BasicAuthApplication) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given SchemeApplicationCredentials and assigns it to the Credentials field.
func (o *BasicAuthApplication) SetCredentials(v SchemeApplicationCredentials) {
	o.Credentials = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BasicAuthApplication) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthApplication) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BasicAuthApplication) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BasicAuthApplication) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *BasicAuthApplication) GetSettings() BasicApplicationSettings {
	if o == nil || o.Settings == nil {
		var ret BasicApplicationSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthApplication) GetSettingsOk() (*BasicApplicationSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *BasicAuthApplication) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given BasicApplicationSettings and assigns it to the Settings field.
func (o *BasicAuthApplication) SetSettings(v BasicApplicationSettings) {
	o.Settings = &v
}

func (o BasicAuthApplication) MarshalJSON() ([]byte, error) {
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

func (o *BasicAuthApplication) UnmarshalJSON(bytes []byte) (err error) {
	type BasicAuthApplicationWithoutEmbeddedStruct struct {
		Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
		// Unique key for the app definition
		Name *string `json:"name,omitempty"`
		Settings *BasicApplicationSettings `json:"settings,omitempty"`
	}

	varBasicAuthApplicationWithoutEmbeddedStruct := BasicAuthApplicationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varBasicAuthApplicationWithoutEmbeddedStruct)
	if err == nil {
		varBasicAuthApplication := _BasicAuthApplication{}
		varBasicAuthApplication.Credentials = varBasicAuthApplicationWithoutEmbeddedStruct.Credentials
		varBasicAuthApplication.Name = varBasicAuthApplicationWithoutEmbeddedStruct.Name
		varBasicAuthApplication.Settings = varBasicAuthApplicationWithoutEmbeddedStruct.Settings
		*o = BasicAuthApplication(varBasicAuthApplication)
	} else {
		return err
	}

	varBasicAuthApplication := _BasicAuthApplication{}

	err = json.Unmarshal(bytes, &varBasicAuthApplication)
	if err == nil {
		o.Application = varBasicAuthApplication.Application
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

type NullableBasicAuthApplication struct {
	value *BasicAuthApplication
	isSet bool
}

func (v NullableBasicAuthApplication) Get() *BasicAuthApplication {
	return v.value
}

func (v *NullableBasicAuthApplication) Set(val *BasicAuthApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicAuthApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicAuthApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicAuthApplication(val *BasicAuthApplication) *NullableBasicAuthApplication {
	return &NullableBasicAuthApplication{value: val, isSet: true}
}

func (v NullableBasicAuthApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicAuthApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

