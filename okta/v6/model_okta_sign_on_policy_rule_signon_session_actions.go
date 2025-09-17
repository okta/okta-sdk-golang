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

// checks if the OktaSignOnPolicyRuleSignonSessionActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaSignOnPolicyRuleSignonSessionActions{}

// OktaSignOnPolicyRuleSignonSessionActions Properties governing the user's session lifetime
type OktaSignOnPolicyRuleSignonSessionActions struct {
	// Maximum number of minutes that a user session can be idle before the session is ended
	MaxSessionIdleMinutes *int32 `json:"maxSessionIdleMinutes,omitempty"`
	// Maximum number of minutes (from when the user signs in) that a user's session is active. Set this to force users to sign in again after the number of specified minutes. Disable by setting to `0`.
	MaxSessionLifetimeMinutes *int32 `json:"maxSessionLifetimeMinutes,omitempty"`
	// If set to `false`, user session cookies only last the length of a browser session. If set to `true`, user session cookies last across browser sessions. This setting doesn't impact administrators who can never have persistent session cookies. This property is read-only for the default rule of the default global session policy.
	UsePersistentCookie  *bool `json:"usePersistentCookie,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSignOnPolicyRuleSignonSessionActions OktaSignOnPolicyRuleSignonSessionActions

// NewOktaSignOnPolicyRuleSignonSessionActions instantiates a new OktaSignOnPolicyRuleSignonSessionActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSignOnPolicyRuleSignonSessionActions() *OktaSignOnPolicyRuleSignonSessionActions {
	this := OktaSignOnPolicyRuleSignonSessionActions{}
	var maxSessionIdleMinutes int32 = 120
	this.MaxSessionIdleMinutes = &maxSessionIdleMinutes
	var maxSessionLifetimeMinutes int32 = 0
	this.MaxSessionLifetimeMinutes = &maxSessionLifetimeMinutes
	var usePersistentCookie bool = false
	this.UsePersistentCookie = &usePersistentCookie
	return &this
}

// NewOktaSignOnPolicyRuleSignonSessionActionsWithDefaults instantiates a new OktaSignOnPolicyRuleSignonSessionActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSignOnPolicyRuleSignonSessionActionsWithDefaults() *OktaSignOnPolicyRuleSignonSessionActions {
	this := OktaSignOnPolicyRuleSignonSessionActions{}
	var maxSessionIdleMinutes int32 = 120
	this.MaxSessionIdleMinutes = &maxSessionIdleMinutes
	var maxSessionLifetimeMinutes int32 = 0
	this.MaxSessionLifetimeMinutes = &maxSessionLifetimeMinutes
	var usePersistentCookie bool = false
	this.UsePersistentCookie = &usePersistentCookie
	return &this
}

// GetMaxSessionIdleMinutes returns the MaxSessionIdleMinutes field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleSignonSessionActions) GetMaxSessionIdleMinutes() int32 {
	if o == nil || IsNil(o.MaxSessionIdleMinutes) {
		var ret int32
		return ret
	}
	return *o.MaxSessionIdleMinutes
}

// GetMaxSessionIdleMinutesOk returns a tuple with the MaxSessionIdleMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleSignonSessionActions) GetMaxSessionIdleMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSessionIdleMinutes) {
		return nil, false
	}
	return o.MaxSessionIdleMinutes, true
}

// HasMaxSessionIdleMinutes returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleSignonSessionActions) HasMaxSessionIdleMinutes() bool {
	if o != nil && !IsNil(o.MaxSessionIdleMinutes) {
		return true
	}

	return false
}

// SetMaxSessionIdleMinutes gets a reference to the given int32 and assigns it to the MaxSessionIdleMinutes field.
func (o *OktaSignOnPolicyRuleSignonSessionActions) SetMaxSessionIdleMinutes(v int32) {
	o.MaxSessionIdleMinutes = &v
}

// GetMaxSessionLifetimeMinutes returns the MaxSessionLifetimeMinutes field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleSignonSessionActions) GetMaxSessionLifetimeMinutes() int32 {
	if o == nil || IsNil(o.MaxSessionLifetimeMinutes) {
		var ret int32
		return ret
	}
	return *o.MaxSessionLifetimeMinutes
}

// GetMaxSessionLifetimeMinutesOk returns a tuple with the MaxSessionLifetimeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleSignonSessionActions) GetMaxSessionLifetimeMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSessionLifetimeMinutes) {
		return nil, false
	}
	return o.MaxSessionLifetimeMinutes, true
}

// HasMaxSessionLifetimeMinutes returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleSignonSessionActions) HasMaxSessionLifetimeMinutes() bool {
	if o != nil && !IsNil(o.MaxSessionLifetimeMinutes) {
		return true
	}

	return false
}

// SetMaxSessionLifetimeMinutes gets a reference to the given int32 and assigns it to the MaxSessionLifetimeMinutes field.
func (o *OktaSignOnPolicyRuleSignonSessionActions) SetMaxSessionLifetimeMinutes(v int32) {
	o.MaxSessionLifetimeMinutes = &v
}

// GetUsePersistentCookie returns the UsePersistentCookie field value if set, zero value otherwise.
func (o *OktaSignOnPolicyRuleSignonSessionActions) GetUsePersistentCookie() bool {
	if o == nil || IsNil(o.UsePersistentCookie) {
		var ret bool
		return ret
	}
	return *o.UsePersistentCookie
}

// GetUsePersistentCookieOk returns a tuple with the UsePersistentCookie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSignOnPolicyRuleSignonSessionActions) GetUsePersistentCookieOk() (*bool, bool) {
	if o == nil || IsNil(o.UsePersistentCookie) {
		return nil, false
	}
	return o.UsePersistentCookie, true
}

// HasUsePersistentCookie returns a boolean if a field has been set.
func (o *OktaSignOnPolicyRuleSignonSessionActions) HasUsePersistentCookie() bool {
	if o != nil && !IsNil(o.UsePersistentCookie) {
		return true
	}

	return false
}

// SetUsePersistentCookie gets a reference to the given bool and assigns it to the UsePersistentCookie field.
func (o *OktaSignOnPolicyRuleSignonSessionActions) SetUsePersistentCookie(v bool) {
	o.UsePersistentCookie = &v
}

func (o OktaSignOnPolicyRuleSignonSessionActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaSignOnPolicyRuleSignonSessionActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxSessionIdleMinutes) {
		toSerialize["maxSessionIdleMinutes"] = o.MaxSessionIdleMinutes
	}
	if !IsNil(o.MaxSessionLifetimeMinutes) {
		toSerialize["maxSessionLifetimeMinutes"] = o.MaxSessionLifetimeMinutes
	}
	if !IsNil(o.UsePersistentCookie) {
		toSerialize["usePersistentCookie"] = o.UsePersistentCookie
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaSignOnPolicyRuleSignonSessionActions) UnmarshalJSON(data []byte) (err error) {
	varOktaSignOnPolicyRuleSignonSessionActions := _OktaSignOnPolicyRuleSignonSessionActions{}

	err = json.Unmarshal(data, &varOktaSignOnPolicyRuleSignonSessionActions)

	if err != nil {
		return err
	}

	*o = OktaSignOnPolicyRuleSignonSessionActions(varOktaSignOnPolicyRuleSignonSessionActions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "maxSessionIdleMinutes")
		delete(additionalProperties, "maxSessionLifetimeMinutes")
		delete(additionalProperties, "usePersistentCookie")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaSignOnPolicyRuleSignonSessionActions struct {
	value *OktaSignOnPolicyRuleSignonSessionActions
	isSet bool
}

func (v NullableOktaSignOnPolicyRuleSignonSessionActions) Get() *OktaSignOnPolicyRuleSignonSessionActions {
	return v.value
}

func (v *NullableOktaSignOnPolicyRuleSignonSessionActions) Set(val *OktaSignOnPolicyRuleSignonSessionActions) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSignOnPolicyRuleSignonSessionActions) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSignOnPolicyRuleSignonSessionActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSignOnPolicyRuleSignonSessionActions(val *OktaSignOnPolicyRuleSignonSessionActions) *NullableOktaSignOnPolicyRuleSignonSessionActions {
	return &NullableOktaSignOnPolicyRuleSignonSessionActions{value: val, isSet: true}
}

func (v NullableOktaSignOnPolicyRuleSignonSessionActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSignOnPolicyRuleSignonSessionActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
