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

// InlineHookChannelConfigAuthSchemeResponse The authentication scheme to use for this request
type InlineHookChannelConfigAuthSchemeResponse struct {
	// The header name for the authorization server
	Key *string `json:"key,omitempty"`
	// The authentication scheme type. Supported type&mdash;`HEADER`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookChannelConfigAuthSchemeResponse InlineHookChannelConfigAuthSchemeResponse

// NewInlineHookChannelConfigAuthSchemeResponse instantiates a new InlineHookChannelConfigAuthSchemeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookChannelConfigAuthSchemeResponse() *InlineHookChannelConfigAuthSchemeResponse {
	this := InlineHookChannelConfigAuthSchemeResponse{}
	return &this
}

// NewInlineHookChannelConfigAuthSchemeResponseWithDefaults instantiates a new InlineHookChannelConfigAuthSchemeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookChannelConfigAuthSchemeResponseWithDefaults() *InlineHookChannelConfigAuthSchemeResponse {
	this := InlineHookChannelConfigAuthSchemeResponse{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *InlineHookChannelConfigAuthSchemeResponse) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigAuthSchemeResponse) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InlineHookChannelConfigAuthSchemeResponse) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *InlineHookChannelConfigAuthSchemeResponse) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineHookChannelConfigAuthSchemeResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigAuthSchemeResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineHookChannelConfigAuthSchemeResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineHookChannelConfigAuthSchemeResponse) SetType(v string) {
	o.Type = &v
}

func (o InlineHookChannelConfigAuthSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineHookChannelConfigAuthSchemeResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInlineHookChannelConfigAuthSchemeResponse := _InlineHookChannelConfigAuthSchemeResponse{}

	err = json.Unmarshal(bytes, &varInlineHookChannelConfigAuthSchemeResponse)
	if err == nil {
		*o = InlineHookChannelConfigAuthSchemeResponse(varInlineHookChannelConfigAuthSchemeResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableInlineHookChannelConfigAuthSchemeResponse struct {
	value *InlineHookChannelConfigAuthSchemeResponse
	isSet bool
}

func (v NullableInlineHookChannelConfigAuthSchemeResponse) Get() *InlineHookChannelConfigAuthSchemeResponse {
	return v.value
}

func (v *NullableInlineHookChannelConfigAuthSchemeResponse) Set(val *InlineHookChannelConfigAuthSchemeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookChannelConfigAuthSchemeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookChannelConfigAuthSchemeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookChannelConfigAuthSchemeResponse(val *InlineHookChannelConfigAuthSchemeResponse) *NullableInlineHookChannelConfigAuthSchemeResponse {
	return &NullableInlineHookChannelConfigAuthSchemeResponse{value: val, isSet: true}
}

func (v NullableInlineHookChannelConfigAuthSchemeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookChannelConfigAuthSchemeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

