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
)

// checks if the ProtectedResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtectedResource{}

// ProtectedResource <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>A resource protected by a custom authorization server. This shape is common to all resource types (see `resourceType`). Type-specific properties, such as `resourceUrl`, are only present for the resource types that have them.
type ProtectedResource struct {
	// Unique identifier of the protected resource
	Id string `json:"id"`
	// <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>The type of a resource protected by an authorization server
	ResourceType string `json:"resourceType"`
	// The URL of the protected resource. Present only for resource types that expose a resource URL.
	ResourceUrl *string `json:"resourceUrl,omitempty"`
	// Current status of the resource in its lifecycle
	Status               string    `json:"status"`
	Links                LinksSelf `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ProtectedResource ProtectedResource

// NewProtectedResource instantiates a new ProtectedResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectedResource(id string, resourceType string, status string, links LinksSelf) *ProtectedResource {
	this := ProtectedResource{}
	this.Id = id
	this.ResourceType = resourceType
	this.Status = status
	this.Links = links
	return &this
}

// NewProtectedResourceWithDefaults instantiates a new ProtectedResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectedResourceWithDefaults() *ProtectedResource {
	this := ProtectedResource{}
	return &this
}

// GetId returns the Id field value
func (o *ProtectedResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProtectedResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProtectedResource) SetId(v string) {
	o.Id = v
}

// GetResourceType returns the ResourceType field value
func (o *ProtectedResource) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ProtectedResource) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *ProtectedResource) SetResourceType(v string) {
	o.ResourceType = v
}

// GetResourceUrl returns the ResourceUrl field value if set, zero value otherwise.
func (o *ProtectedResource) GetResourceUrl() string {
	if o == nil || IsNil(o.ResourceUrl) {
		var ret string
		return ret
	}
	return *o.ResourceUrl
}

// GetResourceUrlOk returns a tuple with the ResourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedResource) GetResourceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceUrl) {
		return nil, false
	}
	return o.ResourceUrl, true
}

// HasResourceUrl returns a boolean if a field has been set.
func (o *ProtectedResource) HasResourceUrl() bool {
	if o != nil && !IsNil(o.ResourceUrl) {
		return true
	}

	return false
}

// SetResourceUrl gets a reference to the given string and assigns it to the ResourceUrl field.
func (o *ProtectedResource) SetResourceUrl(v string) {
	o.ResourceUrl = &v
}

// GetStatus returns the Status field value
func (o *ProtectedResource) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProtectedResource) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ProtectedResource) SetStatus(v string) {
	o.Status = v
}

// GetLinks returns the Links field value
func (o *ProtectedResource) GetLinks() LinksSelf {
	if o == nil {
		var ret LinksSelf
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ProtectedResource) GetLinksOk() (*LinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ProtectedResource) SetLinks(v LinksSelf) {
	o.Links = v
}

func (o ProtectedResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtectedResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["resourceType"] = o.ResourceType
	if !IsNil(o.ResourceUrl) {
		toSerialize["resourceUrl"] = o.ResourceUrl
	}
	toSerialize["status"] = o.Status
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProtectedResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"resourceType",
		"status",
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

	varProtectedResource := _ProtectedResource{}

	err = json.Unmarshal(data, &varProtectedResource)

	if err != nil {
		return err
	}

	*o = ProtectedResource(varProtectedResource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "resourceUrl")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProtectedResource struct {
	value *ProtectedResource
	isSet bool
}

func (v NullableProtectedResource) Get() *ProtectedResource {
	return v.value
}

func (v *NullableProtectedResource) Set(val *ProtectedResource) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectedResource) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectedResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectedResource(val *ProtectedResource) *NullableProtectedResource {
	return &NullableProtectedResource{value: val, isSet: true}
}

func (v NullableProtectedResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectedResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
