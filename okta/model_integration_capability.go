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
	"time"
)

// checks if the IntegrationCapability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationCapability{}

// IntegrationCapability struct for IntegrationCapability
type IntegrationCapability struct {
	// Description of the integration capability
	Description *string `json:"description,omitempty"`
	// URL to the help documentation for the integration capability
	HelpUrl *string `json:"helpUrl,omitempty"`
	// Unique identifier for the integration capability
	Id *string `json:"id,omitempty"`
	// Name of the integration capability
	Name *string `json:"name,omitempty"`
	// The date when the integration capability was released
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
	// Status of the integration capability
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IntegrationCapability IntegrationCapability

// NewIntegrationCapability instantiates a new IntegrationCapability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationCapability() *IntegrationCapability {
	this := IntegrationCapability{}
	return &this
}

// NewIntegrationCapabilityWithDefaults instantiates a new IntegrationCapability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationCapabilityWithDefaults() *IntegrationCapability {
	this := IntegrationCapability{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationCapability) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCapability) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IntegrationCapability) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IntegrationCapability) SetDescription(v string) {
	o.Description = &v
}

// GetHelpUrl returns the HelpUrl field value if set, zero value otherwise.
func (o *IntegrationCapability) GetHelpUrl() string {
	if o == nil || IsNil(o.HelpUrl) {
		var ret string
		return ret
	}
	return *o.HelpUrl
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCapability) GetHelpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HelpUrl) {
		return nil, false
	}
	return o.HelpUrl, true
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *IntegrationCapability) HasHelpUrl() bool {
	if o != nil && !IsNil(o.HelpUrl) {
		return true
	}

	return false
}

// SetHelpUrl gets a reference to the given string and assigns it to the HelpUrl field.
func (o *IntegrationCapability) SetHelpUrl(v string) {
	o.HelpUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationCapability) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCapability) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationCapability) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationCapability) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IntegrationCapability) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCapability) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IntegrationCapability) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IntegrationCapability) SetName(v string) {
	o.Name = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *IntegrationCapability) GetReleaseDate() time.Time {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCapability) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *IntegrationCapability) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *IntegrationCapability) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IntegrationCapability) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationCapability) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IntegrationCapability) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IntegrationCapability) SetStatus(v string) {
	o.Status = &v
}

func (o IntegrationCapability) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationCapability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.HelpUrl) {
		toSerialize["helpUrl"] = o.HelpUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IntegrationCapability) UnmarshalJSON(data []byte) (err error) {
	varIntegrationCapability := _IntegrationCapability{}

	err = json.Unmarshal(data, &varIntegrationCapability)

	if err != nil {
		return err
	}

	*o = IntegrationCapability(varIntegrationCapability)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "helpUrl")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "releaseDate")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntegrationCapability struct {
	value *IntegrationCapability
	isSet bool
}

func (v NullableIntegrationCapability) Get() *IntegrationCapability {
	return v.value
}

func (v *NullableIntegrationCapability) Set(val *IntegrationCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationCapability(val *IntegrationCapability) *NullableIntegrationCapability {
	return &NullableIntegrationCapability{value: val, isSet: true}
}

func (v NullableIntegrationCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
