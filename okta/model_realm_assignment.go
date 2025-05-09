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

// RealmAssignment struct for RealmAssignment
type RealmAssignment struct {
	Actions *Actions `json:"actions,omitempty"`
	Conditions *Conditions `json:"conditions,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Id *string `json:"id,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Name *string `json:"name,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	Status *string `json:"status,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmAssignment RealmAssignment

// NewRealmAssignment instantiates a new RealmAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmAssignment() *RealmAssignment {
	this := RealmAssignment{}
	return &this
}

// NewRealmAssignmentWithDefaults instantiates a new RealmAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmAssignmentWithDefaults() *RealmAssignment {
	this := RealmAssignment{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *RealmAssignment) GetActions() Actions {
	if o == nil || o.Actions == nil {
		var ret Actions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetActionsOk() (*Actions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *RealmAssignment) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given Actions and assigns it to the Actions field.
func (o *RealmAssignment) SetActions(v Actions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *RealmAssignment) GetConditions() Conditions {
	if o == nil || o.Conditions == nil {
		var ret Conditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetConditionsOk() (*Conditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *RealmAssignment) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given Conditions and assigns it to the Conditions field.
func (o *RealmAssignment) SetConditions(v Conditions) {
	o.Conditions = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RealmAssignment) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RealmAssignment) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RealmAssignment) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RealmAssignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RealmAssignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RealmAssignment) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *RealmAssignment) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *RealmAssignment) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *RealmAssignment) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *RealmAssignment) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *RealmAssignment) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *RealmAssignment) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RealmAssignment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RealmAssignment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RealmAssignment) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RealmAssignment) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RealmAssignment) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *RealmAssignment) SetPriority(v int32) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RealmAssignment) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RealmAssignment) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RealmAssignment) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RealmAssignment) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmAssignment) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RealmAssignment) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *RealmAssignment) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o RealmAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RealmAssignment) UnmarshalJSON(bytes []byte) (err error) {
	varRealmAssignment := _RealmAssignment{}

	err = json.Unmarshal(bytes, &varRealmAssignment)
	if err == nil {
		*o = RealmAssignment(varRealmAssignment)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isDefault")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRealmAssignment struct {
	value *RealmAssignment
	isSet bool
}

func (v NullableRealmAssignment) Get() *RealmAssignment {
	return v.value
}

func (v *NullableRealmAssignment) Set(val *RealmAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmAssignment(val *RealmAssignment) *NullableRealmAssignment {
	return &NullableRealmAssignment{value: val, isSet: true}
}

func (v NullableRealmAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

