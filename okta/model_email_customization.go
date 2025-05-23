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
	"time"
)

// EmailCustomization struct for EmailCustomization
type EmailCustomization struct {
	// The HTML body of the email. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references).  <x-lifecycle class=\"ea\"></x-lifecycle> Not required if Custom languages for Okta Email Templates is enabled. A `null` body is replaced with a default value from one of the following in priority order:  1. An existing default email customization, if one exists 2. Okta-provided translated content for the specified language, if one exists 3. Okta-provided translated content for the brand locale, if it's set 4. Okta-provided content in English 
	Body string `json:"body"`
	// The email subject. May contain [variable references](https://velocity.apache.org/engine/1.7/user-guide.html#references).  <x-lifecycle class=\"ea\"></x-lifecycle> Not required if Custom languages for Okta Email Templates is enabled. A `null` subject is replaced with a default value from one of the following in priority order:  1. An existing default email customization, if one exists 2. Okta-provided translated content for the specified language, if one exists 3. Okta-provided translated content for the brand locale, if it's set 4. Okta-provided content in English 
	Subject string `json:"subject"`
	// The UTC time at which this email customization was created.
	Created *time.Time `json:"created,omitempty"`
	// A unique identifier for this email customization
	Id *string `json:"id,omitempty"`
	// Whether this is the default customization for the email template. Each customized email template must have exactly one default customization. Defaults to `true` for the first customization and `false` thereafter.
	IsDefault *bool `json:"isDefault,omitempty"`
	// The language specified as an [IETF BCP 47 language tag](https://datatracker.ietf.org/doc/html/rfc5646)
	Language string `json:"language"`
	// The UTC time at which this email customization was last updated.
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Links *EmailCustomizationAllOfLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailCustomization EmailCustomization

// NewEmailCustomization instantiates a new EmailCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailCustomization(body string, subject string, language string) *EmailCustomization {
	this := EmailCustomization{}
	this.Body = body
	this.Subject = subject
	this.Language = language
	return &this
}

// NewEmailCustomizationWithDefaults instantiates a new EmailCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailCustomizationWithDefaults() *EmailCustomization {
	this := EmailCustomization{}
	return &this
}

// GetBody returns the Body field value
func (o *EmailCustomization) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *EmailCustomization) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *EmailCustomization) SetBody(v string) {
	o.Body = v
}

// GetSubject returns the Subject field value
func (o *EmailCustomization) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *EmailCustomization) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *EmailCustomization) SetSubject(v string) {
	o.Subject = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *EmailCustomization) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomization) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EmailCustomization) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *EmailCustomization) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailCustomization) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomization) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailCustomization) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailCustomization) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *EmailCustomization) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomization) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *EmailCustomization) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *EmailCustomization) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLanguage returns the Language field value
func (o *EmailCustomization) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *EmailCustomization) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *EmailCustomization) SetLanguage(v string) {
	o.Language = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *EmailCustomization) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomization) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *EmailCustomization) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *EmailCustomization) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailCustomization) GetLinks() EmailCustomizationAllOfLinks {
	if o == nil || o.Links == nil {
		var ret EmailCustomizationAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomization) GetLinksOk() (*EmailCustomizationAllOfLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailCustomization) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EmailCustomizationAllOfLinks and assigns it to the Links field.
func (o *EmailCustomization) SetLinks(v EmailCustomizationAllOfLinks) {
	o.Links = &v
}

func (o EmailCustomization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["body"] = o.Body
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailCustomization) UnmarshalJSON(bytes []byte) (err error) {
	varEmailCustomization := _EmailCustomization{}

	err = json.Unmarshal(bytes, &varEmailCustomization)
	if err == nil {
		*o = EmailCustomization(varEmailCustomization)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "body")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isDefault")
		delete(additionalProperties, "language")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailCustomization struct {
	value *EmailCustomization
	isSet bool
}

func (v NullableEmailCustomization) Get() *EmailCustomization {
	return v.value
}

func (v *NullableEmailCustomization) Set(val *EmailCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailCustomization(val *EmailCustomization) *NullableEmailCustomization {
	return &NullableEmailCustomization{value: val, isSet: true}
}

func (v NullableEmailCustomization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

