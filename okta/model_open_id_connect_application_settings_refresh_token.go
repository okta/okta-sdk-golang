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

// OpenIdConnectApplicationSettingsRefreshToken Refresh token configuration for an OAuth 2.0 client  When you create or update an OAuth 2.0 client, you can configure refresh token rotation by setting the `rotation_type` and `leeway` properties. If you don't set these properties when you create an app integration, the default values are used. When you update an app integration, your previously configured values are used. 
type OpenIdConnectApplicationSettingsRefreshToken struct {
	// The leeway, in seconds, allowed for the OAuth 2.0 client. After the refresh token is rotated, the previous token remains valid for the specified period of time so clients can get the new token.  > **Note:** A leeway of 0 doesn't necessarily mean that the previous token is immediately invalidated. The previous token is invalidated after the new token is generated and returned in the response. 
	Leeway *int32 `json:"leeway,omitempty"`
	// The refresh token rotation mode for the OAuth 2.0 client
	RotationType string `json:"rotation_type"`
	AdditionalProperties map[string]interface{}
}

type _OpenIdConnectApplicationSettingsRefreshToken OpenIdConnectApplicationSettingsRefreshToken

// NewOpenIdConnectApplicationSettingsRefreshToken instantiates a new OpenIdConnectApplicationSettingsRefreshToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectApplicationSettingsRefreshToken(rotationType string) *OpenIdConnectApplicationSettingsRefreshToken {
	this := OpenIdConnectApplicationSettingsRefreshToken{}
	var leeway int32 = 30
	this.Leeway = &leeway
	this.RotationType = rotationType
	return &this
}

// NewOpenIdConnectApplicationSettingsRefreshTokenWithDefaults instantiates a new OpenIdConnectApplicationSettingsRefreshToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIdConnectApplicationSettingsRefreshTokenWithDefaults() *OpenIdConnectApplicationSettingsRefreshToken {
	this := OpenIdConnectApplicationSettingsRefreshToken{}
	var leeway int32 = 30
	this.Leeway = &leeway
	return &this
}

// GetLeeway returns the Leeway field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettingsRefreshToken) GetLeeway() int32 {
	if o == nil || o.Leeway == nil {
		var ret int32
		return ret
	}
	return *o.Leeway
}

// GetLeewayOk returns a tuple with the Leeway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsRefreshToken) GetLeewayOk() (*int32, bool) {
	if o == nil || o.Leeway == nil {
		return nil, false
	}
	return o.Leeway, true
}

// HasLeeway returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettingsRefreshToken) HasLeeway() bool {
	if o != nil && o.Leeway != nil {
		return true
	}

	return false
}

// SetLeeway gets a reference to the given int32 and assigns it to the Leeway field.
func (o *OpenIdConnectApplicationSettingsRefreshToken) SetLeeway(v int32) {
	o.Leeway = &v
}

// GetRotationType returns the RotationType field value
func (o *OpenIdConnectApplicationSettingsRefreshToken) GetRotationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RotationType
}

// GetRotationTypeOk returns a tuple with the RotationType field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettingsRefreshToken) GetRotationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RotationType, true
}

// SetRotationType sets field value
func (o *OpenIdConnectApplicationSettingsRefreshToken) SetRotationType(v string) {
	o.RotationType = v
}

func (o OpenIdConnectApplicationSettingsRefreshToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Leeway != nil {
		toSerialize["leeway"] = o.Leeway
	}
	if true {
		toSerialize["rotation_type"] = o.RotationType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenIdConnectApplicationSettingsRefreshToken) UnmarshalJSON(bytes []byte) (err error) {
	varOpenIdConnectApplicationSettingsRefreshToken := _OpenIdConnectApplicationSettingsRefreshToken{}

	err = json.Unmarshal(bytes, &varOpenIdConnectApplicationSettingsRefreshToken)
	if err == nil {
		*o = OpenIdConnectApplicationSettingsRefreshToken(varOpenIdConnectApplicationSettingsRefreshToken)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "leeway")
		delete(additionalProperties, "rotation_type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOpenIdConnectApplicationSettingsRefreshToken struct {
	value *OpenIdConnectApplicationSettingsRefreshToken
	isSet bool
}

func (v NullableOpenIdConnectApplicationSettingsRefreshToken) Get() *OpenIdConnectApplicationSettingsRefreshToken {
	return v.value
}

func (v *NullableOpenIdConnectApplicationSettingsRefreshToken) Set(val *OpenIdConnectApplicationSettingsRefreshToken) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectApplicationSettingsRefreshToken) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectApplicationSettingsRefreshToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectApplicationSettingsRefreshToken(val *OpenIdConnectApplicationSettingsRefreshToken) *NullableOpenIdConnectApplicationSettingsRefreshToken {
	return &NullableOpenIdConnectApplicationSettingsRefreshToken{value: val, isSet: true}
}

func (v NullableOpenIdConnectApplicationSettingsRefreshToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectApplicationSettingsRefreshToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

