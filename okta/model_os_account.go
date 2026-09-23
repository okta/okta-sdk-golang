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

// checks if the OSAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OSAccount{}

// OSAccount struct for OSAccount
type OSAccount struct {
	// Timestamp when the OS account was created
	Created time.Time `json:"created"`
	// Unique identifier of the device this OS account belongs to
	DeviceId string `json:"deviceId"`
	// Unique identifier for the OS account
	Id string `json:"id"`
	// Timestamp when the OS account was last seen
	LastSeenAt NullableTime `json:"lastSeenAt,omitempty"`
	// Timestamp when the OS account was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// OS platform for OS accounts (desktop platforms only)
	Platform            string                `json:"platform"`
	ResourceAlternateId NullableString        `json:"resourceAlternateId,omitempty"`
	ResourceDisplayName *OSAccountDisplayName `json:"resourceDisplayName,omitempty"`
	// Alternate key for the `id`
	ResourceId   *string `json:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	// Status of the OS account
	Status               string             `json:"status"`
	Embedded             *OSAccountEmbedded `json:"_embedded,omitempty"`
	Links                OSAccountLinks     `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _OSAccount OSAccount

// NewOSAccount instantiates a new OSAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOSAccount(created time.Time, deviceId string, id string, lastUpdated time.Time, platform string, status string, links OSAccountLinks) *OSAccount {
	this := OSAccount{}
	this.Created = created
	this.DeviceId = deviceId
	this.Id = id
	this.LastUpdated = lastUpdated
	this.Platform = platform
	this.Status = status
	this.Links = links
	return &this
}

// NewOSAccountWithDefaults instantiates a new OSAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOSAccountWithDefaults() *OSAccount {
	this := OSAccount{}
	return &this
}

// GetCreated returns the Created field value
func (o *OSAccount) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *OSAccount) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *OSAccount) SetCreated(v time.Time) {
	o.Created = v
}

// GetDeviceId returns the DeviceId field value
func (o *OSAccount) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *OSAccount) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *OSAccount) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetId returns the Id field value
func (o *OSAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OSAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OSAccount) SetId(v string) {
	o.Id = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OSAccount) GetLastSeenAt() time.Time {
	if o == nil || IsNil(o.LastSeenAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt.Get()
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OSAccount) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastSeenAt.Get(), o.LastSeenAt.IsSet()
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *OSAccount) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt.IsSet() {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given NullableTime and assigns it to the LastSeenAt field.
func (o *OSAccount) SetLastSeenAt(v time.Time) {
	o.LastSeenAt.Set(&v)
}

// SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil
func (o *OSAccount) SetLastSeenAtNil() {
	o.LastSeenAt.Set(nil)
}

// UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
func (o *OSAccount) UnsetLastSeenAt() {
	o.LastSeenAt.Unset()
}

// GetLastUpdated returns the LastUpdated field value
func (o *OSAccount) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *OSAccount) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *OSAccount) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetPlatform returns the Platform field value
func (o *OSAccount) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *OSAccount) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *OSAccount) SetPlatform(v string) {
	o.Platform = v
}

// GetResourceAlternateId returns the ResourceAlternateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OSAccount) GetResourceAlternateId() string {
	if o == nil || IsNil(o.ResourceAlternateId.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceAlternateId.Get()
}

// GetResourceAlternateIdOk returns a tuple with the ResourceAlternateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OSAccount) GetResourceAlternateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceAlternateId.Get(), o.ResourceAlternateId.IsSet()
}

// HasResourceAlternateId returns a boolean if a field has been set.
func (o *OSAccount) HasResourceAlternateId() bool {
	if o != nil && o.ResourceAlternateId.IsSet() {
		return true
	}

	return false
}

// SetResourceAlternateId gets a reference to the given NullableString and assigns it to the ResourceAlternateId field.
func (o *OSAccount) SetResourceAlternateId(v string) {
	o.ResourceAlternateId.Set(&v)
}

