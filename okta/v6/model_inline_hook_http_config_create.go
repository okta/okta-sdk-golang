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

// InlineHookHttpConfigCreate struct for InlineHookHttpConfigCreate
type InlineHookHttpConfigCreate struct {
	AuthScheme NullableInlineHookChannelConfigAuthSchemeBody `json:"authScheme,omitempty"`
	// An optional list of key/value pairs for headers that you can send with the request to the external service.
	Headers []InlineHookChannelConfigHeaders `json:"headers,omitempty"`
	// The method of the Okta inline hook request
	Method *string `json:"method,omitempty"`
	// The external service endpoint that executes the inline hook handler. It must begin with `https://` and be reachable by Okta. No white space is allowed in the URI.
	Uri *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookHttpConfigCreate InlineHookHttpConfigCreate

// NewInlineHookHttpConfigCreate instantiates a new InlineHookHttpConfigCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookHttpConfigCreate() *InlineHookHttpConfigCreate {
	this := InlineHookHttpConfigCreate{}
	return &this
}

// NewInlineHookHttpConfigCreateWithDefaults instantiates a new InlineHookHttpConfigCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookHttpConfigCreateWithDefaults() *InlineHookHttpConfigCreate {
	this := InlineHookHttpConfigCreate{}
	return &this
}

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineHookHttpConfigCreate) GetAuthScheme() InlineHookChannelConfigAuthSchemeBody {
	if o == nil || o.AuthScheme.Get() == nil {
		var ret InlineHookChannelConfigAuthSchemeBody
		return ret
	}
	return *o.AuthScheme.Get()
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineHookHttpConfigCreate) GetAuthSchemeOk() (*InlineHookChannelConfigAuthSchemeBody, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthScheme.Get(), o.AuthScheme.IsSet()
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *InlineHookHttpConfigCreate) HasAuthScheme() bool {
	if o != nil && o.AuthScheme.IsSet() {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given NullableInlineHookChannelConfigAuthSchemeBody and assigns it to the AuthScheme field.
func (o *InlineHookHttpConfigCreate) SetAuthScheme(v InlineHookChannelConfigAuthSchemeBody) {
	o.AuthScheme.Set(&v)
}
// SetAuthSchemeNil sets the value for AuthScheme to be an explicit nil
func (o *InlineHookHttpConfigCreate) SetAuthSchemeNil() {
	o.AuthScheme.Set(nil)
}

// UnsetAuthScheme ensures that no value is present for AuthScheme, not even an explicit nil
func (o *InlineHookHttpConfigCreate) UnsetAuthScheme() {
	o.AuthScheme.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineHookHttpConfigCreate) GetHeaders() []InlineHookChannelConfigHeaders {
	if o == nil || o.Headers == nil {
		var ret []InlineHookChannelConfigHeaders
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookHttpConfigCreate) GetHeadersOk() ([]InlineHookChannelConfigHeaders, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineHookHttpConfigCreate) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []InlineHookChannelConfigHeaders and assigns it to the Headers field.
func (o *InlineHookHttpConfigCreate) SetHeaders(v []InlineHookChannelConfigHeaders) {
	o.Headers = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *InlineHookHttpConfigCreate) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookHttpConfigCreate) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *InlineHookHttpConfigCreate) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *InlineHookHttpConfigCreate) SetMethod(v string) {
	o.Method = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineHookHttpConfigCreate) GetUri() string {
	if o == nil || o.Uri == nil {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookHttpConfigCreate) GetUriOk() (*string, bool) {
	if o == nil || o.Uri == nil {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineHookHttpConfigCreate) HasUri() bool {
	if o != nil && o.Uri != nil {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineHookHttpConfigCreate) SetUri(v string) {
	o.Uri = &v
}

func (o InlineHookHttpConfigCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthScheme.IsSet() {
		toSerialize["authScheme"] = o.AuthScheme.Get()
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Uri != nil {
		toSerialize["uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookHttpConfigCreate) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookHttpConfigCreate := _InlineHookHttpConfigCreate{}

	err = json.Unmarshal(bytes, &varInlineHookHttpConfigCreate)
	if err == nil {
		*o = InlineHookHttpConfigCreate(varInlineHookHttpConfigCreate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "method")
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookHttpConfigCreate struct {
	value *InlineHookHttpConfigCreate
	isSet bool
}

func (v NullableInlineHookHttpConfigCreate) Get() *InlineHookHttpConfigCreate {
	return v.value
}

func (v *NullableInlineHookHttpConfigCreate) Set(val *InlineHookHttpConfigCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookHttpConfigCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookHttpConfigCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookHttpConfigCreate(val *InlineHookHttpConfigCreate) *NullableInlineHookHttpConfigCreate {
	return &NullableInlineHookHttpConfigCreate{value: val, isSet: true}
}

func (v NullableInlineHookHttpConfigCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookHttpConfigCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

