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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the FulfillmentDataOrderDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentDataOrderDetails{}

// FulfillmentDataOrderDetails Information about the fulfillment order that includes the factor’s make and model, the custom configuration of the factor, and inventory details.
type FulfillmentDataOrderDetails struct {
	// ID for the set of custom configurations of the requested factor
	CustomizationId *string `json:"customizationId,omitempty"`
	// ID for the specific inventory bucket of the requested factor
	InventoryProductId *string `json:"inventoryProductId,omitempty"`
	// ID for the make and model of the requested factor
	ProductId            *string `json:"productId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FulfillmentDataOrderDetails FulfillmentDataOrderDetails

// NewFulfillmentDataOrderDetails instantiates a new FulfillmentDataOrderDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentDataOrderDetails() *FulfillmentDataOrderDetails {
	this := FulfillmentDataOrderDetails{}
	return &this
}

// NewFulfillmentDataOrderDetailsWithDefaults instantiates a new FulfillmentDataOrderDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentDataOrderDetailsWithDefaults() *FulfillmentDataOrderDetails {
	this := FulfillmentDataOrderDetails{}
	return &this
}

// GetCustomizationId returns the CustomizationId field value if set, zero value otherwise.
func (o *FulfillmentDataOrderDetails) GetCustomizationId() string {
	if o == nil || IsNil(o.CustomizationId) {
		var ret string
		return ret
	}
	return *o.CustomizationId
}

// GetCustomizationIdOk returns a tuple with the CustomizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentDataOrderDetails) GetCustomizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomizationId) {
		return nil, false
	}
	return o.CustomizationId, true
}

// HasCustomizationId returns a boolean if a field has been set.
func (o *FulfillmentDataOrderDetails) HasCustomizationId() bool {
	if o != nil && !IsNil(o.CustomizationId) {
		return true
	}

	return false
}

// SetCustomizationId gets a reference to the given string and assigns it to the CustomizationId field.
func (o *FulfillmentDataOrderDetails) SetCustomizationId(v string) {
	o.CustomizationId = &v
}

// GetInventoryProductId returns the InventoryProductId field value if set, zero value otherwise.
func (o *FulfillmentDataOrderDetails) GetInventoryProductId() string {
	if o == nil || IsNil(o.InventoryProductId) {
		var ret string
		return ret
	}
	return *o.InventoryProductId
}

// GetInventoryProductIdOk returns a tuple with the InventoryProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentDataOrderDetails) GetInventoryProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventoryProductId) {
		return nil, false
	}
	return o.InventoryProductId, true
}

// HasInventoryProductId returns a boolean if a field has been set.
func (o *FulfillmentDataOrderDetails) HasInventoryProductId() bool {
	if o != nil && !IsNil(o.InventoryProductId) {
		return true
	}

	return false
}

// SetInventoryProductId gets a reference to the given string and assigns it to the InventoryProductId field.
func (o *FulfillmentDataOrderDetails) SetInventoryProductId(v string) {
	o.InventoryProductId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *FulfillmentDataOrderDetails) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentDataOrderDetails) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *FulfillmentDataOrderDetails) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *FulfillmentDataOrderDetails) SetProductId(v string) {
	o.ProductId = &v
}

func (o FulfillmentDataOrderDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FulfillmentDataOrderDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomizationId) {
		toSerialize["customizationId"] = o.CustomizationId
	}
	if !IsNil(o.InventoryProductId) {
		toSerialize["inventoryProductId"] = o.InventoryProductId
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FulfillmentDataOrderDetails) UnmarshalJSON(data []byte) (err error) {
	varFulfillmentDataOrderDetails := _FulfillmentDataOrderDetails{}

	err = json.Unmarshal(data, &varFulfillmentDataOrderDetails)

	if err != nil {
		return err
	}

	*o = FulfillmentDataOrderDetails(varFulfillmentDataOrderDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customizationId")
		delete(additionalProperties, "inventoryProductId")
		delete(additionalProperties, "productId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFulfillmentDataOrderDetails struct {
	value *FulfillmentDataOrderDetails
	isSet bool
}

func (v NullableFulfillmentDataOrderDetails) Get() *FulfillmentDataOrderDetails {
	return v.value
}

func (v *NullableFulfillmentDataOrderDetails) Set(val *FulfillmentDataOrderDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentDataOrderDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentDataOrderDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentDataOrderDetails(val *FulfillmentDataOrderDetails) *NullableFulfillmentDataOrderDetails {
	return &NullableFulfillmentDataOrderDetails{value: val, isSet: true}
}

func (v NullableFulfillmentDataOrderDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillmentDataOrderDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
