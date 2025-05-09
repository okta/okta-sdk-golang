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
)

// PasswordImportRequest struct for PasswordImportRequest
type PasswordImportRequest struct {
	Data *PasswordImportRequestData `json:"data,omitempty"`
	// The type of inline hook. The password import inline hook type is `com.okta.user.credential.password.import`.
	EventType *string `json:"eventType,omitempty"`
	// The ID and URL of the password import inline hook
	Source *string `json:"source,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordImportRequest PasswordImportRequest

// NewPasswordImportRequest instantiates a new PasswordImportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordImportRequest() *PasswordImportRequest {
	this := PasswordImportRequest{}
	return &this
}

// NewPasswordImportRequestWithDefaults instantiates a new PasswordImportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordImportRequestWithDefaults() *PasswordImportRequest {
	this := PasswordImportRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PasswordImportRequest) GetData() PasswordImportRequestData {
	if o == nil || o.Data == nil {
		var ret PasswordImportRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequest) GetDataOk() (*PasswordImportRequestData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PasswordImportRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PasswordImportRequestData and assigns it to the Data field.
func (o *PasswordImportRequest) SetData(v PasswordImportRequestData) {
	o.Data = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *PasswordImportRequest) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequest) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *PasswordImportRequest) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *PasswordImportRequest) SetEventType(v string) {
	o.EventType = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PasswordImportRequest) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordImportRequest) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PasswordImportRequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PasswordImportRequest) SetSource(v string) {
	o.Source = &v
}

func (o PasswordImportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PasswordImportRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPasswordImportRequest := _PasswordImportRequest{}

	err = json.Unmarshal(bytes, &varPasswordImportRequest)
	if err == nil {
		*o = PasswordImportRequest(varPasswordImportRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "source")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePasswordImportRequest struct {
	value *PasswordImportRequest
	isSet bool
}

func (v NullablePasswordImportRequest) Get() *PasswordImportRequest {
	return v.value
}

func (v *NullablePasswordImportRequest) Set(val *PasswordImportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordImportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordImportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordImportRequest(val *PasswordImportRequest) *NullablePasswordImportRequest {
	return &NullablePasswordImportRequest{value: val, isSet: true}
}

func (v NullablePasswordImportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordImportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

