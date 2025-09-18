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

// checks if the DRStatusItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DRStatusItem{}

// DRStatusItem Status whether a domain has been failed over or not
type DRStatusItem struct {
	// Domain for your org
	Domain *string `json:"domain,omitempty"`
	// Indicates if the domain has been failed over
	IsFailedOver         *bool `json:"isFailedOver,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DRStatusItem DRStatusItem

// NewDRStatusItem instantiates a new DRStatusItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDRStatusItem() *DRStatusItem {
	this := DRStatusItem{}
	return &this
}

// NewDRStatusItemWithDefaults instantiates a new DRStatusItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDRStatusItemWithDefaults() *DRStatusItem {
	this := DRStatusItem{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DRStatusItem) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DRStatusItem) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DRStatusItem) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DRStatusItem) SetDomain(v string) {
	o.Domain = &v
}

// GetIsFailedOver returns the IsFailedOver field value if set, zero value otherwise.
func (o *DRStatusItem) GetIsFailedOver() bool {
	if o == nil || IsNil(o.IsFailedOver) {
		var ret bool
		return ret
	}
	return *o.IsFailedOver
}

// GetIsFailedOverOk returns a tuple with the IsFailedOver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DRStatusItem) GetIsFailedOverOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFailedOver) {
		return nil, false
	}
	return o.IsFailedOver, true
}

// HasIsFailedOver returns a boolean if a field has been set.
func (o *DRStatusItem) HasIsFailedOver() bool {
	if o != nil && !IsNil(o.IsFailedOver) {
		return true
	}

	return false
}

// SetIsFailedOver gets a reference to the given bool and assigns it to the IsFailedOver field.
func (o *DRStatusItem) SetIsFailedOver(v bool) {
	o.IsFailedOver = &v
}

func (o DRStatusItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DRStatusItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.IsFailedOver) {
		toSerialize["isFailedOver"] = o.IsFailedOver
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DRStatusItem) UnmarshalJSON(data []byte) (err error) {
	varDRStatusItem := _DRStatusItem{}

	err = json.Unmarshal(data, &varDRStatusItem)

	if err != nil {
		return err
	}

	*o = DRStatusItem(varDRStatusItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "isFailedOver")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDRStatusItem struct {
	value *DRStatusItem
	isSet bool
}

func (v NullableDRStatusItem) Get() *DRStatusItem {
	return v.value
}

func (v *NullableDRStatusItem) Set(val *DRStatusItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDRStatusItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDRStatusItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDRStatusItem(val *DRStatusItem) *NullableDRStatusItem {
	return &NullableDRStatusItem{value: val, isSet: true}
}

func (v NullableDRStatusItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDRStatusItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
