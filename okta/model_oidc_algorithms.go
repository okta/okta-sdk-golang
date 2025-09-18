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

// checks if the OidcAlgorithms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OidcAlgorithms{}

// OidcAlgorithms struct for OidcAlgorithms
type OidcAlgorithms struct {
	Request              *OidcRequestAlgorithm `json:"request,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OidcAlgorithms OidcAlgorithms

// NewOidcAlgorithms instantiates a new OidcAlgorithms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcAlgorithms() *OidcAlgorithms {
	this := OidcAlgorithms{}
	return &this
}

// NewOidcAlgorithmsWithDefaults instantiates a new OidcAlgorithms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcAlgorithmsWithDefaults() *OidcAlgorithms {
	this := OidcAlgorithms{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *OidcAlgorithms) GetRequest() OidcRequestAlgorithm {
	if o == nil || IsNil(o.Request) {
		var ret OidcRequestAlgorithm
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcAlgorithms) GetRequestOk() (*OidcRequestAlgorithm, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *OidcAlgorithms) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given OidcRequestAlgorithm and assigns it to the Request field.
func (o *OidcAlgorithms) SetRequest(v OidcRequestAlgorithm) {
	o.Request = &v
}

func (o OidcAlgorithms) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OidcAlgorithms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OidcAlgorithms) UnmarshalJSON(data []byte) (err error) {
	varOidcAlgorithms := _OidcAlgorithms{}

	err = json.Unmarshal(data, &varOidcAlgorithms)

	if err != nil {
		return err
	}

	*o = OidcAlgorithms(varOidcAlgorithms)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOidcAlgorithms struct {
	value *OidcAlgorithms
	isSet bool
}

func (v NullableOidcAlgorithms) Get() *OidcAlgorithms {
	return v.value
}

func (v *NullableOidcAlgorithms) Set(val *OidcAlgorithms) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcAlgorithms) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcAlgorithms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcAlgorithms(val *OidcAlgorithms) *NullableOidcAlgorithms {
	return &NullableOidcAlgorithms{value: val, isSet: true}
}

func (v NullableOidcAlgorithms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcAlgorithms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
