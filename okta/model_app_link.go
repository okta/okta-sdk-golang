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

// AppLink struct for AppLink
type AppLink struct {
	Login *HrefObjectAppLink `json:"login,omitempty"`
	Logo *HrefObjectLogoLink `json:"logo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppLink AppLink

// NewAppLink instantiates a new AppLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppLink() *AppLink {
	this := AppLink{}
	return &this
}

// NewAppLinkWithDefaults instantiates a new AppLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppLinkWithDefaults() *AppLink {
	this := AppLink{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *AppLink) GetLogin() HrefObjectAppLink {
	if o == nil || o.Login == nil {
		var ret HrefObjectAppLink
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetLoginOk() (*HrefObjectAppLink, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *AppLink) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given HrefObjectAppLink and assigns it to the Login field.
func (o *AppLink) SetLogin(v HrefObjectAppLink) {
	o.Login = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *AppLink) GetLogo() HrefObjectLogoLink {
	if o == nil || o.Logo == nil {
		var ret HrefObjectLogoLink
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppLink) GetLogoOk() (*HrefObjectLogoLink, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *AppLink) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given HrefObjectLogoLink and assigns it to the Logo field.
func (o *AppLink) SetLogo(v HrefObjectLogoLink) {
	o.Logo = &v
}

func (o AppLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppLink) UnmarshalJSON(bytes []byte) (err error) {
	varAppLink := _AppLink{}

	err = json.Unmarshal(bytes, &varAppLink)
	if err == nil {
		*o = AppLink(varAppLink)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "login")
		delete(additionalProperties, "logo")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppLink struct {
	value *AppLink
	isSet bool
}

func (v NullableAppLink) Get() *AppLink {
	return v.value
}

func (v *NullableAppLink) Set(val *AppLink) {
	v.value = val
	v.isSet = true
}

func (v NullableAppLink) IsSet() bool {
	return v.isSet
}

func (v *NullableAppLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppLink(val *AppLink) *NullableAppLink {
	return &NullableAppLink{value: val, isSet: true}
}

func (v NullableAppLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

