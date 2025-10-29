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

// checks if the RegistrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationResponse{}

// RegistrationResponse struct for RegistrationResponse
type RegistrationResponse struct {
	// The `commands` object lets you invoke commands to modify or add values to the attributes in the Okta user profile that are created for this user. The object also lets you control whether or not the registration attempt is allowed to proceed.  This object is an array, allowing you to send multiple commands in your response. Each array element requires a `type` property and a `value` property. The `type` property is where you specify which of the supported commands you wish to execute, and `value` is where you supply parameters for that command.  The registration inline hook supports these three commands: * `com.okta.user.profile.update`: Change attribute values in the user's Okta user profile. For SSR only. Invalid if used with a Progressive Profile response. * `com.okta.action.update`: Allow or deny the user's registration. * `com.okta.user.progressive.profile.update`: Change attribute values in the user's Okta Progressive Profile.
	Commands             []RegistrationResponseCommandsInner `json:"commands,omitempty"`
	Error                *RegistrationResponseError          `json:"Error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationResponse RegistrationResponse

// NewRegistrationResponse instantiates a new RegistrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationResponse() *RegistrationResponse {
	this := RegistrationResponse{}
	return &this
}

// NewRegistrationResponseWithDefaults instantiates a new RegistrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationResponseWithDefaults() *RegistrationResponse {
	this := RegistrationResponse{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *RegistrationResponse) GetCommands() []RegistrationResponseCommandsInner {
	if o == nil || IsNil(o.Commands) {
		var ret []RegistrationResponseCommandsInner
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponse) GetCommandsOk() ([]RegistrationResponseCommandsInner, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *RegistrationResponse) HasCommands() bool {
	if o != nil && !IsNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []RegistrationResponseCommandsInner and assigns it to the Commands field.
func (o *RegistrationResponse) SetCommands(v []RegistrationResponseCommandsInner) {
	o.Commands = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *RegistrationResponse) GetError() RegistrationResponseError {
	if o == nil || IsNil(o.Error) {
		var ret RegistrationResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponse) GetErrorOk() (*RegistrationResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *RegistrationResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given RegistrationResponseError and assigns it to the Error field.
func (o *RegistrationResponse) SetError(v RegistrationResponseError) {
	o.Error = &v
}

func (o RegistrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commands) {
		toSerialize["commands"] = o.Commands
	}
	if !IsNil(o.Error) {
		toSerialize["Error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationResponse) UnmarshalJSON(data []byte) (err error) {
	varRegistrationResponse := _RegistrationResponse{}

	err = json.Unmarshal(data, &varRegistrationResponse)

	if err != nil {
		return err
	}

	*o = RegistrationResponse(varRegistrationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commands")
		delete(additionalProperties, "Error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationResponse struct {
	value *RegistrationResponse
	isSet bool
}

func (v NullableRegistrationResponse) Get() *RegistrationResponse {
	return v.value
}

func (v *NullableRegistrationResponse) Set(val *RegistrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationResponse(val *RegistrationResponse) *NullableRegistrationResponse {
	return &NullableRegistrationResponse{value: val, isSet: true}
}

func (v NullableRegistrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
