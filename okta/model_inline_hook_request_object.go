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

// InlineHookRequestObject struct for InlineHookRequestObject
type InlineHookRequestObject struct {
	Request *InlineHookRequestObjectRequest `json:"request,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookRequestObject InlineHookRequestObject

// NewInlineHookRequestObject instantiates a new InlineHookRequestObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookRequestObject() *InlineHookRequestObject {
	this := InlineHookRequestObject{}
	return &this
}

// NewInlineHookRequestObjectWithDefaults instantiates a new InlineHookRequestObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookRequestObjectWithDefaults() *InlineHookRequestObject {
	this := InlineHookRequestObject{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *InlineHookRequestObject) GetRequest() InlineHookRequestObjectRequest {
	if o == nil || o.Request == nil {
		var ret InlineHookRequestObjectRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookRequestObject) GetRequestOk() (*InlineHookRequestObjectRequest, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *InlineHookRequestObject) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineHookRequestObjectRequest and assigns it to the Request field.
func (o *InlineHookRequestObject) SetRequest(v InlineHookRequestObjectRequest) {
	o.Request = &v
}

func (o InlineHookRequestObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookRequestObject) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookRequestObject := _InlineHookRequestObject{}

	err = json.Unmarshal(bytes, &varInlineHookRequestObject)
	if err == nil {
		*o = InlineHookRequestObject(varInlineHookRequestObject)
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

type NullableInlineHookRequestObject struct {
	value *InlineHookRequestObject
	isSet bool
}

func (v NullableInlineHookRequestObject) Get() *InlineHookRequestObject {
	return v.value
}

func (v *NullableInlineHookRequestObject) Set(val *InlineHookRequestObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookRequestObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookRequestObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookRequestObject(val *InlineHookRequestObject) *NullableInlineHookRequestObject {
	return &NullableInlineHookRequestObject{value: val, isSet: true}
}

func (v NullableInlineHookRequestObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookRequestObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

