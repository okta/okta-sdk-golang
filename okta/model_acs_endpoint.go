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
	"fmt"
)

// checks if the AcsEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcsEndpoint{}

// AcsEndpoint An array of ACS endpoints. You can configure a maximum of 100 endpoints.
type AcsEndpoint struct {
	// Index of the URL in the array of ACS endpoints
	Index int32 `json:"index"`
	// URL of the ACS
	Url                  string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _AcsEndpoint AcsEndpoint

// NewAcsEndpoint instantiates a new AcsEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcsEndpoint(index int32, url string) *AcsEndpoint {
	this := AcsEndpoint{}
	this.Index = index
	this.Url = url
	return &this
}

// NewAcsEndpointWithDefaults instantiates a new AcsEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcsEndpointWithDefaults() *AcsEndpoint {
	this := AcsEndpoint{}
	return &this
}

// GetIndex returns the Index field value
func (o *AcsEndpoint) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *AcsEndpoint) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *AcsEndpoint) SetIndex(v int32) {
	o.Index = v
}

// GetUrl returns the Url field value
func (o *AcsEndpoint) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AcsEndpoint) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AcsEndpoint) SetUrl(v string) {
	o.Url = v
}

func (o AcsEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcsEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcsEndpoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAcsEndpoint := _AcsEndpoint{}

	err = json.Unmarshal(data, &varAcsEndpoint)

	if err != nil {
		return err
	}

	*o = AcsEndpoint(varAcsEndpoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "index")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcsEndpoint struct {
	value *AcsEndpoint
	isSet bool
}

func (v NullableAcsEndpoint) Get() *AcsEndpoint {
	return v.value
}

func (v *NullableAcsEndpoint) Set(val *AcsEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAcsEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAcsEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcsEndpoint(val *AcsEndpoint) *NullableAcsEndpoint {
	return &NullableAcsEndpoint{value: val, isSet: true}
}

func (v NullableAcsEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcsEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
