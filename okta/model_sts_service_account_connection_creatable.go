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

// checks if the STSServiceAccountConnectionCreatable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &STSServiceAccountConnectionCreatable{}

// STSServiceAccountConnectionCreatable Create an STS connection for a service account
type STSServiceAccountConnectionCreatable struct {
	App IdentityAssertionAppInstanceConnectionCreatableApp `json:"app"`
	// Type of connection authentication method
	ConnectionType string `json:"connectionType"`
	// The authentication protocol type used for the connection
	ProtocolType   *string                                            `json:"protocolType,omitempty"`
	ServiceAccount STSServiceAccountConnectionCreatableServiceAccount `json:"serviceAccount"`
}

type _STSServiceAccountConnectionCreatable STSServiceAccountConnectionCreatable

// NewSTSServiceAccountConnectionCreatable instantiates a new STSServiceAccountConnectionCreatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSTSServiceAccountConnectionCreatable(app IdentityAssertionAppInstanceConnectionCreatableApp, connectionType string, serviceAccount STSServiceAccountConnectionCreatableServiceAccount) *STSServiceAccountConnectionCreatable {
	this := STSServiceAccountConnectionCreatable{}
	this.App = app
	this.ConnectionType = connectionType
	this.ServiceAccount = serviceAccount
	return &this
}

// NewSTSServiceAccountConnectionCreatableWithDefaults instantiates a new STSServiceAccountConnectionCreatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSTSServiceAccountConnectionCreatableWithDefaults() *STSServiceAccountConnectionCreatable {
	this := STSServiceAccountConnectionCreatable{}
	return &this
}

// GetApp returns the App field value
func (o *STSServiceAccountConnectionCreatable) GetApp() IdentityAssertionAppInstanceConnectionCreatableApp {
	if o == nil {
		var ret IdentityAssertionAppInstanceConnectionCreatableApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnectionCreatable) GetAppOk() (*IdentityAssertionAppInstanceConnectionCreatableApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *STSServiceAccountConnectionCreatable) SetApp(v IdentityAssertionAppInstanceConnectionCreatableApp) {
	o.App = v
}

// GetConnectionType returns the ConnectionType field value
func (o *STSServiceAccountConnectionCreatable) GetConnectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnectionCreatable) GetConnectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionType, true
}

// SetConnectionType sets field value
func (o *STSServiceAccountConnectionCreatable) SetConnectionType(v string) {
	o.ConnectionType = v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *STSServiceAccountConnectionCreatable) GetProtocolType() string {
	if o == nil || IsNil(o.ProtocolType) {
		var ret string
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnectionCreatable) GetProtocolTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *STSServiceAccountConnectionCreatable) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given string and assigns it to the ProtocolType field.
func (o *STSServiceAccountConnectionCreatable) SetProtocolType(v string) {
	o.ProtocolType = &v
}

// GetServiceAccount returns the ServiceAccount field value
func (o *STSServiceAccountConnectionCreatable) GetServiceAccount() STSServiceAccountConnectionCreatableServiceAccount {
	if o == nil {
		var ret STSServiceAccountConnectionCreatableServiceAccount
		return ret
	}

	return o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value
// and a boolean to check if the value has been set.
func (o *STSServiceAccountConnectionCreatable) GetServiceAccountOk() (*STSServiceAccountConnectionCreatableServiceAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccount, true
}

// SetServiceAccount sets field value
func (o *STSServiceAccountConnectionCreatable) SetServiceAccount(v STSServiceAccountConnectionCreatableServiceAccount) {
	o.ServiceAccount = v
}

func (o STSServiceAccountConnectionCreatable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o STSServiceAccountConnectionCreatable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	toSerialize["connectionType"] = o.ConnectionType
	if !IsNil(o.ProtocolType) {
		toSerialize["protocolType"] = o.ProtocolType
	}
	toSerialize["serviceAccount"] = o.ServiceAccount
	return toSerialize, nil
}

func (o *STSServiceAccountConnectionCreatable) UnmarshalJSON(data []byte) (err error) {
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

	varSTSServiceAccountConnectionCreatable := _STSServiceAccountConnectionCreatable{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSTSServiceAccountConnectionCreatable)

	if err != nil {
		return err
	}

	*o = STSServiceAccountConnectionCreatable(varSTSServiceAccountConnectionCreatable)

	return err
}

type NullableSTSServiceAccountConnectionCreatable struct {
	value *STSServiceAccountConnectionCreatable
	isSet bool
}

func (v NullableSTSServiceAccountConnectionCreatable) Get() *STSServiceAccountConnectionCreatable {
	return v.value
}

func (v *NullableSTSServiceAccountConnectionCreatable) Set(val *STSServiceAccountConnectionCreatable) {
	v.value = val
	v.isSet = true
}

func (v NullableSTSServiceAccountConnectionCreatable) IsSet() bool {
	return v.isSet
}

func (v *NullableSTSServiceAccountConnectionCreatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSTSServiceAccountConnectionCreatable(val *STSServiceAccountConnectionCreatable) *NullableSTSServiceAccountConnectionCreatable {
	return &NullableSTSServiceAccountConnectionCreatable{value: val, isSet: true}
}

func (v NullableSTSServiceAccountConnectionCreatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSTSServiceAccountConnectionCreatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
