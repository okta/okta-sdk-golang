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

// checks if the AppAssignmentTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAssignmentTask{}

// AppAssignmentTask An app assignment task
type AppAssignmentTask struct {
	// ID of the app instance the task belongs to
	AppInstanceId string `json:"appInstanceId"`
	// ID of the user's assignment to the app instance
	AppUserId string `json:"appUserId"`
	// Type of assignment
	AssignmentType string `json:"assignmentType"`
	// Timestamp when the task was created
	CreatedDate time.Time `json:"createdDate"`
	// Actions that you can currently perform on the task. Empty when none are available, which can also be the case for a `FAILED` task. Use this list instead of working out the actions from `outcome` or `reason`.
	EligibleActions []string `json:"eligibleActions"`
	// Human-readable error detail from the task. `null` when the task has no error text.
	Error NullableString `json:"error"`
	// Unique identifier of the task
	Id string `json:"id"`
	// Timestamp when the task was last updated
	LastUpdatedDate time.Time `json:"lastUpdatedDate"`
	// Outcome of the task
	Outcome string `json:"outcome"`
	// Reason for the outcome. Not returned when the task has no reason.
	Reason *string `json:"reason,omitempty"`
	// ID of the user the task is for
	UserId               string `json:"userId"`
	AdditionalProperties map[string]interface{}
}

type _AppAssignmentTask AppAssignmentTask

// NewAppAssignmentTask instantiates a new AppAssignmentTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAssignmentTask(appInstanceId string, appUserId string, assignmentType string, createdDate time.Time, eligibleActions []string, error_ NullableString, id string, lastUpdatedDate time.Time, outcome string, userId string) *AppAssignmentTask {
	this := AppAssignmentTask{}
	this.AppInstanceId = appInstanceId
	this.AppUserId = appUserId
	this.AssignmentType = assignmentType
	this.CreatedDate = createdDate
	this.EligibleActions = eligibleActions
	this.Error = error_
	this.Id = id
	this.LastUpdatedDate = lastUpdatedDate
	this.Outcome = outcome
	this.UserId = userId
	return &this
}

// NewAppAssignmentTaskWithDefaults instantiates a new AppAssignmentTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAssignmentTaskWithDefaults() *AppAssignmentTask {
	this := AppAssignmentTask{}
	return &this
}

// GetAppInstanceId returns the AppInstanceId field value
func (o *AppAssignmentTask) GetAppInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetAppInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppInstanceId, true
}

// SetAppInstanceId sets field value
func (o *AppAssignmentTask) SetAppInstanceId(v string) {
	o.AppInstanceId = v
}

// GetAppUserId returns the AppUserId field value
func (o *AppAssignmentTask) GetAppUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppUserId
}

// GetAppUserIdOk returns a tuple with the AppUserId field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetAppUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUserId, true
}

// SetAppUserId sets field value
func (o *AppAssignmentTask) SetAppUserId(v string) {
	o.AppUserId = v
}

// GetAssignmentType returns the AssignmentType field value
func (o *AppAssignmentTask) GetAssignmentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignmentType
}

// GetAssignmentTypeOk returns a tuple with the AssignmentType field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetAssignmentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignmentType, true
}

// SetAssignmentType sets field value
func (o *AppAssignmentTask) SetAssignmentType(v string) {
	o.AssignmentType = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *AppAssignmentTask) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *AppAssignmentTask) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetEligibleActions returns the EligibleActions field value
func (o *AppAssignmentTask) GetEligibleActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EligibleActions
}

// GetEligibleActionsOk returns a tuple with the EligibleActions field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetEligibleActionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EligibleActions, true
}

// SetEligibleActions sets field value
func (o *AppAssignmentTask) SetEligibleActions(v []string) {
	o.EligibleActions = v
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AppAssignmentTask) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAssignmentTask) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *AppAssignmentTask) SetError(v string) {
	o.Error.Set(&v)
}

// GetId returns the Id field value
func (o *AppAssignmentTask) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppAssignmentTask) SetId(v string) {
	o.Id = v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value
func (o *AppAssignmentTask) GetLastUpdatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedDate, true
}

// SetLastUpdatedDate sets field value
func (o *AppAssignmentTask) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = v
}

// GetOutcome returns the Outcome field value
func (o *AppAssignmentTask) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *AppAssignmentTask) SetOutcome(v string) {
	o.Outcome = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *AppAssignmentTask) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *AppAssignmentTask) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *AppAssignmentTask) SetReason(v string) {
	o.Reason = &v
}

// GetUserId returns the UserId field value
func (o *AppAssignmentTask) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *AppAssignmentTask) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *AppAssignmentTask) SetUserId(v string) {
	o.UserId = v
}

func (o AppAssignmentTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAssignmentTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appInstanceId"] = o.AppInstanceId
	toSerialize["appUserId"] = o.AppUserId
	toSerialize["assignmentType"] = o.AssignmentType
	toSerialize["createdDate"] = o.CreatedDate
	toSerialize["eligibleActions"] = o.EligibleActions
	toSerialize["error"] = o.Error.Get()
	toSerialize["id"] = o.Id
	toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	toSerialize["outcome"] = o.Outcome
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["userId"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppAssignmentTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appInstanceId",
		"appUserId",
		"assignmentType",
		"createdDate",
		"eligibleActions",
		"error",
		"id",
		"lastUpdatedDate",
		"outcome",
		"userId",
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

	varAppAssignmentTask := _AppAssignmentTask{}

	err = json.Unmarshal(data, &varAppAssignmentTask)

	if err != nil {
		return err
	}

	*o = AppAssignmentTask(varAppAssignmentTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appInstanceId")
		delete(additionalProperties, "appUserId")
		delete(additionalProperties, "assignmentType")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "eligibleActions")
		delete(additionalProperties, "error")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdatedDate")
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppAssignmentTask struct {
	value *AppAssignmentTask
	isSet bool
}

func (v NullableAppAssignmentTask) Get() *AppAssignmentTask {
	return v.value
}

func (v *NullableAppAssignmentTask) Set(val *AppAssignmentTask) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAssignmentTask) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAssignmentTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAssignmentTask(val *AppAssignmentTask) *NullableAppAssignmentTask {
	return &NullableAppAssignmentTask{value: val, isSet: true}
}

func (v NullableAppAssignmentTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAssignmentTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
