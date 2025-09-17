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
	"fmt"
)

// checks if the RealmProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmProfile{}

// RealmProfile struct for RealmProfile
type RealmProfile struct {
	// Array of allowed domains. No user in this realm can be created or updated unless they have a username and email from one of these domains.  The following characters aren't allowed in the domain name: `!$%^&()=*+,:;<>'[]|/?\\`
	Domains []string `json:"domains,omitempty"`
	// Name of a realm
	Name string `json:"name"`
	// Used to store partner users. This property must be set to `PARTNER` to access Okta's external partner portal.
	RealmType            *string `json:"realmType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmProfile RealmProfile

// NewRealmProfile instantiates a new RealmProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmProfile(name string) *RealmProfile {
	this := RealmProfile{}
	this.Name = name
	return &this
}

// NewRealmProfileWithDefaults instantiates a new RealmProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmProfileWithDefaults() *RealmProfile {
	this := RealmProfile{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *RealmProfile) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmProfile) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *RealmProfile) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *RealmProfile) SetDomains(v []string) {
	o.Domains = v
}

// GetName returns the Name field value
func (o *RealmProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RealmProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RealmProfile) SetName(v string) {
	o.Name = v
}

// GetRealmType returns the RealmType field value if set, zero value otherwise.
func (o *RealmProfile) GetRealmType() string {
	if o == nil || IsNil(o.RealmType) {
		var ret string
		return ret
	}
	return *o.RealmType
}

// GetRealmTypeOk returns a tuple with the RealmType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealmProfile) GetRealmTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RealmType) {
		return nil, false
	}
	return o.RealmType, true
}

// HasRealmType returns a boolean if a field has been set.
func (o *RealmProfile) HasRealmType() bool {
	if o != nil && !IsNil(o.RealmType) {
		return true
	}

	return false
}

// SetRealmType gets a reference to the given string and assigns it to the RealmType field.
func (o *RealmProfile) SetRealmType(v string) {
	o.RealmType = &v
}

func (o RealmProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.RealmType) {
		toSerialize["realmType"] = o.RealmType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RealmProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varRealmProfile := _RealmProfile{}

	err = json.Unmarshal(data, &varRealmProfile)

	if err != nil {
		return err
	}

	*o = RealmProfile(varRealmProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domains")
		delete(additionalProperties, "name")
		delete(additionalProperties, "realmType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmProfile struct {
	value *RealmProfile
	isSet bool
}

func (v NullableRealmProfile) Get() *RealmProfile {
	return v.value
}

func (v *NullableRealmProfile) Set(val *RealmProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmProfile(val *RealmProfile) *NullableRealmProfile {
	return &NullableRealmProfile{value: val, isSet: true}
}

func (v NullableRealmProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
