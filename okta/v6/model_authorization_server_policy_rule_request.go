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
	"fmt"
)

// AuthorizationServerPolicyRuleRequest struct for AuthorizationServerPolicyRuleRequest
type AuthorizationServerPolicyRuleRequest struct {
	Actions *AuthorizationServerPolicyRuleActions `json:"actions,omitempty"`
	Conditions AuthorizationServerPolicyRuleConditions `json:"conditions"`
	// Timestamp when the rule was created
	Created *time.Time `json:"created,omitempty"`
	// Identifier of the rule
	Id *string `json:"id,omitempty"`
	// Timestamp when the rule was last modified
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Name of the rule
	Name string `json:"name"`
	// Priority of the rule
	Priority *int32 `json:"priority,omitempty"`
	// Status of the rule
	Status *string `json:"status,omitempty"`
	// Set to `true` for system rules. You can't delete system rules.
	System *bool `json:"system,omitempty"`
	// Rule type
	Type string `json:"type"`
	Links *LinksSelfAndLifecycle `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyRuleRequest AuthorizationServerPolicyRuleRequest

// NewAuthorizationServerPolicyRuleRequest instantiates a new AuthorizationServerPolicyRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyRuleRequest(conditions AuthorizationServerPolicyRuleConditions, name string, type_ string) *AuthorizationServerPolicyRuleRequest {
	this := AuthorizationServerPolicyRuleRequest{}
	this.Conditions = conditions
	this.Name = name
	this.Type = type_
	return &this
}

// NewAuthorizationServerPolicyRuleRequestWithDefaults instantiates a new AuthorizationServerPolicyRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyRuleRequestWithDefaults() *AuthorizationServerPolicyRuleRequest {
	this := AuthorizationServerPolicyRuleRequest{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleRequest) GetActions() AuthorizationServerPolicyRuleActions {
	if o == nil || o.Actions == nil {
		var ret AuthorizationServerPolicyRuleActions
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetActionsOk() (*AuthorizationServerPolicyRuleActions, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleRequest) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given AuthorizationServerPolicyRuleActions and assigns it to the Actions field.
func (o *AuthorizationServerPolicyRuleRequest) SetActions(v AuthorizationServerPolicyRuleActions) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value
func (o *AuthorizationServerPolicyRuleRequest) GetConditions() AuthorizationServerPolicyRuleConditions {
	if o == nil {
		var ret AuthorizationServerPolicyRuleConditions
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetConditionsOk() (*AuthorizationServerPolicyRuleConditions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value
func (o *AuthorizationServerPolicyRuleRequest) SetConditions(v AuthorizationServerPolicyRuleConditions) {
	o.Conditions = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleRequest) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleRequest) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AuthorizationServerPolicyRuleRequest) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthorizationServerPolicyRuleRequest) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleRequest) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleRequest) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AuthorizationServerPolicyRuleRequest) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *AuthorizationServerPolicyRuleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthorizationServerPolicyRuleRequest) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleRequest) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleRequest) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *AuthorizationServerPolicyRuleRequest) SetPriority(v int32) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AuthorizationServerPolicyRuleRequest) SetStatus(v string) {
	o.Status = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleRequest) GetSystem() bool {
	if o == nil || o.System == nil {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetSystemOk() (*bool, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleRequest) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *AuthorizationServerPolicyRuleRequest) SetSystem(v bool) {
	o.System = &v
}

// GetType returns the Type field value
func (o *AuthorizationServerPolicyRuleRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizationServerPolicyRuleRequest) SetType(v string) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyRuleRequest) GetLinks() LinksSelfAndLifecycle {
	if o == nil || o.Links == nil {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyRuleRequest) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyRuleRequest) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *AuthorizationServerPolicyRuleRequest) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o AuthorizationServerPolicyRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if true {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["name"] = o.Name
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
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicyRuleRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicyRuleRequest := _AuthorizationServerPolicyRuleRequest{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyRuleRequest)
	if err == nil {
		*o = AuthorizationServerPolicyRuleRequest(varAuthorizationServerPolicyRuleRequest)
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
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "status")
		delete(additionalProperties, "system")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicyRuleRequest struct {
	value *AuthorizationServerPolicyRuleRequest
	isSet bool
}

func (v NullableAuthorizationServerPolicyRuleRequest) Get() *AuthorizationServerPolicyRuleRequest {
	return v.value
}

func (v *NullableAuthorizationServerPolicyRuleRequest) Set(val *AuthorizationServerPolicyRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyRuleRequest(val *AuthorizationServerPolicyRuleRequest) *NullableAuthorizationServerPolicyRuleRequest {
	return &NullableAuthorizationServerPolicyRuleRequest{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

