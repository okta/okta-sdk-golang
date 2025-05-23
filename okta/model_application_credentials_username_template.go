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

// ApplicationCredentialsUsernameTemplate struct for ApplicationCredentialsUsernameTemplate
type ApplicationCredentialsUsernameTemplate struct {
	PushStatus *string `json:"pushStatus,omitempty"`
	Template *string `json:"template,omitempty"`
	Type *string `json:"type,omitempty"`
	UserSuffix *string `json:"userSuffix,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationCredentialsUsernameTemplate ApplicationCredentialsUsernameTemplate

// NewApplicationCredentialsUsernameTemplate instantiates a new ApplicationCredentialsUsernameTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCredentialsUsernameTemplate() *ApplicationCredentialsUsernameTemplate {
	this := ApplicationCredentialsUsernameTemplate{}
	return &this
}

// NewApplicationCredentialsUsernameTemplateWithDefaults instantiates a new ApplicationCredentialsUsernameTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCredentialsUsernameTemplateWithDefaults() *ApplicationCredentialsUsernameTemplate {
	this := ApplicationCredentialsUsernameTemplate{}
	return &this
}

// GetPushStatus returns the PushStatus field value if set, zero value otherwise.
func (o *ApplicationCredentialsUsernameTemplate) GetPushStatus() string {
	if o == nil || o.PushStatus == nil {
		var ret string
		return ret
	}
	return *o.PushStatus
}

// GetPushStatusOk returns a tuple with the PushStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsUsernameTemplate) GetPushStatusOk() (*string, bool) {
	if o == nil || o.PushStatus == nil {
		return nil, false
	}
	return o.PushStatus, true
}

// HasPushStatus returns a boolean if a field has been set.
func (o *ApplicationCredentialsUsernameTemplate) HasPushStatus() bool {
	if o != nil && o.PushStatus != nil {
		return true
	}

	return false
}

// SetPushStatus gets a reference to the given string and assigns it to the PushStatus field.
func (o *ApplicationCredentialsUsernameTemplate) SetPushStatus(v string) {
	o.PushStatus = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ApplicationCredentialsUsernameTemplate) GetTemplate() string {
	if o == nil || o.Template == nil {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsUsernameTemplate) GetTemplateOk() (*string, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ApplicationCredentialsUsernameTemplate) HasTemplate() bool {
	if o != nil && o.Template != nil {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *ApplicationCredentialsUsernameTemplate) SetTemplate(v string) {
	o.Template = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApplicationCredentialsUsernameTemplate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsUsernameTemplate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationCredentialsUsernameTemplate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApplicationCredentialsUsernameTemplate) SetType(v string) {
	o.Type = &v
}

// GetUserSuffix returns the UserSuffix field value if set, zero value otherwise.
func (o *ApplicationCredentialsUsernameTemplate) GetUserSuffix() string {
	if o == nil || o.UserSuffix == nil {
		var ret string
		return ret
	}
	return *o.UserSuffix
}

// GetUserSuffixOk returns a tuple with the UserSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsUsernameTemplate) GetUserSuffixOk() (*string, bool) {
	if o == nil || o.UserSuffix == nil {
		return nil, false
	}
	return o.UserSuffix, true
}

// HasUserSuffix returns a boolean if a field has been set.
func (o *ApplicationCredentialsUsernameTemplate) HasUserSuffix() bool {
	if o != nil && o.UserSuffix != nil {
		return true
	}

	return false
}

// SetUserSuffix gets a reference to the given string and assigns it to the UserSuffix field.
func (o *ApplicationCredentialsUsernameTemplate) SetUserSuffix(v string) {
	o.UserSuffix = &v
}

func (o ApplicationCredentialsUsernameTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PushStatus != nil {
		toSerialize["pushStatus"] = o.PushStatus
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UserSuffix != nil {
		toSerialize["userSuffix"] = o.UserSuffix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationCredentialsUsernameTemplate) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationCredentialsUsernameTemplate := _ApplicationCredentialsUsernameTemplate{}

	err = json.Unmarshal(bytes, &varApplicationCredentialsUsernameTemplate)
	if err == nil {
		*o = ApplicationCredentialsUsernameTemplate(varApplicationCredentialsUsernameTemplate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "pushStatus")
		delete(additionalProperties, "template")
		delete(additionalProperties, "type")
		delete(additionalProperties, "userSuffix")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationCredentialsUsernameTemplate struct {
	value *ApplicationCredentialsUsernameTemplate
	isSet bool
}

func (v NullableApplicationCredentialsUsernameTemplate) Get() *ApplicationCredentialsUsernameTemplate {
	return v.value
}

func (v *NullableApplicationCredentialsUsernameTemplate) Set(val *ApplicationCredentialsUsernameTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCredentialsUsernameTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCredentialsUsernameTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCredentialsUsernameTemplate(val *ApplicationCredentialsUsernameTemplate) *NullableApplicationCredentialsUsernameTemplate {
	return &NullableApplicationCredentialsUsernameTemplate{value: val, isSet: true}
}

func (v NullableApplicationCredentialsUsernameTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCredentialsUsernameTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

