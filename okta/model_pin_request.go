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

// checks if the PinRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PinRequest{}

// PinRequest Pin request
type PinRequest struct {
	// ID for a WebAuthn preregistration factor in Okta
	AuthenticatorEnrollmentId *string `json:"authenticatorEnrollmentId,omitempty"`
	// Name of the fulfillment provider for the WebAuthn preregistration factor
	FulfillmentProvider *string `json:"fulfillmentProvider,omitempty"`
	// ID of an existing Okta user
	UserId               *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PinRequest PinRequest

// NewPinRequest instantiates a new PinRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinRequest() *PinRequest {
	this := PinRequest{}
	return &this
}

// NewPinRequestWithDefaults instantiates a new PinRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinRequestWithDefaults() *PinRequest {
	this := PinRequest{}
	return &this
}

// GetAuthenticatorEnrollmentId returns the AuthenticatorEnrollmentId field value if set, zero value otherwise.
func (o *PinRequest) GetAuthenticatorEnrollmentId() string {
	if o == nil || IsNil(o.AuthenticatorEnrollmentId) {
		var ret string
		return ret
	}
	return *o.AuthenticatorEnrollmentId
}

// GetAuthenticatorEnrollmentIdOk returns a tuple with the AuthenticatorEnrollmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinRequest) GetAuthenticatorEnrollmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticatorEnrollmentId) {
		return nil, false
	}
	return o.AuthenticatorEnrollmentId, true
}

// HasAuthenticatorEnrollmentId returns a boolean if a field has been set.
func (o *PinRequest) HasAuthenticatorEnrollmentId() bool {
	if o != nil && !IsNil(o.AuthenticatorEnrollmentId) {
		return true
	}

	return false
}

// SetAuthenticatorEnrollmentId gets a reference to the given string and assigns it to the AuthenticatorEnrollmentId field.
func (o *PinRequest) SetAuthenticatorEnrollmentId(v string) {
	o.AuthenticatorEnrollmentId = &v
}

// GetFulfillmentProvider returns the FulfillmentProvider field value if set, zero value otherwise.
func (o *PinRequest) GetFulfillmentProvider() string {
	if o == nil || IsNil(o.FulfillmentProvider) {
		var ret string
		return ret
	}
	return *o.FulfillmentProvider
}

// GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinRequest) GetFulfillmentProviderOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentProvider) {
		return nil, false
	}
	return o.FulfillmentProvider, true
}

// HasFulfillmentProvider returns a boolean if a field has been set.
func (o *PinRequest) HasFulfillmentProvider() bool {
	if o != nil && !IsNil(o.FulfillmentProvider) {
		return true
	}

	return false
}

// SetFulfillmentProvider gets a reference to the given string and assigns it to the FulfillmentProvider field.
func (o *PinRequest) SetFulfillmentProvider(v string) {
	o.FulfillmentProvider = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PinRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PinRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PinRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o PinRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PinRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticatorEnrollmentId) {
		toSerialize["authenticatorEnrollmentId"] = o.AuthenticatorEnrollmentId
	}
	if !IsNil(o.FulfillmentProvider) {
		toSerialize["fulfillmentProvider"] = o.FulfillmentProvider
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PinRequest) UnmarshalJSON(data []byte) (err error) {
	varPinRequest := _PinRequest{}

	err = json.Unmarshal(data, &varPinRequest)

	if err != nil {
		return err
	}

	*o = PinRequest(varPinRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticatorEnrollmentId")
		delete(additionalProperties, "fulfillmentProvider")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePinRequest struct {
	value *PinRequest
	isSet bool
}

func (v NullablePinRequest) Get() *PinRequest {
	return v.value
}

func (v *NullablePinRequest) Set(val *PinRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePinRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePinRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinRequest(val *PinRequest) *NullablePinRequest {
	return &NullablePinRequest{value: val, isSet: true}
}

func (v NullablePinRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
