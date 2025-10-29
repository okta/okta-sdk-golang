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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the RegistrationInlineHookSSRDataAllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInlineHookSSRDataAllOfData{}

// RegistrationInlineHookSSRDataAllOfData struct for RegistrationInlineHookSSRDataAllOfData
type RegistrationInlineHookSSRDataAllOfData struct {
	Context *RegistrationInlineHookSSRDataAllOfDataContext `json:"context,omitempty"`
	// The default action the system will take. Will be `ALLOW`. `DENY` will never be sent to your external service.
	Action *string `json:"action,omitempty"`
	// The name-value pairs for each registration-related attribute supplied by the user in the Profile Enrollment form.
	UserProfile          map[string]interface{} `json:"userProfile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookSSRDataAllOfData RegistrationInlineHookSSRDataAllOfData

// NewRegistrationInlineHookSSRDataAllOfData instantiates a new RegistrationInlineHookSSRDataAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookSSRDataAllOfData() *RegistrationInlineHookSSRDataAllOfData {
	this := RegistrationInlineHookSSRDataAllOfData{}
	return &this
}

// NewRegistrationInlineHookSSRDataAllOfDataWithDefaults instantiates a new RegistrationInlineHookSSRDataAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookSSRDataAllOfDataWithDefaults() *RegistrationInlineHookSSRDataAllOfData {
	this := RegistrationInlineHookSSRDataAllOfData{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRDataAllOfData) GetContext() RegistrationInlineHookSSRDataAllOfDataContext {
	if o == nil || IsNil(o.Context) {
		var ret RegistrationInlineHookSSRDataAllOfDataContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRDataAllOfData) GetContextOk() (*RegistrationInlineHookSSRDataAllOfDataContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRDataAllOfData) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given RegistrationInlineHookSSRDataAllOfDataContext and assigns it to the Context field.
func (o *RegistrationInlineHookSSRDataAllOfData) SetContext(v RegistrationInlineHookSSRDataAllOfDataContext) {
	o.Context = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRDataAllOfData) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRDataAllOfData) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRDataAllOfData) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RegistrationInlineHookSSRDataAllOfData) SetAction(v string) {
	o.Action = &v
}

// GetUserProfile returns the UserProfile field value if set, zero value otherwise.
func (o *RegistrationInlineHookSSRDataAllOfData) GetUserProfile() map[string]interface{} {
	if o == nil || IsNil(o.UserProfile) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserProfile
}

// GetUserProfileOk returns a tuple with the UserProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookSSRDataAllOfData) GetUserProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserProfile) {
		return map[string]interface{}{}, false
	}
	return o.UserProfile, true
}

// HasUserProfile returns a boolean if a field has been set.
func (o *RegistrationInlineHookSSRDataAllOfData) HasUserProfile() bool {
	if o != nil && !IsNil(o.UserProfile) {
		return true
	}

	return false
}

// SetUserProfile gets a reference to the given map[string]interface{} and assigns it to the UserProfile field.
func (o *RegistrationInlineHookSSRDataAllOfData) SetUserProfile(v map[string]interface{}) {
	o.UserProfile = v
}

func (o RegistrationInlineHookSSRDataAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInlineHookSSRDataAllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.UserProfile) {
		toSerialize["userProfile"] = o.UserProfile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationInlineHookSSRDataAllOfData) UnmarshalJSON(data []byte) (err error) {
	varRegistrationInlineHookSSRDataAllOfData := _RegistrationInlineHookSSRDataAllOfData{}

	err = json.Unmarshal(data, &varRegistrationInlineHookSSRDataAllOfData)

	if err != nil {
		return err
	}

	*o = RegistrationInlineHookSSRDataAllOfData(varRegistrationInlineHookSSRDataAllOfData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "action")
		delete(additionalProperties, "userProfile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationInlineHookSSRDataAllOfData struct {
	value *RegistrationInlineHookSSRDataAllOfData
	isSet bool
}

func (v NullableRegistrationInlineHookSSRDataAllOfData) Get() *RegistrationInlineHookSSRDataAllOfData {
	return v.value
}

func (v *NullableRegistrationInlineHookSSRDataAllOfData) Set(val *RegistrationInlineHookSSRDataAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookSSRDataAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookSSRDataAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookSSRDataAllOfData(val *RegistrationInlineHookSSRDataAllOfData) *NullableRegistrationInlineHookSSRDataAllOfData {
	return &NullableRegistrationInlineHookSSRDataAllOfData{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookSSRDataAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookSSRDataAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
