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
	"fmt"
	"time"
)

// checks if the RealmAssignmentOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmAssignmentOperationResponse{}

// RealmAssignmentOperationResponse struct for RealmAssignmentOperationResponse
type RealmAssignmentOperationResponse struct {
	// Timestamp of when the operation completed
	Completed *time.Time `json:"completed,omitempty"`
	// Timestamp of when the operation was created
	Created time.Time `json:"created"`
	// ID of the asynchronous operation
	Id string `json:"id"`
	// Timestamp of when the operation started
	Started *time.Time `json:"started,omitempty"`
	// The status of the asynchronous operation
	Status string `json:"status"`
	// The operation type
	Type                string                                                    `json:"type"`
	AssignmentOperation *RealmAssignmentOperationResponseAllOfAssignmentOperation `json:"assignmentOperation,omitempty"`
	// Number of users moved
	NumUserMoved *float32 `json:"numUserMoved,omitempty"`
	// ID of the realm
	RealmId *string `json:"realmId,omitempty"`
	// Name of the realm
	RealmName            *string    `json:"realmName,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmAssignmentOperationResponse RealmAssignmentOperationResponse

// NewRealmAssignmentOperationResponse instantiates a new RealmAssignmentOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmAssignmentOperationResponse(created time.Time, id string, status string, type_ string) *RealmAssignmentOperationResponse {
	this := RealmAssignmentOperationResponse{}
	this.Created = created
	this.Id = id
	this.Status = status
	this.Type = type_
	return &this
}

// NewRealmAssignmentOperationResponseWithDefaults instantiates a new RealmAssignmentOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmAssignmentOperationResponseWithDefaults() *RealmAssignmentOperationResponse {
	this := RealmAssignmentOperationResponse{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponse) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponse) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *RealmAssignmentOperationResponse) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetCreated returns the Created field value
func (o *RealmAssignmentOperationResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RealmAssignmentOperationResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *RealmAssignmentOperationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RealmAssignmentOperationResponse) SetId(v string) {
	o.Id = v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponse) GetStarted() time.Time {
	if o == nil || IsNil(o.Started) {
		var ret time.Time
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetStartedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponse) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given time.Time and assigns it to the Started field.
func (o *RealmAssignmentOperationResponse) SetStarted(v time.Time) {
	o.Started = &v
}

// GetStatus returns the Status field value
func (o *RealmAssignmentOperationResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RealmAssignmentOperationResponse) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *RealmAssignmentOperationResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RealmAssignmentOperationResponse) SetType(v string) {
	o.Type = v
}

// GetAssignmentOperation returns the AssignmentOperation field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponse) GetAssignmentOperation() RealmAssignmentOperationResponseAllOfAssignmentOperation {
	if o == nil || IsNil(o.AssignmentOperation) {
		var ret RealmAssignmentOperationResponseAllOfAssignmentOperation
		return ret
	}
	return *o.AssignmentOperation
}

// GetAssignmentOperationOk returns a tuple with the AssignmentOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetAssignmentOperationOk() (*RealmAssignmentOperationResponseAllOfAssignmentOperation, bool) {
	if o == nil || IsNil(o.AssignmentOperation) {
		return nil, false
	}
	return o.AssignmentOperation, true
}

// HasAssignmentOperation returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponse) HasAssignmentOperation() bool {
	if o != nil && !IsNil(o.AssignmentOperation) {
		return true
	}

	return false
}

// SetAssignmentOperation gets a reference to the given RealmAssignmentOperationResponseAllOfAssignmentOperation and assigns it to the AssignmentOperation field.
func (o *RealmAssignmentOperationResponse) SetAssignmentOperation(v RealmAssignmentOperationResponseAllOfAssignmentOperation) {
	o.AssignmentOperation = &v
}

// GetNumUserMoved returns the NumUserMoved field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponse) GetNumUserMoved() float32 {
	if o == nil || IsNil(o.NumUserMoved) {
		var ret float32
		return ret
	}
	return *o.NumUserMoved
}

// GetNumUserMovedOk returns a tuple with the NumUserMoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetNumUserMovedOk() (*float32, bool) {
	if o == nil || IsNil(o.NumUserMoved) {
		return nil, false
	}
	return o.NumUserMoved, true
}

// HasNumUserMoved returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponse) HasNumUserMoved() bool {
	if o != nil && !IsNil(o.NumUserMoved) {
		return true
	}

	return false
}

// SetNumUserMoved gets a reference to the given float32 and assigns it to the NumUserMoved field.
func (o *RealmAssignmentOperationResponse) SetNumUserMoved(v float32) {
	o.NumUserMoved = &v
}

// GetRealmId returns the RealmId field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponse) GetRealmId() string {
	if o == nil || IsNil(o.RealmId) {
		var ret string
		return ret
	}
	return *o.RealmId
}

// GetRealmIdOk returns a tuple with the RealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetRealmIdOk() (*string, bool) {
	if o == nil || IsNil(o.RealmId) {
		return nil, false
	}
	return o.RealmId, true
}

// HasRealmId returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponse) HasRealmId() bool {
	if o != nil && !IsNil(o.RealmId) {
		return true
	}

	return false
}

// SetRealmId gets a reference to the given string and assigns it to the RealmId field.
func (o *RealmAssignmentOperationResponse) SetRealmId(v string) {
	o.RealmId = &v
}

// GetRealmName returns the RealmName field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponse) GetRealmName() string {
	if o == nil || IsNil(o.RealmName) {
		var ret string
		return ret
	}
	return *o.RealmName
}

// GetRealmNameOk returns a tuple with the RealmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetRealmNameOk() (*string, bool) {
	if o == nil || IsNil(o.RealmName) {
		return nil, false
	}
	return o.RealmName, true
}

// HasRealmName returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponse) HasRealmName() bool {
	if o != nil && !IsNil(o.RealmName) {
		return true
	}

	return false
}

// SetRealmName gets a reference to the given string and assigns it to the RealmName field.
func (o *RealmAssignmentOperationResponse) SetRealmName(v string) {
	o.RealmName = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RealmAssignmentOperationResponse) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignmentOperationResponse) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RealmAssignmentOperationResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *RealmAssignmentOperationResponse) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o RealmAssignmentOperationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmAssignmentOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	if !IsNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	if !IsNil(o.AssignmentOperation) {
		toSerialize["assignmentOperation"] = o.AssignmentOperation
	}
	if !IsNil(o.NumUserMoved) {
		toSerialize["numUserMoved"] = o.NumUserMoved
	}
	if !IsNil(o.RealmId) {
		toSerialize["realmId"] = o.RealmId
	}
	if !IsNil(o.RealmName) {
		toSerialize["realmName"] = o.RealmName
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RealmAssignmentOperationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"id",
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

	varRealmAssignmentOperationResponse := _RealmAssignmentOperationResponse{}

	err = json.Unmarshal(data, &varRealmAssignmentOperationResponse)

	if err != nil {
		return err
	}

	*o = RealmAssignmentOperationResponse(varRealmAssignmentOperationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "completed")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "started")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "assignmentOperation")
		delete(additionalProperties, "numUserMoved")
		delete(additionalProperties, "realmId")
		delete(additionalProperties, "realmName")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmAssignmentOperationResponse struct {
	value *RealmAssignmentOperationResponse
	isSet bool
}

func (v NullableRealmAssignmentOperationResponse) Get() *RealmAssignmentOperationResponse {
	return v.value
}

func (v *NullableRealmAssignmentOperationResponse) Set(val *RealmAssignmentOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmAssignmentOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmAssignmentOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmAssignmentOperationResponse(val *RealmAssignmentOperationResponse) *NullableRealmAssignmentOperationResponse {
	return &NullableRealmAssignmentOperationResponse{value: val, isSet: true}
}

func (v NullableRealmAssignmentOperationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmAssignmentOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
