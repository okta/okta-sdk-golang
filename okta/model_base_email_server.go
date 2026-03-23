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

// checks if the BaseEmailServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseEmailServer{}

// BaseEmailServer struct for BaseEmailServer
type BaseEmailServer struct {
	// Human-readable name for your SMTP server
	Alias string `json:"alias"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle> <x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>The authentication type that's used by your SMTP server
	AuthType string `json:"authType"`
	// If `true`, all email traffic is routed through your SMTP server
	Enabled bool `json:"enabled"`
	// Hostname or IP address of your SMTP server
	Host string `json:"host"`
	// ID of your SMTP server
	Id *string `json:"id,omitempty"`
	// Port number of your SMTP server
	Port int32 `json:"port"`
	// Username that's used to access your SMTP server
	Username             string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _BaseEmailServer BaseEmailServer

// NewBaseEmailServer instantiates a new BaseEmailServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEmailServer(alias string, authType string, enabled bool, host string, port int32, username string) *BaseEmailServer {
	this := BaseEmailServer{}
	this.Alias = alias
	this.AuthType = authType
	this.Enabled = enabled
	this.Host = host
	this.Port = port
	this.Username = username
	return &this
}

// NewBaseEmailServerWithDefaults instantiates a new BaseEmailServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEmailServerWithDefaults() *BaseEmailServer {
	this := BaseEmailServer{}
	return &this
}

// GetAlias returns the Alias field value
func (o *BaseEmailServer) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *BaseEmailServer) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *BaseEmailServer) SetAlias(v string) {
	o.Alias = v
}

// GetAuthType returns the AuthType field value
func (o *BaseEmailServer) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *BaseEmailServer) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *BaseEmailServer) SetAuthType(v string) {
	o.AuthType = v
}

// GetEnabled returns the Enabled field value
func (o *BaseEmailServer) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *BaseEmailServer) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *BaseEmailServer) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHost returns the Host field value
func (o *BaseEmailServer) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *BaseEmailServer) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *BaseEmailServer) SetHost(v string) {
	o.Host = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseEmailServer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEmailServer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseEmailServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseEmailServer) SetId(v string) {
	o.Id = &v
}

// GetPort returns the Port field value
func (o *BaseEmailServer) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *BaseEmailServer) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *BaseEmailServer) SetPort(v int32) {
	o.Port = v
}

// GetUsername returns the Username field value
func (o *BaseEmailServer) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *BaseEmailServer) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *BaseEmailServer) SetUsername(v string) {
	o.Username = v
}

func (o BaseEmailServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseEmailServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	toSerialize["authType"] = o.AuthType
	toSerialize["enabled"] = o.Enabled
	toSerialize["host"] = o.Host
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["port"] = o.Port
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseEmailServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alias",
		"authType",
		"enabled",
		"host",
		"port",
		"username",
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

	varBaseEmailServer := _BaseEmailServer{}

	err = json.Unmarshal(data, &varBaseEmailServer)

	if err != nil {
		return err
	}

	*o = BaseEmailServer(varBaseEmailServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alias")
		delete(additionalProperties, "authType")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "host")
		delete(additionalProperties, "id")
		delete(additionalProperties, "port")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseEmailServer struct {
	value *BaseEmailServer
	isSet bool
}

func (v NullableBaseEmailServer) Get() *BaseEmailServer {
	return v.value
}

func (v *NullableBaseEmailServer) Set(val *BaseEmailServer) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEmailServer) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEmailServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEmailServer(val *BaseEmailServer) *NullableBaseEmailServer {
	return &NullableBaseEmailServer{value: val, isSet: true}
}

func (v NullableBaseEmailServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEmailServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
