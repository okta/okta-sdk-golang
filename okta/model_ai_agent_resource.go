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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the AIAgentResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AIAgentResource{}

// AIAgentResource The AI agent resource associated with the operation. These properties are available after the operation completes successfully.
type AIAgentResource struct {
	// The ID of the AI agent resource
	Id string `json:"id"`
	// The status of the AI agent resource
	Status string `json:"status"`
	// The type of resource
	Type                 string    `json:"type"`
	Links                LinksSelf `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _AIAgentResource AIAgentResource

// NewAIAgentResource instantiates a new AIAgentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAIAgentResource(id string, status string, type_ string, links LinksSelf) *AIAgentResource {
	this := AIAgentResource{}
	this.Id = id
	this.Status = status
	this.Type = type_
	this.Links = links
	return &this
}

// NewAIAgentResourceWithDefaults instantiates a new AIAgentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAIAgentResourceWithDefaults() *AIAgentResource {
	this := AIAgentResource{}
	return &this
}

// GetId returns the Id field value
func (o *AIAgentResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AIAgentResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AIAgentResource) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *AIAgentResource) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AIAgentResource) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AIAgentResource) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *AIAgentResource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AIAgentResource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AIAgentResource) SetType(v string) {
	o.Type = v
}

// GetLinks returns the Links field value
func (o *AIAgentResource) GetLinks() LinksSelf {
	if o == nil {
		var ret LinksSelf
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AIAgentResource) GetLinksOk() (*LinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AIAgentResource) SetLinks(v LinksSelf) {
	o.Links = v
}

func (o AIAgentResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AIAgentResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AIAgentResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"type",
		"_links",
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

	varAIAgentResource := _AIAgentResource{}

	err = json.Unmarshal(data, &varAIAgentResource)

	if err != nil {
		return err
	}

	*o = AIAgentResource(varAIAgentResource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAIAgentResource struct {
	value *AIAgentResource
	isSet bool
}

func (v NullableAIAgentResource) Get() *AIAgentResource {
	return v.value
}

func (v *NullableAIAgentResource) Set(val *AIAgentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAIAgentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAIAgentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAIAgentResource(val *AIAgentResource) *NullableAIAgentResource {
	return &NullableAIAgentResource{value: val, isSet: true}
}

func (v NullableAIAgentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAIAgentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
