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

// checks if the AutoLoginApplicationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoLoginApplicationSettings{}

// AutoLoginApplicationSettings struct for AutoLoginApplicationSettings
type AutoLoginApplicationSettings struct {
	// The entitlement management opt-in status for the app
	EmOptInStatus *string `json:"emOptInStatus,omitempty"`
	// Identifies an additional identity store app, if your app supports it. The `identityStoreId` value must be a valid identity store app ID. This identity store app must be created in the same org as your app.
	IdentityStoreId *string `json:"identityStoreId,omitempty"`
	// Controls whether Okta automatically assigns users to the app based on the user's role or group membership.
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty"`
	// Identifier of an inline hook. Inline hooks are outbound calls from Okta to your own custom code, triggered at specific points in Okta process flows. They allow you to integrate custom functionality into those flows. See [Inline hooks](/openapi/okta-management/management/tag/InlineHook/).
	InlineHookId         *string                             `json:"inlineHookId,omitempty"`
	Notes                *ApplicationSettingsNotes           `json:"notes,omitempty"`
	Notifications        *ApplicationSettingsNotifications   `json:"notifications,omitempty"`
	SignOn               *AutoLoginApplicationSettingsSignOn `json:"signOn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoLoginApplicationSettings AutoLoginApplicationSettings

// NewAutoLoginApplicationSettings instantiates a new AutoLoginApplicationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoLoginApplicationSettings() *AutoLoginApplicationSettings {
	this := AutoLoginApplicationSettings{}
	return &this
}

// NewAutoLoginApplicationSettingsWithDefaults instantiates a new AutoLoginApplicationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoLoginApplicationSettingsWithDefaults() *AutoLoginApplicationSettings {
	this := AutoLoginApplicationSettings{}
	return &this
}

// GetEmOptInStatus returns the EmOptInStatus field value if set, zero value otherwise.
func (o *AutoLoginApplicationSettings) GetEmOptInStatus() string {
	if o == nil || IsNil(o.EmOptInStatus) {
		var ret string
		return ret
	}
	return *o.EmOptInStatus
}

// GetEmOptInStatusOk returns a tuple with the EmOptInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplicationSettings) GetEmOptInStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EmOptInStatus) {
		return nil, false
	}
	return o.EmOptInStatus, true
}

// HasEmOptInStatus returns a boolean if a field has been set.
func (o *AutoLoginApplicationSettings) HasEmOptInStatus() bool {
	if o != nil && !IsNil(o.EmOptInStatus) {
		return true
	}

	return false
}

// SetEmOptInStatus gets a reference to the given string and assigns it to the EmOptInStatus field.
func (o *AutoLoginApplicationSettings) SetEmOptInStatus(v string) {
	o.EmOptInStatus = &v
}

// GetIdentityStoreId returns the IdentityStoreId field value if set, zero value otherwise.
func (o *AutoLoginApplicationSettings) GetIdentityStoreId() string {
	if o == nil || IsNil(o.IdentityStoreId) {
		var ret string
		return ret
	}
	return *o.IdentityStoreId
}

// GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplicationSettings) GetIdentityStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityStoreId) {
		return nil, false
	}
	return o.IdentityStoreId, true
}

// HasIdentityStoreId returns a boolean if a field has been set.
func (o *AutoLoginApplicationSettings) HasIdentityStoreId() bool {
	if o != nil && !IsNil(o.IdentityStoreId) {
		return true
	}

	return false
}

// SetIdentityStoreId gets a reference to the given string and assigns it to the IdentityStoreId field.
func (o *AutoLoginApplicationSettings) SetIdentityStoreId(v string) {
	o.IdentityStoreId = &v
}

// GetImplicitAssignment returns the ImplicitAssignment field value if set, zero value otherwise.
func (o *AutoLoginApplicationSettings) GetImplicitAssignment() bool {
	if o == nil || IsNil(o.ImplicitAssignment) {
		var ret bool
		return ret
	}
	return *o.ImplicitAssignment
}

// GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplicationSettings) GetImplicitAssignmentOk() (*bool, bool) {
	if o == nil || IsNil(o.ImplicitAssignment) {
		return nil, false
	}
	return o.ImplicitAssignment, true
}

// HasImplicitAssignment returns a boolean if a field has been set.
func (o *AutoLoginApplicationSettings) HasImplicitAssignment() bool {
	if o != nil && !IsNil(o.ImplicitAssignment) {
		return true
	}

	return false
}

// SetImplicitAssignment gets a reference to the given bool and assigns it to the ImplicitAssignment field.
func (o *AutoLoginApplicationSettings) SetImplicitAssignment(v bool) {
	o.ImplicitAssignment = &v
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *AutoLoginApplicationSettings) GetInlineHookId() string {
	if o == nil || IsNil(o.InlineHookId) {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplicationSettings) GetInlineHookIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineHookId) {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *AutoLoginApplicationSettings) HasInlineHookId() bool {
	if o != nil && !IsNil(o.InlineHookId) {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *AutoLoginApplicationSettings) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AutoLoginApplicationSettings) GetNotes() ApplicationSettingsNotes {
	if o == nil || IsNil(o.Notes) {
		var ret ApplicationSettingsNotes
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AutoLoginApplicationSettings) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given ApplicationSettingsNotes and assigns it to the Notes field.
func (o *AutoLoginApplicationSettings) SetNotes(v ApplicationSettingsNotes) {
	o.Notes = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *AutoLoginApplicationSettings) GetNotifications() ApplicationSettingsNotifications {
	if o == nil || IsNil(o.Notifications) {
		var ret ApplicationSettingsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *AutoLoginApplicationSettings) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given ApplicationSettingsNotifications and assigns it to the Notifications field.
func (o *AutoLoginApplicationSettings) SetNotifications(v ApplicationSettingsNotifications) {
	o.Notifications = &v
}

// GetSignOn returns the SignOn field value if set, zero value otherwise.
func (o *AutoLoginApplicationSettings) GetSignOn() AutoLoginApplicationSettingsSignOn {
	if o == nil || IsNil(o.SignOn) {
		var ret AutoLoginApplicationSettingsSignOn
		return ret
	}
	return *o.SignOn
}

// GetSignOnOk returns a tuple with the SignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoLoginApplicationSettings) GetSignOnOk() (*AutoLoginApplicationSettingsSignOn, bool) {
	if o == nil || IsNil(o.SignOn) {
		return nil, false
	}
	return o.SignOn, true
}

// HasSignOn returns a boolean if a field has been set.
func (o *AutoLoginApplicationSettings) HasSignOn() bool {
	if o != nil && !IsNil(o.SignOn) {
		return true
	}

	return false
}

// SetSignOn gets a reference to the given AutoLoginApplicationSettingsSignOn and assigns it to the SignOn field.
func (o *AutoLoginApplicationSettings) SetSignOn(v AutoLoginApplicationSettingsSignOn) {
	o.SignOn = &v
}

func (o AutoLoginApplicationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoLoginApplicationSettings) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.SignOn) {
		toSerialize["signOn"] = o.SignOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutoLoginApplicationSettings) UnmarshalJSON(data []byte) (err error) {
	varAutoLoginApplicationSettings := _AutoLoginApplicationSettings{}

	err = json.Unmarshal(data, &varAutoLoginApplicationSettings)

	if err != nil {
		return err
	}

	*o = AutoLoginApplicationSettings(varAutoLoginApplicationSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "emOptInStatus")
		delete(additionalProperties, "identityStoreId")
		delete(additionalProperties, "implicitAssignment")
		delete(additionalProperties, "inlineHookId")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "signOn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoLoginApplicationSettings struct {
	value *AutoLoginApplicationSettings
	isSet bool
}

func (v NullableAutoLoginApplicationSettings) Get() *AutoLoginApplicationSettings {
	return v.value
}

func (v *NullableAutoLoginApplicationSettings) Set(val *AutoLoginApplicationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoLoginApplicationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoLoginApplicationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoLoginApplicationSettings(val *AutoLoginApplicationSettings) *NullableAutoLoginApplicationSettings {
	return &NullableAutoLoginApplicationSettings{value: val, isSet: true}
}

func (v NullableAutoLoginApplicationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoLoginApplicationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
