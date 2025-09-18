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

// checks if the RegistrationInlineHookSSRDataAllOfDataContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInlineHookSSRDataAllOfDataContext{}

// RegistrationInlineHookSSRDataAllOfDataContext struct for RegistrationInlineHookSSRDataAllOfDataContext
type RegistrationInlineHookSSRDataAllOfDataContext struct {
	Request              *InlineHookRequestObject `json:"request,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookSSRDataAllOfDataContext RegistrationInlineHookSSRDataAllOfDataContext

// NewRegistrationInlineHookSSRDataAllOfDataContext instantiates a new RegistrationInlineHookSSRDataAllOfDataContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookSSRDataAllOfDataContext() *RegistrationInlineHookSSRDataAllOfDataContext {
	this := RegistrationInlineHookSSRDataAllOfDataContext{}
	return &this
}

// NewRegistrationInlineHookSSRDataAllOfDataContextWithDefaults instantiates a new RegistrationInlineHookSSRDataAllOfDataContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookSSRDataAllOfDataContextWithDefaults() *RegistrationInlineHookSSRDataAllOfDataContext {
	this := RegistrationInlineHookSSRDataAllOfDataContext{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRDataAllOfDataContext) GetRequest() InlineHookRequestObject {
	if o == nil || IsNil(o.Request) {
		var ret InlineHookRequestObject
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRDataAllOfDataContext) GetRequestOk() (*InlineHookRequestObject, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRDataAllOfDataContext) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineHookRequestObject and assigns it to the Request field.
func (o *RegistrationInlineHookSSRDataAllOfDataContext) SetRequest(v InlineHookRequestObject) {
	o.Request = &v
}

func (o RegistrationInlineHookSSRDataAllOfDataContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInlineHookSSRDataAllOfDataContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationInlineHookSSRDataAllOfDataContext) UnmarshalJSON(data []byte) (err error) {
	varRegistrationInlineHookSSRDataAllOfDataContext := _RegistrationInlineHookSSRDataAllOfDataContext{}

	err = json.Unmarshal(data, &varRegistrationInlineHookSSRDataAllOfDataContext)

	if err != nil {
		return err
	}

	*o = RegistrationInlineHookSSRDataAllOfDataContext(varRegistrationInlineHookSSRDataAllOfDataContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationInlineHookSSRDataAllOfDataContext struct {
	value *RegistrationInlineHookSSRDataAllOfDataContext
	isSet bool
}

func (v NullableRegistrationInlineHookSSRDataAllOfDataContext) Get() *RegistrationInlineHookSSRDataAllOfDataContext {
	return v.value
}

func (v *NullableRegistrationInlineHookSSRDataAllOfDataContext) Set(val *RegistrationInlineHookSSRDataAllOfDataContext) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookSSRDataAllOfDataContext) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookSSRDataAllOfDataContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookSSRDataAllOfDataContext(val *RegistrationInlineHookSSRDataAllOfDataContext) *NullableRegistrationInlineHookSSRDataAllOfDataContext {
	return &NullableRegistrationInlineHookSSRDataAllOfDataContext{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookSSRDataAllOfDataContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookSSRDataAllOfDataContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
