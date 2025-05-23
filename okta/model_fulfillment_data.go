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

// FulfillmentData Fulfillment provider details
type FulfillmentData struct {
	// ID for the set of custom configurations of the requested Factor
	CustomizationId *string `json:"customizationId,omitempty"`
	// ID for the specific inventory bucket of the requested Factor
	InventoryProductId *string `json:"inventoryProductId,omitempty"`
	// ID for the make and model of the requested Factor
	ProductId *string `json:"productId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FulfillmentData FulfillmentData

// NewFulfillmentData instantiates a new FulfillmentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentData() *FulfillmentData {
	this := FulfillmentData{}
	return &this
}

// NewFulfillmentDataWithDefaults instantiates a new FulfillmentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentDataWithDefaults() *FulfillmentData {
	this := FulfillmentData{}
	return &this
}

// GetCustomizationId returns the CustomizationId field value if set, zero value otherwise.
func (o *FulfillmentData) GetCustomizationId() string {
	if o == nil || o.CustomizationId == nil {
		var ret string
		return ret
	}
	return *o.CustomizationId
}

// GetCustomizationIdOk returns a tuple with the CustomizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentData) GetCustomizationIdOk() (*string, bool) {
	if o == nil || o.CustomizationId == nil {
		return nil, false
	}
	return o.CustomizationId, true
}

// HasCustomizationId returns a boolean if a field has been set.
func (o *FulfillmentData) HasCustomizationId() bool {
	if o != nil && o.CustomizationId != nil {
		return true
	}

	return false
}

// SetCustomizationId gets a reference to the given string and assigns it to the CustomizationId field.
func (o *FulfillmentData) SetCustomizationId(v string) {
	o.CustomizationId = &v
}

// GetInventoryProductId returns the InventoryProductId field value if set, zero value otherwise.
func (o *FulfillmentData) GetInventoryProductId() string {
	if o == nil || o.InventoryProductId == nil {
		var ret string
		return ret
	}
	return *o.InventoryProductId
}

// GetInventoryProductIdOk returns a tuple with the InventoryProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentData) GetInventoryProductIdOk() (*string, bool) {
	if o == nil || o.InventoryProductId == nil {
		return nil, false
	}
	return o.InventoryProductId, true
}

// HasInventoryProductId returns a boolean if a field has been set.
func (o *FulfillmentData) HasInventoryProductId() bool {
	if o != nil && o.InventoryProductId != nil {
		return true
	}

	return false
}

// SetInventoryProductId gets a reference to the given string and assigns it to the InventoryProductId field.
func (o *FulfillmentData) SetInventoryProductId(v string) {
	o.InventoryProductId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *FulfillmentData) GetProductId() string {
	if o == nil || o.ProductId == nil {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentData) GetProductIdOk() (*string, bool) {
	if o == nil || o.ProductId == nil {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *FulfillmentData) HasProductId() bool {
	if o != nil && o.ProductId != nil {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *FulfillmentData) SetProductId(v string) {
	o.ProductId = &v
}

func (o FulfillmentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomizationId != nil {
		toSerialize["customizationId"] = o.CustomizationId
	}
	if o.InventoryProductId != nil {
		toSerialize["inventoryProductId"] = o.InventoryProductId
	}
	if o.ProductId != nil {
		toSerialize["productId"] = o.ProductId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FulfillmentData) UnmarshalJSON(bytes []byte) (err error) {
	varFulfillmentData := _FulfillmentData{}

	err = json.Unmarshal(bytes, &varFulfillmentData)
	if err == nil {
		*o = FulfillmentData(varFulfillmentData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "customizationId")
		delete(additionalProperties, "inventoryProductId")
		delete(additionalProperties, "productId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFulfillmentData struct {
	value *FulfillmentData
	isSet bool
}

func (v NullableFulfillmentData) Get() *FulfillmentData {
	return v.value
}

func (v *NullableFulfillmentData) Set(val *FulfillmentData) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentData) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentData(val *FulfillmentData) *NullableFulfillmentData {
	return &NullableFulfillmentData{value: val, isSet: true}
}

func (v NullableFulfillmentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillmentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

