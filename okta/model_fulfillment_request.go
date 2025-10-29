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

// checks if the FulfillmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentRequest{}

// FulfillmentRequest Fulfillment request
type FulfillmentRequest struct {
	// List of fulfillment order details
	FulfillmentData []FulfillmentDataOrderDetails `json:"fulfillmentData,omitempty"`
	// Name of the fulfillment provider for the WebAuthn preregistration factor
	FulfillmentProvider *string `json:"fulfillmentProvider,omitempty"`
	// ID of an existing Okta user
	UserId               *string `json:"userId,omitempty"`
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
func (o *FulfillmentRequest) GetFulfillmentData() []FulfillmentDataOrderDetails {
	if o == nil || IsNil(o.FulfillmentData) {
		var ret []FulfillmentDataOrderDetails
		return ret
	}
	return o.FulfillmentData
}

// GetFulfillmentDataOk returns a tuple with the FulfillmentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentRequest) GetFulfillmentDataOk() ([]FulfillmentDataOrderDetails, bool) {
	if o == nil || IsNil(o.FulfillmentData) {
		return nil, false
	}
	return o.FulfillmentData, true
}

// HasFulfillmentData returns a boolean if a field has been set.
func (o *FulfillmentRequest) HasFulfillmentData() bool {
	if o != nil && !IsNil(o.FulfillmentData) {
		return true
	}

	return false
}

// SetFulfillmentData gets a reference to the given []FulfillmentDataOrderDetails and assigns it to the FulfillmentData field.
func (o *FulfillmentRequest) SetFulfillmentData(v []FulfillmentDataOrderDetails) {
	o.FulfillmentData = v
}

// GetFulfillmentProvider returns the FulfillmentProvider field value if set, zero value otherwise.
func (o *FulfillmentRequest) GetFulfillmentProvider() string {
	if o == nil || IsNil(o.FulfillmentProvider) {
		var ret string
		return ret
	}
	return *o.FulfillmentProvider
}

// GetFulfillmentProviderOk returns a tuple with the FulfillmentProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentRequest) GetFulfillmentProviderOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentProvider) {
		return nil, false
	}
	return o.FulfillmentProvider, true
}

// HasFulfillmentProvider returns a boolean if a field has been set.
func (o *FulfillmentRequest) HasFulfillmentProvider() bool {
	if o != nil && !IsNil(o.FulfillmentProvider) {
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
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FulfillmentRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FulfillmentRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o FulfillmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FulfillmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FulfillmentData) {
		toSerialize["fulfillmentData"] = o.FulfillmentData
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

func (o *FulfillmentRequest) UnmarshalJSON(data []byte) (err error) {
	varFulfillmentRequest := _FulfillmentRequest{}

	err = json.Unmarshal(data, &varFulfillmentRequest)

	if err != nil {
		return err
	}

	*o = FulfillmentRequest(varFulfillmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fulfillmentData")
		delete(additionalProperties, "fulfillmentProvider")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
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
