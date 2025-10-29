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
	"encoding/json"
	"fmt"
)

// checks if the OpenIdConnectApplicationNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenIdConnectApplicationNetwork{}

// OpenIdConnectApplicationNetwork <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>The network restrictions of the client
type OpenIdConnectApplicationNetwork struct {
	// The connection type of the network. Can be `ANYWHERE` or `ZONE`.
	Connection string `json:"connection"`
	// If `ZONE` is specified as a connection, then specify the excluded IP network zones here. Value can be \"ALL_IP_ZONES\" or an array of zone IDs.
	Exclude []string `json:"exclude,omitempty"`
	// If `ZONE` is specified as a connection, then specify the included IP network zones here. Value can be \"ALL_IP_ZONES\" or an array of zone IDs.
	Include              []string `json:"include,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenIdConnectApplicationNetwork OpenIdConnectApplicationNetwork

// NewOpenIdConnectApplicationNetwork instantiates a new OpenIdConnectApplicationNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenIdConnectApplicationNetwork(connection string) *OpenIdConnectApplicationNetwork {
	this := OpenIdConnectApplicationNetwork{}
	this.Connection = connection
	return &this
}

// NewOpenIdConnectApplicationNetworkWithDefaults instantiates a new OpenIdConnectApplicationNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenIdConnectApplicationNetworkWithDefaults() *OpenIdConnectApplicationNetwork {
	this := OpenIdConnectApplicationNetwork{}
	return &this
}

// GetConnection returns the Connection field value
func (o *OpenIdConnectApplicationNetwork) GetConnection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationNetwork) GetConnectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connection, true
}

// SetConnection sets field value
func (o *OpenIdConnectApplicationNetwork) SetConnection(v string) {
	o.Connection = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationNetwork) GetExclude() []string {
	if o == nil || IsNil(o.Exclude) {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationNetwork) GetExcludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationNetwork) HasExclude() bool {
	if o != nil && !IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *OpenIdConnectApplicationNetwork) SetExclude(v []string) {
	o.Exclude = v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *OpenIdConnectApplicationNetwork) GetInclude() []string {
	if o == nil || IsNil(o.Include) {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenIdConnectApplicationNetwork) GetIncludeOk() ([]string, bool) {
	if o == nil || IsNil(o.Include) {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *OpenIdConnectApplicationNetwork) HasInclude() bool {
	if o != nil && !IsNil(o.Include) {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *OpenIdConnectApplicationNetwork) SetInclude(v []string) {
	o.Include = v
}

func (o OpenIdConnectApplicationNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenIdConnectApplicationNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connection"] = o.Connection
	if !IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !IsNil(o.Include) {
		toSerialize["include"] = o.Include
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenIdConnectApplicationNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connection",
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

	varOpenIdConnectApplicationNetwork := _OpenIdConnectApplicationNetwork{}

	err = json.Unmarshal(data, &varOpenIdConnectApplicationNetwork)

	if err != nil {
		return err
	}

	*o = OpenIdConnectApplicationNetwork(varOpenIdConnectApplicationNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connection")
		delete(additionalProperties, "exclude")
		delete(additionalProperties, "include")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenIdConnectApplicationNetwork struct {
	value *OpenIdConnectApplicationNetwork
	isSet bool
}

func (v NullableOpenIdConnectApplicationNetwork) Get() *OpenIdConnectApplicationNetwork {
	return v.value
}

func (v *NullableOpenIdConnectApplicationNetwork) Set(val *OpenIdConnectApplicationNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenIdConnectApplicationNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenIdConnectApplicationNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenIdConnectApplicationNetwork(val *OpenIdConnectApplicationNetwork) *NullableOpenIdConnectApplicationNetwork {
	return &NullableOpenIdConnectApplicationNetwork{value: val, isSet: true}
}

func (v NullableOpenIdConnectApplicationNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenIdConnectApplicationNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
