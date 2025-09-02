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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// RegistrationResponseCommandsInner struct for RegistrationResponseCommandsInner
type RegistrationResponseCommandsInner struct {
	// The location where you specify the command. To set attributes in the user's Okta profile, supply a `type` property set to `com.okta.user.profile.update`, together with a `value` property set to a list of key-value pairs corresponding to the Okta user profile attributes you want to set. The attributes must already exist in your user profile schema.  To explicitly allow or deny registration to the user, supply a type property set to `com.okta.action.update`, together with a value property set to `{\"registration\": \"ALLOW\"}` or `{\"registration\": \"DENY\"}`. The default is to allow registration.  In Okta Identity Engine, to set attributes in the user's profile, supply a `type` property set to `com.okta.user.progressive.profile.update`, together with a `value` property set to a list of key-value pairs corresponding to the Progressive Enrollment attributes that you want to set. See [Registration inline hook - Send response](https://developer.okta.com/docs/guides/registration-inline-hook/nodejs/main/#send-response).  Commands are applied in the order that they appear in the array. Within a single `com.okta.user.profile.update` or `com.okta.user.progressive.profile.update command`, attributes are updated in the order that they appear in the `value` object.  You can never use a command to update the user's password, but you are allowed to set the values of attributes other than password that are designated sensitive in your Okta user schema. However, the values of those sensitive attributes, if included as fields in the Profile Enrollment form, aren't included in the `data.userProfile` object sent to your external service by Okta. See [data.userProfile](/openapi/okta-management/management/tag/InlineHook/#tag/InlineHook/operation/create-registration-hook!path=0/data/userProfile&t=request).
	Type *string `json:"type,omitempty"`
	// The `value` object is the parameter to pass to the command.  For `com.okta.user.profile.update` commands, `value` should be an object containing one or more name-value pairs for the attributes you wish to update.  For `com.okta.action.update` commands, the value should be an object containing the attribute `action` set to a value of either `ALLOW` or `DENY`, indicating whether the registration should be permitted or not.  Registrations are allowed by default, so setting a value of `ALLOW` for the action field is valid but superfluous.
	Value map[string]interface{} `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationResponseCommandsInner RegistrationResponseCommandsInner

// NewRegistrationResponseCommandsInner instantiates a new RegistrationResponseCommandsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationResponseCommandsInner() *RegistrationResponseCommandsInner {
	this := RegistrationResponseCommandsInner{}
	return &this
}

// NewRegistrationResponseCommandsInnerWithDefaults instantiates a new RegistrationResponseCommandsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationResponseCommandsInnerWithDefaults() *RegistrationResponseCommandsInner {
	this := RegistrationResponseCommandsInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegistrationResponseCommandsInner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseCommandsInner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegistrationResponseCommandsInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RegistrationResponseCommandsInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RegistrationResponseCommandsInner) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseCommandsInner) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RegistrationResponseCommandsInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *RegistrationResponseCommandsInner) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o RegistrationResponseCommandsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RegistrationResponseCommandsInner) UnmarshalJSON(bytes []byte) (err error) {
	varRegistrationResponseCommandsInner := _RegistrationResponseCommandsInner{}

	err = json.Unmarshal(bytes, &varRegistrationResponseCommandsInner)
	if err == nil {
		*o = RegistrationResponseCommandsInner(varRegistrationResponseCommandsInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRegistrationResponseCommandsInner struct {
	value *RegistrationResponseCommandsInner
	isSet bool
}

func (v NullableRegistrationResponseCommandsInner) Get() *RegistrationResponseCommandsInner {
	return v.value
}

func (v *NullableRegistrationResponseCommandsInner) Set(val *RegistrationResponseCommandsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationResponseCommandsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationResponseCommandsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationResponseCommandsInner(val *RegistrationResponseCommandsInner) *NullableRegistrationResponseCommandsInner {
	return &NullableRegistrationResponseCommandsInner{value: val, isSet: true}
}

func (v NullableRegistrationResponseCommandsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationResponseCommandsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

