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

// EmailServerPost struct for EmailServerPost
type EmailServerPost struct {
	// Human-readable name for your SMTP server
	Alias string `json:"alias"`
	// If `true`, routes all email traffic through your SMTP server
	Enabled *bool `json:"enabled,omitempty"`
	// Hostname or IP address of your SMTP server
	Host string `json:"host"`
	// Port number of your SMTP server
	Port int32 `json:"port"`
	// Username used to access your SMTP server
	Username string `json:"username"`
	// Password used to access your SMTP server
	Password string `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _EmailServerPost EmailServerPost

// NewEmailServerPost instantiates a new EmailServerPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailServerPost(alias string, host string, port int32, username string, password string) *EmailServerPost {
	this := EmailServerPost{}
	this.Alias = alias
	this.Host = host
	this.Port = port
	this.Username = username
	this.Password = password
	return &this
}

// NewEmailServerPostWithDefaults instantiates a new EmailServerPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailServerPostWithDefaults() *EmailServerPost {
	this := EmailServerPost{}
	return &this
}

// GetAlias returns the Alias field value
func (o *EmailServerPost) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *EmailServerPost) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *EmailServerPost) SetAlias(v string) {
	o.Alias = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EmailServerPost) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerPost) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EmailServerPost) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EmailServerPost) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value
func (o *EmailServerPost) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *EmailServerPost) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *EmailServerPost) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
func (o *EmailServerPost) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *EmailServerPost) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *EmailServerPost) SetPort(v int32) {
	o.Port = v
}

// GetUsername returns the Username field value
func (o *EmailServerPost) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *EmailServerPost) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *EmailServerPost) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *EmailServerPost) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *EmailServerPost) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *EmailServerPost) SetPassword(v string) {
	o.Password = v
}

func (o EmailServerPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["alias"] = o.Alias
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailServerPost) UnmarshalJSON(bytes []byte) (err error) {
	varEmailServerPost := _EmailServerPost{}

	err = json.Unmarshal(bytes, &varEmailServerPost)
	if err == nil {
		*o = EmailServerPost(varEmailServerPost)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "alias")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailServerPost struct {
	value *EmailServerPost
	isSet bool
}

func (v NullableEmailServerPost) Get() *EmailServerPost {
	return v.value
}

func (v *NullableEmailServerPost) Set(val *EmailServerPost) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailServerPost) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailServerPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailServerPost(val *EmailServerPost) *NullableEmailServerPost {
	return &NullableEmailServerPost{value: val, isSet: true}
}

func (v NullableEmailServerPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailServerPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

