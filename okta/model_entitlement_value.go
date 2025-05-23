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

// EntitlementValue struct for EntitlementValue
type EntitlementValue struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValue EntitlementValue

// NewEntitlementValue instantiates a new EntitlementValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValue() *EntitlementValue {
	this := EntitlementValue{}
	return &this
}

// NewEntitlementValueWithDefaults instantiates a new EntitlementValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValueWithDefaults() *EntitlementValue {
	this := EntitlementValue{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementValue) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValue) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementValue) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementValue) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementValue) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValue) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementValue) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementValue) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValue) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EntitlementValue) SetValue(v string) {
	o.Value = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EntitlementValue) GetLinks() map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValue) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EntitlementValue) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *EntitlementValue) SetLinks(v map[string]interface{}) {
	o.Links = v
}

func (o EntitlementValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementValue) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementValue := _EntitlementValue{}

	err = json.Unmarshal(bytes, &varEntitlementValue)
	if err == nil {
		*o = EntitlementValue(varEntitlementValue)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementValue struct {
	value *EntitlementValue
	isSet bool
}

func (v NullableEntitlementValue) Get() *EntitlementValue {
	return v.value
}

func (v *NullableEntitlementValue) Set(val *EntitlementValue) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValue) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValue(val *EntitlementValue) *NullableEntitlementValue {
	return &NullableEntitlementValue{value: val, isSet: true}
}

func (v NullableEntitlementValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

