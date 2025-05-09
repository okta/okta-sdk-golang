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

// EnrollmentActivationResponse Enrollment Initialization Response
type EnrollmentActivationResponse struct {
	// List of IDs for preregistered WebAuthn Factors in Okta
	AuthenticatorEnrollmentIds []string `json:"authenticatorEnrollmentIds,omitempty"`
	// Name of the fulfillment provider for the WebAuthn Preregistration Factor
	FulfillmentProvider *string `json:"fulfillmentProvider,omitempty"`
	// ID of an existing Okta user
	UserId *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnrollmentActivationResponse EnrollmentActivationResponse

// NewEnrollmentActivationResponse instantiates a new EnrollmentActivationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentActivationResponse() *EnrollmentActivationResponse {
	this := EnrollmentActivationResponse{}
	return &this
}

// NewEnrollmentActivationResponseWithDefaults instantiates a new EnrollmentActivationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentActivationResponseWithDefaults() *EnrollmentActivationResponse {
	this := EnrollmentActivationResponse{}
	return &this
}

// GetAuthenticatorEnrollmentIds returns the AuthenticatorEnrollmentIds field value if set, zero value otherwise.
func (o *EnrollmentActivationResponse) GetAuthenticatorEnrollmentIds() []string {
	if o == nil || o.AuthenticatorEnrollmentIds == nil {
		var ret []string
		return ret
	}
	return o.AuthenticatorEnrollmentIds
}

// GetAuthenticatorEnrollmentIdsOk returns a tuple with the AuthenticatorEnrollmentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationResponse) GetAuthenticatorEnrollmentIdsOk() ([]string, bool) {
	if o == nil || o.AuthenticatorEnrollmentIds == nil {
		return nil, false
	}
	return o.AuthenticatorEnrollmentIds, true
}

// HasAuthenticatorEnrollmentIds returns a boolean if a field has been set.
func (o *EnrollmentActivationResponse) HasAuthenticatorEnrollmentIds() bool {
	if o != nil && o.AuthenticatorEnrollmentIds != nil {
		return true
	}

	return false
}

// SetAuthenticatorEnrollmentIds gets a reference to the given []string and assigns it to the AuthenticatorEnrollmentIds field.
func (o *EnrollmentActivationResponse) SetAuthenticatorEnrollmentIds(v []string) {
	o.AuthenticatorEnrollmentIds = v
}

// GetFulfillmentProvider returns the FulfillmentProvider field value if set, zero value otherwise.
func (o *EnrollmentActivationResponse) GetFulfillmentProvider() string {
	if o == nil || o.FulfillmentProvider == nil {
		var ret string
		return ret
	}
	return *o.FulfillmentProvider
}

// GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationResponse) GetFulfillmentProviderOk() (*string, bool) {
	if o == nil || o.FulfillmentProvider == nil {
		return nil, false
	}
	return o.FulfillmentProvider, true
}

// HasFulfillmentProvider returns a boolean if a field has been set.
func (o *EnrollmentActivationResponse) HasFulfillmentProvider() bool {
	if o != nil && o.FulfillmentProvider != nil {
		return true
	}

	return false
}

// SetFulfillmentProvider gets a reference to the given string and assigns it to the FulfillmentProvider field.
func (o *EnrollmentActivationResponse) SetFulfillmentProvider(v string) {
	o.FulfillmentProvider = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *EnrollmentActivationResponse) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentActivationResponse) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *EnrollmentActivationResponse) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *EnrollmentActivationResponse) SetUserId(v string) {
	o.UserId = &v
}

func (o EnrollmentActivationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticatorEnrollmentIds != nil {
		toSerialize["authenticatorEnrollmentIds"] = o.AuthenticatorEnrollmentIds
	}
	if o.FulfillmentProvider != nil {
		toSerialize["fulfillmentProvider"] = o.FulfillmentProvider
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EnrollmentActivationResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEnrollmentActivationResponse := _EnrollmentActivationResponse{}

	err = json.Unmarshal(bytes, &varEnrollmentActivationResponse)
	if err == nil {
		*o = EnrollmentActivationResponse(varEnrollmentActivationResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticatorEnrollmentIds")
		delete(additionalProperties, "fulfillmentProvider")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEnrollmentActivationResponse struct {
	value *EnrollmentActivationResponse
	isSet bool
}

func (v NullableEnrollmentActivationResponse) Get() *EnrollmentActivationResponse {
	return v.value
}

func (v *NullableEnrollmentActivationResponse) Set(val *EnrollmentActivationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentActivationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentActivationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentActivationResponse(val *EnrollmentActivationResponse) *NullableEnrollmentActivationResponse {
	return &NullableEnrollmentActivationResponse{value: val, isSet: true}
}

func (v NullableEnrollmentActivationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentActivationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

