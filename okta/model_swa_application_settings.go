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

// checks if the SwaApplicationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwaApplicationSettings{}

// SwaApplicationSettings struct for SwaApplicationSettings
type SwaApplicationSettings struct {
	// The Governance Engine opt-in status for the app
	EmOptInStatus *string `json:"emOptInStatus,omitempty"`
	// Identifies an additional identity store app, if your app supports it. The `identityStoreId` value must be a valid identity store app ID. This identity store app must be created in the same org as your app.
	IdentityStoreId *string `json:"identityStoreId,omitempty"`
	// Controls whether Okta automatically assigns users to the app based on the user's role or group membership.
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty"`
	// Identifier of an inline hook. Inline hooks are outbound calls from Okta to your own custom code, triggered at specific points in Okta process flows. They allow you to integrate custom functionality into those flows. See [Inline hooks](/openapi/okta-management/management/tag/InlineHook/).
	InlineHookId         *string                            `json:"inlineHookId,omitempty"`
	Notes                *ApplicationSettingsNotes          `json:"notes,omitempty"`
	Notifications        *ApplicationSettingsNotifications  `json:"notifications,omitempty"`
	App                  *SwaApplicationSettingsApplication `json:"app,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwaApplicationSettings SwaApplicationSettings

// NewSwaApplicationSettings instantiates a new SwaApplicationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwaApplicationSettings() *SwaApplicationSettings {
	this := SwaApplicationSettings{}
	return &this
}

// NewSwaApplicationSettingsWithDefaults instantiates a new SwaApplicationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwaApplicationSettingsWithDefaults() *SwaApplicationSettings {
	this := SwaApplicationSettings{}
	return &this
}

// GetEmOptInStatus returns the EmOptInStatus field value if set, zero value otherwise.
func (o *SwaApplicationSettings) GetEmOptInStatus() string {
	if o == nil || IsNil(o.EmOptInStatus) {
		var ret string
		return ret
	}
	return *o.EmOptInStatus
}

// GetEmOptInStatusOk returns a tuple with the EmOptInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettings) GetEmOptInStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EmOptInStatus) {
		return nil, false
	}
	return o.EmOptInStatus, true
}

// HasEmOptInStatus returns a boolean if a field has been set.
func (o *SwaApplicationSettings) HasEmOptInStatus() bool {
	if o != nil && !IsNil(o.EmOptInStatus) {
		return true
	}

	return false
}

// SetEmOptInStatus gets a reference to the given string and assigns it to the EmOptInStatus field.
func (o *SwaApplicationSettings) SetEmOptInStatus(v string) {
	o.EmOptInStatus = &v
}

// GetIdentityStoreId returns the IdentityStoreId field value if set, zero value otherwise.
func (o *SwaApplicationSettings) GetIdentityStoreId() string {
	if o == nil || IsNil(o.IdentityStoreId) {
		var ret string
		return ret
	}
	return *o.IdentityStoreId
}

// GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettings) GetIdentityStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityStoreId) {
		return nil, false
	}
	return o.IdentityStoreId, true
}

// HasIdentityStoreId returns a boolean if a field has been set.
func (o *SwaApplicationSettings) HasIdentityStoreId() bool {
	if o != nil && !IsNil(o.IdentityStoreId) {
		return true
	}

	return false
}

// SetIdentityStoreId gets a reference to the given string and assigns it to the IdentityStoreId field.
func (o *SwaApplicationSettings) SetIdentityStoreId(v string) {
	o.IdentityStoreId = &v
}

// GetImplicitAssignment returns the ImplicitAssignment field value if set, zero value otherwise.
func (o *SwaApplicationSettings) GetImplicitAssignment() bool {
	if o == nil || IsNil(o.ImplicitAssignment) {
		var ret bool
		return ret
	}
	return *o.ImplicitAssignment
}

// GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettings) GetImplicitAssignmentOk() (*bool, bool) {
	if o == nil || IsNil(o.ImplicitAssignment) {
		return nil, false
	}
	return o.ImplicitAssignment, true
}

// HasImplicitAssignment returns a boolean if a field has been set.
func (o *SwaApplicationSettings) HasImplicitAssignment() bool {
	if o != nil && !IsNil(o.ImplicitAssignment) {
		return true
	}

	return false
}

// SetImplicitAssignment gets a reference to the given bool and assigns it to the ImplicitAssignment field.
func (o *SwaApplicationSettings) SetImplicitAssignment(v bool) {
	o.ImplicitAssignment = &v
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *SwaApplicationSettings) GetInlineHookId() string {
	if o == nil || IsNil(o.InlineHookId) {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettings) GetInlineHookIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineHookId) {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *SwaApplicationSettings) HasInlineHookId() bool {
	if o != nil && !IsNil(o.InlineHookId) {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *SwaApplicationSettings) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SwaApplicationSettings) GetNotes() ApplicationSettingsNotes {
	if o == nil || IsNil(o.Notes) {
		var ret ApplicationSettingsNotes
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SwaApplicationSettings) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given ApplicationSettingsNotes and assigns it to the Notes field.
func (o *SwaApplicationSettings) SetNotes(v ApplicationSettingsNotes) {
	o.Notes = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *SwaApplicationSettings) GetNotifications() ApplicationSettingsNotifications {
	if o == nil || IsNil(o.Notifications) {
		var ret ApplicationSettingsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *SwaApplicationSettings) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given ApplicationSettingsNotifications and assigns it to the Notifications field.
func (o *SwaApplicationSettings) SetNotifications(v ApplicationSettingsNotifications) {
	o.Notifications = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *SwaApplicationSettings) GetApp() SwaApplicationSettingsApplication {
	if o == nil || IsNil(o.App) {
		var ret SwaApplicationSettingsApplication
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettings) GetAppOk() (*SwaApplicationSettingsApplication, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *SwaApplicationSettings) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given SwaApplicationSettingsApplication and assigns it to the App field.
func (o *SwaApplicationSettings) SetApp(v SwaApplicationSettingsApplication) {
	o.App = &v
}

func (o SwaApplicationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwaApplicationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmOptInStatus) {
		toSerialize["emOptInStatus"] = o.EmOptInStatus
	}
	if !IsNil(o.IdentityStoreId) {
		toSerialize["identityStoreId"] = o.IdentityStoreId
	}
	if !IsNil(o.ImplicitAssignment) {
		toSerialize["implicitAssignment"] = o.ImplicitAssignment
	}
	if !IsNil(o.InlineHookId) {
		toSerialize["inlineHookId"] = o.InlineHookId
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SwaApplicationSettings) UnmarshalJSON(data []byte) (err error) {
	varSwaApplicationSettings := _SwaApplicationSettings{}

	err = json.Unmarshal(data, &varSwaApplicationSettings)

	if err != nil {
		return err
	}

	*o = SwaApplicationSettings(varSwaApplicationSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "emOptInStatus")
		delete(additionalProperties, "identityStoreId")
		delete(additionalProperties, "implicitAssignment")
		delete(additionalProperties, "inlineHookId")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "app")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSwaApplicationSettings struct {
	value *SwaApplicationSettings
	isSet bool
}

func (v NullableSwaApplicationSettings) Get() *SwaApplicationSettings {
	return v.value
}

func (v *NullableSwaApplicationSettings) Set(val *SwaApplicationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSwaApplicationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSwaApplicationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwaApplicationSettings(val *SwaApplicationSettings) *NullableSwaApplicationSettings {
	return &NullableSwaApplicationSettings{value: val, isSet: true}
}

func (v NullableSwaApplicationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwaApplicationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
