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

// ApplicationCredentials Credentials for the specified `signOnMode`
type ApplicationCredentials struct {
	Signing *ApplicationCredentialsSigning `json:"signing,omitempty"`
	UserNameTemplate *ApplicationCredentialsUsernameTemplate `json:"userNameTemplate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationCredentials ApplicationCredentials

// NewApplicationCredentials instantiates a new ApplicationCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCredentials() *ApplicationCredentials {
	this := ApplicationCredentials{}
	return &this
}

// NewApplicationCredentialsWithDefaults instantiates a new ApplicationCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCredentialsWithDefaults() *ApplicationCredentials {
	this := ApplicationCredentials{}
	return &this
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *ApplicationCredentials) GetSigning() ApplicationCredentialsSigning {
	if o == nil || o.Signing == nil {
		var ret ApplicationCredentialsSigning
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentials) GetSigningOk() (*ApplicationCredentialsSigning, bool) {
	if o == nil || o.Signing == nil {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *ApplicationCredentials) HasSigning() bool {
	if o != nil && o.Signing != nil {
		return true
	}

	return false
}

// SetSigning gets a reference to the given ApplicationCredentialsSigning and assigns it to the Signing field.
func (o *ApplicationCredentials) SetSigning(v ApplicationCredentialsSigning) {
	o.Signing = &v
}

// GetUserNameTemplate returns the UserNameTemplate field value if set, zero value otherwise.
func (o *ApplicationCredentials) GetUserNameTemplate() ApplicationCredentialsUsernameTemplate {
	if o == nil || o.UserNameTemplate == nil {
		var ret ApplicationCredentialsUsernameTemplate
		return ret
	}
	return *o.UserNameTemplate
}

// GetUserNameTemplateOk returns a tuple with the UserNameTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentials) GetUserNameTemplateOk() (*ApplicationCredentialsUsernameTemplate, bool) {
	if o == nil || o.UserNameTemplate == nil {
		return nil, false
	}
	return o.UserNameTemplate, true
}

// HasUserNameTemplate returns a boolean if a field has been set.
func (o *ApplicationCredentials) HasUserNameTemplate() bool {
	if o != nil && o.UserNameTemplate != nil {
		return true
	}

	return false
}

// SetUserNameTemplate gets a reference to the given ApplicationCredentialsUsernameTemplate and assigns it to the UserNameTemplate field.
func (o *ApplicationCredentials) SetUserNameTemplate(v ApplicationCredentialsUsernameTemplate) {
	o.UserNameTemplate = &v
}

func (o ApplicationCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Signing != nil {
		toSerialize["signing"] = o.Signing
	}
	if o.UserNameTemplate != nil {
		toSerialize["userNameTemplate"] = o.UserNameTemplate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationCredentials := _ApplicationCredentials{}

	err = json.Unmarshal(bytes, &varApplicationCredentials)
	if err == nil {
		*o = ApplicationCredentials(varApplicationCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "signing")
		delete(additionalProperties, "userNameTemplate")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationCredentials struct {
	value *ApplicationCredentials
	isSet bool
}

func (v NullableApplicationCredentials) Get() *ApplicationCredentials {
	return v.value
}

func (v *NullableApplicationCredentials) Set(val *ApplicationCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCredentials(val *ApplicationCredentials) *NullableApplicationCredentials {
	return &NullableApplicationCredentials{value: val, isSet: true}
}

func (v NullableApplicationCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

