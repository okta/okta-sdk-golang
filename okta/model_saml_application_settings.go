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

// SamlApplicationSettings struct for SamlApplicationSettings
type SamlApplicationSettings struct {
	IdentityStoreId *string `json:"identityStoreId,omitempty"`
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty"`
	InlineHookId *string `json:"inlineHookId,omitempty"`
	Notes *ApplicationSettingsNotes `json:"notes,omitempty"`
	Notifications *ApplicationSettingsNotifications `json:"notifications,omitempty"`
	App *SamlApplicationSettingsApplication `json:"app,omitempty"`
	SignOn *SamlApplicationSettingsSignOn `json:"signOn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlApplicationSettings SamlApplicationSettings

// NewSamlApplicationSettings instantiates a new SamlApplicationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlApplicationSettings() *SamlApplicationSettings {
	this := SamlApplicationSettings{}
	return &this
}

// NewSamlApplicationSettingsWithDefaults instantiates a new SamlApplicationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlApplicationSettingsWithDefaults() *SamlApplicationSettings {
	this := SamlApplicationSettings{}
	return &this
}

// GetIdentityStoreId returns the IdentityStoreId field value if set, zero value otherwise.
func (o *SamlApplicationSettings) GetIdentityStoreId() string {
	if o == nil || o.IdentityStoreId == nil {
		var ret string
		return ret
	}
	return *o.IdentityStoreId
}

// GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettings) GetIdentityStoreIdOk() (*string, bool) {
	if o == nil || o.IdentityStoreId == nil {
		return nil, false
	}
	return o.IdentityStoreId, true
}

// HasIdentityStoreId returns a boolean if a field has been set.
func (o *SamlApplicationSettings) HasIdentityStoreId() bool {
	if o != nil && o.IdentityStoreId != nil {
		return true
	}

	return false
}

// SetIdentityStoreId gets a reference to the given string and assigns it to the IdentityStoreId field.
func (o *SamlApplicationSettings) SetIdentityStoreId(v string) {
	o.IdentityStoreId = &v
}

// GetImplicitAssignment returns the ImplicitAssignment field value if set, zero value otherwise.
func (o *SamlApplicationSettings) GetImplicitAssignment() bool {
	if o == nil || o.ImplicitAssignment == nil {
		var ret bool
		return ret
	}
	return *o.ImplicitAssignment
}

// GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettings) GetImplicitAssignmentOk() (*bool, bool) {
	if o == nil || o.ImplicitAssignment == nil {
		return nil, false
	}
	return o.ImplicitAssignment, true
}

// HasImplicitAssignment returns a boolean if a field has been set.
func (o *SamlApplicationSettings) HasImplicitAssignment() bool {
	if o != nil && o.ImplicitAssignment != nil {
		return true
	}

	return false
}

// SetImplicitAssignment gets a reference to the given bool and assigns it to the ImplicitAssignment field.
func (o *SamlApplicationSettings) SetImplicitAssignment(v bool) {
	o.ImplicitAssignment = &v
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *SamlApplicationSettings) GetInlineHookId() string {
	if o == nil || o.InlineHookId == nil {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettings) GetInlineHookIdOk() (*string, bool) {
	if o == nil || o.InlineHookId == nil {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *SamlApplicationSettings) HasInlineHookId() bool {
	if o != nil && o.InlineHookId != nil {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *SamlApplicationSettings) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SamlApplicationSettings) GetNotes() ApplicationSettingsNotes {
	if o == nil || o.Notes == nil {
		var ret ApplicationSettingsNotes
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SamlApplicationSettings) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given ApplicationSettingsNotes and assigns it to the Notes field.
func (o *SamlApplicationSettings) SetNotes(v ApplicationSettingsNotes) {
	o.Notes = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *SamlApplicationSettings) GetNotifications() ApplicationSettingsNotifications {
	if o == nil || o.Notifications == nil {
		var ret ApplicationSettingsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *SamlApplicationSettings) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given ApplicationSettingsNotifications and assigns it to the Notifications field.
func (o *SamlApplicationSettings) SetNotifications(v ApplicationSettingsNotifications) {
	o.Notifications = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *SamlApplicationSettings) GetApp() SamlApplicationSettingsApplication {
	if o == nil || o.App == nil {
		var ret SamlApplicationSettingsApplication
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettings) GetAppOk() (*SamlApplicationSettingsApplication, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *SamlApplicationSettings) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given SamlApplicationSettingsApplication and assigns it to the App field.
func (o *SamlApplicationSettings) SetApp(v SamlApplicationSettingsApplication) {
	o.App = &v
}

// GetSignOn returns the SignOn field value if set, zero value otherwise.
func (o *SamlApplicationSettings) GetSignOn() SamlApplicationSettingsSignOn {
	if o == nil || o.SignOn == nil {
		var ret SamlApplicationSettingsSignOn
		return ret
	}
	return *o.SignOn
}

// GetSignOnOk returns a tuple with the SignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApplicationSettings) GetSignOnOk() (*SamlApplicationSettingsSignOn, bool) {
	if o == nil || o.SignOn == nil {
		return nil, false
	}
	return o.SignOn, true
}

// HasSignOn returns a boolean if a field has been set.
func (o *SamlApplicationSettings) HasSignOn() bool {
	if o != nil && o.SignOn != nil {
		return true
	}

	return false
}

// SetSignOn gets a reference to the given SamlApplicationSettingsSignOn and assigns it to the SignOn field.
func (o *SamlApplicationSettings) SetSignOn(v SamlApplicationSettingsSignOn) {
	o.SignOn = &v
}

func (o SamlApplicationSettings) MarshalJSON() ([]byte, error) {
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
	if o.SignOn != nil {
		toSerialize["signOn"] = o.SignOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlApplicationSettings) UnmarshalJSON(bytes []byte) (err error) {
	varSamlApplicationSettings := _SamlApplicationSettings{}

	err = json.Unmarshal(bytes, &varSamlApplicationSettings)
	if err == nil {
		*o = SamlApplicationSettings(varSamlApplicationSettings)
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
		delete(additionalProperties, "signOn")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlApplicationSettings struct {
	value *SamlApplicationSettings
	isSet bool
}

func (v NullableSamlApplicationSettings) Get() *SamlApplicationSettings {
	return v.value
}

func (v *NullableSamlApplicationSettings) Set(val *SamlApplicationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlApplicationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlApplicationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlApplicationSettings(val *SamlApplicationSettings) *NullableSamlApplicationSettings {
	return &NullableSamlApplicationSettings{value: val, isSet: true}
}

func (v NullableSamlApplicationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlApplicationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

