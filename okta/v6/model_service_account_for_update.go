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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// ServiceAccountForUpdate struct for ServiceAccountForUpdate
type ServiceAccountForUpdate struct {
	// The description of the service account
	Description *string `json:"description,omitempty"`
	// The human-readable name for the service account
	Name *string `json:"name,omitempty" validate:"regexp=^[\\\\w\\\\-_. ]+$"`
	// A list of IDs of the Okta groups who own the service account
	OwnerGroupIds []string `json:"ownerGroupIds,omitempty"`
	// A list of IDs of the Okta users who own the service account
	OwnerUserIds []string `json:"ownerUserIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceAccountForUpdate ServiceAccountForUpdate

// NewServiceAccountForUpdate instantiates a new ServiceAccountForUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountForUpdate() *ServiceAccountForUpdate {
	this := ServiceAccountForUpdate{}
	return &this
}

// NewServiceAccountForUpdateWithDefaults instantiates a new ServiceAccountForUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountForUpdateWithDefaults() *ServiceAccountForUpdate {
	this := ServiceAccountForUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccountForUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountForUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountForUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountForUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAccountForUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountForUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountForUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountForUpdate) SetName(v string) {
	o.Name = &v
}

// GetOwnerGroupIds returns the OwnerGroupIds field value if set, zero value otherwise.
func (o *ServiceAccountForUpdate) GetOwnerGroupIds() []string {
	if o == nil || o.OwnerGroupIds == nil {
		var ret []string
		return ret
	}
	return o.OwnerGroupIds
}

// GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountForUpdate) GetOwnerGroupIdsOk() ([]string, bool) {
	if o == nil || o.OwnerGroupIds == nil {
		return nil, false
	}
	return o.OwnerGroupIds, true
}

// HasOwnerGroupIds returns a boolean if a field has been set.
func (o *ServiceAccountForUpdate) HasOwnerGroupIds() bool {
	if o != nil && o.OwnerGroupIds != nil {
		return true
	}

	return false
}

// SetOwnerGroupIds gets a reference to the given []string and assigns it to the OwnerGroupIds field.
func (o *ServiceAccountForUpdate) SetOwnerGroupIds(v []string) {
	o.OwnerGroupIds = v
}

// GetOwnerUserIds returns the OwnerUserIds field value if set, zero value otherwise.
func (o *ServiceAccountForUpdate) GetOwnerUserIds() []string {
	if o == nil || o.OwnerUserIds == nil {
		var ret []string
		return ret
	}
	return o.OwnerUserIds
}

// GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountForUpdate) GetOwnerUserIdsOk() ([]string, bool) {
	if o == nil || o.OwnerUserIds == nil {
		return nil, false
	}
	return o.OwnerUserIds, true
}

// HasOwnerUserIds returns a boolean if a field has been set.
func (o *ServiceAccountForUpdate) HasOwnerUserIds() bool {
	if o != nil && o.OwnerUserIds != nil {
		return true
	}

	return false
}

// SetOwnerUserIds gets a reference to the given []string and assigns it to the OwnerUserIds field.
func (o *ServiceAccountForUpdate) SetOwnerUserIds(v []string) {
	o.OwnerUserIds = v
}

func (o ServiceAccountForUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OwnerGroupIds != nil {
		toSerialize["ownerGroupIds"] = o.OwnerGroupIds
	}
	if o.OwnerUserIds != nil {
		toSerialize["ownerUserIds"] = o.OwnerUserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceAccountForUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varServiceAccountForUpdate := _ServiceAccountForUpdate{}

	err = json.Unmarshal(bytes, &varServiceAccountForUpdate)
	if err == nil {
		*o = ServiceAccountForUpdate(varServiceAccountForUpdate)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ownerGroupIds")
		delete(additionalProperties, "ownerUserIds")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableServiceAccountForUpdate struct {
	value *ServiceAccountForUpdate
	isSet bool
}

func (v NullableServiceAccountForUpdate) Get() *ServiceAccountForUpdate {
	return v.value
}

func (v *NullableServiceAccountForUpdate) Set(val *ServiceAccountForUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountForUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountForUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountForUpdate(val *ServiceAccountForUpdate) *NullableServiceAccountForUpdate {
	return &NullableServiceAccountForUpdate{value: val, isSet: true}
}

func (v NullableServiceAccountForUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountForUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

