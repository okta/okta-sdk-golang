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

// AuthenticatorMethodTotpAllOfSettings struct for AuthenticatorMethodTotpAllOfSettings
type AuthenticatorMethodTotpAllOfSettings struct {
	// Time interval for TOTP in seconds
	TimeIntervalInSeconds *int32 `json:"timeIntervalInSeconds,omitempty"`
	// The shared secret encoding
	Encoding *string `json:"encoding,omitempty"`
	// HMAC algorithm
	Algorithm *string `json:"algorithm,omitempty"`
	// Number of digits in an OTP value
	PassCodeLength *int32 `json:"passCodeLength,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodTotpAllOfSettings AuthenticatorMethodTotpAllOfSettings

// NewAuthenticatorMethodTotpAllOfSettings instantiates a new AuthenticatorMethodTotpAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodTotpAllOfSettings() *AuthenticatorMethodTotpAllOfSettings {
	this := AuthenticatorMethodTotpAllOfSettings{}
	return &this
}

// NewAuthenticatorMethodTotpAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodTotpAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodTotpAllOfSettingsWithDefaults() *AuthenticatorMethodTotpAllOfSettings {
	this := AuthenticatorMethodTotpAllOfSettings{}
	return &this
}

// GetTimeIntervalInSeconds returns the TimeIntervalInSeconds field value if set, zero value otherwise.
func (o *AuthenticatorMethodTotpAllOfSettings) GetTimeIntervalInSeconds() int32 {
	if o == nil || o.TimeIntervalInSeconds == nil {
		var ret int32
		return ret
	}
	return *o.TimeIntervalInSeconds
}

// GetTimeIntervalInSecondsOk returns a tuple with the TimeIntervalInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodTotpAllOfSettings) GetTimeIntervalInSecondsOk() (*int32, bool) {
	if o == nil || o.TimeIntervalInSeconds == nil {
		return nil, false
	}
	return o.TimeIntervalInSeconds, true
}

// HasTimeIntervalInSeconds returns a boolean if a field has been set.
func (o *AuthenticatorMethodTotpAllOfSettings) HasTimeIntervalInSeconds() bool {
	if o != nil && o.TimeIntervalInSeconds != nil {
		return true
	}

	return false
}

// SetTimeIntervalInSeconds gets a reference to the given int32 and assigns it to the TimeIntervalInSeconds field.
func (o *AuthenticatorMethodTotpAllOfSettings) SetTimeIntervalInSeconds(v int32) {
	o.TimeIntervalInSeconds = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *AuthenticatorMethodTotpAllOfSettings) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodTotpAllOfSettings) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *AuthenticatorMethodTotpAllOfSettings) HasEncoding() bool {
	if o != nil && o.Encoding != nil {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *AuthenticatorMethodTotpAllOfSettings) SetEncoding(v string) {
	o.Encoding = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *AuthenticatorMethodTotpAllOfSettings) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodTotpAllOfSettings) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *AuthenticatorMethodTotpAllOfSettings) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *AuthenticatorMethodTotpAllOfSettings) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetPassCodeLength returns the PassCodeLength field value if set, zero value otherwise.
func (o *AuthenticatorMethodTotpAllOfSettings) GetPassCodeLength() int32 {
	if o == nil || o.PassCodeLength == nil {
		var ret int32
		return ret
	}
	return *o.PassCodeLength
}

// GetPassCodeLengthOk returns a tuple with the PassCodeLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodTotpAllOfSettings) GetPassCodeLengthOk() (*int32, bool) {
	if o == nil || o.PassCodeLength == nil {
		return nil, false
	}
	return o.PassCodeLength, true
}

// HasPassCodeLength returns a boolean if a field has been set.
func (o *AuthenticatorMethodTotpAllOfSettings) HasPassCodeLength() bool {
	if o != nil && o.PassCodeLength != nil {
		return true
	}

	return false
}

// SetPassCodeLength gets a reference to the given int32 and assigns it to the PassCodeLength field.
func (o *AuthenticatorMethodTotpAllOfSettings) SetPassCodeLength(v int32) {
	o.PassCodeLength = &v
}

func (o AuthenticatorMethodTotpAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TimeIntervalInSeconds != nil {
		toSerialize["timeIntervalInSeconds"] = o.TimeIntervalInSeconds
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.PassCodeLength != nil {
		toSerialize["passCodeLength"] = o.PassCodeLength
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorMethodTotpAllOfSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorMethodTotpAllOfSettings := _AuthenticatorMethodTotpAllOfSettings{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodTotpAllOfSettings)
	if err == nil {
		*o = AuthenticatorMethodTotpAllOfSettings(varAuthenticatorMethodTotpAllOfSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "timeIntervalInSeconds")
		delete(additionalProperties, "encoding")
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "passCodeLength")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorMethodTotpAllOfSettings struct {
	value *AuthenticatorMethodTotpAllOfSettings
	isSet bool
}

func (v NullableAuthenticatorMethodTotpAllOfSettings) Get() *AuthenticatorMethodTotpAllOfSettings {
	return v.value
}

func (v *NullableAuthenticatorMethodTotpAllOfSettings) Set(val *AuthenticatorMethodTotpAllOfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodTotpAllOfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodTotpAllOfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodTotpAllOfSettings(val *AuthenticatorMethodTotpAllOfSettings) *NullableAuthenticatorMethodTotpAllOfSettings {
	return &NullableAuthenticatorMethodTotpAllOfSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodTotpAllOfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodTotpAllOfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

