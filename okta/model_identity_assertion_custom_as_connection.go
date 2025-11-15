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

// checks if the IdentityAssertionCustomASConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityAssertionCustomASConnection{}

// IdentityAssertionCustomASConnection Identity assertion connection for a custom authorization server
type IdentityAssertionCustomASConnection struct {
	AuthorizationServer CustomAuthorizationServer `json:"authorizationServer"`
	// Type of connection authentication method
	ConnectionType string `json:"connectionType"`
	// Unique identifier for the managed connection. Only present for managed connections.
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection
	Orn *string `json:"orn,omitempty"`
	// The authentication protocol type used for the connection
	ProtocolType *string `json:"protocolType,omitempty"`
	// Determines how Okta evaluates requested scopes for the connection.
	ScopeCondition *string `json:"scopeCondition,omitempty"`
	// List of scopes for the connection based on scopeCondition
	Scopes []string `json:"scopes,omitempty"`
	// The status of the connection
	Status *string    `json:"status,omitempty"`
	Links  *LinksSelf `json:"_links,omitempty"`
}

type _IdentityAssertionCustomASConnection IdentityAssertionCustomASConnection

// NewIdentityAssertionCustomASConnection instantiates a new IdentityAssertionCustomASConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityAssertionCustomASConnection(authorizationServer CustomAuthorizationServer, connectionType string) *IdentityAssertionCustomASConnection {
	this := IdentityAssertionCustomASConnection{}
	this.AuthorizationServer = authorizationServer
	this.ConnectionType = connectionType
	return &this
}

// NewIdentityAssertionCustomASConnectionWithDefaults instantiates a new IdentityAssertionCustomASConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityAssertionCustomASConnectionWithDefaults() *IdentityAssertionCustomASConnection {
	this := IdentityAssertionCustomASConnection{}
	return &this
}

// GetAuthorizationServer returns the AuthorizationServer field value
func (o *IdentityAssertionCustomASConnection) GetAuthorizationServer() CustomAuthorizationServer {
	if o == nil {
		var ret CustomAuthorizationServer
		return ret
	}

	return o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnection) GetAuthorizationServerOk() (*CustomAuthorizationServer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationServer, true
}

// SetAuthorizationServer sets field value
func (o *IdentityAssertionCustomASConnection) SetAuthorizationServer(v CustomAuthorizationServer) {
	o.AuthorizationServer = v
}

// GetConnectionType returns the ConnectionType field value
func (o *IdentityAssertionCustomASConnection) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnection) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *IdentityAssertionCustomASConnection) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityAssertionCustomASConnection) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnection) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnection) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnection) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *IdentityAssertionCustomASConnection) SetOrn(v string) {
	o.Orn = &v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnection) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnection) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnection) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *IdentityAssertionCustomASConnection) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetScopeCondition returns the ScopeCondition field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnection) GetScopeCondition() string {
	if o == nil || IsNil(o.ScopeCondition) {
		var ret string
		return ret
	}
	return *o.ScopeCondition
}

// GetScopeConditionOk returns a tuple with the ScopeCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnection) GetScopeConditionOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeCondition) {
		return nil, false
	}
	return o.ScopeCondition, true
}

// HasScopeCondition returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnection) HasScopeCondition() bool {
	if o != nil && !IsNil(o.ScopeCondition) {
		return true
	}

	return false
}

// SetScopeCondition gets a reference to the given string and assigns it to the ScopeCondition field.
func (o *IdentityAssertionCustomASConnection) SetScopeCondition(v string) {
	o.ScopeCondition = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnection) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnection) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnection) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *IdentityAssertionCustomASConnection) SetScopes(v []string) {
	o.Scopes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IdentityAssertionCustomASConnection) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IdentityAssertionCustomASConnection) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityAssertionCustomASConnection) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IdentityAssertionCustomASConnection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *IdentityAssertionCustomASConnection) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o IdentityAssertionCustomASConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityAssertionCustomASConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorizationServer"] = o.AuthorizationServer
	toSerialize["connectionType"] = o.ConnectionType
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.ProtocolType) {
		toSerialize["protocolType"] = o.ProtocolType
	}
	if !IsNil(o.ScopeCondition) {
		toSerialize["scopeCondition"] = o.ScopeCondition
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

func (o *IdentityAssertionCustomASConnection) UnmarshalJSON(data []byte) (err error) {
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

	varIdentityAssertionCustomASConnection := _IdentityAssertionCustomASConnection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentityAssertionCustomASConnection)

	if err != nil {
		return err
	}

	*o = IdentityAssertionCustomASConnection(varIdentityAssertionCustomASConnection)

	return err
}

type NullableIdentityAssertionCustomASConnection struct {
	value *IdentityAssertionCustomASConnection
	isSet bool
}

func (v NullableIdentityAssertionCustomASConnection) Get() *IdentityAssertionCustomASConnection {
	return v.value
}

func (v *NullableIdentityAssertionCustomASConnection) Set(val *IdentityAssertionCustomASConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityAssertionCustomASConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityAssertionCustomASConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityAssertionCustomASConnection(val *IdentityAssertionCustomASConnection) *NullableIdentityAssertionCustomASConnection {
	return &NullableIdentityAssertionCustomASConnection{value: val, isSet: true}
}

func (v NullableIdentityAssertionCustomASConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityAssertionCustomASConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
