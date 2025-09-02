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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// SAMLPayLoadData struct for SAMLPayLoadData
type SAMLPayLoadData struct {
	Context *SAMLPayLoadDataContext `json:"context,omitempty"`
	Assertion *SAMLPayLoadDataAssertion `json:"assertion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadData SAMLPayLoadData

// NewSAMLPayLoadData instantiates a new SAMLPayLoadData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadData() *SAMLPayLoadData {
	this := SAMLPayLoadData{}
	return &this
}

// NewSAMLPayLoadDataWithDefaults instantiates a new SAMLPayLoadData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataWithDefaults() *SAMLPayLoadData {
	this := SAMLPayLoadData{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *SAMLPayLoadData) GetContext() SAMLPayLoadDataContext {
	if o == nil || o.Context == nil {
		var ret SAMLPayLoadDataContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadData) GetContextOk() (*SAMLPayLoadDataContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *SAMLPayLoadData) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given SAMLPayLoadDataContext and assigns it to the Context field.
func (o *SAMLPayLoadData) SetContext(v SAMLPayLoadDataContext) {
	o.Context = &v
}

// GetAssertion returns the Assertion field value if set, zero value otherwise.
func (o *SAMLPayLoadData) GetAssertion() SAMLPayLoadDataAssertion {
	if o == nil || o.Assertion == nil {
		var ret SAMLPayLoadDataAssertion
		return ret
	}
	return *o.Assertion
}

// GetAssertionOk returns a tuple with the Assertion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadData) GetAssertionOk() (*SAMLPayLoadDataAssertion, bool) {
	if o == nil || o.Assertion == nil {
		return nil, false
	}
	return o.Assertion, true
}

// HasAssertion returns a boolean if a field has been set.
func (o *SAMLPayLoadData) HasAssertion() bool {
	if o != nil && o.Assertion != nil {
		return true
	}

	return false
}

// SetAssertion gets a reference to the given SAMLPayLoadDataAssertion and assigns it to the Assertion field.
func (o *SAMLPayLoadData) SetAssertion(v SAMLPayLoadDataAssertion) {
	o.Assertion = &v
}

func (o SAMLPayLoadData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Assertion != nil {
		toSerialize["assertion"] = o.Assertion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SAMLPayLoadData) UnmarshalJSON(bytes []byte) (err error) {
	varSAMLPayLoadData := _SAMLPayLoadData{}

	err = json.Unmarshal(bytes, &varSAMLPayLoadData)
	if err == nil {
		*o = SAMLPayLoadData(varSAMLPayLoadData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "assertion")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSAMLPayLoadData struct {
	value *SAMLPayLoadData
	isSet bool
}

func (v NullableSAMLPayLoadData) Get() *SAMLPayLoadData {
	return v.value
}

func (v *NullableSAMLPayLoadData) Set(val *SAMLPayLoadData) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadData) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadData(val *SAMLPayLoadData) *NullableSAMLPayLoadData {
	return &NullableSAMLPayLoadData{value: val, isSet: true}
}

func (v NullableSAMLPayLoadData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

