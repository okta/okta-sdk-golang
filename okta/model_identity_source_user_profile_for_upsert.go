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

// IdentitySourceUserProfileForUpsert struct for IdentitySourceUserProfileForUpsert
type IdentitySourceUserProfileForUpsert struct {
	Email *string `json:"email,omitempty"`
	FirstName NullableString `json:"firstName,omitempty"`
	HomeAddress NullableString `json:"homeAddress,omitempty"`
	LastName NullableString `json:"lastName,omitempty"`
	MobilePhone NullableString `json:"mobilePhone,omitempty"`
	SecondEmail *string `json:"secondEmail,omitempty"`
	UserName *string `json:"userName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySourceUserProfileForUpsert IdentitySourceUserProfileForUpsert

// NewIdentitySourceUserProfileForUpsert instantiates a new IdentitySourceUserProfileForUpsert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySourceUserProfileForUpsert() *IdentitySourceUserProfileForUpsert {
	this := IdentitySourceUserProfileForUpsert{}
	return &this
}

// NewIdentitySourceUserProfileForUpsertWithDefaults instantiates a new IdentitySourceUserProfileForUpsert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySourceUserProfileForUpsertWithDefaults() *IdentitySourceUserProfileForUpsert {
	this := IdentitySourceUserProfileForUpsert{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *IdentitySourceUserProfileForUpsert) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceUserProfileForUpsert) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsert) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *IdentitySourceUserProfileForUpsert) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourceUserProfileForUpsert) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourceUserProfileForUpsert) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsert) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *IdentitySourceUserProfileForUpsert) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *IdentitySourceUserProfileForUpsert) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *IdentitySourceUserProfileForUpsert) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetHomeAddress returns the HomeAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourceUserProfileForUpsert) GetHomeAddress() string {
	if o == nil || o.HomeAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.HomeAddress.Get()
}

// GetHomeAddressOk returns a tuple with the HomeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourceUserProfileForUpsert) GetHomeAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HomeAddress.Get(), o.HomeAddress.IsSet()
}

// HasHomeAddress returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsert) HasHomeAddress() bool {
	if o != nil && o.HomeAddress.IsSet() {
		return true
	}

	return false
}

// SetHomeAddress gets a reference to the given NullableString and assigns it to the HomeAddress field.
func (o *IdentitySourceUserProfileForUpsert) SetHomeAddress(v string) {
	o.HomeAddress.Set(&v)
}
// SetHomeAddressNil sets the value for HomeAddress to be an explicit nil
func (o *IdentitySourceUserProfileForUpsert) SetHomeAddressNil() {
	o.HomeAddress.Set(nil)
}

// UnsetHomeAddress ensures that no value is present for HomeAddress, not even an explicit nil
func (o *IdentitySourceUserProfileForUpsert) UnsetHomeAddress() {
	o.HomeAddress.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourceUserProfileForUpsert) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourceUserProfileForUpsert) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsert) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *IdentitySourceUserProfileForUpsert) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *IdentitySourceUserProfileForUpsert) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *IdentitySourceUserProfileForUpsert) UnsetLastName() {
	o.LastName.Unset()
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourceUserProfileForUpsert) GetMobilePhone() string {
	if o == nil || o.MobilePhone.Get() == nil {
		var ret string
		return ret
	}
	return *o.MobilePhone.Get()
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourceUserProfileForUpsert) GetMobilePhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobilePhone.Get(), o.MobilePhone.IsSet()
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsert) HasMobilePhone() bool {
	if o != nil && o.MobilePhone.IsSet() {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given NullableString and assigns it to the MobilePhone field.
func (o *IdentitySourceUserProfileForUpsert) SetMobilePhone(v string) {
	o.MobilePhone.Set(&v)
}
// SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil
func (o *IdentitySourceUserProfileForUpsert) SetMobilePhoneNil() {
	o.MobilePhone.Set(nil)
}

// UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
func (o *IdentitySourceUserProfileForUpsert) UnsetMobilePhone() {
	o.MobilePhone.Unset()
}

// GetSecondEmail returns the SecondEmail field value if set, zero value otherwise.
func (o *IdentitySourceUserProfileForUpsert) GetSecondEmail() string {
	if o == nil || o.SecondEmail == nil {
		var ret string
		return ret
	}
	return *o.SecondEmail
}

// GetSecondEmailOk returns a tuple with the SecondEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceUserProfileForUpsert) GetSecondEmailOk() (*string, bool) {
	if o == nil || o.SecondEmail == nil {
		return nil, false
	}
	return o.SecondEmail, true
}

// HasSecondEmail returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsert) HasSecondEmail() bool {
	if o != nil && o.SecondEmail != nil {
		return true
	}

	return false
}

// SetSecondEmail gets a reference to the given string and assigns it to the SecondEmail field.
func (o *IdentitySourceUserProfileForUpsert) SetSecondEmail(v string) {
	o.SecondEmail = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *IdentitySourceUserProfileForUpsert) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceUserProfileForUpsert) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsert) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *IdentitySourceUserProfileForUpsert) SetUserName(v string) {
	o.UserName = &v
}

func (o IdentitySourceUserProfileForUpsert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.HomeAddress.IsSet() {
		toSerialize["homeAddress"] = o.HomeAddress.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.MobilePhone.IsSet() {
		toSerialize["mobilePhone"] = o.MobilePhone.Get()
	}
	if o.SecondEmail != nil {
		toSerialize["secondEmail"] = o.SecondEmail
	}
	if o.UserName != nil {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentitySourceUserProfileForUpsert) UnmarshalJSON(bytes []byte) (err error) {
	varIdentitySourceUserProfileForUpsert := _IdentitySourceUserProfileForUpsert{}

	err = json.Unmarshal(bytes, &varIdentitySourceUserProfileForUpsert)
	if err == nil {
		*o = IdentitySourceUserProfileForUpsert(varIdentitySourceUserProfileForUpsert)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "homeAddress")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "mobilePhone")
		delete(additionalProperties, "secondEmail")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentitySourceUserProfileForUpsert struct {
	value *IdentitySourceUserProfileForUpsert
	isSet bool
}

func (v NullableIdentitySourceUserProfileForUpsert) Get() *IdentitySourceUserProfileForUpsert {
	return v.value
}

func (v *NullableIdentitySourceUserProfileForUpsert) Set(val *IdentitySourceUserProfileForUpsert) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySourceUserProfileForUpsert) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySourceUserProfileForUpsert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySourceUserProfileForUpsert(val *IdentitySourceUserProfileForUpsert) *NullableIdentitySourceUserProfileForUpsert {
	return &NullableIdentitySourceUserProfileForUpsert{value: val, isSet: true}
}

func (v NullableIdentitySourceUserProfileForUpsert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySourceUserProfileForUpsert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

