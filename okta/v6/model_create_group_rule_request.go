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
)

// CreateGroupRuleRequest struct for CreateGroupRuleRequest
type CreateGroupRuleRequest struct {
	Actions *GroupRuleAction `json:"actions,omitempty"`
	Conditions *GroupRuleConditions `json:"conditions,omitempty"`
	// Name of the group rule
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateGroupRuleRequest CreateGroupRuleRequest

// NewCreateGroupRuleRequest instantiates a new CreateGroupRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGroupRuleRequest() *CreateGroupRuleRequest {
	this := CreateGroupRuleRequest{}
	return &this
}

// NewCreateGroupRuleRequestWithDefaults instantiates a new CreateGroupRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGroupRuleRequestWithDefaults() *CreateGroupRuleRequest {
	this := CreateGroupRuleRequest{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CreateGroupRuleRequest) GetActions() GroupRuleAction {
	if o == nil || o.Actions == nil {
		var ret GroupRuleAction
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupRuleRequest) GetActionsOk() (*GroupRuleAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CreateGroupRuleRequest) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given GroupRuleAction and assigns it to the Actions field.
func (o *CreateGroupRuleRequest) SetActions(v GroupRuleAction) {
	o.Actions = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *CreateGroupRuleRequest) GetConditions() GroupRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret GroupRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupRuleRequest) GetConditionsOk() (*GroupRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *CreateGroupRuleRequest) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given GroupRuleConditions and assigns it to the Conditions field.
func (o *CreateGroupRuleRequest) SetConditions(v GroupRuleConditions) {
	o.Conditions = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateGroupRuleRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupRuleRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateGroupRuleRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateGroupRuleRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateGroupRuleRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGroupRuleRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateGroupRuleRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateGroupRuleRequest) SetType(v string) {
	o.Type = &v
}

func (o CreateGroupRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateGroupRuleRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateGroupRuleRequest := _CreateGroupRuleRequest{}

	err = json.Unmarshal(bytes, &varCreateGroupRuleRequest)
	if err == nil {
		*o = CreateGroupRuleRequest(varCreateGroupRuleRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCreateGroupRuleRequest struct {
	value *CreateGroupRuleRequest
	isSet bool
}

func (v NullableCreateGroupRuleRequest) Get() *CreateGroupRuleRequest {
	return v.value
}

func (v *NullableCreateGroupRuleRequest) Set(val *CreateGroupRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGroupRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGroupRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGroupRuleRequest(val *CreateGroupRuleRequest) *NullableCreateGroupRuleRequest {
	return &NullableCreateGroupRuleRequest{value: val, isSet: true}
}

func (v NullableCreateGroupRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGroupRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

