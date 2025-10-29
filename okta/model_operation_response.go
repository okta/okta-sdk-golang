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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// checks if the OperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationResponse{}

// OperationResponse struct for OperationResponse
type OperationResponse struct {
	AssignmentOperation *OperationResponseAssignmentOperation `json:"assignmentOperation,omitempty"`
	// Timestamp when the realm assignment operation completed
	Completed *time.Time `json:"completed,omitempty"`
	// Timestamp when the realm assignment operation was created
	Created *time.Time `json:"created,omitempty"`
	// ID of the realm
	Id *string `json:"id,omitempty"`
	// Number of users moved
	NumUserMoved *float32 `json:"numUserMoved,omitempty"`
	// ID of the realm
	RealmId *string `json:"realmId,omitempty"`
	// Name of the realm
	RealmName *string `json:"realmName,omitempty"`
	// Timestamp when the realm assignment operation started
	Started *time.Time `json:"started,omitempty"`
	// Current status of the operation
	Status *string `json:"status,omitempty"`
	// Realm type
	Type                 *string    `json:"type,omitempty"`
	Links                *LinksSelf `json:"_links,omitempty"`
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
	if o == nil || IsNil(o.AssignmentOperation) {
		var ret OperationResponseAssignmentOperation
		return ret
	}
	return *o.AssignmentOperation
}

// GetAssignmentOperationOk returns a tuple with the AssignmentOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetAssignmentOperationOk() (*OperationResponseAssignmentOperation, bool) {
	if o == nil || IsNil(o.AssignmentOperation) {
		return nil, false
	}
	return o.AssignmentOperation, true
}

// HasAssignmentOperation returns a boolean if a field has been set.
func (o *OperationResponse) HasAssignmentOperation() bool {
	if o != nil && !IsNil(o.AssignmentOperation) {
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
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *OperationResponse) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
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
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OperationResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
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
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OperationResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
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
	if o == nil || IsNil(o.NumUserMoved) {
		var ret float32
		return ret
	}
	return *o.NumUserMoved
}

// GetNumUserMovedOk returns a tuple with the NumUserMoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetNumUserMovedOk() (*float32, bool) {
	if o == nil || IsNil(o.NumUserMoved) {
		return nil, false
	}
	return o.NumUserMoved, true
}

// HasNumUserMoved returns a boolean if a field has been set.
func (o *OperationResponse) HasNumUserMoved() bool {
	if o != nil && !IsNil(o.NumUserMoved) {
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
	if o == nil || IsNil(o.RealmId) {
		var ret string
		return ret
	}
	return *o.RealmId
}

// GetRealmIdOk returns a tuple with the RealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetRealmIdOk() (*string, bool) {
	if o == nil || IsNil(o.RealmId) {
		return nil, false
	}
	return o.RealmId, true
}

// HasRealmId returns a boolean if a field has been set.
func (o *OperationResponse) HasRealmId() bool {
	if o != nil && !IsNil(o.RealmId) {
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
	if o == nil || IsNil(o.RealmName) {
		var ret string
		return ret
	}
	return *o.RealmName
}

// GetRealmNameOk returns a tuple with the RealmName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetRealmNameOk() (*string, bool) {
	if o == nil || IsNil(o.RealmName) {
		return nil, false
	}
	return o.RealmName, true
}

// HasRealmName returns a boolean if a field has been set.
func (o *OperationResponse) HasRealmName() bool {
	if o != nil && !IsNil(o.RealmName) {
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
	if o == nil || IsNil(o.Started) {
		var ret time.Time
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetStartedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *OperationResponse) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
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
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OperationResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
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
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OperationResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
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
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OperationResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *OperationResponse) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o OperationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignmentOperation) {
		toSerialize["assignmentOperation"] = o.AssignmentOperation
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !IsNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OperationResponse) UnmarshalJSON(data []byte) (err error) {
	varOperationResponse := _OperationResponse{}

	err = json.Unmarshal(data, &varOperationResponse)

	if err != nil {
		return err
	}

	*o = OperationResponse(varOperationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
