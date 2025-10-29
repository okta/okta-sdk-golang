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
	"fmt"
)

// checks if the OrgCreationAdmin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCreationAdmin{}

// OrgCreationAdmin Profile and credential information for the first super admin user of the child org. If you plan to configure and manage the org programmatically, create a system user with a dedicated email address and a strong password. > **Note:** If you don't provide `credentials`, the super admin user is prompted to set up their credentials when they sign in to the org for the first time.
type OrgCreationAdmin struct {
	Credentials          *OrgCreationAdminCredentials `json:"credentials,omitempty"`
	Profile              OrgCreationAdminProfile      `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _OrgCreationAdmin OrgCreationAdmin

// NewOrgCreationAdmin instantiates a new OrgCreationAdmin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCreationAdmin(profile OrgCreationAdminProfile) *OrgCreationAdmin {
	this := OrgCreationAdmin{}
	this.Profile = profile
	return &this
}

// NewOrgCreationAdminWithDefaults instantiates a new OrgCreationAdmin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCreationAdminWithDefaults() *OrgCreationAdmin {
	this := OrgCreationAdmin{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *OrgCreationAdmin) GetCredentials() OrgCreationAdminCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret OrgCreationAdminCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCreationAdmin) GetCredentialsOk() (*OrgCreationAdminCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *OrgCreationAdmin) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given OrgCreationAdminCredentials and assigns it to the Credentials field.
func (o *OrgCreationAdmin) SetCredentials(v OrgCreationAdminCredentials) {
	o.Credentials = &v
}

// GetProfile returns the Profile field value
func (o *OrgCreationAdmin) GetProfile() OrgCreationAdminProfile {
	if o == nil {
		var ret OrgCreationAdminProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *OrgCreationAdmin) GetProfileOk() (*OrgCreationAdminProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *OrgCreationAdmin) SetProfile(v OrgCreationAdminProfile) {
	o.Profile = v
}

func (o OrgCreationAdmin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCreationAdmin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	toSerialize["profile"] = o.Profile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCreationAdmin) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile",
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

	varOrgCreationAdmin := _OrgCreationAdmin{}

	err = json.Unmarshal(data, &varOrgCreationAdmin)

	if err != nil {
		return err
	}

	*o = OrgCreationAdmin(varOrgCreationAdmin)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCreationAdmin struct {
	value *OrgCreationAdmin
	isSet bool
}

func (v NullableOrgCreationAdmin) Get() *OrgCreationAdmin {
	return v.value
}

func (v *NullableOrgCreationAdmin) Set(val *OrgCreationAdmin) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCreationAdmin) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCreationAdmin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCreationAdmin(val *OrgCreationAdmin) *NullableOrgCreationAdmin {
	return &NullableOrgCreationAdmin{value: val, isSet: true}
}

func (v NullableOrgCreationAdmin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCreationAdmin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
