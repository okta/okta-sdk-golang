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

// checks if the UserImportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportResponse{}

// UserImportResponse struct for UserImportResponse
type UserImportResponse struct {
	// The `commands` object is where you can provide commands to Okta. It is an array that allows you to send multiple commands. Each array element needs to consist of a type-value pair.
	Commands             []UserImportResponseCommandsInner `json:"commands,omitempty"`
	Error                *UserImportResponseError          `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserImportResponse UserImportResponse

// NewUserImportResponse instantiates a new UserImportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportResponse() *UserImportResponse {
	this := UserImportResponse{}
	return &this
}

// NewUserImportResponseWithDefaults instantiates a new UserImportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportResponseWithDefaults() *UserImportResponse {
	this := UserImportResponse{}
	return &this
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *UserImportResponse) GetCommands() []UserImportResponseCommandsInner {
	if o == nil || IsNil(o.Commands) {
		var ret []UserImportResponseCommandsInner
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportResponse) GetCommandsOk() ([]UserImportResponseCommandsInner, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *UserImportResponse) HasCommands() bool {
	if o != nil && !IsNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []UserImportResponseCommandsInner and assigns it to the Commands field.
func (o *UserImportResponse) SetCommands(v []UserImportResponseCommandsInner) {
	o.Commands = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UserImportResponse) GetError() UserImportResponseError {
	if o == nil || IsNil(o.Error) {
		var ret UserImportResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportResponse) GetErrorOk() (*UserImportResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UserImportResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given UserImportResponseError and assigns it to the Error field.
func (o *UserImportResponse) SetError(v UserImportResponseError) {
	o.Error = &v
}

func (o UserImportResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportResponse) ToMap() (map[string]interface{}, error) {
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

func (o *UserImportResponse) UnmarshalJSON(data []byte) (err error) {
	varUserImportResponse := _UserImportResponse{}

	err = json.Unmarshal(data, &varUserImportResponse)

	if err != nil {
		return err
	}

	*o = UserImportResponse(varUserImportResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "commands")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserImportResponse struct {
	value *UserImportResponse
	isSet bool
}

func (v NullableUserImportResponse) Get() *UserImportResponse {
	return v.value
}

func (v *NullableUserImportResponse) Set(val *UserImportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportResponse(val *UserImportResponse) *NullableUserImportResponse {
	return &NullableUserImportResponse{value: val, isSet: true}
}

func (v NullableUserImportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
