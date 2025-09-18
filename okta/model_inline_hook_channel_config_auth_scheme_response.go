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

// checks if the InlineHookChannelConfigAuthSchemeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookChannelConfigAuthSchemeResponse{}

// InlineHookChannelConfigAuthSchemeResponse The authentication scheme to use for this request
type InlineHookChannelConfigAuthSchemeResponse struct {
	// The header name for the authorization server
	Key *string `json:"key,omitempty"`
	// The authentication scheme type. Supported type&mdash;`HEADER`
	Type                 *string `json:"type,omitempty"`
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
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigAuthSchemeResponse) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *InlineHookChannelConfigAuthSchemeResponse) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookChannelConfigAuthSchemeResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineHookChannelConfigAuthSchemeResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineHookChannelConfigAuthSchemeResponse) SetType(v string) {
	o.Type = &v
}

func (o InlineHookChannelConfigAuthSchemeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookChannelConfigAuthSchemeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookChannelConfigAuthSchemeResponse) UnmarshalJSON(data []byte) (err error) {
	varInlineHookChannelConfigAuthSchemeResponse := _InlineHookChannelConfigAuthSchemeResponse{}

	err = json.Unmarshal(data, &varInlineHookChannelConfigAuthSchemeResponse)

	if err != nil {
		return err
	}

	*o = InlineHookChannelConfigAuthSchemeResponse(varInlineHookChannelConfigAuthSchemeResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
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
