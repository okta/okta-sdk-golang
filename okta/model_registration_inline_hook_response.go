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

// checks if the RegistrationInlineHookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInlineHookResponse{}

// RegistrationInlineHookResponse Registration inline hook response
type RegistrationInlineHookResponse struct {
	Commands             []string `json:"commands,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookResponse RegistrationInlineHookResponse

// NewRegistrationInlineHookResponse instantiates a new RegistrationInlineHookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookResponse() *RegistrationInlineHookResponse {
	this := RegistrationInlineHookResponse{}
	return &this
}

// NewRegistrationInlineHookResponseWithDefaults instantiates a new RegistrationInlineHookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookResponseWithDefaults() *RegistrationInlineHookResponse {
	this := RegistrationInlineHookResponse{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *RegistrationInlineHookResponse) GetCommands() []string {
	if o == nil || IsNil(o.Commands) {
		var ret []string
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookResponse) GetCommandsOk() ([]string, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *RegistrationInlineHookResponse) HasCommands() bool {
	if o != nil && !IsNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []string and assigns it to the Commands field.
func (o *RegistrationInlineHookResponse) SetCommands(v []string) {
	o.Commands = v
}

func (o RegistrationInlineHookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInlineHookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commands) {
		toSerialize["commands"] = o.Commands
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationInlineHookResponse) UnmarshalJSON(data []byte) (err error) {
	varRegistrationInlineHookResponse := _RegistrationInlineHookResponse{}

	err = json.Unmarshal(data, &varRegistrationInlineHookResponse)

	if err != nil {
		return err
	}

	*o = RegistrationInlineHookResponse(varRegistrationInlineHookResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commands")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationInlineHookResponse struct {
	value *RegistrationInlineHookResponse
	isSet bool
}

func (v NullableRegistrationInlineHookResponse) Get() *RegistrationInlineHookResponse {
	return v.value
}

func (v *NullableRegistrationInlineHookResponse) Set(val *RegistrationInlineHookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookResponse(val *RegistrationInlineHookResponse) *NullableRegistrationInlineHookResponse {
	return &NullableRegistrationInlineHookResponse{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
