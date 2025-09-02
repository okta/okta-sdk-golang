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

// AuthenticationMethodChain struct for AuthenticationMethodChain
type AuthenticationMethodChain struct {
	AuthenticationMethods []AuthenticationMethod `json:"authenticationMethods,omitempty"`
	// The next steps of the authentication method chain. This is an array of `AuthenticationMethodChain`. Only supports one item in the array.
	Next []map[string]interface{} `json:"next,omitempty"`
	// Specifies how often the user is prompted for authentication using duration format for the time period. For example, `PT2H30M` for two and a half hours. This parameter can't be set at the same time as the `reauthenticateIn` property on the `verificationMethod`.
	ReauthenticateIn *string `json:"reauthenticateIn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticationMethodChain AuthenticationMethodChain

// NewAuthenticationMethodChain instantiates a new AuthenticationMethodChain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationMethodChain() *AuthenticationMethodChain {
	this := AuthenticationMethodChain{}
	return &this
}

// NewAuthenticationMethodChainWithDefaults instantiates a new AuthenticationMethodChain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationMethodChainWithDefaults() *AuthenticationMethodChain {
	this := AuthenticationMethodChain{}
	return &this
}

// GetAuthenticationMethods returns the AuthenticationMethods field value if set, zero value otherwise.
func (o *AuthenticationMethodChain) GetAuthenticationMethods() []AuthenticationMethod {
	if o == nil || o.AuthenticationMethods == nil {
		var ret []AuthenticationMethod
		return ret
	}
	return o.AuthenticationMethods
}

// GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodChain) GetAuthenticationMethodsOk() ([]AuthenticationMethod, bool) {
	if o == nil || o.AuthenticationMethods == nil {
		return nil, false
	}
	return o.AuthenticationMethods, true
}

// HasAuthenticationMethods returns a boolean if a field has been set.
func (o *AuthenticationMethodChain) HasAuthenticationMethods() bool {
	if o != nil && o.AuthenticationMethods != nil {
		return true
	}

	return false
}

// SetAuthenticationMethods gets a reference to the given []AuthenticationMethod and assigns it to the AuthenticationMethods field.
func (o *AuthenticationMethodChain) SetAuthenticationMethods(v []AuthenticationMethod) {
	o.AuthenticationMethods = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *AuthenticationMethodChain) GetNext() []map[string]interface{} {
	if o == nil || o.Next == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodChain) GetNextOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *AuthenticationMethodChain) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given []map[string]interface{} and assigns it to the Next field.
func (o *AuthenticationMethodChain) SetNext(v []map[string]interface{}) {
	o.Next = v
}

// GetReauthenticateIn returns the ReauthenticateIn field value if set, zero value otherwise.
func (o *AuthenticationMethodChain) GetReauthenticateIn() string {
	if o == nil || o.ReauthenticateIn == nil {
		var ret string
		return ret
	}
	return *o.ReauthenticateIn
}

// GetReauthenticateInOk returns a tuple with the ReauthenticateIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethodChain) GetReauthenticateInOk() (*string, bool) {
	if o == nil || o.ReauthenticateIn == nil {
		return nil, false
	}
	return o.ReauthenticateIn, true
}

// HasReauthenticateIn returns a boolean if a field has been set.
func (o *AuthenticationMethodChain) HasReauthenticateIn() bool {
	if o != nil && o.ReauthenticateIn != nil {
		return true
	}

	return false
}

// SetReauthenticateIn gets a reference to the given string and assigns it to the ReauthenticateIn field.
func (o *AuthenticationMethodChain) SetReauthenticateIn(v string) {
	o.ReauthenticateIn = &v
}

func (o AuthenticationMethodChain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationMethods != nil {
		toSerialize["authenticationMethods"] = o.AuthenticationMethods
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.ReauthenticateIn != nil {
		toSerialize["reauthenticateIn"] = o.ReauthenticateIn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticationMethodChain) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticationMethodChain := _AuthenticationMethodChain{}

	err = json.Unmarshal(bytes, &varAuthenticationMethodChain)
	if err == nil {
		*o = AuthenticationMethodChain(varAuthenticationMethodChain)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticationMethods")
		delete(additionalProperties, "next")
		delete(additionalProperties, "reauthenticateIn")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticationMethodChain struct {
	value *AuthenticationMethodChain
	isSet bool
}

func (v NullableAuthenticationMethodChain) Get() *AuthenticationMethodChain {
	return v.value
}

func (v *NullableAuthenticationMethodChain) Set(val *AuthenticationMethodChain) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationMethodChain) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationMethodChain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationMethodChain(val *AuthenticationMethodChain) *NullableAuthenticationMethodChain {
	return &NullableAuthenticationMethodChain{value: val, isSet: true}
}

func (v NullableAuthenticationMethodChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationMethodChain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

