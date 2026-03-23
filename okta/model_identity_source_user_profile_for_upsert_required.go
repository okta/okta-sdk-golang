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
	"fmt"
)

// checks if the IdentitySourceUserProfileForUpsertRequired type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentitySourceUserProfileForUpsertRequired{}

// IdentitySourceUserProfileForUpsertRequired struct for IdentitySourceUserProfileForUpsertRequired
type IdentitySourceUserProfileForUpsertRequired struct {
	// Email address of the user
	Email string `json:"email"`
	// First name of the user
	FirstName NullableString `json:"firstName,omitempty"`
	// Home address of the user
	HomeAddress NullableString `json:"homeAddress,omitempty"`
	// Last name of the user
	LastName NullableString `json:"lastName,omitempty"`
	// Mobile phone number of the user
	MobilePhone NullableString `json:"mobilePhone,omitempty"`
	// Alternative email address of the user
	SecondEmail *string `json:"secondEmail,omitempty"`
	// Username of the user
	UserName             string `json:"userName"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySourceUserProfileForUpsertRequired IdentitySourceUserProfileForUpsertRequired

// NewIdentitySourceUserProfileForUpsertRequired instantiates a new IdentitySourceUserProfileForUpsertRequired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySourceUserProfileForUpsertRequired(email string, userName string) *IdentitySourceUserProfileForUpsertRequired {
	this := IdentitySourceUserProfileForUpsertRequired{}
	this.Email = email
	this.UserName = userName
	return &this
}

// NewIdentitySourceUserProfileForUpsertRequiredWithDefaults instantiates a new IdentitySourceUserProfileForUpsertRequired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySourceUserProfileForUpsertRequiredWithDefaults() *IdentitySourceUserProfileForUpsertRequired {
	this := IdentitySourceUserProfileForUpsertRequired{}
	return &this
}

// GetEmail returns the Email field value
func (o *IdentitySourceUserProfileForUpsertRequired) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *IdentitySourceUserProfileForUpsertRequired) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *IdentitySourceUserProfileForUpsertRequired) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourceUserProfileForUpsertRequired) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourceUserProfileForUpsertRequired) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsertRequired) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *IdentitySourceUserProfileForUpsertRequired) SetFirstName(v string) {
	o.FirstName.Set(&v)
}

// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *IdentitySourceUserProfileForUpsertRequired) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *IdentitySourceUserProfileForUpsertRequired) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetHomeAddress returns the HomeAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourceUserProfileForUpsertRequired) GetHomeAddress() string {
	if o == nil || IsNil(o.HomeAddress.Get()) {
		var ret string
		return ret
	}
	return *o.HomeAddress.Get()
}

// GetHomeAddressOk returns a tuple with the HomeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourceUserProfileForUpsertRequired) GetHomeAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HomeAddress.Get(), o.HomeAddress.IsSet()
}

// HasHomeAddress returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsertRequired) HasHomeAddress() bool {
	if o != nil && o.HomeAddress.IsSet() {
		return true
	}

	return false
}

// SetHomeAddress gets a reference to the given NullableString and assigns it to the HomeAddress field.
func (o *IdentitySourceUserProfileForUpsertRequired) SetHomeAddress(v string) {
	o.HomeAddress.Set(&v)
}

// SetHomeAddressNil sets the value for HomeAddress to be an explicit nil
func (o *IdentitySourceUserProfileForUpsertRequired) SetHomeAddressNil() {
	o.HomeAddress.Set(nil)
}

// UnsetHomeAddress ensures that no value is present for HomeAddress, not even an explicit nil
func (o *IdentitySourceUserProfileForUpsertRequired) UnsetHomeAddress() {
	o.HomeAddress.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourceUserProfileForUpsertRequired) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourceUserProfileForUpsertRequired) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsertRequired) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *IdentitySourceUserProfileForUpsertRequired) SetLastName(v string) {
	o.LastName.Set(&v)
}

// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *IdentitySourceUserProfileForUpsertRequired) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *IdentitySourceUserProfileForUpsertRequired) UnsetLastName() {
	o.LastName.Unset()
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourceUserProfileForUpsertRequired) GetMobilePhone() string {
	if o == nil || IsNil(o.MobilePhone.Get()) {
		var ret string
		return ret
	}
	return *o.MobilePhone.Get()
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourceUserProfileForUpsertRequired) GetMobilePhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobilePhone.Get(), o.MobilePhone.IsSet()
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsertRequired) HasMobilePhone() bool {
	if o != nil && o.MobilePhone.IsSet() {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given NullableString and assigns it to the MobilePhone field.
func (o *IdentitySourceUserProfileForUpsertRequired) SetMobilePhone(v string) {
	o.MobilePhone.Set(&v)
}

// SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil
func (o *IdentitySourceUserProfileForUpsertRequired) SetMobilePhoneNil() {
	o.MobilePhone.Set(nil)
}

// UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
func (o *IdentitySourceUserProfileForUpsertRequired) UnsetMobilePhone() {
	o.MobilePhone.Unset()
}

// GetSecondEmail returns the SecondEmail field value if set, zero value otherwise.
func (o *IdentitySourceUserProfileForUpsertRequired) GetSecondEmail() string {
	if o == nil || IsNil(o.SecondEmail) {
		var ret string
		return ret
	}
	return *o.SecondEmail
}

// GetSecondEmailOk returns a tuple with the SecondEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceUserProfileForUpsertRequired) GetSecondEmailOk() (*string, bool) {
	if o == nil || IsNil(o.SecondEmail) {
		return nil, false
	}
	return o.SecondEmail, true
}

// HasSecondEmail returns a boolean if a field has been set.
func (o *IdentitySourceUserProfileForUpsertRequired) HasSecondEmail() bool {
	if o != nil && !IsNil(o.SecondEmail) {
		return true
	}

	return false
}

// SetSecondEmail gets a reference to the given string and assigns it to the SecondEmail field.
func (o *IdentitySourceUserProfileForUpsertRequired) SetSecondEmail(v string) {
	o.SecondEmail = &v
}

// GetUserName returns the UserName field value
func (o *IdentitySourceUserProfileForUpsertRequired) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *IdentitySourceUserProfileForUpsertRequired) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *IdentitySourceUserProfileForUpsertRequired) SetUserName(v string) {
	o.UserName = v
}

func (o IdentitySourceUserProfileForUpsertRequired) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentitySourceUserProfileForUpsertRequired) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
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
	if !IsNil(o.SecondEmail) {
		toSerialize["secondEmail"] = o.SecondEmail
	}
	toSerialize["userName"] = o.UserName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentitySourceUserProfileForUpsertRequired) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"userName",
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

	varIdentitySourceUserProfileForUpsertRequired := _IdentitySourceUserProfileForUpsertRequired{}

	err = json.Unmarshal(data, &varIdentitySourceUserProfileForUpsertRequired)

	if err != nil {
		return err
	}

	*o = IdentitySourceUserProfileForUpsertRequired(varIdentitySourceUserProfileForUpsertRequired)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "homeAddress")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "mobilePhone")
		delete(additionalProperties, "secondEmail")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentitySourceUserProfileForUpsertRequired struct {
	value *IdentitySourceUserProfileForUpsertRequired
	isSet bool
}

func (v NullableIdentitySourceUserProfileForUpsertRequired) Get() *IdentitySourceUserProfileForUpsertRequired {
	return v.value
}

func (v *NullableIdentitySourceUserProfileForUpsertRequired) Set(val *IdentitySourceUserProfileForUpsertRequired) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySourceUserProfileForUpsertRequired) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySourceUserProfileForUpsertRequired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySourceUserProfileForUpsertRequired(val *IdentitySourceUserProfileForUpsertRequired) *NullableIdentitySourceUserProfileForUpsertRequired {
	return &NullableIdentitySourceUserProfileForUpsertRequired{value: val, isSet: true}
}

func (v NullableIdentitySourceUserProfileForUpsertRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySourceUserProfileForUpsertRequired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
