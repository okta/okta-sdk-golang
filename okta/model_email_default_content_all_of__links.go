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

// EmailDefaultContentAllOfLinks struct for EmailDefaultContentAllOfLinks
type EmailDefaultContentAllOfLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Template *HrefObject `json:"template,omitempty"`
	Preview *HrefObject `json:"preview,omitempty"`
	Test *HrefObject `json:"test,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailDefaultContentAllOfLinks EmailDefaultContentAllOfLinks

// NewEmailDefaultContentAllOfLinks instantiates a new EmailDefaultContentAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDefaultContentAllOfLinks() *EmailDefaultContentAllOfLinks {
	this := EmailDefaultContentAllOfLinks{}
	return &this
}

// NewEmailDefaultContentAllOfLinksWithDefaults instantiates a new EmailDefaultContentAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDefaultContentAllOfLinksWithDefaults() *EmailDefaultContentAllOfLinks {
	this := EmailDefaultContentAllOfLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EmailDefaultContentAllOfLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDefaultContentAllOfLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EmailDefaultContentAllOfLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *EmailDefaultContentAllOfLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EmailDefaultContentAllOfLinks) GetTemplate() HrefObject {
	if o == nil || o.Template == nil {
		var ret HrefObject
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDefaultContentAllOfLinks) GetTemplateOk() (*HrefObject, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EmailDefaultContentAllOfLinks) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given HrefObject and assigns it to the Template field.
func (o *EmailDefaultContentAllOfLinks) SetTemplate(v HrefObject) {
	o.Template = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *EmailDefaultContentAllOfLinks) GetPreview() HrefObject {
	if o == nil || o.Preview == nil {
		var ret HrefObject
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDefaultContentAllOfLinks) GetPreviewOk() (*HrefObject, bool) {
	if o == nil || o.Preview == nil {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *EmailDefaultContentAllOfLinks) HasPreview() bool {
	if o != nil && o.Preview != nil {
		return true
	}

	return false
}

// SetPreview gets a reference to the given HrefObject and assigns it to the Preview field.
func (o *EmailDefaultContentAllOfLinks) SetPreview(v HrefObject) {
	o.Preview = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *EmailDefaultContentAllOfLinks) GetTest() HrefObject {
	if o == nil || o.Test == nil {
		var ret HrefObject
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDefaultContentAllOfLinks) GetTestOk() (*HrefObject, bool) {
	if o == nil || o.Test == nil {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *EmailDefaultContentAllOfLinks) HasTest() bool {
	if o != nil && o.Test != nil {
		return true
	}

	return false
}

// SetTest gets a reference to the given HrefObject and assigns it to the Test field.
func (o *EmailDefaultContentAllOfLinks) SetTest(v HrefObject) {
	o.Test = &v
}

func (o EmailDefaultContentAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Preview != nil {
		toSerialize["preview"] = o.Preview
	}
	if o.Test != nil {
		toSerialize["test"] = o.Test
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailDefaultContentAllOfLinks) UnmarshalJSON(bytes []byte) (err error) {
	varEmailDefaultContentAllOfLinks := _EmailDefaultContentAllOfLinks{}

	err = json.Unmarshal(bytes, &varEmailDefaultContentAllOfLinks)
	if err == nil {
		*o = EmailDefaultContentAllOfLinks(varEmailDefaultContentAllOfLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "template")
		delete(additionalProperties, "preview")
		delete(additionalProperties, "test")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailDefaultContentAllOfLinks struct {
	value *EmailDefaultContentAllOfLinks
	isSet bool
}

func (v NullableEmailDefaultContentAllOfLinks) Get() *EmailDefaultContentAllOfLinks {
	return v.value
}

func (v *NullableEmailDefaultContentAllOfLinks) Set(val *EmailDefaultContentAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDefaultContentAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDefaultContentAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDefaultContentAllOfLinks(val *EmailDefaultContentAllOfLinks) *NullableEmailDefaultContentAllOfLinks {
	return &NullableEmailDefaultContentAllOfLinks{value: val, isSet: true}
}

func (v NullableEmailDefaultContentAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDefaultContentAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

