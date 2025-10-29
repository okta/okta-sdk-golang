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

// checks if the PolicyMappingLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyMappingLinks{}

// PolicyMappingLinks struct for PolicyMappingLinks
type PolicyMappingLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Link to the mapped application
	Application *HrefObject `json:"application,omitempty"`
	// Link to the mapped policy
	Policy               *HrefObject `json:"policy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyMappingLinks PolicyMappingLinks

// NewPolicyMappingLinks instantiates a new PolicyMappingLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyMappingLinks() *PolicyMappingLinks {
	this := PolicyMappingLinks{}
	return &this
}

// NewPolicyMappingLinksWithDefaults instantiates a new PolicyMappingLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyMappingLinksWithDefaults() *PolicyMappingLinks {
	this := PolicyMappingLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PolicyMappingLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PolicyMappingLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *PolicyMappingLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *PolicyMappingLinks) GetApplication() HrefObject {
	if o == nil || IsNil(o.Application) {
		var ret HrefObject
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingLinks) GetApplicationOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *PolicyMappingLinks) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given HrefObject and assigns it to the Application field.
func (o *PolicyMappingLinks) SetApplication(v HrefObject) {
	o.Application = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *PolicyMappingLinks) GetPolicy() HrefObject {
	if o == nil || IsNil(o.Policy) {
		var ret HrefObject
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingLinks) GetPolicyOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *PolicyMappingLinks) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given HrefObject and assigns it to the Policy field.
func (o *PolicyMappingLinks) SetPolicy(v HrefObject) {
	o.Policy = &v
}

func (o PolicyMappingLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyMappingLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyMappingLinks) UnmarshalJSON(data []byte) (err error) {
	varPolicyMappingLinks := _PolicyMappingLinks{}

	err = json.Unmarshal(data, &varPolicyMappingLinks)

	if err != nil {
		return err
	}

	*o = PolicyMappingLinks(varPolicyMappingLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "application")
		delete(additionalProperties, "policy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyMappingLinks struct {
	value *PolicyMappingLinks
	isSet bool
}

func (v NullablePolicyMappingLinks) Get() *PolicyMappingLinks {
	return v.value
}

func (v *NullablePolicyMappingLinks) Set(val *PolicyMappingLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyMappingLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyMappingLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyMappingLinks(val *PolicyMappingLinks) *NullablePolicyMappingLinks {
	return &NullablePolicyMappingLinks{value: val, isSet: true}
}

func (v NullablePolicyMappingLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyMappingLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
