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

// AuthorizationServerPolicy struct for AuthorizationServerPolicy
type AuthorizationServerPolicy struct {
	// ID of the Policy
	Id *string `json:"id,omitempty"`
	// Indicates that the Policy is an authorization server Policy
	Type *string `json:"type,omitempty"`
	// Name of the Policy
	Name *string `json:"name,omitempty"`
	Conditions *AuthorizationServerPolicyConditions `json:"conditions,omitempty"`
	// Description of the Policy
	Description *string `json:"description,omitempty"`
	// Specifies the order in which this Policy is evaluated in relation to the other Policies in a custom authorization server
	Priority *int32 `json:"priority,omitempty"`
	// Specifies whether requests have access to this Policy
	Status *string `json:"status,omitempty"`
	// Specifies whether Okta created this Policy
	System *bool `json:"system,omitempty"`
	// Timestamp when the Policy was created
	Created *time.Time `json:"created,omitempty"`
	// Timestamp when the Policy was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Links *AuthorizationServerPolicyAllOfLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicy AuthorizationServerPolicy

// NewAuthorizationServerPolicy instantiates a new AuthorizationServerPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicy() *AuthorizationServerPolicy {
	this := AuthorizationServerPolicy{}
	return &this
}

// NewAuthorizationServerPolicyWithDefaults instantiates a new AuthorizationServerPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyWithDefaults() *AuthorizationServerPolicy {
	this := AuthorizationServerPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthorizationServerPolicy) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AuthorizationServerPolicy) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthorizationServerPolicy) SetName(v string) {
	o.Name = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetConditions() AuthorizationServerPolicyConditions {
	if o == nil || o.Conditions == nil {
		var ret AuthorizationServerPolicyConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetConditionsOk() (*AuthorizationServerPolicyConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given AuthorizationServerPolicyConditions and assigns it to the Conditions field.
func (o *AuthorizationServerPolicy) SetConditions(v AuthorizationServerPolicyConditions) {
	o.Conditions = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizationServerPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *AuthorizationServerPolicy) SetPriority(v int32) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthorizationServerPolicy) SetStatus(v string) {
	o.Status = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetSystem() bool {
	if o == nil || o.System == nil {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetSystemOk() (*bool, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *AuthorizationServerPolicy) SetSystem(v bool) {
	o.System = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AuthorizationServerPolicy) SetCreated(v time.Time) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AuthorizationServerPolicy) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizationServerPolicy) GetLinks() AuthorizationServerPolicyAllOfLinks {
	if o == nil || o.Links == nil {
		var ret AuthorizationServerPolicyAllOfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicy) GetLinksOk() (*AuthorizationServerPolicyAllOfLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizationServerPolicy) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AuthorizationServerPolicyAllOfLinks and assigns it to the Links field.
func (o *AuthorizationServerPolicy) SetLinks(v AuthorizationServerPolicyAllOfLinks) {
	o.Links = &v
}

func (o AuthorizationServerPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.System != nil {
		toSerialize["system"] = o.System
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicy) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicy := _AuthorizationServerPolicy{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicy)
	if err == nil {
		*o = AuthorizationServerPolicy(varAuthorizationServerPolicy)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "description")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "status")
		delete(additionalProperties, "system")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicy struct {
	value *AuthorizationServerPolicy
	isSet bool
}

func (v NullableAuthorizationServerPolicy) Get() *AuthorizationServerPolicy {
	return v.value
}

func (v *NullableAuthorizationServerPolicy) Set(val *AuthorizationServerPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicy(val *AuthorizationServerPolicy) *NullableAuthorizationServerPolicy {
	return &NullableAuthorizationServerPolicy{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

