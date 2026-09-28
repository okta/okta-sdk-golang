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

// checks if the AccountLinkedEnrollment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountLinkedEnrollment{}

// AccountLinkedEnrollment An enrollment linked to an OS account
type AccountLinkedEnrollment struct {
	// Timestamp when the enrollment was created
	Created *time.Time `json:"created,omitempty"`
	// Unique identifier of the enrollment
	Id *string `json:"id,omitempty"`
	// Timestamp when the enrollment was last updated
	LastUpdated *time.Time                      `json:"lastUpdated,omitempty"`
	Profile     *AccountLinkedEnrollmentProfile `json:"profile,omitempty"`
	// Status of the enrollment. Possible values depend on `type`. For `platform_sso` - `ACTIVE`, `SUSPENDED`, `REVOKED`. For `desktop_mfa` - `ACTIVE`, `DELETED`.
	Status *string `json:"status,omitempty"`
	// Type of the linked enrollment
	Type                 *string                          `json:"type,omitempty"`
	Embedded             *AccountLinkedEnrollmentEmbedded `json:"_embedded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountLinkedEnrollment AccountLinkedEnrollment

// NewAccountLinkedEnrollment instantiates a new AccountLinkedEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountLinkedEnrollment() *AccountLinkedEnrollment {
	this := AccountLinkedEnrollment{}
	return &this
}

// NewAccountLinkedEnrollmentWithDefaults instantiates a new AccountLinkedEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountLinkedEnrollmentWithDefaults() *AccountLinkedEnrollment {
	this := AccountLinkedEnrollment{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AccountLinkedEnrollment) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollment) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AccountLinkedEnrollment) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AccountLinkedEnrollment) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountLinkedEnrollment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountLinkedEnrollment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountLinkedEnrollment) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AccountLinkedEnrollment) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AccountLinkedEnrollment) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AccountLinkedEnrollment) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AccountLinkedEnrollment) GetProfile() AccountLinkedEnrollmentProfile {
	if o == nil || IsNil(o.Profile) {
		var ret AccountLinkedEnrollmentProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollment) GetProfileOk() (*AccountLinkedEnrollmentProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AccountLinkedEnrollment) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given AccountLinkedEnrollmentProfile and assigns it to the Profile field.
func (o *AccountLinkedEnrollment) SetProfile(v AccountLinkedEnrollmentProfile) {
	o.Profile = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountLinkedEnrollment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountLinkedEnrollment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AccountLinkedEnrollment) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountLinkedEnrollment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountLinkedEnrollment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountLinkedEnrollment) SetType(v string) {
	o.Type = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *AccountLinkedEnrollment) GetEmbedded() AccountLinkedEnrollmentEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret AccountLinkedEnrollmentEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountLinkedEnrollment) GetEmbeddedOk() (*AccountLinkedEnrollmentEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *AccountLinkedEnrollment) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given AccountLinkedEnrollmentEmbedded and assigns it to the Embedded field.
func (o *AccountLinkedEnrollment) SetEmbedded(v AccountLinkedEnrollmentEmbedded) {
	o.Embedded = &v
}

func (o AccountLinkedEnrollment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountLinkedEnrollment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountLinkedEnrollment) UnmarshalJSON(data []byte) (err error) {
	varAccountLinkedEnrollment := _AccountLinkedEnrollment{}

	err = json.Unmarshal(data, &varAccountLinkedEnrollment)

	if err != nil {
		return err
	}

	*o = AccountLinkedEnrollment(varAccountLinkedEnrollment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_embedded")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountLinkedEnrollment struct {
	value *AccountLinkedEnrollment
	isSet bool
}

func (v NullableAccountLinkedEnrollment) Get() *AccountLinkedEnrollment {
	return v.value
}

func (v *NullableAccountLinkedEnrollment) Set(val *AccountLinkedEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountLinkedEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountLinkedEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountLinkedEnrollment(val *AccountLinkedEnrollment) *NullableAccountLinkedEnrollment {
	return &NullableAccountLinkedEnrollment{value: val, isSet: true}
}

func (v NullableAccountLinkedEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountLinkedEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
