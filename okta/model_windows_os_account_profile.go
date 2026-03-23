/*
Okta Admin Management API

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

// checks if the WindowsOSAccountProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WindowsOSAccountProfile{}

// WindowsOSAccountProfile struct for WindowsOSAccountProfile
type WindowsOSAccountProfile struct {
	// Active Directory join status
	DirectoryJoinStatus *string `json:"directoryJoinStatus,omitempty"`
	// Domain\\username format (down-level logon name)
	DownLevelUsername *string `json:"downLevelUsername,omitempty"`
	// Full name of the account user
	FullName *string `json:"fullName,omitempty"`
	// Globally Unique Identifier for the account
	GUID *string `json:"GUID,omitempty"`
	// Fully qualified username
	QualifiedUsername *string `json:"qualifiedUsername,omitempty"`
	// Windows Security Identifier (SID)
	SecurityId *string `json:"securityId,omitempty"`
	// User principal name
	Upn                  *string `json:"upn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WindowsOSAccountProfile WindowsOSAccountProfile

// NewWindowsOSAccountProfile instantiates a new WindowsOSAccountProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWindowsOSAccountProfile() *WindowsOSAccountProfile {
	this := WindowsOSAccountProfile{}
	return &this
}

// NewWindowsOSAccountProfileWithDefaults instantiates a new WindowsOSAccountProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWindowsOSAccountProfileWithDefaults() *WindowsOSAccountProfile {
	this := WindowsOSAccountProfile{}
	return &this
}

// GetDirectoryJoinStatus returns the DirectoryJoinStatus field value if set, zero value otherwise.
func (o *WindowsOSAccountProfile) GetDirectoryJoinStatus() string {
	if o == nil || IsNil(o.DirectoryJoinStatus) {
		var ret string
		return ret
	}
	return *o.DirectoryJoinStatus
}

// GetDirectoryJoinStatusOk returns a tuple with the DirectoryJoinStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsOSAccountProfile) GetDirectoryJoinStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DirectoryJoinStatus) {
		return nil, false
	}
	return o.DirectoryJoinStatus, true
}

// HasDirectoryJoinStatus returns a boolean if a field has been set.
func (o *WindowsOSAccountProfile) HasDirectoryJoinStatus() bool {
	if o != nil && !IsNil(o.DirectoryJoinStatus) {
		return true
	}

	return false
}

// SetDirectoryJoinStatus gets a reference to the given string and assigns it to the DirectoryJoinStatus field.
func (o *WindowsOSAccountProfile) SetDirectoryJoinStatus(v string) {
	o.DirectoryJoinStatus = &v
}

// GetDownLevelUsername returns the DownLevelUsername field value if set, zero value otherwise.
func (o *WindowsOSAccountProfile) GetDownLevelUsername() string {
	if o == nil || IsNil(o.DownLevelUsername) {
		var ret string
		return ret
	}
	return *o.DownLevelUsername
}

// GetDownLevelUsernameOk returns a tuple with the DownLevelUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsOSAccountProfile) GetDownLevelUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.DownLevelUsername) {
		return nil, false
	}
	return o.DownLevelUsername, true
}

// HasDownLevelUsername returns a boolean if a field has been set.
func (o *WindowsOSAccountProfile) HasDownLevelUsername() bool {
	if o != nil && !IsNil(o.DownLevelUsername) {
		return true
	}

	return false
}

// SetDownLevelUsername gets a reference to the given string and assigns it to the DownLevelUsername field.
func (o *WindowsOSAccountProfile) SetDownLevelUsername(v string) {
	o.DownLevelUsername = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *WindowsOSAccountProfile) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsOSAccountProfile) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *WindowsOSAccountProfile) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *WindowsOSAccountProfile) SetFullName(v string) {
	o.FullName = &v
}

// GetGUID returns the GUID field value if set, zero value otherwise.
func (o *WindowsOSAccountProfile) GetGUID() string {
	if o == nil || IsNil(o.GUID) {
		var ret string
		return ret
	}
	return *o.GUID
}

// GetGUIDOk returns a tuple with the GUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsOSAccountProfile) GetGUIDOk() (*string, bool) {
	if o == nil || IsNil(o.GUID) {
		return nil, false
	}
	return o.GUID, true
}

// HasGUID returns a boolean if a field has been set.
func (o *WindowsOSAccountProfile) HasGUID() bool {
	if o != nil && !IsNil(o.GUID) {
		return true
	}

	return false
}

// SetGUID gets a reference to the given string and assigns it to the GUID field.
func (o *WindowsOSAccountProfile) SetGUID(v string) {
	o.GUID = &v
}

// GetQualifiedUsername returns the QualifiedUsername field value if set, zero value otherwise.
func (o *WindowsOSAccountProfile) GetQualifiedUsername() string {
	if o == nil || IsNil(o.QualifiedUsername) {
		var ret string
		return ret
	}
	return *o.QualifiedUsername
}

// GetQualifiedUsernameOk returns a tuple with the QualifiedUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsOSAccountProfile) GetQualifiedUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.QualifiedUsername) {
		return nil, false
	}
	return o.QualifiedUsername, true
}

// HasQualifiedUsername returns a boolean if a field has been set.
func (o *WindowsOSAccountProfile) HasQualifiedUsername() bool {
	if o != nil && !IsNil(o.QualifiedUsername) {
		return true
	}

	return false
}

// SetQualifiedUsername gets a reference to the given string and assigns it to the QualifiedUsername field.
func (o *WindowsOSAccountProfile) SetQualifiedUsername(v string) {
	o.QualifiedUsername = &v
}

// GetSecurityId returns the SecurityId field value if set, zero value otherwise.
func (o *WindowsOSAccountProfile) GetSecurityId() string {
	if o == nil || IsNil(o.SecurityId) {
		var ret string
		return ret
	}
	return *o.SecurityId
}

// GetSecurityIdOk returns a tuple with the SecurityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsOSAccountProfile) GetSecurityIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityId) {
		return nil, false
	}
	return o.SecurityId, true
}

// HasSecurityId returns a boolean if a field has been set.
func (o *WindowsOSAccountProfile) HasSecurityId() bool {
	if o != nil && !IsNil(o.SecurityId) {
		return true
	}

	return false
}

// SetSecurityId gets a reference to the given string and assigns it to the SecurityId field.
func (o *WindowsOSAccountProfile) SetSecurityId(v string) {
	o.SecurityId = &v
}

// GetUpn returns the Upn field value if set, zero value otherwise.
func (o *WindowsOSAccountProfile) GetUpn() string {
	if o == nil || IsNil(o.Upn) {
		var ret string
		return ret
	}
	return *o.Upn
}

// GetUpnOk returns a tuple with the Upn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WindowsOSAccountProfile) GetUpnOk() (*string, bool) {
	if o == nil || IsNil(o.Upn) {
		return nil, false
	}
	return o.Upn, true
}

// HasUpn returns a boolean if a field has been set.
func (o *WindowsOSAccountProfile) HasUpn() bool {
	if o != nil && !IsNil(o.Upn) {
		return true
	}

	return false
}

// SetUpn gets a reference to the given string and assigns it to the Upn field.
func (o *WindowsOSAccountProfile) SetUpn(v string) {
	o.Upn = &v
}

func (o WindowsOSAccountProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WindowsOSAccountProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DirectoryJoinStatus) {
		toSerialize["directoryJoinStatus"] = o.DirectoryJoinStatus
	}
	if !IsNil(o.DownLevelUsername) {
		toSerialize["downLevelUsername"] = o.DownLevelUsername
	}
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.GUID) {
		toSerialize["GUID"] = o.GUID
	}
	if !IsNil(o.QualifiedUsername) {
		toSerialize["qualifiedUsername"] = o.QualifiedUsername
	}
	if !IsNil(o.SecurityId) {
		toSerialize["securityId"] = o.SecurityId
	}
	if !IsNil(o.Upn) {
		toSerialize["upn"] = o.Upn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WindowsOSAccountProfile) UnmarshalJSON(data []byte) (err error) {
	varWindowsOSAccountProfile := _WindowsOSAccountProfile{}

	err = json.Unmarshal(data, &varWindowsOSAccountProfile)

	if err != nil {
		return err
	}

	*o = WindowsOSAccountProfile(varWindowsOSAccountProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "directoryJoinStatus")
		delete(additionalProperties, "downLevelUsername")
		delete(additionalProperties, "fullName")
		delete(additionalProperties, "GUID")
		delete(additionalProperties, "qualifiedUsername")
		delete(additionalProperties, "securityId")
		delete(additionalProperties, "upn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWindowsOSAccountProfile struct {
	value *WindowsOSAccountProfile
	isSet bool
}

func (v NullableWindowsOSAccountProfile) Get() *WindowsOSAccountProfile {
	return v.value
}

func (v *NullableWindowsOSAccountProfile) Set(val *WindowsOSAccountProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableWindowsOSAccountProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableWindowsOSAccountProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWindowsOSAccountProfile(val *WindowsOSAccountProfile) *NullableWindowsOSAccountProfile {
	return &NullableWindowsOSAccountProfile{value: val, isSet: true}
}

func (v NullableWindowsOSAccountProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWindowsOSAccountProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
