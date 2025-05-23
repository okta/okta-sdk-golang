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

// AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate struct for AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate
type AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate struct {
	// The Duo Security user template name
	Template *string `json:"template,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate

// NewAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate instantiates a new AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate() *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate {
	this := AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate{}
	return &this
}

// NewAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplateWithDefaults instantiates a new AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplateWithDefaults() *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate {
	this := AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) SetTemplate(v string) {
	o.Template = &v
}

func (o AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate := _AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate)
	if err == nil {
		*o = AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate(varAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "template")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate struct {
	value *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate
	isSet bool
}

func (v NullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) Get() *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate {
	return v.value
}

func (v *NullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) Set(val *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate(val *AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) *NullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate {
	return &NullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

