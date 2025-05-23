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

// LogStreamSettingsSplunkPut Specifies the configuration for the `splunk_cloud_logstreaming` Log Stream type.
type LogStreamSettingsSplunkPut struct {
	// Edition of the Splunk Cloud instance
	Edition string `json:"edition"`
	// The domain name for your Splunk Cloud instance. Don't include `http` or `https` in the string. For example: `acme.splunkcloud.com`
	Host string `json:"host"`
	AdditionalProperties map[string]interface{}
}

type _LogStreamSettingsSplunkPut LogStreamSettingsSplunkPut

// NewLogStreamSettingsSplunkPut instantiates a new LogStreamSettingsSplunkPut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStreamSettingsSplunkPut(edition string, host string) *LogStreamSettingsSplunkPut {
	this := LogStreamSettingsSplunkPut{}
	this.Edition = edition
	this.Host = host
	return &this
}

// NewLogStreamSettingsSplunkPutWithDefaults instantiates a new LogStreamSettingsSplunkPut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamSettingsSplunkPutWithDefaults() *LogStreamSettingsSplunkPut {
	this := LogStreamSettingsSplunkPut{}
	return &this
}

// GetEdition returns the Edition field value
func (o *LogStreamSettingsSplunkPut) GetEdition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Edition
}

// GetEditionOk returns a tuple with the Edition field value
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsSplunkPut) GetEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edition, true
}

// SetEdition sets field value
func (o *LogStreamSettingsSplunkPut) SetEdition(v string) {
	o.Edition = v
}

// GetHost returns the Host field value
func (o *LogStreamSettingsSplunkPut) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *LogStreamSettingsSplunkPut) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *LogStreamSettingsSplunkPut) SetHost(v string) {
	o.Host = v
}

func (o LogStreamSettingsSplunkPut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["edition"] = o.Edition
	}
	if true {
		toSerialize["host"] = o.Host
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogStreamSettingsSplunkPut) UnmarshalJSON(bytes []byte) (err error) {
	varLogStreamSettingsSplunkPut := _LogStreamSettingsSplunkPut{}

	err = json.Unmarshal(bytes, &varLogStreamSettingsSplunkPut)
	if err == nil {
		*o = LogStreamSettingsSplunkPut(varLogStreamSettingsSplunkPut)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "edition")
		delete(additionalProperties, "host")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogStreamSettingsSplunkPut struct {
	value *LogStreamSettingsSplunkPut
	isSet bool
}

func (v NullableLogStreamSettingsSplunkPut) Get() *LogStreamSettingsSplunkPut {
	return v.value
}

func (v *NullableLogStreamSettingsSplunkPut) Set(val *LogStreamSettingsSplunkPut) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStreamSettingsSplunkPut) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStreamSettingsSplunkPut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStreamSettingsSplunkPut(val *LogStreamSettingsSplunkPut) *NullableLogStreamSettingsSplunkPut {
	return &NullableLogStreamSettingsSplunkPut{value: val, isSet: true}
}

func (v NullableLogStreamSettingsSplunkPut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStreamSettingsSplunkPut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

