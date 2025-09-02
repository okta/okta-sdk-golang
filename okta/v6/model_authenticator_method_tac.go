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

// AuthenticatorMethodTac struct for AuthenticatorMethodTac
type AuthenticatorMethodTac struct {
	Status *string `json:"status,omitempty"`
	// The type of authenticator method
	Type *string `json:"type,omitempty"`
	Links *LinksSelfAndLifecycle `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodTac AuthenticatorMethodTac

// NewAuthenticatorMethodTac instantiates a new AuthenticatorMethodTac object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodTac() *AuthenticatorMethodTac {
	this := AuthenticatorMethodTac{}
	return &this
}

// NewAuthenticatorMethodTacWithDefaults instantiates a new AuthenticatorMethodTac object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodTacWithDefaults() *AuthenticatorMethodTac {
	this := AuthenticatorMethodTac{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthenticatorMethodTac) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodTac) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthenticatorMethodTac) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthenticatorMethodTac) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticatorMethodTac) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodTac) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticatorMethodTac) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthenticatorMethodTac) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthenticatorMethodTac) GetLinks() LinksSelfAndLifecycle {
	if o == nil || o.Links == nil {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodTac) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthenticatorMethodTac) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *AuthenticatorMethodTac) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o AuthenticatorMethodTac) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorMethodTac) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorMethodTac := _AuthenticatorMethodTac{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodTac)
	if err == nil {
		*o = AuthenticatorMethodTac(varAuthenticatorMethodTac)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorMethodTac struct {
	value *AuthenticatorMethodTac
	isSet bool
}

func (v NullableAuthenticatorMethodTac) Get() *AuthenticatorMethodTac {
	return v.value
}

func (v *NullableAuthenticatorMethodTac) Set(val *AuthenticatorMethodTac) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodTac) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodTac) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodTac(val *AuthenticatorMethodTac) *NullableAuthenticatorMethodTac {
	return &NullableAuthenticatorMethodTac{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodTac) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodTac) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

