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

// EmailSettings struct for EmailSettings
type EmailSettings struct {
	Recipients string `json:"recipients"`
	AdditionalProperties map[string]interface{}
}

type _EmailSettings EmailSettings

// NewEmailSettings instantiates a new EmailSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSettings(recipients string) *EmailSettings {
	this := EmailSettings{}
	this.Recipients = recipients
	return &this
}

// NewEmailSettingsWithDefaults instantiates a new EmailSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSettingsWithDefaults() *EmailSettings {
	this := EmailSettings{}
	return &this
}

// GetRecipients returns the Recipients field value
func (o *EmailSettings) GetRecipients() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *EmailSettings) GetRecipientsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *EmailSettings) SetRecipients(v string) {
	o.Recipients = v
}

func (o EmailSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipients"] = o.Recipients
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailSettings) UnmarshalJSON(bytes []byte) (err error) {
	varEmailSettings := _EmailSettings{}

	err = json.Unmarshal(bytes, &varEmailSettings)
	if err == nil {
		*o = EmailSettings(varEmailSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "recipients")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailSettings struct {
	value *EmailSettings
	isSet bool
}

func (v NullableEmailSettings) Get() *EmailSettings {
	return v.value
}

func (v *NullableEmailSettings) Set(val *EmailSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSettings(val *EmailSettings) *NullableEmailSettings {
	return &NullableEmailSettings{value: val, isSet: true}
}

func (v NullableEmailSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

