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

// checks if the ImportRoadblock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportRoadblock{}

// ImportRoadblock A safeguard that paused user imports because an import would unassign an unusually large share of users
type ImportRoadblock struct {
	// The apps whose assignments the import would remove, each checked against its own app-level threshold. Only returned for app-level roadblocks (`type` is `GROUP_DELETE` or `GROUP_UPDATE` and the threshold crossed is an app-level one).  These are the apps that lose assignments, not the app instance whose import triggered the roadblock (`appInstanceId`). For example, deleting a group in an Active Directory source removes the assignments that the group grants to other apps.
	AffectedApps []ImportRoadblockAffectedApp `json:"affectedApps,omitempty"`
	// ID of the app instance whose import triggered the roadblock (the import source)
	AppInstanceId string `json:"appInstanceId"`
	// Label of the app instance
	AppInstanceLabel *string `json:"appInstanceLabel,omitempty"`
	// Key name of the app, for example, `active_directory`
	AppName *string `json:"appName,omitempty"`
	// Number of app assignments in the org used as the denominator for `unassignmentPercentage`. Only returned for org-level roadblocks.
	Assignments *int32 `json:"assignments,omitempty"`
	// Timestamp when the roadblock was created
	Created             time.Time                           `json:"created"`
	EntitlementChanges  *ImportRoadblockEntitlementChanges  `json:"entitlementChanges,omitempty"`
	EntitlementSettings *ImportRoadblockEntitlementSettings `json:"entitlementSettings,omitempty"`
	// Unique identifier for the roadblock
	Id string `json:"id"`
	// Indicates which user imports the roadblock pauses:  * `ORG`: Every user import in the org is paused. * `APP_INSTANCE`: Only the user imports of this app instance are paused.
	PauseScope string `json:"pauseScope"`
	// A human-readable message that describes why the roadblock was created, in the org's language. Use it for display only. Don't parse it; use `type` and the count fields instead.
	Reason *string `json:"reason,omitempty"`
	// Status of the roadblock
	Status string `json:"status"`
	// The safeguard that triggered the roadblock
	Type string `json:"type"`
	// The share of app assignments in the org that the import would remove (`unassignments` / `assignments` × 100). Only returned for org-level roadblocks.
	UnassignmentPercentage *float32 `json:"unassignmentPercentage,omitempty"`
	// Number of app assignments in the org that the import would remove. Only returned for org-level roadblocks (`pauseScope` is `ORG` and `type` is `USER_DELETE`, `GROUP_DELETE`, or `GROUP_UPDATE`).
	Unassignments        *int32 `json:"unassignments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ImportRoadblock ImportRoadblock

// NewImportRoadblock instantiates a new ImportRoadblock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRoadblock(appInstanceId string, created time.Time, id string, pauseScope string, status string, type_ string) *ImportRoadblock {
	this := ImportRoadblock{}
	this.AppInstanceId = appInstanceId
	this.Created = created
	this.Id = id
	this.PauseScope = pauseScope
	this.Status = status
	this.Type = type_
	return &this
}

// NewImportRoadblockWithDefaults instantiates a new ImportRoadblock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRoadblockWithDefaults() *ImportRoadblock {
	this := ImportRoadblock{}
	return &this
}

// GetAffectedApps returns the AffectedApps field value if set, zero value otherwise.
func (o *ImportRoadblock) GetAffectedApps() []ImportRoadblockAffectedApp {
	if o == nil || IsNil(o.AffectedApps) {
		var ret []ImportRoadblockAffectedApp
		return ret
	}
	return o.AffectedApps
}

// GetAffectedAppsOk returns a tuple with the AffectedApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetAffectedAppsOk() ([]ImportRoadblockAffectedApp, bool) {
	if o == nil || IsNil(o.AffectedApps) {
		return nil, false
	}
	return o.AffectedApps, true
}

// HasAffectedApps returns a boolean if a field has been set.
func (o *ImportRoadblock) HasAffectedApps() bool {
	if o != nil && !IsNil(o.AffectedApps) {
		return true
	}

	return false
}

// SetAffectedApps gets a reference to the given []ImportRoadblockAffectedApp and assigns it to the AffectedApps field.
func (o *ImportRoadblock) SetAffectedApps(v []ImportRoadblockAffectedApp) {
	o.AffectedApps = v
}

// GetAppInstanceId returns the AppInstanceId field value
func (o *ImportRoadblock) GetAppInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetAppInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppInstanceId, true
}

// SetAppInstanceId sets field value
func (o *ImportRoadblock) SetAppInstanceId(v string) {
	o.AppInstanceId = v
}

// GetAppInstanceLabel returns the AppInstanceLabel field value if set, zero value otherwise.
func (o *ImportRoadblock) GetAppInstanceLabel() string {
	if o == nil || IsNil(o.AppInstanceLabel) {
		var ret string
		return ret
	}
	return *o.AppInstanceLabel
}

// GetAppInstanceLabelOk returns a tuple with the AppInstanceLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetAppInstanceLabelOk() (*string, bool) {
	if o == nil || IsNil(o.AppInstanceLabel) {
		return nil, false
	}
	return o.AppInstanceLabel, true
}

// HasAppInstanceLabel returns a boolean if a field has been set.
func (o *ImportRoadblock) HasAppInstanceLabel() bool {
	if o != nil && !IsNil(o.AppInstanceLabel) {
		return true
	}

	return false
}

// SetAppInstanceLabel gets a reference to the given string and assigns it to the AppInstanceLabel field.
func (o *ImportRoadblock) SetAppInstanceLabel(v string) {
	o.AppInstanceLabel = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *ImportRoadblock) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *ImportRoadblock) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *ImportRoadblock) SetAppName(v string) {
	o.AppName = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *ImportRoadblock) GetAssignments() int32 {
	if o == nil || IsNil(o.Assignments) {
		var ret int32
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetAssignmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *ImportRoadblock) HasAssignments() bool {
	if o != nil && !IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given int32 and assigns it to the Assignments field.
func (o *ImportRoadblock) SetAssignments(v int32) {
	o.Assignments = &v
}

// GetCreated returns the Created field value
func (o *ImportRoadblock) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *ImportRoadblock) SetCreated(v time.Time) {
	o.Created = v
}

// GetEntitlementChanges returns the EntitlementChanges field value if set, zero value otherwise.
func (o *ImportRoadblock) GetEntitlementChanges() ImportRoadblockEntitlementChanges {
	if o == nil || IsNil(o.EntitlementChanges) {
		var ret ImportRoadblockEntitlementChanges
		return ret
	}
	return *o.EntitlementChanges
}

// GetEntitlementChangesOk returns a tuple with the EntitlementChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetEntitlementChangesOk() (*ImportRoadblockEntitlementChanges, bool) {
	if o == nil || IsNil(o.EntitlementChanges) {
		return nil, false
	}
	return o.EntitlementChanges, true
}

// HasEntitlementChanges returns a boolean if a field has been set.
func (o *ImportRoadblock) HasEntitlementChanges() bool {
	if o != nil && !IsNil(o.EntitlementChanges) {
		return true
	}

	return false
}

// SetEntitlementChanges gets a reference to the given ImportRoadblockEntitlementChanges and assigns it to the EntitlementChanges field.
func (o *ImportRoadblock) SetEntitlementChanges(v ImportRoadblockEntitlementChanges) {
	o.EntitlementChanges = &v
}

// GetEntitlementSettings returns the EntitlementSettings field value if set, zero value otherwise.
func (o *ImportRoadblock) GetEntitlementSettings() ImportRoadblockEntitlementSettings {
	if o == nil || IsNil(o.EntitlementSettings) {
		var ret ImportRoadblockEntitlementSettings
		return ret
	}
	return *o.EntitlementSettings
}

// GetEntitlementSettingsOk returns a tuple with the EntitlementSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetEntitlementSettingsOk() (*ImportRoadblockEntitlementSettings, bool) {
	if o == nil || IsNil(o.EntitlementSettings) {
		return nil, false
	}
	return o.EntitlementSettings, true
}

// HasEntitlementSettings returns a boolean if a field has been set.
func (o *ImportRoadblock) HasEntitlementSettings() bool {
	if o != nil && !IsNil(o.EntitlementSettings) {
		return true
	}

	return false
}

// SetEntitlementSettings gets a reference to the given ImportRoadblockEntitlementSettings and assigns it to the EntitlementSettings field.
func (o *ImportRoadblock) SetEntitlementSettings(v ImportRoadblockEntitlementSettings) {
	o.EntitlementSettings = &v
}

// GetId returns the Id field value
func (o *ImportRoadblock) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ImportRoadblock) SetId(v string) {
	o.Id = v
}

// GetPauseScope returns the PauseScope field value
func (o *ImportRoadblock) GetPauseScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PauseScope
}

// GetPauseScopeOk returns a tuple with the PauseScope field value
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetPauseScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PauseScope, true
}

// SetPauseScope sets field value
func (o *ImportRoadblock) SetPauseScope(v string) {
	o.PauseScope = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ImportRoadblock) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ImportRoadblock) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ImportRoadblock) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value
func (o *ImportRoadblock) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ImportRoadblock) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *ImportRoadblock) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImportRoadblock) SetType(v string) {
	o.Type = v
}

// GetUnassignmentPercentage returns the UnassignmentPercentage field value if set, zero value otherwise.
func (o *ImportRoadblock) GetUnassignmentPercentage() float32 {
	if o == nil || IsNil(o.UnassignmentPercentage) {
		var ret float32
		return ret
	}
	return *o.UnassignmentPercentage
}

// GetUnassignmentPercentageOk returns a tuple with the UnassignmentPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetUnassignmentPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.UnassignmentPercentage) {
		return nil, false
	}
	return o.UnassignmentPercentage, true
}

// HasUnassignmentPercentage returns a boolean if a field has been set.
func (o *ImportRoadblock) HasUnassignmentPercentage() bool {
	if o != nil && !IsNil(o.UnassignmentPercentage) {
		return true
	}

	return false
}

// SetUnassignmentPercentage gets a reference to the given float32 and assigns it to the UnassignmentPercentage field.
func (o *ImportRoadblock) SetUnassignmentPercentage(v float32) {
	o.UnassignmentPercentage = &v
}

// GetUnassignments returns the Unassignments field value if set, zero value otherwise.
func (o *ImportRoadblock) GetUnassignments() int32 {
	if o == nil || IsNil(o.Unassignments) {
		var ret int32
		return ret
	}
	return *o.Unassignments
}

// GetUnassignmentsOk returns a tuple with the Unassignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRoadblock) GetUnassignmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.Unassignments) {
		return nil, false
	}
	return o.Unassignments, true
}

// HasUnassignments returns a boolean if a field has been set.
func (o *ImportRoadblock) HasUnassignments() bool {
	if o != nil && !IsNil(o.Unassignments) {
		return true
	}

	return false
}

// SetUnassignments gets a reference to the given int32 and assigns it to the Unassignments field.
func (o *ImportRoadblock) SetUnassignments(v int32) {
	o.Unassignments = &v
}

func (o ImportRoadblock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportRoadblock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedApps) {
		toSerialize["affectedApps"] = o.AffectedApps
	}
	toSerialize["appInstanceId"] = o.AppInstanceId
	if !IsNil(o.AppInstanceLabel) {
		toSerialize["appInstanceLabel"] = o.AppInstanceLabel
	}
	if !IsNil(o.AppName) {
		toSerialize["appName"] = o.AppName
	}
	if !IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}
	toSerialize["created"] = o.Created
	if !IsNil(o.EntitlementChanges) {
		toSerialize["entitlementChanges"] = o.EntitlementChanges
	}
	if !IsNil(o.EntitlementSettings) {
		toSerialize["entitlementSettings"] = o.EntitlementSettings
	}
	toSerialize["id"] = o.Id
	toSerialize["pauseScope"] = o.PauseScope
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	if !IsNil(o.UnassignmentPercentage) {
		toSerialize["unassignmentPercentage"] = o.UnassignmentPercentage
	}
	if !IsNil(o.Unassignments) {
		toSerialize["unassignments"] = o.Unassignments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ImportRoadblock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appInstanceId",
		"created",
		"id",
		"pauseScope",
		"status",
		"type",
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

	varImportRoadblock := _ImportRoadblock{}

	err = json.Unmarshal(data, &varImportRoadblock)

	if err != nil {
		return err
	}

	*o = ImportRoadblock(varImportRoadblock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "affectedApps")
		delete(additionalProperties, "appInstanceId")
		delete(additionalProperties, "appInstanceLabel")
		delete(additionalProperties, "appName")
		delete(additionalProperties, "assignments")
		delete(additionalProperties, "created")
		delete(additionalProperties, "entitlementChanges")
		delete(additionalProperties, "entitlementSettings")
		delete(additionalProperties, "id")
		delete(additionalProperties, "pauseScope")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "unassignmentPercentage")
		delete(additionalProperties, "unassignments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableImportRoadblock struct {
	value *ImportRoadblock
	isSet bool
}

func (v NullableImportRoadblock) Get() *ImportRoadblock {
	return v.value
}

func (v *NullableImportRoadblock) Set(val *ImportRoadblock) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRoadblock) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRoadblock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRoadblock(val *ImportRoadblock) *NullableImportRoadblock {
	return &NullableImportRoadblock{value: val, isSet: true}
}

func (v NullableImportRoadblock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRoadblock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
