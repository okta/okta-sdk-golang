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

// checks if the AIAgentOperationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIAgentOperationResponse{}

// AIAgentOperationResponse struct for AIAgentOperationResponse
type AIAgentOperationResponse struct {
	// Timestamp of when the AI agent operation completed
	Completed *time.Time `json:"completed,omitempty"`
	// Timestamp of when the AI agent operation was created
	Created      time.Time     `json:"created"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// ID of the AI agent operation
	Id       string           `json:"id"`
	Resource *AIAgentResource `json:"resource,omitempty"`
	// Timestamp of when the AI agent operation started
	Started *time.Time `json:"started,omitempty"`
	// The status of the AI agent operation
	Status string `json:"status"`
	// The AI agent operation type
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _AIAgentOperationResponse AIAgentOperationResponse

// NewAIAgentOperationResponse instantiates a new AIAgentOperationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIAgentOperationResponse(created time.Time, id string, status string, type_ string) *AIAgentOperationResponse {
	this := AIAgentOperationResponse{}
	this.Created = created
	this.Id = id
	this.Status = status
	this.Type = type_
	return &this
}

// NewAIAgentOperationResponseWithDefaults instantiates a new AIAgentOperationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIAgentOperationResponseWithDefaults() *AIAgentOperationResponse {
	this := AIAgentOperationResponse{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *AIAgentOperationResponse) GetCompleted() time.Time {
	if o == nil || IsNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIAgentOperationResponse) GetCompletedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *AIAgentOperationResponse) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *AIAgentOperationResponse) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetCreated returns the Created field value
func (o *AIAgentOperationResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *AIAgentOperationResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *AIAgentOperationResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *AIAgentOperationResponse) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIAgentOperationResponse) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *AIAgentOperationResponse) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *AIAgentOperationResponse) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetId returns the Id field value
func (o *AIAgentOperationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AIAgentOperationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AIAgentOperationResponse) SetId(v string) {
	o.Id = v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *AIAgentOperationResponse) GetResource() AIAgentResource {
	if o == nil || IsNil(o.Resource) {
		var ret AIAgentResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIAgentOperationResponse) GetResourceOk() (*AIAgentResource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *AIAgentOperationResponse) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given AIAgentResource and assigns it to the Resource field.
func (o *AIAgentOperationResponse) SetResource(v AIAgentResource) {
	o.Resource = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *AIAgentOperationResponse) GetStarted() time.Time {
	if o == nil || IsNil(o.Started) {
		var ret time.Time
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AIAgentOperationResponse) GetStartedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *AIAgentOperationResponse) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given time.Time and assigns it to the Started field.
func (o *AIAgentOperationResponse) SetStarted(v time.Time) {
	o.Started = &v
}

// GetStatus returns the Status field value
func (o *AIAgentOperationResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AIAgentOperationResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AIAgentOperationResponse) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *AIAgentOperationResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AIAgentOperationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AIAgentOperationResponse) SetType(v string) {
	o.Type = v
}

func (o AIAgentOperationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIAgentOperationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	toSerialize["created"] = o.Created
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
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

func (o *AIAgentOperationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAIAgentOperationResponse := _AIAgentOperationResponse{}

	err = json.Unmarshal(data, &varAIAgentOperationResponse)

	if err != nil {
		return err
	}

	*o = AIAgentOperationResponse(varAIAgentOperationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "completed")
		delete(additionalProperties, "created")
		delete(additionalProperties, "errorDetails")
		delete(additionalProperties, "id")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "started")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIAgentOperationResponse struct {
	value *AIAgentOperationResponse
	isSet bool
}

func (v NullableAIAgentOperationResponse) Get() *AIAgentOperationResponse {
	return v.value
}

func (v *NullableAIAgentOperationResponse) Set(val *AIAgentOperationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAIAgentOperationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAIAgentOperationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIAgentOperationResponse(val *AIAgentOperationResponse) *NullableAIAgentOperationResponse {
	return &NullableAIAgentOperationResponse{value: val, isSet: true}
}

func (v NullableAIAgentOperationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIAgentOperationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
