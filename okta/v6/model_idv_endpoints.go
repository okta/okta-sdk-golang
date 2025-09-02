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

// IDVEndpoints Contains the endpoints for the IDV
type IDVEndpoints struct {
	Authorization *IDVAuthorizationEndpoint `json:"authorization,omitempty"`
	Par *IDVParEndpoint `json:"par,omitempty"`
	Token *IDVTokenEndpoint `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IDVEndpoints IDVEndpoints

// NewIDVEndpoints instantiates a new IDVEndpoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDVEndpoints() *IDVEndpoints {
	this := IDVEndpoints{}
	return &this
}

// NewIDVEndpointsWithDefaults instantiates a new IDVEndpoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDVEndpointsWithDefaults() *IDVEndpoints {
	this := IDVEndpoints{}
	return &this
}

// GetAuthorization returns the Authorization field value if set, zero value otherwise.
func (o *IDVEndpoints) GetAuthorization() IDVAuthorizationEndpoint {
	if o == nil || o.Authorization == nil {
		var ret IDVAuthorizationEndpoint
		return ret
	}
	return *o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVEndpoints) GetAuthorizationOk() (*IDVAuthorizationEndpoint, bool) {
	if o == nil || o.Authorization == nil {
		return nil, false
	}
	return o.Authorization, true
}

// HasAuthorization returns a boolean if a field has been set.
func (o *IDVEndpoints) HasAuthorization() bool {
	if o != nil && o.Authorization != nil {
		return true
	}

	return false
}

// SetAuthorization gets a reference to the given IDVAuthorizationEndpoint and assigns it to the Authorization field.
func (o *IDVEndpoints) SetAuthorization(v IDVAuthorizationEndpoint) {
	o.Authorization = &v
}

// GetPar returns the Par field value if set, zero value otherwise.
func (o *IDVEndpoints) GetPar() IDVParEndpoint {
	if o == nil || o.Par == nil {
		var ret IDVParEndpoint
		return ret
	}
	return *o.Par
}

// GetParOk returns a tuple with the Par field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVEndpoints) GetParOk() (*IDVParEndpoint, bool) {
	if o == nil || o.Par == nil {
		return nil, false
	}
	return o.Par, true
}

// HasPar returns a boolean if a field has been set.
func (o *IDVEndpoints) HasPar() bool {
	if o != nil && o.Par != nil {
		return true
	}

	return false
}

// SetPar gets a reference to the given IDVParEndpoint and assigns it to the Par field.
func (o *IDVEndpoints) SetPar(v IDVParEndpoint) {
	o.Par = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *IDVEndpoints) GetToken() IDVTokenEndpoint {
	if o == nil || o.Token == nil {
		var ret IDVTokenEndpoint
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IDVEndpoints) GetTokenOk() (*IDVTokenEndpoint, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *IDVEndpoints) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given IDVTokenEndpoint and assigns it to the Token field.
func (o *IDVEndpoints) SetToken(v IDVTokenEndpoint) {
	o.Token = &v
}

func (o IDVEndpoints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authorization != nil {
		toSerialize["authorization"] = o.Authorization
	}
	if o.Par != nil {
		toSerialize["par"] = o.Par
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IDVEndpoints) UnmarshalJSON(bytes []byte) (err error) {
	varIDVEndpoints := _IDVEndpoints{}

	err = json.Unmarshal(bytes, &varIDVEndpoints)
	if err == nil {
		*o = IDVEndpoints(varIDVEndpoints)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authorization")
		delete(additionalProperties, "par")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIDVEndpoints struct {
	value *IDVEndpoints
	isSet bool
}

func (v NullableIDVEndpoints) Get() *IDVEndpoints {
	return v.value
}

func (v *NullableIDVEndpoints) Set(val *IDVEndpoints) {
	v.value = val
	v.isSet = true
}

func (v NullableIDVEndpoints) IsSet() bool {
	return v.isSet
}

func (v *NullableIDVEndpoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIDVEndpoints(val *IDVEndpoints) *NullableIDVEndpoints {
	return &NullableIDVEndpoints{value: val, isSet: true}
}

func (v NullableIDVEndpoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIDVEndpoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

