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
)

// checks if the AgentJsonWebKeyResponseBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentJsonWebKeyResponseBase{}

// AgentJsonWebKeyResponseBase struct for AgentJsonWebKeyResponseBase
type AgentJsonWebKeyResponseBase struct {
	// Timestamp of when the AI agent JSON Web Key was created
	Created *string `json:"created,omitempty"`
	// The unique ID of the AI agent JSON Web Key
	Id *string `json:"id,omitempty"`
	// Timestamp of when the AI agent JSON Web Key was last updated
	LastUpdated          *string           `json:"lastUpdated,omitempty"`
	Links                *AgentSecretLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentJsonWebKeyResponseBase AgentJsonWebKeyResponseBase

// NewAgentJsonWebKeyResponseBase instantiates a new AgentJsonWebKeyResponseBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentJsonWebKeyResponseBase() *AgentJsonWebKeyResponseBase {
	this := AgentJsonWebKeyResponseBase{}
	return &this
}

// NewAgentJsonWebKeyResponseBaseWithDefaults instantiates a new AgentJsonWebKeyResponseBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentJsonWebKeyResponseBaseWithDefaults() *AgentJsonWebKeyResponseBase {
	this := AgentJsonWebKeyResponseBase{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AgentJsonWebKeyResponseBase) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyResponseBase) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AgentJsonWebKeyResponseBase) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *AgentJsonWebKeyResponseBase) SetCreated(v string) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentJsonWebKeyResponseBase) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyResponseBase) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentJsonWebKeyResponseBase) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentJsonWebKeyResponseBase) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AgentJsonWebKeyResponseBase) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyResponseBase) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AgentJsonWebKeyResponseBase) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *AgentJsonWebKeyResponseBase) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgentJsonWebKeyResponseBase) GetLinks() AgentSecretLinks {
	if o == nil || IsNil(o.Links) {
		var ret AgentSecretLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyResponseBase) GetLinksOk() (*AgentSecretLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgentJsonWebKeyResponseBase) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AgentSecretLinks and assigns it to the Links field.
func (o *AgentJsonWebKeyResponseBase) SetLinks(v AgentSecretLinks) {
	o.Links = &v
}

func (o AgentJsonWebKeyResponseBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentJsonWebKeyResponseBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentJsonWebKeyResponseBase) UnmarshalJSON(data []byte) (err error) {
	varAgentJsonWebKeyResponseBase := _AgentJsonWebKeyResponseBase{}

	err = json.Unmarshal(data, &varAgentJsonWebKeyResponseBase)

	if err != nil {
		return err
	}

	*o = AgentJsonWebKeyResponseBase(varAgentJsonWebKeyResponseBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentJsonWebKeyResponseBase struct {
	value *AgentJsonWebKeyResponseBase
	isSet bool
}

func (v NullableAgentJsonWebKeyResponseBase) Get() *AgentJsonWebKeyResponseBase {
	return v.value
}

func (v *NullableAgentJsonWebKeyResponseBase) Set(val *AgentJsonWebKeyResponseBase) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentJsonWebKeyResponseBase) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentJsonWebKeyResponseBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentJsonWebKeyResponseBase(val *AgentJsonWebKeyResponseBase) *NullableAgentJsonWebKeyResponseBase {
	return &NullableAgentJsonWebKeyResponseBase{value: val, isSet: true}
}

func (v NullableAgentJsonWebKeyResponseBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentJsonWebKeyResponseBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
