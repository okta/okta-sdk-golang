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
	"time"
)

// checks if the Group type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Group{}

// Group struct for Group
type Group struct {
	// Timestamp when the group was created
	Created *time.Time `json:"created,omitempty"`
	// Unique ID for the group
	Id *string `json:"id,omitempty"`
	// Timestamp when the groups memberships were last updated
	LastMembershipUpdated *time.Time `json:"lastMembershipUpdated,omitempty"`
	// Timestamp when the group's profile was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Determines the group's `profile`
	ObjectClass []string      `json:"objectClass,omitempty"`
	Profile     *GroupProfile `json:"profile,omitempty"`
	// Determines how a group's profile and memberships are managed
	Type *string `json:"type,omitempty"`
	// Embedded resources related to the group
	Embedded             map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links                *GroupLinks                       `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Group Group

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup() *Group {
	this := Group{}
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Group) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Group) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Group) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Group) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Group) SetId(v string) {
	o.Id = &v
}

// GetLastMembershipUpdated returns the LastMembershipUpdated field value if set, zero value otherwise.
func (o *Group) GetLastMembershipUpdated() time.Time {
	if o == nil || IsNil(o.LastMembershipUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastMembershipUpdated
}

// GetLastMembershipUpdatedOk returns a tuple with the LastMembershipUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLastMembershipUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastMembershipUpdated) {
		return nil, false
	}
	return o.LastMembershipUpdated, true
}

// HasLastMembershipUpdated returns a boolean if a field has been set.
func (o *Group) HasLastMembershipUpdated() bool {
	if o != nil && !IsNil(o.LastMembershipUpdated) {
		return true
	}

	return false
}

// SetLastMembershipUpdated gets a reference to the given time.Time and assigns it to the LastMembershipUpdated field.
func (o *Group) SetLastMembershipUpdated(v time.Time) {
	o.LastMembershipUpdated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Group) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Group) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Group) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetObjectClass returns the ObjectClass field value if set, zero value otherwise.
func (o *Group) GetObjectClass() []string {
	if o == nil || IsNil(o.ObjectClass) {
		var ret []string
		return ret
	}
	return o.ObjectClass
}

// GetObjectClassOk returns a tuple with the ObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetObjectClassOk() ([]string, bool) {
	if o == nil || IsNil(o.ObjectClass) {
		return nil, false
	}
	return o.ObjectClass, true
}

// HasObjectClass returns a boolean if a field has been set.
func (o *Group) HasObjectClass() bool {
	if o != nil && !IsNil(o.ObjectClass) {
		return true
	}

	return false
}

// SetObjectClass gets a reference to the given []string and assigns it to the ObjectClass field.
func (o *Group) SetObjectClass(v []string) {
	o.ObjectClass = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *Group) GetProfile() GroupProfile {
	if o == nil || IsNil(o.Profile) {
		var ret GroupProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetProfileOk() (*GroupProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *Group) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given GroupProfile and assigns it to the Profile field.
func (o *Group) SetProfile(v GroupProfile) {
	o.Profile = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Group) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Group) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Group) SetType(v string) {
	o.Type = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *Group) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *Group) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *Group) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Group) GetLinks() GroupLinks {
	if o == nil || IsNil(o.Links) {
		var ret GroupLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Group) GetLinksOk() (*GroupLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Group) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GroupLinks and assigns it to the Links field.
func (o *Group) SetLinks(v GroupLinks) {
	o.Links = &v
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Group) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastMembershipUpdated) {
		toSerialize["lastMembershipUpdated"] = o.LastMembershipUpdated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.ObjectClass) {
		toSerialize["objectClass"] = o.ObjectClass
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Group) UnmarshalJSON(data []byte) (err error) {
	varGroup := _Group{}

	err = json.Unmarshal(data, &varGroup)

	if err != nil {
		return err
	}

	*o = Group(varGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastMembershipUpdated")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "objectClass")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
