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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the MtlsEndpoints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtlsEndpoints{}

// MtlsEndpoints struct for MtlsEndpoints
type MtlsEndpoints struct {
	Sso                  *MtlsSsoEndpoint `json:"sso,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MtlsEndpoints MtlsEndpoints

// NewMtlsEndpoints instantiates a new MtlsEndpoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtlsEndpoints() *MtlsEndpoints {
	this := MtlsEndpoints{}
	return &this
}

// NewMtlsEndpointsWithDefaults instantiates a new MtlsEndpoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtlsEndpointsWithDefaults() *MtlsEndpoints {
	this := MtlsEndpoints{}
	return &this
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *MtlsEndpoints) GetSso() MtlsSsoEndpoint {
	if o == nil || IsNil(o.Sso) {
		var ret MtlsSsoEndpoint
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtlsEndpoints) GetSsoOk() (*MtlsSsoEndpoint, bool) {
	if o == nil || IsNil(o.Sso) {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *MtlsEndpoints) HasSso() bool {
	if o != nil && !IsNil(o.Sso) {
		return true
	}

	return false
}

// SetSso gets a reference to the given MtlsSsoEndpoint and assigns it to the Sso field.
func (o *MtlsEndpoints) SetSso(v MtlsSsoEndpoint) {
	o.Sso = &v
}

func (o MtlsEndpoints) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtlsEndpoints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sso) {
		toSerialize["sso"] = o.Sso
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MtlsEndpoints) UnmarshalJSON(data []byte) (err error) {
	varMtlsEndpoints := _MtlsEndpoints{}

	err = json.Unmarshal(data, &varMtlsEndpoints)

	if err != nil {
		return err
	}

	*o = MtlsEndpoints(varMtlsEndpoints)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sso")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMtlsEndpoints struct {
	value *MtlsEndpoints
	isSet bool
}

func (v NullableMtlsEndpoints) Get() *MtlsEndpoints {
	return v.value
}

func (v *NullableMtlsEndpoints) Set(val *MtlsEndpoints) {
	v.value = val
	v.isSet = true
}

func (v NullableMtlsEndpoints) IsSet() bool {
	return v.isSet
}

func (v *NullableMtlsEndpoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtlsEndpoints(val *MtlsEndpoints) *NullableMtlsEndpoints {
	return &NullableMtlsEndpoints{value: val, isSet: true}
}

func (v NullableMtlsEndpoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtlsEndpoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
