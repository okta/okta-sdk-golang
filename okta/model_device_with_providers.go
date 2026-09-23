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

// checks if the DeviceWithProviders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceWithProviders{}

// DeviceWithProviders struct for DeviceWithProviders
type DeviceWithProviders struct {
	// Timestamp when the device was created
	Created *time.Time `json:"created,omitempty"`
	// Unique key for the device record. This identifier must be unique across all devices. If two or more physical devices report a non-unique device profile attribute (such as a truncated `profile.udid`), Okta may map them to the same device record (shared `id`), which can cause unexpected behavior.
	Id *string `json:"id,omitempty"`
	// Timestamp when the device record was last updated. Updates occur when Okta collects and saves device signals during authentication, and when the lifecycle state of the device changes.
	LastUpdated *time.Time     `json:"lastUpdated,omitempty"`
	Profile     *DeviceProfile `json:"profile,omitempty"`
	// Registration grants associated with this device. Present only when the device was pre-registered. Omitted (not an empty array) if no grants exist for the device.
	RegistrationGrants  []RegistrationGrant `json:"registrationGrants,omitempty"`
	ResourceAlternateId *string             `json:"resourceAlternateId,omitempty"`
	ResourceDisplayName *DeviceDisplayName  `json:"resourceDisplayName,omitempty"`
	// Alternate key for the `id`
	ResourceId   *string `json:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	// The state object of the device
	Status *string                         `json:"status,omitempty"`
	Links  *LinksSelfAndFullUsersLifecycle `json:"_links,omitempty"`
	// List of providers for the device when the `expand=providers` query parameter is specified
	Providers            []DeviceProvider `json:"providers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceWithProviders DeviceWithProviders

// NewDeviceWithProviders instantiates a new DeviceWithProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceWithProviders() *DeviceWithProviders {
	this := DeviceWithProviders{}
	return &this
}

// NewDeviceWithProvidersWithDefaults instantiates a new DeviceWithProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceWithProvidersWithDefaults() *DeviceWithProviders {
	this := DeviceWithProviders{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *DeviceWithProviders) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceWithProviders) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *DeviceWithProviders) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetProfile() DeviceProfile {
	if o == nil || IsNil(o.Profile) {
		var ret DeviceProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetProfileOk() (*DeviceProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given DeviceProfile and assigns it to the Profile field.
func (o *DeviceWithProviders) SetProfile(v DeviceProfile) {
	o.Profile = &v
}

// GetRegistrationGrants returns the RegistrationGrants field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetRegistrationGrants() []RegistrationGrant {
	if o == nil || IsNil(o.RegistrationGrants) {
		var ret []RegistrationGrant
		return ret
	}
	return o.RegistrationGrants
}

// GetRegistrationGrantsOk returns a tuple with the RegistrationGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetRegistrationGrantsOk() ([]RegistrationGrant, bool) {
	if o == nil || IsNil(o.RegistrationGrants) {
		return nil, false
	}
	return o.RegistrationGrants, true
}

// HasRegistrationGrants returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasRegistrationGrants() bool {
	if o != nil && !IsNil(o.RegistrationGrants) {
		return true
	}

	return false
}

// SetRegistrationGrants gets a reference to the given []RegistrationGrant and assigns it to the RegistrationGrants field.
func (o *DeviceWithProviders) SetRegistrationGrants(v []RegistrationGrant) {
	o.RegistrationGrants = v
}

// GetResourceAlternateId returns the ResourceAlternateId field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetResourceAlternateId() string {
	if o == nil || IsNil(o.ResourceAlternateId) {
		var ret string
		return ret
	}
	return *o.ResourceAlternateId
}

// GetResourceAlternateIdOk returns a tuple with the ResourceAlternateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetResourceAlternateIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceAlternateId) {
		return nil, false
	}
	return o.ResourceAlternateId, true
}

// HasResourceAlternateId returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasResourceAlternateId() bool {
	if o != nil && !IsNil(o.ResourceAlternateId) {
		return true
	}

	return false
}

// SetResourceAlternateId gets a reference to the given string and assigns it to the ResourceAlternateId field.
func (o *DeviceWithProviders) SetResourceAlternateId(v string) {
	o.ResourceAlternateId = &v
}

// GetResourceDisplayName returns the ResourceDisplayName field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetResourceDisplayName() DeviceDisplayName {
	if o == nil || IsNil(o.ResourceDisplayName) {
		var ret DeviceDisplayName
		return ret
	}
	return *o.ResourceDisplayName
}

// GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetResourceDisplayNameOk() (*DeviceDisplayName, bool) {
	if o == nil || IsNil(o.ResourceDisplayName) {
		return nil, false
	}
	return o.ResourceDisplayName, true
}

// HasResourceDisplayName returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasResourceDisplayName() bool {
	if o != nil && !IsNil(o.ResourceDisplayName) {
		return true
	}

	return false
}

// SetResourceDisplayName gets a reference to the given DeviceDisplayName and assigns it to the ResourceDisplayName field.
func (o *DeviceWithProviders) SetResourceDisplayName(v DeviceDisplayName) {
	o.ResourceDisplayName = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *DeviceWithProviders) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *DeviceWithProviders) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceWithProviders) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetLinks() LinksSelfAndFullUsersLifecycle {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelfAndFullUsersLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetLinksOk() (*LinksSelfAndFullUsersLifecycle, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndFullUsersLifecycle and assigns it to the Links field.
func (o *DeviceWithProviders) SetLinks(v LinksSelfAndFullUsersLifecycle) {
	o.Links = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *DeviceWithProviders) GetProviders() []DeviceProvider {
	if o == nil || IsNil(o.Providers) {
		var ret []DeviceProvider
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceWithProviders) GetProvidersOk() ([]DeviceProvider, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *DeviceWithProviders) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []DeviceProvider and assigns it to the Providers field.
func (o *DeviceWithProviders) SetProviders(v []DeviceProvider) {
	o.Providers = v
}

func (o DeviceWithProviders) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceWithProviders) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RegistrationGrants) {
		toSerialize["registrationGrants"] = o.RegistrationGrants
	}
	if !IsNil(o.ResourceAlternateId) {
		toSerialize["resourceAlternateId"] = o.ResourceAlternateId
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceWithProviders) UnmarshalJSON(data []byte) (err error) {
	varDeviceWithProviders := _DeviceWithProviders{}

	err = json.Unmarshal(data, &varDeviceWithProviders)

	if err != nil {
		return err
	}

	*o = DeviceWithProviders(varDeviceWithProviders)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "registrationGrants")
		delete(additionalProperties, "resourceAlternateId")
		delete(additionalProperties, "resourceDisplayName")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "providers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceWithProviders struct {
	value *DeviceWithProviders
	isSet bool
}

func (v NullableDeviceWithProviders) Get() *DeviceWithProviders {
	return v.value
}

func (v *NullableDeviceWithProviders) Set(val *DeviceWithProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceWithProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceWithProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceWithProviders(val *DeviceWithProviders) *NullableDeviceWithProviders {
	return &NullableDeviceWithProviders{value: val, isSet: true}
}

func (v NullableDeviceWithProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceWithProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
