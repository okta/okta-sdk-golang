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
	"reflect"
	"strings"
)

// checks if the SecurePasswordStoreApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurePasswordStoreApplication{}

// SecurePasswordStoreApplication struct for SecurePasswordStoreApplication
type SecurePasswordStoreApplication struct {
	Application
	Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
	// `template_sps` is the key name for a SWA app instance that uses HTTP POST and doesn't require a browser plugin
	Name                 string                                 `json:"name"`
	Settings             SecurePasswordStoreApplicationSettings `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _SecurePasswordStoreApplication SecurePasswordStoreApplication

// NewSecurePasswordStoreApplication instantiates a new SecurePasswordStoreApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurePasswordStoreApplication(name string, settings SecurePasswordStoreApplicationSettings, label string, signOnMode string) *SecurePasswordStoreApplication {
	this := SecurePasswordStoreApplication{}
	this.Label = label
	this.SignOnMode = signOnMode
	this.Name = name
	this.Settings = settings
	return &this
}

// NewSecurePasswordStoreApplicationWithDefaults instantiates a new SecurePasswordStoreApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurePasswordStoreApplicationWithDefaults() *SecurePasswordStoreApplication {
	this := SecurePasswordStoreApplication{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *SecurePasswordStoreApplication) GetCredentials() SchemeApplicationCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret SchemeApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *SecurePasswordStoreApplication) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given SchemeApplicationCredentials and assigns it to the Credentials field.
func (o *SecurePasswordStoreApplication) SetCredentials(v SchemeApplicationCredentials) {
	o.Credentials = &v
}

// GetName returns the Name field value
func (o *SecurePasswordStoreApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurePasswordStoreApplication) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value
func (o *SecurePasswordStoreApplication) GetSettings() SecurePasswordStoreApplicationSettings {
	if o == nil {
		var ret SecurePasswordStoreApplicationSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *SecurePasswordStoreApplication) GetSettingsOk() (*SecurePasswordStoreApplicationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *SecurePasswordStoreApplication) SetSettings(v SecurePasswordStoreApplicationSettings) {
	o.Settings = v
}

func (o SecurePasswordStoreApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurePasswordStoreApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedApplication, errApplication := json.Marshal(o.Application)
	if errApplication != nil {
		return map[string]interface{}{}, errApplication
	}
	errApplication = json.Unmarshal([]byte(serializedApplication), &toSerialize)
	if errApplication != nil {
		return map[string]interface{}{}, errApplication
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurePasswordStoreApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"settings",
		"label",
		"signOnMode",
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

	type SecurePasswordStoreApplicationWithoutEmbeddedStruct struct {
		Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
		// `template_sps` is the key name for a SWA app instance that uses HTTP POST and doesn't require a browser plugin
		Name     string                                 `json:"name"`
		Settings SecurePasswordStoreApplicationSettings `json:"settings"`
	}

	varSecurePasswordStoreApplicationWithoutEmbeddedStruct := SecurePasswordStoreApplicationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varSecurePasswordStoreApplicationWithoutEmbeddedStruct)
	if err == nil {
		varSecurePasswordStoreApplication := _SecurePasswordStoreApplication{}
		varSecurePasswordStoreApplication.Credentials = varSecurePasswordStoreApplicationWithoutEmbeddedStruct.Credentials
		varSecurePasswordStoreApplication.Name = varSecurePasswordStoreApplicationWithoutEmbeddedStruct.Name
		varSecurePasswordStoreApplication.Settings = varSecurePasswordStoreApplicationWithoutEmbeddedStruct.Settings
		*o = SecurePasswordStoreApplication(varSecurePasswordStoreApplication)
	} else {
		return err
	}

	varSecurePasswordStoreApplication := _SecurePasswordStoreApplication{}

	err = json.Unmarshal(data, &varSecurePasswordStoreApplication)
	if err == nil {
		o.Application = varSecurePasswordStoreApplication.Application
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
	}

	return err
}

type NullableSecurePasswordStoreApplication struct {
	value *SecurePasswordStoreApplication
	isSet bool
}

func (v NullableSecurePasswordStoreApplication) Get() *SecurePasswordStoreApplication {
	return v.value
}

func (v *NullableSecurePasswordStoreApplication) Set(val *SecurePasswordStoreApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurePasswordStoreApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurePasswordStoreApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurePasswordStoreApplication(val *SecurePasswordStoreApplication) *NullableSecurePasswordStoreApplication {
	return &NullableSecurePasswordStoreApplication{value: val, isSet: true}
}

func (v NullableSecurePasswordStoreApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurePasswordStoreApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
