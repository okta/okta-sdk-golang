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
	"time"
)

// checks if the AppUserAssignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppUserAssignRequest{}

// AppUserAssignRequest struct for AppUserAssignRequest
type AppUserAssignRequest struct {
	// Timestamp when the object was created
	Created     *time.Time          `json:"created,omitempty"`
	Credentials *AppUserCredentials `json:"credentials,omitempty"`
	// The ID of the user in the target app that's linked to the Okta application user object. This value is the native app-specific identifier or primary key for the user in the target app.  The `externalId` is set during import when the user is confirmed (reconciled) or during provisioning when the user is created in the target app. This value isn't populated for SSO app assignments (for example, SAML or SWA) because it isn't synchronized with a target app.
	ExternalId *string `json:"externalId,omitempty"`
	// Unique identifier for the Okta user
	Id *string `json:"id,omitempty"`
	// Timestamp of the last synchronization operation. This value is only updated for apps with the `IMPORT_PROFILE_UPDATES` or `PUSH PROFILE_UPDATES` feature.
	LastSync *time.Time `json:"lastSync,omitempty"`
	// Timestamp when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Timestamp when the application user password was last changed
	PasswordChanged NullableTime `json:"passwordChanged,omitempty"`
	// Specifies the default and custom profile properties for a user. Properties that are visible in the Admin Console for an app assignment can also be assigned through the API. Some properties are reference properties that are imported from the target app and can't be configured. See [profile](/openapi/okta-management/management/tag/User/#tag/User/operation/getUser!c=200&path=profile&t=response).
	Profile map[string]interface{} `json:"profile,omitempty"`
	// Indicates if the assignment is direct (`USER`) or by group membership (`GROUP`).
	Scope *string `json:"scope,omitempty"`
	// Status of an application user
	Status *string `json:"status,omitempty"`
	// Timestamp when the application user status was last changed
	StatusChanged *time.Time `json:"statusChanged,omitempty"`
	// The synchronization state for the application user. The application user's `syncState` depends on whether the `PROFILE_MASTERING` feature is enabled for the app.  > **Note:** User provisioning currently must be configured through the Admin Console.
	SyncState *string `json:"syncState,omitempty"`
	// Embedded resources related to the application user using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
	Embedded             map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links                *LinksAppAndUser                  `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppUserAssignRequest AppUserAssignRequest

// NewAppUserAssignRequest instantiates a new AppUserAssignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppUserAssignRequest() *AppUserAssignRequest {
	this := AppUserAssignRequest{}
	return &this
}

// NewAppUserAssignRequestWithDefaults instantiates a new AppUserAssignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppUserAssignRequestWithDefaults() *AppUserAssignRequest {
	this := AppUserAssignRequest{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AppUserAssignRequest) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetCredentials() AppUserCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret AppUserCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetCredentialsOk() (*AppUserCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given AppUserCredentials and assigns it to the Credentials field.
func (o *AppUserAssignRequest) SetCredentials(v AppUserCredentials) {
	o.Credentials = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AppUserAssignRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppUserAssignRequest) SetId(v string) {
	o.Id = &v
}

// GetLastSync returns the LastSync field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetLastSync() time.Time {
	if o == nil || IsNil(o.LastSync) {
		var ret time.Time
		return ret
	}
	return *o.LastSync
}

// GetLastSyncOk returns a tuple with the LastSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetLastSyncOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSync) {
		return nil, false
	}
	return o.LastSync, true
}

// HasLastSync returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasLastSync() bool {
	if o != nil && !IsNil(o.LastSync) {
		return true
	}

	return false
}

// SetLastSync gets a reference to the given time.Time and assigns it to the LastSync field.
func (o *AppUserAssignRequest) SetLastSync(v time.Time) {
	o.LastSync = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AppUserAssignRequest) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetPasswordChanged returns the PasswordChanged field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppUserAssignRequest) GetPasswordChanged() time.Time {
	if o == nil || IsNil(o.PasswordChanged.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PasswordChanged.Get()
}

// GetPasswordChangedOk returns a tuple with the PasswordChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppUserAssignRequest) GetPasswordChangedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordChanged.Get(), o.PasswordChanged.IsSet()
}

// HasPasswordChanged returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasPasswordChanged() bool {
	if o != nil && o.PasswordChanged.IsSet() {
		return true
	}

	return false
}

// SetPasswordChanged gets a reference to the given NullableTime and assigns it to the PasswordChanged field.
func (o *AppUserAssignRequest) SetPasswordChanged(v time.Time) {
	o.PasswordChanged.Set(&v)
}

// SetPasswordChangedNil sets the value for PasswordChanged to be an explicit nil
func (o *AppUserAssignRequest) SetPasswordChangedNil() {
	o.PasswordChanged.Set(nil)
}

// UnsetPasswordChanged ensures that no value is present for PasswordChanged, not even an explicit nil
func (o *AppUserAssignRequest) UnsetPasswordChanged() {
	o.PasswordChanged.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetProfile() map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *AppUserAssignRequest) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AppUserAssignRequest) SetScope(v string) {
	o.Scope = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppUserAssignRequest) SetStatus(v string) {
	o.Status = &v
}

// GetStatusChanged returns the StatusChanged field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetStatusChanged() time.Time {
	if o == nil || IsNil(o.StatusChanged) {
		var ret time.Time
		return ret
	}
	return *o.StatusChanged
}

// GetStatusChangedOk returns a tuple with the StatusChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetStatusChangedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StatusChanged) {
		return nil, false
	}
	return o.StatusChanged, true
}

// HasStatusChanged returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasStatusChanged() bool {
	if o != nil && !IsNil(o.StatusChanged) {
		return true
	}

	return false
}

// SetStatusChanged gets a reference to the given time.Time and assigns it to the StatusChanged field.
func (o *AppUserAssignRequest) SetStatusChanged(v time.Time) {
	o.StatusChanged = &v
}

// GetSyncState returns the SyncState field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetSyncState() string {
	if o == nil || IsNil(o.SyncState) {
		var ret string
		return ret
	}
	return *o.SyncState
}

// GetSyncStateOk returns a tuple with the SyncState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetSyncStateOk() (*string, bool) {
	if o == nil || IsNil(o.SyncState) {
		return nil, false
	}
	return o.SyncState, true
}

// HasSyncState returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasSyncState() bool {
	if o != nil && !IsNil(o.SyncState) {
		return true
	}

	return false
}

// SetSyncState gets a reference to the given string and assigns it to the SyncState field.
func (o *AppUserAssignRequest) SetSyncState(v string) {
	o.SyncState = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *AppUserAssignRequest) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppUserAssignRequest) GetLinks() LinksAppAndUser {
	if o == nil || IsNil(o.Links) {
		var ret LinksAppAndUser
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppUserAssignRequest) GetLinksOk() (*LinksAppAndUser, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppUserAssignRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksAppAndUser and assigns it to the Links field.
func (o *AppUserAssignRequest) SetLinks(v LinksAppAndUser) {
	o.Links = &v
}

func (o AppUserAssignRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppUserAssignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastSync) {
		toSerialize["lastSync"] = o.LastSync
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.PasswordChanged.IsSet() {
		toSerialize["passwordChanged"] = o.PasswordChanged.Get()
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusChanged) {
		toSerialize["statusChanged"] = o.StatusChanged
	}
	if !IsNil(o.SyncState) {
		toSerialize["syncState"] = o.SyncState
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppUserAssignRequest) UnmarshalJSON(data []byte) (err error) {
	varAppUserAssignRequest := _AppUserAssignRequest{}

	err = json.Unmarshal(data, &varAppUserAssignRequest)

	if err != nil {
		return err
	}

	*o = AppUserAssignRequest(varAppUserAssignRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
	}

	return err
}

type NullableAppUserAssignRequest struct {
	value *AppUserAssignRequest
	isSet bool
}

func (v NullableAppUserAssignRequest) Get() *AppUserAssignRequest {
	return v.value
}

func (v *NullableAppUserAssignRequest) Set(val *AppUserAssignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppUserAssignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppUserAssignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppUserAssignRequest(val *AppUserAssignRequest) *NullableAppUserAssignRequest {
	return &NullableAppUserAssignRequest{value: val, isSet: true}
}

func (v NullableAppUserAssignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppUserAssignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
