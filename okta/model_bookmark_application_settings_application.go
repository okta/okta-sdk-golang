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
)

// BookmarkApplicationSettingsApplication struct for BookmarkApplicationSettingsApplication
type BookmarkApplicationSettingsApplication struct {
	RequestIntegration *bool `json:"requestIntegration,omitempty"`
	Url *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BookmarkApplicationSettingsApplication BookmarkApplicationSettingsApplication

// NewBookmarkApplicationSettingsApplication instantiates a new BookmarkApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkApplicationSettingsApplication() *BookmarkApplicationSettingsApplication {
	this := BookmarkApplicationSettingsApplication{}
	return &this
}

// NewBookmarkApplicationSettingsApplicationWithDefaults instantiates a new BookmarkApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkApplicationSettingsApplicationWithDefaults() *BookmarkApplicationSettingsApplication {
	this := BookmarkApplicationSettingsApplication{}
	return &this
}

// GetRequestIntegration returns the RequestIntegration field value if set, zero value otherwise.
func (o *BookmarkApplicationSettingsApplication) GetRequestIntegration() bool {
	if o == nil || o.RequestIntegration == nil {
		var ret bool
		return ret
	}
	return *o.RequestIntegration
}

// GetRequestIntegrationOk returns a tuple with the RequestIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettingsApplication) GetRequestIntegrationOk() (*bool, bool) {
	if o == nil || o.RequestIntegration == nil {
		return nil, false
	}
	return o.RequestIntegration, true
}

// HasRequestIntegration returns a boolean if a field has been set.
func (o *BookmarkApplicationSettingsApplication) HasRequestIntegration() bool {
	if o != nil && o.RequestIntegration != nil {
		return true
	}

	return false
}

// SetRequestIntegration gets a reference to the given bool and assigns it to the RequestIntegration field.
func (o *BookmarkApplicationSettingsApplication) SetRequestIntegration(v bool) {
	o.RequestIntegration = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BookmarkApplicationSettingsApplication) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettingsApplication) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BookmarkApplicationSettingsApplication) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BookmarkApplicationSettingsApplication) SetUrl(v string) {
	o.Url = &v
}

func (o BookmarkApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestIntegration != nil {
		toSerialize["requestIntegration"] = o.RequestIntegration
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BookmarkApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varBookmarkApplicationSettingsApplication := _BookmarkApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varBookmarkApplicationSettingsApplication)
	if err == nil {
		*o = BookmarkApplicationSettingsApplication(varBookmarkApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestIntegration")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

