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

// RegistrationInlineHookCommand struct for RegistrationInlineHookCommand
type RegistrationInlineHookCommand struct {
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookCommand RegistrationInlineHookCommand

// NewRegistrationInlineHookCommand instantiates a new RegistrationInlineHookCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookCommand() *RegistrationInlineHookCommand {
	this := RegistrationInlineHookCommand{}
	return &this
}

// NewRegistrationInlineHookCommandWithDefaults instantiates a new RegistrationInlineHookCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookCommandWithDefaults() *RegistrationInlineHookCommand {
	this := RegistrationInlineHookCommand{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegistrationInlineHookCommand) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookCommand) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegistrationInlineHookCommand) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RegistrationInlineHookCommand) SetType(v string) {
	o.Type = &v
}

func (o RegistrationInlineHookCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RegistrationInlineHookCommand) UnmarshalJSON(bytes []byte) (err error) {
	varRegistrationInlineHookCommand := _RegistrationInlineHookCommand{}

	err = json.Unmarshal(bytes, &varRegistrationInlineHookCommand)
	if err == nil {
		*o = RegistrationInlineHookCommand(varRegistrationInlineHookCommand)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRegistrationInlineHookCommand struct {
	value *RegistrationInlineHookCommand
	isSet bool
}

func (v NullableRegistrationInlineHookCommand) Get() *RegistrationInlineHookCommand {
	return v.value
}

func (v *NullableRegistrationInlineHookCommand) Set(val *RegistrationInlineHookCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookCommand(val *RegistrationInlineHookCommand) *NullableRegistrationInlineHookCommand {
	return &NullableRegistrationInlineHookCommand{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

