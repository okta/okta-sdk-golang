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

// AuthenticatorMethodBase struct for AuthenticatorMethodBase
type AuthenticatorMethodBase struct {
	Status *string `json:"status,omitempty"`
	// The type of authenticator method
	Type *string `json:"type,omitempty"`
	Links *LinksSelfAndLifecycle `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorMethodBase AuthenticatorMethodBase

// NewAuthenticatorMethodBase instantiates a new AuthenticatorMethodBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodBase() *AuthenticatorMethodBase {
	this := AuthenticatorMethodBase{}
	return &this
}

// NewAuthenticatorMethodBaseWithDefaults instantiates a new AuthenticatorMethodBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodBaseWithDefaults() *AuthenticatorMethodBase {
	this := AuthenticatorMethodBase{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthenticatorMethodBase) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodBase) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthenticatorMethodBase) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthenticatorMethodBase) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticatorMethodBase) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodBase) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticatorMethodBase) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthenticatorMethodBase) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthenticatorMethodBase) GetLinks() LinksSelfAndLifecycle {
	if o == nil || o.Links == nil {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodBase) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthenticatorMethodBase) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *AuthenticatorMethodBase) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o AuthenticatorMethodBase) MarshalJSON() ([]byte, error) {
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

func (o *AuthenticatorMethodBase) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorMethodBase := _AuthenticatorMethodBase{}

	err = json.Unmarshal(bytes, &varAuthenticatorMethodBase)
	if err == nil {
		*o = AuthenticatorMethodBase(varAuthenticatorMethodBase)
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

type NullableAuthenticatorMethodBase struct {
	value *AuthenticatorMethodBase
	isSet bool
}

func (v NullableAuthenticatorMethodBase) Get() *AuthenticatorMethodBase {
	return v.value
}

func (v *NullableAuthenticatorMethodBase) Set(val *AuthenticatorMethodBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodBase(val *AuthenticatorMethodBase) *NullableAuthenticatorMethodBase {
	return &NullableAuthenticatorMethodBase{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

