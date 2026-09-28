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

// checks if the PreRegisteredDeviceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreRegisteredDeviceResponse{}

// PreRegisteredDeviceResponse struct for PreRegisteredDeviceResponse
type PreRegisteredDeviceResponse struct {
	// Indicates whether the device record was created or updated by this call
	Action *string `json:"action,omitempty"`
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

type _PreRegisteredDeviceResponse PreRegisteredDeviceResponse

// NewPreRegisteredDeviceResponse instantiates a new PreRegisteredDeviceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreRegisteredDeviceResponse() *PreRegisteredDeviceResponse {
	this := PreRegisteredDeviceResponse{}
	return &this
}

// NewPreRegisteredDeviceResponseWithDefaults instantiates a new PreRegisteredDeviceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreRegisteredDeviceResponseWithDefaults() *PreRegisteredDeviceResponse {
	this := PreRegisteredDeviceResponse{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PreRegisteredDeviceResponse) SetAction(v string) {
	o.Action = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PreRegisteredDeviceResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PreRegisteredDeviceResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PreRegisteredDeviceResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetProfile() DeviceProfile {
	if o == nil || IsNil(o.Profile) {
		var ret DeviceProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetProfileOk() (*DeviceProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given DeviceProfile and assigns it to the Profile field.
func (o *PreRegisteredDeviceResponse) SetProfile(v DeviceProfile) {
	o.Profile = &v
}

// GetRegistrationGrants returns the RegistrationGrants field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetRegistrationGrants() []RegistrationGrant {
	if o == nil || IsNil(o.RegistrationGrants) {
		var ret []RegistrationGrant
		return ret
	}
	return o.RegistrationGrants
}

// GetRegistrationGrantsOk returns a tuple with the RegistrationGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetRegistrationGrantsOk() ([]RegistrationGrant, bool) {
	if o == nil || IsNil(o.RegistrationGrants) {
		return nil, false
	}
	return o.RegistrationGrants, true
}

// HasRegistrationGrants returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasRegistrationGrants() bool {
	if o != nil && !IsNil(o.RegistrationGrants) {
		return true
	}

	return false
}

// SetRegistrationGrants gets a reference to the given []RegistrationGrant and assigns it to the RegistrationGrants field.
func (o *PreRegisteredDeviceResponse) SetRegistrationGrants(v []RegistrationGrant) {
	o.RegistrationGrants = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *PreRegisteredDeviceResponse) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *PreRegisteredDeviceResponse) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PreRegisteredDeviceResponse) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PreRegisteredDeviceResponse) GetLinks() LinksSelfAndLifecycle {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreRegisteredDeviceResponse) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PreRegisteredDeviceResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *PreRegisteredDeviceResponse) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o PreRegisteredDeviceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreRegisteredDeviceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
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

func (o *PreRegisteredDeviceResponse) UnmarshalJSON(data []byte) (err error) {
	varPreRegisteredDeviceResponse := _PreRegisteredDeviceResponse{}

	err = json.Unmarshal(data, &varPreRegisteredDeviceResponse)

	if err != nil {
		return err
	}

	*o = PreRegisteredDeviceResponse(varPreRegisteredDeviceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
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

type NullablePreRegisteredDeviceResponse struct {
	value *PreRegisteredDeviceResponse
	isSet bool
}

func (v NullablePreRegisteredDeviceResponse) Get() *PreRegisteredDeviceResponse {
	return v.value
}

func (v *NullablePreRegisteredDeviceResponse) Set(val *PreRegisteredDeviceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePreRegisteredDeviceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePreRegisteredDeviceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreRegisteredDeviceResponse(val *PreRegisteredDeviceResponse) *NullablePreRegisteredDeviceResponse {
	return &NullablePreRegisteredDeviceResponse{value: val, isSet: true}
}

func (v NullablePreRegisteredDeviceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreRegisteredDeviceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
