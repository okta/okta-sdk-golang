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

// checks if the PatchAIAgentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchAIAgentRequest{}

// PatchAIAgentRequest JSON Merge Patch for AI agent. Send only the fields to update. Use null to remove a value.
type PatchAIAgentRequest struct {
	// The ID of the connected app for the AI Agent
	AppId                NullableString       `json:"appId,omitempty"`
	Profile              *PatchAIAgentProfile `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchAIAgentRequest PatchAIAgentRequest

// NewPatchAIAgentRequest instantiates a new PatchAIAgentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchAIAgentRequest() *PatchAIAgentRequest {
	this := PatchAIAgentRequest{}
	return &this
}

// NewPatchAIAgentRequestWithDefaults instantiates a new PatchAIAgentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchAIAgentRequestWithDefaults() *PatchAIAgentRequest {
	this := PatchAIAgentRequest{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchAIAgentRequest) GetAppId() string {
	if o == nil || IsNil(o.AppId.Get()) {
		var ret string
		return ret
	}
	return *o.AppId.Get()
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchAIAgentRequest) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppId.Get(), o.AppId.IsSet()
}

// HasAppId returns a boolean if a field has been set.
func (o *PatchAIAgentRequest) HasAppId() bool {
	if o != nil && o.AppId.IsSet() {
		return true
	}

	return false
}

// SetAppId gets a reference to the given NullableString and assigns it to the AppId field.
func (o *PatchAIAgentRequest) SetAppId(v string) {
	o.AppId.Set(&v)
}

// SetAppIdNil sets the value for AppId to be an explicit nil
func (o *PatchAIAgentRequest) SetAppIdNil() {
	o.AppId.Set(nil)
}

// UnsetAppId ensures that no value is present for AppId, not even an explicit nil
func (o *PatchAIAgentRequest) UnsetAppId() {
	o.AppId.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PatchAIAgentRequest) GetProfile() PatchAIAgentProfile {
	if o == nil || IsNil(o.Profile) {
		var ret PatchAIAgentProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAIAgentRequest) GetProfileOk() (*PatchAIAgentProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PatchAIAgentRequest) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given PatchAIAgentProfile and assigns it to the Profile field.
func (o *PatchAIAgentRequest) SetProfile(v PatchAIAgentProfile) {
	o.Profile = &v
}

func (o PatchAIAgentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchAIAgentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppId.IsSet() {
		toSerialize["appId"] = o.AppId.Get()
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchAIAgentRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchAIAgentRequest := _PatchAIAgentRequest{}

	err = json.Unmarshal(data, &varPatchAIAgentRequest)

	if err != nil {
		return err
	}

	*o = PatchAIAgentRequest(varPatchAIAgentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appId")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchAIAgentRequest struct {
	value *PatchAIAgentRequest
	isSet bool
}

func (v NullablePatchAIAgentRequest) Get() *PatchAIAgentRequest {
	return v.value
}

func (v *NullablePatchAIAgentRequest) Set(val *PatchAIAgentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchAIAgentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchAIAgentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchAIAgentRequest(val *PatchAIAgentRequest) *NullablePatchAIAgentRequest {
	return &NullablePatchAIAgentRequest{value: val, isSet: true}
}

func (v NullablePatchAIAgentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchAIAgentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
