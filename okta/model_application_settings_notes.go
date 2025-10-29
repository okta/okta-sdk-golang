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
)

// checks if the ApplicationSettingsNotes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationSettingsNotes{}

// ApplicationSettingsNotes App notes visible to either the admin or end user
type ApplicationSettingsNotes struct {
	// An app message that's visible to admins
	Admin *string `json:"admin,omitempty"`
	// A message that's visible in the End-User Dashboard
	Enduser              *string `json:"enduser,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationSettingsNotes ApplicationSettingsNotes

// NewApplicationSettingsNotes instantiates a new ApplicationSettingsNotes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSettingsNotes() *ApplicationSettingsNotes {
	this := ApplicationSettingsNotes{}
	return &this
}

// NewApplicationSettingsNotesWithDefaults instantiates a new ApplicationSettingsNotes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSettingsNotesWithDefaults() *ApplicationSettingsNotes {
	this := ApplicationSettingsNotes{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *ApplicationSettingsNotes) GetAdmin() string {
	if o == nil || IsNil(o.Admin) {
		var ret string
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotes) GetAdminOk() (*string, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *ApplicationSettingsNotes) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given string and assigns it to the Admin field.
func (o *ApplicationSettingsNotes) SetAdmin(v string) {
	o.Admin = &v
}

// GetEnduser returns the Enduser field value if set, zero value otherwise.
func (o *ApplicationSettingsNotes) GetEnduser() string {
	if o == nil || IsNil(o.Enduser) {
		var ret string
		return ret
	}
	return *o.Enduser
}

// GetEnduserOk returns a tuple with the Enduser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSettingsNotes) GetEnduserOk() (*string, bool) {
	if o == nil || IsNil(o.Enduser) {
		return nil, false
	}
	return o.Enduser, true
}

// HasEnduser returns a boolean if a field has been set.
func (o *ApplicationSettingsNotes) HasEnduser() bool {
	if o != nil && !IsNil(o.Enduser) {
		return true
	}

	return false
}

// SetEnduser gets a reference to the given string and assigns it to the Enduser field.
func (o *ApplicationSettingsNotes) SetEnduser(v string) {
	o.Enduser = &v
}

func (o ApplicationSettingsNotes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationSettingsNotes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.Enduser) {
		toSerialize["enduser"] = o.Enduser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationSettingsNotes) UnmarshalJSON(data []byte) (err error) {
	varApplicationSettingsNotes := _ApplicationSettingsNotes{}

	err = json.Unmarshal(data, &varApplicationSettingsNotes)

	if err != nil {
		return err
	}

	*o = ApplicationSettingsNotes(varApplicationSettingsNotes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "admin")
		delete(additionalProperties, "enduser")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationSettingsNotes struct {
	value *ApplicationSettingsNotes
	isSet bool
}

func (v NullableApplicationSettingsNotes) Get() *ApplicationSettingsNotes {
	return v.value
}

func (v *NullableApplicationSettingsNotes) Set(val *ApplicationSettingsNotes) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSettingsNotes) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSettingsNotes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSettingsNotes(val *ApplicationSettingsNotes) *NullableApplicationSettingsNotes {
	return &NullableApplicationSettingsNotes{value: val, isSet: true}
}

func (v NullableApplicationSettingsNotes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSettingsNotes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
