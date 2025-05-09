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

// AuthorizationServerPolicyConditions struct for AuthorizationServerPolicyConditions
type AuthorizationServerPolicyConditions struct {
	Clients *ClientPolicyCondition `json:"clients,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizationServerPolicyConditions AuthorizationServerPolicyConditions

// NewAuthorizationServerPolicyConditions instantiates a new AuthorizationServerPolicyConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationServerPolicyConditions() *AuthorizationServerPolicyConditions {
	this := AuthorizationServerPolicyConditions{}
	return &this
}

// NewAuthorizationServerPolicyConditionsWithDefaults instantiates a new AuthorizationServerPolicyConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationServerPolicyConditionsWithDefaults() *AuthorizationServerPolicyConditions {
	this := AuthorizationServerPolicyConditions{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *AuthorizationServerPolicyConditions) GetClients() ClientPolicyCondition {
	if o == nil || o.Clients == nil {
		var ret ClientPolicyCondition
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationServerPolicyConditions) GetClientsOk() (*ClientPolicyCondition, bool) {
	if o == nil || o.Clients == nil {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *AuthorizationServerPolicyConditions) HasClients() bool {
	if o != nil && o.Clients != nil {
		return true
	}

	return false
}

// SetClients gets a reference to the given ClientPolicyCondition and assigns it to the Clients field.
func (o *AuthorizationServerPolicyConditions) SetClients(v ClientPolicyCondition) {
	o.Clients = &v
}

func (o AuthorizationServerPolicyConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Clients != nil {
		toSerialize["clients"] = o.Clients
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthorizationServerPolicyConditions) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizationServerPolicyConditions := _AuthorizationServerPolicyConditions{}

	err = json.Unmarshal(bytes, &varAuthorizationServerPolicyConditions)
	if err == nil {
		*o = AuthorizationServerPolicyConditions(varAuthorizationServerPolicyConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "clients")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthorizationServerPolicyConditions struct {
	value *AuthorizationServerPolicyConditions
	isSet bool
}

func (v NullableAuthorizationServerPolicyConditions) Get() *AuthorizationServerPolicyConditions {
	return v.value
}

func (v *NullableAuthorizationServerPolicyConditions) Set(val *AuthorizationServerPolicyConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationServerPolicyConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationServerPolicyConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationServerPolicyConditions(val *AuthorizationServerPolicyConditions) *NullableAuthorizationServerPolicyConditions {
	return &NullableAuthorizationServerPolicyConditions{value: val, isSet: true}
}

func (v NullableAuthorizationServerPolicyConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationServerPolicyConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

