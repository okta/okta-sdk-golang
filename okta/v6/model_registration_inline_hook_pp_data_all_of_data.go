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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the RegistrationInlineHookPPDataAllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInlineHookPPDataAllOfData{}

// RegistrationInlineHookPPDataAllOfData struct for RegistrationInlineHookPPDataAllOfData
type RegistrationInlineHookPPDataAllOfData struct {
	Context *RegistrationInlineHookPPDataAllOfDataContext `json:"context,omitempty"`
	// The default action the system takes. Set to `ALLOW`. `DENY` is never sent to your external service
	Action *string `json:"action,omitempty"`
	// Name-value pairs for each new attribute supplied by the user in the Progressive Profile form
	UserProfileUpdate    map[string]interface{} `json:"userProfileUpdate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationInlineHookPPDataAllOfData RegistrationInlineHookPPDataAllOfData

// NewRegistrationInlineHookPPDataAllOfData instantiates a new RegistrationInlineHookPPDataAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInlineHookPPDataAllOfData() *RegistrationInlineHookPPDataAllOfData {
	this := RegistrationInlineHookPPDataAllOfData{}
	return &this
}

// NewRegistrationInlineHookPPDataAllOfDataWithDefaults instantiates a new RegistrationInlineHookPPDataAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInlineHookPPDataAllOfDataWithDefaults() *RegistrationInlineHookPPDataAllOfData {
	this := RegistrationInlineHookPPDataAllOfData{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPDataAllOfData) GetContext() RegistrationInlineHookPPDataAllOfDataContext {
	if o == nil || IsNil(o.Context) {
		var ret RegistrationInlineHookPPDataAllOfDataContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPDataAllOfData) GetContextOk() (*RegistrationInlineHookPPDataAllOfDataContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPDataAllOfData) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given RegistrationInlineHookPPDataAllOfDataContext and assigns it to the Context field.
func (o *RegistrationInlineHookPPDataAllOfData) SetContext(v RegistrationInlineHookPPDataAllOfDataContext) {
	o.Context = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPDataAllOfData) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPDataAllOfData) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPDataAllOfData) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *RegistrationInlineHookPPDataAllOfData) SetAction(v string) {
	o.Action = &v
}

// GetUserProfileUpdate returns the UserProfileUpdate field value if set, zero value otherwise.
func (o *RegistrationInlineHookPPDataAllOfData) GetUserProfileUpdate() map[string]interface{} {
	if o == nil || IsNil(o.UserProfileUpdate) {
		var ret map[string]interface{}
		return ret
	}
	return o.UserProfileUpdate
}

// GetUserProfileUpdateOk returns a tuple with the UserProfileUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInlineHookPPDataAllOfData) GetUserProfileUpdateOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UserProfileUpdate) {
		return map[string]interface{}{}, false
	}
	return o.UserProfileUpdate, true
}

// HasUserProfileUpdate returns a boolean if a field has been set.
func (o *RegistrationInlineHookPPDataAllOfData) HasUserProfileUpdate() bool {
	if o != nil && !IsNil(o.UserProfileUpdate) {
		return true
	}

	return false
}

// SetUserProfileUpdate gets a reference to the given map[string]interface{} and assigns it to the UserProfileUpdate field.
func (o *RegistrationInlineHookPPDataAllOfData) SetUserProfileUpdate(v map[string]interface{}) {
	o.UserProfileUpdate = v
}

func (o RegistrationInlineHookPPDataAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInlineHookPPDataAllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.UserProfileUpdate) {
		toSerialize["userProfileUpdate"] = o.UserProfileUpdate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationInlineHookPPDataAllOfData) UnmarshalJSON(data []byte) (err error) {
	varRegistrationInlineHookPPDataAllOfData := _RegistrationInlineHookPPDataAllOfData{}

	err = json.Unmarshal(data, &varRegistrationInlineHookPPDataAllOfData)

	if err != nil {
		return err
	}

	*o = RegistrationInlineHookPPDataAllOfData(varRegistrationInlineHookPPDataAllOfData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "action")
		delete(additionalProperties, "userProfileUpdate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationInlineHookPPDataAllOfData struct {
	value *RegistrationInlineHookPPDataAllOfData
	isSet bool
}

func (v NullableRegistrationInlineHookPPDataAllOfData) Get() *RegistrationInlineHookPPDataAllOfData {
	return v.value
}

func (v *NullableRegistrationInlineHookPPDataAllOfData) Set(val *RegistrationInlineHookPPDataAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInlineHookPPDataAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInlineHookPPDataAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInlineHookPPDataAllOfData(val *RegistrationInlineHookPPDataAllOfData) *NullableRegistrationInlineHookPPDataAllOfData {
	return &NullableRegistrationInlineHookPPDataAllOfData{value: val, isSet: true}
}

func (v NullableRegistrationInlineHookPPDataAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInlineHookPPDataAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
