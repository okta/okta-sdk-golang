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

// checks if the PreRegisteredDeviceListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreRegisteredDeviceListResponse{}

// PreRegisteredDeviceListResponse Device object in the pre-registration listing (same shape as POST response without the `action` field)
type PreRegisteredDeviceListResponse struct {
	// Timestamp when the device was created
	Created *time.Time `json:"created,omitempty"`
	// Unique key for the device
	Id *string `json:"id,omitempty"`
	// Timestamp when the device record was last updated
	LastUpdated *time.Time     `json:"lastUpdated,omitempty"`
	Profile     *DeviceProfile `json:"profile,omitempty"`
	// Registration grants associated with this device
	RegistrationGrants []RegistrationGrant `json:"registrationGrants,omitempty"`
	// Alternate key for the `id`
	ResourceId   *string `json:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	// The state object of the device
	Status               *string                `json:"status,omitempty"`
	Links                *LinksSelfAndLifecycle `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PreRegisteredDeviceListResponse PreRegisteredDeviceListResponse

// NewPreRegisteredDeviceListResponse instantiates a new PreRegisteredDeviceListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreRegisteredDeviceListResponse() *PreRegisteredDeviceListResponse {
	this := PreRegisteredDeviceListResponse{}
	return &this
}

// NewPreRegisteredDeviceListResponseWithDefaults instantiates a new PreRegisteredDeviceListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreRegisteredDeviceListResponseWithDefaults() *PreRegisteredDeviceListResponse {
	this := PreRegisteredDeviceListResponse{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PreRegisteredDeviceListResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceListResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PreRegisteredDeviceListResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PreRegisteredDeviceListResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PreRegisteredDeviceListResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceListResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PreRegisteredDeviceListResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PreRegisteredDeviceListResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PreRegisteredDeviceListResponse) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceListResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PreRegisteredDeviceListResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PreRegisteredDeviceListResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PreRegisteredDeviceListResponse) GetProfile() DeviceProfile {
	if o == nil || IsNil(o.Profile) {
		var ret DeviceProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceListResponse) GetProfileOk() (*DeviceProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PreRegisteredDeviceListResponse) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given DeviceProfile and assigns it to the Profile field.
func (o *PreRegisteredDeviceListResponse) SetProfile(v DeviceProfile) {
	o.Profile = &v
}

// GetRegistrationGrants returns the RegistrationGrants field value if set, zero value otherwise.
func (o *PreRegisteredDeviceListResponse) GetRegistrationGrants() []RegistrationGrant {
	if o == nil || IsNil(o.RegistrationGrants) {
		var ret []RegistrationGrant
		return ret
	}
	return o.RegistrationGrants
}

// GetRegistrationGrantsOk returns a tuple with the RegistrationGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceListResponse) GetRegistrationGrantsOk() ([]RegistrationGrant, bool) {
	if o == nil || IsNil(o.RegistrationGrants) {
		return nil, false
	}
	return o.RegistrationGrants, true
}

// HasRegistrationGrants returns a boolean if a field has been set.
func (o *PreRegisteredDeviceListResponse) HasRegistrationGrants() bool {
	if o != nil && !IsNil(o.RegistrationGrants) {
		return true
	}

	return false
}

// SetRegistrationGrants gets a reference to the given []RegistrationGrant and assigns it to the RegistrationGrants field.
func (o *PreRegisteredDeviceListResponse) SetRegistrationGrants(v []RegistrationGrant) {
	o.RegistrationGrants = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *PreRegisteredDeviceListResponse) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceListResponse) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *PreRegisteredDeviceListResponse) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *PreRegisteredDeviceListResponse) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *PreRegisteredDeviceListResponse) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceListResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *PreRegisteredDeviceListResponse) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *PreRegisteredDeviceListResponse) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PreRegisteredDeviceListResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceListResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PreRegisteredDeviceListResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PreRegisteredDeviceListResponse) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PreRegisteredDeviceListResponse) GetLinks() LinksSelfAndLifecycle {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceListResponse) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PreRegisteredDeviceListResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *PreRegisteredDeviceListResponse) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o PreRegisteredDeviceListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreRegisteredDeviceListResponse) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PreRegisteredDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	varPreRegisteredDeviceListResponse := _PreRegisteredDeviceListResponse{}

	err = json.Unmarshal(data, &varPreRegisteredDeviceListResponse)

	if err != nil {
		return err
	}

	*o = PreRegisteredDeviceListResponse(varPreRegisteredDeviceListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "registrationGrants")
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePreRegisteredDeviceListResponse struct {
	value *PreRegisteredDeviceListResponse
	isSet bool
}

func (v NullablePreRegisteredDeviceListResponse) Get() *PreRegisteredDeviceListResponse {
	return v.value
}

func (v *NullablePreRegisteredDeviceListResponse) Set(val *PreRegisteredDeviceListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePreRegisteredDeviceListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePreRegisteredDeviceListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreRegisteredDeviceListResponse(val *PreRegisteredDeviceListResponse) *NullablePreRegisteredDeviceListResponse {
	return &NullablePreRegisteredDeviceListResponse{value: val, isSet: true}
}

func (v NullablePreRegisteredDeviceListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreRegisteredDeviceListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
