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

// ApplicationSettings App settings
type ApplicationSettings struct {
	IdentityStoreId *string `json:"identityStoreId,omitempty"`
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty"`
	InlineHookId *string `json:"inlineHookId,omitempty"`
	Notes *ApplicationSettingsNotes `json:"notes,omitempty"`
	Notifications *ApplicationSettingsNotifications `json:"notifications,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationSettings ApplicationSettings

// NewApplicationSettings instantiates a new ApplicationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSettings() *ApplicationSettings {
	this := ApplicationSettings{}
	return &this
}

// NewApplicationSettingsWithDefaults instantiates a new ApplicationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSettingsWithDefaults() *ApplicationSettings {
	this := ApplicationSettings{}
	return &this
}

// GetIdentityStoreId returns the IdentityStoreId field value if set, zero value otherwise.
func (o *ApplicationSettings) GetIdentityStoreId() string {
	if o == nil || o.IdentityStoreId == nil {
		var ret string
		return ret
	}
	return *o.IdentityStoreId
}

// GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettings) GetIdentityStoreIdOk() (*string, bool) {
	if o == nil || o.IdentityStoreId == nil {
		return nil, false
	}
	return o.IdentityStoreId, true
}

// HasIdentityStoreId returns a boolean if a field has been set.
func (o *ApplicationSettings) HasIdentityStoreId() bool {
	if o != nil && o.IdentityStoreId != nil {
		return true
	}

	return false
}

// SetIdentityStoreId gets a reference to the given string and assigns it to the IdentityStoreId field.
func (o *ApplicationSettings) SetIdentityStoreId(v string) {
	o.IdentityStoreId = &v
}

// GetImplicitAssignment returns the ImplicitAssignment field value if set, zero value otherwise.
func (o *ApplicationSettings) GetImplicitAssignment() bool {
	if o == nil || o.ImplicitAssignment == nil {
		var ret bool
		return ret
	}
	return *o.ImplicitAssignment
}

// GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettings) GetImplicitAssignmentOk() (*bool, bool) {
	if o == nil || o.ImplicitAssignment == nil {
		return nil, false
	}
	return o.ImplicitAssignment, true
}

// HasImplicitAssignment returns a boolean if a field has been set.
func (o *ApplicationSettings) HasImplicitAssignment() bool {
	if o != nil && o.ImplicitAssignment != nil {
		return true
	}

	return false
}

// SetImplicitAssignment gets a reference to the given bool and assigns it to the ImplicitAssignment field.
func (o *ApplicationSettings) SetImplicitAssignment(v bool) {
	o.ImplicitAssignment = &v
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *ApplicationSettings) GetInlineHookId() string {
	if o == nil || o.InlineHookId == nil {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettings) GetInlineHookIdOk() (*string, bool) {
	if o == nil || o.InlineHookId == nil {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *ApplicationSettings) HasInlineHookId() bool {
	if o != nil && o.InlineHookId != nil {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *ApplicationSettings) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApplicationSettings) GetNotes() ApplicationSettingsNotes {
	if o == nil || o.Notes == nil {
		var ret ApplicationSettingsNotes
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApplicationSettings) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given ApplicationSettingsNotes and assigns it to the Notes field.
func (o *ApplicationSettings) SetNotes(v ApplicationSettingsNotes) {
	o.Notes = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *ApplicationSettings) GetNotifications() ApplicationSettingsNotifications {
	if o == nil || o.Notifications == nil {
		var ret ApplicationSettingsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *ApplicationSettings) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given ApplicationSettingsNotifications and assigns it to the Notifications field.
func (o *ApplicationSettings) SetNotifications(v ApplicationSettingsNotifications) {
	o.Notifications = &v
}

func (o ApplicationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IdentityStoreId != nil {
		toSerialize["identityStoreId"] = o.IdentityStoreId
	}
	if o.ImplicitAssignment != nil {
		toSerialize["implicitAssignment"] = o.ImplicitAssignment
	}
	if o.InlineHookId != nil {
		toSerialize["inlineHookId"] = o.InlineHookId
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationSettings) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationSettings := _ApplicationSettings{}

	err = json.Unmarshal(bytes, &varApplicationSettings)
	if err == nil {
		*o = ApplicationSettings(varApplicationSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "identityStoreId")
		delete(additionalProperties, "implicitAssignment")
		delete(additionalProperties, "inlineHookId")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "notifications")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationSettings struct {
	value *ApplicationSettings
	isSet bool
}

func (v NullableApplicationSettings) Get() *ApplicationSettings {
	return v.value
}

func (v *NullableApplicationSettings) Set(val *ApplicationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSettings(val *ApplicationSettings) *NullableApplicationSettings {
	return &NullableApplicationSettings{value: val, isSet: true}
}

func (v NullableApplicationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

