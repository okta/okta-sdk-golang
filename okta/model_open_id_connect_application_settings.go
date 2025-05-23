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

// OpenIdConnectApplicationSettings struct for OpenIdConnectApplicationSettings
type OpenIdConnectApplicationSettings struct {
	IdentityStoreId *string `json:"identityStoreId,omitempty"`
	ImplicitAssignment *bool `json:"implicitAssignment,omitempty"`
	InlineHookId *string `json:"inlineHookId,omitempty"`
	Notes *ApplicationSettingsNotes `json:"notes,omitempty"`
	Notifications *ApplicationSettingsNotifications `json:"notifications,omitempty"`
	OauthClient *OpenIdConnectApplicationSettingsClient `json:"oauthClient,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenIdConnectApplicationSettings OpenIdConnectApplicationSettings

// NewOpenIdConnectApplicationSettings instantiates a new OpenIdConnectApplicationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectApplicationSettings() *OpenIdConnectApplicationSettings {
	this := OpenIdConnectApplicationSettings{}
	return &this
}

// NewOpenIdConnectApplicationSettingsWithDefaults instantiates a new OpenIdConnectApplicationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIdConnectApplicationSettingsWithDefaults() *OpenIdConnectApplicationSettings {
	this := OpenIdConnectApplicationSettings{}
	return &this
}

// GetIdentityStoreId returns the IdentityStoreId field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettings) GetIdentityStoreId() string {
	if o == nil || o.IdentityStoreId == nil {
		var ret string
		return ret
	}
	return *o.IdentityStoreId
}

// GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettings) GetIdentityStoreIdOk() (*string, bool) {
	if o == nil || o.IdentityStoreId == nil {
		return nil, false
	}
	return o.IdentityStoreId, true
}

// HasIdentityStoreId returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettings) HasIdentityStoreId() bool {
	if o != nil && o.IdentityStoreId != nil {
		return true
	}

	return false
}

// SetIdentityStoreId gets a reference to the given string and assigns it to the IdentityStoreId field.
func (o *OpenIdConnectApplicationSettings) SetIdentityStoreId(v string) {
	o.IdentityStoreId = &v
}

// GetImplicitAssignment returns the ImplicitAssignment field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettings) GetImplicitAssignment() bool {
	if o == nil || o.ImplicitAssignment == nil {
		var ret bool
		return ret
	}
	return *o.ImplicitAssignment
}

// GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettings) GetImplicitAssignmentOk() (*bool, bool) {
	if o == nil || o.ImplicitAssignment == nil {
		return nil, false
	}
	return o.ImplicitAssignment, true
}

// HasImplicitAssignment returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettings) HasImplicitAssignment() bool {
	if o != nil && o.ImplicitAssignment != nil {
		return true
	}

	return false
}

// SetImplicitAssignment gets a reference to the given bool and assigns it to the ImplicitAssignment field.
func (o *OpenIdConnectApplicationSettings) SetImplicitAssignment(v bool) {
	o.ImplicitAssignment = &v
}

// GetInlineHookId returns the InlineHookId field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettings) GetInlineHookId() string {
	if o == nil || o.InlineHookId == nil {
		var ret string
		return ret
	}
	return *o.InlineHookId
}

// GetInlineHookIdOk returns a tuple with the InlineHookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettings) GetInlineHookIdOk() (*string, bool) {
	if o == nil || o.InlineHookId == nil {
		return nil, false
	}
	return o.InlineHookId, true
}

// HasInlineHookId returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettings) HasInlineHookId() bool {
	if o != nil && o.InlineHookId != nil {
		return true
	}

	return false
}

// SetInlineHookId gets a reference to the given string and assigns it to the InlineHookId field.
func (o *OpenIdConnectApplicationSettings) SetInlineHookId(v string) {
	o.InlineHookId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettings) GetNotes() ApplicationSettingsNotes {
	if o == nil || o.Notes == nil {
		var ret ApplicationSettingsNotes
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettings) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given ApplicationSettingsNotes and assigns it to the Notes field.
func (o *OpenIdConnectApplicationSettings) SetNotes(v ApplicationSettingsNotes) {
	o.Notes = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettings) GetNotifications() ApplicationSettingsNotifications {
	if o == nil || o.Notifications == nil {
		var ret ApplicationSettingsNotifications
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettings) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given ApplicationSettingsNotifications and assigns it to the Notifications field.
func (o *OpenIdConnectApplicationSettings) SetNotifications(v ApplicationSettingsNotifications) {
	o.Notifications = &v
}

// GetOauthClient returns the OauthClient field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationSettings) GetOauthClient() OpenIdConnectApplicationSettingsClient {
	if o == nil || o.OauthClient == nil {
		var ret OpenIdConnectApplicationSettingsClient
		return ret
	}
	return *o.OauthClient
}

// GetOauthClientOk returns a tuple with the OauthClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationSettings) GetOauthClientOk() (*OpenIdConnectApplicationSettingsClient, bool) {
	if o == nil || o.OauthClient == nil {
		return nil, false
	}
	return o.OauthClient, true
}

// HasOauthClient returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationSettings) HasOauthClient() bool {
	if o != nil && o.OauthClient != nil {
		return true
	}

	return false
}

// SetOauthClient gets a reference to the given OpenIdConnectApplicationSettingsClient and assigns it to the OauthClient field.
func (o *OpenIdConnectApplicationSettings) SetOauthClient(v OpenIdConnectApplicationSettingsClient) {
	o.OauthClient = &v
}

func (o OpenIdConnectApplicationSettings) MarshalJSON() ([]byte, error) {
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
	if o.OauthClient != nil {
		toSerialize["oauthClient"] = o.OauthClient
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OpenIdConnectApplicationSettings) UnmarshalJSON(bytes []byte) (err error) {
	varOpenIdConnectApplicationSettings := _OpenIdConnectApplicationSettings{}

	err = json.Unmarshal(bytes, &varOpenIdConnectApplicationSettings)
	if err == nil {
		*o = OpenIdConnectApplicationSettings(varOpenIdConnectApplicationSettings)
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
		delete(additionalProperties, "oauthClient")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOpenIdConnectApplicationSettings struct {
	value *OpenIdConnectApplicationSettings
	isSet bool
}

func (v NullableOpenIdConnectApplicationSettings) Get() *OpenIdConnectApplicationSettings {
	return v.value
}

func (v *NullableOpenIdConnectApplicationSettings) Set(val *OpenIdConnectApplicationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectApplicationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectApplicationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectApplicationSettings(val *OpenIdConnectApplicationSettings) *NullableOpenIdConnectApplicationSettings {
	return &NullableOpenIdConnectApplicationSettings{value: val, isSet: true}
}

func (v NullableOpenIdConnectApplicationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectApplicationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

