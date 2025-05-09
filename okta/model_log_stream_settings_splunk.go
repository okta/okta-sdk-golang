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

// LogStreamSettingsSplunk Specifies the configuration for the `splunk_cloud_logstreaming` Log Stream type.
type LogStreamSettingsSplunk struct {
	// Edition of the Splunk Cloud instance
	Edition string `json:"edition"`
	// The domain name for your Splunk Cloud instance. Don't include `http` or `https` in the string. For example: `acme.splunkcloud.com`
	Host string `json:"host"`
	// The HEC token for your Splunk Cloud HTTP Event Collector. The token value is set at object creation, but isn't returned.
	Token string `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamSettingsSplunk LogStreamSettingsSplunk

// NewLogStreamSettingsSplunk instantiates a new LogStreamSettingsSplunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamSettingsSplunk(edition string, host string, token string) *LogStreamSettingsSplunk {
	this := LogStreamSettingsSplunk{}
	this.Edition = edition
	this.Host = host
	this.Token = token
	return &this
}

// NewLogStreamSettingsSplunkWithDefaults instantiates a new LogStreamSettingsSplunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamSettingsSplunkWithDefaults() *LogStreamSettingsSplunk {
	this := LogStreamSettingsSplunk{}
	return &this
}

// GetEdition returns the Edition field value
func (o *LogStreamSettingsSplunk) GetEdition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Edition
}

// GetEditionOk returns a tuple with the Edition field value
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsSplunk) GetEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edition, true
}

// SetEdition sets field value
func (o *LogStreamSettingsSplunk) SetEdition(v string) {
	o.Edition = v
}

// GetHost returns the Host field value
func (o *LogStreamSettingsSplunk) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsSplunk) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *LogStreamSettingsSplunk) SetHost(v string) {
	o.Host = v
}

// GetToken returns the Token field value
func (o *LogStreamSettingsSplunk) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsSplunk) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *LogStreamSettingsSplunk) SetToken(v string) {
	o.Token = v
}

func (o LogStreamSettingsSplunk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["edition"] = o.Edition
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogStreamSettingsSplunk) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamSettingsSplunk := _LogStreamSettingsSplunk{}

	err = json.Unmarshal(bytes, &varLogStreamSettingsSplunk)
	if err == nil {
		*o = LogStreamSettingsSplunk(varLogStreamSettingsSplunk)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "edition")
		delete(additionalProperties, "host")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogStreamSettingsSplunk struct {
	value *LogStreamSettingsSplunk
	isSet bool
}

func (v NullableLogStreamSettingsSplunk) Get() *LogStreamSettingsSplunk {
	return v.value
}

func (v *NullableLogStreamSettingsSplunk) Set(val *LogStreamSettingsSplunk) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamSettingsSplunk) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamSettingsSplunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamSettingsSplunk(val *LogStreamSettingsSplunk) *NullableLogStreamSettingsSplunk {
	return &NullableLogStreamSettingsSplunk{value: val, isSet: true}
}

func (v NullableLogStreamSettingsSplunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamSettingsSplunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

