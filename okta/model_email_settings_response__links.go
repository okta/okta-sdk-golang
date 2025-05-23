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

// EmailSettingsResponseLinks struct for EmailSettingsResponseLinks
type EmailSettingsResponseLinks struct {
	Self *HrefObject `json:"self,omitempty"`
	Template *HrefObject `json:"template,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailSettingsResponseLinks EmailSettingsResponseLinks

// NewEmailSettingsResponseLinks instantiates a new EmailSettingsResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSettingsResponseLinks() *EmailSettingsResponseLinks {
	this := EmailSettingsResponseLinks{}
	return &this
}

// NewEmailSettingsResponseLinksWithDefaults instantiates a new EmailSettingsResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSettingsResponseLinksWithDefaults() *EmailSettingsResponseLinks {
	this := EmailSettingsResponseLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EmailSettingsResponseLinks) GetSelf() HrefObject {
	if o == nil || o.Self == nil {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsResponseLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EmailSettingsResponseLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *EmailSettingsResponseLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EmailSettingsResponseLinks) GetTemplate() HrefObject {
	if o == nil || o.Template == nil {
		var ret HrefObject
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSettingsResponseLinks) GetTemplateOk() (*HrefObject, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EmailSettingsResponseLinks) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given HrefObject and assigns it to the Template field.
func (o *EmailSettingsResponseLinks) SetTemplate(v HrefObject) {
	o.Template = &v
}

func (o EmailSettingsResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailSettingsResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	varEmailSettingsResponseLinks := _EmailSettingsResponseLinks{}

	err = json.Unmarshal(bytes, &varEmailSettingsResponseLinks)
	if err == nil {
		*o = EmailSettingsResponseLinks(varEmailSettingsResponseLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "template")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailSettingsResponseLinks struct {
	value *EmailSettingsResponseLinks
	isSet bool
}

func (v NullableEmailSettingsResponseLinks) Get() *EmailSettingsResponseLinks {
	return v.value
}

func (v *NullableEmailSettingsResponseLinks) Set(val *EmailSettingsResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSettingsResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSettingsResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSettingsResponseLinks(val *EmailSettingsResponseLinks) *NullableEmailSettingsResponseLinks {
	return &NullableEmailSettingsResponseLinks{value: val, isSet: true}
}

func (v NullableEmailSettingsResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSettingsResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

