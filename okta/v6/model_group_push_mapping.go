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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// GroupPushMapping struct for GroupPushMapping
type GroupPushMapping struct {
	AppConfig AppConfig `json:"appConfig,omitempty"`
	// Timestamp when the group push mapping was created
	Created *time.Time `json:"created,omitempty"`
	// The error message summary if the latest push failed
	ErrorSummary *string `json:"errorSummary,omitempty"`
	// The ID of the group push mapping
	Id *string `json:"id,omitempty"`
	// Timestamp when the group push mapping was pushed
	LastPush *time.Time `json:"lastPush,omitempty"`
	// Timestamp when the group push mapping was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The ID of the source group for the group push mapping
	SourceGroupId *string `json:"sourceGroupId,omitempty"`
	// The status of the group push mapping
	Status *string `json:"status,omitempty"`
	// The ID of the target group for the group push mapping
	TargetGroupId        *string                `json:"targetGroupId,omitempty"`
	Links                *GroupPushMappingLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupPushMapping GroupPushMapping

// NewGroupPushMapping instantiates a new GroupPushMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPushMapping() *GroupPushMapping {
	this := GroupPushMapping{}
	return &this
}

// NewGroupPushMappingWithDefaults instantiates a new GroupPushMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPushMappingWithDefaults() *GroupPushMapping {
	this := GroupPushMapping{}
	return &this
}

// GetAppConfig returns the AppConfig field value if set, zero value otherwise.
func (o *GroupPushMapping) GetAppConfig() AppConfig {
	if o == nil || o.AppConfig == nil {
		var ret AppConfig
		return ret
	}
	return o.AppConfig
}

// GetAppConfigOk returns a tuple with the AppConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetAppConfigOk() (AppConfig, bool) {
	if o == nil || o.AppConfig == nil {
		return nil, false
	}
	return o.AppConfig, true
}

// HasAppConfig returns a boolean if a field has been set.
func (o *GroupPushMapping) HasAppConfig() bool {
	if o != nil && o.AppConfig != nil {
		return true
	}

	return false
}

// SetAppConfig gets a reference to the given AppConfig and assigns it to the AppConfig field.
func (o *GroupPushMapping) SetAppConfig(v AppConfig) {
	o.AppConfig = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GroupPushMapping) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GroupPushMapping) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GroupPushMapping) SetCreated(v time.Time) {
	o.Created = &v
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *GroupPushMapping) GetErrorSummary() string {
	if o == nil || o.ErrorSummary == nil {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetErrorSummaryOk() (*string, bool) {
	if o == nil || o.ErrorSummary == nil {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *GroupPushMapping) HasErrorSummary() bool {
	if o != nil && o.ErrorSummary != nil {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *GroupPushMapping) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupPushMapping) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupPushMapping) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupPushMapping) SetId(v string) {
	o.Id = &v
}

// GetLastPush returns the LastPush field value if set, zero value otherwise.
func (o *GroupPushMapping) GetLastPush() time.Time {
	if o == nil || o.LastPush == nil {
		var ret time.Time
		return ret
	}
	return *o.LastPush
}

// GetLastPushOk returns a tuple with the LastPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetLastPushOk() (*time.Time, bool) {
	if o == nil || o.LastPush == nil {
		return nil, false
	}
	return o.LastPush, true
}

// HasLastPush returns a boolean if a field has been set.
func (o *GroupPushMapping) HasLastPush() bool {
	if o != nil && o.LastPush != nil {
		return true
	}

	return false
}

// SetLastPush gets a reference to the given time.Time and assigns it to the LastPush field.
func (o *GroupPushMapping) SetLastPush(v time.Time) {
	o.LastPush = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GroupPushMapping) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *GroupPushMapping) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *GroupPushMapping) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetSourceGroupId returns the SourceGroupId field value if set, zero value otherwise.
func (o *GroupPushMapping) GetSourceGroupId() string {
	if o == nil || o.SourceGroupId == nil {
		var ret string
		return ret
	}
	return *o.SourceGroupId
}

// GetSourceGroupIdOk returns a tuple with the SourceGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetSourceGroupIdOk() (*string, bool) {
	if o == nil || o.SourceGroupId == nil {
		return nil, false
	}
	return o.SourceGroupId, true
}

// HasSourceGroupId returns a boolean if a field has been set.
func (o *GroupPushMapping) HasSourceGroupId() bool {
	if o != nil && o.SourceGroupId != nil {
		return true
	}

	return false
}

// SetSourceGroupId gets a reference to the given string and assigns it to the SourceGroupId field.
func (o *GroupPushMapping) SetSourceGroupId(v string) {
	o.SourceGroupId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GroupPushMapping) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GroupPushMapping) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GroupPushMapping) SetStatus(v string) {
	o.Status = &v
}

// GetTargetGroupId returns the TargetGroupId field value if set, zero value otherwise.
func (o *GroupPushMapping) GetTargetGroupId() string {
	if o == nil || o.TargetGroupId == nil {
		var ret string
		return ret
	}
	return *o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetTargetGroupIdOk() (*string, bool) {
	if o == nil || o.TargetGroupId == nil {
		return nil, false
	}
	return o.TargetGroupId, true
}

// HasTargetGroupId returns a boolean if a field has been set.
func (o *GroupPushMapping) HasTargetGroupId() bool {
	if o != nil && o.TargetGroupId != nil {
		return true
	}

	return false
}

// SetTargetGroupId gets a reference to the given string and assigns it to the TargetGroupId field.
func (o *GroupPushMapping) SetTargetGroupId(v string) {
	o.TargetGroupId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GroupPushMapping) GetLinks() GroupPushMappingLinks {
	if o == nil || o.Links == nil {
		var ret GroupPushMappingLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPushMapping) GetLinksOk() (*GroupPushMappingLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GroupPushMapping) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GroupPushMappingLinks and assigns it to the Links field.
func (o *GroupPushMapping) SetLinks(v GroupPushMappingLinks) {
	o.Links = &v
}

func (o GroupPushMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppConfig != nil {
		toSerialize["appConfig"] = o.AppConfig
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.ErrorSummary != nil {
		toSerialize["errorSummary"] = o.ErrorSummary
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastPush != nil {
		toSerialize["lastPush"] = o.LastPush
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.SourceGroupId != nil {
		toSerialize["sourceGroupId"] = o.SourceGroupId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TargetGroupId != nil {
		toSerialize["targetGroupId"] = o.TargetGroupId
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GroupPushMapping) UnmarshalJSON(bytes []byte) (err error) {
	varGroupPushMapping := _GroupPushMapping{}

	err = json.Unmarshal(bytes, &varGroupPushMapping)
	if err == nil {
		*o = GroupPushMapping(varGroupPushMapping)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appConfig")
		delete(additionalProperties, "created")
		delete(additionalProperties, "errorSummary")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastPush")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "sourceGroupId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "targetGroupId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGroupPushMapping struct {
	value *GroupPushMapping
	isSet bool
}

func (v NullableGroupPushMapping) Get() *GroupPushMapping {
	return v.value
}

func (v *NullableGroupPushMapping) Set(val *GroupPushMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPushMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPushMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPushMapping(val *GroupPushMapping) *NullableGroupPushMapping {
	return &NullableGroupPushMapping{value: val, isSet: true}
}

func (v NullableGroupPushMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPushMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
