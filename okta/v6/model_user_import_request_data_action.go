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

// UserImportRequestDataAction The object that specifies the default action Okta is set to take
type UserImportRequestDataAction struct {
	// The current default action that results when Okta imports a user. The two possible values are `CREATE_USER` and `LINK_USER`. You can change the action that is taken by means of the commands object you return.
	Result *string `json:"result,omitempty"`
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
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataAction) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UserImportRequestDataAction) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *UserImportRequestDataAction) SetResult(v string) {
	o.Result = &v
}

func (o UserImportRequestDataAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserImportRequestDataAction) UnmarshalJSON(bytes []byte) (err error) {
	varUserImportRequestDataAction := _UserImportRequestDataAction{}

	err = json.Unmarshal(bytes, &varUserImportRequestDataAction)
	if err == nil {
		*o = UserImportRequestDataAction(varUserImportRequestDataAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

