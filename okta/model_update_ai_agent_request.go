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

// checks if the UpdateAIAgentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAIAgentRequest{}

// UpdateAIAgentRequest struct for UpdateAIAgentRequest
type UpdateAIAgentRequest struct {
	// The ID of the connected app for the AI agent
	AppId                *string         `json:"appId,omitempty"`
	Profile              *AIAgentProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateAIAgentRequest UpdateAIAgentRequest

// NewUpdateAIAgentRequest instantiates a new UpdateAIAgentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAIAgentRequest() *UpdateAIAgentRequest {
	this := UpdateAIAgentRequest{}
	return &this
}

// NewUpdateAIAgentRequestWithDefaults instantiates a new UpdateAIAgentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAIAgentRequestWithDefaults() *UpdateAIAgentRequest {
	this := UpdateAIAgentRequest{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *UpdateAIAgentRequest) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIAgentRequest) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *UpdateAIAgentRequest) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *UpdateAIAgentRequest) SetAppId(v string) {
	o.AppId = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UpdateAIAgentRequest) GetProfile() AIAgentProfile {
	if o == nil || IsNil(o.Profile) {
		var ret AIAgentProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAIAgentRequest) GetProfileOk() (*AIAgentProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UpdateAIAgentRequest) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given AIAgentProfile and assigns it to the Profile field.
func (o *UpdateAIAgentRequest) SetProfile(v AIAgentProfile) {
	o.Profile = &v
}

func (o UpdateAIAgentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAIAgentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateAIAgentRequest) UnmarshalJSON(data []byte) (err error) {
	varUpdateAIAgentRequest := _UpdateAIAgentRequest{}

	err = json.Unmarshal(data, &varUpdateAIAgentRequest)

	if err != nil {
		return err
	}

	*o = UpdateAIAgentRequest(varUpdateAIAgentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appId")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateAIAgentRequest struct {
	value *UpdateAIAgentRequest
	isSet bool
}

func (v NullableUpdateAIAgentRequest) Get() *UpdateAIAgentRequest {
	return v.value
}

func (v *NullableUpdateAIAgentRequest) Set(val *UpdateAIAgentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAIAgentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAIAgentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAIAgentRequest(val *UpdateAIAgentRequest) *NullableUpdateAIAgentRequest {
	return &NullableUpdateAIAgentRequest{value: val, isSet: true}
}

func (v NullableUpdateAIAgentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAIAgentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
