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
)

// checks if the AppAccountContainerLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAccountContainerLink{}

// AppAccountContainerLink struct for AppAccountContainerLink
type AppAccountContainerLink struct {
	Login                *HrefObjectAppLink  `json:"login,omitempty"`
	Logo                 *HrefObjectLogoLink `json:"logo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppAccountContainerLink AppAccountContainerLink

// NewAppAccountContainerLink instantiates a new AppAccountContainerLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAccountContainerLink() *AppAccountContainerLink {
	this := AppAccountContainerLink{}
	return &this
}

// NewAppAccountContainerLinkWithDefaults instantiates a new AppAccountContainerLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAccountContainerLinkWithDefaults() *AppAccountContainerLink {
	this := AppAccountContainerLink{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *AppAccountContainerLink) GetLogin() HrefObjectAppLink {
	if o == nil || IsNil(o.Login) {
		var ret HrefObjectAppLink
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerLink) GetLoginOk() (*HrefObjectAppLink, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *AppAccountContainerLink) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given HrefObjectAppLink and assigns it to the Login field.
func (o *AppAccountContainerLink) SetLogin(v HrefObjectAppLink) {
	o.Login = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *AppAccountContainerLink) GetLogo() HrefObjectLogoLink {
	if o == nil || IsNil(o.Logo) {
		var ret HrefObjectLogoLink
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerLink) GetLogoOk() (*HrefObjectLogoLink, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *AppAccountContainerLink) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given HrefObjectLogoLink and assigns it to the Logo field.
func (o *AppAccountContainerLink) SetLogo(v HrefObjectLogoLink) {
	o.Logo = &v
}

func (o AppAccountContainerLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAccountContainerLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppAccountContainerLink) UnmarshalJSON(data []byte) (err error) {
	varAppAccountContainerLink := _AppAccountContainerLink{}

	err = json.Unmarshal(data, &varAppAccountContainerLink)

	if err != nil {
		return err
	}

	*o = AppAccountContainerLink(varAppAccountContainerLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "login")
		delete(additionalProperties, "logo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppAccountContainerLink struct {
	value *AppAccountContainerLink
	isSet bool
}

func (v NullableAppAccountContainerLink) Get() *AppAccountContainerLink {
	return v.value
}

func (v *NullableAppAccountContainerLink) Set(val *AppAccountContainerLink) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAccountContainerLink) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAccountContainerLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAccountContainerLink(val *AppAccountContainerLink) *NullableAppAccountContainerLink {
	return &NullableAppAccountContainerLink{value: val, isSet: true}
}

func (v NullableAppAccountContainerLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAccountContainerLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
