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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// OperationResponse struct for OperationResponse
type OperationResponse struct {
	AssignmentOperation *OperationResponseAssignmentOperation `json:"assignmentOperation,omitempty"`
	Completed *time.Time `json:"completed,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Id *string `json:"id,omitempty"`
	NumUserMoved *float32 `json:"numUserMoved,omitempty"`
	RealmId *string `json:"realmId,omitempty"`
	RealmName *string `json:"realmName,omitempty"`
	Started *time.Time `json:"started,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OperationResponse OperationResponse

// NewOperationResponse instantiates a new OperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationResponse() *OperationResponse {
	this := OperationResponse{}
	return &this
}

// NewOperationResponseWithDefaults instantiates a new OperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationResponseWithDefaults() *OperationResponse {
	this := OperationResponse{}
	return &this
}

// GetAssignmentOperation returns the AssignmentOperation field value if set, zero value otherwise.
func (o *OperationResponse) GetAssignmentOperation() OperationResponseAssignmentOperation {
	if o == nil || o.AssignmentOperation == nil {
		var ret OperationResponseAssignmentOperation
		return ret
	}
	return *o.AssignmentOperation
}

// GetAssignmentOperationOk returns a tuple with the AssignmentOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetAssignmentOperationOk() (*OperationResponseAssignmentOperation, bool) {
	if o == nil || o.AssignmentOperation == nil {
		return nil, false
	}
	return o.AssignmentOperation, true
}

// HasAssignmentOperation returns a boolean if a field has been set.
func (o *OperationResponse) HasAssignmentOperation() bool {
	if o != nil && o.AssignmentOperation != nil {
		return true
	}

	return false
}

// SetAssignmentOperation gets a reference to the given OperationResponseAssignmentOperation and assigns it to the AssignmentOperation field.
func (o *OperationResponse) SetAssignmentOperation(v OperationResponseAssignmentOperation) {
	o.AssignmentOperation = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *OperationResponse) GetCompleted() time.Time {
	if o == nil || o.Completed == nil {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetCompletedOk() (*time.Time, bool) {
	if o == nil || o.Completed == nil {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *OperationResponse) HasCompleted() bool {
	if o != nil && o.Completed != nil {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *OperationResponse) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OperationResponse) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OperationResponse) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *OperationResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OperationResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OperationResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OperationResponse) SetId(v string) {
	o.Id = &v
}

// GetNumUserMoved returns the NumUserMoved field value if set, zero value otherwise.
func (o *OperationResponse) GetNumUserMoved() float32 {
	if o == nil || o.NumUserMoved == nil {
		var ret float32
		return ret
	}
	return *o.NumUserMoved
}

// GetNumUserMovedOk returns a tuple with the NumUserMoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetNumUserMovedOk() (*float32, bool) {
	if o == nil || o.NumUserMoved == nil {
		return nil, false
	}
	return o.NumUserMoved, true
}

// HasNumUserMoved returns a boolean if a field has been set.
func (o *OperationResponse) HasNumUserMoved() bool {
	if o != nil && o.NumUserMoved != nil {
		return true
	}

	return false
}

// SetNumUserMoved gets a reference to the given float32 and assigns it to the NumUserMoved field.
func (o *OperationResponse) SetNumUserMoved(v float32) {
	o.NumUserMoved = &v
}

// GetRealmId returns the RealmId field value if set, zero value otherwise.
func (o *OperationResponse) GetRealmId() string {
	if o == nil || o.RealmId == nil {
		var ret string
		return ret
	}
	return *o.RealmId
}

// GetRealmIdOk returns a tuple with the RealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetRealmIdOk() (*string, bool) {
	if o == nil || o.RealmId == nil {
		return nil, false
	}
	return o.RealmId, true
}

// HasRealmId returns a boolean if a field has been set.
func (o *OperationResponse) HasRealmId() bool {
	if o != nil && o.RealmId != nil {
		return true
	}

	return false
}

// SetRealmId gets a reference to the given string and assigns it to the RealmId field.
func (o *OperationResponse) SetRealmId(v string) {
	o.RealmId = &v
}

// GetRealmName returns the RealmName field value if set, zero value otherwise.
func (o *OperationResponse) GetRealmName() string {
	if o == nil || o.RealmName == nil {
		var ret string
		return ret
	}
	return *o.RealmName
}

// GetRealmNameOk returns a tuple with the RealmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetRealmNameOk() (*string, bool) {
	if o == nil || o.RealmName == nil {
		return nil, false
	}
	return o.RealmName, true
}

// HasRealmName returns a boolean if a field has been set.
func (o *OperationResponse) HasRealmName() bool {
	if o != nil && o.RealmName != nil {
		return true
	}

	return false
}

// SetRealmName gets a reference to the given string and assigns it to the RealmName field.
func (o *OperationResponse) SetRealmName(v string) {
	o.RealmName = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *OperationResponse) GetStarted() time.Time {
	if o == nil || o.Started == nil {
		var ret time.Time
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetStartedOk() (*time.Time, bool) {
	if o == nil || o.Started == nil {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *OperationResponse) HasStarted() bool {
	if o != nil && o.Started != nil {
		return true
	}

	return false
}

// SetStarted gets a reference to the given time.Time and assigns it to the Started field.
func (o *OperationResponse) SetStarted(v time.Time) {
	o.Started = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OperationResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OperationResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OperationResponse) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OperationResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OperationResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OperationResponse) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OperationResponse) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OperationResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *OperationResponse) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o OperationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignmentOperation != nil {
		toSerialize["assignmentOperation"] = o.AssignmentOperation
	}
	if o.Completed != nil {
		toSerialize["completed"] = o.Completed
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NumUserMoved != nil {
		toSerialize["numUserMoved"] = o.NumUserMoved
	}
	if o.RealmId != nil {
		toSerialize["realmId"] = o.RealmId
	}
	if o.RealmName != nil {
		toSerialize["realmName"] = o.RealmName
	}
	if o.Started != nil {
		toSerialize["started"] = o.Started
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OperationResponse) UnmarshalJSON(bytes []byte) (err error) {
	varOperationResponse := _OperationResponse{}

	err = json.Unmarshal(bytes, &varOperationResponse)
	if err == nil {
		*o = OperationResponse(varOperationResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "assignmentOperation")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "numUserMoved")
		delete(additionalProperties, "realmId")
		delete(additionalProperties, "realmName")
		delete(additionalProperties, "started")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOperationResponse struct {
	value *OperationResponse
	isSet bool
}

func (v NullableOperationResponse) Get() *OperationResponse {
	return v.value
}

func (v *NullableOperationResponse) Set(val *OperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationResponse(val *OperationResponse) *NullableOperationResponse {
	return &NullableOperationResponse{value: val, isSet: true}
}

func (v NullableOperationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

