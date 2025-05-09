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

// Device struct for Device
type Device struct {
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
	AdditionalProperties map[string]interface{}
}

type _Device Device

// NewDevice instantiates a new Device object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevice() *Device {
	this := Device{}
	return &this
}

// NewDeviceWithDefaults instantiates a new Device object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithDefaults() *Device {
	this := Device{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Device) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Device) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Device) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Device) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Device) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Device) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Device) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Device) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Device) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *Device) GetProfile() DeviceProfile {
	if o == nil || o.Profile == nil {
		var ret DeviceProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetProfileOk() (*DeviceProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *Device) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given DeviceProfile and assigns it to the Profile field.
func (o *Device) SetProfile(v DeviceProfile) {
	o.Profile = &v
}

// GetResourceAlternateId returns the ResourceAlternateId field value if set, zero value otherwise.
func (o *Device) GetResourceAlternateId() string {
	if o == nil || o.ResourceAlternateId == nil {
		var ret string
		return ret
	}
	return *o.ResourceAlternateId
}

// GetResourceAlternateIdOk returns a tuple with the ResourceAlternateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetResourceAlternateIdOk() (*string, bool) {
	if o == nil || o.ResourceAlternateId == nil {
		return nil, false
	}
	return o.ResourceAlternateId, true
}

// HasResourceAlternateId returns a boolean if a field has been set.
func (o *Device) HasResourceAlternateId() bool {
	if o != nil && o.ResourceAlternateId != nil {
		return true
	}

	return false
}

// SetResourceAlternateId gets a reference to the given string and assigns it to the ResourceAlternateId field.
func (o *Device) SetResourceAlternateId(v string) {
	o.ResourceAlternateId = &v
}

// GetResourceDisplayName returns the ResourceDisplayName field value if set, zero value otherwise.
func (o *Device) GetResourceDisplayName() DeviceDisplayName {
	if o == nil || o.ResourceDisplayName == nil {
		var ret DeviceDisplayName
		return ret
	}
	return *o.ResourceDisplayName
}

// GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetResourceDisplayNameOk() (*DeviceDisplayName, bool) {
	if o == nil || o.ResourceDisplayName == nil {
		return nil, false
	}
	return o.ResourceDisplayName, true
}

// HasResourceDisplayName returns a boolean if a field has been set.
func (o *Device) HasResourceDisplayName() bool {
	if o != nil && o.ResourceDisplayName != nil {
		return true
	}

	return false
}

// SetResourceDisplayName gets a reference to the given DeviceDisplayName and assigns it to the ResourceDisplayName field.
func (o *Device) SetResourceDisplayName(v DeviceDisplayName) {
	o.ResourceDisplayName = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *Device) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *Device) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *Device) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *Device) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *Device) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *Device) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Device) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Device) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Device) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Device) GetLinks() LinksSelfAndFullUsersLifecycle {
	if o == nil || o.Links == nil {
		var ret LinksSelfAndFullUsersLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Device) GetLinksOk() (*LinksSelfAndFullUsersLifecycle, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Device) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndFullUsersLifecycle and assigns it to the Links field.
func (o *Device) SetLinks(v LinksSelfAndFullUsersLifecycle) {
	o.Links = &v
}

func (o Device) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Device) UnmarshalJSON(bytes []byte) (err error) {
	varDevice := _Device{}

	err = json.Unmarshal(bytes, &varDevice)
	if err == nil {
		*o = Device(varDevice)
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
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDevice struct {
	value *Device
	isSet bool
}

func (v NullableDevice) Get() *Device {
	return v.value
}

func (v *NullableDevice) Set(val *Device) {
	v.value = val
	v.isSet = true
}

func (v NullableDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevice(val *Device) *NullableDevice {
	return &NullableDevice{value: val, isSet: true}
}

func (v NullableDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

