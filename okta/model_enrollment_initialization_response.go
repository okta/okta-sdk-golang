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

// EnrollmentInitializationResponse Yubico Transport Key in the form of a JWK, used to encrypt our fulfillment request to Yubico. The currently agreed protocol uses P-384.
type EnrollmentInitializationResponse struct {
	// List of credential requests for the fulfillment provider
	CredRequests []WebAuthnCredRequest `json:"credRequests,omitempty"`
	// Name of the fulfillment provider for the WebAuthn Preregistration Factor
	FulfillmentProvider *string `json:"fulfillmentProvider,omitempty"`
	// Encrypted JWE of PIN request for the fulfillment provider
	PinRequestJwe *string `json:"pinRequestJwe,omitempty"`
	// ID of an existing Okta user
	UserId *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnrollmentInitializationResponse EnrollmentInitializationResponse

// NewEnrollmentInitializationResponse instantiates a new EnrollmentInitializationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentInitializationResponse() *EnrollmentInitializationResponse {
	this := EnrollmentInitializationResponse{}
	return &this
}

// NewEnrollmentInitializationResponseWithDefaults instantiates a new EnrollmentInitializationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentInitializationResponseWithDefaults() *EnrollmentInitializationResponse {
	this := EnrollmentInitializationResponse{}
	return &this
}

// GetCredRequests returns the CredRequests field value if set, zero value otherwise.
func (o *EnrollmentInitializationResponse) GetCredRequests() []WebAuthnCredRequest {
	if o == nil || o.CredRequests == nil {
		var ret []WebAuthnCredRequest
		return ret
	}
	return o.CredRequests
}

// GetCredRequestsOk returns a tuple with the CredRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentInitializationResponse) GetCredRequestsOk() ([]WebAuthnCredRequest, bool) {
	if o == nil || o.CredRequests == nil {
		return nil, false
	}
	return o.CredRequests, true
}

// HasCredRequests returns a boolean if a field has been set.
func (o *EnrollmentInitializationResponse) HasCredRequests() bool {
	if o != nil && o.CredRequests != nil {
		return true
	}

	return false
}

// SetCredRequests gets a reference to the given []WebAuthnCredRequest and assigns it to the CredRequests field.
func (o *EnrollmentInitializationResponse) SetCredRequests(v []WebAuthnCredRequest) {
	o.CredRequests = v
}

// GetFulfillmentProvider returns the FulfillmentProvider field value if set, zero value otherwise.
func (o *EnrollmentInitializationResponse) GetFulfillmentProvider() string {
	if o == nil || o.FulfillmentProvider == nil {
		var ret string
		return ret
	}
	return *o.FulfillmentProvider
}

// GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentInitializationResponse) GetFulfillmentProviderOk() (*string, bool) {
	if o == nil || o.FulfillmentProvider == nil {
		return nil, false
	}
	return o.FulfillmentProvider, true
}

// HasFulfillmentProvider returns a boolean if a field has been set.
func (o *EnrollmentInitializationResponse) HasFulfillmentProvider() bool {
	if o != nil && o.FulfillmentProvider != nil {
		return true
	}

	return false
}

// SetFulfillmentProvider gets a reference to the given string and assigns it to the FulfillmentProvider field.
func (o *EnrollmentInitializationResponse) SetFulfillmentProvider(v string) {
	o.FulfillmentProvider = &v
}

// GetPinRequestJwe returns the PinRequestJwe field value if set, zero value otherwise.
func (o *EnrollmentInitializationResponse) GetPinRequestJwe() string {
	if o == nil || o.PinRequestJwe == nil {
		var ret string
		return ret
	}
	return *o.PinRequestJwe
}

// GetPinRequestJweOk returns a tuple with the PinRequestJwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentInitializationResponse) GetPinRequestJweOk() (*string, bool) {
	if o == nil || o.PinRequestJwe == nil {
		return nil, false
	}
	return o.PinRequestJwe, true
}

// HasPinRequestJwe returns a boolean if a field has been set.
func (o *EnrollmentInitializationResponse) HasPinRequestJwe() bool {
	if o != nil && o.PinRequestJwe != nil {
		return true
	}

	return false
}

// SetPinRequestJwe gets a reference to the given string and assigns it to the PinRequestJwe field.
func (o *EnrollmentInitializationResponse) SetPinRequestJwe(v string) {
	o.PinRequestJwe = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EnrollmentInitializationResponse) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentInitializationResponse) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EnrollmentInitializationResponse) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EnrollmentInitializationResponse) SetUserId(v string) {
	o.UserId = &v
}

func (o EnrollmentInitializationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CredRequests != nil {
		toSerialize["credRequests"] = o.CredRequests
	}
	if o.FulfillmentProvider != nil {
		toSerialize["fulfillmentProvider"] = o.FulfillmentProvider
	}
	if o.PinRequestJwe != nil {
		toSerialize["pinRequestJwe"] = o.PinRequestJwe
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EnrollmentInitializationResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEnrollmentInitializationResponse := _EnrollmentInitializationResponse{}

	err = json.Unmarshal(bytes, &varEnrollmentInitializationResponse)
	if err == nil {
		*o = EnrollmentInitializationResponse(varEnrollmentInitializationResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credRequests")
		delete(additionalProperties, "fulfillmentProvider")
		delete(additionalProperties, "pinRequestJwe")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEnrollmentInitializationResponse struct {
	value *EnrollmentInitializationResponse
	isSet bool
}

func (v NullableEnrollmentInitializationResponse) Get() *EnrollmentInitializationResponse {
	return v.value
}

func (v *NullableEnrollmentInitializationResponse) Set(val *EnrollmentInitializationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentInitializationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentInitializationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentInitializationResponse(val *EnrollmentInitializationResponse) *NullableEnrollmentInitializationResponse {
	return &NullableEnrollmentInitializationResponse{value: val, isSet: true}
}

func (v NullableEnrollmentInitializationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentInitializationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

