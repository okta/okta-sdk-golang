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

// OktaUserServiceAccountCredentials Credentials for an Okta user
type OktaUserServiceAccountCredentials struct {
	// The username associated with the service account
	Username *string `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaUserServiceAccountCredentials OktaUserServiceAccountCredentials

// NewOktaUserServiceAccountCredentials instantiates a new OktaUserServiceAccountCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaUserServiceAccountCredentials() *OktaUserServiceAccountCredentials {
	this := OktaUserServiceAccountCredentials{}
	return &this
}

// NewOktaUserServiceAccountCredentialsWithDefaults instantiates a new OktaUserServiceAccountCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaUserServiceAccountCredentialsWithDefaults() *OktaUserServiceAccountCredentials {
	this := OktaUserServiceAccountCredentials{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OktaUserServiceAccountCredentials) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaUserServiceAccountCredentials) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OktaUserServiceAccountCredentials) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OktaUserServiceAccountCredentials) SetUsername(v string) {
	o.Username = &v
}

func (o OktaUserServiceAccountCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaUserServiceAccountCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varOktaUserServiceAccountCredentials := _OktaUserServiceAccountCredentials{}

	err = json.Unmarshal(bytes, &varOktaUserServiceAccountCredentials)
	if err == nil {
		*o = OktaUserServiceAccountCredentials(varOktaUserServiceAccountCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaUserServiceAccountCredentials struct {
	value *OktaUserServiceAccountCredentials
	isSet bool
}

func (v NullableOktaUserServiceAccountCredentials) Get() *OktaUserServiceAccountCredentials {
	return v.value
}

func (v *NullableOktaUserServiceAccountCredentials) Set(val *OktaUserServiceAccountCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaUserServiceAccountCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaUserServiceAccountCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaUserServiceAccountCredentials(val *OktaUserServiceAccountCredentials) *NullableOktaUserServiceAccountCredentials {
	return &NullableOktaUserServiceAccountCredentials{value: val, isSet: true}
}

func (v NullableOktaUserServiceAccountCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaUserServiceAccountCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

