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

// checks if the UserImportRequestDataAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportRequestDataAction{}

// UserImportRequestDataAction The object that specifies the default action Okta is set to take
type UserImportRequestDataAction struct {
	// The current default action that results when Okta imports a user. The two possible values are `CREATE_USER` and `LINK_USER`. You can change the action that is taken by means of the commands object you return.
	Result               *string `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserImportRequestDataAction UserImportRequestDataAction

// NewUserImportRequestDataAction instantiates a new UserImportRequestDataAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportRequestDataAction() *UserImportRequestDataAction {
	this := UserImportRequestDataAction{}
	return &this
}

// NewUserImportRequestDataActionWithDefaults instantiates a new UserImportRequestDataAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportRequestDataActionWithDefaults() *UserImportRequestDataAction {
	this := UserImportRequestDataAction{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UserImportRequestDataAction) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataAction) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UserImportRequestDataAction) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *UserImportRequestDataAction) SetResult(v string) {
	o.Result = &v
}

func (o UserImportRequestDataAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportRequestDataAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserImportRequestDataAction) UnmarshalJSON(data []byte) (err error) {
	varUserImportRequestDataAction := _UserImportRequestDataAction{}

	err = json.Unmarshal(data, &varUserImportRequestDataAction)

	if err != nil {
		return err
	}

	*o = UserImportRequestDataAction(varUserImportRequestDataAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserImportRequestDataAction struct {
	value *UserImportRequestDataAction
	isSet bool
}

func (v NullableUserImportRequestDataAction) Get() *UserImportRequestDataAction {
	return v.value
}

func (v *NullableUserImportRequestDataAction) Set(val *UserImportRequestDataAction) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportRequestDataAction) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportRequestDataAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportRequestDataAction(val *UserImportRequestDataAction) *NullableUserImportRequestDataAction {
	return &NullableUserImportRequestDataAction{value: val, isSet: true}
}

func (v NullableUserImportRequestDataAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportRequestDataAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
