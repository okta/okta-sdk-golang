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

// checks if the ManagedConnectionServiceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedConnectionServiceAccount{}

// ManagedConnectionServiceAccount Service account for the managed connection
type ManagedConnectionServiceAccount struct {
	// Display name of the service account
	Name string `json:"name"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the service account
	Orn                  string                         `json:"orn"`
	Links                CustomAuthorizationServerLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ManagedConnectionServiceAccount ManagedConnectionServiceAccount

// NewManagedConnectionServiceAccount instantiates a new ManagedConnectionServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedConnectionServiceAccount(name string, orn string, links CustomAuthorizationServerLinks) *ManagedConnectionServiceAccount {
	this := ManagedConnectionServiceAccount{}
	this.Name = name
	this.Orn = orn
	this.Links = links
	return &this
}

// NewManagedConnectionServiceAccountWithDefaults instantiates a new ManagedConnectionServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedConnectionServiceAccountWithDefaults() *ManagedConnectionServiceAccount {
	this := ManagedConnectionServiceAccount{}
	return &this
}

// GetName returns the Name field value
func (o *ManagedConnectionServiceAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ManagedConnectionServiceAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ManagedConnectionServiceAccount) SetName(v string) {
	o.Name = v
}

// GetOrn returns the Orn field value
func (o *ManagedConnectionServiceAccount) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *ManagedConnectionServiceAccount) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *ManagedConnectionServiceAccount) SetOrn(v string) {
	o.Orn = v
}

// GetLinks returns the Links field value
func (o *ManagedConnectionServiceAccount) GetLinks() CustomAuthorizationServerLinks {
	if o == nil {
		var ret CustomAuthorizationServerLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ManagedConnectionServiceAccount) GetLinksOk() (*CustomAuthorizationServerLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ManagedConnectionServiceAccount) SetLinks(v CustomAuthorizationServerLinks) {
	o.Links = v
}

func (o ManagedConnectionServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedConnectionServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["orn"] = o.Orn
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ManagedConnectionServiceAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"orn",
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

	varManagedConnectionServiceAccount := _ManagedConnectionServiceAccount{}

	err = json.Unmarshal(data, &varManagedConnectionServiceAccount)

	if err != nil {
		return err
	}

	*o = ManagedConnectionServiceAccount(varManagedConnectionServiceAccount)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableManagedConnectionServiceAccount struct {
	value *ManagedConnectionServiceAccount
	isSet bool
}

func (v NullableManagedConnectionServiceAccount) Get() *ManagedConnectionServiceAccount {
	return v.value
}

func (v *NullableManagedConnectionServiceAccount) Set(val *ManagedConnectionServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedConnectionServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedConnectionServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedConnectionServiceAccount(val *ManagedConnectionServiceAccount) *NullableManagedConnectionServiceAccount {
	return &NullableManagedConnectionServiceAccount{value: val, isSet: true}
}

func (v NullableManagedConnectionServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedConnectionServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
