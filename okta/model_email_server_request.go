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

// EmailServerRequest struct for EmailServerRequest
type EmailServerRequest struct {
	// Human-readable name for your SMTP server
	Alias *string `json:"alias,omitempty"`
	// If `true`, routes all email traffic through your SMTP server
	Enabled *bool `json:"enabled,omitempty"`
	// Hostname or IP address of your SMTP server
	Host *string `json:"host,omitempty"`
	// Port number of your SMTP server
	Port *int32 `json:"port,omitempty"`
	// Username used to access your SMTP server
	Username *string `json:"username,omitempty"`
	// Password used to access your SMTP server
	Password *string `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailServerRequest EmailServerRequest

// NewEmailServerRequest instantiates a new EmailServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailServerRequest() *EmailServerRequest {
	this := EmailServerRequest{}
	return &this
}

// NewEmailServerRequestWithDefaults instantiates a new EmailServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailServerRequestWithDefaults() *EmailServerRequest {
	this := EmailServerRequest{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *EmailServerRequest) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerRequest) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *EmailServerRequest) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *EmailServerRequest) SetAlias(v string) {
	o.Alias = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EmailServerRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EmailServerRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EmailServerRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *EmailServerRequest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerRequest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *EmailServerRequest) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *EmailServerRequest) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *EmailServerRequest) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerRequest) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *EmailServerRequest) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *EmailServerRequest) SetPort(v int32) {
	o.Port = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *EmailServerRequest) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerRequest) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *EmailServerRequest) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *EmailServerRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *EmailServerRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailServerRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *EmailServerRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *EmailServerRequest) SetPassword(v string) {
	o.Password = &v
}

func (o EmailServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailServerRequest) UnmarshalJSON(bytes []byte) (err error) {
	varEmailServerRequest := _EmailServerRequest{}

	err = json.Unmarshal(bytes, &varEmailServerRequest)
	if err == nil {
		*o = EmailServerRequest(varEmailServerRequest)
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

type NullableEmailServerRequest struct {
	value *EmailServerRequest
	isSet bool
}

func (v NullableEmailServerRequest) Get() *EmailServerRequest {
	return v.value
}

func (v *NullableEmailServerRequest) Set(val *EmailServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailServerRequest(val *EmailServerRequest) *NullableEmailServerRequest {
	return &NullableEmailServerRequest{value: val, isSet: true}
}

func (v NullableEmailServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

