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

// checks if the Provisioning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Provisioning{}

// Provisioning Specifies the behavior for just-in-time (JIT) provisioning of an IdP user as a new Okta user and their group memberships
type Provisioning struct {
	// Specifies the user provisioning action during authentication when an IdP user isn't linked to an existing Okta user. * To successfully provision a new Okta user, you must enable just-in-time (JIT) provisioning in your org security settings. * If the target username isn't unique or the resulting Okta user profile is missing a required profile attribute, JIT provisioning may fail. * New Okta users are provisioned with either a `FEDERATION` or `SOCIAL` authentication provider depending on the IdP type.
	Action     *string                 `json:"action,omitempty"`
	Conditions *ProvisioningConditions `json:"conditions,omitempty"`
	Groups     *ProvisioningGroups     `json:"groups,omitempty"`
	// Determines if the IdP should act as a source of truth for user profile attributes
	ProfileMaster        *bool `json:"profileMaster,omitempty"`
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
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provisioning) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *Provisioning) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
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
	if o == nil || IsNil(o.Conditions) {
		var ret ProvisioningConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provisioning) GetConditionsOk() (*ProvisioningConditions, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *Provisioning) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
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
	if o == nil || IsNil(o.Groups) {
		var ret ProvisioningGroups
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provisioning) GetGroupsOk() (*ProvisioningGroups, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Provisioning) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
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
	if o == nil || IsNil(o.ProfileMaster) {
		var ret bool
		return ret
	}
	return *o.ProfileMaster
}

// GetProfileMasterOk returns a tuple with the ProfileMaster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Provisioning) GetProfileMasterOk() (*bool, bool) {
	if o == nil || IsNil(o.ProfileMaster) {
		return nil, false
	}
	return o.ProfileMaster, true
}

// HasProfileMaster returns a boolean if a field has been set.
func (o *Provisioning) HasProfileMaster() bool {
	if o != nil && !IsNil(o.ProfileMaster) {
		return true
	}

	return false
}

// SetProfileMaster gets a reference to the given bool and assigns it to the ProfileMaster field.
func (o *Provisioning) SetProfileMaster(v bool) {
	o.ProfileMaster = &v
}

func (o Provisioning) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Provisioning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.ProfileMaster) {
		toSerialize["profileMaster"] = o.ProfileMaster
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Provisioning) UnmarshalJSON(data []byte) (err error) {
	varProvisioning := _Provisioning{}

	err = json.Unmarshal(data, &varProvisioning)

	if err != nil {
		return err
	}

	*o = Provisioning(varProvisioning)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "profileMaster")
		o.AdditionalProperties = additionalProperties
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
