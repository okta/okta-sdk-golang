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

// checks if the UserImportResponseCommandsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportResponseCommandsInner{}

// UserImportResponseCommandsInner struct for UserImportResponseCommandsInner
type UserImportResponseCommandsInner struct {
	// The command types supported for the import inline hook. When using the `com.okta.action.update` command to specify that the user should be treated as a match, you need to also provide a `com.okta.user.update` command that sets the ID of the Okta user.
	Type *string `json:"type,omitempty"`
	// The `value` object is the parameter to pass to the command. In the case of the `com.okta.appUser.profile.update` and `com.okta.user.profile.update` commands, the parameter should be a list of one or more profile attributes and the values you wish to set them to. In the case of the `com.okta.action.update` command, the parameter should be a `result` property set to either `CREATE_USER` or `LINK_USER`.
	Value                *map[string]string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserImportResponseCommandsInner UserImportResponseCommandsInner

// NewUserImportResponseCommandsInner instantiates a new UserImportResponseCommandsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportResponseCommandsInner() *UserImportResponseCommandsInner {
	this := UserImportResponseCommandsInner{}
	return &this
}

// NewUserImportResponseCommandsInnerWithDefaults instantiates a new UserImportResponseCommandsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportResponseCommandsInnerWithDefaults() *UserImportResponseCommandsInner {
	this := UserImportResponseCommandsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserImportResponseCommandsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportResponseCommandsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserImportResponseCommandsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserImportResponseCommandsInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UserImportResponseCommandsInner) GetValue() map[string]string {
	if o == nil || IsNil(o.Value) {
		var ret map[string]string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportResponseCommandsInner) GetValueOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UserImportResponseCommandsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]string and assigns it to the Value field.
func (o *UserImportResponseCommandsInner) SetValue(v map[string]string) {
	o.Value = &v
}

func (o UserImportResponseCommandsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportResponseCommandsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserImportResponseCommandsInner) UnmarshalJSON(data []byte) (err error) {
	varUserImportResponseCommandsInner := _UserImportResponseCommandsInner{}

	err = json.Unmarshal(data, &varUserImportResponseCommandsInner)

	if err != nil {
		return err
	}

	*o = UserImportResponseCommandsInner(varUserImportResponseCommandsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserImportResponseCommandsInner struct {
	value *UserImportResponseCommandsInner
	isSet bool
}

func (v NullableUserImportResponseCommandsInner) Get() *UserImportResponseCommandsInner {
	return v.value
}

func (v *NullableUserImportResponseCommandsInner) Set(val *UserImportResponseCommandsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportResponseCommandsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportResponseCommandsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportResponseCommandsInner(val *UserImportResponseCommandsInner) *NullableUserImportResponseCommandsInner {
	return &NullableUserImportResponseCommandsInner{value: val, isSet: true}
}

func (v NullableUserImportResponseCommandsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportResponseCommandsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
