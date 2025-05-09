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

// ProtocolAlgorithms struct for ProtocolAlgorithms
type ProtocolAlgorithms struct {
	Request *ProtocolAlgorithmType `json:"request,omitempty"`
	Response *ProtocolAlgorithmType `json:"response,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolAlgorithms ProtocolAlgorithms

// NewProtocolAlgorithms instantiates a new ProtocolAlgorithms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolAlgorithms() *ProtocolAlgorithms {
	this := ProtocolAlgorithms{}
	return &this
}

// NewProtocolAlgorithmsWithDefaults instantiates a new ProtocolAlgorithms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolAlgorithmsWithDefaults() *ProtocolAlgorithms {
	this := ProtocolAlgorithms{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *ProtocolAlgorithms) GetRequest() ProtocolAlgorithmType {
	if o == nil || o.Request == nil {
		var ret ProtocolAlgorithmType
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolAlgorithms) GetRequestOk() (*ProtocolAlgorithmType, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *ProtocolAlgorithms) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given ProtocolAlgorithmType and assigns it to the Request field.
func (o *ProtocolAlgorithms) SetRequest(v ProtocolAlgorithmType) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ProtocolAlgorithms) GetResponse() ProtocolAlgorithmType {
	if o == nil || o.Response == nil {
		var ret ProtocolAlgorithmType
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolAlgorithms) GetResponseOk() (*ProtocolAlgorithmType, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ProtocolAlgorithms) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ProtocolAlgorithmType and assigns it to the Response field.
func (o *ProtocolAlgorithms) SetResponse(v ProtocolAlgorithmType) {
	o.Response = &v
}

func (o ProtocolAlgorithms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProtocolAlgorithms) UnmarshalJSON(bytes []byte) (err error) {
	varProtocolAlgorithms := _ProtocolAlgorithms{}

	err = json.Unmarshal(bytes, &varProtocolAlgorithms)
	if err == nil {
		*o = ProtocolAlgorithms(varProtocolAlgorithms)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "response")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProtocolAlgorithms struct {
	value *ProtocolAlgorithms
	isSet bool
}

func (v NullableProtocolAlgorithms) Get() *ProtocolAlgorithms {
	return v.value
}

func (v *NullableProtocolAlgorithms) Set(val *ProtocolAlgorithms) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolAlgorithms) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolAlgorithms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolAlgorithms(val *ProtocolAlgorithms) *NullableProtocolAlgorithms {
	return &NullableProtocolAlgorithms{value: val, isSet: true}
}

func (v NullableProtocolAlgorithms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolAlgorithms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

