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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the AuthenticationMethodObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationMethodObject{}

// AuthenticationMethodObject struct for AuthenticationMethodObject
type AuthenticationMethodObject struct {
	// <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Authenticator ID
	Id *string `json:"id,omitempty"`
	// A label that identifies the authenticator
	Key string `json:"key"`
	// Specifies the method used for the authenticator
	Method               *string `json:"method,omitempty"`
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationMethodObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodObject) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *AuthenticationMethodObject) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *AuthenticationMethodObject) SetMethod(v string) {
	o.Method = &v
}

func (o AuthenticationMethodObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationMethodObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticationMethodObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAuthenticationMethodObject := _AuthenticationMethodObject{}

	err = json.Unmarshal(data, &varAuthenticationMethodObject)

	if err != nil {
		return err
	}

	*o = AuthenticationMethodObject(varAuthenticationMethodObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "method")
		o.AdditionalProperties = additionalProperties
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
