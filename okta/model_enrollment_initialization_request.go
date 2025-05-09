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

// EnrollmentInitializationRequest Enrollment Initialization Request
type EnrollmentInitializationRequest struct {
	// List of Relying Party hostnames to register on the YubiKey.
	EnrollmentRpIds []string `json:"enrollmentRpIds,omitempty"`
	// Name of the fulfillment provider for the WebAuthn Preregistration Factor
	FulfillmentProvider *string `json:"fulfillmentProvider,omitempty"`
	// ID of an existing Okta user
	UserId *string `json:"userId,omitempty"`
	YubicoTransportKeyJWK *ECKeyJWK `json:"yubicoTransportKeyJWK,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnrollmentInitializationRequest EnrollmentInitializationRequest

// NewEnrollmentInitializationRequest instantiates a new EnrollmentInitializationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentInitializationRequest() *EnrollmentInitializationRequest {
	this := EnrollmentInitializationRequest{}
	return &this
}

// NewEnrollmentInitializationRequestWithDefaults instantiates a new EnrollmentInitializationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentInitializationRequestWithDefaults() *EnrollmentInitializationRequest {
	this := EnrollmentInitializationRequest{}
	return &this
}

// GetEnrollmentRpIds returns the EnrollmentRpIds field value if set, zero value otherwise.
func (o *EnrollmentInitializationRequest) GetEnrollmentRpIds() []string {
	if o == nil || o.EnrollmentRpIds == nil {
		var ret []string
		return ret
	}
	return o.EnrollmentRpIds
}

// GetEnrollmentRpIdsOk returns a tuple with the EnrollmentRpIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentInitializationRequest) GetEnrollmentRpIdsOk() ([]string, bool) {
	if o == nil || o.EnrollmentRpIds == nil {
		return nil, false
	}
	return o.EnrollmentRpIds, true
}

// HasEnrollmentRpIds returns a boolean if a field has been set.
func (o *EnrollmentInitializationRequest) HasEnrollmentRpIds() bool {
	if o != nil && o.EnrollmentRpIds != nil {
		return true
	}

	return false
}

// SetEnrollmentRpIds gets a reference to the given []string and assigns it to the EnrollmentRpIds field.
func (o *EnrollmentInitializationRequest) SetEnrollmentRpIds(v []string) {
	o.EnrollmentRpIds = v
}

// GetFulfillmentProvider returns the FulfillmentProvider field value if set, zero value otherwise.
func (o *EnrollmentInitializationRequest) GetFulfillmentProvider() string {
	if o == nil || o.FulfillmentProvider == nil {
		var ret string
		return ret
	}
	return *o.FulfillmentProvider
}

// GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentInitializationRequest) GetFulfillmentProviderOk() (*string, bool) {
	if o == nil || o.FulfillmentProvider == nil {
		return nil, false
	}
	return o.FulfillmentProvider, true
}

// HasFulfillmentProvider returns a boolean if a field has been set.
func (o *EnrollmentInitializationRequest) HasFulfillmentProvider() bool {
	if o != nil && o.FulfillmentProvider != nil {
		return true
	}

	return false
}

// SetFulfillmentProvider gets a reference to the given string and assigns it to the FulfillmentProvider field.
func (o *EnrollmentInitializationRequest) SetFulfillmentProvider(v string) {
	o.FulfillmentProvider = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EnrollmentInitializationRequest) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentInitializationRequest) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EnrollmentInitializationRequest) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EnrollmentInitializationRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetYubicoTransportKeyJWK returns the YubicoTransportKeyJWK field value if set, zero value otherwise.
func (o *EnrollmentInitializationRequest) GetYubicoTransportKeyJWK() ECKeyJWK {
	if o == nil || o.YubicoTransportKeyJWK == nil {
		var ret ECKeyJWK
		return ret
	}
	return *o.YubicoTransportKeyJWK
}

// GetYubicoTransportKeyJWKOk returns a tuple with the YubicoTransportKeyJWK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentInitializationRequest) GetYubicoTransportKeyJWKOk() (*ECKeyJWK, bool) {
	if o == nil || o.YubicoTransportKeyJWK == nil {
		return nil, false
	}
	return o.YubicoTransportKeyJWK, true
}

// HasYubicoTransportKeyJWK returns a boolean if a field has been set.
func (o *EnrollmentInitializationRequest) HasYubicoTransportKeyJWK() bool {
	if o != nil && o.YubicoTransportKeyJWK != nil {
		return true
	}

	return false
}

// SetYubicoTransportKeyJWK gets a reference to the given ECKeyJWK and assigns it to the YubicoTransportKeyJWK field.
func (o *EnrollmentInitializationRequest) SetYubicoTransportKeyJWK(v ECKeyJWK) {
	o.YubicoTransportKeyJWK = &v
}

func (o EnrollmentInitializationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EnrollmentRpIds != nil {
		toSerialize["enrollmentRpIds"] = o.EnrollmentRpIds
	}
	if o.FulfillmentProvider != nil {
		toSerialize["fulfillmentProvider"] = o.FulfillmentProvider
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.YubicoTransportKeyJWK != nil {
		toSerialize["yubicoTransportKeyJWK"] = o.YubicoTransportKeyJWK
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EnrollmentInitializationRequest) UnmarshalJSON(bytes []byte) (err error) {
	varEnrollmentInitializationRequest := _EnrollmentInitializationRequest{}

	err = json.Unmarshal(bytes, &varEnrollmentInitializationRequest)
	if err == nil {
		*o = EnrollmentInitializationRequest(varEnrollmentInitializationRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "enrollmentRpIds")
		delete(additionalProperties, "fulfillmentProvider")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "yubicoTransportKeyJWK")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEnrollmentInitializationRequest struct {
	value *EnrollmentInitializationRequest
	isSet bool
}

func (v NullableEnrollmentInitializationRequest) Get() *EnrollmentInitializationRequest {
	return v.value
}

func (v *NullableEnrollmentInitializationRequest) Set(val *EnrollmentInitializationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentInitializationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentInitializationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentInitializationRequest(val *EnrollmentInitializationRequest) *NullableEnrollmentInitializationRequest {
	return &NullableEnrollmentInitializationRequest{value: val, isSet: true}
}

func (v NullableEnrollmentInitializationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentInitializationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

