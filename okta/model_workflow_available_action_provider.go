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

// checks if the WorkflowAvailableActionProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowAvailableActionProvider{}

// WorkflowAvailableActionProvider struct for WorkflowAvailableActionProvider
type WorkflowAvailableActionProvider struct {
	// The name of the action flow
	ActionName string `json:"actionName"`
	// The unique identifier of the action flow in the provider system
	ExternalId string `json:"externalId"`
	// Type of action provider
	Type string `json:"type"`
	// The URL to the action flow
	Url                  string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowAvailableActionProvider WorkflowAvailableActionProvider

// NewWorkflowAvailableActionProvider instantiates a new WorkflowAvailableActionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowAvailableActionProvider(actionName string, externalId string, type_ string, url string) *WorkflowAvailableActionProvider {
	this := WorkflowAvailableActionProvider{}
	this.ActionName = actionName
	this.ExternalId = externalId
	this.Type = type_
	this.Url = url
	return &this
}

// NewWorkflowAvailableActionProviderWithDefaults instantiates a new WorkflowAvailableActionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowAvailableActionProviderWithDefaults() *WorkflowAvailableActionProvider {
	this := WorkflowAvailableActionProvider{}
	return &this
}

// GetActionName returns the ActionName field value
func (o *WorkflowAvailableActionProvider) GetActionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionName
}

// GetActionNameOk returns a tuple with the ActionName field value
// and a boolean to check if the value has been set.
func (o *WorkflowAvailableActionProvider) GetActionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionName, true
}

// SetActionName sets field value
func (o *WorkflowAvailableActionProvider) SetActionName(v string) {
	o.ActionName = v
}

// GetExternalId returns the ExternalId field value
func (o *WorkflowAvailableActionProvider) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *WorkflowAvailableActionProvider) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *WorkflowAvailableActionProvider) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *WorkflowAvailableActionProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowAvailableActionProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WorkflowAvailableActionProvider) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *WorkflowAvailableActionProvider) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WorkflowAvailableActionProvider) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WorkflowAvailableActionProvider) SetUrl(v string) {
	o.Url = v
}

func (o WorkflowAvailableActionProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowAvailableActionProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionName"] = o.ActionName
	toSerialize["externalId"] = o.ExternalId
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkflowAvailableActionProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actionName",
		"externalId",
		"type",
		"url",
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

	varWorkflowAvailableActionProvider := _WorkflowAvailableActionProvider{}

	err = json.Unmarshal(data, &varWorkflowAvailableActionProvider)

	if err != nil {
		return err
	}

	*o = WorkflowAvailableActionProvider(varWorkflowAvailableActionProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actionName")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowAvailableActionProvider struct {
	value *WorkflowAvailableActionProvider
	isSet bool
}

func (v NullableWorkflowAvailableActionProvider) Get() *WorkflowAvailableActionProvider {
	return v.value
}

func (v *NullableWorkflowAvailableActionProvider) Set(val *WorkflowAvailableActionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowAvailableActionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowAvailableActionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowAvailableActionProvider(val *WorkflowAvailableActionProvider) *NullableWorkflowAvailableActionProvider {
	return &NullableWorkflowAvailableActionProvider{value: val, isSet: true}
}

func (v NullableWorkflowAvailableActionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowAvailableActionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
