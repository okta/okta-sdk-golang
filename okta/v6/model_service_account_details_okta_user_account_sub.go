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
	"fmt"
)

// checks if the ServiceAccountDetailsOktaUserAccountSub type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountDetailsOktaUserAccountSub{}

// ServiceAccountDetailsOktaUserAccountSub Details for managing an Okta user as a service account
type ServiceAccountDetailsOktaUserAccountSub struct {
	Credentials *OktaUserServiceAccountCredentials `json:"credentials,omitempty"`
	// The email address for the Okta user
	Email *string `json:"email,omitempty"`
	// The ID of the Okta user to manage as a service account
	OktaUserId           string `json:"oktaUserId"`
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
	if o == nil || IsNil(o.Credentials) {
		var ret OktaUserServiceAccountCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsOktaUserAccountSub) GetCredentialsOk() (*OktaUserServiceAccountCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ServiceAccountDetailsOktaUserAccountSub) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
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
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsOktaUserAccountSub) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ServiceAccountDetailsOktaUserAccountSub) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountDetailsOktaUserAccountSub) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["oktaUserId"] = o.OktaUserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceAccountDetailsOktaUserAccountSub) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"oktaUserId",
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

	varServiceAccountDetailsOktaUserAccountSub := _ServiceAccountDetailsOktaUserAccountSub{}

	err = json.Unmarshal(data, &varServiceAccountDetailsOktaUserAccountSub)

	if err != nil {
		return err
	}

	*o = ServiceAccountDetailsOktaUserAccountSub(varServiceAccountDetailsOktaUserAccountSub)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "email")
		delete(additionalProperties, "oktaUserId")
		o.AdditionalProperties = additionalProperties
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
