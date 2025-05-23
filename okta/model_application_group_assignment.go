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
	"time"
)

// ApplicationGroupAssignment The Application Group object that defines a group of users' app-specific profile and credentials for an app
type ApplicationGroupAssignment struct {
	// ID of the [Group](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/)
	Id *string `json:"id,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Priority assigned to the group. If an app has more than one group assigned to the same user, then the group with the higher priority has its profile applied to the [Application User](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/ApplicationUsers/). If a priority value isn't specified, then the next highest priority is assigned by default. See [Assign attribute group priority](https://help.okta.com/okta_help.htm?type=oie&id=ext-usgp-app-group-priority) and the [sample priority use case](https://help.okta.com/okta_help.htm?type=oie&id=ext-usgp-combine-values-use).
	Priority *int32 `json:"priority,omitempty"`
	// Specifies the profile properties applied to [Application Users](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/ApplicationUsers/) that are assigned to the app through group membership. Some reference properties are imported from the target app and can't be configured. See [profile](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/User/#tag/User/operation/getUser!c=200&path=profile&t=response).
	Profile map[string]interface{} `json:"profile,omitempty"`
	// Embedded resource related to the Application Group using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. If the `expand=group` query parameter is specified, then the [Group](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/) object is embedded. If the `expand=metadata` query parameter is specified, then the group assignment metadata is embedded.
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links *ApplicationGroupAssignmentLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationGroupAssignment ApplicationGroupAssignment

// NewApplicationGroupAssignment instantiates a new ApplicationGroupAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationGroupAssignment() *ApplicationGroupAssignment {
	this := ApplicationGroupAssignment{}
	return &this
}

// NewApplicationGroupAssignmentWithDefaults instantiates a new ApplicationGroupAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationGroupAssignmentWithDefaults() *ApplicationGroupAssignment {
	this := ApplicationGroupAssignment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationGroupAssignment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGroupAssignment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationGroupAssignment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationGroupAssignment) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ApplicationGroupAssignment) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGroupAssignment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ApplicationGroupAssignment) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ApplicationGroupAssignment) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ApplicationGroupAssignment) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGroupAssignment) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ApplicationGroupAssignment) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ApplicationGroupAssignment) SetPriority(v int32) {
	o.Priority = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ApplicationGroupAssignment) GetProfile() map[string]interface{} {
	if o == nil || o.Profile == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGroupAssignment) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ApplicationGroupAssignment) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *ApplicationGroupAssignment) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *ApplicationGroupAssignment) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || o.Embedded == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGroupAssignment) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *ApplicationGroupAssignment) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *ApplicationGroupAssignment) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApplicationGroupAssignment) GetLinks() ApplicationGroupAssignmentLinks {
	if o == nil || o.Links == nil {
		var ret ApplicationGroupAssignmentLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGroupAssignment) GetLinksOk() (*ApplicationGroupAssignmentLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApplicationGroupAssignment) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ApplicationGroupAssignmentLinks and assigns it to the Links field.
func (o *ApplicationGroupAssignment) SetLinks(v ApplicationGroupAssignmentLinks) {
	o.Links = &v
}

func (o ApplicationGroupAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationGroupAssignment) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationGroupAssignment := _ApplicationGroupAssignment{}

	err = json.Unmarshal(bytes, &varApplicationGroupAssignment)
	if err == nil {
		*o = ApplicationGroupAssignment(varApplicationGroupAssignment)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationGroupAssignment struct {
	value *ApplicationGroupAssignment
	isSet bool
}

func (v NullableApplicationGroupAssignment) Get() *ApplicationGroupAssignment {
	return v.value
}

func (v *NullableApplicationGroupAssignment) Set(val *ApplicationGroupAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationGroupAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationGroupAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationGroupAssignment(val *ApplicationGroupAssignment) *NullableApplicationGroupAssignment {
	return &NullableApplicationGroupAssignment{value: val, isSet: true}
}

func (v NullableApplicationGroupAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationGroupAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

