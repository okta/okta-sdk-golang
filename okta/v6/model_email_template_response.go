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
)

// checks if the EmailTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateResponse{}

// EmailTemplateResponse struct for EmailTemplateResponse
type EmailTemplateResponse struct {
	// The name of this email template
	Name                 *string                        `json:"name,omitempty"`
	Embedded             *EmailTemplateResponseEmbedded `json:"_embedded,omitempty"`
	Links                *EmailTemplateResponseLinks    `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailTemplateResponse EmailTemplateResponse

// NewEmailTemplateResponse instantiates a new EmailTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplateResponse() *EmailTemplateResponse {
	this := EmailTemplateResponse{}
	return &this
}

// NewEmailTemplateResponseWithDefaults instantiates a new EmailTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateResponseWithDefaults() *EmailTemplateResponse {
	this := EmailTemplateResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EmailTemplateResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EmailTemplateResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EmailTemplateResponse) SetName(v string) {
	o.Name = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *EmailTemplateResponse) GetEmbedded() EmailTemplateResponseEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret EmailTemplateResponseEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponse) GetEmbeddedOk() (*EmailTemplateResponseEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *EmailTemplateResponse) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EmailTemplateResponseEmbedded and assigns it to the Embedded field.
func (o *EmailTemplateResponse) SetEmbedded(v EmailTemplateResponseEmbedded) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailTemplateResponse) GetLinks() EmailTemplateResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret EmailTemplateResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponse) GetLinksOk() (*EmailTemplateResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailTemplateResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EmailTemplateResponseLinks and assigns it to the Links field.
func (o *EmailTemplateResponse) SetLinks(v EmailTemplateResponseLinks) {
	o.Links = &v
}

func (o EmailTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailTemplateResponse) UnmarshalJSON(data []byte) (err error) {
	varEmailTemplateResponse := _EmailTemplateResponse{}

	err = json.Unmarshal(data, &varEmailTemplateResponse)

	if err != nil {
		return err
	}

	*o = EmailTemplateResponse(varEmailTemplateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailTemplateResponse struct {
	value *EmailTemplateResponse
	isSet bool
}

func (v NullableEmailTemplateResponse) Get() *EmailTemplateResponse {
	return v.value
}

func (v *NullableEmailTemplateResponse) Set(val *EmailTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateResponse(val *EmailTemplateResponse) *NullableEmailTemplateResponse {
	return &NullableEmailTemplateResponse{value: val, isSet: true}
}

func (v NullableEmailTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
