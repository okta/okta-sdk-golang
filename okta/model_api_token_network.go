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

// ApiTokenNetwork The Network Condition of the API Token
type ApiTokenNetwork struct {
	// The connection type of the Network Condition
	Connection *string `json:"connection,omitempty"`
	// List of included IP network zones
	Include []string `json:"include,omitempty"`
	// List of excluded IP network zones
	Exclude []string `json:"exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiTokenNetwork ApiTokenNetwork

// NewApiTokenNetwork instantiates a new ApiTokenNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTokenNetwork() *ApiTokenNetwork {
	this := ApiTokenNetwork{}
	return &this
}

// NewApiTokenNetworkWithDefaults instantiates a new ApiTokenNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenNetworkWithDefaults() *ApiTokenNetwork {
	this := ApiTokenNetwork{}
	return &this
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *ApiTokenNetwork) GetConnection() string {
	if o == nil || o.Connection == nil {
		var ret string
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenNetwork) GetConnectionOk() (*string, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *ApiTokenNetwork) HasConnection() bool {
	if o != nil && o.Connection != nil {
		return true
	}

	return false
}

// SetConnection gets a reference to the given string and assigns it to the Connection field.
func (o *ApiTokenNetwork) SetConnection(v string) {
	o.Connection = &v
}

// GetInclude returns the Include field value if set, zero value otherwise.
func (o *ApiTokenNetwork) GetInclude() []string {
	if o == nil || o.Include == nil {
		var ret []string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenNetwork) GetIncludeOk() ([]string, bool) {
	if o == nil || o.Include == nil {
		return nil, false
	}
	return o.Include, true
}

// HasInclude returns a boolean if a field has been set.
func (o *ApiTokenNetwork) HasInclude() bool {
	if o != nil && o.Include != nil {
		return true
	}

	return false
}

// SetInclude gets a reference to the given []string and assigns it to the Include field.
func (o *ApiTokenNetwork) SetInclude(v []string) {
	o.Include = v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *ApiTokenNetwork) GetExclude() []string {
	if o == nil || o.Exclude == nil {
		var ret []string
		return ret
	}
	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenNetwork) GetExcludeOk() ([]string, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *ApiTokenNetwork) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given []string and assigns it to the Exclude field.
func (o *ApiTokenNetwork) SetExclude(v []string) {
	o.Exclude = v
}

func (o ApiTokenNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	if o.Include != nil {
		toSerialize["include"] = o.Include
	}
	if o.Exclude != nil {
		toSerialize["exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApiTokenNetwork) UnmarshalJSON(bytes []byte) (err error) {
	varApiTokenNetwork := _ApiTokenNetwork{}

	err = json.Unmarshal(bytes, &varApiTokenNetwork)
	if err == nil {
		*o = ApiTokenNetwork(varApiTokenNetwork)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "connection")
		delete(additionalProperties, "include")
		delete(additionalProperties, "exclude")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApiTokenNetwork struct {
	value *ApiTokenNetwork
	isSet bool
}

func (v NullableApiTokenNetwork) Get() *ApiTokenNetwork {
	return v.value
}

func (v *NullableApiTokenNetwork) Set(val *ApiTokenNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTokenNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTokenNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTokenNetwork(val *ApiTokenNetwork) *NullableApiTokenNetwork {
	return &NullableApiTokenNetwork{value: val, isSet: true}
}

func (v NullableApiTokenNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTokenNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

