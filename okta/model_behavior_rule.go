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

// checks if the BehaviorRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BehaviorRule{}

// BehaviorRule struct for BehaviorRule
type BehaviorRule struct {
	// Timestamp when the Behavior Detection Rule was created
	Created *string `json:"created,omitempty"`
	// ID of the Behavior Detection Rule
	Id *string `json:"id,omitempty"`
	// Timestamp when the Behavior Detection Rule was last modified
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// Name of the Behavior Detection Rule
	Name                 string     `json:"name"`
	Status               *string    `json:"status,omitempty"`
	Type                 string     `json:"type"`
	Link                 *LinksSelf `json:"_link,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BehaviorRule BehaviorRule

// NewBehaviorRule instantiates a new BehaviorRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBehaviorRule(name string, type_ string) *BehaviorRule {
	this := BehaviorRule{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewBehaviorRuleWithDefaults instantiates a new BehaviorRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBehaviorRuleWithDefaults() *BehaviorRule {
	this := BehaviorRule{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *BehaviorRule) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRule) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *BehaviorRule) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *BehaviorRule) SetCreated(v string) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BehaviorRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BehaviorRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BehaviorRule) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *BehaviorRule) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRule) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *BehaviorRule) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *BehaviorRule) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *BehaviorRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BehaviorRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BehaviorRule) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BehaviorRule) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRule) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BehaviorRule) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BehaviorRule) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value
func (o *BehaviorRule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BehaviorRule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BehaviorRule) SetType(v string) {
	o.Type = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *BehaviorRule) GetLink() LinksSelf {
	if o == nil || IsNil(o.Link) {
		var ret LinksSelf
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BehaviorRule) GetLinkOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *BehaviorRule) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given LinksSelf and assigns it to the Link field.
func (o *BehaviorRule) SetLink(v LinksSelf) {
	o.Link = &v
}

func (o BehaviorRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BehaviorRule) ToMap() (map[string]interface{}, error) {
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
	toSerialize["name"] = o.Name
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Link) {
		toSerialize["_link"] = o.Link
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BehaviorRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
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

	varBehaviorRule := _BehaviorRule{}

	err = json.Unmarshal(data, &varBehaviorRule)

	if err != nil {
		return err
	}

	*o = BehaviorRule(varBehaviorRule)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_link")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBehaviorRule struct {
	value *BehaviorRule
	isSet bool
}

func (v NullableBehaviorRule) Get() *BehaviorRule {
	return v.value
}

func (v *NullableBehaviorRule) Set(val *BehaviorRule) {
	v.value = val
	v.isSet = true
}

func (v NullableBehaviorRule) IsSet() bool {
	return v.isSet
}

func (v *NullableBehaviorRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBehaviorRule(val *BehaviorRule) *NullableBehaviorRule {
	return &NullableBehaviorRule{value: val, isSet: true}
}

func (v NullableBehaviorRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBehaviorRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
