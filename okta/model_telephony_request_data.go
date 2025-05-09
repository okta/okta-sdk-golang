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

// TelephonyRequestData struct for TelephonyRequestData
type TelephonyRequestData struct {
	Context *TelephonyRequestDataContext `json:"context,omitempty"`
	MessageProfile *TelephonyRequestDataMessageProfile `json:"messageProfile,omitempty"`
	UserProfile *TelephonyRequestDataUserProfile `json:"userProfile,omitempty"`
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
func (o *TelephonyRequestData) GetContext() TelephonyRequestDataContext {
	if o == nil || o.Context == nil {
		var ret TelephonyRequestDataContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestData) GetContextOk() (*TelephonyRequestDataContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelephonyRequestData) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelephonyRequestDataContext and assigns it to the Context field.
func (o *TelephonyRequestData) SetContext(v TelephonyRequestDataContext) {
	o.Context = &v
}

// GetMessageProfile returns the MessageProfile field value if set, zero value otherwise.
func (o *TelephonyRequestData) GetMessageProfile() TelephonyRequestDataMessageProfile {
	if o == nil || o.MessageProfile == nil {
		var ret TelephonyRequestDataMessageProfile
		return ret
	}
	return *o.MessageProfile
}

// GetMessageProfileOk returns a tuple with the MessageProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestData) GetMessageProfileOk() (*TelephonyRequestDataMessageProfile, bool) {
	if o == nil || o.MessageProfile == nil {
		return nil, false
	}
	return o.MessageProfile, true
}

// HasMessageProfile returns a boolean if a field has been set.
func (o *TelephonyRequestData) HasMessageProfile() bool {
	if o != nil && o.MessageProfile != nil {
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
	if o == nil || o.UserProfile == nil {
		var ret TelephonyRequestDataUserProfile
		return ret
	}
	return *o.UserProfile
}

// GetUserProfileOk returns a tuple with the UserProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestData) GetUserProfileOk() (*TelephonyRequestDataUserProfile, bool) {
	if o == nil || o.UserProfile == nil {
		return nil, false
	}
	return o.UserProfile, true
}

// HasUserProfile returns a boolean if a field has been set.
func (o *TelephonyRequestData) HasUserProfile() bool {
	if o != nil && o.UserProfile != nil {
		return true
	}

	return false
}

// SetUserProfile gets a reference to the given TelephonyRequestDataUserProfile and assigns it to the UserProfile field.
func (o *TelephonyRequestData) SetUserProfile(v TelephonyRequestDataUserProfile) {
	o.UserProfile = &v
}

func (o TelephonyRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.MessageProfile != nil {
		toSerialize["messageProfile"] = o.MessageProfile
	}
	if o.UserProfile != nil {
		toSerialize["userProfile"] = o.UserProfile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelephonyRequestData) UnmarshalJSON(bytes []byte) (err error) {
	varTelephonyRequestData := _TelephonyRequestData{}

	err = json.Unmarshal(bytes, &varTelephonyRequestData)
	if err == nil {
		*o = TelephonyRequestData(varTelephonyRequestData)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "context")
		delete(additionalProperties, "messageProfile")
		delete(additionalProperties, "userProfile")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

