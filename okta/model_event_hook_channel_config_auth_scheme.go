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

// EventHookChannelConfigAuthScheme The authentication scheme used for this request.  To use Basic Auth for authentication, set `type` to `HEADER`, `key` to `Authorization`, and `value` to the Base64-encoded string of \"username:password\". Ensure that you include the scheme (including space) as part of the `value` parameter. For example, `Basic YWRtaW46c3VwZXJzZWNyZXQ=`. See [HTTP Basic Authentication](/books/api-security/authn/api-authentication-options/#http-basic-authentication).
type EventHookChannelConfigAuthScheme struct {
	// The name for the authorization header
	Key *string `json:"key,omitempty"`
	// The authentication scheme type. Currently only supports `HEADER`.
	Type *string `json:"type,omitempty"`
	// The header value. This secret key is passed to your external service endpoint for security verification. This property is not returned in the response.
	Value *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventHookChannelConfigAuthScheme EventHookChannelConfigAuthScheme

// NewEventHookChannelConfigAuthScheme instantiates a new EventHookChannelConfigAuthScheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventHookChannelConfigAuthScheme() *EventHookChannelConfigAuthScheme {
	this := EventHookChannelConfigAuthScheme{}
	return &this
}

// NewEventHookChannelConfigAuthSchemeWithDefaults instantiates a new EventHookChannelConfigAuthScheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventHookChannelConfigAuthSchemeWithDefaults() *EventHookChannelConfigAuthScheme {
	this := EventHookChannelConfigAuthScheme{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EventHookChannelConfigAuthScheme) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookChannelConfigAuthScheme) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EventHookChannelConfigAuthScheme) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EventHookChannelConfigAuthScheme) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EventHookChannelConfigAuthScheme) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookChannelConfigAuthScheme) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EventHookChannelConfigAuthScheme) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EventHookChannelConfigAuthScheme) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EventHookChannelConfigAuthScheme) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventHookChannelConfigAuthScheme) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EventHookChannelConfigAuthScheme) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EventHookChannelConfigAuthScheme) SetValue(v string) {
	o.Value = &v
}

func (o EventHookChannelConfigAuthScheme) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EventHookChannelConfigAuthScheme) UnmarshalJSON(bytes []byte) (err error) {
	varEventHookChannelConfigAuthScheme := _EventHookChannelConfigAuthScheme{}

	err = json.Unmarshal(bytes, &varEventHookChannelConfigAuthScheme)
	if err == nil {
		*o = EventHookChannelConfigAuthScheme(varEventHookChannelConfigAuthScheme)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEventHookChannelConfigAuthScheme struct {
	value *EventHookChannelConfigAuthScheme
	isSet bool
}

func (v NullableEventHookChannelConfigAuthScheme) Get() *EventHookChannelConfigAuthScheme {
	return v.value
}

func (v *NullableEventHookChannelConfigAuthScheme) Set(val *EventHookChannelConfigAuthScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableEventHookChannelConfigAuthScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableEventHookChannelConfigAuthScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventHookChannelConfigAuthScheme(val *EventHookChannelConfigAuthScheme) *NullableEventHookChannelConfigAuthScheme {
	return &NullableEventHookChannelConfigAuthScheme{value: val, isSet: true}
}

func (v NullableEventHookChannelConfigAuthScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventHookChannelConfigAuthScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

