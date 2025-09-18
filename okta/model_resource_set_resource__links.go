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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ResourceSetResourceLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSetResourceLinks{}

// ResourceSetResourceLinks Related discoverable resources
type ResourceSetResourceLinks struct {
	// The REST API URL of the related resource
	Self *HrefObject `json:"self,omitempty"`
	// Link to this resource set resource object (self)
	Resource *HrefObject `json:"resource,omitempty"`
	// If applicable, the REST API URL of the related groups resource
	Groups *HrefObject `json:"groups,omitempty"`
	// If applicable, the REST API URL of the related users resource
	Users                *HrefObject `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetResourceLinks ResourceSetResourceLinks

// NewResourceSetResourceLinks instantiates a new ResourceSetResourceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetResourceLinks() *ResourceSetResourceLinks {
	this := ResourceSetResourceLinks{}
	return &this
}

// NewResourceSetResourceLinksWithDefaults instantiates a new ResourceSetResourceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetResourceLinksWithDefaults() *ResourceSetResourceLinks {
	this := ResourceSetResourceLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceSetResourceLinks) GetSelf() HrefObject {
	if o == nil || IsNil(o.Self) {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResourceLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceSetResourceLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *ResourceSetResourceLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourceSetResourceLinks) GetResource() HrefObject {
	if o == nil || IsNil(o.Resource) {
		var ret HrefObject
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResourceLinks) GetResourceOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourceSetResourceLinks) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given HrefObject and assigns it to the Resource field.
func (o *ResourceSetResourceLinks) SetResource(v HrefObject) {
	o.Resource = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ResourceSetResourceLinks) GetGroups() HrefObject {
	if o == nil || IsNil(o.Groups) {
		var ret HrefObject
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResourceLinks) GetGroupsOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ResourceSetResourceLinks) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given HrefObject and assigns it to the Groups field.
func (o *ResourceSetResourceLinks) SetGroups(v HrefObject) {
	o.Groups = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ResourceSetResourceLinks) GetUsers() HrefObject {
	if o == nil || IsNil(o.Users) {
		var ret HrefObject
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSetResourceLinks) GetUsersOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ResourceSetResourceLinks) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given HrefObject and assigns it to the Users field.
func (o *ResourceSetResourceLinks) SetUsers(v HrefObject) {
	o.Users = &v
}

func (o ResourceSetResourceLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSetResourceLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSetResourceLinks) UnmarshalJSON(data []byte) (err error) {
	varResourceSetResourceLinks := _ResourceSetResourceLinks{}

	err = json.Unmarshal(data, &varResourceSetResourceLinks)

	if err != nil {
		return err
	}

	*o = ResourceSetResourceLinks(varResourceSetResourceLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSetResourceLinks struct {
	value *ResourceSetResourceLinks
	isSet bool
}

func (v NullableResourceSetResourceLinks) Get() *ResourceSetResourceLinks {
	return v.value
}

func (v *NullableResourceSetResourceLinks) Set(val *ResourceSetResourceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetResourceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetResourceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetResourceLinks(val *ResourceSetResourceLinks) *NullableResourceSetResourceLinks {
	return &NullableResourceSetResourceLinks{value: val, isSet: true}
}

func (v NullableResourceSetResourceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetResourceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
