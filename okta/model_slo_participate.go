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

// checks if the SloParticipate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloParticipate{}

// SloParticipate Determines if the app participates in Single Logout (SLO)
type SloParticipate struct {
	// Request binding type
	BindingType *string `json:"bindingType,omitempty"`
	// Indicates whether the app is allowed to participate in front-channel SLO
	Enabled *bool `json:"enabled,omitempty"`
	// URL where Okta sends the logout request
	LogoutRequestUrl *string `json:"logoutRequestUrl,omitempty"`
	// Determines whether Okta sends the `SessionIndex` elements in the logout request
	SessionIndexRequired *bool `json:"sessionIndexRequired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SloParticipate SloParticipate

// NewSloParticipate instantiates a new SloParticipate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloParticipate() *SloParticipate {
	this := SloParticipate{}
	return &this
}

// NewSloParticipateWithDefaults instantiates a new SloParticipate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloParticipateWithDefaults() *SloParticipate {
	this := SloParticipate{}
	return &this
}

// GetBindingType returns the BindingType field value if set, zero value otherwise.
func (o *SloParticipate) GetBindingType() string {
	if o == nil || IsNil(o.BindingType) {
		var ret string
		return ret
	}
	return *o.BindingType
}

// GetBindingTypeOk returns a tuple with the BindingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloParticipate) GetBindingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BindingType) {
		return nil, false
	}
	return o.BindingType, true
}

// HasBindingType returns a boolean if a field has been set.
func (o *SloParticipate) HasBindingType() bool {
	if o != nil && !IsNil(o.BindingType) {
		return true
	}

	return false
}

// SetBindingType gets a reference to the given string and assigns it to the BindingType field.
func (o *SloParticipate) SetBindingType(v string) {
	o.BindingType = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SloParticipate) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloParticipate) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SloParticipate) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SloParticipate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLogoutRequestUrl returns the LogoutRequestUrl field value if set, zero value otherwise.
func (o *SloParticipate) GetLogoutRequestUrl() string {
	if o == nil || IsNil(o.LogoutRequestUrl) {
		var ret string
		return ret
	}
	return *o.LogoutRequestUrl
}

// GetLogoutRequestUrlOk returns a tuple with the LogoutRequestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloParticipate) GetLogoutRequestUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoutRequestUrl) {
		return nil, false
	}
	return o.LogoutRequestUrl, true
}

// HasLogoutRequestUrl returns a boolean if a field has been set.
func (o *SloParticipate) HasLogoutRequestUrl() bool {
	if o != nil && !IsNil(o.LogoutRequestUrl) {
		return true
	}

	return false
}

// SetLogoutRequestUrl gets a reference to the given string and assigns it to the LogoutRequestUrl field.
func (o *SloParticipate) SetLogoutRequestUrl(v string) {
	o.LogoutRequestUrl = &v
}

// GetSessionIndexRequired returns the SessionIndexRequired field value if set, zero value otherwise.
func (o *SloParticipate) GetSessionIndexRequired() bool {
	if o == nil || IsNil(o.SessionIndexRequired) {
		var ret bool
		return ret
	}
	return *o.SessionIndexRequired
}

// GetSessionIndexRequiredOk returns a tuple with the SessionIndexRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloParticipate) GetSessionIndexRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.SessionIndexRequired) {
		return nil, false
	}
	return o.SessionIndexRequired, true
}

// HasSessionIndexRequired returns a boolean if a field has been set.
func (o *SloParticipate) HasSessionIndexRequired() bool {
	if o != nil && !IsNil(o.SessionIndexRequired) {
		return true
	}

	return false
}

// SetSessionIndexRequired gets a reference to the given bool and assigns it to the SessionIndexRequired field.
func (o *SloParticipate) SetSessionIndexRequired(v bool) {
	o.SessionIndexRequired = &v
}

func (o SloParticipate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloParticipate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BindingType) {
		toSerialize["bindingType"] = o.BindingType
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.LogoutRequestUrl) {
		toSerialize["logoutRequestUrl"] = o.LogoutRequestUrl
	}
	if !IsNil(o.SessionIndexRequired) {
		toSerialize["sessionIndexRequired"] = o.SessionIndexRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SloParticipate) UnmarshalJSON(data []byte) (err error) {
	varSloParticipate := _SloParticipate{}

	err = json.Unmarshal(data, &varSloParticipate)

	if err != nil {
		return err
	}

	*o = SloParticipate(varSloParticipate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bindingType")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "logoutRequestUrl")
		delete(additionalProperties, "sessionIndexRequired")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSloParticipate struct {
	value *SloParticipate
	isSet bool
}

func (v NullableSloParticipate) Get() *SloParticipate {
	return v.value
}

func (v *NullableSloParticipate) Set(val *SloParticipate) {
	v.value = val
	v.isSet = true
}

func (v NullableSloParticipate) IsSet() bool {
	return v.isSet
}

func (v *NullableSloParticipate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloParticipate(val *SloParticipate) *NullableSloParticipate {
	return &NullableSloParticipate{value: val, isSet: true}
}

func (v NullableSloParticipate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloParticipate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
