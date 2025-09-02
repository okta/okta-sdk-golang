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
	"fmt"
)

// AuthenticationMethodObject struct for AuthenticationMethodObject
type AuthenticationMethodObject struct {
	// <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Authenticator ID
	Id *string `json:"id,omitempty"`
	// A label that identifies the authenticator
	Key string `json:"key"`
	// Specifies the method used for the authenticator
	Method *string `json:"method,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticationMethodObject AuthenticationMethodObject

// NewAuthenticationMethodObject instantiates a new AuthenticationMethodObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationMethodObject(key string) *AuthenticationMethodObject {
	this := AuthenticationMethodObject{}
	this.Key = key
	return &this
}

// NewAuthenticationMethodObjectWithDefaults instantiates a new AuthenticationMethodObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationMethodObjectWithDefaults() *AuthenticationMethodObject {
	this := AuthenticationMethodObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationMethodObject) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodObject) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationMethodObject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticationMethodObject) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value
func (o *AuthenticationMethodObject) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodObject) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AuthenticationMethodObject) SetKey(v string) {
	o.Key = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *AuthenticationMethodObject) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodObject) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *AuthenticationMethodObject) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *AuthenticationMethodObject) SetMethod(v string) {
	o.Method = &v
}

func (o AuthenticationMethodObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticationMethodObject) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticationMethodObject := _AuthenticationMethodObject{}

	err = json.Unmarshal(bytes, &varAuthenticationMethodObject)
	if err == nil {
		*o = AuthenticationMethodObject(varAuthenticationMethodObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "method")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticationMethodObject struct {
	value *AuthenticationMethodObject
	isSet bool
}

func (v NullableAuthenticationMethodObject) Get() *AuthenticationMethodObject {
	return v.value
}

func (v *NullableAuthenticationMethodObject) Set(val *AuthenticationMethodObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationMethodObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationMethodObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationMethodObject(val *AuthenticationMethodObject) *NullableAuthenticationMethodObject {
	return &NullableAuthenticationMethodObject{value: val, isSet: true}
}

func (v NullableAuthenticationMethodObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationMethodObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

