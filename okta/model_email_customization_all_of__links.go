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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the EmailCustomizationAllOfLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailCustomizationAllOfLinks{}

// EmailCustomizationAllOfLinks struct for EmailCustomizationAllOfLinks
type EmailCustomizationAllOfLinks struct {
	Self                 *HrefObject `json:"self,omitempty"`
	Template             *HrefObject `json:"template,omitempty"`
	Preview              *HrefObject `json:"preview,omitempty"`
	Test                 *HrefObject `json:"test,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailCustomizationAllOfLinks EmailCustomizationAllOfLinks

// NewEmailCustomizationAllOfLinks instantiates a new EmailCustomizationAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailCustomizationAllOfLinks() *EmailCustomizationAllOfLinks {
	this := EmailCustomizationAllOfLinks{}
	return &this
}

// NewEmailCustomizationAllOfLinksWithDefaults instantiates a new EmailCustomizationAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailCustomizationAllOfLinksWithDefaults() *EmailCustomizationAllOfLinks {
	this := EmailCustomizationAllOfLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EmailCustomizationAllOfLinks) GetSelf() HrefObject {
	if o == nil || IsNil(o.Self) {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomizationAllOfLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EmailCustomizationAllOfLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *EmailCustomizationAllOfLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EmailCustomizationAllOfLinks) GetTemplate() HrefObject {
	if o == nil || IsNil(o.Template) {
		var ret HrefObject
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomizationAllOfLinks) GetTemplateOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EmailCustomizationAllOfLinks) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given HrefObject and assigns it to the Template field.
func (o *EmailCustomizationAllOfLinks) SetTemplate(v HrefObject) {
	o.Template = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *EmailCustomizationAllOfLinks) GetPreview() HrefObject {
	if o == nil || IsNil(o.Preview) {
		var ret HrefObject
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomizationAllOfLinks) GetPreviewOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Preview) {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *EmailCustomizationAllOfLinks) HasPreview() bool {
	if o != nil && !IsNil(o.Preview) {
		return true
	}

	return false
}

// SetPreview gets a reference to the given HrefObject and assigns it to the Preview field.
func (o *EmailCustomizationAllOfLinks) SetPreview(v HrefObject) {
	o.Preview = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *EmailCustomizationAllOfLinks) GetTest() HrefObject {
	if o == nil || IsNil(o.Test) {
		var ret HrefObject
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailCustomizationAllOfLinks) GetTestOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *EmailCustomizationAllOfLinks) HasTest() bool {
	if o != nil && !IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given HrefObject and assigns it to the Test field.
func (o *EmailCustomizationAllOfLinks) SetTest(v HrefObject) {
	o.Test = &v
}

func (o EmailCustomizationAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailCustomizationAllOfLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Preview) {
		toSerialize["preview"] = o.Preview
	}
	if !IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailCustomizationAllOfLinks) UnmarshalJSON(data []byte) (err error) {
	varEmailCustomizationAllOfLinks := _EmailCustomizationAllOfLinks{}

	err = json.Unmarshal(data, &varEmailCustomizationAllOfLinks)

	if err != nil {
		return err
	}

	*o = EmailCustomizationAllOfLinks(varEmailCustomizationAllOfLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "template")
		delete(additionalProperties, "preview")
		delete(additionalProperties, "test")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailCustomizationAllOfLinks struct {
	value *EmailCustomizationAllOfLinks
	isSet bool
}

func (v NullableEmailCustomizationAllOfLinks) Get() *EmailCustomizationAllOfLinks {
	return v.value
}

func (v *NullableEmailCustomizationAllOfLinks) Set(val *EmailCustomizationAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailCustomizationAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailCustomizationAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailCustomizationAllOfLinks(val *EmailCustomizationAllOfLinks) *NullableEmailCustomizationAllOfLinks {
	return &NullableEmailCustomizationAllOfLinks{value: val, isSet: true}
}

func (v NullableEmailCustomizationAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailCustomizationAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
