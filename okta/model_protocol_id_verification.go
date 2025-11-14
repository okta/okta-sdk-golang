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
)

// checks if the ProtocolIdVerification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtocolIdVerification{}

// ProtocolIdVerification Protocol settings for the IDV vendor
type ProtocolIdVerification struct {
	Credentials *IDVCredentials `json:"credentials,omitempty"`
	Endpoints   *IDVEndpoints   `json:"endpoints,omitempty"`
	// IdP-defined permission bundles to request delegated access from the user. > **Note:** The [identity provider type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path=type&t=request) table lists the scopes that are supported for each IdP.
	Scopes []string `json:"scopes,omitempty"`
	// ID verification protocol
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtocolIdVerification ProtocolIdVerification

// NewProtocolIdVerification instantiates a new ProtocolIdVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocolIdVerification() *ProtocolIdVerification {
	this := ProtocolIdVerification{}
	return &this
}

// NewProtocolIdVerificationWithDefaults instantiates a new ProtocolIdVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolIdVerificationWithDefaults() *ProtocolIdVerification {
	this := ProtocolIdVerification{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ProtocolIdVerification) GetCredentials() IDVCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret IDVCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolIdVerification) GetCredentialsOk() (*IDVCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ProtocolIdVerification) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given IDVCredentials and assigns it to the Credentials field.
func (o *ProtocolIdVerification) SetCredentials(v IDVCredentials) {
	o.Credentials = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ProtocolIdVerification) GetEndpoints() IDVEndpoints {
	if o == nil || IsNil(o.Endpoints) {
		var ret IDVEndpoints
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolIdVerification) GetEndpointsOk() (*IDVEndpoints, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ProtocolIdVerification) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given IDVEndpoints and assigns it to the Endpoints field.
func (o *ProtocolIdVerification) SetEndpoints(v IDVEndpoints) {
	o.Endpoints = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ProtocolIdVerification) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolIdVerification) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ProtocolIdVerification) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ProtocolIdVerification) SetScopes(v []string) {
	o.Scopes = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProtocolIdVerification) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtocolIdVerification) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProtocolIdVerification) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProtocolIdVerification) SetType(v string) {
	o.Type = &v
}

func (o ProtocolIdVerification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtocolIdVerification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProtocolIdVerification) UnmarshalJSON(data []byte) (err error) {
	varProtocolIdVerification := _ProtocolIdVerification{}

	err = json.Unmarshal(data, &varProtocolIdVerification)

	if err != nil {
		return err
	}

	*o = ProtocolIdVerification(varProtocolIdVerification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "endpoints")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProtocolIdVerification struct {
	value *ProtocolIdVerification
	isSet bool
}

func (v NullableProtocolIdVerification) Get() *ProtocolIdVerification {
	return v.value
}

func (v *NullableProtocolIdVerification) Set(val *ProtocolIdVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocolIdVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocolIdVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocolIdVerification(val *ProtocolIdVerification) *NullableProtocolIdVerification {
	return &NullableProtocolIdVerification{value: val, isSet: true}
}

func (v NullableProtocolIdVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocolIdVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
