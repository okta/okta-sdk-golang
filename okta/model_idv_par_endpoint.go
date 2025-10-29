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

// checks if the IDVParEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IDVParEndpoint{}

// IDVParEndpoint IDV [PAR](https://datatracker.ietf.org/doc/html/rfc9126) endpoint
type IDVParEndpoint struct {
	Binding *string `json:"binding,omitempty"`
	// URL of the `par` endpoint of the IDV vendor
	Url                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IDVParEndpoint IDVParEndpoint

// NewIDVParEndpoint instantiates a new IDVParEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDVParEndpoint() *IDVParEndpoint {
	this := IDVParEndpoint{}
	return &this
}

// NewIDVParEndpointWithDefaults instantiates a new IDVParEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDVParEndpointWithDefaults() *IDVParEndpoint {
	this := IDVParEndpoint{}
	return &this
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *IDVParEndpoint) GetBinding() string {
	if o == nil || IsNil(o.Binding) {
		var ret string
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVParEndpoint) GetBindingOk() (*string, bool) {
	if o == nil || IsNil(o.Binding) {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *IDVParEndpoint) HasBinding() bool {
	if o != nil && !IsNil(o.Binding) {
		return true
	}

	return false
}

// SetBinding gets a reference to the given string and assigns it to the Binding field.
func (o *IDVParEndpoint) SetBinding(v string) {
	o.Binding = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *IDVParEndpoint) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVParEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *IDVParEndpoint) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *IDVParEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o IDVParEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IDVParEndpoint) ToMap() (map[string]interface{}, error) {
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

func (o *IDVParEndpoint) UnmarshalJSON(data []byte) (err error) {
	varIDVParEndpoint := _IDVParEndpoint{}

	err = json.Unmarshal(data, &varIDVParEndpoint)

	if err != nil {
		return err
	}

	*o = IDVParEndpoint(varIDVParEndpoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "binding")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIDVParEndpoint struct {
	value *IDVParEndpoint
	isSet bool
}

func (v NullableIDVParEndpoint) Get() *IDVParEndpoint {
	return v.value
}

func (v *NullableIDVParEndpoint) Set(val *IDVParEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableIDVParEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableIDVParEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDVParEndpoint(val *IDVParEndpoint) *NullableIDVParEndpoint {
	return &NullableIDVParEndpoint{value: val, isSet: true}
}

func (v NullableIDVParEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDVParEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
