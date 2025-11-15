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

// checks if the IdentityAssertionAppInstanceConnectionCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAssertionAppInstanceConnectionCreatable{}

// IdentityAssertionAppInstanceConnectionCreatable Create an identity assertion connection for an app instance
type IdentityAssertionAppInstanceConnectionCreatable struct {
	App                 IdentityAssertionAppInstanceConnectionCreatableApp                 `json:"app"`
	AuthorizationServer IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer `json:"authorizationServer"`
	// Type of connection authentication method
	ConnectionType string `json:"connectionType"`
	// The authentication protocol type used for the connection
	ProtocolType *string `json:"protocolType,omitempty"`
	// Determines how Okta evaluates requested scopes for the connection.
	ScopeCondition *string `json:"scopeCondition,omitempty"`
	// Optional array of scopes
	Scopes []string `json:"scopes,omitempty"`
}

type _IdentityAssertionAppInstanceConnectionCreatable IdentityAssertionAppInstanceConnectionCreatable

// NewIdentityAssertionAppInstanceConnectionCreatable instantiates a new IdentityAssertionAppInstanceConnectionCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAssertionAppInstanceConnectionCreatable(app IdentityAssertionAppInstanceConnectionCreatableApp, authorizationServer IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer, connectionType string) *IdentityAssertionAppInstanceConnectionCreatable {
	this := IdentityAssertionAppInstanceConnectionCreatable{}
	this.App = app
	this.AuthorizationServer = authorizationServer
	this.ConnectionType = connectionType
	return &this
}

// NewIdentityAssertionAppInstanceConnectionCreatableWithDefaults instantiates a new IdentityAssertionAppInstanceConnectionCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAssertionAppInstanceConnectionCreatableWithDefaults() *IdentityAssertionAppInstanceConnectionCreatable {
	this := IdentityAssertionAppInstanceConnectionCreatable{}
	return &this
}

// GetApp returns the App field value
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetApp() IdentityAssertionAppInstanceConnectionCreatableApp {
	if o == nil {
		var ret IdentityAssertionAppInstanceConnectionCreatableApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetAppOk() (*IdentityAssertionAppInstanceConnectionCreatableApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *IdentityAssertionAppInstanceConnectionCreatable) SetApp(v IdentityAssertionAppInstanceConnectionCreatableApp) {
	o.App = v
}

// GetAuthorizationServer returns the AuthorizationServer field value
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetAuthorizationServer() IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer {
	if o == nil {
		var ret IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer
		return ret
	}

	return o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetAuthorizationServerOk() (*IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationServer, true
}

// SetAuthorizationServer sets field value
func (o *IdentityAssertionAppInstanceConnectionCreatable) SetAuthorizationServer(v IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer) {
	o.AuthorizationServer = v
}

// GetConnectionType returns the ConnectionType field value
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *IdentityAssertionAppInstanceConnectionCreatable) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *IdentityAssertionAppInstanceConnectionCreatable) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *IdentityAssertionAppInstanceConnectionCreatable) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetScopeCondition returns the ScopeCondition field value if set, zero value otherwise.
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetScopeCondition() string {
	if o == nil || IsNil(o.ScopeCondition) {
		var ret string
		return ret
	}
	return *o.ScopeCondition
}

// GetScopeConditionOk returns a tuple with the ScopeCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetScopeConditionOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeCondition) {
		return nil, false
	}
	return o.ScopeCondition, true
}

// HasScopeCondition returns a boolean if a field has been set.
func (o *IdentityAssertionAppInstanceConnectionCreatable) HasScopeCondition() bool {
	if o != nil && !IsNil(o.ScopeCondition) {
		return true
	}

	return false
}

// SetScopeCondition gets a reference to the given string and assigns it to the ScopeCondition field.
func (o *IdentityAssertionAppInstanceConnectionCreatable) SetScopeCondition(v string) {
	o.ScopeCondition = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionAppInstanceConnectionCreatable) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *IdentityAssertionAppInstanceConnectionCreatable) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *IdentityAssertionAppInstanceConnectionCreatable) SetScopes(v []string) {
	o.Scopes = v
}

func (o IdentityAssertionAppInstanceConnectionCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAssertionAppInstanceConnectionCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
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

func (o *IdentityAssertionAppInstanceConnectionCreatable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app",
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

	varIdentityAssertionAppInstanceConnectionCreatable := _IdentityAssertionAppInstanceConnectionCreatable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentityAssertionAppInstanceConnectionCreatable)

	if err != nil {
		return err
	}

	*o = IdentityAssertionAppInstanceConnectionCreatable(varIdentityAssertionAppInstanceConnectionCreatable)

	return err
}

type NullableIdentityAssertionAppInstanceConnectionCreatable struct {
	value *IdentityAssertionAppInstanceConnectionCreatable
	isSet bool
}

func (v NullableIdentityAssertionAppInstanceConnectionCreatable) Get() *IdentityAssertionAppInstanceConnectionCreatable {
	return v.value
}

func (v *NullableIdentityAssertionAppInstanceConnectionCreatable) Set(val *IdentityAssertionAppInstanceConnectionCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAssertionAppInstanceConnectionCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAssertionAppInstanceConnectionCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAssertionAppInstanceConnectionCreatable(val *IdentityAssertionAppInstanceConnectionCreatable) *NullableIdentityAssertionAppInstanceConnectionCreatable {
	return &NullableIdentityAssertionAppInstanceConnectionCreatable{value: val, isSet: true}
}

func (v NullableIdentityAssertionAppInstanceConnectionCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAssertionAppInstanceConnectionCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
