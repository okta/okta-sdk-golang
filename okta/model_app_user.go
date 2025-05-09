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
	"time"
)

// AppUser The Application User object defines a user's app-specific profile and credentials for an app
type AppUser struct {
	Created *time.Time `json:"created,omitempty"`
	Credentials *AppUserCredentials `json:"credentials,omitempty"`
	// The ID of the user in the target app that's linked to the Okta Application User object. This value is the native app-specific identifier or primary key for the user in the target app.  The `externalId` is set during import when the user is confirmed (reconciled) or during provisioning when the user is created in the target app. This value isn't populated for SSO app assignments (for example, SAML or SWA) because it isn't synchronized with a target app.
	ExternalId *string `json:"externalId,omitempty"`
	// Unique identifier for the Okta User
	Id *string `json:"id,omitempty"`
	// Timestamp of the last synchronization operation. This value is only updated for apps with the `IMPORT_PROFILE_UPDATES` or `PUSH PROFILE_UPDATES` feature.
	LastSync *time.Time `json:"lastSync,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Timestamp when the Application User password was last changed
	PasswordChanged NullableTime `json:"passwordChanged,omitempty"`
	// Specifies the default and custom profile properties for a user. Properties that are visible in the Admin Console for an app assignment can also be assigned through the API. Some properties are reference properties that are imported from the target app and can't be configured. See [profile](/openapi/okta-management/management/tag/User/#tag/User/operation/getUser!c=200&path=profile&t=response). 
	Profile map[string]interface{} `json:"profile,omitempty"`
	// Indicates if the assignment is direct (`USER`) or by group membership (`GROUP`).
	Scope *string `json:"scope,omitempty"`
	// Status of an Application User
	Status *string `json:"status,omitempty"`
	// Timestamp when the Application User status was last changed
	StatusChanged *time.Time `json:"statusChanged,omitempty"`
	// The synchronization state for the Application User. The Application User's `syncState` depends on whether the `PROFILE_MASTERING` feature is enabled for the app.  > **Note:** User provisioning currently must be configured through the Admin Console.
	SyncState *string `json:"syncState,omitempty"`
	// Embedded resources related to the Application User using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links *LinksAppAndUser `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppUser AppUser

// NewAppUser instantiates a new AppUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUser() *AppUser {
	this := AppUser{}
	return &this
}

// NewAppUserWithDefaults instantiates a new AppUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserWithDefaults() *AppUser {
	this := AppUser{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AppUser) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AppUser) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AppUser) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *AppUser) GetCredentials() AppUserCredentials {
	if o == nil || o.Credentials == nil {
		var ret AppUserCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetCredentialsOk() (*AppUserCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *AppUser) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given AppUserCredentials and assigns it to the Credentials field.
func (o *AppUser) SetCredentials(v AppUserCredentials) {
	o.Credentials = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AppUser) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AppUser) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AppUser) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppUser) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppUser) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppUser) SetId(v string) {
	o.Id = &v
}

// GetLastSync returns the LastSync field value if set, zero value otherwise.
func (o *AppUser) GetLastSync() time.Time {
	if o == nil || o.LastSync == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSync
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetLastSyncOk() (*time.Time, bool) {
	if o == nil || o.LastSync == nil {
		return nil, false
	}
	return o.LastSync, true
}

// HasLastSync returns a boolean if a field has been set.
func (o *AppUser) HasLastSync() bool {
	if o != nil && o.LastSync != nil {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given time.Time and assigns it to the LastSync field.
func (o *AppUser) SetLastSync(v time.Time) {
	o.LastSync = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AppUser) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AppUser) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AppUser) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetPasswordChanged returns the PasswordChanged field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppUser) GetPasswordChanged() time.Time {
	if o == nil || o.PasswordChanged.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PasswordChanged.Get()
}

// GetPasswordChangedOk returns a tuple with the PasswordChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUser) GetPasswordChangedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordChanged.Get(), o.PasswordChanged.IsSet()
}

