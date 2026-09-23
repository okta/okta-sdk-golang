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

// checks if the CimdDocumentRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CimdDocumentRequirements{}

// CimdDocumentRequirements Defines a named set of metadata requirements that CIMD clients matching a CIMD client entity must satisfy
type CimdDocumentRequirements struct {
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	// Description of the CIMD requirement
	Description *string `json:"description,omitempty"`
	// Unique identifier for the CIMD requirement
	Id *string `json:"id,omitempty"`
	// Timestamp when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Human-readable name for the CIMD requirement
	Name string `json:"name"`
	// List of requirement entries. Currently, you can only have one entry.
	Requirements         []CimdRequirementEntry `json:"requirements"`
	Links                *LinksSelf             `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CimdDocumentRequirements CimdDocumentRequirements

// NewCimdDocumentRequirements instantiates a new CimdDocumentRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCimdDocumentRequirements(name string, requirements []CimdRequirementEntry) *CimdDocumentRequirements {
	this := CimdDocumentRequirements{}
	this.Name = name
	this.Requirements = requirements
	return &this
}

// NewCimdDocumentRequirementsWithDefaults instantiates a new CimdDocumentRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCimdDocumentRequirementsWithDefaults() *CimdDocumentRequirements {
	this := CimdDocumentRequirements{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CimdDocumentRequirements) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdDocumentRequirements) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CimdDocumentRequirements) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CimdDocumentRequirements) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CimdDocumentRequirements) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdDocumentRequirements) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CimdDocumentRequirements) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CimdDocumentRequirements) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CimdDocumentRequirements) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdDocumentRequirements) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CimdDocumentRequirements) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CimdDocumentRequirements) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CimdDocumentRequirements) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdDocumentRequirements) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *CimdDocumentRequirements) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *CimdDocumentRequirements) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *CimdDocumentRequirements) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CimdDocumentRequirements) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CimdDocumentRequirements) SetName(v string) {
	o.Name = v
}

// GetRequirements returns the Requirements field value
func (o *CimdDocumentRequirements) GetRequirements() []CimdRequirementEntry {
	if o == nil {
		var ret []CimdRequirementEntry
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *CimdDocumentRequirements) GetRequirementsOk() ([]CimdRequirementEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requirements, true
}

// SetRequirements sets field value
func (o *CimdDocumentRequirements) SetRequirements(v []CimdRequirementEntry) {
	o.Requirements = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CimdDocumentRequirements) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CimdDocumentRequirements) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CimdDocumentRequirements) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *CimdDocumentRequirements) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o CimdDocumentRequirements) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CimdDocumentRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	toSerialize["name"] = o.Name
	toSerialize["requirements"] = o.Requirements
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CimdDocumentRequirements) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"requirements",
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

	varCimdDocumentRequirements := _CimdDocumentRequirements{}

	err = json.Unmarshal(data, &varCimdDocumentRequirements)

	if err != nil {
		return err
	}

	*o = CimdDocumentRequirements(varCimdDocumentRequirements)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "requirements")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCimdDocumentRequirements struct {
	value *CimdDocumentRequirements
	isSet bool
}

func (v NullableCimdDocumentRequirements) Get() *CimdDocumentRequirements {
	return v.value
}

func (v *NullableCimdDocumentRequirements) Set(val *CimdDocumentRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableCimdDocumentRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableCimdDocumentRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCimdDocumentRequirements(val *CimdDocumentRequirements) *NullableCimdDocumentRequirements {
	return &NullableCimdDocumentRequirements{value: val, isSet: true}
}

func (v NullableCimdDocumentRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCimdDocumentRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
