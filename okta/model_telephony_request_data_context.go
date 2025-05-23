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

// TelephonyRequestDataContext struct for TelephonyRequestDataContext
type TelephonyRequestDataContext struct {
	Request *InlineHookRequestObject `json:"request,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelephonyRequestDataContext TelephonyRequestDataContext

// NewTelephonyRequestDataContext instantiates a new TelephonyRequestDataContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephonyRequestDataContext() *TelephonyRequestDataContext {
	this := TelephonyRequestDataContext{}
	return &this
}

// NewTelephonyRequestDataContextWithDefaults instantiates a new TelephonyRequestDataContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephonyRequestDataContextWithDefaults() *TelephonyRequestDataContext {
	this := TelephonyRequestDataContext{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *TelephonyRequestDataContext) GetRequest() InlineHookRequestObject {
	if o == nil || o.Request == nil {
		var ret InlineHookRequestObject
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataContext) GetRequestOk() (*InlineHookRequestObject, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *TelephonyRequestDataContext) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineHookRequestObject and assigns it to the Request field.
func (o *TelephonyRequestDataContext) SetRequest(v InlineHookRequestObject) {
	o.Request = &v
}

func (o TelephonyRequestDataContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelephonyRequestDataContext) UnmarshalJSON(bytes []byte) (err error) {
	varTelephonyRequestDataContext := _TelephonyRequestDataContext{}

	err = json.Unmarshal(bytes, &varTelephonyRequestDataContext)
	if err == nil {
		*o = TelephonyRequestDataContext(varTelephonyRequestDataContext)
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

type NullableTelephonyRequestDataContext struct {
	value *TelephonyRequestDataContext
	isSet bool
}

func (v NullableTelephonyRequestDataContext) Get() *TelephonyRequestDataContext {
	return v.value
}

func (v *NullableTelephonyRequestDataContext) Set(val *TelephonyRequestDataContext) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephonyRequestDataContext) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephonyRequestDataContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephonyRequestDataContext(val *TelephonyRequestDataContext) *NullableTelephonyRequestDataContext {
	return &NullableTelephonyRequestDataContext{value: val, isSet: true}
}

func (v NullableTelephonyRequestDataContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephonyRequestDataContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

