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
)

// checks if the ImportRoadblockEntitlementChanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRoadblockEntitlementChanges{}

// ImportRoadblockEntitlementChanges The entitlement changes that the import would apply. Only returned for `ENTITLEMENT_VALUE_CHANGE` roadblocks.
type ImportRoadblockEntitlementChanges struct {
	// The absolute threshold (number of users) in effect when the roadblock was created. Only returned when the absolute threshold is enabled.
	AbsoluteThreshold *int32 `json:"absoluteThreshold,omitempty"`
	// Number of users whose entitlements the import would change or remove (`modifiedUsers` + `deletedUsers`). Compared with `absoluteThreshold`.
	ChangedUsers *int32 `json:"changedUsers,omitempty"`
	// The share of users whose entitlements the import would change or remove (`changedUsers` / `totalUsers` × 100). Compared with `percentageThreshold`.
	ChangePercentage *float32 `json:"changePercentage,omitempty"`
	// Number of users whose entitlements the import would remove because they're no longer in the source
	DeletedUsers *int32 `json:"deletedUsers,omitempty"`
	// Number of users whose entitlements the import would change
	ModifiedUsers *int32 `json:"modifiedUsers,omitempty"`
	// The percentage threshold in effect when the roadblock was created. Only returned when the percentage threshold is enabled.
	PercentageThreshold *float32 `json:"percentageThreshold,omitempty"`
	// The threshold that the import crossed. If both thresholds were crossed, `PERCENTAGE` is returned.
	ThresholdType *string `json:"thresholdType,omitempty"`
	// Number of users with entitlements before the import (the baseline). Used as the denominator for `changePercentage`.
	TotalUsers           *int32 `json:"totalUsers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportRoadblockEntitlementChanges ImportRoadblockEntitlementChanges

// NewImportRoadblockEntitlementChanges instantiates a new ImportRoadblockEntitlementChanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRoadblockEntitlementChanges() *ImportRoadblockEntitlementChanges {
	this := ImportRoadblockEntitlementChanges{}
	return &this
}

// NewImportRoadblockEntitlementChangesWithDefaults instantiates a new ImportRoadblockEntitlementChanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRoadblockEntitlementChangesWithDefaults() *ImportRoadblockEntitlementChanges {
	this := ImportRoadblockEntitlementChanges{}
	return &this
}

// GetAbsoluteThreshold returns the AbsoluteThreshold field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementChanges) GetAbsoluteThreshold() int32 {
	if o == nil || IsNil(o.AbsoluteThreshold) {
		var ret int32
		return ret
	}
	return *o.AbsoluteThreshold
}

// GetAbsoluteThresholdOk returns a tuple with the AbsoluteThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementChanges) GetAbsoluteThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.AbsoluteThreshold) {
		return nil, false
	}
	return o.AbsoluteThreshold, true
}

// HasAbsoluteThreshold returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementChanges) HasAbsoluteThreshold() bool {
	if o != nil && !IsNil(o.AbsoluteThreshold) {
		return true
	}

	return false
}

// SetAbsoluteThreshold gets a reference to the given int32 and assigns it to the AbsoluteThreshold field.
func (o *ImportRoadblockEntitlementChanges) SetAbsoluteThreshold(v int32) {
	o.AbsoluteThreshold = &v
}

// GetChangedUsers returns the ChangedUsers field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementChanges) GetChangedUsers() int32 {
	if o == nil || IsNil(o.ChangedUsers) {
		var ret int32
		return ret
	}
	return *o.ChangedUsers
}

// GetChangedUsersOk returns a tuple with the ChangedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementChanges) GetChangedUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.ChangedUsers) {
		return nil, false
	}
	return o.ChangedUsers, true
}

// HasChangedUsers returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementChanges) HasChangedUsers() bool {
	if o != nil && !IsNil(o.ChangedUsers) {
		return true
	}

	return false
}

// SetChangedUsers gets a reference to the given int32 and assigns it to the ChangedUsers field.
func (o *ImportRoadblockEntitlementChanges) SetChangedUsers(v int32) {
	o.ChangedUsers = &v
}

// GetChangePercentage returns the ChangePercentage field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementChanges) GetChangePercentage() float32 {
	if o == nil || IsNil(o.ChangePercentage) {
		var ret float32
		return ret
	}
	return *o.ChangePercentage
}

// GetChangePercentageOk returns a tuple with the ChangePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementChanges) GetChangePercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.ChangePercentage) {
		return nil, false
	}
	return o.ChangePercentage, true
}

// HasChangePercentage returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementChanges) HasChangePercentage() bool {
	if o != nil && !IsNil(o.ChangePercentage) {
		return true
	}

	return false
}

// SetChangePercentage gets a reference to the given float32 and assigns it to the ChangePercentage field.
func (o *ImportRoadblockEntitlementChanges) SetChangePercentage(v float32) {
	o.ChangePercentage = &v
}

// GetDeletedUsers returns the DeletedUsers field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementChanges) GetDeletedUsers() int32 {
	if o == nil || IsNil(o.DeletedUsers) {
		var ret int32
		return ret
	}
	return *o.DeletedUsers
}

// GetDeletedUsersOk returns a tuple with the DeletedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementChanges) GetDeletedUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.DeletedUsers) {
		return nil, false
	}
	return o.DeletedUsers, true
}

// HasDeletedUsers returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementChanges) HasDeletedUsers() bool {
	if o != nil && !IsNil(o.DeletedUsers) {
		return true
	}

	return false
}

// SetDeletedUsers gets a reference to the given int32 and assigns it to the DeletedUsers field.
func (o *ImportRoadblockEntitlementChanges) SetDeletedUsers(v int32) {
	o.DeletedUsers = &v
}

// GetModifiedUsers returns the ModifiedUsers field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementChanges) GetModifiedUsers() int32 {
	if o == nil || IsNil(o.ModifiedUsers) {
		var ret int32
		return ret
	}
	return *o.ModifiedUsers
}

// GetModifiedUsersOk returns a tuple with the ModifiedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementChanges) GetModifiedUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.ModifiedUsers) {
		return nil, false
	}
	return o.ModifiedUsers, true
}

// HasModifiedUsers returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementChanges) HasModifiedUsers() bool {
	if o != nil && !IsNil(o.ModifiedUsers) {
		return true
	}

	return false
}

// SetModifiedUsers gets a reference to the given int32 and assigns it to the ModifiedUsers field.
func (o *ImportRoadblockEntitlementChanges) SetModifiedUsers(v int32) {
	o.ModifiedUsers = &v
}

// GetPercentageThreshold returns the PercentageThreshold field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementChanges) GetPercentageThreshold() float32 {
	if o == nil || IsNil(o.PercentageThreshold) {
		var ret float32
		return ret
	}
	return *o.PercentageThreshold
}

// GetPercentageThresholdOk returns a tuple with the PercentageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementChanges) GetPercentageThresholdOk() (*float32, bool) {
	if o == nil || IsNil(o.PercentageThreshold) {
		return nil, false
	}
	return o.PercentageThreshold, true
}

// HasPercentageThreshold returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementChanges) HasPercentageThreshold() bool {
	if o != nil && !IsNil(o.PercentageThreshold) {
		return true
	}

	return false
}

// SetPercentageThreshold gets a reference to the given float32 and assigns it to the PercentageThreshold field.
func (o *ImportRoadblockEntitlementChanges) SetPercentageThreshold(v float32) {
	o.PercentageThreshold = &v
}

// GetThresholdType returns the ThresholdType field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementChanges) GetThresholdType() string {
	if o == nil || IsNil(o.ThresholdType) {
		var ret string
		return ret
	}
	return *o.ThresholdType
}

// GetThresholdTypeOk returns a tuple with the ThresholdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementChanges) GetThresholdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ThresholdType) {
		return nil, false
	}
	return o.ThresholdType, true
}

// HasThresholdType returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementChanges) HasThresholdType() bool {
	if o != nil && !IsNil(o.ThresholdType) {
		return true
	}

	return false
}

// SetThresholdType gets a reference to the given string and assigns it to the ThresholdType field.
func (o *ImportRoadblockEntitlementChanges) SetThresholdType(v string) {
	o.ThresholdType = &v
}

// GetTotalUsers returns the TotalUsers field value if set, zero value otherwise.
func (o *ImportRoadblockEntitlementChanges) GetTotalUsers() int32 {
	if o == nil || IsNil(o.TotalUsers) {
		var ret int32
		return ret
	}
	return *o.TotalUsers
}

// GetTotalUsersOk returns a tuple with the TotalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblockEntitlementChanges) GetTotalUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalUsers) {
		return nil, false
	}
	return o.TotalUsers, true
}

// HasTotalUsers returns a boolean if a field has been set.
func (o *ImportRoadblockEntitlementChanges) HasTotalUsers() bool {
	if o != nil && !IsNil(o.TotalUsers) {
		return true
	}

	return false
}

// SetTotalUsers gets a reference to the given int32 and assigns it to the TotalUsers field.
func (o *ImportRoadblockEntitlementChanges) SetTotalUsers(v int32) {
	o.TotalUsers = &v
}

func (o ImportRoadblockEntitlementChanges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRoadblockEntitlementChanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbsoluteThreshold) {
		toSerialize["absoluteThreshold"] = o.AbsoluteThreshold
	}
	if !IsNil(o.ChangedUsers) {
		toSerialize["changedUsers"] = o.ChangedUsers
	}
	if !IsNil(o.ChangePercentage) {
		toSerialize["changePercentage"] = o.ChangePercentage
	}
	if !IsNil(o.DeletedUsers) {
		toSerialize["deletedUsers"] = o.DeletedUsers
	}
	if !IsNil(o.ModifiedUsers) {
		toSerialize["modifiedUsers"] = o.ModifiedUsers
	}
	if !IsNil(o.PercentageThreshold) {
		toSerialize["percentageThreshold"] = o.PercentageThreshold
	}
	if !IsNil(o.ThresholdType) {
		toSerialize["thresholdType"] = o.ThresholdType
	}
	if !IsNil(o.TotalUsers) {
		toSerialize["totalUsers"] = o.TotalUsers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportRoadblockEntitlementChanges) UnmarshalJSON(data []byte) (err error) {
	varImportRoadblockEntitlementChanges := _ImportRoadblockEntitlementChanges{}

	err = json.Unmarshal(data, &varImportRoadblockEntitlementChanges)

	if err != nil {
		return err
	}

	*o = ImportRoadblockEntitlementChanges(varImportRoadblockEntitlementChanges)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "absoluteThreshold")
		delete(additionalProperties, "changedUsers")
		delete(additionalProperties, "changePercentage")
		delete(additionalProperties, "deletedUsers")
		delete(additionalProperties, "modifiedUsers")
		delete(additionalProperties, "percentageThreshold")
		delete(additionalProperties, "thresholdType")
		delete(additionalProperties, "totalUsers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportRoadblockEntitlementChanges struct {
	value *ImportRoadblockEntitlementChanges
	isSet bool
}

func (v NullableImportRoadblockEntitlementChanges) Get() *ImportRoadblockEntitlementChanges {
	return v.value
}

func (v *NullableImportRoadblockEntitlementChanges) Set(val *ImportRoadblockEntitlementChanges) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRoadblockEntitlementChanges) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRoadblockEntitlementChanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRoadblockEntitlementChanges(val *ImportRoadblockEntitlementChanges) *NullableImportRoadblockEntitlementChanges {
	return &NullableImportRoadblockEntitlementChanges{value: val, isSet: true}
}

func (v NullableImportRoadblockEntitlementChanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRoadblockEntitlementChanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
