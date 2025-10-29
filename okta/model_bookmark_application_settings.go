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

// checks if the BookmarkApplicationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkApplicationSettings{}

// BookmarkApplicationSettings struct for BookmarkApplicationSettings
type BookmarkApplicationSettings struct {
	// The Governance Engine opt-in status for the app
	EmOptInStatus *string `json:"emOptInStatus,omitempty"`
	// Identifies an additional identity store app, if your app supports it. The `identityStoreId` value must be a valid identity store app ID. This identity store app must be created in the same org as your app.
	IdentityStoreId *string `json:"identityStoreId,omitempty"`
	// Controls whether Okta automatically assigns users to the app based on the user's role or group membership.
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty"`
	// Identifier of an inline hook. Inline hooks are outbound calls from Okta to your own custom code, triggered at specific points in Okta process flows. They allow you to integrate custom functionality into those flows. See [Inline hooks](/openapi/okta-management/management/tag/InlineHook/).
	InlineHookId         *string                                 `json:"inlineHookId,omitempty"`
	Notes                *ApplicationSettingsNotes               `json:"notes,omitempty"`
	Notifications        *ApplicationSettingsNotifications       `json:"notifications,omitempty"`
	App                  *BookmarkApplicationSettingsApplication `json:"app,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BookmarkApplicationSettings BookmarkApplicationSettings

// NewBookmarkApplicationSettings instantiates a new BookmarkApplicationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkApplicationSettings() *BookmarkApplicationSettings {
	this := BookmarkApplicationSettings{}
	return &this
}

// NewBookmarkApplicationSettingsWithDefaults instantiates a new BookmarkApplicationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkApplicationSettingsWithDefaults() *BookmarkApplicationSettings {
	this := BookmarkApplicationSettings{}
	return &this
}

// GetEmOptInStatus returns the EmOptInStatus field value if set, zero value otherwise.
func (o *BookmarkApplicationSettings) GetEmOptInStatus() string {
	if o == nil || IsNil(o.EmOptInStatus) {
		var ret string
		return ret
	}
	return *o.EmOptInStatus
}

// GetEmOptInStatusOk returns a tuple with the EmOptInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettings) GetEmOptInStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EmOptInStatus) {
		return nil, false
	}
	return o.EmOptInStatus, true
}

// HasEmOptInStatus returns a boolean if a field has been set.
func (o *BookmarkApplicationSettings) HasEmOptInStatus() bool {
	if o != nil && !IsNil(o.EmOptInStatus) {
		return true
	}

	return false
}

// SetEmOptInStatus gets a reference to the given string and assigns it to the EmOptInStatus field.
func (o *BookmarkApplicationSettings) SetEmOptInStatus(v string) {
	o.EmOptInStatus = &v
}

// GetIdentityStoreId returns the IdentityStoreId field value if set, zero value otherwise.
func (o *BookmarkApplicationSettings) GetIdentityStoreId() string {
	if o == nil || IsNil(o.IdentityStoreId) {
		var ret string
		return ret
	}
	return *o.IdentityStoreId
}

// GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettings) GetIdentityStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityStoreId) {
		return nil, false
	}
	return o.IdentityStoreId, true
}

// HasIdentityStoreId returns a boolean if a field has been set.
func (o *BookmarkApplicationSettings) HasIdentityStoreId() bool {
	if o != nil && !IsNil(o.IdentityStoreId) {
		return true
	}

	return false
}

// SetIdentityStoreId gets a reference to the given string and assigns it to the IdentityStoreId field.
func (o *BookmarkApplicationSettings) SetIdentityStoreId(v string) {
	o.IdentityStoreId = &v
}

// GetImplicitAssignment returns the ImplicitAssignment field value if set, zero value otherwise.
func (o *BookmarkApplicationSettings) GetImplicitAssignment() bool {
	if o == nil || IsNil(o.ImplicitAssignment) {
		var ret bool
		return ret
	}
	return *o.ImplicitAssignment
}

// GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettings) GetImplicitAssignmentOk() (*bool, bool) {
	if o == nil || IsNil(o.ImplicitAssignment) {
		return nil, false
	}
	return o.ImplicitAssignment, true
}

// HasImplicitAssignment returns a boolean if a field has been set.
func (o *BookmarkApplicationSettings) HasImplicitAssignment() bool {
	if o != nil && !IsNil(o.ImplicitAssignment) {
		return true
	}

	return false
}

// SetImplicitAssignment gets a reference to the given bool and assigns it to the ImplicitAssignment field.
func (o *BookmarkApplicationSettings) SetImplicitAssignment(v bool) {
	o.ImplicitAssignment = &v
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *BookmarkApplicationSettings) GetInlineHookId() string {
	if o == nil || IsNil(o.InlineHookId) {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettings) GetInlineHookIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineHookId) {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *BookmarkApplicationSettings) HasInlineHookId() bool {
	if o != nil && !IsNil(o.InlineHookId) {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *BookmarkApplicationSettings) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *BookmarkApplicationSettings) GetNotes() ApplicationSettingsNotes {
	if o == nil || IsNil(o.Notes) {
		var ret ApplicationSettingsNotes
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *BookmarkApplicationSettings) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given ApplicationSettingsNotes and assigns it to the Notes field.
func (o *BookmarkApplicationSettings) SetNotes(v ApplicationSettingsNotes) {
	o.Notes = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *BookmarkApplicationSettings) GetNotifications() ApplicationSettingsNotifications {
	if o == nil || IsNil(o.Notifications) {
		var ret ApplicationSettingsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *BookmarkApplicationSettings) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given ApplicationSettingsNotifications and assigns it to the Notifications field.
func (o *BookmarkApplicationSettings) SetNotifications(v ApplicationSettingsNotifications) {
	o.Notifications = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *BookmarkApplicationSettings) GetApp() BookmarkApplicationSettingsApplication {
	if o == nil || IsNil(o.App) {
		var ret BookmarkApplicationSettingsApplication
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkApplicationSettings) GetAppOk() (*BookmarkApplicationSettingsApplication, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *BookmarkApplicationSettings) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given BookmarkApplicationSettingsApplication and assigns it to the App field.
func (o *BookmarkApplicationSettings) SetApp(v BookmarkApplicationSettingsApplication) {
	o.App = &v
}

func (o BookmarkApplicationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkApplicationSettings) ToMap() (map[string]interface{}, error) {
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

func (o *BookmarkApplicationSettings) UnmarshalJSON(data []byte) (err error) {
	varBookmarkApplicationSettings := _BookmarkApplicationSettings{}

	err = json.Unmarshal(data, &varBookmarkApplicationSettings)

	if err != nil {
		return err
	}

	*o = BookmarkApplicationSettings(varBookmarkApplicationSettings)

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

type NullableBookmarkApplicationSettings struct {
	value *BookmarkApplicationSettings
	isSet bool
}

func (v NullableBookmarkApplicationSettings) Get() *BookmarkApplicationSettings {
	return v.value
}

func (v *NullableBookmarkApplicationSettings) Set(val *BookmarkApplicationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkApplicationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkApplicationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkApplicationSettings(val *BookmarkApplicationSettings) *NullableBookmarkApplicationSettings {
	return &NullableBookmarkApplicationSettings{value: val, isSet: true}
}

func (v NullableBookmarkApplicationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkApplicationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
