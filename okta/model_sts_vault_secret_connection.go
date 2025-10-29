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
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the STSVaultSecretConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &STSVaultSecretConnection{}

// STSVaultSecretConnection STS connection to a vaulted secret
type STSVaultSecretConnection struct {
	// Type of connection authentication method
	ConnectionType string `json:"connectionType"`
	// Unique identifier for the managed connection. Only present for managed connections.
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection
	Orn *string `json:"orn,omitempty"`
	// The authentication protocol type used for the connection
	ProtocolType *string                        `json:"protocolType,omitempty"`
	Secret       ManagedConnectionVaultedSecret `json:"secret"`
	// The status of the connection
	Status *string    `json:"status,omitempty"`
	Links  *LinksSelf `json:"_links,omitempty"`
}

type _STSVaultSecretConnection STSVaultSecretConnection

// NewSTSVaultSecretConnection instantiates a new STSVaultSecretConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSTSVaultSecretConnection(connectionType string, secret ManagedConnectionVaultedSecret) *STSVaultSecretConnection {
	this := STSVaultSecretConnection{}
	this.ConnectionType = connectionType
	this.Secret = secret
	return &this
}

// NewSTSVaultSecretConnectionWithDefaults instantiates a new STSVaultSecretConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSTSVaultSecretConnectionWithDefaults() *STSVaultSecretConnection {
	this := STSVaultSecretConnection{}
	return &this
}

// GetConnectionType returns the ConnectionType field value
func (o *STSVaultSecretConnection) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnection) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *STSVaultSecretConnection) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *STSVaultSecretConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *STSVaultSecretConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *STSVaultSecretConnection) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *STSVaultSecretConnection) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnection) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *STSVaultSecretConnection) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *STSVaultSecretConnection) SetOrn(v string) {
	o.Orn = &v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *STSVaultSecretConnection) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnection) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *STSVaultSecretConnection) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *STSVaultSecretConnection) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetSecret returns the Secret field value
func (o *STSVaultSecretConnection) GetSecret() ManagedConnectionVaultedSecret {
	if o == nil {
		var ret ManagedConnectionVaultedSecret
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnection) GetSecretOk() (*ManagedConnectionVaultedSecret, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *STSVaultSecretConnection) SetSecret(v ManagedConnectionVaultedSecret) {
	o.Secret = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *STSVaultSecretConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *STSVaultSecretConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *STSVaultSecretConnection) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *STSVaultSecretConnection) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSVaultSecretConnection) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *STSVaultSecretConnection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *STSVaultSecretConnection) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o STSVaultSecretConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o STSVaultSecretConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	toSerialize["secret"] = o.Secret
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

func (o *STSVaultSecretConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectionType",
		"secret",
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

	varSTSVaultSecretConnection := _STSVaultSecretConnection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSTSVaultSecretConnection)

	if err != nil {
		return err
	}

	*o = STSVaultSecretConnection(varSTSVaultSecretConnection)

	return err
}

type NullableSTSVaultSecretConnection struct {
	value *STSVaultSecretConnection
	isSet bool
}

func (v NullableSTSVaultSecretConnection) Get() *STSVaultSecretConnection {
	return v.value
}

func (v *NullableSTSVaultSecretConnection) Set(val *STSVaultSecretConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableSTSVaultSecretConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableSTSVaultSecretConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSTSVaultSecretConnection(val *STSVaultSecretConnection) *NullableSTSVaultSecretConnection {
	return &NullableSTSVaultSecretConnection{value: val, isSet: true}
}

func (v NullableSTSVaultSecretConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSTSVaultSecretConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
