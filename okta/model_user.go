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

// User struct for User
type User struct {
	// The timestamp when the user status transitioned to `ACTIVE`
	Activated NullableTime `json:"activated,omitempty"`
	// The timestamp when the user was created
	Created *time.Time `json:"created,omitempty"`
	Credentials *UserCredentials `json:"credentials,omitempty"`
	// The unique key for the user
	Id *string `json:"id,omitempty"`
	// The timestamp of the last login
	LastLogin NullableTime `json:"lastLogin,omitempty"`
	// The timestamp when the user was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The timestamp when the user's password was last updated
	PasswordChanged NullableTime `json:"passwordChanged,omitempty"`
	Profile *UserProfile `json:"profile,omitempty"`
	// <div class=\"x-lifecycle-container\"><x-lifecycle class=\"ea\"></x-lifecycle></div>The ID of the Realm in which the user is residing
	RealmId *string `json:"realmId,omitempty"`
	// The current status of the user
	Status *string `json:"status,omitempty"`
	// The timestamp when the status of the user last changed
	StatusChanged NullableTime `json:"statusChanged,omitempty"`
	// The target status of an in-progress asynchronous status transition. This property is only returned if the user's state is transitioning.
	TransitioningToStatus NullableString `json:"transitioningToStatus,omitempty"`
	Type *UserType `json:"type,omitempty"`
	// If specified, includes embedded resources related to the user
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links *UserLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetActivated returns the Activated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetActivated() time.Time {
	if o == nil || o.Activated.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Activated.Get()
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetActivatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Activated.Get(), o.Activated.IsSet()
}

// HasActivated returns a boolean if a field has been set.
func (o *User) HasActivated() bool {
	if o != nil && o.Activated.IsSet() {
		return true
	}

	return false
}

// SetActivated gets a reference to the given NullableTime and assigns it to the Activated field.
func (o *User) SetActivated(v time.Time) {
	o.Activated.Set(&v)
}
// SetActivatedNil sets the value for Activated to be an explicit nil
func (o *User) SetActivatedNil() {
	o.Activated.Set(nil)
}

// UnsetActivated ensures that no value is present for Activated, not even an explicit nil
func (o *User) UnsetActivated() {
	o.Activated.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *User) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *User) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *User) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *User) GetCredentials() UserCredentials {
	if o == nil || o.Credentials == nil {
		var ret UserCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCredentialsOk() (*UserCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *User) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given UserCredentials and assigns it to the Credentials field.
func (o *User) SetCredentials(v UserCredentials) {
	o.Credentials = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLastLogin() time.Time {
	if o == nil || o.LastLogin.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastLogin.Get()
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLastLoginOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLogin.Get(), o.LastLogin.IsSet()
}

// HasLastLogin returns a boolean if a field has been set.
func (o *User) HasLastLogin() bool {
	if o != nil && o.LastLogin.IsSet() {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given NullableTime and assigns it to the LastLogin field.
func (o *User) SetLastLogin(v time.Time) {
	o.LastLogin.Set(&v)
}
// SetLastLoginNil sets the value for LastLogin to be an explicit nil
func (o *User) SetLastLoginNil() {
	o.LastLogin.Set(nil)
}

// UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
func (o *User) UnsetLastLogin() {
	o.LastLogin.Unset()
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *User) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *User) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *User) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetPasswordChanged returns the PasswordChanged field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPasswordChanged() time.Time {
	if o == nil || o.PasswordChanged.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PasswordChanged.Get()
}

// GetPasswordChangedOk returns a tuple with the PasswordChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPasswordChangedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordChanged.Get(), o.PasswordChanged.IsSet()
}

// HasPasswordChanged returns a boolean if a field has been set.
func (o *User) HasPasswordChanged() bool {
	if o != nil && o.PasswordChanged.IsSet() {
		return true
	}

	return false
}

// SetPasswordChanged gets a reference to the given NullableTime and assigns it to the PasswordChanged field.
func (o *User) SetPasswordChanged(v time.Time) {
	o.PasswordChanged.Set(&v)
}
// SetPasswordChangedNil sets the value for PasswordChanged to be an explicit nil
func (o *User) SetPasswordChangedNil() {
	o.PasswordChanged.Set(nil)
}

// UnsetPasswordChanged ensures that no value is present for PasswordChanged, not even an explicit nil
func (o *User) UnsetPasswordChanged() {
	o.PasswordChanged.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *User) GetProfile() UserProfile {
	if o == nil || o.Profile == nil {
		var ret UserProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetProfileOk() (*UserProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *User) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserProfile and assigns it to the Profile field.
func (o *User) SetProfile(v UserProfile) {
	o.Profile = &v
}

// GetRealmId returns the RealmId field value if set, zero value otherwise.
func (o *User) GetRealmId() string {
	if o == nil || o.RealmId == nil {
		var ret string
		return ret
	}
	return *o.RealmId
}

// GetRealmIdOk returns a tuple with the RealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRealmIdOk() (*string, bool) {
	if o == nil || o.RealmId == nil {
		return nil, false
	}
	return o.RealmId, true
}

// HasRealmId returns a boolean if a field has been set.
func (o *User) HasRealmId() bool {
	if o != nil && o.RealmId != nil {
		return true
	}

	return false
}

// SetRealmId gets a reference to the given string and assigns it to the RealmId field.
func (o *User) SetRealmId(v string) {
	o.RealmId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *User) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *User) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *User) SetStatus(v string) {
	o.Status = &v
}

// GetStatusChanged returns the StatusChanged field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetStatusChanged() time.Time {
	if o == nil || o.StatusChanged.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusChanged.Get()
}

