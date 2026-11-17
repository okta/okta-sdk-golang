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

// checks if the DeprovisionTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprovisionTask{}

// DeprovisionTask A deprovision task
type DeprovisionTask struct {
	// ID of the app instance the task belongs to
	AppInstanceId string `json:"appInstanceId"`
	// ID of the user's assignment to the app instance
	AppUserId string `json:"appUserId"`
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
	// Why the deprovision task was created
	Purpose string `json:"purpose"`
	// Reason for the outcome. Not returned when the task has no reason.
	Reason *string `json:"reason,omitempty"`
	// ID of the user the task is for
	UserId               string `json:"userId"`
	AdditionalProperties map[string]interface{}
}

type _DeprovisionTask DeprovisionTask

// NewDeprovisionTask instantiates a new DeprovisionTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprovisionTask(appInstanceId string, appUserId string, createdDate time.Time, eligibleActions []string, error_ NullableString, id string, lastUpdatedDate time.Time, outcome string, purpose string, userId string) *DeprovisionTask {
	this := DeprovisionTask{}
	this.AppInstanceId = appInstanceId
	this.AppUserId = appUserId
	this.CreatedDate = createdDate
	this.EligibleActions = eligibleActions
	this.Error = error_
	this.Id = id
	this.LastUpdatedDate = lastUpdatedDate
	this.Outcome = outcome
	this.Purpose = purpose
	this.UserId = userId
	return &this
}

// NewDeprovisionTaskWithDefaults instantiates a new DeprovisionTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprovisionTaskWithDefaults() *DeprovisionTask {
	this := DeprovisionTask{}
	return &this
}

// GetAppInstanceId returns the AppInstanceId field value
func (o *DeprovisionTask) GetAppInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetAppInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppInstanceId, true
}

// SetAppInstanceId sets field value
func (o *DeprovisionTask) SetAppInstanceId(v string) {
	o.AppInstanceId = v
}

// GetAppUserId returns the AppUserId field value
func (o *DeprovisionTask) GetAppUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppUserId
}

// GetAppUserIdOk returns a tuple with the AppUserId field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetAppUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppUserId, true
}

// SetAppUserId sets field value
func (o *DeprovisionTask) SetAppUserId(v string) {
	o.AppUserId = v
}

// GetCreatedDate returns the CreatedDate field value
func (o *DeprovisionTask) GetCreatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDate, true
}

// SetCreatedDate sets field value
func (o *DeprovisionTask) SetCreatedDate(v time.Time) {
	o.CreatedDate = v
}

// GetEligibleActions returns the EligibleActions field value
func (o *DeprovisionTask) GetEligibleActions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EligibleActions
}

// GetEligibleActionsOk returns a tuple with the EligibleActions field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetEligibleActionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EligibleActions, true
}

// SetEligibleActions sets field value
func (o *DeprovisionTask) SetEligibleActions(v []string) {
	o.EligibleActions = v
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeprovisionTask) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeprovisionTask) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *DeprovisionTask) SetError(v string) {
	o.Error.Set(&v)
}

// GetId returns the Id field value
func (o *DeprovisionTask) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeprovisionTask) SetId(v string) {
	o.Id = v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value
func (o *DeprovisionTask) GetLastUpdatedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedDate, true
}

// SetLastUpdatedDate sets field value
func (o *DeprovisionTask) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = v
}

// GetOutcome returns the Outcome field value
func (o *DeprovisionTask) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *DeprovisionTask) SetOutcome(v string) {
	o.Outcome = v
}

// GetPurpose returns the Purpose field value
func (o *DeprovisionTask) GetPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *DeprovisionTask) SetPurpose(v string) {
	o.Purpose = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DeprovisionTask) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DeprovisionTask) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DeprovisionTask) SetReason(v string) {
	o.Reason = &v
}

// GetUserId returns the UserId field value
func (o *DeprovisionTask) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DeprovisionTask) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DeprovisionTask) SetUserId(v string) {
	o.UserId = v
}

func (o DeprovisionTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprovisionTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appInstanceId"] = o.AppInstanceId
	toSerialize["appUserId"] = o.AppUserId
	toSerialize["createdDate"] = o.CreatedDate
	toSerialize["eligibleActions"] = o.EligibleActions
	toSerialize["error"] = o.Error.Get()
	toSerialize["id"] = o.Id
	toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	toSerialize["outcome"] = o.Outcome
	toSerialize["purpose"] = o.Purpose
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["userId"] = o.UserId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeprovisionTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appInstanceId",
		"appUserId",
		"createdDate",
		"eligibleActions",
		"error",
		"id",
		"lastUpdatedDate",
		"outcome",
		"purpose",
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

	varDeprovisionTask := _DeprovisionTask{}

	err = json.Unmarshal(data, &varDeprovisionTask)

	if err != nil {
		return err
	}

	*o = DeprovisionTask(varDeprovisionTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appInstanceId")
		delete(additionalProperties, "appUserId")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "eligibleActions")
		delete(additionalProperties, "error")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdatedDate")
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "purpose")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "userId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeprovisionTask struct {
	value *DeprovisionTask
	isSet bool
}

func (v NullableDeprovisionTask) Get() *DeprovisionTask {
	return v.value
}

func (v *NullableDeprovisionTask) Set(val *DeprovisionTask) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprovisionTask) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprovisionTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprovisionTask(val *DeprovisionTask) *NullableDeprovisionTask {
	return &NullableDeprovisionTask{value: val, isSet: true}
}

func (v NullableDeprovisionTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprovisionTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
