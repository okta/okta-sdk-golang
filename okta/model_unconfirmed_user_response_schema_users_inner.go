/*
Okta Admin Management API

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

// checks if the UnconfirmedUserResponseSchemaUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnconfirmedUserResponseSchemaUsersInner{}

// UnconfirmedUserResponseSchemaUsersInner struct for UnconfirmedUserResponseSchemaUsersInner
type UnconfirmedUserResponseSchemaUsersInner struct {
	// The ID of the user in the target app that's linked to the Okta application user object. This value is the native app-specific identifier or primary key for the user in the target app.  The `externalId` is set during import when the user is confirmed (reconciled) or during provisioning when the user is created in the target app. This value isn't populated for SSO app assignments (for example, SAML or SWA) because it isn't synchronized with a target app.
	ExternalId  *string             `json:"externalId,omitempty"`
	Credentials *AppUserCredentials `json:"credentials,omitempty"`
	// Specifies the default and custom profile properties for a user. Properties that are visible in the Admin Console for an app assignment can also be assigned through the API. Some properties are reference properties that are imported from the target app and can't be configured. See [profile](/openapi/okta-management/management/user/getuser#user/getuser/t=response&c=200&path=profile).
	Profile              map[string]interface{} `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UnconfirmedUserResponseSchemaUsersInner UnconfirmedUserResponseSchemaUsersInner

// NewUnconfirmedUserResponseSchemaUsersInner instantiates a new UnconfirmedUserResponseSchemaUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnconfirmedUserResponseSchemaUsersInner() *UnconfirmedUserResponseSchemaUsersInner {
	this := UnconfirmedUserResponseSchemaUsersInner{}
	return &this
}

// NewUnconfirmedUserResponseSchemaUsersInnerWithDefaults instantiates a new UnconfirmedUserResponseSchemaUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnconfirmedUserResponseSchemaUsersInnerWithDefaults() *UnconfirmedUserResponseSchemaUsersInner {
	this := UnconfirmedUserResponseSchemaUsersInner{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *UnconfirmedUserResponseSchemaUsersInner) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnconfirmedUserResponseSchemaUsersInner) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *UnconfirmedUserResponseSchemaUsersInner) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *UnconfirmedUserResponseSchemaUsersInner) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *UnconfirmedUserResponseSchemaUsersInner) GetCredentials() AppUserCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret AppUserCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnconfirmedUserResponseSchemaUsersInner) GetCredentialsOk() (*AppUserCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *UnconfirmedUserResponseSchemaUsersInner) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given AppUserCredentials and assigns it to the Credentials field.
func (o *UnconfirmedUserResponseSchemaUsersInner) SetCredentials(v AppUserCredentials) {
	o.Credentials = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UnconfirmedUserResponseSchemaUsersInner) GetProfile() map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnconfirmedUserResponseSchemaUsersInner) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UnconfirmedUserResponseSchemaUsersInner) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *UnconfirmedUserResponseSchemaUsersInner) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

func (o UnconfirmedUserResponseSchemaUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnconfirmedUserResponseSchemaUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UnconfirmedUserResponseSchemaUsersInner) UnmarshalJSON(data []byte) (err error) {
	varUnconfirmedUserResponseSchemaUsersInner := _UnconfirmedUserResponseSchemaUsersInner{}

	err = json.Unmarshal(data, &varUnconfirmedUserResponseSchemaUsersInner)

	if err != nil {
		return err
	}

	*o = UnconfirmedUserResponseSchemaUsersInner(varUnconfirmedUserResponseSchemaUsersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUnconfirmedUserResponseSchemaUsersInner struct {
	value *UnconfirmedUserResponseSchemaUsersInner
	isSet bool
}

func (v NullableUnconfirmedUserResponseSchemaUsersInner) Get() *UnconfirmedUserResponseSchemaUsersInner {
	return v.value
}

func (v *NullableUnconfirmedUserResponseSchemaUsersInner) Set(val *UnconfirmedUserResponseSchemaUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUnconfirmedUserResponseSchemaUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUnconfirmedUserResponseSchemaUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnconfirmedUserResponseSchemaUsersInner(val *UnconfirmedUserResponseSchemaUsersInner) *NullableUnconfirmedUserResponseSchemaUsersInner {
	return &NullableUnconfirmedUserResponseSchemaUsersInner{value: val, isSet: true}
}

func (v NullableUnconfirmedUserResponseSchemaUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnconfirmedUserResponseSchemaUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
