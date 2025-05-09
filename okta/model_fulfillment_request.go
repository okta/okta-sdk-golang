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

// FulfillmentRequest Fulfillment Request
type FulfillmentRequest struct {
	FulfillmentData *FulfillmentData `json:"fulfillmentData,omitempty"`
	// Name of the fulfillment provider for the WebAuthn Preregistration Factor
	FulfillmentProvider *string `json:"fulfillmentProvider,omitempty"`
	// ID of an existing Okta user
	UserId *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FulfillmentRequest FulfillmentRequest

// NewFulfillmentRequest instantiates a new FulfillmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentRequest() *FulfillmentRequest {
	this := FulfillmentRequest{}
	return &this
}

// NewFulfillmentRequestWithDefaults instantiates a new FulfillmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentRequestWithDefaults() *FulfillmentRequest {
	this := FulfillmentRequest{}
	return &this
}

// GetFulfillmentData returns the FulfillmentData field value if set, zero value otherwise.
func (o *FulfillmentRequest) GetFulfillmentData() FulfillmentData {
	if o == nil || o.FulfillmentData == nil {
		var ret FulfillmentData
		return ret
	}
	return *o.FulfillmentData
}

// GetFulfillmentDataOk returns a tuple with the FulfillmentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentRequest) GetFulfillmentDataOk() (*FulfillmentData, bool) {
	if o == nil || o.FulfillmentData == nil {
		return nil, false
	}
	return o.FulfillmentData, true
}

// HasFulfillmentData returns a boolean if a field has been set.
func (o *FulfillmentRequest) HasFulfillmentData() bool {
	if o != nil && o.FulfillmentData != nil {
		return true
	}

	return false
}

// SetFulfillmentData gets a reference to the given FulfillmentData and assigns it to the FulfillmentData field.
func (o *FulfillmentRequest) SetFulfillmentData(v FulfillmentData) {
	o.FulfillmentData = &v
}

// GetFulfillmentProvider returns the FulfillmentProvider field value if set, zero value otherwise.
func (o *FulfillmentRequest) GetFulfillmentProvider() string {
	if o == nil || o.FulfillmentProvider == nil {
		var ret string
		return ret
	}
	return *o.FulfillmentProvider
}

// GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentRequest) GetFulfillmentProviderOk() (*string, bool) {
	if o == nil || o.FulfillmentProvider == nil {
		return nil, false
	}
	return o.FulfillmentProvider, true
}

// HasFulfillmentProvider returns a boolean if a field has been set.
func (o *FulfillmentRequest) HasFulfillmentProvider() bool {
	if o != nil && o.FulfillmentProvider != nil {
		return true
	}

	return false
}

// SetFulfillmentProvider gets a reference to the given string and assigns it to the FulfillmentProvider field.
func (o *FulfillmentRequest) SetFulfillmentProvider(v string) {
	o.FulfillmentProvider = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FulfillmentRequest) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentRequest) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FulfillmentRequest) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FulfillmentRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o FulfillmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FulfillmentData != nil {
		toSerialize["fulfillmentData"] = o.FulfillmentData
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

func (o *FulfillmentRequest) UnmarshalJSON(bytes []byte) (err error) {
	varFulfillmentRequest := _FulfillmentRequest{}

	err = json.Unmarshal(bytes, &varFulfillmentRequest)
	if err == nil {
		*o = FulfillmentRequest(varFulfillmentRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "fulfillmentData")
		delete(additionalProperties, "fulfillmentProvider")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFulfillmentRequest struct {
	value *FulfillmentRequest
	isSet bool
}

func (v NullableFulfillmentRequest) Get() *FulfillmentRequest {
	return v.value
}

func (v *NullableFulfillmentRequest) Set(val *FulfillmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentRequest(val *FulfillmentRequest) *NullableFulfillmentRequest {
	return &NullableFulfillmentRequest{value: val, isSet: true}
}

func (v NullableFulfillmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

