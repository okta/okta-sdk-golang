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
)

// Provisioning struct for Provisioning
type Provisioning struct {
	Action *string `json:"action,omitempty"`
	Conditions *ProvisioningConditions `json:"conditions,omitempty"`
	Groups *ProvisioningGroups `json:"groups,omitempty"`
	ProfileMaster *bool `json:"profileMaster,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Provisioning Provisioning

// NewProvisioning instantiates a new Provisioning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioning() *Provisioning {
	this := Provisioning{}
	return &this
}

// NewProvisioningWithDefaults instantiates a new Provisioning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningWithDefaults() *Provisioning {
	this := Provisioning{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *Provisioning) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provisioning) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Provisioning) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *Provisioning) SetAction(v string) {
	o.Action = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *Provisioning) GetConditions() ProvisioningConditions {
	if o == nil || o.Conditions == nil {
		var ret ProvisioningConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provisioning) GetConditionsOk() (*ProvisioningConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *Provisioning) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given ProvisioningConditions and assigns it to the Conditions field.
func (o *Provisioning) SetConditions(v ProvisioningConditions) {
	o.Conditions = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *Provisioning) GetGroups() ProvisioningGroups {
	if o == nil || o.Groups == nil {
		var ret ProvisioningGroups
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provisioning) GetGroupsOk() (*ProvisioningGroups, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Provisioning) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given ProvisioningGroups and assigns it to the Groups field.
func (o *Provisioning) SetGroups(v ProvisioningGroups) {
	o.Groups = &v
}

// GetProfileMaster returns the ProfileMaster field value if set, zero value otherwise.
func (o *Provisioning) GetProfileMaster() bool {
	if o == nil || o.ProfileMaster == nil {
		var ret bool
		return ret
	}
	return *o.ProfileMaster
}

// GetProfileMasterOk returns a tuple with the ProfileMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provisioning) GetProfileMasterOk() (*bool, bool) {
	if o == nil || o.ProfileMaster == nil {
		return nil, false
	}
	return o.ProfileMaster, true
}

// HasProfileMaster returns a boolean if a field has been set.
func (o *Provisioning) HasProfileMaster() bool {
	if o != nil && o.ProfileMaster != nil {
		return true
	}

	return false
}

// SetProfileMaster gets a reference to the given bool and assigns it to the ProfileMaster field.
func (o *Provisioning) SetProfileMaster(v bool) {
	o.ProfileMaster = &v
}

func (o Provisioning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.ProfileMaster != nil {
		toSerialize["profileMaster"] = o.ProfileMaster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Provisioning) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioning := _Provisioning{}

	err = json.Unmarshal(bytes, &varProvisioning)
	if err == nil {
		*o = Provisioning(varProvisioning)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "profileMaster")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProvisioning struct {
	value *Provisioning
	isSet bool
}

func (v NullableProvisioning) Get() *Provisioning {
	return v.value
}

func (v *NullableProvisioning) Set(val *Provisioning) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioning) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioning(val *Provisioning) *NullableProvisioning {
	return &NullableProvisioning{value: val, isSet: true}
}

func (v NullableProvisioning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

