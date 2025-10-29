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

// checks if the ApplicationCredentialsUsernameTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCredentialsUsernameTemplate{}

// ApplicationCredentialsUsernameTemplate The template used to generate the username when the app is assigned through a group or directly to a user
type ApplicationCredentialsUsernameTemplate struct {
	// Determines if the username is pushed to the app on updates for CUSTOM `type`
	PushStatus *string `json:"pushStatus,omitempty"`
	// Mapping expression used to generate usernames.  The following are supported mapping expressions that are used with the `BUILT_IN` template type:  | Name                            | Template Expression                            | | ------------------------------- | ---------------------------------------------- | | AD Employee ID                  | `${source.employeeID}`                         | | AD SAM Account Name             | `${source.samAccountName}`                     | | AD SAM Account Name (lowercase) | `${fn:toLowerCase(source.samAccountName)}`     | | AD User Principal Name          | `${source.userName}`                           | | AD User Principal Name prefix   | `${fn:substringBefore(source.userName, \"@\")}`  | | Email                           | `${source.email}`                              | | Email (lowercase)               | `${fn:toLowerCase(source.email)}`              | | Email prefix                    | `${fn:substringBefore(source.email, \"@\")}`     | | LDAP UID + custom suffix        | `${source.userName}${instance.userSuffix}`     | | Okta username                   | `${source.login}`                              | | Okta username prefix            | `${fn:substringBefore(source.login, \"@\")}`     |
	Template *string `json:"template,omitempty"`
	// Type of mapping expression. Empty string is allowed.
	Type *string `json:"type,omitempty"`
	// An optional suffix appended to usernames for `BUILT_IN` mapping expressions
	UserSuffix           *string `json:"userSuffix,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationCredentialsUsernameTemplate ApplicationCredentialsUsernameTemplate

// NewApplicationCredentialsUsernameTemplate instantiates a new ApplicationCredentialsUsernameTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCredentialsUsernameTemplate() *ApplicationCredentialsUsernameTemplate {
	this := ApplicationCredentialsUsernameTemplate{}
	var template string = "${source.login}"
	this.Template = &template
	var type_ string = "BUILT_IN"
	this.Type = &type_
	return &this
}

// NewApplicationCredentialsUsernameTemplateWithDefaults instantiates a new ApplicationCredentialsUsernameTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCredentialsUsernameTemplateWithDefaults() *ApplicationCredentialsUsernameTemplate {
	this := ApplicationCredentialsUsernameTemplate{}
	var template string = "${source.login}"
	this.Template = &template
	var type_ string = "BUILT_IN"
	this.Type = &type_
	return &this
}

// GetPushStatus returns the PushStatus field value if set, zero value otherwise.
func (o *ApplicationCredentialsUsernameTemplate) GetPushStatus() string {
	if o == nil || IsNil(o.PushStatus) {
		var ret string
		return ret
	}
	return *o.PushStatus
}

// GetPushStatusOk returns a tuple with the PushStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsUsernameTemplate) GetPushStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PushStatus) {
		return nil, false
	}
	return o.PushStatus, true
}

// HasPushStatus returns a boolean if a field has been set.
func (o *ApplicationCredentialsUsernameTemplate) HasPushStatus() bool {
	if o != nil && !IsNil(o.PushStatus) {
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
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsUsernameTemplate) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ApplicationCredentialsUsernameTemplate) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsUsernameTemplate) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationCredentialsUsernameTemplate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
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
	if o == nil || IsNil(o.UserSuffix) {
		var ret string
		return ret
	}
	return *o.UserSuffix
}

// GetUserSuffixOk returns a tuple with the UserSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialsUsernameTemplate) GetUserSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.UserSuffix) {
		return nil, false
	}
	return o.UserSuffix, true
}

// HasUserSuffix returns a boolean if a field has been set.
func (o *ApplicationCredentialsUsernameTemplate) HasUserSuffix() bool {
	if o != nil && !IsNil(o.UserSuffix) {
		return true
	}

	return false
}

// SetUserSuffix gets a reference to the given string and assigns it to the UserSuffix field.
func (o *ApplicationCredentialsUsernameTemplate) SetUserSuffix(v string) {
	o.UserSuffix = &v
}

func (o ApplicationCredentialsUsernameTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCredentialsUsernameTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PushStatus) {
		toSerialize["pushStatus"] = o.PushStatus
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserSuffix) {
		toSerialize["userSuffix"] = o.UserSuffix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationCredentialsUsernameTemplate) UnmarshalJSON(data []byte) (err error) {
	varApplicationCredentialsUsernameTemplate := _ApplicationCredentialsUsernameTemplate{}

	err = json.Unmarshal(data, &varApplicationCredentialsUsernameTemplate)

	if err != nil {
		return err
	}

	*o = ApplicationCredentialsUsernameTemplate(varApplicationCredentialsUsernameTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pushStatus")
		delete(additionalProperties, "template")
		delete(additionalProperties, "type")
		delete(additionalProperties, "userSuffix")
		o.AdditionalProperties = additionalProperties
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
