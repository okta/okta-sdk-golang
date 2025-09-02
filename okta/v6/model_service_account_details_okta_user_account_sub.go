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

// ServiceAccountDetailsOktaUserAccountSub Details for managing an Okta user as a service account
type ServiceAccountDetailsOktaUserAccountSub struct {
	Credentials *OktaUserServiceAccountCredentials `json:"credentials,omitempty"`
	// The email address for the Okta user
	Email *string `json:"email,omitempty"`
	// The ID of the Okta user to manage as a service account
	OktaUserId string `json:"oktaUserId"`
	AdditionalProperties map[string]interface{}
}

type _ServiceAccountDetailsOktaUserAccountSub ServiceAccountDetailsOktaUserAccountSub

// NewServiceAccountDetailsOktaUserAccountSub instantiates a new ServiceAccountDetailsOktaUserAccountSub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountDetailsOktaUserAccountSub(oktaUserId string) *ServiceAccountDetailsOktaUserAccountSub {
	this := ServiceAccountDetailsOktaUserAccountSub{}
	this.OktaUserId = oktaUserId
	return &this
}

// NewServiceAccountDetailsOktaUserAccountSubWithDefaults instantiates a new ServiceAccountDetailsOktaUserAccountSub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountDetailsOktaUserAccountSubWithDefaults() *ServiceAccountDetailsOktaUserAccountSub {
	this := ServiceAccountDetailsOktaUserAccountSub{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ServiceAccountDetailsOktaUserAccountSub) GetCredentials() OktaUserServiceAccountCredentials {
	if o == nil || o.Credentials == nil {
		var ret OktaUserServiceAccountCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsOktaUserAccountSub) GetCredentialsOk() (*OktaUserServiceAccountCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ServiceAccountDetailsOktaUserAccountSub) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given OktaUserServiceAccountCredentials and assigns it to the Credentials field.
func (o *ServiceAccountDetailsOktaUserAccountSub) SetCredentials(v OktaUserServiceAccountCredentials) {
	o.Credentials = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ServiceAccountDetailsOktaUserAccountSub) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsOktaUserAccountSub) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ServiceAccountDetailsOktaUserAccountSub) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ServiceAccountDetailsOktaUserAccountSub) SetEmail(v string) {
	o.Email = &v
}

// GetOktaUserId returns the OktaUserId field value
func (o *ServiceAccountDetailsOktaUserAccountSub) GetOktaUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OktaUserId
}

// GetOktaUserIdOk returns a tuple with the OktaUserId field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsOktaUserAccountSub) GetOktaUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OktaUserId, true
}

// SetOktaUserId sets field value
func (o *ServiceAccountDetailsOktaUserAccountSub) SetOktaUserId(v string) {
	o.OktaUserId = v
}

func (o ServiceAccountDetailsOktaUserAccountSub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["oktaUserId"] = o.OktaUserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceAccountDetailsOktaUserAccountSub) UnmarshalJSON(bytes []byte) (err error) {
	varServiceAccountDetailsOktaUserAccountSub := _ServiceAccountDetailsOktaUserAccountSub{}

	err = json.Unmarshal(bytes, &varServiceAccountDetailsOktaUserAccountSub)
	if err == nil {
		*o = ServiceAccountDetailsOktaUserAccountSub(varServiceAccountDetailsOktaUserAccountSub)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "email")
		delete(additionalProperties, "oktaUserId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableServiceAccountDetailsOktaUserAccountSub struct {
	value *ServiceAccountDetailsOktaUserAccountSub
	isSet bool
}

func (v NullableServiceAccountDetailsOktaUserAccountSub) Get() *ServiceAccountDetailsOktaUserAccountSub {
	return v.value
}

func (v *NullableServiceAccountDetailsOktaUserAccountSub) Set(val *ServiceAccountDetailsOktaUserAccountSub) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountDetailsOktaUserAccountSub) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountDetailsOktaUserAccountSub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountDetailsOktaUserAccountSub(val *ServiceAccountDetailsOktaUserAccountSub) *NullableServiceAccountDetailsOktaUserAccountSub {
	return &NullableServiceAccountDetailsOktaUserAccountSub{value: val, isSet: true}
}

func (v NullableServiceAccountDetailsOktaUserAccountSub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountDetailsOktaUserAccountSub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

