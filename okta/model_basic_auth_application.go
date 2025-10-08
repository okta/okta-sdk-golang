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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the BasicAuthApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicAuthApplication{}

// BasicAuthApplication struct for BasicAuthApplication
type BasicAuthApplication struct {
	Application
	Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
	// `template_basic_auth` is the key name for a Basic Authentication scheme app instance
	Name                 string                   `json:"name"`
	Settings             BasicApplicationSettings `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _BasicAuthApplication BasicAuthApplication

// NewBasicAuthApplication instantiates a new BasicAuthApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicAuthApplication(name string, settings BasicApplicationSettings, label string, signOnMode string) *BasicAuthApplication {
	this := BasicAuthApplication{}
	this.Label = label
	this.SignOnMode = signOnMode
	this.Name = name
	this.Settings = settings
	return &this
}

// NewBasicAuthApplicationWithDefaults instantiates a new BasicAuthApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicAuthApplicationWithDefaults() *BasicAuthApplication {
	this := BasicAuthApplication{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *BasicAuthApplication) GetCredentials() SchemeApplicationCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret SchemeApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *BasicAuthApplication) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given SchemeApplicationCredentials and assigns it to the Credentials field.
func (o *BasicAuthApplication) SetCredentials(v SchemeApplicationCredentials) {
	o.Credentials = &v
}

// GetName returns the Name field value
func (o *BasicAuthApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BasicAuthApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BasicAuthApplication) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value
func (o *BasicAuthApplication) GetSettings() BasicApplicationSettings {
	if o == nil {
		var ret BasicApplicationSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *BasicAuthApplication) GetSettingsOk() (*BasicApplicationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *BasicAuthApplication) SetSettings(v BasicApplicationSettings) {
	o.Settings = v
}

func (o BasicAuthApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicAuthApplication) ToMap() (map[string]interface{}, error) {
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

func (o *BasicAuthApplication) UnmarshalJSON(data []byte) (err error) {
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

	type BasicAuthApplicationWithoutEmbeddedStruct struct {
		Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
		// `template_basic_auth` is the key name for a Basic Authentication scheme app instance
		Name     string                   `json:"name"`
		Settings BasicApplicationSettings `json:"settings"`
	}

	varBasicAuthApplicationWithoutEmbeddedStruct := BasicAuthApplicationWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varBasicAuthApplicationWithoutEmbeddedStruct)
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

	err = json.Unmarshal(data, &varBasicAuthApplication)
	if err == nil {
		o.Application = varBasicAuthApplication.Application
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
