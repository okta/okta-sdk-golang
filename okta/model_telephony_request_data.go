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

// checks if the TelephonyRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TelephonyRequestData{}

// TelephonyRequestData struct for TelephonyRequestData
type TelephonyRequestData struct {
	Context              *RegistrationInlineHookSSRDataAllOfDataContext `json:"context,omitempty"`
	MessageProfile       *TelephonyRequestDataMessageProfile            `json:"messageProfile,omitempty"`
	UserProfile          *TelephonyRequestDataUserProfile               `json:"userProfile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelephonyRequestData TelephonyRequestData

// NewTelephonyRequestData instantiates a new TelephonyRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephonyRequestData() *TelephonyRequestData {
	this := TelephonyRequestData{}
	return &this
}

// NewTelephonyRequestDataWithDefaults instantiates a new TelephonyRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephonyRequestDataWithDefaults() *TelephonyRequestData {
	this := TelephonyRequestData{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelephonyRequestData) GetContext() RegistrationInlineHookSSRDataAllOfDataContext {
	if o == nil || IsNil(o.Context) {
		var ret RegistrationInlineHookSSRDataAllOfDataContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestData) GetContextOk() (*RegistrationInlineHookSSRDataAllOfDataContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelephonyRequestData) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given RegistrationInlineHookSSRDataAllOfDataContext and assigns it to the Context field.
func (o *TelephonyRequestData) SetContext(v RegistrationInlineHookSSRDataAllOfDataContext) {
	o.Context = &v
}

// GetMessageProfile returns the MessageProfile field value if set, zero value otherwise.
func (o *TelephonyRequestData) GetMessageProfile() TelephonyRequestDataMessageProfile {
	if o == nil || IsNil(o.MessageProfile) {
		var ret TelephonyRequestDataMessageProfile
		return ret
	}
	return *o.MessageProfile
}

// GetMessageProfileOk returns a tuple with the MessageProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestData) GetMessageProfileOk() (*TelephonyRequestDataMessageProfile, bool) {
	if o == nil || IsNil(o.MessageProfile) {
		return nil, false
	}
	return o.MessageProfile, true
}

// HasMessageProfile returns a boolean if a field has been set.
func (o *TelephonyRequestData) HasMessageProfile() bool {
	if o != nil && !IsNil(o.MessageProfile) {
		return true
	}

	return false
}

// SetMessageProfile gets a reference to the given TelephonyRequestDataMessageProfile and assigns it to the MessageProfile field.
func (o *TelephonyRequestData) SetMessageProfile(v TelephonyRequestDataMessageProfile) {
	o.MessageProfile = &v
}

// GetUserProfile returns the UserProfile field value if set, zero value otherwise.
func (o *TelephonyRequestData) GetUserProfile() TelephonyRequestDataUserProfile {
	if o == nil || IsNil(o.UserProfile) {
		var ret TelephonyRequestDataUserProfile
		return ret
	}
	return *o.UserProfile
}

// GetUserProfileOk returns a tuple with the UserProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestData) GetUserProfileOk() (*TelephonyRequestDataUserProfile, bool) {
	if o == nil || IsNil(o.UserProfile) {
		return nil, false
	}
	return o.UserProfile, true
}

// HasUserProfile returns a boolean if a field has been set.
func (o *TelephonyRequestData) HasUserProfile() bool {
	if o != nil && !IsNil(o.UserProfile) {
		return true
	}

	return false
}

// SetUserProfile gets a reference to the given TelephonyRequestDataUserProfile and assigns it to the UserProfile field.
func (o *TelephonyRequestData) SetUserProfile(v TelephonyRequestDataUserProfile) {
	o.UserProfile = &v
}

func (o TelephonyRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TelephonyRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.MessageProfile) {
		toSerialize["messageProfile"] = o.MessageProfile
	}
	if !IsNil(o.UserProfile) {
		toSerialize["userProfile"] = o.UserProfile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TelephonyRequestData) UnmarshalJSON(data []byte) (err error) {
	varTelephonyRequestData := _TelephonyRequestData{}

	err = json.Unmarshal(data, &varTelephonyRequestData)

	if err != nil {
		return err
	}

	*o = TelephonyRequestData(varTelephonyRequestData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "messageProfile")
		delete(additionalProperties, "userProfile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelephonyRequestData struct {
	value *TelephonyRequestData
	isSet bool
}

func (v NullableTelephonyRequestData) Get() *TelephonyRequestData {
	return v.value
}

func (v *NullableTelephonyRequestData) Set(val *TelephonyRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephonyRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephonyRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephonyRequestData(val *TelephonyRequestData) *NullableTelephonyRequestData {
	return &NullableTelephonyRequestData{value: val, isSet: true}
}

func (v NullableTelephonyRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephonyRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
