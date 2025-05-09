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

// EnrollmentActivationRequest Enrollment Initialization Request
type EnrollmentActivationRequest struct {
	// List of credential responses from the fulfillment provider
	CredResponses []WebAuthnCredResponse `json:"credResponses,omitempty"`
	// Name of the fulfillment provider for the WebAuthn Preregistration Factor
	FulfillmentProvider *string `json:"fulfillmentProvider,omitempty"`
	// Encrypted JWE of PIN response from the fulfillment provider
	PinResponseJwe *string `json:"pinResponseJwe,omitempty"`
	// Serial number of the YubiKey
	Serial *string `json:"serial,omitempty"`
	// ID of an existing Okta user
	UserId *string `json:"userId,omitempty"`
	// Firmware version of the YubiKey
	Version *string `json:"version,omitempty"`
	// List of usable signing keys from Yubico (in JWKS format) used to verify the JWS inside the JWE
	YubicoSigningJwks []ECKeyJWK `json:"yubicoSigningJwks,omitempty"`
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
	if o == nil || o.CredResponses == nil {
		var ret []WebAuthnCredResponse
		return ret
	}
	return o.CredResponses
}

// GetCredResponsesOk returns a tuple with the CredResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetCredResponsesOk() ([]WebAuthnCredResponse, bool) {
	if o == nil || o.CredResponses == nil {
		return nil, false
	}
	return o.CredResponses, true
}

// HasCredResponses returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasCredResponses() bool {
	if o != nil && o.CredResponses != nil {
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
	if o == nil || o.FulfillmentProvider == nil {
		var ret string
		return ret
	}
	return *o.FulfillmentProvider
}

// GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetFulfillmentProviderOk() (*string, bool) {
	if o == nil || o.FulfillmentProvider == nil {
		return nil, false
	}
	return o.FulfillmentProvider, true
}

// HasFulfillmentProvider returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasFulfillmentProvider() bool {
	if o != nil && o.FulfillmentProvider != nil {
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
	if o == nil || o.PinResponseJwe == nil {
		var ret string
		return ret
	}
	return *o.PinResponseJwe
}

// GetPinResponseJweOk returns a tuple with the PinResponseJwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetPinResponseJweOk() (*string, bool) {
	if o == nil || o.PinResponseJwe == nil {
		return nil, false
	}
	return o.PinResponseJwe, true
}

// HasPinResponseJwe returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasPinResponseJwe() bool {
	if o != nil && o.PinResponseJwe != nil {
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
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasSerial() bool {
	if o != nil && o.Serial != nil {
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
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasUserId() bool {
	if o != nil && o.UserId != nil {
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
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
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
	if o == nil || o.YubicoSigningJwks == nil {
		var ret []ECKeyJWK
		return ret
	}
	return o.YubicoSigningJwks
}

// GetYubicoSigningJwksOk returns a tuple with the YubicoSigningJwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationRequest) GetYubicoSigningJwksOk() ([]ECKeyJWK, bool) {
	if o == nil || o.YubicoSigningJwks == nil {
		return nil, false
	}
	return o.YubicoSigningJwks, true
}

// HasYubicoSigningJwks returns a boolean if a field has been set.
func (o *EnrollmentActivationRequest) HasYubicoSigningJwks() bool {
	if o != nil && o.YubicoSigningJwks != nil {
		return true
	}

	return false
}

// SetYubicoSigningJwks gets a reference to the given []ECKeyJWK and assigns it to the YubicoSigningJwks field.
func (o *EnrollmentActivationRequest) SetYubicoSigningJwks(v []ECKeyJWK) {
	o.YubicoSigningJwks = v
}

func (o EnrollmentActivationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredResponses != nil {
		toSerialize["credResponses"] = o.CredResponses
	}
	if o.FulfillmentProvider != nil {
		toSerialize["fulfillmentProvider"] = o.FulfillmentProvider
	}
	if o.PinResponseJwe != nil {
		toSerialize["pinResponseJwe"] = o.PinResponseJwe
	}
	if o.Serial != nil {
		toSerialize["serial"] = o.Serial
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.YubicoSigningJwks != nil {
		toSerialize["yubicoSigningJwks"] = o.YubicoSigningJwks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EnrollmentActivationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varEnrollmentActivationRequest := _EnrollmentActivationRequest{}

	err = json.Unmarshal(bytes, &varEnrollmentActivationRequest)
	if err == nil {
		*o = EnrollmentActivationRequest(varEnrollmentActivationRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credResponses")
		delete(additionalProperties, "fulfillmentProvider")
		delete(additionalProperties, "pinResponseJwe")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "version")
		delete(additionalProperties, "yubicoSigningJwks")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

