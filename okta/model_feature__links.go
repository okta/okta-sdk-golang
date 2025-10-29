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

// checks if the FeatureLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureLinks{}

// FeatureLinks struct for FeatureLinks
type FeatureLinks struct {
	Self                 *HrefObjectSelfLink            `json:"self,omitempty"`
	Dependents           *FeatureLinksAllOfDependents   `json:"dependents,omitempty"`
	Dependencies         *FeatureLinksAllOfDependencies `json:"dependencies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeatureLinks FeatureLinks

// NewFeatureLinks instantiates a new FeatureLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureLinks() *FeatureLinks {
	this := FeatureLinks{}
	return &this
}

// NewFeatureLinksWithDefaults instantiates a new FeatureLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureLinksWithDefaults() *FeatureLinks {
	this := FeatureLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *FeatureLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *FeatureLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *FeatureLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetDependents returns the Dependents field value if set, zero value otherwise.
func (o *FeatureLinks) GetDependents() FeatureLinksAllOfDependents {
	if o == nil || IsNil(o.Dependents) {
		var ret FeatureLinksAllOfDependents
		return ret
	}
	return *o.Dependents
}

// GetDependentsOk returns a tuple with the Dependents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureLinks) GetDependentsOk() (*FeatureLinksAllOfDependents, bool) {
	if o == nil || IsNil(o.Dependents) {
		return nil, false
	}
	return o.Dependents, true
}

// HasDependents returns a boolean if a field has been set.
func (o *FeatureLinks) HasDependents() bool {
	if o != nil && !IsNil(o.Dependents) {
		return true
	}

	return false
}

// SetDependents gets a reference to the given FeatureLinksAllOfDependents and assigns it to the Dependents field.
func (o *FeatureLinks) SetDependents(v FeatureLinksAllOfDependents) {
	o.Dependents = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *FeatureLinks) GetDependencies() FeatureLinksAllOfDependencies {
	if o == nil || IsNil(o.Dependencies) {
		var ret FeatureLinksAllOfDependencies
		return ret
	}
	return *o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureLinks) GetDependenciesOk() (*FeatureLinksAllOfDependencies, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *FeatureLinks) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given FeatureLinksAllOfDependencies and assigns it to the Dependencies field.
func (o *FeatureLinks) SetDependencies(v FeatureLinksAllOfDependencies) {
	o.Dependencies = &v
}

func (o FeatureLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Dependents) {
		toSerialize["dependents"] = o.Dependents
	}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FeatureLinks) UnmarshalJSON(data []byte) (err error) {
	varFeatureLinks := _FeatureLinks{}

	err = json.Unmarshal(data, &varFeatureLinks)

	if err != nil {
		return err
	}

	*o = FeatureLinks(varFeatureLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "dependents")
		delete(additionalProperties, "dependencies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeatureLinks struct {
	value *FeatureLinks
	isSet bool
}

func (v NullableFeatureLinks) Get() *FeatureLinks {
	return v.value
}

func (v *NullableFeatureLinks) Set(val *FeatureLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureLinks(val *FeatureLinks) *NullableFeatureLinks {
	return &NullableFeatureLinks{value: val, isSet: true}
}

func (v NullableFeatureLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
