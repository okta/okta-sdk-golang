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
	"time"
)

// checks if the AuthenticatorEnrollment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorEnrollment{}

// AuthenticatorEnrollment struct for AuthenticatorEnrollment
type AuthenticatorEnrollment struct {
	// Timestamp when the authenticator enrollment was created
	Created *time.Time `json:"created,omitempty"`
	// The unique identifier of the authenticator enrollment
	Id *string `json:"id,omitempty"`
	// A human-readable string that identifies the authenticator
	Key *string `json:"key,omitempty"`
	// Timestamp when the authenticator enrollment was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The authenticator display name
	Name    *string               `json:"name,omitempty"`
	Profile *AuthenticatorProfile `json:"profile,omitempty"`
	// Status of the enrollment
	Status *string `json:"status,omitempty"`
	// The type of authenticator
	Type                 *string                       `json:"type,omitempty"`
	Links                *AuthenticatorEnrollmentLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorEnrollment AuthenticatorEnrollment

// NewAuthenticatorEnrollment instantiates a new AuthenticatorEnrollment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorEnrollment() *AuthenticatorEnrollment {
	this := AuthenticatorEnrollment{}
	return &this
}

// NewAuthenticatorEnrollmentWithDefaults instantiates a new AuthenticatorEnrollment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorEnrollmentWithDefaults() *AuthenticatorEnrollment {
	this := AuthenticatorEnrollment{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AuthenticatorEnrollment) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollment) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AuthenticatorEnrollment) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AuthenticatorEnrollment) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticatorEnrollment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticatorEnrollment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticatorEnrollment) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AuthenticatorEnrollment) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollment) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AuthenticatorEnrollment) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AuthenticatorEnrollment) SetKey(v string) {
	o.Key = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AuthenticatorEnrollment) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AuthenticatorEnrollment) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AuthenticatorEnrollment) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthenticatorEnrollment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthenticatorEnrollment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthenticatorEnrollment) SetName(v string) {
	o.Name = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *AuthenticatorEnrollment) GetProfile() AuthenticatorProfile {
	if o == nil || IsNil(o.Profile) {
		var ret AuthenticatorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollment) GetProfileOk() (*AuthenticatorProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *AuthenticatorEnrollment) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given AuthenticatorProfile and assigns it to the Profile field.
func (o *AuthenticatorEnrollment) SetProfile(v AuthenticatorProfile) {
	o.Profile = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthenticatorEnrollment) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollment) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthenticatorEnrollment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthenticatorEnrollment) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthenticatorEnrollment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthenticatorEnrollment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthenticatorEnrollment) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthenticatorEnrollment) GetLinks() AuthenticatorEnrollmentLinks {
	if o == nil || IsNil(o.Links) {
		var ret AuthenticatorEnrollmentLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorEnrollment) GetLinksOk() (*AuthenticatorEnrollmentLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthenticatorEnrollment) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AuthenticatorEnrollmentLinks and assigns it to the Links field.
func (o *AuthenticatorEnrollment) SetLinks(v AuthenticatorEnrollmentLinks) {
	o.Links = &v
}

func (o AuthenticatorEnrollment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorEnrollment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorEnrollment) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorEnrollment := _AuthenticatorEnrollment{}

	err = json.Unmarshal(data, &varAuthenticatorEnrollment)

	if err != nil {
		return err
	}

	*o = AuthenticatorEnrollment(varAuthenticatorEnrollment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorEnrollment struct {
	value *AuthenticatorEnrollment
	isSet bool
}

func (v NullableAuthenticatorEnrollment) Get() *AuthenticatorEnrollment {
	return v.value
}

func (v *NullableAuthenticatorEnrollment) Set(val *AuthenticatorEnrollment) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorEnrollment) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorEnrollment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorEnrollment(val *AuthenticatorEnrollment) *NullableAuthenticatorEnrollment {
	return &NullableAuthenticatorEnrollment{value: val, isSet: true}
}

func (v NullableAuthenticatorEnrollment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorEnrollment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
