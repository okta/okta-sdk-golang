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
	"fmt"
	"time"
)

// checks if the DeviceEnrollmentAuthenticator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceEnrollmentAuthenticator{}

// DeviceEnrollmentAuthenticator struct for DeviceEnrollmentAuthenticator
type DeviceEnrollmentAuthenticator struct {
	// Timestamp when the enrollment authenticator was created
	Created time.Time `json:"created"`
	// The unique identifier of the enrollment authenticator
	Id string `json:"id"`
	// The authenticator key. For example: `okta_verify`.
	Key string `json:"key"`
	// Timestamp when the enrollment authenticator was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// Timestamp when the enrollment authenticator was last verified
	LastVerified *time.Time `json:"lastVerified,omitempty"`
	// The authenticator display name
	Name string `json:"name"`
	// A user-friendly name for the enrollment authenticator
	Nickname *string                               `json:"nickname,omitempty"`
	Profile  *DeviceEnrollmentAuthenticatorProfile `json:"profile,omitempty"`
	// Status of the enrollment
	Status string `json:"status"`
	// The type of authenticator
	Type                 string                             `json:"type"`
	Links                DeviceEnrollmentAuthenticatorLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _DeviceEnrollmentAuthenticator DeviceEnrollmentAuthenticator

// NewDeviceEnrollmentAuthenticator instantiates a new DeviceEnrollmentAuthenticator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceEnrollmentAuthenticator(created time.Time, id string, key string, lastUpdated time.Time, name string, status string, type_ string, links DeviceEnrollmentAuthenticatorLinks) *DeviceEnrollmentAuthenticator {
	this := DeviceEnrollmentAuthenticator{}
	this.Created = created
	this.Id = id
	this.Key = key
	this.LastUpdated = lastUpdated
	this.Name = name
	this.Status = status
	this.Type = type_
	this.Links = links
	return &this
}

// NewDeviceEnrollmentAuthenticatorWithDefaults instantiates a new DeviceEnrollmentAuthenticator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceEnrollmentAuthenticatorWithDefaults() *DeviceEnrollmentAuthenticator {
	this := DeviceEnrollmentAuthenticator{}
	return &this
}

// GetCreated returns the Created field value
func (o *DeviceEnrollmentAuthenticator) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *DeviceEnrollmentAuthenticator) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *DeviceEnrollmentAuthenticator) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceEnrollmentAuthenticator) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *DeviceEnrollmentAuthenticator) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DeviceEnrollmentAuthenticator) SetKey(v string) {
	o.Key = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *DeviceEnrollmentAuthenticator) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *DeviceEnrollmentAuthenticator) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastVerified returns the LastVerified field value if set, zero value otherwise.
func (o *DeviceEnrollmentAuthenticator) GetLastVerified() time.Time {
	if o == nil || IsNil(o.LastVerified) {
		var ret time.Time
		return ret
	}
	return *o.LastVerified
}

// GetLastVerifiedOk returns a tuple with the LastVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetLastVerifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastVerified) {
		return nil, false
	}
	return o.LastVerified, true
}

// HasLastVerified returns a boolean if a field has been set.
func (o *DeviceEnrollmentAuthenticator) HasLastVerified() bool {
	if o != nil && !IsNil(o.LastVerified) {
		return true
	}

	return false
}

// SetLastVerified gets a reference to the given time.Time and assigns it to the LastVerified field.
func (o *DeviceEnrollmentAuthenticator) SetLastVerified(v time.Time) {
	o.LastVerified = &v
}

// GetName returns the Name field value
func (o *DeviceEnrollmentAuthenticator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceEnrollmentAuthenticator) SetName(v string) {
	o.Name = v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *DeviceEnrollmentAuthenticator) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *DeviceEnrollmentAuthenticator) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *DeviceEnrollmentAuthenticator) SetNickname(v string) {
	o.Nickname = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *DeviceEnrollmentAuthenticator) GetProfile() DeviceEnrollmentAuthenticatorProfile {
	if o == nil || IsNil(o.Profile) {
		var ret DeviceEnrollmentAuthenticatorProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetProfileOk() (*DeviceEnrollmentAuthenticatorProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *DeviceEnrollmentAuthenticator) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given DeviceEnrollmentAuthenticatorProfile and assigns it to the Profile field.
func (o *DeviceEnrollmentAuthenticator) SetProfile(v DeviceEnrollmentAuthenticatorProfile) {
	o.Profile = &v
}

// GetStatus returns the Status field value
func (o *DeviceEnrollmentAuthenticator) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeviceEnrollmentAuthenticator) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *DeviceEnrollmentAuthenticator) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeviceEnrollmentAuthenticator) SetType(v string) {
	o.Type = v
}

// GetLinks returns the Links field value
func (o *DeviceEnrollmentAuthenticator) GetLinks() DeviceEnrollmentAuthenticatorLinks {
	if o == nil {
		var ret DeviceEnrollmentAuthenticatorLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DeviceEnrollmentAuthenticator) GetLinksOk() (*DeviceEnrollmentAuthenticatorLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DeviceEnrollmentAuthenticator) SetLinks(v DeviceEnrollmentAuthenticatorLinks) {
	o.Links = v
}

func (o DeviceEnrollmentAuthenticator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceEnrollmentAuthenticator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["lastUpdated"] = o.LastUpdated
	if !IsNil(o.LastVerified) {
		toSerialize["lastVerified"] = o.LastVerified
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceEnrollmentAuthenticator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"id",
		"key",
		"lastUpdated",
		"name",
		"status",
		"type",
		"_links",
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

	varDeviceEnrollmentAuthenticator := _DeviceEnrollmentAuthenticator{}

	err = json.Unmarshal(data, &varDeviceEnrollmentAuthenticator)

	if err != nil {
		return err
	}

	*o = DeviceEnrollmentAuthenticator(varDeviceEnrollmentAuthenticator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "key")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastVerified")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceEnrollmentAuthenticator struct {
	value *DeviceEnrollmentAuthenticator
	isSet bool
}

func (v NullableDeviceEnrollmentAuthenticator) Get() *DeviceEnrollmentAuthenticator {
	return v.value
}

func (v *NullableDeviceEnrollmentAuthenticator) Set(val *DeviceEnrollmentAuthenticator) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceEnrollmentAuthenticator) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceEnrollmentAuthenticator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceEnrollmentAuthenticator(val *DeviceEnrollmentAuthenticator) *NullableDeviceEnrollmentAuthenticator {
	return &NullableDeviceEnrollmentAuthenticator{value: val, isSet: true}
}

func (v NullableDeviceEnrollmentAuthenticator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceEnrollmentAuthenticator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
