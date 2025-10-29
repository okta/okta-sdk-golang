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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the TokenHookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenHookResponse{}

// TokenHookResponse For the token inline hook, the `commands` and `error` objects that you can return in the JSON payload of your response are defined in the following sections. > **Note:** The size of your response payload must be less than 256 KB.
type TokenHookResponse struct {
	// You can use the `commands` object to provide commands to Okta. It's where you can tell Okta to add more claims to the token. The `commands` object is an array, allowing you to send multiple commands. In each array element, there needs to be a `type` property and `value` property. The `type` property is where you specify which of the supported commands you want to execute, and `value` is where you supply an operand for that command. In the case of the token hook type, the `value` property is itself a nested object in which you specify a particular operation, a path to act on, and a value.
	Commands             []TokenHookResponseCommandsInner `json:"commands,omitempty"`
	Error                *TokenHookResponseError          `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenHookResponse TokenHookResponse

// NewTokenHookResponse instantiates a new TokenHookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenHookResponse() *TokenHookResponse {
	this := TokenHookResponse{}
	return &this
}

// NewTokenHookResponseWithDefaults instantiates a new TokenHookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenHookResponseWithDefaults() *TokenHookResponse {
	this := TokenHookResponse{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *TokenHookResponse) GetCommands() []TokenHookResponseCommandsInner {
	if o == nil || IsNil(o.Commands) {
		var ret []TokenHookResponseCommandsInner
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenHookResponse) GetCommandsOk() ([]TokenHookResponseCommandsInner, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *TokenHookResponse) HasCommands() bool {
	if o != nil && !IsNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []TokenHookResponseCommandsInner and assigns it to the Commands field.
func (o *TokenHookResponse) SetCommands(v []TokenHookResponseCommandsInner) {
	o.Commands = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *TokenHookResponse) GetError() TokenHookResponseError {
	if o == nil || IsNil(o.Error) {
		var ret TokenHookResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenHookResponse) GetErrorOk() (*TokenHookResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TokenHookResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given TokenHookResponseError and assigns it to the Error field.
func (o *TokenHookResponse) SetError(v TokenHookResponseError) {
	o.Error = &v
}

func (o TokenHookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenHookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commands) {
		toSerialize["commands"] = o.Commands
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenHookResponse) UnmarshalJSON(data []byte) (err error) {
	varTokenHookResponse := _TokenHookResponse{}

	err = json.Unmarshal(data, &varTokenHookResponse)

	if err != nil {
		return err
	}

	*o = TokenHookResponse(varTokenHookResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commands")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenHookResponse struct {
	value *TokenHookResponse
	isSet bool
}

func (v NullableTokenHookResponse) Get() *TokenHookResponse {
	return v.value
}

func (v *NullableTokenHookResponse) Set(val *TokenHookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenHookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenHookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenHookResponse(val *TokenHookResponse) *NullableTokenHookResponse {
	return &NullableTokenHookResponse{value: val, isSet: true}
}

func (v NullableTokenHookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenHookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