// GetStatusChangedOk returns a tuple with the StatusChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetStatusChangedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusChanged.Get(), o.StatusChanged.IsSet()
}

// HasStatusChanged returns a boolean if a field has been set.
func (o *User) HasStatusChanged() bool {
	if o != nil && o.StatusChanged.IsSet() {
		return true
	}

	return false
}

// SetStatusChanged gets a reference to the given NullableTime and assigns it to the StatusChanged field.
func (o *User) SetStatusChanged(v time.Time) {
	o.StatusChanged.Set(&v)
}
// SetStatusChangedNil sets the value for StatusChanged to be an explicit nil
func (o *User) SetStatusChangedNil() {
	o.StatusChanged.Set(nil)
}

// UnsetStatusChanged ensures that no value is present for StatusChanged, not even an explicit nil
func (o *User) UnsetStatusChanged() {
	o.StatusChanged.Unset()
}

// GetTransitioningToStatus returns the TransitioningToStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetTransitioningToStatus() string {
	if o == nil || o.TransitioningToStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.TransitioningToStatus.Get()
}

// GetTransitioningToStatusOk returns a tuple with the TransitioningToStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetTransitioningToStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransitioningToStatus.Get(), o.TransitioningToStatus.IsSet()
}

// HasTransitioningToStatus returns a boolean if a field has been set.
func (o *User) HasTransitioningToStatus() bool {
	if o != nil && o.TransitioningToStatus.IsSet() {
		return true
	}

	return false
}

// SetTransitioningToStatus gets a reference to the given NullableString and assigns it to the TransitioningToStatus field.
func (o *User) SetTransitioningToStatus(v string) {
	o.TransitioningToStatus.Set(&v)
}
// SetTransitioningToStatusNil sets the value for TransitioningToStatus to be an explicit nil
func (o *User) SetTransitioningToStatusNil() {
	o.TransitioningToStatus.Set(nil)
}

// UnsetTransitioningToStatus ensures that no value is present for TransitioningToStatus, not even an explicit nil
func (o *User) UnsetTransitioningToStatus() {
	o.TransitioningToStatus.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *User) GetType() UserType {
	if o == nil || o.Type == nil {
		var ret UserType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTypeOk() (*UserType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *User) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given UserType and assigns it to the Type field.
func (o *User) SetType(v UserType) {
	o.Type = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *User) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || o.Embedded == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *User) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *User) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *User) GetLinks() UserLinks {
	if o == nil || o.Links == nil {
		var ret UserLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLinksOk() (*UserLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *User) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserLinks and assigns it to the Links field.
func (o *User) SetLinks(v UserLinks) {
	o.Links = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activated.IsSet() {
		toSerialize["activated"] = o.Activated.Get()
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastLogin.IsSet() {
		toSerialize["lastLogin"] = o.LastLogin.Get()
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
	if o.RealmId != nil {
		toSerialize["realmId"] = o.RealmId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusChanged.IsSet() {
		toSerialize["statusChanged"] = o.StatusChanged.Get()
	}
	if o.TransitioningToStatus.IsSet() {
		toSerialize["transitioningToStatus"] = o.TransitioningToStatus.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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

func (o *User) UnmarshalJSON(bytes []byte) (err error) {
	varUser := _User{}

	err = json.Unmarshal(bytes, &varUser)
	if err == nil {
		*o = User(varUser)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "activated")
		delete(additionalProperties, "created")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastLogin")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "passwordChanged")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "realmId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusChanged")
		delete(additionalProperties, "transitioningToStatus")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

