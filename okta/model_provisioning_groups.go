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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the ProvisioningGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningGroups{}

// ProvisioningGroups Provisioning settings for a user's group memberships
type ProvisioningGroups struct {
	// Provisioning action for the IdP user's group memberships  | Enum     | Description                                                                                                                                                       | Existing OKTA_GROUP Memberships                                                                       | Existing APP_GROUP Memberships | Existing BUILT_IN Memberships | | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------------------ | ----------------------------- | | `APPEND` | Adds a user to any group defined by the IdP as a value of the `sourceAttributeName` array that matches the name of the allow listed group defined in the `filter` | Unchanged                                                                                             | Unchanged                      | Unchanged                     | | `ASSIGN` | Assigns a user to groups defined in the `assignments` array                                                                                                       | Unchanged                                                                                             | Unchanged                      | Unchanged                     | | `NONE`   | Skips processing of group memberships                                                                                                                             | Unchanged                                                                                             | Unchanged                      | Unchanged                     | | `SYNC`   | Group memberships are sourced by the IdP as a value of the `sourceAttributeName` array that matches the name of the group defined in the `filter`                 | Removed if not defined by the IdP in `sourceAttributeName` and matching name of the group in `filter` | Unchanged                      | Unchanged                     |  > **Note:** Group provisioning action is processed independently from profile sourcing. You can sync group memberships through SAML with profile sourcing disabled.
	Action *string `json:"action,omitempty"`
	// List of `OKTA_GROUP` group identifiers to add an IdP user as a member with the `ASSIGN` action
	Assignments []string `json:"assignments,omitempty"`
	// Allowlist of `OKTA_GROUP` group identifiers for the `APPEND` or `SYNC` provisioning action
	Filter []string `json:"filter,omitempty"`
	// IdP user profile attribute name (case-insensitive) for an array value that contains group memberships
	SourceAttributeName  *string `json:"sourceAttributeName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningGroups ProvisioningGroups

// NewProvisioningGroups instantiates a new ProvisioningGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningGroups() *ProvisioningGroups {
	this := ProvisioningGroups{}
	return &this
}

// NewProvisioningGroupsWithDefaults instantiates a new ProvisioningGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningGroupsWithDefaults() *ProvisioningGroups {
	this := ProvisioningGroups{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ProvisioningGroups) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningGroups) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ProvisioningGroups) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ProvisioningGroups) SetAction(v string) {
	o.Action = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *ProvisioningGroups) GetAssignments() []string {
	if o == nil || IsNil(o.Assignments) {
		var ret []string
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningGroups) GetAssignmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *ProvisioningGroups) HasAssignments() bool {
	if o != nil && !IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []string and assigns it to the Assignments field.
func (o *ProvisioningGroups) SetAssignments(v []string) {
	o.Assignments = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ProvisioningGroups) GetFilter() []string {
	if o == nil || IsNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningGroups) GetFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ProvisioningGroups) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *ProvisioningGroups) SetFilter(v []string) {
	o.Filter = v
}

// GetSourceAttributeName returns the SourceAttributeName field value if set, zero value otherwise.
func (o *ProvisioningGroups) GetSourceAttributeName() string {
	if o == nil || IsNil(o.SourceAttributeName) {
		var ret string
		return ret
	}
	return *o.SourceAttributeName
}

// GetSourceAttributeNameOk returns a tuple with the SourceAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningGroups) GetSourceAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.SourceAttributeName) {
		return nil, false
	}
	return o.SourceAttributeName, true
}

// HasSourceAttributeName returns a boolean if a field has been set.
func (o *ProvisioningGroups) HasSourceAttributeName() bool {
	if o != nil && !IsNil(o.SourceAttributeName) {
		return true
	}

	return false
}

// SetSourceAttributeName gets a reference to the given string and assigns it to the SourceAttributeName field.
func (o *ProvisioningGroups) SetSourceAttributeName(v string) {
	o.SourceAttributeName = &v
}

func (o ProvisioningGroups) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.SourceAttributeName) {
		toSerialize["sourceAttributeName"] = o.SourceAttributeName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningGroups) UnmarshalJSON(data []byte) (err error) {
	varProvisioningGroups := _ProvisioningGroups{}

	err = json.Unmarshal(data, &varProvisioningGroups)

	if err != nil {
		return err
	}

	*o = ProvisioningGroups(varProvisioningGroups)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "assignments")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "sourceAttributeName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningGroups struct {
	value *ProvisioningGroups
	isSet bool
}

func (v NullableProvisioningGroups) Get() *ProvisioningGroups {
	return v.value
}

func (v *NullableProvisioningGroups) Set(val *ProvisioningGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningGroups(val *ProvisioningGroups) *NullableProvisioningGroups {
	return &NullableProvisioningGroups{value: val, isSet: true}
}

func (v NullableProvisioningGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
