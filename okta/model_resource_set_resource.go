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
	"time"
)

// checks if the ResourceSetResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSetResource{}

// ResourceSetResource struct for ResourceSetResource
type ResourceSetResource struct {
	Conditions *ResourceConditions `json:"conditions,omitempty"`
	// Timestamp when the resource set resource object was created
	Created *time.Time `json:"created,omitempty"`
	// Unique ID of the resource set resource object
	Id *string `json:"id,omitempty"`
	// Timestamp when this object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The Okta Resource Name (ORN) of the resource
	Orn                  *string                   `json:"orn,omitempty"`
	Links                *ResourceSetResourceLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetResource ResourceSetResource

// NewResourceSetResource instantiates a new ResourceSetResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetResource() *ResourceSetResource {
	this := ResourceSetResource{}
	return &this
}

// NewResourceSetResourceWithDefaults instantiates a new ResourceSetResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetResourceWithDefaults() *ResourceSetResource {
	this := ResourceSetResource{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ResourceSetResource) GetConditions() ResourceConditions {
	if o == nil || IsNil(o.Conditions) {
		var ret ResourceConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResource) GetConditionsOk() (*ResourceConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ResourceSetResource) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given ResourceConditions and assigns it to the Conditions field.
func (o *ResourceSetResource) SetConditions(v ResourceConditions) {
	o.Conditions = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ResourceSetResource) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResource) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ResourceSetResource) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ResourceSetResource) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceSetResource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceSetResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceSetResource) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ResourceSetResource) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResource) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ResourceSetResource) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ResourceSetResource) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ResourceSetResource) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResource) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ResourceSetResource) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ResourceSetResource) SetOrn(v string) {
	o.Orn = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ResourceSetResource) GetLinks() ResourceSetResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceSetResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResource) GetLinksOk() (*ResourceSetResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ResourceSetResource) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceSetResourceLinks and assigns it to the Links field.
func (o *ResourceSetResource) SetLinks(v ResourceSetResourceLinks) {
	o.Links = &v
}

func (o ResourceSetResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSetResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSetResource) UnmarshalJSON(data []byte) (err error) {
	varResourceSetResource := _ResourceSetResource{}

	err = json.Unmarshal(data, &varResourceSetResource)

	if err != nil {
		return err
	}

	*o = ResourceSetResource(varResourceSetResource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSetResource struct {
	value *ResourceSetResource
	isSet bool
}

func (v NullableResourceSetResource) Get() *ResourceSetResource {
	return v.value
}

func (v *NullableResourceSetResource) Set(val *ResourceSetResource) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetResource(val *ResourceSetResource) *NullableResourceSetResource {
	return &NullableResourceSetResource{value: val, isSet: true}
}

func (v NullableResourceSetResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
