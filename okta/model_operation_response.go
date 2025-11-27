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

// checks if the OperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationResponse{}

// OperationResponse struct for OperationResponse
type OperationResponse struct {
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
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _OperationResponse OperationResponse

// NewOperationResponse instantiates a new OperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationResponse(created time.Time, id string, status string, type_ string) *OperationResponse {
	this := OperationResponse{}
	this.Created = created
	this.Id = id
	this.Status = status
	this.Type = type_
	return &this
}

// NewOperationResponseWithDefaults instantiates a new OperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationResponseWithDefaults() *OperationResponse {
	this := OperationResponse{}
	return &this
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

// GetCreated returns the Created field value
func (o *OperationResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *OperationResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *OperationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OperationResponse) SetId(v string) {
	o.Id = v
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

// GetStatus returns the Status field value
func (o *OperationResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OperationResponse) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *OperationResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OperationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OperationResponse) SetType(v string) {
	o.Type = v
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OperationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varOperationResponse := _OperationResponse{}

	err = json.Unmarshal(data, &varOperationResponse)

	if err != nil {
		return err
	}

	*o = OperationResponse(varOperationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "completed")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "started")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
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
