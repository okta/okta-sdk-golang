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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// UserImportRequestDataContext struct for UserImportRequestDataContext
type UserImportRequestDataContext struct {
	// An array of user profile attributes that are in conflict
	Conflicts []map[string]interface{} `json:"conflicts,omitempty"`
	Application *UserImportRequestDataContextApplication `json:"application,omitempty"`
	Job *UserImportRequestDataContextJob `json:"job,omitempty"`
	// The list of Okta users currently matched to the app user based on import matching. There can be more than one match.
	Matches []map[string]interface{} `json:"matches,omitempty"`
	// The list of any policies that apply to the import matching
	Policy []map[string]interface{} `json:"policy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserImportRequestDataContext UserImportRequestDataContext

// NewUserImportRequestDataContext instantiates a new UserImportRequestDataContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportRequestDataContext() *UserImportRequestDataContext {
	this := UserImportRequestDataContext{}
	return &this
}

// NewUserImportRequestDataContextWithDefaults instantiates a new UserImportRequestDataContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportRequestDataContextWithDefaults() *UserImportRequestDataContext {
	this := UserImportRequestDataContext{}
	return &this
}

// GetConflicts returns the Conflicts field value if set, zero value otherwise.
func (o *UserImportRequestDataContext) GetConflicts() []map[string]interface{} {
	if o == nil || o.Conflicts == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Conflicts
}

// GetConflictsOk returns a tuple with the Conflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataContext) GetConflictsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Conflicts == nil {
		return nil, false
	}
	return o.Conflicts, true
}

// HasConflicts returns a boolean if a field has been set.
func (o *UserImportRequestDataContext) HasConflicts() bool {
	if o != nil && o.Conflicts != nil {
		return true
	}

	return false
}

// SetConflicts gets a reference to the given []map[string]interface{} and assigns it to the Conflicts field.
func (o *UserImportRequestDataContext) SetConflicts(v []map[string]interface{}) {
	o.Conflicts = v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *UserImportRequestDataContext) GetApplication() UserImportRequestDataContextApplication {
	if o == nil || o.Application == nil {
		var ret UserImportRequestDataContextApplication
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataContext) GetApplicationOk() (*UserImportRequestDataContextApplication, bool) {
	if o == nil || o.Application == nil {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *UserImportRequestDataContext) HasApplication() bool {
	if o != nil && o.Application != nil {
		return true
	}

	return false
}

// SetApplication gets a reference to the given UserImportRequestDataContextApplication and assigns it to the Application field.
func (o *UserImportRequestDataContext) SetApplication(v UserImportRequestDataContextApplication) {
	o.Application = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *UserImportRequestDataContext) GetJob() UserImportRequestDataContextJob {
	if o == nil || o.Job == nil {
		var ret UserImportRequestDataContextJob
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataContext) GetJobOk() (*UserImportRequestDataContextJob, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *UserImportRequestDataContext) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given UserImportRequestDataContextJob and assigns it to the Job field.
func (o *UserImportRequestDataContext) SetJob(v UserImportRequestDataContextJob) {
	o.Job = &v
}

// GetMatches returns the Matches field value if set, zero value otherwise.
func (o *UserImportRequestDataContext) GetMatches() []map[string]interface{} {
	if o == nil || o.Matches == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataContext) GetMatchesOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Matches == nil {
		return nil, false
	}
	return o.Matches, true
}

// HasMatches returns a boolean if a field has been set.
func (o *UserImportRequestDataContext) HasMatches() bool {
	if o != nil && o.Matches != nil {
		return true
	}

	return false
}

// SetMatches gets a reference to the given []map[string]interface{} and assigns it to the Matches field.
func (o *UserImportRequestDataContext) SetMatches(v []map[string]interface{}) {
	o.Matches = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *UserImportRequestDataContext) GetPolicy() []map[string]interface{} {
	if o == nil || o.Policy == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataContext) GetPolicyOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *UserImportRequestDataContext) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given []map[string]interface{} and assigns it to the Policy field.
func (o *UserImportRequestDataContext) SetPolicy(v []map[string]interface{}) {
	o.Policy = v
}

func (o UserImportRequestDataContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conflicts != nil {
		toSerialize["conflicts"] = o.Conflicts
	}
	if o.Application != nil {
		toSerialize["application"] = o.Application
	}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	if o.Matches != nil {
		toSerialize["matches"] = o.Matches
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserImportRequestDataContext) UnmarshalJSON(bytes []byte) (err error) {
	varUserImportRequestDataContext := _UserImportRequestDataContext{}

	err = json.Unmarshal(bytes, &varUserImportRequestDataContext)
	if err == nil {
		*o = UserImportRequestDataContext(varUserImportRequestDataContext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conflicts")
		delete(additionalProperties, "application")
		delete(additionalProperties, "job")
		delete(additionalProperties, "matches")
		delete(additionalProperties, "policy")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserImportRequestDataContext struct {
	value *UserImportRequestDataContext
	isSet bool
}

func (v NullableUserImportRequestDataContext) Get() *UserImportRequestDataContext {
	return v.value
}

func (v *NullableUserImportRequestDataContext) Set(val *UserImportRequestDataContext) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportRequestDataContext) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportRequestDataContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportRequestDataContext(val *UserImportRequestDataContext) *NullableUserImportRequestDataContext {
	return &NullableUserImportRequestDataContext{value: val, isSet: true}
}

func (v NullableUserImportRequestDataContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportRequestDataContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

