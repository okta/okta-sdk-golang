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
	"time"
)

// AuthorizationServerCredentialsSigningConfig struct for AuthorizationServerCredentialsSigningConfig
type AuthorizationServerCredentialsSigningConfig struct {
	// The ID of the JSON Web Key used for signing tokens issued by the authorization server
	Kid *string `json:"kid,omitempty"`
	// The timestamp when the authorization server started using the `kid` for signing tokens
	LastRotated *time.Time `json:"lastRotated,omitempty"`
	// The timestamp when the authorization server changes the Key for signing tokens. This is only returned when `rotationMode` is set to `AUTO`.
	NextRotation *time.Time `json:"nextRotation,omitempty"`
	// The Key rotation mode for the authorization server
	RotationMode *string `json:"rotationMode,omitempty"`
	// How the key is used
	Use *string `json:"use,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerCredentialsSigningConfig AuthorizationServerCredentialsSigningConfig

// NewAuthorizationServerCredentialsSigningConfig instantiates a new AuthorizationServerCredentialsSigningConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerCredentialsSigningConfig() *AuthorizationServerCredentialsSigningConfig {
	this := AuthorizationServerCredentialsSigningConfig{}
	return &this
}

// NewAuthorizationServerCredentialsSigningConfigWithDefaults instantiates a new AuthorizationServerCredentialsSigningConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerCredentialsSigningConfigWithDefaults() *AuthorizationServerCredentialsSigningConfig {
	this := AuthorizationServerCredentialsSigningConfig{}
	return &this
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *AuthorizationServerCredentialsSigningConfig) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerCredentialsSigningConfig) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *AuthorizationServerCredentialsSigningConfig) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *AuthorizationServerCredentialsSigningConfig) SetKid(v string) {
	o.Kid = &v
}

// GetLastRotated returns the LastRotated field value if set, zero value otherwise.
func (o *AuthorizationServerCredentialsSigningConfig) GetLastRotated() time.Time {
	if o == nil || o.LastRotated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastRotated
}

// GetLastRotatedOk returns a tuple with the LastRotated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerCredentialsSigningConfig) GetLastRotatedOk() (*time.Time, bool) {
	if o == nil || o.LastRotated == nil {
		return nil, false
	}
	return o.LastRotated, true
}

// HasLastRotated returns a boolean if a field has been set.
func (o *AuthorizationServerCredentialsSigningConfig) HasLastRotated() bool {
	if o != nil && o.LastRotated != nil {
		return true
	}

	return false
}

// SetLastRotated gets a reference to the given time.Time and assigns it to the LastRotated field.
func (o *AuthorizationServerCredentialsSigningConfig) SetLastRotated(v time.Time) {
	o.LastRotated = &v
}

// GetNextRotation returns the NextRotation field value if set, zero value otherwise.
func (o *AuthorizationServerCredentialsSigningConfig) GetNextRotation() time.Time {
	if o == nil || o.NextRotation == nil {
		var ret time.Time
		return ret
	}
	return *o.NextRotation
}

// GetNextRotationOk returns a tuple with the NextRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerCredentialsSigningConfig) GetNextRotationOk() (*time.Time, bool) {
	if o == nil || o.NextRotation == nil {
		return nil, false
	}
	return o.NextRotation, true
}

// HasNextRotation returns a boolean if a field has been set.
func (o *AuthorizationServerCredentialsSigningConfig) HasNextRotation() bool {
	if o != nil && o.NextRotation != nil {
		return true
	}

	return false
}

// SetNextRotation gets a reference to the given time.Time and assigns it to the NextRotation field.
func (o *AuthorizationServerCredentialsSigningConfig) SetNextRotation(v time.Time) {
	o.NextRotation = &v
}

// GetRotationMode returns the RotationMode field value if set, zero value otherwise.
func (o *AuthorizationServerCredentialsSigningConfig) GetRotationMode() string {
	if o == nil || o.RotationMode == nil {
		var ret string
		return ret
	}
	return *o.RotationMode
}

// GetRotationModeOk returns a tuple with the RotationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerCredentialsSigningConfig) GetRotationModeOk() (*string, bool) {
	if o == nil || o.RotationMode == nil {
		return nil, false
	}
	return o.RotationMode, true
}

// HasRotationMode returns a boolean if a field has been set.
func (o *AuthorizationServerCredentialsSigningConfig) HasRotationMode() bool {
	if o != nil && o.RotationMode != nil {
		return true
	}

	return false
}

// SetRotationMode gets a reference to the given string and assigns it to the RotationMode field.
func (o *AuthorizationServerCredentialsSigningConfig) SetRotationMode(v string) {
	o.RotationMode = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *AuthorizationServerCredentialsSigningConfig) GetUse() string {
	if o == nil || o.Use == nil {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerCredentialsSigningConfig) GetUseOk() (*string, bool) {
	if o == nil || o.Use == nil {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *AuthorizationServerCredentialsSigningConfig) HasUse() bool {
	if o != nil && o.Use != nil {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *AuthorizationServerCredentialsSigningConfig) SetUse(v string) {
	o.Use = &v
}

func (o AuthorizationServerCredentialsSigningConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kid != nil {
		toSerialize["kid"] = o.Kid
	}
	if o.LastRotated != nil {
		toSerialize["lastRotated"] = o.LastRotated
	}
	if o.NextRotation != nil {
		toSerialize["nextRotation"] = o.NextRotation
	}
	if o.RotationMode != nil {
		toSerialize["rotationMode"] = o.RotationMode
	}
	if o.Use != nil {
		toSerialize["use"] = o.Use
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerCredentialsSigningConfig) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerCredentialsSigningConfig := _AuthorizationServerCredentialsSigningConfig{}

	err = json.Unmarshal(bytes, &varAuthorizationServerCredentialsSigningConfig)
	if err == nil {
		*o = AuthorizationServerCredentialsSigningConfig(varAuthorizationServerCredentialsSigningConfig)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "kid")
		delete(additionalProperties, "lastRotated")
		delete(additionalProperties, "nextRotation")
		delete(additionalProperties, "rotationMode")
		delete(additionalProperties, "use")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerCredentialsSigningConfig struct {
	value *AuthorizationServerCredentialsSigningConfig
	isSet bool
}

func (v NullableAuthorizationServerCredentialsSigningConfig) Get() *AuthorizationServerCredentialsSigningConfig {
	return v.value
}

func (v *NullableAuthorizationServerCredentialsSigningConfig) Set(val *AuthorizationServerCredentialsSigningConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerCredentialsSigningConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerCredentialsSigningConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerCredentialsSigningConfig(val *AuthorizationServerCredentialsSigningConfig) *NullableAuthorizationServerCredentialsSigningConfig {
	return &NullableAuthorizationServerCredentialsSigningConfig{value: val, isSet: true}
}

func (v NullableAuthorizationServerCredentialsSigningConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerCredentialsSigningConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

