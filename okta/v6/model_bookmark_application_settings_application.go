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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the BookmarkApplicationSettingsApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkApplicationSettingsApplication{}

// BookmarkApplicationSettingsApplication struct for BookmarkApplicationSettingsApplication
type BookmarkApplicationSettingsApplication struct {
	// Would you like Okta to add an integration for this app?
	RequestIntegration *bool `json:"requestIntegration,omitempty"`
	// The URL of the launch page for this app
	Url                  string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _BookmarkApplicationSettingsApplication BookmarkApplicationSettingsApplication

// NewBookmarkApplicationSettingsApplication instantiates a new BookmarkApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkApplicationSettingsApplication(url string) *BookmarkApplicationSettingsApplication {
	this := BookmarkApplicationSettingsApplication{}
	var requestIntegration bool = false
	this.RequestIntegration = &requestIntegration
	this.Url = url
	return &this
}

// NewBookmarkApplicationSettingsApplicationWithDefaults instantiates a new BookmarkApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkApplicationSettingsApplicationWithDefaults() *BookmarkApplicationSettingsApplication {
	this := BookmarkApplicationSettingsApplication{}
	var requestIntegration bool = false
	this.RequestIntegration = &requestIntegration
	return &this
}

// GetRequestIntegration returns the RequestIntegration field value if set, zero value otherwise.
func (o *BookmarkApplicationSettingsApplication) GetRequestIntegration() bool {
	if o == nil || IsNil(o.RequestIntegration) {
		var ret bool
		return ret
	}
	return *o.RequestIntegration
}

// GetRequestIntegrationOk returns a tuple with the RequestIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettingsApplication) GetRequestIntegrationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestIntegration) {
		return nil, false
	}
	return o.RequestIntegration, true
}

// HasRequestIntegration returns a boolean if a field has been set.
func (o *BookmarkApplicationSettingsApplication) HasRequestIntegration() bool {
	if o != nil && !IsNil(o.RequestIntegration) {
		return true
	}

	return false
}

// SetRequestIntegration gets a reference to the given bool and assigns it to the RequestIntegration field.
func (o *BookmarkApplicationSettingsApplication) SetRequestIntegration(v bool) {
	o.RequestIntegration = &v
}

// GetUrl returns the Url field value
func (o *BookmarkApplicationSettingsApplication) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettingsApplication) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BookmarkApplicationSettingsApplication) SetUrl(v string) {
	o.Url = v
}

func (o BookmarkApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkApplicationSettingsApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestIntegration) {
		toSerialize["requestIntegration"] = o.RequestIntegration
	}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BookmarkApplicationSettingsApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varBookmarkApplicationSettingsApplication := _BookmarkApplicationSettingsApplication{}

	err = json.Unmarshal(data, &varBookmarkApplicationSettingsApplication)

	if err != nil {
		return err
	}

	*o = BookmarkApplicationSettingsApplication(varBookmarkApplicationSettingsApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "requestIntegration")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBookmarkApplicationSettingsApplication struct {
	value *BookmarkApplicationSettingsApplication
	isSet bool
}

func (v NullableBookmarkApplicationSettingsApplication) Get() *BookmarkApplicationSettingsApplication {
	return v.value
}

func (v *NullableBookmarkApplicationSettingsApplication) Set(val *BookmarkApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkApplicationSettingsApplication(val *BookmarkApplicationSettingsApplication) *NullableBookmarkApplicationSettingsApplication {
	return &NullableBookmarkApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableBookmarkApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