// SetResourceAlternateIdNil sets the value for ResourceAlternateId to be an explicit nil
func (o *OSAccount) SetResourceAlternateIdNil() {
	o.ResourceAlternateId.Set(nil)
}

// UnsetResourceAlternateId ensures that no value is present for ResourceAlternateId, not even an explicit nil
func (o *OSAccount) UnsetResourceAlternateId() {
	o.ResourceAlternateId.Unset()
}

// GetResourceDisplayName returns the ResourceDisplayName field value if set, zero value otherwise.
func (o *OSAccount) GetResourceDisplayName() OSAccountDisplayName {
	if o == nil || IsNil(o.ResourceDisplayName) {
		var ret OSAccountDisplayName
		return ret
	}
	return *o.ResourceDisplayName
}

// GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSAccount) GetResourceDisplayNameOk() (*OSAccountDisplayName, bool) {
	if o == nil || IsNil(o.ResourceDisplayName) {
		return nil, false
	}
	return o.ResourceDisplayName, true
}

// HasResourceDisplayName returns a boolean if a field has been set.
func (o *OSAccount) HasResourceDisplayName() bool {
	if o != nil && !IsNil(o.ResourceDisplayName) {
		return true
	}

	return false
}

// SetResourceDisplayName gets a reference to the given OSAccountDisplayName and assigns it to the ResourceDisplayName field.
func (o *OSAccount) SetResourceDisplayName(v OSAccountDisplayName) {
	o.ResourceDisplayName = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *OSAccount) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSAccount) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *OSAccount) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *OSAccount) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *OSAccount) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSAccount) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *OSAccount) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *OSAccount) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStatus returns the Status field value
func (o *OSAccount) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OSAccount) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OSAccount) SetStatus(v string) {
	o.Status = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *OSAccount) GetEmbedded() OSAccountEmbedded {
	if o == nil || IsNil(o.Embedded) {
		var ret OSAccountEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OSAccount) GetEmbeddedOk() (*OSAccountEmbedded, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *OSAccount) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given OSAccountEmbedded and assigns it to the Embedded field.
func (o *OSAccount) SetEmbedded(v OSAccountEmbedded) {
	o.Embedded = &v
}

// GetLinks returns the Links field value
func (o *OSAccount) GetLinks() OSAccountLinks {
	if o == nil {
		var ret OSAccountLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *OSAccount) GetLinksOk() (*OSAccountLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *OSAccount) SetLinks(v OSAccountLinks) {
	o.Links = v
}

func (o OSAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OSAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["deviceId"] = o.DeviceId
	toSerialize["id"] = o.Id
	if o.LastSeenAt.IsSet() {
		toSerialize["lastSeenAt"] = o.LastSeenAt.Get()
	}
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["platform"] = o.Platform
	if o.ResourceAlternateId.IsSet() {
		toSerialize["resourceAlternateId"] = o.ResourceAlternateId.Get()
	}
	if !IsNil(o.ResourceDisplayName) {
		toSerialize["resourceDisplayName"] = o.ResourceDisplayName
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OSAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"deviceId",
		"id",
		"lastUpdated",
		"platform",
		"status",
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

	varOSAccount := _OSAccount{}

	err = json.Unmarshal(data, &varOSAccount)

	if err != nil {
		return err
	}

	*o = OSAccount(varOSAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "deviceId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastSeenAt")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "resourceAlternateId")
		delete(additionalProperties, "resourceDisplayName")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOSAccount struct {
	value *OSAccount
	isSet bool
}

func (v NullableOSAccount) Get() *OSAccount {
	return v.value
}

func (v *NullableOSAccount) Set(val *OSAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableOSAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableOSAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOSAccount(val *OSAccount) *NullableOSAccount {
	return &NullableOSAccount{value: val, isSet: true}
}

func (v NullableOSAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOSAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
