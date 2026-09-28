/*
Okta Admin Management API

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
	"time"
)

// checks if the CimdClientEntityBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CimdClientEntityBinding{}

// CimdClientEntityBinding Represents a binding between a CIMD client entity and another entity, such as an app instance
type CimdClientEntityBinding struct {
	// Unique identifier of the entity that's bound to the CIMD client entity
	BoundEntityId string `json:"boundEntityId"`
	// The type of resource that's bound to the CIMD client entity
	BoundEntityType string `json:"boundEntityType"`
	// Timestamp when the binding was created
	Created *time.Time `json:"created,omitempty"`
	// Unique identifier for the CIMD client entity binding
	Id *string `json:"id,omitempty"`
	// Timestamp when the binding was last updated
	LastUpdated          *time.Time                    `json:"lastUpdated,omitempty"`
	Links                *CimdClientEntityBindingLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CimdClientEntityBinding CimdClientEntityBinding

// NewCimdClientEntityBinding instantiates a new CimdClientEntityBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCimdClientEntityBinding(boundEntityId string, boundEntityType string) *CimdClientEntityBinding {
	this := CimdClientEntityBinding{}
	this.BoundEntityId = boundEntityId
	this.BoundEntityType = boundEntityType
	return &this
}

// NewCimdClientEntityBindingWithDefaults instantiates a new CimdClientEntityBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCimdClientEntityBindingWithDefaults() *CimdClientEntityBinding {
	this := CimdClientEntityBinding{}
	return &this
}

// GetBoundEntityId returns the BoundEntityId field value
func (o *CimdClientEntityBinding) GetBoundEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoundEntityId
}

// GetBoundEntityIdOk returns a tuple with the BoundEntityId field value
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBinding) GetBoundEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundEntityId, true
}

// SetBoundEntityId sets field value
func (o *CimdClientEntityBinding) SetBoundEntityId(v string) {
	o.BoundEntityId = v
}

// GetBoundEntityType returns the BoundEntityType field value
func (o *CimdClientEntityBinding) GetBoundEntityType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoundEntityType
}

// GetBoundEntityTypeOk returns a tuple with the BoundEntityType field value
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBinding) GetBoundEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoundEntityType, true
}

// SetBoundEntityType sets field value
func (o *CimdClientEntityBinding) SetBoundEntityType(v string) {
	o.BoundEntityType = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CimdClientEntityBinding) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBinding) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CimdClientEntityBinding) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CimdClientEntityBinding) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CimdClientEntityBinding) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBinding) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CimdClientEntityBinding) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CimdClientEntityBinding) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CimdClientEntityBinding) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBinding) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *CimdClientEntityBinding) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *CimdClientEntityBinding) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CimdClientEntityBinding) GetLinks() CimdClientEntityBindingLinks {
	if o == nil || IsNil(o.Links) {
		var ret CimdClientEntityBindingLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdClientEntityBinding) GetLinksOk() (*CimdClientEntityBindingLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CimdClientEntityBinding) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CimdClientEntityBindingLinks and assigns it to the Links field.
func (o *CimdClientEntityBinding) SetLinks(v CimdClientEntityBindingLinks) {
	o.Links = &v
}

func (o CimdClientEntityBinding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CimdClientEntityBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boundEntityId"] = o.BoundEntityId
	toSerialize["boundEntityType"] = o.BoundEntityType
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

func (o *CimdClientEntityBinding) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"boundEntityId",
		"boundEntityType",
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

	varCimdClientEntityBinding := _CimdClientEntityBinding{}

	err = json.Unmarshal(data, &varCimdClientEntityBinding)

	if err != nil {
		return err
	}

	*o = CimdClientEntityBinding(varCimdClientEntityBinding)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "boundEntityId")
		delete(additionalProperties, "boundEntityType")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCimdClientEntityBinding struct {
	value *CimdClientEntityBinding
	isSet bool
}

func (v NullableCimdClientEntityBinding) Get() *CimdClientEntityBinding {
	return v.value
}

func (v *NullableCimdClientEntityBinding) Set(val *CimdClientEntityBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableCimdClientEntityBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableCimdClientEntityBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCimdClientEntityBinding(val *CimdClientEntityBinding) *NullableCimdClientEntityBinding {
	return &NullableCimdClientEntityBinding{value: val, isSet: true}
}

func (v NullableCimdClientEntityBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCimdClientEntityBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
