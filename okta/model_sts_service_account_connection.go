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

// checks if the STSServiceAccountConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &STSServiceAccountConnection{}

// STSServiceAccountConnection STS connection to a service account
type STSServiceAccountConnection struct {
	App ManagedConnectionAppInstance `json:"app"`
	// Type of connection authentication method
	ConnectionType string `json:"connectionType"`
	// Unique identifier for the managed connection. Only present for managed connections.
	Id *string `json:"id,omitempty"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection
	Orn *string `json:"orn,omitempty"`
	// The authentication protocol type used for the connection
	ProtocolType   *string                         `json:"protocolType,omitempty"`
	ServiceAccount ManagedConnectionServiceAccount `json:"serviceAccount"`
	// The status of the connection
	Status *string    `json:"status,omitempty"`
	Links  *LinksSelf `json:"_links,omitempty"`
}

type _STSServiceAccountConnection STSServiceAccountConnection

// NewSTSServiceAccountConnection instantiates a new STSServiceAccountConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSTSServiceAccountConnection(app ManagedConnectionAppInstance, connectionType string, serviceAccount ManagedConnectionServiceAccount) *STSServiceAccountConnection {
	this := STSServiceAccountConnection{}
	this.App = app
	this.ConnectionType = connectionType
	this.ServiceAccount = serviceAccount
	return &this
}

// NewSTSServiceAccountConnectionWithDefaults instantiates a new STSServiceAccountConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSTSServiceAccountConnectionWithDefaults() *STSServiceAccountConnection {
	this := STSServiceAccountConnection{}
	return &this
}

// GetApp returns the App field value
func (o *STSServiceAccountConnection) GetApp() ManagedConnectionAppInstance {
	if o == nil {
		var ret ManagedConnectionAppInstance
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnection) GetAppOk() (*ManagedConnectionAppInstance, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *STSServiceAccountConnection) SetApp(v ManagedConnectionAppInstance) {
	o.App = v
}

// GetConnectionType returns the ConnectionType field value
func (o *STSServiceAccountConnection) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnection) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *STSServiceAccountConnection) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *STSServiceAccountConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *STSServiceAccountConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *STSServiceAccountConnection) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *STSServiceAccountConnection) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnection) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *STSServiceAccountConnection) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *STSServiceAccountConnection) SetOrn(v string) {
	o.Orn = &v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *STSServiceAccountConnection) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnection) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *STSServiceAccountConnection) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *STSServiceAccountConnection) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetServiceAccount returns the ServiceAccount field value
func (o *STSServiceAccountConnection) GetServiceAccount() ManagedConnectionServiceAccount {
	if o == nil {
		var ret ManagedConnectionServiceAccount
		return ret
	}

	return o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnection) GetServiceAccountOk() (*ManagedConnectionServiceAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccount, true
}

// SetServiceAccount sets field value
func (o *STSServiceAccountConnection) SetServiceAccount(v ManagedConnectionServiceAccount) {
	o.ServiceAccount = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *STSServiceAccountConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *STSServiceAccountConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *STSServiceAccountConnection) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *STSServiceAccountConnection) GetLinks() LinksSelf {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnection) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *STSServiceAccountConnection) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *STSServiceAccountConnection) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o STSServiceAccountConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o STSServiceAccountConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
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
	toSerialize["serviceAccount"] = o.ServiceAccount
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

func (o *STSServiceAccountConnection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app",
		"connectionType",
		"serviceAccount",
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

	varSTSServiceAccountConnection := _STSServiceAccountConnection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSTSServiceAccountConnection)

	if err != nil {
		return err
	}

	*o = STSServiceAccountConnection(varSTSServiceAccountConnection)

	return err
}

type NullableSTSServiceAccountConnection struct {
	value *STSServiceAccountConnection
	isSet bool
}

func (v NullableSTSServiceAccountConnection) Get() *STSServiceAccountConnection {
	return v.value
}

func (v *NullableSTSServiceAccountConnection) Set(val *STSServiceAccountConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableSTSServiceAccountConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableSTSServiceAccountConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSTSServiceAccountConnection(val *STSServiceAccountConnection) *NullableSTSServiceAccountConnection {
	return &NullableSTSServiceAccountConnection{value: val, isSet: true}
}

func (v NullableSTSServiceAccountConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSTSServiceAccountConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
