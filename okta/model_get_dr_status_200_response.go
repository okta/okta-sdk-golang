/*
Okta Admin Management API

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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the GetDRStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDRStatus200Response{}

// GetDRStatus200Response struct for GetDRStatus200Response
type GetDRStatus200Response struct {
	// List of domains and their disaster recovery status
	Status               []DRStatusItem `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDRStatus200Response GetDRStatus200Response

// NewGetDRStatus200Response instantiates a new GetDRStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDRStatus200Response() *GetDRStatus200Response {
	this := GetDRStatus200Response{}
	return &this
}

// NewGetDRStatus200ResponseWithDefaults instantiates a new GetDRStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDRStatus200ResponseWithDefaults() *GetDRStatus200Response {
	this := GetDRStatus200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetDRStatus200Response) GetStatus() []DRStatusItem {
	if o == nil || IsNil(o.Status) {
		var ret []DRStatusItem
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDRStatus200Response) GetStatusOk() ([]DRStatusItem, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetDRStatus200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []DRStatusItem and assigns it to the Status field.
func (o *GetDRStatus200Response) SetStatus(v []DRStatusItem) {
	o.Status = v
}

func (o GetDRStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDRStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetDRStatus200Response) UnmarshalJSON(data []byte) (err error) {
	varGetDRStatus200Response := _GetDRStatus200Response{}

	err = json.Unmarshal(data, &varGetDRStatus200Response)

	if err != nil {
		return err
	}

	*o = GetDRStatus200Response(varGetDRStatus200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDRStatus200Response struct {
	value *GetDRStatus200Response
	isSet bool
}

func (v NullableGetDRStatus200Response) Get() *GetDRStatus200Response {
	return v.value
}

func (v *NullableGetDRStatus200Response) Set(val *GetDRStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDRStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDRStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDRStatus200Response(val *GetDRStatus200Response) *NullableGetDRStatus200Response {
	return &NullableGetDRStatus200Response{value: val, isSet: true}
}

func (v NullableGetDRStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDRStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
