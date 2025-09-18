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

// checks if the SAMLHookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLHookResponse{}

// SAMLHookResponse struct for SAMLHookResponse
type SAMLHookResponse struct {
	// The `commands` object is where you tell Okta to add additional claims to the assertion or to modify the existing assertion statements.  `commands` is an array, allowing you to send multiple commands. In each array element, include a `type` property and a `value` property. The `type` property is where you specify which of the supported commands you want to execute, and `value` is where you supply an operand for that command. In the case of the SAML assertion inline hook, the `value` property is itself a nested object, in which you specify a particular operation, a path to act on, and a value.
	Commands             []SAMLHookResponseCommandsInner `json:"commands,omitempty"`
	Error                *SAMLHookResponseError          `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLHookResponse SAMLHookResponse

// NewSAMLHookResponse instantiates a new SAMLHookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLHookResponse() *SAMLHookResponse {
	this := SAMLHookResponse{}
	return &this
}

// NewSAMLHookResponseWithDefaults instantiates a new SAMLHookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLHookResponseWithDefaults() *SAMLHookResponse {
	this := SAMLHookResponse{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *SAMLHookResponse) GetCommands() []SAMLHookResponseCommandsInner {
	if o == nil || IsNil(o.Commands) {
		var ret []SAMLHookResponseCommandsInner
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLHookResponse) GetCommandsOk() ([]SAMLHookResponseCommandsInner, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *SAMLHookResponse) HasCommands() bool {
	if o != nil && !IsNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []SAMLHookResponseCommandsInner and assigns it to the Commands field.
func (o *SAMLHookResponse) SetCommands(v []SAMLHookResponseCommandsInner) {
	o.Commands = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SAMLHookResponse) GetError() SAMLHookResponseError {
	if o == nil || IsNil(o.Error) {
		var ret SAMLHookResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLHookResponse) GetErrorOk() (*SAMLHookResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SAMLHookResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given SAMLHookResponseError and assigns it to the Error field.
func (o *SAMLHookResponse) SetError(v SAMLHookResponseError) {
	o.Error = &v
}

func (o SAMLHookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLHookResponse) ToMap() (map[string]interface{}, error) {
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

func (o *SAMLHookResponse) UnmarshalJSON(data []byte) (err error) {
	varSAMLHookResponse := _SAMLHookResponse{}

	err = json.Unmarshal(data, &varSAMLHookResponse)

	if err != nil {
		return err
	}

	*o = SAMLHookResponse(varSAMLHookResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commands")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLHookResponse struct {
	value *SAMLHookResponse
	isSet bool
}

func (v NullableSAMLHookResponse) Get() *SAMLHookResponse {
	return v.value
}

func (v *NullableSAMLHookResponse) Set(val *SAMLHookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLHookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLHookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLHookResponse(val *SAMLHookResponse) *NullableSAMLHookResponse {
	return &NullableSAMLHookResponse{value: val, isSet: true}
}

func (v NullableSAMLHookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLHookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
