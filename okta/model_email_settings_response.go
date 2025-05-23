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

// EmailSettingsResponse struct for EmailSettingsResponse
type EmailSettingsResponse struct {
	Recipients *string `json:"recipients,omitempty"`
	Links *EmailSettingsResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailSettingsResponse EmailSettingsResponse

// NewEmailSettingsResponse instantiates a new EmailSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSettingsResponse() *EmailSettingsResponse {
	this := EmailSettingsResponse{}
	return &this
}

// NewEmailSettingsResponseWithDefaults instantiates a new EmailSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSettingsResponseWithDefaults() *EmailSettingsResponse {
	this := EmailSettingsResponse{}
	return &this
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *EmailSettingsResponse) GetRecipients() string {
	if o == nil || o.Recipients == nil {
		var ret string
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsResponse) GetRecipientsOk() (*string, bool) {
	if o == nil || o.Recipients == nil {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *EmailSettingsResponse) HasRecipients() bool {
	if o != nil && o.Recipients != nil {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given string and assigns it to the Recipients field.
func (o *EmailSettingsResponse) SetRecipients(v string) {
	o.Recipients = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailSettingsResponse) GetLinks() EmailSettingsResponseLinks {
	if o == nil || o.Links == nil {
		var ret EmailSettingsResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsResponse) GetLinksOk() (*EmailSettingsResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailSettingsResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EmailSettingsResponseLinks and assigns it to the Links field.
func (o *EmailSettingsResponse) SetLinks(v EmailSettingsResponseLinks) {
	o.Links = &v
}

func (o EmailSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Recipients != nil {
		toSerialize["recipients"] = o.Recipients
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEmailSettingsResponse := _EmailSettingsResponse{}

	err = json.Unmarshal(bytes, &varEmailSettingsResponse)
	if err == nil {
		*o = EmailSettingsResponse(varEmailSettingsResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailSettingsResponse struct {
	value *EmailSettingsResponse
	isSet bool
}

func (v NullableEmailSettingsResponse) Get() *EmailSettingsResponse {
	return v.value
}

func (v *NullableEmailSettingsResponse) Set(val *EmailSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSettingsResponse(val *EmailSettingsResponse) *NullableEmailSettingsResponse {
	return &NullableEmailSettingsResponse{value: val, isSet: true}
}

func (v NullableEmailSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

