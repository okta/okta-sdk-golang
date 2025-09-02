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

// TokenHookResponseCommandsInner struct for TokenHookResponseCommandsInner
type TokenHookResponseCommandsInner struct {
	// One of the supported commands:   `com.okta.identity.patch`: Modify an ID token   `com.okta.access.patch`: Modify an access token > **Note:** The `commands` array should only contain commands that can be applied to the requested tokens. For example, if only an ID token is requested, the `commands` array shouldn't contain commands of the type `com.okta.access.patch`.
	Type *string `json:"type,omitempty"`
	// The `value` object is where you specify the operation to perform. It's an array, which allows you to request more than one operation.
	Value []TokenHookResponseCommandsInnerValueInner `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenHookResponseCommandsInner TokenHookResponseCommandsInner

// NewTokenHookResponseCommandsInner instantiates a new TokenHookResponseCommandsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenHookResponseCommandsInner() *TokenHookResponseCommandsInner {
	this := TokenHookResponseCommandsInner{}
	return &this
}

// NewTokenHookResponseCommandsInnerWithDefaults instantiates a new TokenHookResponseCommandsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenHookResponseCommandsInnerWithDefaults() *TokenHookResponseCommandsInner {
	this := TokenHookResponseCommandsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TokenHookResponseCommandsInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenHookResponseCommandsInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TokenHookResponseCommandsInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TokenHookResponseCommandsInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TokenHookResponseCommandsInner) GetValue() []TokenHookResponseCommandsInnerValueInner {
	if o == nil || o.Value == nil {
		var ret []TokenHookResponseCommandsInnerValueInner
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenHookResponseCommandsInner) GetValueOk() ([]TokenHookResponseCommandsInnerValueInner, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TokenHookResponseCommandsInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []TokenHookResponseCommandsInnerValueInner and assigns it to the Value field.
func (o *TokenHookResponseCommandsInner) SetValue(v []TokenHookResponseCommandsInnerValueInner) {
	o.Value = v
}

func (o TokenHookResponseCommandsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *TokenHookResponseCommandsInner) UnmarshalJSON(bytes []byte) (err error) {
	varTokenHookResponseCommandsInner := _TokenHookResponseCommandsInner{}

	err = json.Unmarshal(bytes, &varTokenHookResponseCommandsInner)
	if err == nil {
		*o = TokenHookResponseCommandsInner(varTokenHookResponseCommandsInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenHookResponseCommandsInner struct {
	value *TokenHookResponseCommandsInner
	isSet bool
}

func (v NullableTokenHookResponseCommandsInner) Get() *TokenHookResponseCommandsInner {
	return v.value
}

func (v *NullableTokenHookResponseCommandsInner) Set(val *TokenHookResponseCommandsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenHookResponseCommandsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenHookResponseCommandsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenHookResponseCommandsInner(val *TokenHookResponseCommandsInner) *NullableTokenHookResponseCommandsInner {
	return &NullableTokenHookResponseCommandsInner{value: val, isSet: true}
}

func (v NullableTokenHookResponseCommandsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenHookResponseCommandsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

