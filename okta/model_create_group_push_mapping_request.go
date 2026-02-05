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
)

// checks if the CreateGroupPushMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGroupPushMappingRequest{}

// CreateGroupPushMappingRequest struct for CreateGroupPushMappingRequest
type CreateGroupPushMappingRequest struct {
	AppConfig *AppConfig `json:"appConfig,omitempty"`
	// The ID of the source group for the group push mapping
	SourceGroupId string `json:"sourceGroupId"`
	// The status of the group push mapping
	Status *string `json:"status,omitempty"`
	// The ID of the existing target group for the group push mapping. This is used to link to an existing group. Required if `targetGroupName` is not provided.
	TargetGroupId *string `json:"targetGroupId,omitempty"`
	// The name of the target group for the group push mapping. This is used when creating a new downstream group. If the group already exists, it links to the existing group. Required if `targetGroupId` is not provided.
	TargetGroupName      *string `json:"targetGroupName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateGroupPushMappingRequest CreateGroupPushMappingRequest

// NewCreateGroupPushMappingRequest instantiates a new CreateGroupPushMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupPushMappingRequest(sourceGroupId string) *CreateGroupPushMappingRequest {
	this := CreateGroupPushMappingRequest{}
	this.SourceGroupId = sourceGroupId
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewCreateGroupPushMappingRequestWithDefaults instantiates a new CreateGroupPushMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupPushMappingRequestWithDefaults() *CreateGroupPushMappingRequest {
	this := CreateGroupPushMappingRequest{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetAppConfig returns the AppConfig field value if set, zero value otherwise.
func (o *CreateGroupPushMappingRequest) GetAppConfig() AppConfig {
	if o == nil || IsNil(o.AppConfig) {
		var ret AppConfig
		return ret
	}
	return *o.AppConfig
}

// GetAppConfigOk returns a tuple with the AppConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupPushMappingRequest) GetAppConfigOk() (*AppConfig, bool) {
	if o == nil || IsNil(o.AppConfig) {
		return nil, false
	}
	return o.AppConfig, true
}

// HasAppConfig returns a boolean if a field has been set.
func (o *CreateGroupPushMappingRequest) HasAppConfig() bool {
	if o != nil && !IsNil(o.AppConfig) {
		return true
	}

	return false
}

// SetAppConfig gets a reference to the given AppConfig and assigns it to the AppConfig field.
func (o *CreateGroupPushMappingRequest) SetAppConfig(v AppConfig) {
	o.AppConfig = &v
}

// GetSourceGroupId returns the SourceGroupId field value
func (o *CreateGroupPushMappingRequest) GetSourceGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceGroupId
}

// GetSourceGroupIdOk returns a tuple with the SourceGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateGroupPushMappingRequest) GetSourceGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceGroupId, true
}

// SetSourceGroupId sets field value
func (o *CreateGroupPushMappingRequest) SetSourceGroupId(v string) {
	o.SourceGroupId = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateGroupPushMappingRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupPushMappingRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateGroupPushMappingRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateGroupPushMappingRequest) SetStatus(v string) {
	o.Status = &v
}

// GetTargetGroupId returns the TargetGroupId field value if set, zero value otherwise.
func (o *CreateGroupPushMappingRequest) GetTargetGroupId() string {
	if o == nil || IsNil(o.TargetGroupId) {
		var ret string
		return ret
	}
	return *o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupPushMappingRequest) GetTargetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetGroupId) {
		return nil, false
	}
	return o.TargetGroupId, true
}

// HasTargetGroupId returns a boolean if a field has been set.
func (o *CreateGroupPushMappingRequest) HasTargetGroupId() bool {
	if o != nil && !IsNil(o.TargetGroupId) {
		return true
	}

	return false
}

// SetTargetGroupId gets a reference to the given string and assigns it to the TargetGroupId field.
func (o *CreateGroupPushMappingRequest) SetTargetGroupId(v string) {
	o.TargetGroupId = &v
}

// GetTargetGroupName returns the TargetGroupName field value if set, zero value otherwise.
func (o *CreateGroupPushMappingRequest) GetTargetGroupName() string {
	if o == nil || IsNil(o.TargetGroupName) {
		var ret string
		return ret
	}
	return *o.TargetGroupName
}

// GetTargetGroupNameOk returns a tuple with the TargetGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupPushMappingRequest) GetTargetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetGroupName) {
		return nil, false
	}
	return o.TargetGroupName, true
}

// HasTargetGroupName returns a boolean if a field has been set.
func (o *CreateGroupPushMappingRequest) HasTargetGroupName() bool {
	if o != nil && !IsNil(o.TargetGroupName) {
		return true
	}

	return false
}

// SetTargetGroupName gets a reference to the given string and assigns it to the TargetGroupName field.
func (o *CreateGroupPushMappingRequest) SetTargetGroupName(v string) {
	o.TargetGroupName = &v
}

func (o CreateGroupPushMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGroupPushMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppConfig) {
		toSerialize["appConfig"] = o.AppConfig
	}
	toSerialize["sourceGroupId"] = o.SourceGroupId
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TargetGroupId) {
		toSerialize["targetGroupId"] = o.TargetGroupId
	}
	if !IsNil(o.TargetGroupName) {
		toSerialize["targetGroupName"] = o.TargetGroupName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateGroupPushMappingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sourceGroupId",
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

	varCreateGroupPushMappingRequest := _CreateGroupPushMappingRequest{}

	err = json.Unmarshal(data, &varCreateGroupPushMappingRequest)

	if err != nil {
		return err
	}

	*o = CreateGroupPushMappingRequest(varCreateGroupPushMappingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appConfig")
		delete(additionalProperties, "sourceGroupId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "targetGroupId")
		delete(additionalProperties, "targetGroupName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateGroupPushMappingRequest struct {
	value *CreateGroupPushMappingRequest
	isSet bool
}

func (v NullableCreateGroupPushMappingRequest) Get() *CreateGroupPushMappingRequest {
	return v.value
}

func (v *NullableCreateGroupPushMappingRequest) Set(val *CreateGroupPushMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupPushMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupPushMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupPushMappingRequest(val *CreateGroupPushMappingRequest) *NullableCreateGroupPushMappingRequest {
	return &NullableCreateGroupPushMappingRequest{value: val, isSet: true}
}

func (v NullableCreateGroupPushMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupPushMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
