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

// PolicyLinks struct for PolicyLinks
type PolicyLinks struct {
	Activate *HrefObjectActivateLink `json:"activate,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	Mappings *HrefObjectMappingsLink `json:"mappings,omitempty"`
	Rules *HrefObjectRulesLink `json:"rules,omitempty"`
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyLinks PolicyLinks

// NewPolicyLinks instantiates a new PolicyLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyLinks() *PolicyLinks {
	this := PolicyLinks{}
	return &this
}

// NewPolicyLinksWithDefaults instantiates a new PolicyLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyLinksWithDefaults() *PolicyLinks {
	this := PolicyLinks{}
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *PolicyLinks) GetActivate() HrefObjectActivateLink {
	if o == nil || o.Activate == nil {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyLinks) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *PolicyLinks) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *PolicyLinks) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *PolicyLinks) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || o.Deactivate == nil {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *PolicyLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *PolicyLinks) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *PolicyLinks) GetMappings() HrefObjectMappingsLink {
	if o == nil || o.Mappings == nil {
		var ret HrefObjectMappingsLink
		return ret
	}
	return *o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyLinks) GetMappingsOk() (*HrefObjectMappingsLink, bool) {
	if o == nil || o.Mappings == nil {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *PolicyLinks) HasMappings() bool {
	if o != nil && o.Mappings != nil {
		return true
	}

	return false
}

// SetMappings gets a reference to the given HrefObjectMappingsLink and assigns it to the Mappings field.
func (o *PolicyLinks) SetMappings(v HrefObjectMappingsLink) {
	o.Mappings = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *PolicyLinks) GetRules() HrefObjectRulesLink {
	if o == nil || o.Rules == nil {
		var ret HrefObjectRulesLink
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyLinks) GetRulesOk() (*HrefObjectRulesLink, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *PolicyLinks) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given HrefObjectRulesLink and assigns it to the Rules field.
func (o *PolicyLinks) SetRules(v HrefObjectRulesLink) {
	o.Rules = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PolicyLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PolicyLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *PolicyLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

func (o PolicyLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyLinks) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyLinks := _PolicyLinks{}

	err = json.Unmarshal(bytes, &varPolicyLinks)
	if err == nil {
		*o = PolicyLinks(varPolicyLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "mappings")
		delete(additionalProperties, "rules")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyLinks struct {
	value *PolicyLinks
	isSet bool
}

func (v NullablePolicyLinks) Get() *PolicyLinks {
	return v.value
}

func (v *NullablePolicyLinks) Set(val *PolicyLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyLinks(val *PolicyLinks) *NullablePolicyLinks {
	return &NullablePolicyLinks{value: val, isSet: true}
}

func (v NullablePolicyLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

