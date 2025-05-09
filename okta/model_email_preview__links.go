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

// EmailPreviewLinks struct for EmailPreviewLinks
type EmailPreviewLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	ContentSource *HrefObject `json:"contentSource,omitempty"`
	Template *HrefObject `json:"template,omitempty"`
	Test *HrefObject `json:"test,omitempty"`
	DefaultContent *HrefObject `json:"defaultContent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailPreviewLinks EmailPreviewLinks

// NewEmailPreviewLinks instantiates a new EmailPreviewLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailPreviewLinks() *EmailPreviewLinks {
	this := EmailPreviewLinks{}
	return &this
}

// NewEmailPreviewLinksWithDefaults instantiates a new EmailPreviewLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailPreviewLinksWithDefaults() *EmailPreviewLinks {
	this := EmailPreviewLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EmailPreviewLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPreviewLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EmailPreviewLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *EmailPreviewLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetContentSource returns the ContentSource field value if set, zero value otherwise.
func (o *EmailPreviewLinks) GetContentSource() HrefObject {
	if o == nil || o.ContentSource == nil {
		var ret HrefObject
		return ret
	}
	return *o.ContentSource
}

// GetContentSourceOk returns a tuple with the ContentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPreviewLinks) GetContentSourceOk() (*HrefObject, bool) {
	if o == nil || o.ContentSource == nil {
		return nil, false
	}
	return o.ContentSource, true
}

// HasContentSource returns a boolean if a field has been set.
func (o *EmailPreviewLinks) HasContentSource() bool {
	if o != nil && o.ContentSource != nil {
		return true
	}

	return false
}

// SetContentSource gets a reference to the given HrefObject and assigns it to the ContentSource field.
func (o *EmailPreviewLinks) SetContentSource(v HrefObject) {
	o.ContentSource = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EmailPreviewLinks) GetTemplate() HrefObject {
	if o == nil || o.Template == nil {
		var ret HrefObject
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPreviewLinks) GetTemplateOk() (*HrefObject, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EmailPreviewLinks) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given HrefObject and assigns it to the Template field.
func (o *EmailPreviewLinks) SetTemplate(v HrefObject) {
	o.Template = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *EmailPreviewLinks) GetTest() HrefObject {
	if o == nil || o.Test == nil {
		var ret HrefObject
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPreviewLinks) GetTestOk() (*HrefObject, bool) {
	if o == nil || o.Test == nil {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *EmailPreviewLinks) HasTest() bool {
	if o != nil && o.Test != nil {
		return true
	}

	return false
}

// SetTest gets a reference to the given HrefObject and assigns it to the Test field.
func (o *EmailPreviewLinks) SetTest(v HrefObject) {
	o.Test = &v
}

// GetDefaultContent returns the DefaultContent field value if set, zero value otherwise.
func (o *EmailPreviewLinks) GetDefaultContent() HrefObject {
	if o == nil || o.DefaultContent == nil {
		var ret HrefObject
		return ret
	}
	return *o.DefaultContent
}

// GetDefaultContentOk returns a tuple with the DefaultContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailPreviewLinks) GetDefaultContentOk() (*HrefObject, bool) {
	if o == nil || o.DefaultContent == nil {
		return nil, false
	}
	return o.DefaultContent, true
}

// HasDefaultContent returns a boolean if a field has been set.
func (o *EmailPreviewLinks) HasDefaultContent() bool {
	if o != nil && o.DefaultContent != nil {
		return true
	}

	return false
}

// SetDefaultContent gets a reference to the given HrefObject and assigns it to the DefaultContent field.
func (o *EmailPreviewLinks) SetDefaultContent(v HrefObject) {
	o.DefaultContent = &v
}

func (o EmailPreviewLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.ContentSource != nil {
		toSerialize["contentSource"] = o.ContentSource
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Test != nil {
		toSerialize["test"] = o.Test
	}
	if o.DefaultContent != nil {
		toSerialize["defaultContent"] = o.DefaultContent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailPreviewLinks) UnmarshalJSON(bytes []byte) (err error) {
	varEmailPreviewLinks := _EmailPreviewLinks{}

	err = json.Unmarshal(bytes, &varEmailPreviewLinks)
	if err == nil {
		*o = EmailPreviewLinks(varEmailPreviewLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "contentSource")
		delete(additionalProperties, "template")
		delete(additionalProperties, "test")
		delete(additionalProperties, "defaultContent")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailPreviewLinks struct {
	value *EmailPreviewLinks
	isSet bool
}

func (v NullableEmailPreviewLinks) Get() *EmailPreviewLinks {
	return v.value
}

func (v *NullableEmailPreviewLinks) Set(val *EmailPreviewLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailPreviewLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailPreviewLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailPreviewLinks(val *EmailPreviewLinks) *NullableEmailPreviewLinks {
	return &NullableEmailPreviewLinks{value: val, isSet: true}
}

func (v NullableEmailPreviewLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailPreviewLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

