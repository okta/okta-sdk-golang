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
)

// checks if the AppServiceAccountForUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppServiceAccountForUpdate{}

// AppServiceAccountForUpdate struct for AppServiceAccountForUpdate
type AppServiceAccountForUpdate struct {
	// The description of the app service account
	Description *string `json:"description,omitempty"`
	// The user-defined name for the app service account
	Name *string `json:"name,omitempty" validate:"regexp=^[\\\\w\\\\-_. ]+$"`
	// A list of IDs of the Okta groups who own the app service account
	OwnerGroupIds []string `json:"ownerGroupIds,omitempty"`
	// A list of IDs of the Okta users who own the app service account
	OwnerUserIds         []string `json:"ownerUserIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppServiceAccountForUpdate AppServiceAccountForUpdate

// NewAppServiceAccountForUpdate instantiates a new AppServiceAccountForUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppServiceAccountForUpdate() *AppServiceAccountForUpdate {
	this := AppServiceAccountForUpdate{}
	return &this
}

// NewAppServiceAccountForUpdateWithDefaults instantiates a new AppServiceAccountForUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppServiceAccountForUpdateWithDefaults() *AppServiceAccountForUpdate {
	this := AppServiceAccountForUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppServiceAccountForUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccountForUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppServiceAccountForUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppServiceAccountForUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppServiceAccountForUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccountForUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppServiceAccountForUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppServiceAccountForUpdate) SetName(v string) {
	o.Name = &v
}

// GetOwnerGroupIds returns the OwnerGroupIds field value if set, zero value otherwise.
func (o *AppServiceAccountForUpdate) GetOwnerGroupIds() []string {
	if o == nil || IsNil(o.OwnerGroupIds) {
		var ret []string
		return ret
	}
	return o.OwnerGroupIds
}

// GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccountForUpdate) GetOwnerGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerGroupIds) {
		return nil, false
	}
	return o.OwnerGroupIds, true
}

// HasOwnerGroupIds returns a boolean if a field has been set.
func (o *AppServiceAccountForUpdate) HasOwnerGroupIds() bool {
	if o != nil && !IsNil(o.OwnerGroupIds) {
		return true
	}

	return false
}

// SetOwnerGroupIds gets a reference to the given []string and assigns it to the OwnerGroupIds field.
func (o *AppServiceAccountForUpdate) SetOwnerGroupIds(v []string) {
	o.OwnerGroupIds = v
}

// GetOwnerUserIds returns the OwnerUserIds field value if set, zero value otherwise.
func (o *AppServiceAccountForUpdate) GetOwnerUserIds() []string {
	if o == nil || IsNil(o.OwnerUserIds) {
		var ret []string
		return ret
	}
	return o.OwnerUserIds
}

// GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccountForUpdate) GetOwnerUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerUserIds) {
		return nil, false
	}
	return o.OwnerUserIds, true
}

// HasOwnerUserIds returns a boolean if a field has been set.
func (o *AppServiceAccountForUpdate) HasOwnerUserIds() bool {
	if o != nil && !IsNil(o.OwnerUserIds) {
		return true
	}

	return false
}

// SetOwnerUserIds gets a reference to the given []string and assigns it to the OwnerUserIds field.
func (o *AppServiceAccountForUpdate) SetOwnerUserIds(v []string) {
	o.OwnerUserIds = v
}

func (o AppServiceAccountForUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppServiceAccountForUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OwnerGroupIds) {
		toSerialize["ownerGroupIds"] = o.OwnerGroupIds
	}
	if !IsNil(o.OwnerUserIds) {
		toSerialize["ownerUserIds"] = o.OwnerUserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppServiceAccountForUpdate) UnmarshalJSON(data []byte) (err error) {
	varAppServiceAccountForUpdate := _AppServiceAccountForUpdate{}

	err = json.Unmarshal(data, &varAppServiceAccountForUpdate)

	if err != nil {
		return err
	}

	*o = AppServiceAccountForUpdate(varAppServiceAccountForUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ownerGroupIds")
		delete(additionalProperties, "ownerUserIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppServiceAccountForUpdate struct {
	value *AppServiceAccountForUpdate
	isSet bool
}

func (v NullableAppServiceAccountForUpdate) Get() *AppServiceAccountForUpdate {
	return v.value
}

func (v *NullableAppServiceAccountForUpdate) Set(val *AppServiceAccountForUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppServiceAccountForUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppServiceAccountForUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppServiceAccountForUpdate(val *AppServiceAccountForUpdate) *NullableAppServiceAccountForUpdate {
	return &NullableAppServiceAccountForUpdate{value: val, isSet: true}
}

func (v NullableAppServiceAccountForUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppServiceAccountForUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
