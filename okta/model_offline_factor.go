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
	"time"
)

// checks if the OfflineFactor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfflineFactor{}

// OfflineFactor A factor that's enrolled for offline authentication
type OfflineFactor struct {
	// Human-readable name of the offline factor
	DisplayName *string `json:"displayName,omitempty"`
	// Timestamp when the offline factor was enrolled
	EnrollDate *time.Time `json:"enrollDate,omitempty"`
	// Unique identifier of the offline factor
	Id      *string               `json:"id,omitempty"`
	Profile *OfflineFactorProfile `json:"profile,omitempty"`
	// The type of offline factor. `TOTP` is the only supported offline factor for macOS enrollments, currently.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OfflineFactor OfflineFactor

// NewOfflineFactor instantiates a new OfflineFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfflineFactor() *OfflineFactor {
	this := OfflineFactor{}
	return &this
}

// NewOfflineFactorWithDefaults instantiates a new OfflineFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfflineFactorWithDefaults() *OfflineFactor {
	this := OfflineFactor{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OfflineFactor) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineFactor) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OfflineFactor) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OfflineFactor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnrollDate returns the EnrollDate field value if set, zero value otherwise.
func (o *OfflineFactor) GetEnrollDate() time.Time {
	if o == nil || IsNil(o.EnrollDate) {
		var ret time.Time
		return ret
	}
	return *o.EnrollDate
}

// GetEnrollDateOk returns a tuple with the EnrollDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineFactor) GetEnrollDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EnrollDate) {
		return nil, false
	}
	return o.EnrollDate, true
}

// HasEnrollDate returns a boolean if a field has been set.
func (o *OfflineFactor) HasEnrollDate() bool {
	if o != nil && !IsNil(o.EnrollDate) {
		return true
	}

	return false
}

// SetEnrollDate gets a reference to the given time.Time and assigns it to the EnrollDate field.
func (o *OfflineFactor) SetEnrollDate(v time.Time) {
	o.EnrollDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OfflineFactor) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineFactor) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OfflineFactor) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OfflineFactor) SetId(v string) {
	o.Id = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *OfflineFactor) GetProfile() OfflineFactorProfile {
	if o == nil || IsNil(o.Profile) {
		var ret OfflineFactorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineFactor) GetProfileOk() (*OfflineFactorProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *OfflineFactor) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given OfflineFactorProfile and assigns it to the Profile field.
func (o *OfflineFactor) SetProfile(v OfflineFactorProfile) {
	o.Profile = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OfflineFactor) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineFactor) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OfflineFactor) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OfflineFactor) SetType(v string) {
	o.Type = &v
}

func (o OfflineFactor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfflineFactor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.EnrollDate) {
		toSerialize["enrollDate"] = o.EnrollDate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OfflineFactor) UnmarshalJSON(data []byte) (err error) {
	varOfflineFactor := _OfflineFactor{}

	err = json.Unmarshal(data, &varOfflineFactor)

	if err != nil {
		return err
	}

	*o = OfflineFactor(varOfflineFactor)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "enrollDate")
		delete(additionalProperties, "id")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOfflineFactor struct {
	value *OfflineFactor
	isSet bool
}

func (v NullableOfflineFactor) Get() *OfflineFactor {
	return v.value
}

func (v *NullableOfflineFactor) Set(val *OfflineFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableOfflineFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableOfflineFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfflineFactor(val *OfflineFactor) *NullableOfflineFactor {
	return &NullableOfflineFactor{value: val, isSet: true}
}

func (v NullableOfflineFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfflineFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
