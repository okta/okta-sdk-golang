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
)

// checks if the BasicApplicationSettingsApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicApplicationSettingsApplication{}

// BasicApplicationSettingsApplication struct for BasicApplicationSettingsApplication
type BasicApplicationSettingsApplication struct {
	// The URL of the authenticating site for this app
	AuthURL string `json:"authURL"`
	// The URL of the sign-in page for this app
	Url                  string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _BasicApplicationSettingsApplication BasicApplicationSettingsApplication

// NewBasicApplicationSettingsApplication instantiates a new BasicApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicApplicationSettingsApplication(authURL string, url string) *BasicApplicationSettingsApplication {
	this := BasicApplicationSettingsApplication{}
	this.AuthURL = authURL
	this.Url = url
	return &this
}

// NewBasicApplicationSettingsApplicationWithDefaults instantiates a new BasicApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicApplicationSettingsApplicationWithDefaults() *BasicApplicationSettingsApplication {
	this := BasicApplicationSettingsApplication{}
	return &this
}

// GetAuthURL returns the AuthURL field value
func (o *BasicApplicationSettingsApplication) GetAuthURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthURL
}

// GetAuthURLOk returns a tuple with the AuthURL field value
// and a boolean to check if the value has been set.
func (o *BasicApplicationSettingsApplication) GetAuthURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthURL, true
}

// SetAuthURL sets field value
func (o *BasicApplicationSettingsApplication) SetAuthURL(v string) {
	o.AuthURL = v
}

// GetUrl returns the Url field value
func (o *BasicApplicationSettingsApplication) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BasicApplicationSettingsApplication) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BasicApplicationSettingsApplication) SetUrl(v string) {
	o.Url = v
}

func (o BasicApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicApplicationSettingsApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authURL"] = o.AuthURL
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BasicApplicationSettingsApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authURL",
		"url",
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

	varBasicApplicationSettingsApplication := _BasicApplicationSettingsApplication{}

	err = json.Unmarshal(data, &varBasicApplicationSettingsApplication)

	if err != nil {
		return err
	}

	*o = BasicApplicationSettingsApplication(varBasicApplicationSettingsApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authURL")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBasicApplicationSettingsApplication struct {
	value *BasicApplicationSettingsApplication
	isSet bool
}

func (v NullableBasicApplicationSettingsApplication) Get() *BasicApplicationSettingsApplication {
	return v.value
}

func (v *NullableBasicApplicationSettingsApplication) Set(val *BasicApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicApplicationSettingsApplication(val *BasicApplicationSettingsApplication) *NullableBasicApplicationSettingsApplication {
	return &NullableBasicApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableBasicApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
