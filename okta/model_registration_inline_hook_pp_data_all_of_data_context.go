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

// checks if the RegistrationInlineHookPPDataAllOfDataContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInlineHookPPDataAllOfDataContext{}

// RegistrationInlineHookPPDataAllOfDataContext struct for RegistrationInlineHookPPDataAllOfDataContext
type RegistrationInlineHookPPDataAllOfDataContext struct {
	Request              *InlineHookRequestObject                          `json:"request,omitempty"`
	User                 *RegistrationInlineHookPPDataAllOfDataContextUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookPPDataAllOfDataContext RegistrationInlineHookPPDataAllOfDataContext

// NewRegistrationInlineHookPPDataAllOfDataContext instantiates a new RegistrationInlineHookPPDataAllOfDataContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookPPDataAllOfDataContext() *RegistrationInlineHookPPDataAllOfDataContext {
	this := RegistrationInlineHookPPDataAllOfDataContext{}
	return &this
}

// NewRegistrationInlineHookPPDataAllOfDataContextWithDefaults instantiates a new RegistrationInlineHookPPDataAllOfDataContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookPPDataAllOfDataContextWithDefaults() *RegistrationInlineHookPPDataAllOfDataContext {
	this := RegistrationInlineHookPPDataAllOfDataContext{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPDataAllOfDataContext) GetRequest() InlineHookRequestObject {
	if o == nil || IsNil(o.Request) {
		var ret InlineHookRequestObject
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContext) GetRequestOk() (*InlineHookRequestObject, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContext) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given InlineHookRequestObject and assigns it to the Request field.
func (o *RegistrationInlineHookPPDataAllOfDataContext) SetRequest(v InlineHookRequestObject) {
	o.Request = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPDataAllOfDataContext) GetUser() RegistrationInlineHookPPDataAllOfDataContextUser {
	if o == nil || IsNil(o.User) {
		var ret RegistrationInlineHookPPDataAllOfDataContextUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContext) GetUserOk() (*RegistrationInlineHookPPDataAllOfDataContextUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPDataAllOfDataContext) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given RegistrationInlineHookPPDataAllOfDataContextUser and assigns it to the User field.
func (o *RegistrationInlineHookPPDataAllOfDataContext) SetUser(v RegistrationInlineHookPPDataAllOfDataContextUser) {
	o.User = &v
}

func (o RegistrationInlineHookPPDataAllOfDataContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInlineHookPPDataAllOfDataContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationInlineHookPPDataAllOfDataContext) UnmarshalJSON(data []byte) (err error) {
	varRegistrationInlineHookPPDataAllOfDataContext := _RegistrationInlineHookPPDataAllOfDataContext{}

	err = json.Unmarshal(data, &varRegistrationInlineHookPPDataAllOfDataContext)

	if err != nil {
		return err
	}

	*o = RegistrationInlineHookPPDataAllOfDataContext(varRegistrationInlineHookPPDataAllOfDataContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationInlineHookPPDataAllOfDataContext struct {
	value *RegistrationInlineHookPPDataAllOfDataContext
	isSet bool
}

func (v NullableRegistrationInlineHookPPDataAllOfDataContext) Get() *RegistrationInlineHookPPDataAllOfDataContext {
	return v.value
}

func (v *NullableRegistrationInlineHookPPDataAllOfDataContext) Set(val *RegistrationInlineHookPPDataAllOfDataContext) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookPPDataAllOfDataContext) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookPPDataAllOfDataContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookPPDataAllOfDataContext(val *RegistrationInlineHookPPDataAllOfDataContext) *NullableRegistrationInlineHookPPDataAllOfDataContext {
	return &NullableRegistrationInlineHookPPDataAllOfDataContext{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookPPDataAllOfDataContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookPPDataAllOfDataContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
