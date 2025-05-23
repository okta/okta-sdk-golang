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

// ProtocolEndpoints struct for ProtocolEndpoints
type ProtocolEndpoints struct {
	Acs *ProtocolEndpoint `json:"acs,omitempty"`
	Authorization *ProtocolEndpoint `json:"authorization,omitempty"`
	Jwks *ProtocolEndpoint `json:"jwks,omitempty"`
	Metadata *ProtocolEndpoint `json:"metadata,omitempty"`
	Slo *ProtocolEndpoint `json:"slo,omitempty"`
	Sso *ProtocolEndpoint `json:"sso,omitempty"`
	Token *ProtocolEndpoint `json:"token,omitempty"`
	UserInfo *ProtocolEndpoint `json:"userInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolEndpoints ProtocolEndpoints

// NewProtocolEndpoints instantiates a new ProtocolEndpoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolEndpoints() *ProtocolEndpoints {
	this := ProtocolEndpoints{}
	return &this
}

// NewProtocolEndpointsWithDefaults instantiates a new ProtocolEndpoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolEndpointsWithDefaults() *ProtocolEndpoints {
	this := ProtocolEndpoints{}
	return &this
}

// GetAcs returns the Acs field value if set, zero value otherwise.
func (o *ProtocolEndpoints) GetAcs() ProtocolEndpoint {
	if o == nil || o.Acs == nil {
		var ret ProtocolEndpoint
		return ret
	}
	return *o.Acs
}

// GetAcsOk returns a tuple with the Acs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoints) GetAcsOk() (*ProtocolEndpoint, bool) {
	if o == nil || o.Acs == nil {
		return nil, false
	}
	return o.Acs, true
}

// HasAcs returns a boolean if a field has been set.
func (o *ProtocolEndpoints) HasAcs() bool {
	if o != nil && o.Acs != nil {
		return true
	}

	return false
}

// SetAcs gets a reference to the given ProtocolEndpoint and assigns it to the Acs field.
func (o *ProtocolEndpoints) SetAcs(v ProtocolEndpoint) {
	o.Acs = &v
}

// GetAuthorization returns the Authorization field value if set, zero value otherwise.
func (o *ProtocolEndpoints) GetAuthorization() ProtocolEndpoint {
	if o == nil || o.Authorization == nil {
		var ret ProtocolEndpoint
		return ret
	}
	return *o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoints) GetAuthorizationOk() (*ProtocolEndpoint, bool) {
	if o == nil || o.Authorization == nil {
		return nil, false
	}
	return o.Authorization, true
}

// HasAuthorization returns a boolean if a field has been set.
func (o *ProtocolEndpoints) HasAuthorization() bool {
	if o != nil && o.Authorization != nil {
		return true
	}

	return false
}

// SetAuthorization gets a reference to the given ProtocolEndpoint and assigns it to the Authorization field.
func (o *ProtocolEndpoints) SetAuthorization(v ProtocolEndpoint) {
	o.Authorization = &v
}

// GetJwks returns the Jwks field value if set, zero value otherwise.
func (o *ProtocolEndpoints) GetJwks() ProtocolEndpoint {
	if o == nil || o.Jwks == nil {
		var ret ProtocolEndpoint
		return ret
	}
	return *o.Jwks
}

// GetJwksOk returns a tuple with the Jwks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoints) GetJwksOk() (*ProtocolEndpoint, bool) {
	if o == nil || o.Jwks == nil {
		return nil, false
	}
	return o.Jwks, true
}

// HasJwks returns a boolean if a field has been set.
func (o *ProtocolEndpoints) HasJwks() bool {
	if o != nil && o.Jwks != nil {
		return true
	}

	return false
}

// SetJwks gets a reference to the given ProtocolEndpoint and assigns it to the Jwks field.
func (o *ProtocolEndpoints) SetJwks(v ProtocolEndpoint) {
	o.Jwks = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ProtocolEndpoints) GetMetadata() ProtocolEndpoint {
	if o == nil || o.Metadata == nil {
		var ret ProtocolEndpoint
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoints) GetMetadataOk() (*ProtocolEndpoint, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ProtocolEndpoints) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ProtocolEndpoint and assigns it to the Metadata field.
func (o *ProtocolEndpoints) SetMetadata(v ProtocolEndpoint) {
	o.Metadata = &v
}

// GetSlo returns the Slo field value if set, zero value otherwise.
func (o *ProtocolEndpoints) GetSlo() ProtocolEndpoint {
	if o == nil || o.Slo == nil {
		var ret ProtocolEndpoint
		return ret
	}
	return *o.Slo
}

// GetSloOk returns a tuple with the Slo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoints) GetSloOk() (*ProtocolEndpoint, bool) {
	if o == nil || o.Slo == nil {
		return nil, false
	}
	return o.Slo, true
}

// HasSlo returns a boolean if a field has been set.
func (o *ProtocolEndpoints) HasSlo() bool {
	if o != nil && o.Slo != nil {
		return true
	}

	return false
}

// SetSlo gets a reference to the given ProtocolEndpoint and assigns it to the Slo field.
func (o *ProtocolEndpoints) SetSlo(v ProtocolEndpoint) {
	o.Slo = &v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *ProtocolEndpoints) GetSso() ProtocolEndpoint {
	if o == nil || o.Sso == nil {
		var ret ProtocolEndpoint
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoints) GetSsoOk() (*ProtocolEndpoint, bool) {
	if o == nil || o.Sso == nil {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *ProtocolEndpoints) HasSso() bool {
	if o != nil && o.Sso != nil {
		return true
	}

	return false
}

// SetSso gets a reference to the given ProtocolEndpoint and assigns it to the Sso field.
func (o *ProtocolEndpoints) SetSso(v ProtocolEndpoint) {
	o.Sso = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProtocolEndpoints) GetToken() ProtocolEndpoint {
	if o == nil || o.Token == nil {
		var ret ProtocolEndpoint
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoints) GetTokenOk() (*ProtocolEndpoint, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProtocolEndpoints) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given ProtocolEndpoint and assigns it to the Token field.
func (o *ProtocolEndpoints) SetToken(v ProtocolEndpoint) {
	o.Token = &v
}

// GetUserInfo returns the UserInfo field value if set, zero value otherwise.
func (o *ProtocolEndpoints) GetUserInfo() ProtocolEndpoint {
	if o == nil || o.UserInfo == nil {
		var ret ProtocolEndpoint
		return ret
	}
	return *o.UserInfo
}

// GetUserInfoOk returns a tuple with the UserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolEndpoints) GetUserInfoOk() (*ProtocolEndpoint, bool) {
	if o == nil || o.UserInfo == nil {
		return nil, false
	}
	return o.UserInfo, true
}

// HasUserInfo returns a boolean if a field has been set.
func (o *ProtocolEndpoints) HasUserInfo() bool {
	if o != nil && o.UserInfo != nil {
		return true
	}

	return false
}

// SetUserInfo gets a reference to the given ProtocolEndpoint and assigns it to the UserInfo field.
func (o *ProtocolEndpoints) SetUserInfo(v ProtocolEndpoint) {
	o.UserInfo = &v
}

func (o ProtocolEndpoints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Acs != nil {
		toSerialize["acs"] = o.Acs
	}
	if o.Authorization != nil {
		toSerialize["authorization"] = o.Authorization
	}
	if o.Jwks != nil {
		toSerialize["jwks"] = o.Jwks
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Slo != nil {
		toSerialize["slo"] = o.Slo
	}
	if o.Sso != nil {
		toSerialize["sso"] = o.Sso
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UserInfo != nil {
		toSerialize["userInfo"] = o.UserInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProtocolEndpoints) UnmarshalJSON(bytes []byte) (err error) {
	varProtocolEndpoints := _ProtocolEndpoints{}

	err = json.Unmarshal(bytes, &varProtocolEndpoints)
	if err == nil {
		*o = ProtocolEndpoints(varProtocolEndpoints)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "acs")
		delete(additionalProperties, "authorization")
		delete(additionalProperties, "jwks")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "slo")
		delete(additionalProperties, "sso")
		delete(additionalProperties, "token")
		delete(additionalProperties, "userInfo")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProtocolEndpoints struct {
	value *ProtocolEndpoints
	isSet bool
}

func (v NullableProtocolEndpoints) Get() *ProtocolEndpoints {
	return v.value
}

func (v *NullableProtocolEndpoints) Set(val *ProtocolEndpoints) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolEndpoints) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolEndpoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolEndpoints(val *ProtocolEndpoints) *NullableProtocolEndpoints {
	return &NullableProtocolEndpoints{value: val, isSet: true}
}

func (v NullableProtocolEndpoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolEndpoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

