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

// checks if the IDVTokenEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IDVTokenEndpoint{}

// IDVTokenEndpoint IDV token endpoint
type IDVTokenEndpoint struct {
	Binding *string `json:"binding,omitempty"`
	// URL of the IDV `token` endpoint
	Url                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IDVTokenEndpoint IDVTokenEndpoint

// NewIDVTokenEndpoint instantiates a new IDVTokenEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDVTokenEndpoint() *IDVTokenEndpoint {
	this := IDVTokenEndpoint{}
	return &this
}

// NewIDVTokenEndpointWithDefaults instantiates a new IDVTokenEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDVTokenEndpointWithDefaults() *IDVTokenEndpoint {
	this := IDVTokenEndpoint{}
	return &this
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *IDVTokenEndpoint) GetBinding() string {
	if o == nil || IsNil(o.Binding) {
		var ret string
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVTokenEndpoint) GetBindingOk() (*string, bool) {
	if o == nil || IsNil(o.Binding) {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *IDVTokenEndpoint) HasBinding() bool {
	if o != nil && !IsNil(o.Binding) {
		return true
	}

	return false
}

// SetBinding gets a reference to the given string and assigns it to the Binding field.
func (o *IDVTokenEndpoint) SetBinding(v string) {
	o.Binding = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IDVTokenEndpoint) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVTokenEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IDVTokenEndpoint) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *IDVTokenEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o IDVTokenEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IDVTokenEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Binding) {
		toSerialize["binding"] = o.Binding
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IDVTokenEndpoint) UnmarshalJSON(data []byte) (err error) {
	varIDVTokenEndpoint := _IDVTokenEndpoint{}

	err = json.Unmarshal(data, &varIDVTokenEndpoint)

	if err != nil {
		return err
	}

	*o = IDVTokenEndpoint(varIDVTokenEndpoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "binding")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIDVTokenEndpoint struct {
	value *IDVTokenEndpoint
	isSet bool
}

func (v NullableIDVTokenEndpoint) Get() *IDVTokenEndpoint {
	return v.value
}

func (v *NullableIDVTokenEndpoint) Set(val *IDVTokenEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableIDVTokenEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableIDVTokenEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDVTokenEndpoint(val *IDVTokenEndpoint) *NullableIDVTokenEndpoint {
	return &NullableIDVTokenEndpoint{value: val, isSet: true}
}

func (v NullableIDVTokenEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDVTokenEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
