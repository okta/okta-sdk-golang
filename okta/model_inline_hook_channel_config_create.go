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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the InlineHookChannelConfigCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookChannelConfigCreate{}

// InlineHookChannelConfigCreate Properties of the communications channel that are used to contact your external service
type InlineHookChannelConfigCreate struct {
	// An optional list of key/value pairs for headers that you can send with the request to the external service.
	Headers []InlineHookChannelConfigHeaders `json:"headers,omitempty"`
	// The method of the Okta inline hook request
	Method *string `json:"method,omitempty"`
	// The external service endpoint that executes the inline hook handler. It must begin with `https://` and be reachable by Okta. No white space is allowed in the URI.
	Uri                  *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannelConfigCreate InlineHookChannelConfigCreate

// NewInlineHookChannelConfigCreate instantiates a new InlineHookChannelConfigCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannelConfigCreate() *InlineHookChannelConfigCreate {
	this := InlineHookChannelConfigCreate{}
	return &this
}

// NewInlineHookChannelConfigCreateWithDefaults instantiates a new InlineHookChannelConfigCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelConfigCreateWithDefaults() *InlineHookChannelConfigCreate {
	this := InlineHookChannelConfigCreate{}
	return &this
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineHookChannelConfigCreate) GetHeaders() []InlineHookChannelConfigHeaders {
	if o == nil || IsNil(o.Headers) {
		var ret []InlineHookChannelConfigHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigCreate) GetHeadersOk() ([]InlineHookChannelConfigHeaders, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineHookChannelConfigCreate) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []InlineHookChannelConfigHeaders and assigns it to the Headers field.
func (o *InlineHookChannelConfigCreate) SetHeaders(v []InlineHookChannelConfigHeaders) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineHookChannelConfigCreate) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigCreate) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookChannelConfigCreate) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineHookChannelConfigCreate) SetMethod(v string) {
	o.Method = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineHookChannelConfigCreate) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigCreate) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineHookChannelConfigCreate) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineHookChannelConfigCreate) SetUri(v string) {
	o.Uri = &v
}

func (o InlineHookChannelConfigCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookChannelConfigCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookChannelConfigCreate) UnmarshalJSON(data []byte) (err error) {
	varInlineHookChannelConfigCreate := _InlineHookChannelConfigCreate{}

	err = json.Unmarshal(data, &varInlineHookChannelConfigCreate)

	if err != nil {
		return err
	}

	*o = InlineHookChannelConfigCreate(varInlineHookChannelConfigCreate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "headers")
		delete(additionalProperties, "method")
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookChannelConfigCreate struct {
	value *InlineHookChannelConfigCreate
	isSet bool
}

func (v NullableInlineHookChannelConfigCreate) Get() *InlineHookChannelConfigCreate {
	return v.value
}

func (v *NullableInlineHookChannelConfigCreate) Set(val *InlineHookChannelConfigCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannelConfigCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannelConfigCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannelConfigCreate(val *InlineHookChannelConfigCreate) *NullableInlineHookChannelConfigCreate {
	return &NullableInlineHookChannelConfigCreate{value: val, isSet: true}
}

func (v NullableInlineHookChannelConfigCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannelConfigCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
