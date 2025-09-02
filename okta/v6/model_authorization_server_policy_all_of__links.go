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

// AuthorizationServerPolicyAllOfLinks struct for AuthorizationServerPolicyAllOfLinks
type AuthorizationServerPolicyAllOfLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Activate *HrefObjectActivateLink `json:"activate,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	// Link to the authorization server policy's rules
	Rules *HrefObject `json:"rules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyAllOfLinks AuthorizationServerPolicyAllOfLinks

// NewAuthorizationServerPolicyAllOfLinks instantiates a new AuthorizationServerPolicyAllOfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyAllOfLinks() *AuthorizationServerPolicyAllOfLinks {
	this := AuthorizationServerPolicyAllOfLinks{}
	return &this
}

// NewAuthorizationServerPolicyAllOfLinksWithDefaults instantiates a new AuthorizationServerPolicyAllOfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyAllOfLinksWithDefaults() *AuthorizationServerPolicyAllOfLinks {
	this := AuthorizationServerPolicyAllOfLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyAllOfLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyAllOfLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyAllOfLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *AuthorizationServerPolicyAllOfLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyAllOfLinks) GetActivate() HrefObjectActivateLink {
	if o == nil || o.Activate == nil {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyAllOfLinks) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyAllOfLinks) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *AuthorizationServerPolicyAllOfLinks) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyAllOfLinks) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || o.Deactivate == nil {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyAllOfLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyAllOfLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *AuthorizationServerPolicyAllOfLinks) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyAllOfLinks) GetRules() HrefObject {
	if o == nil || o.Rules == nil {
		var ret HrefObject
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyAllOfLinks) GetRulesOk() (*HrefObject, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyAllOfLinks) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given HrefObject and assigns it to the Rules field.
func (o *AuthorizationServerPolicyAllOfLinks) SetRules(v HrefObject) {
	o.Rules = &v
}

func (o AuthorizationServerPolicyAllOfLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicyAllOfLinks) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicyAllOfLinks := _AuthorizationServerPolicyAllOfLinks{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyAllOfLinks)
	if err == nil {
		*o = AuthorizationServerPolicyAllOfLinks(varAuthorizationServerPolicyAllOfLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicyAllOfLinks struct {
	value *AuthorizationServerPolicyAllOfLinks
	isSet bool
}

func (v NullableAuthorizationServerPolicyAllOfLinks) Get() *AuthorizationServerPolicyAllOfLinks {
	return v.value
}

func (v *NullableAuthorizationServerPolicyAllOfLinks) Set(val *AuthorizationServerPolicyAllOfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyAllOfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyAllOfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyAllOfLinks(val *AuthorizationServerPolicyAllOfLinks) *NullableAuthorizationServerPolicyAllOfLinks {
	return &NullableAuthorizationServerPolicyAllOfLinks{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyAllOfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyAllOfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

