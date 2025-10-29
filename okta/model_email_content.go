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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the EmailContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailContent{}

// EmailContent struct for EmailContent
type EmailContent struct {
	// The HTML body of the email. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references).  <x-lifecycle class=\"ea\"></x-lifecycle> Not required if Custom languages for Okta Email Templates is enabled. A `null` body is replaced with a default value from one of the following in priority order:  1. An existing default email customization, if one exists 2. Okta-provided translated content for the specified language, if one exists 3. Okta-provided translated content for the brand locale, if it's set 4. Okta-provided content in English
	Body string `json:"body"`
	// The email subject. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references).  <x-lifecycle class=\"ea\"></x-lifecycle> Not required if Custom languages for Okta Email Templates is enabled. A `null` subject is replaced with a default value from one of the following in priority order:  1. An existing default email customization, if one exists 2. Okta-provided translated content for the specified language, if one exists 3. Okta-provided translated content for the brand locale, if it's set 4. Okta-provided content in English
	Subject              string `json:"subject"`
	AdditionalProperties map[string]interface{}
}

type _EmailContent EmailContent

// NewEmailContent instantiates a new EmailContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailContent(body string, subject string) *EmailContent {
	this := EmailContent{}
	this.Body = body
	this.Subject = subject
	return &this
}

// NewEmailContentWithDefaults instantiates a new EmailContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailContentWithDefaults() *EmailContent {
	this := EmailContent{}
	return &this
}

// GetBody returns the Body field value
func (o *EmailContent) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EmailContent) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *EmailContent) SetBody(v string) {
	o.Body = v
}

// GetSubject returns the Subject field value
func (o *EmailContent) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EmailContent) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EmailContent) SetSubject(v string) {
	o.Subject = v
}

func (o EmailContent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["subject"] = o.Subject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"body",
		"subject",
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

	varEmailContent := _EmailContent{}

	err = json.Unmarshal(data, &varEmailContent)

	if err != nil {
		return err
	}

	*o = EmailContent(varEmailContent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailContent struct {
	value *EmailContent
	isSet bool
}

func (v NullableEmailContent) Get() *EmailContent {
	return v.value
}

func (v *NullableEmailContent) Set(val *EmailContent) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailContent) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailContent(val *EmailContent) *NullableEmailContent {
	return &NullableEmailContent{value: val, isSet: true}
}

func (v NullableEmailContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
