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

// checks if the UserImportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportRequest{}

// UserImportRequest struct for UserImportRequest
type UserImportRequest struct {
	Data *UserImportRequestData `json:"data,omitempty"`
	// The type of inline hook. The user import inline hook type is `com.okta.import.transform`.
	EventType *string `json:"eventType,omitempty"`
	// The ID of the user import inline hook
	Source               *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserImportRequest UserImportRequest

// NewUserImportRequest instantiates a new UserImportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportRequest() *UserImportRequest {
	this := UserImportRequest{}
	return &this
}

// NewUserImportRequestWithDefaults instantiates a new UserImportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportRequestWithDefaults() *UserImportRequest {
	this := UserImportRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserImportRequest) GetData() UserImportRequestData {
	if o == nil || IsNil(o.Data) {
		var ret UserImportRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequest) GetDataOk() (*UserImportRequestData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserImportRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UserImportRequestData and assigns it to the Data field.
func (o *UserImportRequest) SetData(v UserImportRequestData) {
	o.Data = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *UserImportRequest) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequest) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *UserImportRequest) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *UserImportRequest) SetEventType(v string) {
	o.EventType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *UserImportRequest) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequest) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *UserImportRequest) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *UserImportRequest) SetSource(v string) {
	o.Source = &v
}

func (o UserImportRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserImportRequest) UnmarshalJSON(data []byte) (err error) {
	varUserImportRequest := _UserImportRequest{}

	err = json.Unmarshal(data, &varUserImportRequest)

	if err != nil {
		return err
	}

	*o = UserImportRequest(varUserImportRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserImportRequest struct {
	value *UserImportRequest
	isSet bool
}

func (v NullableUserImportRequest) Get() *UserImportRequest {
	return v.value
}

func (v *NullableUserImportRequest) Set(val *UserImportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportRequest(val *UserImportRequest) *NullableUserImportRequest {
	return &NullableUserImportRequest{value: val, isSet: true}
}

func (v NullableUserImportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
