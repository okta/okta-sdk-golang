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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// ProfileMappingTarget The parameter is the target of a profile mapping and is a valid [JSON Schema Draft 4](https://datatracker.ietf.org/doc/html/draft-zyp-json-schema-04) document with the following properties. The data type can be an app instance or an Okta object.   > **Note:** If the target is Okta and the UserTypes feature isn't enabled, then the target `_links` only has a link to the schema.
type ProfileMappingTarget struct {
	// Unique identifier for the application instance or UserType
	Id *string `json:"id,omitempty"`
	// Variable name of the application instance or name of the referenced userType
	Name *string `json:"name,omitempty"`
	// Type of user referenced in the mapping
	Type *string `json:"type,omitempty"`
	Links *SourceLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProfileMappingTarget ProfileMappingTarget

// NewProfileMappingTarget instantiates a new ProfileMappingTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileMappingTarget() *ProfileMappingTarget {
	this := ProfileMappingTarget{}
	return &this
}

// NewProfileMappingTargetWithDefaults instantiates a new ProfileMappingTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileMappingTargetWithDefaults() *ProfileMappingTarget {
	this := ProfileMappingTarget{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProfileMappingTarget) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingTarget) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProfileMappingTarget) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProfileMappingTarget) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProfileMappingTarget) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingTarget) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProfileMappingTarget) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProfileMappingTarget) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProfileMappingTarget) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingTarget) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProfileMappingTarget) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProfileMappingTarget) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProfileMappingTarget) GetLinks() SourceLinks {
	if o == nil || o.Links == nil {
		var ret SourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileMappingTarget) GetLinksOk() (*SourceLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProfileMappingTarget) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SourceLinks and assigns it to the Links field.
func (o *ProfileMappingTarget) SetLinks(v SourceLinks) {
	o.Links = &v
}

func (o ProfileMappingTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
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

func (o *ProfileMappingTarget) UnmarshalJSON(bytes []byte) (err error) {
	varProfileMappingTarget := _ProfileMappingTarget{}

	err = json.Unmarshal(bytes, &varProfileMappingTarget)
	if err == nil {
		*o = ProfileMappingTarget(varProfileMappingTarget)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProfileMappingTarget struct {
	value *ProfileMappingTarget
	isSet bool
}

func (v NullableProfileMappingTarget) Get() *ProfileMappingTarget {
	return v.value
}

func (v *NullableProfileMappingTarget) Set(val *ProfileMappingTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileMappingTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileMappingTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileMappingTarget(val *ProfileMappingTarget) *NullableProfileMappingTarget {
	return &NullableProfileMappingTarget{value: val, isSet: true}
}

func (v NullableProfileMappingTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileMappingTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

