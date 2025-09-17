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

// checks if the EmailSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailSettings{}

// EmailSettings struct for EmailSettings
type EmailSettings struct {
	Recipients           string `json:"recipients"`
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recipients"] = o.Recipients

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recipients",
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

	varEmailSettings := _EmailSettings{}

	err = json.Unmarshal(data, &varEmailSettings)

	if err != nil {
		return err
	}

	*o = EmailSettings(varEmailSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recipients")
		o.AdditionalProperties = additionalProperties
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
