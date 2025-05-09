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

// DeviceList struct for DeviceList
type DeviceList struct {
	// Timestamp when the device was created
	Created *time.Time `json:"created,omitempty"`
	// Unique key for the device
	Id *string `json:"id,omitempty"`
	// Timestamp when the device record was last updated. Updates occur when Okta collects and saves device signals during authentication, and when the lifecycle state of the device changes.
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Profile *DeviceProfile `json:"profile,omitempty"`
	ResourceAlternateId *string `json:"resourceAlternateId,omitempty"`
	ResourceDisplayName *DeviceDisplayName `json:"resourceDisplayName,omitempty"`
	// Alternate key for the `id`
	ResourceId *string `json:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	// The state object of the device
	Status *string `json:"status,omitempty"`
	Links *LinksSelfAndFullUsersLifecycle `json:"_links,omitempty"`
	Embedded *DeviceListAllOfEmbedded `json:"_embedded,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceList DeviceList

// NewDeviceList instantiates a new DeviceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceList() *DeviceList {
	this := DeviceList{}
	return &this
}

// NewDeviceListWithDefaults instantiates a new DeviceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceListWithDefaults() *DeviceList {
	this := DeviceList{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DeviceList) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DeviceList) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DeviceList) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceList) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceList) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *DeviceList) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DeviceList) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *DeviceList) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *DeviceList) GetProfile() DeviceProfile {
	if o == nil || o.Profile == nil {
		var ret DeviceProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetProfileOk() (*DeviceProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *DeviceList) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given DeviceProfile and assigns it to the Profile field.
func (o *DeviceList) SetProfile(v DeviceProfile) {
	o.Profile = &v
}

// GetResourceAlternateId returns the ResourceAlternateId field value if set, zero value otherwise.
func (o *DeviceList) GetResourceAlternateId() string {
	if o == nil || o.ResourceAlternateId == nil {
		var ret string
		return ret
	}
	return *o.ResourceAlternateId
}

// GetResourceAlternateIdOk returns a tuple with the ResourceAlternateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetResourceAlternateIdOk() (*string, bool) {
	if o == nil || o.ResourceAlternateId == nil {
		return nil, false
	}
	return o.ResourceAlternateId, true
}

// HasResourceAlternateId returns a boolean if a field has been set.
func (o *DeviceList) HasResourceAlternateId() bool {
	if o != nil && o.ResourceAlternateId != nil {
		return true
	}

	return false
}

// SetResourceAlternateId gets a reference to the given string and assigns it to the ResourceAlternateId field.
func (o *DeviceList) SetResourceAlternateId(v string) {
	o.ResourceAlternateId = &v
}

// GetResourceDisplayName returns the ResourceDisplayName field value if set, zero value otherwise.
func (o *DeviceList) GetResourceDisplayName() DeviceDisplayName {
	if o == nil || o.ResourceDisplayName == nil {
		var ret DeviceDisplayName
		return ret
	}
	return *o.ResourceDisplayName
}

// GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetResourceDisplayNameOk() (*DeviceDisplayName, bool) {
	if o == nil || o.ResourceDisplayName == nil {
		return nil, false
	}
	return o.ResourceDisplayName, true
}

// HasResourceDisplayName returns a boolean if a field has been set.
func (o *DeviceList) HasResourceDisplayName() bool {
	if o != nil && o.ResourceDisplayName != nil {
		return true
	}

	return false
}

// SetResourceDisplayName gets a reference to the given DeviceDisplayName and assigns it to the ResourceDisplayName field.
func (o *DeviceList) SetResourceDisplayName(v DeviceDisplayName) {
	o.ResourceDisplayName = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *DeviceList) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *DeviceList) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *DeviceList) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *DeviceList) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *DeviceList) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *DeviceList) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceList) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceList) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceList) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DeviceList) GetLinks() LinksSelfAndFullUsersLifecycle {
	if o == nil || o.Links == nil {
		var ret LinksSelfAndFullUsersLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetLinksOk() (*LinksSelfAndFullUsersLifecycle, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DeviceList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndFullUsersLifecycle and assigns it to the Links field.
func (o *DeviceList) SetLinks(v LinksSelfAndFullUsersLifecycle) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *DeviceList) GetEmbedded() DeviceListAllOfEmbedded {
	if o == nil || o.Embedded == nil {
		var ret DeviceListAllOfEmbedded
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceList) GetEmbeddedOk() (*DeviceListAllOfEmbedded, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *DeviceList) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given DeviceListAllOfEmbedded and assigns it to the Embedded field.
func (o *DeviceList) SetEmbedded(v DeviceListAllOfEmbedded) {
	o.Embedded = &v
}

func (o DeviceList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.ResourceAlternateId != nil {
		toSerialize["resourceAlternateId"] = o.ResourceAlternateId
	}
	if o.ResourceDisplayName != nil {
		toSerialize["resourceDisplayName"] = o.ResourceDisplayName
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeviceList) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceList := _DeviceList{}

	err = json.Unmarshal(bytes, &varDeviceList)
	if err == nil {
		*o = DeviceList(varDeviceList)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "resourceAlternateId")
		delete(additionalProperties, "resourceDisplayName")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "_embedded")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDeviceList struct {
	value *DeviceList
	isSet bool
}

func (v NullableDeviceList) Get() *DeviceList {
	return v.value
}

func (v *NullableDeviceList) Set(val *DeviceList) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceList) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceList(val *DeviceList) *NullableDeviceList {
	return &NullableDeviceList{value: val, isSet: true}
}

func (v NullableDeviceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

