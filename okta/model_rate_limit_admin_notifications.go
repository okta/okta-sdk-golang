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

// RateLimitAdminNotifications 
type RateLimitAdminNotifications struct {
	NotificationsEnabled bool `json:"notificationsEnabled"`
	AdditionalProperties map[string]interface{}
}

type _RateLimitAdminNotifications RateLimitAdminNotifications

// NewRateLimitAdminNotifications instantiates a new RateLimitAdminNotifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitAdminNotifications(notificationsEnabled bool) *RateLimitAdminNotifications {
	this := RateLimitAdminNotifications{}
	this.NotificationsEnabled = notificationsEnabled
	return &this
}

// NewRateLimitAdminNotificationsWithDefaults instantiates a new RateLimitAdminNotifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitAdminNotificationsWithDefaults() *RateLimitAdminNotifications {
	this := RateLimitAdminNotifications{}
	return &this
}

// GetNotificationsEnabled returns the NotificationsEnabled field value
func (o *RateLimitAdminNotifications) GetNotificationsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NotificationsEnabled
}

// GetNotificationsEnabledOk returns a tuple with the NotificationsEnabled field value
// and a boolean to check if the value has been set.
func (o *RateLimitAdminNotifications) GetNotificationsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationsEnabled, true
}

// SetNotificationsEnabled sets field value
func (o *RateLimitAdminNotifications) SetNotificationsEnabled(v bool) {
	o.NotificationsEnabled = v
}

func (o RateLimitAdminNotifications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["notificationsEnabled"] = o.NotificationsEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RateLimitAdminNotifications) UnmarshalJSON(bytes []byte) (err error) {
	varRateLimitAdminNotifications := _RateLimitAdminNotifications{}

	err = json.Unmarshal(bytes, &varRateLimitAdminNotifications)
	if err == nil {
		*o = RateLimitAdminNotifications(varRateLimitAdminNotifications)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "notificationsEnabled")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRateLimitAdminNotifications struct {
	value *RateLimitAdminNotifications
	isSet bool
}

func (v NullableRateLimitAdminNotifications) Get() *RateLimitAdminNotifications {
	return v.value
}

func (v *NullableRateLimitAdminNotifications) Set(val *RateLimitAdminNotifications) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitAdminNotifications) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitAdminNotifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitAdminNotifications(val *RateLimitAdminNotifications) *NullableRateLimitAdminNotifications {
	return &NullableRateLimitAdminNotifications{value: val, isSet: true}
}

func (v NullableRateLimitAdminNotifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitAdminNotifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

