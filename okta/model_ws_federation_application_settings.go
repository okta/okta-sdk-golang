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

// WsFederationApplicationSettings struct for WsFederationApplicationSettings
type WsFederationApplicationSettings struct {
	IdentityStoreId *string `json:"identityStoreId,omitempty"`
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty"`
	InlineHookId *string `json:"inlineHookId,omitempty"`
	Notes *ApplicationSettingsNotes `json:"notes,omitempty"`
	Notifications *ApplicationSettingsNotifications `json:"notifications,omitempty"`
	App *WsFederationApplicationSettingsApplication `json:"app,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WsFederationApplicationSettings WsFederationApplicationSettings

// NewWsFederationApplicationSettings instantiates a new WsFederationApplicationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWsFederationApplicationSettings() *WsFederationApplicationSettings {
	this := WsFederationApplicationSettings{}
	return &this
}

// NewWsFederationApplicationSettingsWithDefaults instantiates a new WsFederationApplicationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWsFederationApplicationSettingsWithDefaults() *WsFederationApplicationSettings {
	this := WsFederationApplicationSettings{}
	return &this
}

// GetIdentityStoreId returns the IdentityStoreId field value if set, zero value otherwise.
func (o *WsFederationApplicationSettings) GetIdentityStoreId() string {
	if o == nil || o.IdentityStoreId == nil {
		var ret string
		return ret
	}
	return *o.IdentityStoreId
}

// GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettings) GetIdentityStoreIdOk() (*string, bool) {
	if o == nil || o.IdentityStoreId == nil {
		return nil, false
	}
	return o.IdentityStoreId, true
}

// HasIdentityStoreId returns a boolean if a field has been set.
func (o *WsFederationApplicationSettings) HasIdentityStoreId() bool {
	if o != nil && o.IdentityStoreId != nil {
		return true
	}

	return false
}

// SetIdentityStoreId gets a reference to the given string and assigns it to the IdentityStoreId field.
func (o *WsFederationApplicationSettings) SetIdentityStoreId(v string) {
	o.IdentityStoreId = &v
}

// GetImplicitAssignment returns the ImplicitAssignment field value if set, zero value otherwise.
func (o *WsFederationApplicationSettings) GetImplicitAssignment() bool {
	if o == nil || o.ImplicitAssignment == nil {
		var ret bool
		return ret
	}
	return *o.ImplicitAssignment
}

// GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettings) GetImplicitAssignmentOk() (*bool, bool) {
	if o == nil || o.ImplicitAssignment == nil {
		return nil, false
	}
	return o.ImplicitAssignment, true
}

// HasImplicitAssignment returns a boolean if a field has been set.
func (o *WsFederationApplicationSettings) HasImplicitAssignment() bool {
	if o != nil && o.ImplicitAssignment != nil {
		return true
	}

	return false
}

// SetImplicitAssignment gets a reference to the given bool and assigns it to the ImplicitAssignment field.
func (o *WsFederationApplicationSettings) SetImplicitAssignment(v bool) {
	o.ImplicitAssignment = &v
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *WsFederationApplicationSettings) GetInlineHookId() string {
	if o == nil || o.InlineHookId == nil {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettings) GetInlineHookIdOk() (*string, bool) {
	if o == nil || o.InlineHookId == nil {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *WsFederationApplicationSettings) HasInlineHookId() bool {
	if o != nil && o.InlineHookId != nil {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *WsFederationApplicationSettings) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *WsFederationApplicationSettings) GetNotes() ApplicationSettingsNotes {
	if o == nil || o.Notes == nil {
		var ret ApplicationSettingsNotes
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *WsFederationApplicationSettings) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given ApplicationSettingsNotes and assigns it to the Notes field.
func (o *WsFederationApplicationSettings) SetNotes(v ApplicationSettingsNotes) {
	o.Notes = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *WsFederationApplicationSettings) GetNotifications() ApplicationSettingsNotifications {
	if o == nil || o.Notifications == nil {
		var ret ApplicationSettingsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *WsFederationApplicationSettings) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given ApplicationSettingsNotifications and assigns it to the Notifications field.
func (o *WsFederationApplicationSettings) SetNotifications(v ApplicationSettingsNotifications) {
	o.Notifications = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *WsFederationApplicationSettings) GetApp() WsFederationApplicationSettingsApplication {
	if o == nil || o.App == nil {
		var ret WsFederationApplicationSettingsApplication
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsFederationApplicationSettings) GetAppOk() (*WsFederationApplicationSettingsApplication, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *WsFederationApplicationSettings) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given WsFederationApplicationSettingsApplication and assigns it to the App field.
func (o *WsFederationApplicationSettings) SetApp(v WsFederationApplicationSettingsApplication) {
	o.App = &v
}

func (o WsFederationApplicationSettings) MarshalJSON() ([]byte, error) {
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
	if o.App != nil {
		toSerialize["app"] = o.App
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WsFederationApplicationSettings) UnmarshalJSON(bytes []byte) (err error) {
	varWsFederationApplicationSettings := _WsFederationApplicationSettings{}

	err = json.Unmarshal(bytes, &varWsFederationApplicationSettings)
	if err == nil {
		*o = WsFederationApplicationSettings(varWsFederationApplicationSettings)
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
		delete(additionalProperties, "app")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWsFederationApplicationSettings struct {
	value *WsFederationApplicationSettings
	isSet bool
}

func (v NullableWsFederationApplicationSettings) Get() *WsFederationApplicationSettings {
	return v.value
}

func (v *NullableWsFederationApplicationSettings) Set(val *WsFederationApplicationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableWsFederationApplicationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableWsFederationApplicationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWsFederationApplicationSettings(val *WsFederationApplicationSettings) *NullableWsFederationApplicationSettings {
	return &NullableWsFederationApplicationSettings{value: val, isSet: true}
}

func (v NullableWsFederationApplicationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWsFederationApplicationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