// HasPasswordChanged returns a boolean if a field has been set.
func (o *AppUser) HasPasswordChanged() bool {
	if o != nil && o.PasswordChanged.IsSet() {
		return true
	}

	return false
}

// SetPasswordChanged gets a reference to the given NullableTime and assigns it to the PasswordChanged field.
func (o *AppUser) SetPasswordChanged(v time.Time) {
	o.PasswordChanged.Set(&v)
}
// SetPasswordChangedNil sets the value for PasswordChanged to be an explicit nil
func (o *AppUser) SetPasswordChangedNil() {
	o.PasswordChanged.Set(nil)
}

// UnsetPasswordChanged ensures that no value is present for PasswordChanged, not even an explicit nil
func (o *AppUser) UnsetPasswordChanged() {
	o.PasswordChanged.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AppUser) GetProfile() map[string]interface{} {
	if o == nil || o.Profile == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AppUser) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *AppUser) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AppUser) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AppUser) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AppUser) SetScope(v string) {
	o.Scope = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppUser) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppUser) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppUser) SetStatus(v string) {
	o.Status = &v
}

// GetStatusChanged returns the StatusChanged field value if set, zero value otherwise.
func (o *AppUser) GetStatusChanged() time.Time {
	if o == nil || o.StatusChanged == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusChanged
}

// GetStatusChangedOk returns a tuple with the StatusChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetStatusChangedOk() (*time.Time, bool) {
	if o == nil || o.StatusChanged == nil {
		return nil, false
	}
	return o.StatusChanged, true
}

// HasStatusChanged returns a boolean if a field has been set.
func (o *AppUser) HasStatusChanged() bool {
	if o != nil && o.StatusChanged != nil {
		return true
	}

	return false
}

// SetStatusChanged gets a reference to the given time.Time and assigns it to the StatusChanged field.
func (o *AppUser) SetStatusChanged(v time.Time) {
	o.StatusChanged = &v
}

// GetSyncState returns the SyncState field value if set, zero value otherwise.
func (o *AppUser) GetSyncState() string {
	if o == nil || o.SyncState == nil {
		var ret string
		return ret
	}
	return *o.SyncState
}

// GetSyncStateOk returns a tuple with the SyncState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetSyncStateOk() (*string, bool) {
	if o == nil || o.SyncState == nil {
		return nil, false
	}
	return o.SyncState, true
}

// HasSyncState returns a boolean if a field has been set.
func (o *AppUser) HasSyncState() bool {
	if o != nil && o.SyncState != nil {
		return true
	}

	return false
}

// SetSyncState gets a reference to the given string and assigns it to the SyncState field.
func (o *AppUser) SetSyncState(v string) {
	o.SyncState = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *AppUser) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || o.Embedded == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *AppUser) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *AppUser) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppUser) GetLinks() LinksAppAndUser {
	if o == nil || o.Links == nil {
		var ret LinksAppAndUser
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUser) GetLinksOk() (*LinksAppAndUser, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppUser) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksAppAndUser and assigns it to the Links field.
func (o *AppUser) SetLinks(v LinksAppAndUser) {
	o.Links = &v
}

func (o AppUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastSync != nil {
		toSerialize["lastSync"] = o.LastSync
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.PasswordChanged.IsSet() {
		toSerialize["passwordChanged"] = o.PasswordChanged.Get()
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusChanged != nil {
		toSerialize["statusChanged"] = o.StatusChanged
	}
	if o.SyncState != nil {
		toSerialize["syncState"] = o.SyncState
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppUser) UnmarshalJSON(bytes []byte) (err error) {
	varAppUser := _AppUser{}

	err = json.Unmarshal(bytes, &varAppUser)
	if err == nil {
		*o = AppUser(varAppUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastSync")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "passwordChanged")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusChanged")
		delete(additionalProperties, "syncState")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppUser struct {
	value *AppUser
	isSet bool
}

func (v NullableAppUser) Get() *AppUser {
	return v.value
}

func (v *NullableAppUser) Set(val *AppUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUser(val *AppUser) *NullableAppUser {
	return &NullableAppUser{value: val, isSet: true}
}

func (v NullableAppUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

