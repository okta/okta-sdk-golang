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

// checks if the SamlAlgorithms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlAlgorithms{}

// SamlAlgorithms Settings for signing and verifying SAML messages
type SamlAlgorithms struct {
	Request              *SamlRequestAlgorithm  `json:"request,omitempty"`
	Response             *SamlResponseAlgorithm `json:"response,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlAlgorithms SamlAlgorithms

// NewSamlAlgorithms instantiates a new SamlAlgorithms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlAlgorithms() *SamlAlgorithms {
	this := SamlAlgorithms{}
	return &this
}

// NewSamlAlgorithmsWithDefaults instantiates a new SamlAlgorithms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAlgorithmsWithDefaults() *SamlAlgorithms {
	this := SamlAlgorithms{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SamlAlgorithms) GetRequest() SamlRequestAlgorithm {
	if o == nil || IsNil(o.Request) {
		var ret SamlRequestAlgorithm
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAlgorithms) GetRequestOk() (*SamlRequestAlgorithm, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SamlAlgorithms) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given SamlRequestAlgorithm and assigns it to the Request field.
func (o *SamlAlgorithms) SetRequest(v SamlRequestAlgorithm) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SamlAlgorithms) GetResponse() SamlResponseAlgorithm {
	if o == nil || IsNil(o.Response) {
		var ret SamlResponseAlgorithm
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAlgorithms) GetResponseOk() (*SamlResponseAlgorithm, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SamlAlgorithms) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given SamlResponseAlgorithm and assigns it to the Response field.
func (o *SamlAlgorithms) SetResponse(v SamlResponseAlgorithm) {
	o.Response = &v
}

func (o SamlAlgorithms) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlAlgorithms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlAlgorithms) UnmarshalJSON(data []byte) (err error) {
	varSamlAlgorithms := _SamlAlgorithms{}

	err = json.Unmarshal(data, &varSamlAlgorithms)

	if err != nil {
		return err
	}

	*o = SamlAlgorithms(varSamlAlgorithms)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "response")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlAlgorithms struct {
	value *SamlAlgorithms
	isSet bool
}

func (v NullableSamlAlgorithms) Get() *SamlAlgorithms {
	return v.value
}

func (v *NullableSamlAlgorithms) Set(val *SamlAlgorithms) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlAlgorithms) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlAlgorithms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlAlgorithms(val *SamlAlgorithms) *NullableSamlAlgorithms {
	return &NullableSamlAlgorithms{value: val, isSet: true}
}

func (v NullableSamlAlgorithms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlAlgorithms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
