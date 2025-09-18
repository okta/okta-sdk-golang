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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the AuthenticatorEnrollmentLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollmentLinks{}

// AuthenticatorEnrollmentLinks struct for AuthenticatorEnrollmentLinks
type AuthenticatorEnrollmentLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Returns information on the specified user
	User *HrefObject `json:"user,omitempty"`
	// Returns information about a specific authenticator. See [Retrieve an authenticator](/openapi/okta-management/management/tag/Authenticator/#tag/Authenticator/operation/getAuthenticator).
	Authenticator        *HrefObject `json:"authenticator,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollmentLinks AuthenticatorEnrollmentLinks

// NewAuthenticatorEnrollmentLinks instantiates a new AuthenticatorEnrollmentLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollmentLinks() *AuthenticatorEnrollmentLinks {
	this := AuthenticatorEnrollmentLinks{}
	return &this
}

// NewAuthenticatorEnrollmentLinksWithDefaults instantiates a new AuthenticatorEnrollmentLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentLinksWithDefaults() *AuthenticatorEnrollmentLinks {
	this := AuthenticatorEnrollmentLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *AuthenticatorEnrollmentLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentLinks) GetUser() HrefObject {
	if o == nil || IsNil(o.User) {
		var ret HrefObject
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentLinks) GetUserOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentLinks) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given HrefObject and assigns it to the User field.
func (o *AuthenticatorEnrollmentLinks) SetUser(v HrefObject) {
	o.User = &v
}

// GetAuthenticator returns the Authenticator field value if set, zero value otherwise.
func (o *AuthenticatorEnrollmentLinks) GetAuthenticator() HrefObject {
	if o == nil || IsNil(o.Authenticator) {
		var ret HrefObject
		return ret
	}
	return *o.Authenticator
}

// GetAuthenticatorOk returns a tuple with the Authenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollmentLinks) GetAuthenticatorOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Authenticator) {
		return nil, false
	}
	return o.Authenticator, true
}

// HasAuthenticator returns a boolean if a field has been set.
func (o *AuthenticatorEnrollmentLinks) HasAuthenticator() bool {
	if o != nil && !IsNil(o.Authenticator) {
		return true
	}

	return false
}

// SetAuthenticator gets a reference to the given HrefObject and assigns it to the Authenticator field.
func (o *AuthenticatorEnrollmentLinks) SetAuthenticator(v HrefObject) {
	o.Authenticator = &v
}

func (o AuthenticatorEnrollmentLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollmentLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Authenticator) {
		toSerialize["authenticator"] = o.Authenticator
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollmentLinks) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollmentLinks := _AuthenticatorEnrollmentLinks{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollmentLinks)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollmentLinks(varAuthenticatorEnrollmentLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "user")
		delete(additionalProperties, "authenticator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollmentLinks struct {
	value *AuthenticatorEnrollmentLinks
	isSet bool
}

func (v NullableAuthenticatorEnrollmentLinks) Get() *AuthenticatorEnrollmentLinks {
	return v.value
}

func (v *NullableAuthenticatorEnrollmentLinks) Set(val *AuthenticatorEnrollmentLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollmentLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollmentLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollmentLinks(val *AuthenticatorEnrollmentLinks) *NullableAuthenticatorEnrollmentLinks {
	return &NullableAuthenticatorEnrollmentLinks{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollmentLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollmentLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
