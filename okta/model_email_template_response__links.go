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

// checks if the EmailTemplateResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailTemplateResponseLinks{}

// EmailTemplateResponseLinks struct for EmailTemplateResponseLinks
type EmailTemplateResponseLinks struct {
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	Settings             *HrefObject         `json:"settings,omitempty"`
	DefaultContent       *HrefObject         `json:"defaultContent,omitempty"`
	Customizations       *HrefObject         `json:"customizations,omitempty"`
	Test                 *HrefObject         `json:"test,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailTemplateResponseLinks EmailTemplateResponseLinks

// NewEmailTemplateResponseLinks instantiates a new EmailTemplateResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailTemplateResponseLinks() *EmailTemplateResponseLinks {
	this := EmailTemplateResponseLinks{}
	return &this
}

// NewEmailTemplateResponseLinksWithDefaults instantiates a new EmailTemplateResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailTemplateResponseLinksWithDefaults() *EmailTemplateResponseLinks {
	this := EmailTemplateResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EmailTemplateResponseLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponseLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EmailTemplateResponseLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *EmailTemplateResponseLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *EmailTemplateResponseLinks) GetSettings() HrefObject {
	if o == nil || IsNil(o.Settings) {
		var ret HrefObject
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponseLinks) GetSettingsOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *EmailTemplateResponseLinks) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given HrefObject and assigns it to the Settings field.
func (o *EmailTemplateResponseLinks) SetSettings(v HrefObject) {
	o.Settings = &v
}

// GetDefaultContent returns the DefaultContent field value if set, zero value otherwise.
func (o *EmailTemplateResponseLinks) GetDefaultContent() HrefObject {
	if o == nil || IsNil(o.DefaultContent) {
		var ret HrefObject
		return ret
	}
	return *o.DefaultContent
}

// GetDefaultContentOk returns a tuple with the DefaultContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponseLinks) GetDefaultContentOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.DefaultContent) {
		return nil, false
	}
	return o.DefaultContent, true
}

// HasDefaultContent returns a boolean if a field has been set.
func (o *EmailTemplateResponseLinks) HasDefaultContent() bool {
	if o != nil && !IsNil(o.DefaultContent) {
		return true
	}

	return false
}

// SetDefaultContent gets a reference to the given HrefObject and assigns it to the DefaultContent field.
func (o *EmailTemplateResponseLinks) SetDefaultContent(v HrefObject) {
	o.DefaultContent = &v
}

// GetCustomizations returns the Customizations field value if set, zero value otherwise.
func (o *EmailTemplateResponseLinks) GetCustomizations() HrefObject {
	if o == nil || IsNil(o.Customizations) {
		var ret HrefObject
		return ret
	}
	return *o.Customizations
}

// GetCustomizationsOk returns a tuple with the Customizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponseLinks) GetCustomizationsOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Customizations) {
		return nil, false
	}
	return o.Customizations, true
}

// HasCustomizations returns a boolean if a field has been set.
func (o *EmailTemplateResponseLinks) HasCustomizations() bool {
	if o != nil && !IsNil(o.Customizations) {
		return true
	}

	return false
}

// SetCustomizations gets a reference to the given HrefObject and assigns it to the Customizations field.
func (o *EmailTemplateResponseLinks) SetCustomizations(v HrefObject) {
	o.Customizations = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *EmailTemplateResponseLinks) GetTest() HrefObject {
	if o == nil || IsNil(o.Test) {
		var ret HrefObject
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailTemplateResponseLinks) GetTestOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *EmailTemplateResponseLinks) HasTest() bool {
	if o != nil && !IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given HrefObject and assigns it to the Test field.
func (o *EmailTemplateResponseLinks) SetTest(v HrefObject) {
	o.Test = &v
}

func (o EmailTemplateResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailTemplateResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.DefaultContent) {
		toSerialize["defaultContent"] = o.DefaultContent
	}
	if !IsNil(o.Customizations) {
		toSerialize["customizations"] = o.Customizations
	}
	if !IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailTemplateResponseLinks) UnmarshalJSON(data []byte) (err error) {
	varEmailTemplateResponseLinks := _EmailTemplateResponseLinks{}

	err = json.Unmarshal(data, &varEmailTemplateResponseLinks)

	if err != nil {
		return err
	}

	*o = EmailTemplateResponseLinks(varEmailTemplateResponseLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "defaultContent")
		delete(additionalProperties, "customizations")
		delete(additionalProperties, "test")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailTemplateResponseLinks struct {
	value *EmailTemplateResponseLinks
	isSet bool
}

func (v NullableEmailTemplateResponseLinks) Get() *EmailTemplateResponseLinks {
	return v.value
}

func (v *NullableEmailTemplateResponseLinks) Set(val *EmailTemplateResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailTemplateResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailTemplateResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailTemplateResponseLinks(val *EmailTemplateResponseLinks) *NullableEmailTemplateResponseLinks {
	return &NullableEmailTemplateResponseLinks{value: val, isSet: true}
}

func (v NullableEmailTemplateResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailTemplateResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
