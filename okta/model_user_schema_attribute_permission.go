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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// UserSchemaAttributePermission struct for UserSchemaAttributePermission
type UserSchemaAttributePermission struct {
	Action *string `json:"action,omitempty"`
	Principal *string `json:"principal,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaAttributePermission UserSchemaAttributePermission

// NewUserSchemaAttributePermission instantiates a new UserSchemaAttributePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaAttributePermission() *UserSchemaAttributePermission {
	this := UserSchemaAttributePermission{}
	return &this
}

// NewUserSchemaAttributePermissionWithDefaults instantiates a new UserSchemaAttributePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaAttributePermissionWithDefaults() *UserSchemaAttributePermission {
	this := UserSchemaAttributePermission{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UserSchemaAttributePermission) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttributePermission) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UserSchemaAttributePermission) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *UserSchemaAttributePermission) SetAction(v string) {
	o.Action = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *UserSchemaAttributePermission) GetPrincipal() string {
	if o == nil || o.Principal == nil {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaAttributePermission) GetPrincipalOk() (*string, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *UserSchemaAttributePermission) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *UserSchemaAttributePermission) SetPrincipal(v string) {
	o.Principal = &v
}

func (o UserSchemaAttributePermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Principal != nil {
		toSerialize["principal"] = o.Principal
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSchemaAttributePermission) UnmarshalJSON(bytes []byte) (err error) {
	varUserSchemaAttributePermission := _UserSchemaAttributePermission{}

	err = json.Unmarshal(bytes, &varUserSchemaAttributePermission)
	if err == nil {
		*o = UserSchemaAttributePermission(varUserSchemaAttributePermission)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "principal")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserSchemaAttributePermission struct {
	value *UserSchemaAttributePermission
	isSet bool
}

func (v NullableUserSchemaAttributePermission) Get() *UserSchemaAttributePermission {
	return v.value
}

func (v *NullableUserSchemaAttributePermission) Set(val *UserSchemaAttributePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaAttributePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaAttributePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaAttributePermission(val *UserSchemaAttributePermission) *NullableUserSchemaAttributePermission {
	return &NullableUserSchemaAttributePermission{value: val, isSet: true}
}

func (v NullableUserSchemaAttributePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaAttributePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

