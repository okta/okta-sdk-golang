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

// AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm struct for AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm
type AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm struct {
	// ID of the FCM (Firebase Cloud Messaging Service) [configurations](https://developer.okta.com/docs/reference/api/push-providers/)
	Id *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm

// NewAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm instantiates a new AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm() *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm {
	this := AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm{}
	return &this
}

// NewAuthenticatorKeyCustomAppAllOfProviderConfigurationFcmWithDefaults instantiates a new AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyCustomAppAllOfProviderConfigurationFcmWithDefaults() *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm {
	this := AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) SetId(v string) {
	o.Id = &v
}

func (o AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm := _AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm)
	if err == nil {
		*o = AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm(varAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm struct {
	value *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm
	isSet bool
}

func (v NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) Get() *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm {
	return v.value
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) Set(val *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm(val *AuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) *NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm {
	return &NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationFcm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

