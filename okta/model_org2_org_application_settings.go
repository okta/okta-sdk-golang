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
	"fmt"
)

// checks if the Org2OrgApplicationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Org2OrgApplicationSettings{}

// Org2OrgApplicationSettings struct for Org2OrgApplicationSettings
type Org2OrgApplicationSettings struct {
	// The entitlement management opt-in status for the app
	EmOptInStatus *string `json:"emOptInStatus,omitempty"`
	// Identifies an additional identity store app, if your app supports it. The `identityStoreId` value must be a valid identity store app ID. This identity store app must be created in the same org as your app.
	IdentityStoreId *string `json:"identityStoreId,omitempty"`
	// Controls whether Okta automatically assigns users to the app based on the user's role or group membership.
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty"`
	// Identifier of an inline hook. Inline hooks are outbound calls from Okta to your own custom code, triggered at specific points in Okta process flows. They allow you to integrate custom functionality into those flows. See [Inline hooks](/openapi/okta-management/management/tag/InlineHook/).
	InlineHookId         *string                               `json:"inlineHookId,omitempty"`
	Notes                *ApplicationSettingsNotes             `json:"notes,omitempty"`
	Notifications        *ApplicationSettingsNotifications     `json:"notifications,omitempty"`
	App                  Org2OrgApplicationSettingsApplication `json:"app"`
	SignOn               *OINSaml20ApplicationSettingsSignOn   `json:"signOn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Org2OrgApplicationSettings Org2OrgApplicationSettings

// NewOrg2OrgApplicationSettings instantiates a new Org2OrgApplicationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrg2OrgApplicationSettings(app Org2OrgApplicationSettingsApplication) *Org2OrgApplicationSettings {
	this := Org2OrgApplicationSettings{}
	this.App = app
	return &this
}

// NewOrg2OrgApplicationSettingsWithDefaults instantiates a new Org2OrgApplicationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrg2OrgApplicationSettingsWithDefaults() *Org2OrgApplicationSettings {
	this := Org2OrgApplicationSettings{}
	return &this
}

// GetEmOptInStatus returns the EmOptInStatus field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettings) GetEmOptInStatus() string {
	if o == nil || IsNil(o.EmOptInStatus) {
		var ret string
		return ret
	}
	return *o.EmOptInStatus
}

// GetEmOptInStatusOk returns a tuple with the EmOptInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettings) GetEmOptInStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EmOptInStatus) {
		return nil, false
	}
	return o.EmOptInStatus, true
}

// HasEmOptInStatus returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettings) HasEmOptInStatus() bool {
	if o != nil && !IsNil(o.EmOptInStatus) {
		return true
	}

	return false
}

// SetEmOptInStatus gets a reference to the given string and assigns it to the EmOptInStatus field.
func (o *Org2OrgApplicationSettings) SetEmOptInStatus(v string) {
	o.EmOptInStatus = &v
}

// GetIdentityStoreId returns the IdentityStoreId field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettings) GetIdentityStoreId() string {
	if o == nil || IsNil(o.IdentityStoreId) {
		var ret string
		return ret
	}
	return *o.IdentityStoreId
}

// GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettings) GetIdentityStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityStoreId) {
		return nil, false
	}
	return o.IdentityStoreId, true
}

// HasIdentityStoreId returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettings) HasIdentityStoreId() bool {
	if o != nil && !IsNil(o.IdentityStoreId) {
		return true
	}

	return false
}

// SetIdentityStoreId gets a reference to the given string and assigns it to the IdentityStoreId field.
func (o *Org2OrgApplicationSettings) SetIdentityStoreId(v string) {
	o.IdentityStoreId = &v
}

// GetImplicitAssignment returns the ImplicitAssignment field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettings) GetImplicitAssignment() bool {
	if o == nil || IsNil(o.ImplicitAssignment) {
		var ret bool
		return ret
	}
	return *o.ImplicitAssignment
}

// GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettings) GetImplicitAssignmentOk() (*bool, bool) {
	if o == nil || IsNil(o.ImplicitAssignment) {
		return nil, false
	}
	return o.ImplicitAssignment, true
}

// HasImplicitAssignment returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettings) HasImplicitAssignment() bool {
	if o != nil && !IsNil(o.ImplicitAssignment) {
		return true
	}

	return false
}

// SetImplicitAssignment gets a reference to the given bool and assigns it to the ImplicitAssignment field.
func (o *Org2OrgApplicationSettings) SetImplicitAssignment(v bool) {
	o.ImplicitAssignment = &v
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettings) GetInlineHookId() string {
	if o == nil || IsNil(o.InlineHookId) {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettings) GetInlineHookIdOk() (*string, bool) {
	if o == nil || IsNil(o.InlineHookId) {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettings) HasInlineHookId() bool {
	if o != nil && !IsNil(o.InlineHookId) {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *Org2OrgApplicationSettings) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettings) GetNotes() ApplicationSettingsNotes {
	if o == nil || IsNil(o.Notes) {
		var ret ApplicationSettingsNotes
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettings) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given ApplicationSettingsNotes and assigns it to the Notes field.
func (o *Org2OrgApplicationSettings) SetNotes(v ApplicationSettingsNotes) {
	o.Notes = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettings) GetNotifications() ApplicationSettingsNotifications {
	if o == nil || IsNil(o.Notifications) {
		var ret ApplicationSettingsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettings) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given ApplicationSettingsNotifications and assigns it to the Notifications field.
func (o *Org2OrgApplicationSettings) SetNotifications(v ApplicationSettingsNotifications) {
	o.Notifications = &v
}

// GetApp returns the App field value
func (o *Org2OrgApplicationSettings) GetApp() Org2OrgApplicationSettingsApplication {
	if o == nil {
		var ret Org2OrgApplicationSettingsApplication
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettings) GetAppOk() (*Org2OrgApplicationSettingsApplication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *Org2OrgApplicationSettings) SetApp(v Org2OrgApplicationSettingsApplication) {
	o.App = v
}

// GetSignOn returns the SignOn field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettings) GetSignOn() OINSaml20ApplicationSettingsSignOn {
	if o == nil || IsNil(o.SignOn) {
		var ret OINSaml20ApplicationSettingsSignOn
		return ret
	}
	return *o.SignOn
}

// GetSignOnOk returns a tuple with the SignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettings) GetSignOnOk() (*OINSaml20ApplicationSettingsSignOn, bool) {
	if o == nil || IsNil(o.SignOn) {
		return nil, false
	}
	return o.SignOn, true
}

// HasSignOn returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettings) HasSignOn() bool {
	if o != nil && !IsNil(o.SignOn) {
		return true
	}

	return false
}

// SetSignOn gets a reference to the given OINSaml20ApplicationSettingsSignOn and assigns it to the SignOn field.
func (o *Org2OrgApplicationSettings) SetSignOn(v OINSaml20ApplicationSettingsSignOn) {
	o.SignOn = &v
}

func (o Org2OrgApplicationSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Org2OrgApplicationSettings) ToMap() (map[string]interface{}, error) {
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
	toSerialize["app"] = o.App
	if !IsNil(o.SignOn) {
		toSerialize["signOn"] = o.SignOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Org2OrgApplicationSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrg2OrgApplicationSettings := _Org2OrgApplicationSettings{}

	err = json.Unmarshal(data, &varOrg2OrgApplicationSettings)

	if err != nil {
		return err
	}

	*o = Org2OrgApplicationSettings(varOrg2OrgApplicationSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "emOptInStatus")
		delete(additionalProperties, "identityStoreId")
		delete(additionalProperties, "implicitAssignment")
		delete(additionalProperties, "inlineHookId")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "app")
		delete(additionalProperties, "signOn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrg2OrgApplicationSettings struct {
	value *Org2OrgApplicationSettings
	isSet bool
}

func (v NullableOrg2OrgApplicationSettings) Get() *Org2OrgApplicationSettings {
	return v.value
}

func (v *NullableOrg2OrgApplicationSettings) Set(val *Org2OrgApplicationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOrg2OrgApplicationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOrg2OrgApplicationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrg2OrgApplicationSettings(val *Org2OrgApplicationSettings) *NullableOrg2OrgApplicationSettings {
	return &NullableOrg2OrgApplicationSettings{value: val, isSet: true}
}

func (v NullableOrg2OrgApplicationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrg2OrgApplicationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
