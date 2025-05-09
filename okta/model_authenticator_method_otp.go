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
	"reflect"
	"strings"
)

// AuthenticatorMethodOtp struct for AuthenticatorMethodOtp
type AuthenticatorMethodOtp struct {
	AuthenticatorMethodWithVerifiableProperties
	// The number of acceptable adjacent intervals, also known as the clock drift interval. This setting allows you to build in tolerance for any time difference between the token and the server. For example, with a `timeIntervalInSeconds` of 60 seconds and an `acceptableAdjacentIntervals` value of 5, Okta accepts passcodes within 300 seconds (60 * 5) before or after the end user enters their code.
	AcceptableAdjacentIntervals *int32 `json:"acceptableAdjacentIntervals,omitempty"`
	// HMAC algorithm
	Algorithm *string `json:"algorithm,omitempty"`
	// The shared secret encoding
	Encoding *string `json:"encoding,omitempty"`
	// The `id` value of the factor profile
	FactorProfileId *string `json:"factorProfileId,omitempty"`
	// Number of digits in an OTP value
	PassCodeLength *int32 `json:"passCodeLength,omitempty"`
	// The protocol used
	Protocol *string `json:"protocol,omitempty"`
	// Time interval for TOTP in seconds
	TimeIntervalInSeconds *int32 `json:"timeIntervalInSeconds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodOtp AuthenticatorMethodOtp

// NewAuthenticatorMethodOtp instantiates a new AuthenticatorMethodOtp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodOtp() *AuthenticatorMethodOtp {
	this := AuthenticatorMethodOtp{}
	return &this
}

// NewAuthenticatorMethodOtpWithDefaults instantiates a new AuthenticatorMethodOtp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodOtpWithDefaults() *AuthenticatorMethodOtp {
	this := AuthenticatorMethodOtp{}
	return &this
}

// GetAcceptableAdjacentIntervals returns the AcceptableAdjacentIntervals field value if set, zero value otherwise.
func (o *AuthenticatorMethodOtp) GetAcceptableAdjacentIntervals() int32 {
	if o == nil || o.AcceptableAdjacentIntervals == nil {
		var ret int32
		return ret
	}
	return *o.AcceptableAdjacentIntervals
}

// GetAcceptableAdjacentIntervalsOk returns a tuple with the AcceptableAdjacentIntervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodOtp) GetAcceptableAdjacentIntervalsOk() (*int32, bool) {
	if o == nil || o.AcceptableAdjacentIntervals == nil {
		return nil, false
	}
	return o.AcceptableAdjacentIntervals, true
}

// HasAcceptableAdjacentIntervals returns a boolean if a field has been set.
func (o *AuthenticatorMethodOtp) HasAcceptableAdjacentIntervals() bool {
	if o != nil && o.AcceptableAdjacentIntervals != nil {
		return true
	}

	return false
}

// SetAcceptableAdjacentIntervals gets a reference to the given int32 and assigns it to the AcceptableAdjacentIntervals field.
func (o *AuthenticatorMethodOtp) SetAcceptableAdjacentIntervals(v int32) {
	o.AcceptableAdjacentIntervals = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *AuthenticatorMethodOtp) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodOtp) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *AuthenticatorMethodOtp) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *AuthenticatorMethodOtp) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *AuthenticatorMethodOtp) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodOtp) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *AuthenticatorMethodOtp) HasEncoding() bool {
	if o != nil && o.Encoding != nil {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *AuthenticatorMethodOtp) SetEncoding(v string) {
	o.Encoding = &v
}

// GetFactorProfileId returns the FactorProfileId field value if set, zero value otherwise.
func (o *AuthenticatorMethodOtp) GetFactorProfileId() string {
	if o == nil || o.FactorProfileId == nil {
		var ret string
		return ret
	}
	return *o.FactorProfileId
}

// GetFactorProfileIdOk returns a tuple with the FactorProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodOtp) GetFactorProfileIdOk() (*string, bool) {
	if o == nil || o.FactorProfileId == nil {
		return nil, false
	}
	return o.FactorProfileId, true
}

// HasFactorProfileId returns a boolean if a field has been set.
func (o *AuthenticatorMethodOtp) HasFactorProfileId() bool {
	if o != nil && o.FactorProfileId != nil {
		return true
	}

	return false
}

// SetFactorProfileId gets a reference to the given string and assigns it to the FactorProfileId field.
func (o *AuthenticatorMethodOtp) SetFactorProfileId(v string) {
	o.FactorProfileId = &v
}

// GetPassCodeLength returns the PassCodeLength field value if set, zero value otherwise.
func (o *AuthenticatorMethodOtp) GetPassCodeLength() int32 {
	if o == nil || o.PassCodeLength == nil {
		var ret int32
		return ret
	}
	return *o.PassCodeLength
}

// GetPassCodeLengthOk returns a tuple with the PassCodeLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodOtp) GetPassCodeLengthOk() (*int32, bool) {
	if o == nil || o.PassCodeLength == nil {
		return nil, false
	}
	return o.PassCodeLength, true
}

// HasPassCodeLength returns a boolean if a field has been set.
func (o *AuthenticatorMethodOtp) HasPassCodeLength() bool {
	if o != nil && o.PassCodeLength != nil {
		return true
	}

	return false
}

// SetPassCodeLength gets a reference to the given int32 and assigns it to the PassCodeLength field.
func (o *AuthenticatorMethodOtp) SetPassCodeLength(v int32) {
	o.PassCodeLength = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AuthenticatorMethodOtp) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodOtp) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AuthenticatorMethodOtp) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *AuthenticatorMethodOtp) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTimeIntervalInSeconds returns the TimeIntervalInSeconds field value if set, zero value otherwise.
func (o *AuthenticatorMethodOtp) GetTimeIntervalInSeconds() int32 {
	if o == nil || o.TimeIntervalInSeconds == nil {
		var ret int32
		return ret
	}
	return *o.TimeIntervalInSeconds
}

// GetTimeIntervalInSecondsOk returns a tuple with the TimeIntervalInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodOtp) GetTimeIntervalInSecondsOk() (*int32, bool) {
	if o == nil || o.TimeIntervalInSeconds == nil {
		return nil, false
	}
	return o.TimeIntervalInSeconds, true
}

// HasTimeIntervalInSeconds returns a boolean if a field has been set.
func (o *AuthenticatorMethodOtp) HasTimeIntervalInSeconds() bool {
	if o != nil && o.TimeIntervalInSeconds != nil {
		return true
	}

	return false
}

// SetTimeIntervalInSeconds gets a reference to the given int32 and assigns it to the TimeIntervalInSeconds field.
func (o *AuthenticatorMethodOtp) SetTimeIntervalInSeconds(v int32) {
	o.TimeIntervalInSeconds = &v
}

func (o AuthenticatorMethodOtp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthenticatorMethodWithVerifiableProperties, errAuthenticatorMethodWithVerifiableProperties := json.Marshal(o.AuthenticatorMethodWithVerifiableProperties)
	if errAuthenticatorMethodWithVerifiableProperties != nil {
		return []byte{}, errAuthenticatorMethodWithVerifiableProperties
	}
	errAuthenticatorMethodWithVerifiableProperties = json.Unmarshal([]byte(serializedAuthenticatorMethodWithVerifiableProperties), &toSerialize)
	if errAuthenticatorMethodWithVerifiableProperties != nil {
		return []byte{}, errAuthenticatorMethodWithVerifiableProperties
	}
	if o.AcceptableAdjacentIntervals != nil {
		toSerialize["acceptableAdjacentIntervals"] = o.AcceptableAdjacentIntervals
	}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.FactorProfileId != nil {
		toSerialize["factorProfileId"] = o.FactorProfileId
	}
	if o.PassCodeLength != nil {
		toSerialize["passCodeLength"] = o.PassCodeLength
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.TimeIntervalInSeconds != nil {
		toSerialize["timeIntervalInSeconds"] = o.TimeIntervalInSeconds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorMethodOtp) UnmarshalJSON(bytes []byte) (err error) {
	type AuthenticatorMethodOtpWithoutEmbeddedStruct struct {
		// The number of acceptable adjacent intervals, also known as the clock drift interval. This setting allows you to build in tolerance for any time difference between the token and the server. For example, with a `timeIntervalInSeconds` of 60 seconds and an `acceptableAdjacentIntervals` value of 5, Okta accepts passcodes within 300 seconds (60 * 5) before or after the end user enters their code.
		AcceptableAdjacentIntervals *int32 `json:"acceptableAdjacentIntervals,omitempty"`
		// HMAC algorithm
		Algorithm *string `json:"algorithm,omitempty"`
		// The shared secret encoding
		Encoding *string `json:"encoding,omitempty"`
		// The `id` value of the factor profile
		FactorProfileId *string `json:"factorProfileId,omitempty"`
		// Number of digits in an OTP value
		PassCodeLength *int32 `json:"passCodeLength,omitempty"`
		// The protocol used
		Protocol *string `json:"protocol,omitempty"`
		// Time interval for TOTP in seconds
		TimeIntervalInSeconds *int32 `json:"timeIntervalInSeconds,omitempty"`
	}

	varAuthenticatorMethodOtpWithoutEmbeddedStruct := AuthenticatorMethodOtpWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodOtpWithoutEmbeddedStruct)
	if err == nil {
		varAuthenticatorMethodOtp := _AuthenticatorMethodOtp{}
		varAuthenticatorMethodOtp.AcceptableAdjacentIntervals = varAuthenticatorMethodOtpWithoutEmbeddedStruct.AcceptableAdjacentIntervals
		varAuthenticatorMethodOtp.Algorithm = varAuthenticatorMethodOtpWithoutEmbeddedStruct.Algorithm
		varAuthenticatorMethodOtp.Encoding = varAuthenticatorMethodOtpWithoutEmbeddedStruct.Encoding
		varAuthenticatorMethodOtp.FactorProfileId = varAuthenticatorMethodOtpWithoutEmbeddedStruct.FactorProfileId
		varAuthenticatorMethodOtp.PassCodeLength = varAuthenticatorMethodOtpWithoutEmbeddedStruct.PassCodeLength
		varAuthenticatorMethodOtp.Protocol = varAuthenticatorMethodOtpWithoutEmbeddedStruct.Protocol
		varAuthenticatorMethodOtp.TimeIntervalInSeconds = varAuthenticatorMethodOtpWithoutEmbeddedStruct.TimeIntervalInSeconds
		*o = AuthenticatorMethodOtp(varAuthenticatorMethodOtp)
	} else {
		return err
	}

	varAuthenticatorMethodOtp := _AuthenticatorMethodOtp{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodOtp)
	if err == nil {
		o.AuthenticatorMethodWithVerifiableProperties = varAuthenticatorMethodOtp.AuthenticatorMethodWithVerifiableProperties
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "acceptableAdjacentIntervals")
		delete(additionalProperties, "algorithm")
		delete(additionalProperties, "encoding")
		delete(additionalProperties, "factorProfileId")
		delete(additionalProperties, "passCodeLength")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "timeIntervalInSeconds")

		// remove fields from embedded structs
		reflectAuthenticatorMethodWithVerifiableProperties := reflect.ValueOf(o.AuthenticatorMethodWithVerifiableProperties)
		for i := 0; i < reflectAuthenticatorMethodWithVerifiableProperties.Type().NumField(); i++ {
			t := reflectAuthenticatorMethodWithVerifiableProperties.Type().Field(i)

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
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorMethodOtp struct {
	value *AuthenticatorMethodOtp
	isSet bool
}

func (v NullableAuthenticatorMethodOtp) Get() *AuthenticatorMethodOtp {
	return v.value
}

func (v *NullableAuthenticatorMethodOtp) Set(val *AuthenticatorMethodOtp) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodOtp) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodOtp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodOtp(val *AuthenticatorMethodOtp) *NullableAuthenticatorMethodOtp {
	return &NullableAuthenticatorMethodOtp{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodOtp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodOtp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

