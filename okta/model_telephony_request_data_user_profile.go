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

// TelephonyRequestDataUserProfile User profile specifies information about the Okta user
type TelephonyRequestDataUserProfile struct {
	// The user's first name
	FirstName *string `json:"firstName,omitempty"`
	// The user's last name
	LastName *string `json:"lastName,omitempty"`
	// The user's Okta login
	Login *string `json:"login,omitempty"`
	// The user's Okta user ID
	UserId *string `json:"userId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelephonyRequestDataUserProfile TelephonyRequestDataUserProfile

// NewTelephonyRequestDataUserProfile instantiates a new TelephonyRequestDataUserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephonyRequestDataUserProfile() *TelephonyRequestDataUserProfile {
	this := TelephonyRequestDataUserProfile{}
	return &this
}

// NewTelephonyRequestDataUserProfileWithDefaults instantiates a new TelephonyRequestDataUserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephonyRequestDataUserProfileWithDefaults() *TelephonyRequestDataUserProfile {
	this := TelephonyRequestDataUserProfile{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *TelephonyRequestDataUserProfile) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataUserProfile) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *TelephonyRequestDataUserProfile) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *TelephonyRequestDataUserProfile) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *TelephonyRequestDataUserProfile) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataUserProfile) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *TelephonyRequestDataUserProfile) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *TelephonyRequestDataUserProfile) SetLastName(v string) {
	o.LastName = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *TelephonyRequestDataUserProfile) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataUserProfile) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *TelephonyRequestDataUserProfile) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *TelephonyRequestDataUserProfile) SetLogin(v string) {
	o.Login = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *TelephonyRequestDataUserProfile) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataUserProfile) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *TelephonyRequestDataUserProfile) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *TelephonyRequestDataUserProfile) SetUserId(v string) {
	o.UserId = &v
}

func (o TelephonyRequestDataUserProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelephonyRequestDataUserProfile) UnmarshalJSON(bytes []byte) (err error) {
	varTelephonyRequestDataUserProfile := _TelephonyRequestDataUserProfile{}

	err = json.Unmarshal(bytes, &varTelephonyRequestDataUserProfile)
	if err == nil {
		*o = TelephonyRequestDataUserProfile(varTelephonyRequestDataUserProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "login")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTelephonyRequestDataUserProfile struct {
	value *TelephonyRequestDataUserProfile
	isSet bool
}

func (v NullableTelephonyRequestDataUserProfile) Get() *TelephonyRequestDataUserProfile {
	return v.value
}

func (v *NullableTelephonyRequestDataUserProfile) Set(val *TelephonyRequestDataUserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephonyRequestDataUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephonyRequestDataUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephonyRequestDataUserProfile(val *TelephonyRequestDataUserProfile) *NullableTelephonyRequestDataUserProfile {
	return &NullableTelephonyRequestDataUserProfile{value: val, isSet: true}
}

func (v NullableTelephonyRequestDataUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephonyRequestDataUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

