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

// RegistrationInlineHookSSRDataAllOfDataContext struct for RegistrationInlineHookSSRDataAllOfDataContext
type RegistrationInlineHookSSRDataAllOfDataContext struct {
	Request *InlineHookRequestObject `json:"request,omitempty"`
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
	if o == nil || o.Request == nil {
		var ret InlineHookRequestObject
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRDataAllOfDataContext) GetRequestOk() (*InlineHookRequestObject, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRDataAllOfDataContext) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineHookRequestObject and assigns it to the Request field.
func (o *RegistrationInlineHookSSRDataAllOfDataContext) SetRequest(v InlineHookRequestObject) {
	o.Request = &v
}

func (o RegistrationInlineHookSSRDataAllOfDataContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RegistrationInlineHookSSRDataAllOfDataContext) UnmarshalJSON(bytes []byte) (err error) {
	varRegistrationInlineHookSSRDataAllOfDataContext := _RegistrationInlineHookSSRDataAllOfDataContext{}

	err = json.Unmarshal(bytes, &varRegistrationInlineHookSSRDataAllOfDataContext)
	if err == nil {
		*o = RegistrationInlineHookSSRDataAllOfDataContext(varRegistrationInlineHookSSRDataAllOfDataContext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "request")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

