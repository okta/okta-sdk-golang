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

// GroupOwner struct for GroupOwner
type GroupOwner struct {
	// The display name of the group owner
	DisplayName *string `json:"displayName,omitempty"`
	// The `id` of the group owner
	Id *string `json:"id,omitempty"`
	// Timestamp when the group owner was last updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// The ID of the app instance if the `originType` is `APPLICATION`. This value is `NULL` if `originType` is `OKTA_DIRECTORY`.
	OriginId *string `json:"originId,omitempty"`
	// The source where group ownership is managed
	OriginType *string `json:"originType,omitempty"`
	// If `originType`is APPLICATION, this parameter is set to `FALSE` until the owner’s `originId` is reconciled with an associated Okta ID.
	Resolved *bool `json:"resolved,omitempty"`
	// The entity type of the owner
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupOwner GroupOwner

// NewGroupOwner instantiates a new GroupOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupOwner() *GroupOwner {
	this := GroupOwner{}
	return &this
}

// NewGroupOwnerWithDefaults instantiates a new GroupOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupOwnerWithDefaults() *GroupOwner {
	this := GroupOwner{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GroupOwner) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOwner) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GroupOwner) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GroupOwner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupOwner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOwner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupOwner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupOwner) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GroupOwner) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOwner) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *GroupOwner) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *GroupOwner) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetOriginId returns the OriginId field value if set, zero value otherwise.
func (o *GroupOwner) GetOriginId() string {
	if o == nil || o.OriginId == nil {
		var ret string
		return ret
	}
	return *o.OriginId
}

// GetOriginIdOk returns a tuple with the OriginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOwner) GetOriginIdOk() (*string, bool) {
	if o == nil || o.OriginId == nil {
		return nil, false
	}
	return o.OriginId, true
}

// HasOriginId returns a boolean if a field has been set.
func (o *GroupOwner) HasOriginId() bool {
	if o != nil && o.OriginId != nil {
		return true
	}

	return false
}

// SetOriginId gets a reference to the given string and assigns it to the OriginId field.
func (o *GroupOwner) SetOriginId(v string) {
	o.OriginId = &v
}

// GetOriginType returns the OriginType field value if set, zero value otherwise.
func (o *GroupOwner) GetOriginType() string {
	if o == nil || o.OriginType == nil {
		var ret string
		return ret
	}
	return *o.OriginType
}

// GetOriginTypeOk returns a tuple with the OriginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOwner) GetOriginTypeOk() (*string, bool) {
	if o == nil || o.OriginType == nil {
		return nil, false
	}
	return o.OriginType, true
}

// HasOriginType returns a boolean if a field has been set.
func (o *GroupOwner) HasOriginType() bool {
	if o != nil && o.OriginType != nil {
		return true
	}

	return false
}

// SetOriginType gets a reference to the given string and assigns it to the OriginType field.
func (o *GroupOwner) SetOriginType(v string) {
	o.OriginType = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *GroupOwner) GetResolved() bool {
	if o == nil || o.Resolved == nil {
		var ret bool
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOwner) GetResolvedOk() (*bool, bool) {
	if o == nil || o.Resolved == nil {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *GroupOwner) HasResolved() bool {
	if o != nil && o.Resolved != nil {
		return true
	}

	return false
}

// SetResolved gets a reference to the given bool and assigns it to the Resolved field.
func (o *GroupOwner) SetResolved(v bool) {
	o.Resolved = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupOwner) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupOwner) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupOwner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GroupOwner) SetType(v string) {
	o.Type = &v
}

func (o GroupOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.OriginId != nil {
		toSerialize["originId"] = o.OriginId
	}
	if o.OriginType != nil {
		toSerialize["originType"] = o.OriginType
	}
	if o.Resolved != nil {
		toSerialize["resolved"] = o.Resolved
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GroupOwner) UnmarshalJSON(bytes []byte) (err error) {
	varGroupOwner := _GroupOwner{}

	err = json.Unmarshal(bytes, &varGroupOwner)
	if err == nil {
		*o = GroupOwner(varGroupOwner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "originId")
		delete(additionalProperties, "originType")
		delete(additionalProperties, "resolved")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGroupOwner struct {
	value *GroupOwner
	isSet bool
}

func (v NullableGroupOwner) Get() *GroupOwner {
	return v.value
}

func (v *NullableGroupOwner) Set(val *GroupOwner) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupOwner) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupOwner(val *GroupOwner) *NullableGroupOwner {
	return &NullableGroupOwner{value: val, isSet: true}
}

func (v NullableGroupOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

