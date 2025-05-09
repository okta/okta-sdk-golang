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

// EmailDefaultContent struct for EmailDefaultContent
type EmailDefaultContent struct {
	// The HTML body of the email. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references).  <x-lifecycle class=\"ea\"></x-lifecycle> Not required if Custom languages for Okta Email Templates is enabled. A `null` body is replaced with a default value from one of the following in priority order:  1. An existing default email customization, if one exists 2. Okta-provided translated content for the specified language, if one exists 3. Okta-provided translated content for the brand locale, if it's set 4. Okta-provided content in English 
	Body string `json:"body"`
	// The email subject. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references).  <x-lifecycle class=\"ea\"></x-lifecycle> Not required if Custom languages for Okta Email Templates is enabled. A `null` subject is replaced with a default value from one of the following in priority order:  1. An existing default email customization, if one exists 2. Okta-provided translated content for the specified language, if one exists 3. Okta-provided translated content for the brand locale, if it's set 4. Okta-provided content in English 
	Subject string `json:"subject"`
	Links *EmailDefaultContentAllOfLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailDefaultContent EmailDefaultContent

// NewEmailDefaultContent instantiates a new EmailDefaultContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDefaultContent(body string, subject string) *EmailDefaultContent {
	this := EmailDefaultContent{}
	this.Body = body
	this.Subject = subject
	return &this
}

// NewEmailDefaultContentWithDefaults instantiates a new EmailDefaultContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDefaultContentWithDefaults() *EmailDefaultContent {
	this := EmailDefaultContent{}
	return &this
}

// GetBody returns the Body field value
func (o *EmailDefaultContent) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EmailDefaultContent) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *EmailDefaultContent) SetBody(v string) {
	o.Body = v
}

// GetSubject returns the Subject field value
func (o *EmailDefaultContent) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EmailDefaultContent) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EmailDefaultContent) SetSubject(v string) {
	o.Subject = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailDefaultContent) GetLinks() EmailDefaultContentAllOfLinks {
	if o == nil || o.Links == nil {
		var ret EmailDefaultContentAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDefaultContent) GetLinksOk() (*EmailDefaultContentAllOfLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailDefaultContent) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EmailDefaultContentAllOfLinks and assigns it to the Links field.
func (o *EmailDefaultContent) SetLinks(v EmailDefaultContentAllOfLinks) {
	o.Links = &v
}

func (o EmailDefaultContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["body"] = o.Body
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailDefaultContent) UnmarshalJSON(bytes []byte) (err error) {
	varEmailDefaultContent := _EmailDefaultContent{}

	err = json.Unmarshal(bytes, &varEmailDefaultContent)
	if err == nil {
		*o = EmailDefaultContent(varEmailDefaultContent)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailDefaultContent struct {
	value *EmailDefaultContent
	isSet bool
}

func (v NullableEmailDefaultContent) Get() *EmailDefaultContent {
	return v.value
}

func (v *NullableEmailDefaultContent) Set(val *EmailDefaultContent) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDefaultContent) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDefaultContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDefaultContent(val *EmailDefaultContent) *NullableEmailDefaultContent {
	return &NullableEmailDefaultContent{value: val, isSet: true}
}

func (v NullableEmailDefaultContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDefaultContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

