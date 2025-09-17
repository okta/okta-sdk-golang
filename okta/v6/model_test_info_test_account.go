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

// checks if the TestInfoTestAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestInfoTestAccount{}

// TestInfoTestAccount An account on a test instance of your app with admin privileges. A test admin account is required by Okta for integration testing. During OIN QA testing, an Okta analyst uses this admin account to configure your app for the various test case flows.
type TestInfoTestAccount struct {
	// The sign-in URL to a test instance of your app
	Url string `json:"url"`
	// The username for your app admin account
	Username string `json:"username"`
	// The password for your app admin account
	Password string `json:"password"`
	// Additional instructions to test the app integration, including instructions for obtaining test accounts
	Instructions         *string `json:"instructions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TestInfoTestAccount TestInfoTestAccount

// NewTestInfoTestAccount instantiates a new TestInfoTestAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestInfoTestAccount(url string, username string, password string) *TestInfoTestAccount {
	this := TestInfoTestAccount{}
	this.Url = url
	this.Username = username
	this.Password = password
	return &this
}

// NewTestInfoTestAccountWithDefaults instantiates a new TestInfoTestAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestInfoTestAccountWithDefaults() *TestInfoTestAccount {
	this := TestInfoTestAccount{}
	return &this
}

// GetUrl returns the Url field value
func (o *TestInfoTestAccount) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *TestInfoTestAccount) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *TestInfoTestAccount) SetUrl(v string) {
	o.Url = v
}

// GetUsername returns the Username field value
func (o *TestInfoTestAccount) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *TestInfoTestAccount) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *TestInfoTestAccount) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *TestInfoTestAccount) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *TestInfoTestAccount) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *TestInfoTestAccount) SetPassword(v string) {
	o.Password = v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *TestInfoTestAccount) GetInstructions() string {
	if o == nil || IsNil(o.Instructions) {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfoTestAccount) GetInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *TestInfoTestAccount) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *TestInfoTestAccount) SetInstructions(v string) {
	o.Instructions = &v
}

func (o TestInfoTestAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestInfoTestAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TestInfoTestAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"username",
		"password",
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

	varTestInfoTestAccount := _TestInfoTestAccount{}

	err = json.Unmarshal(data, &varTestInfoTestAccount)

	if err != nil {
		return err
	}

	*o = TestInfoTestAccount(varTestInfoTestAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "instructions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTestInfoTestAccount struct {
	value *TestInfoTestAccount
	isSet bool
}

func (v NullableTestInfoTestAccount) Get() *TestInfoTestAccount {
	return v.value
}

func (v *NullableTestInfoTestAccount) Set(val *TestInfoTestAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInfoTestAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInfoTestAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInfoTestAccount(val *TestInfoTestAccount) *NullableTestInfoTestAccount {
	return &NullableTestInfoTestAccount{value: val, isSet: true}
}

func (v NullableTestInfoTestAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInfoTestAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
