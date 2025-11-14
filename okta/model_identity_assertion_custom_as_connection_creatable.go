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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the IdentityAssertionCustomASConnectionCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAssertionCustomASConnectionCreatable{}

// IdentityAssertionCustomASConnectionCreatable Create an identity assertion connection for a custom authorization server
type IdentityAssertionCustomASConnectionCreatable struct {
	AuthorizationServer IdentityAssertionCustomASConnectionCreatableAuthorizationServer `json:"authorizationServer"`
	// Type of connection authentication method
	ConnectionType string `json:"connectionType"`
	// The authentication protocol type used for the connection
	ProtocolType *string `json:"protocolType,omitempty"`
	// Determines how Okta evaluates requested scopes for the connection.
	ScopeCondition *string `json:"scopeCondition,omitempty"`
	// Optional array of scopes
	Scopes []string `json:"scopes,omitempty"`
}

type _IdentityAssertionCustomASConnectionCreatable IdentityAssertionCustomASConnectionCreatable

// NewIdentityAssertionCustomASConnectionCreatable instantiates a new IdentityAssertionCustomASConnectionCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAssertionCustomASConnectionCreatable(authorizationServer IdentityAssertionCustomASConnectionCreatableAuthorizationServer, connectionType string) *IdentityAssertionCustomASConnectionCreatable {
	this := IdentityAssertionCustomASConnectionCreatable{}
	this.AuthorizationServer = authorizationServer
	this.ConnectionType = connectionType
	return &this
}

// NewIdentityAssertionCustomASConnectionCreatableWithDefaults instantiates a new IdentityAssertionCustomASConnectionCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAssertionCustomASConnectionCreatableWithDefaults() *IdentityAssertionCustomASConnectionCreatable {
	this := IdentityAssertionCustomASConnectionCreatable{}
	return &this
}

// GetAuthorizationServer returns the AuthorizationServer field value
func (o *IdentityAssertionCustomASConnectionCreatable) GetAuthorizationServer() IdentityAssertionCustomASConnectionCreatableAuthorizationServer {
	if o == nil {
		var ret IdentityAssertionCustomASConnectionCreatableAuthorizationServer
		return ret
	}

	return o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnectionCreatable) GetAuthorizationServerOk() (*IdentityAssertionCustomASConnectionCreatableAuthorizationServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationServer, true
}

// SetAuthorizationServer sets field value
func (o *IdentityAssertionCustomASConnectionCreatable) SetAuthorizationServer(v IdentityAssertionCustomASConnectionCreatableAuthorizationServer) {
	o.AuthorizationServer = v
}

// GetConnectionType returns the ConnectionType field value
func (o *IdentityAssertionCustomASConnectionCreatable) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnectionCreatable) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *IdentityAssertionCustomASConnectionCreatable) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnectionCreatable) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnectionCreatable) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnectionCreatable) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *IdentityAssertionCustomASConnectionCreatable) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetScopeCondition returns the ScopeCondition field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnectionCreatable) GetScopeCondition() string {
	if o == nil || IsNil(o.ScopeCondition) {
		var ret string
		return ret
	}
	return *o.ScopeCondition
}

// GetScopeConditionOk returns a tuple with the ScopeCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnectionCreatable) GetScopeConditionOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeCondition) {
		return nil, false
	}
	return o.ScopeCondition, true
}

// HasScopeCondition returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnectionCreatable) HasScopeCondition() bool {
	if o != nil && !IsNil(o.ScopeCondition) {
		return true
	}

	return false
}

// SetScopeCondition gets a reference to the given string and assigns it to the ScopeCondition field.
func (o *IdentityAssertionCustomASConnectionCreatable) SetScopeCondition(v string) {
	o.ScopeCondition = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnectionCreatable) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnectionCreatable) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnectionCreatable) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *IdentityAssertionCustomASConnectionCreatable) SetScopes(v []string) {
	o.Scopes = v
}

func (o IdentityAssertionCustomASConnectionCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAssertionCustomASConnectionCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorizationServer"] = o.AuthorizationServer
	toSerialize["connectionType"] = o.ConnectionType
	if !IsNil(o.ProtocolType) {
		toSerialize["protocolType"] = o.ProtocolType
	}
	if !IsNil(o.ScopeCondition) {
		toSerialize["scopeCondition"] = o.ScopeCondition
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

func (o *IdentityAssertionCustomASConnectionCreatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorizationServer",
		"connectionType",
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

	varIdentityAssertionCustomASConnectionCreatable := _IdentityAssertionCustomASConnectionCreatable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentityAssertionCustomASConnectionCreatable)

	if err != nil {
		return err
	}

	*o = IdentityAssertionCustomASConnectionCreatable(varIdentityAssertionCustomASConnectionCreatable)

	return err
}

type NullableIdentityAssertionCustomASConnectionCreatable struct {
	value *IdentityAssertionCustomASConnectionCreatable
	isSet bool
}

func (v NullableIdentityAssertionCustomASConnectionCreatable) Get() *IdentityAssertionCustomASConnectionCreatable {
	return v.value
}

func (v *NullableIdentityAssertionCustomASConnectionCreatable) Set(val *IdentityAssertionCustomASConnectionCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAssertionCustomASConnectionCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAssertionCustomASConnectionCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAssertionCustomASConnectionCreatable(val *IdentityAssertionCustomASConnectionCreatable) *NullableIdentityAssertionCustomASConnectionCreatable {
	return &NullableIdentityAssertionCustomASConnectionCreatable{value: val, isSet: true}
}

func (v NullableIdentityAssertionCustomASConnectionCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAssertionCustomASConnectionCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
