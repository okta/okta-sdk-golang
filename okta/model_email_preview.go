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

// EmailPreview struct for EmailPreview
type EmailPreview struct {
	// The email's HTML body
	Body *string `json:"body,omitempty"`
	// The email's subject
	Subject *string `json:"subject,omitempty"`
	Links *EmailPreviewLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailPreview EmailPreview

// NewEmailPreview instantiates a new EmailPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailPreview() *EmailPreview {
	this := EmailPreview{}
	return &this
}

// NewEmailPreviewWithDefaults instantiates a new EmailPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailPreviewWithDefaults() *EmailPreview {
	this := EmailPreview{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *EmailPreview) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPreview) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *EmailPreview) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *EmailPreview) SetBody(v string) {
	o.Body = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *EmailPreview) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPreview) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *EmailPreview) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *EmailPreview) SetSubject(v string) {
	o.Subject = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailPreview) GetLinks() EmailPreviewLinks {
	if o == nil || o.Links == nil {
		var ret EmailPreviewLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPreview) GetLinksOk() (*EmailPreviewLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailPreview) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EmailPreviewLinks and assigns it to the Links field.
func (o *EmailPreview) SetLinks(v EmailPreviewLinks) {
	o.Links = &v
}

func (o EmailPreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Subject != nil {
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

func (o *EmailPreview) UnmarshalJSON(bytes []byte) (err error) {
	varEmailPreview := _EmailPreview{}

	err = json.Unmarshal(bytes, &varEmailPreview)
	if err == nil {
		*o = EmailPreview(varEmailPreview)
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

type NullableEmailPreview struct {
	value *EmailPreview
	isSet bool
}

func (v NullableEmailPreview) Get() *EmailPreview {
	return v.value
}

func (v *NullableEmailPreview) Set(val *EmailPreview) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailPreview) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailPreview(val *EmailPreview) *NullableEmailPreview {
	return &NullableEmailPreview{value: val, isSet: true}
}

func (v NullableEmailPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

