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

// Saml11ApplicationSettings struct for Saml11ApplicationSettings
type Saml11ApplicationSettings struct {
	IdentityStoreId *string `json:"identityStoreId,omitempty"`
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty"`
	InlineHookId *string `json:"inlineHookId,omitempty"`
	Notes *ApplicationSettingsNotes `json:"notes,omitempty"`
	Notifications *ApplicationSettingsNotifications `json:"notifications,omitempty"`
	App map[string]interface{} `json:"app,omitempty"`
	SignOn *Saml11ApplicationSettingsSignOn `json:"signOn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Saml11ApplicationSettings Saml11ApplicationSettings

// NewSaml11ApplicationSettings instantiates a new Saml11ApplicationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaml11ApplicationSettings() *Saml11ApplicationSettings {
	this := Saml11ApplicationSettings{}
	return &this
}

// NewSaml11ApplicationSettingsWithDefaults instantiates a new Saml11ApplicationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaml11ApplicationSettingsWithDefaults() *Saml11ApplicationSettings {
	this := Saml11ApplicationSettings{}
	return &this
}

// GetIdentityStoreId returns the IdentityStoreId field value if set, zero value otherwise.
func (o *Saml11ApplicationSettings) GetIdentityStoreId() string {
	if o == nil || o.IdentityStoreId == nil {
		var ret string
		return ret
	}
	return *o.IdentityStoreId
}

// GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettings) GetIdentityStoreIdOk() (*string, bool) {
	if o == nil || o.IdentityStoreId == nil {
		return nil, false
	}
	return o.IdentityStoreId, true
}

// HasIdentityStoreId returns a boolean if a field has been set.
func (o *Saml11ApplicationSettings) HasIdentityStoreId() bool {
	if o != nil && o.IdentityStoreId != nil {
		return true
	}

	return false
}

// SetIdentityStoreId gets a reference to the given string and assigns it to the IdentityStoreId field.
func (o *Saml11ApplicationSettings) SetIdentityStoreId(v string) {
	o.IdentityStoreId = &v
}

// GetImplicitAssignment returns the ImplicitAssignment field value if set, zero value otherwise.
func (o *Saml11ApplicationSettings) GetImplicitAssignment() bool {
	if o == nil || o.ImplicitAssignment == nil {
		var ret bool
		return ret
	}
	return *o.ImplicitAssignment
}

// GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettings) GetImplicitAssignmentOk() (*bool, bool) {
	if o == nil || o.ImplicitAssignment == nil {
		return nil, false
	}
	return o.ImplicitAssignment, true
}

// HasImplicitAssignment returns a boolean if a field has been set.
func (o *Saml11ApplicationSettings) HasImplicitAssignment() bool {
	if o != nil && o.ImplicitAssignment != nil {
		return true
	}

	return false
}

// SetImplicitAssignment gets a reference to the given bool and assigns it to the ImplicitAssignment field.
func (o *Saml11ApplicationSettings) SetImplicitAssignment(v bool) {
	o.ImplicitAssignment = &v
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *Saml11ApplicationSettings) GetInlineHookId() string {
	if o == nil || o.InlineHookId == nil {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettings) GetInlineHookIdOk() (*string, bool) {
	if o == nil || o.InlineHookId == nil {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *Saml11ApplicationSettings) HasInlineHookId() bool {
	if o != nil && o.InlineHookId != nil {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *Saml11ApplicationSettings) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Saml11ApplicationSettings) GetNotes() ApplicationSettingsNotes {
	if o == nil || o.Notes == nil {
		var ret ApplicationSettingsNotes
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Saml11ApplicationSettings) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given ApplicationSettingsNotes and assigns it to the Notes field.
func (o *Saml11ApplicationSettings) SetNotes(v ApplicationSettingsNotes) {
	o.Notes = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *Saml11ApplicationSettings) GetNotifications() ApplicationSettingsNotifications {
	if o == nil || o.Notifications == nil {
		var ret ApplicationSettingsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *Saml11ApplicationSettings) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given ApplicationSettingsNotifications and assigns it to the Notifications field.
func (o *Saml11ApplicationSettings) SetNotifications(v ApplicationSettingsNotifications) {
	o.Notifications = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *Saml11ApplicationSettings) GetApp() map[string]interface{} {
	if o == nil || o.App == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettings) GetAppOk() (map[string]interface{}, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *Saml11ApplicationSettings) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given map[string]interface{} and assigns it to the App field.
func (o *Saml11ApplicationSettings) SetApp(v map[string]interface{}) {
	o.App = v
}

// GetSignOn returns the SignOn field value if set, zero value otherwise.
func (o *Saml11ApplicationSettings) GetSignOn() Saml11ApplicationSettingsSignOn {
	if o == nil || o.SignOn == nil {
		var ret Saml11ApplicationSettingsSignOn
		return ret
	}
	return *o.SignOn
}

// GetSignOnOk returns a tuple with the SignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml11ApplicationSettings) GetSignOnOk() (*Saml11ApplicationSettingsSignOn, bool) {
	if o == nil || o.SignOn == nil {
		return nil, false
	}
	return o.SignOn, true
}

// HasSignOn returns a boolean if a field has been set.
func (o *Saml11ApplicationSettings) HasSignOn() bool {
	if o != nil && o.SignOn != nil {
		return true
	}

	return false
}

// SetSignOn gets a reference to the given Saml11ApplicationSettingsSignOn and assigns it to the SignOn field.
func (o *Saml11ApplicationSettings) SetSignOn(v Saml11ApplicationSettingsSignOn) {
	o.SignOn = &v
}

func (o Saml11ApplicationSettings) MarshalJSON() ([]byte, error) {
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

func (o *Saml11ApplicationSettings) UnmarshalJSON(bytes []byte) (err error) {
	varSaml11ApplicationSettings := _Saml11ApplicationSettings{}

	err = json.Unmarshal(bytes, &varSaml11ApplicationSettings)
	if err == nil {
		*o = Saml11ApplicationSettings(varSaml11ApplicationSettings)
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

type NullableSaml11ApplicationSettings struct {
	value *Saml11ApplicationSettings
	isSet bool
}

func (v NullableSaml11ApplicationSettings) Get() *Saml11ApplicationSettings {
	return v.value
}

func (v *NullableSaml11ApplicationSettings) Set(val *Saml11ApplicationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSaml11ApplicationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSaml11ApplicationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaml11ApplicationSettings(val *Saml11ApplicationSettings) *NullableSaml11ApplicationSettings {
	return &NullableSaml11ApplicationSettings{value: val, isSet: true}
}

func (v NullableSaml11ApplicationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaml11ApplicationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

