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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the EnrollmentActivationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrollmentActivationRequest{}

// EnrollmentActivationRequest Enrollment Initialization Request
type EnrollmentActivationRequest struct {
	// List of credential responses from the fulfillment provider
	CredResponses []WebAuthnCredResponse `json:"credResponses,omitempty"`
	// Name of the fulfillment provider for the WebAuthn preregistration factor
	FulfillmentProvider *string `json:"fulfillmentProvider,omitempty"`
	// Encrypted JWE of the PIN response from the fulfillment provider
	PinResponseJwe *string `json:"pinResponseJwe,omitempty"`
	// Serial number of the YubiKey
	Serial *string `json:"serial,omitempty"`
	// ID of an existing Okta user
	UserId *string `json:"userId,omitempty"`
	// Firmware version of the YubiKey
	Version *string `json:"version,omitempty"`
	// List of usable signing keys from Yubico (in JSON Web Key Sets (JWKS) format). The signing keys are used to verify the JSON Web Signature (JWS) inside the JWE.
	YubicoSigningJwks    []ECKeyJWK `json:"yubicoSigningJwks,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnrollmentActivationRequest EnrollmentActivationRequest

// NewEnrollmentActivationRequest instantiates a new EnrollmentActivationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentActivationRequest() *EnrollmentActivationRequest {
	this := EnrollmentActivationRequest{}
	return &this
}

// NewEnrollmentActivationRequestWithDefaults instantiates a new EnrollmentActivationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentActivationRequestWithDefaults() *EnrollmentActivationRequest {
	this := EnrollmentActivationRequest{}
	return &this
}

// GetCredResponses returns the CredResponses field value if set, zero value otherwise.
func (o *EnrollmentActivationRequest) GetCredResponses() []WebAuthnCredResponse {
	if o == nil || IsNil(o.CredResponses) {
		var ret []WebAuthnCredResponse
		return ret
	}
	return o.CredResponses
}

// GetCredResponsesOk returns a tuple with the CredResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetCredResponsesOk() ([]WebAuthnCredResponse, bool) {
	if o == nil || IsNil(o.CredResponses) {
		return nil, false
	}
	return o.CredResponses, true
}

// HasCredResponses returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasCredResponses() bool {
	if o != nil && !IsNil(o.CredResponses) {
		return true
	}

	return false
}

// SetCredResponses gets a reference to the given []WebAuthnCredResponse and assigns it to the CredResponses field.
func (o *EnrollmentActivationRequest) SetCredResponses(v []WebAuthnCredResponse) {
	o.CredResponses = v
}

// GetFulfillmentProvider returns the FulfillmentProvider field value if set, zero value otherwise.
func (o *EnrollmentActivationRequest) GetFulfillmentProvider() string {
	if o == nil || IsNil(o.FulfillmentProvider) {
		var ret string
		return ret
	}
	return *o.FulfillmentProvider
}

// GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetFulfillmentProviderOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentProvider) {
		return nil, false
	}
	return o.FulfillmentProvider, true
}

// HasFulfillmentProvider returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasFulfillmentProvider() bool {
	if o != nil && !IsNil(o.FulfillmentProvider) {
		return true
	}

	return false
}

// SetFulfillmentProvider gets a reference to the given string and assigns it to the FulfillmentProvider field.
func (o *EnrollmentActivationRequest) SetFulfillmentProvider(v string) {
	o.FulfillmentProvider = &v
}

// GetPinResponseJwe returns the PinResponseJwe field value if set, zero value otherwise.
func (o *EnrollmentActivationRequest) GetPinResponseJwe() string {
	if o == nil || IsNil(o.PinResponseJwe) {
		var ret string
		return ret
	}
	return *o.PinResponseJwe
}

// GetPinResponseJweOk returns a tuple with the PinResponseJwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetPinResponseJweOk() (*string, bool) {
	if o == nil || IsNil(o.PinResponseJwe) {
		return nil, false
	}
	return o.PinResponseJwe, true
}

// HasPinResponseJwe returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasPinResponseJwe() bool {
	if o != nil && !IsNil(o.PinResponseJwe) {
		return true
	}

	return false
}

// SetPinResponseJwe gets a reference to the given string and assigns it to the PinResponseJwe field.
func (o *EnrollmentActivationRequest) SetPinResponseJwe(v string) {
	o.PinResponseJwe = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *EnrollmentActivationRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *EnrollmentActivationRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EnrollmentActivationRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EnrollmentActivationRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EnrollmentActivationRequest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EnrollmentActivationRequest) SetVersion(v string) {
	o.Version = &v
}

// GetYubicoSigningJwks returns the YubicoSigningJwks field value if set, zero value otherwise.
func (o *EnrollmentActivationRequest) GetYubicoSigningJwks() []ECKeyJWK {
	if o == nil || IsNil(o.YubicoSigningJwks) {
		var ret []ECKeyJWK
		return ret
	}
	return o.YubicoSigningJwks
}

// GetYubicoSigningJwksOk returns a tuple with the YubicoSigningJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetYubicoSigningJwksOk() ([]ECKeyJWK, bool) {
	if o == nil || IsNil(o.YubicoSigningJwks) {
		return nil, false
	}
	return o.YubicoSigningJwks, true
}

// HasYubicoSigningJwks returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasYubicoSigningJwks() bool {
	if o != nil && !IsNil(o.YubicoSigningJwks) {
		return true
	}

	return false
}

// SetYubicoSigningJwks gets a reference to the given []ECKeyJWK and assigns it to the YubicoSigningJwks field.
func (o *EnrollmentActivationRequest) SetYubicoSigningJwks(v []ECKeyJWK) {
	o.YubicoSigningJwks = v
}

func (o EnrollmentActivationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollmentActivationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CredResponses) {
		toSerialize["credResponses"] = o.CredResponses
	}
	if !IsNil(o.FulfillmentProvider) {
		toSerialize["fulfillmentProvider"] = o.FulfillmentProvider
	}
	if !IsNil(o.PinResponseJwe) {
		toSerialize["pinResponseJwe"] = o.PinResponseJwe
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.YubicoSigningJwks) {
		toSerialize["yubicoSigningJwks"] = o.YubicoSigningJwks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnrollmentActivationRequest) UnmarshalJSON(data []byte) (err error) {
	varEnrollmentActivationRequest := _EnrollmentActivationRequest{}

	err = json.Unmarshal(data, &varEnrollmentActivationRequest)

	if err != nil {
		return err
	}

	*o = EnrollmentActivationRequest(varEnrollmentActivationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credResponses")
		delete(additionalProperties, "fulfillmentProvider")
		delete(additionalProperties, "pinResponseJwe")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "version")
		delete(additionalProperties, "yubicoSigningJwks")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnrollmentActivationRequest struct {
	value *EnrollmentActivationRequest
	isSet bool
}

func (v NullableEnrollmentActivationRequest) Get() *EnrollmentActivationRequest {
	return v.value
}

func (v *NullableEnrollmentActivationRequest) Set(val *EnrollmentActivationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentActivationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentActivationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentActivationRequest(val *EnrollmentActivationRequest) *NullableEnrollmentActivationRequest {
	return &NullableEnrollmentActivationRequest{value: val, isSet: true}
}

func (v NullableEnrollmentActivationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentActivationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
