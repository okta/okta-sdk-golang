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
	"encoding/json"
	"fmt"
)

// checks if the IDVEndpoints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IDVEndpoints{}

// IDVEndpoints Contains endpoints for the IDV vendor. When you create an `IDV_STANDARD` IdP, you must include the `par`, `authorization`, `token`, and `jwks` endpoints in the request body.
type IDVEndpoints struct {
	Authorization        IDVAuthorizationEndpoint `json:"authorization"`
	Jwks                 OidcJwksEndpoint         `json:"jwks"`
	Par                  IDVParEndpoint           `json:"par"`
	Token                IDVTokenEndpoint         `json:"token"`
	AdditionalProperties map[string]interface{}
}

type _IDVEndpoints IDVEndpoints

// NewIDVEndpoints instantiates a new IDVEndpoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIDVEndpoints(authorization IDVAuthorizationEndpoint, jwks OidcJwksEndpoint, par IDVParEndpoint, token IDVTokenEndpoint) *IDVEndpoints {
	this := IDVEndpoints{}
	this.Authorization = authorization
	this.Jwks = jwks
	this.Par = par
	this.Token = token
	return &this
}

// NewIDVEndpointsWithDefaults instantiates a new IDVEndpoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIDVEndpointsWithDefaults() *IDVEndpoints {
	this := IDVEndpoints{}
	return &this
}

// GetAuthorization returns the Authorization field value
func (o *IDVEndpoints) GetAuthorization() IDVAuthorizationEndpoint {
	if o == nil {
		var ret IDVAuthorizationEndpoint
		return ret
	}

	return o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value
// and a boolean to check if the value has been set.
func (o *IDVEndpoints) GetAuthorizationOk() (*IDVAuthorizationEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authorization, true
}

// SetAuthorization sets field value
func (o *IDVEndpoints) SetAuthorization(v IDVAuthorizationEndpoint) {
	o.Authorization = v
}

// GetJwks returns the Jwks field value
func (o *IDVEndpoints) GetJwks() OidcJwksEndpoint {
	if o == nil {
		var ret OidcJwksEndpoint
		return ret
	}

	return o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value
// and a boolean to check if the value has been set.
func (o *IDVEndpoints) GetJwksOk() (*OidcJwksEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jwks, true
}

// SetJwks sets field value
func (o *IDVEndpoints) SetJwks(v OidcJwksEndpoint) {
	o.Jwks = v
}

// GetPar returns the Par field value
func (o *IDVEndpoints) GetPar() IDVParEndpoint {
	if o == nil {
		var ret IDVParEndpoint
		return ret
	}

	return o.Par
}

// GetParOk returns a tuple with the Par field value
// and a boolean to check if the value has been set.
func (o *IDVEndpoints) GetParOk() (*IDVParEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Par, true
}

// SetPar sets field value
func (o *IDVEndpoints) SetPar(v IDVParEndpoint) {
	o.Par = v
}

// GetToken returns the Token field value
func (o *IDVEndpoints) GetToken() IDVTokenEndpoint {
	if o == nil {
		var ret IDVTokenEndpoint
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *IDVEndpoints) GetTokenOk() (*IDVTokenEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *IDVEndpoints) SetToken(v IDVTokenEndpoint) {
	o.Token = v
}

func (o IDVEndpoints) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IDVEndpoints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorization"] = o.Authorization
	toSerialize["jwks"] = o.Jwks
	toSerialize["par"] = o.Par
	toSerialize["token"] = o.Token

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IDVEndpoints) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorization",
		"jwks",
		"par",
		"token",
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

	varIDVEndpoints := _IDVEndpoints{}

	err = json.Unmarshal(data, &varIDVEndpoints)

	if err != nil {
		return err
	}

	*o = IDVEndpoints(varIDVEndpoints)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorization")
		delete(additionalProperties, "jwks")
		delete(additionalProperties, "par")
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
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
