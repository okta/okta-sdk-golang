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

// SamlEndpoints SAML 2.0 HTTP binding settings for IdP and SP (Okta)
type SamlEndpoints struct {
	Acs *SamlAcsEndpoint `json:"acs,omitempty"`
	Slo *SamlSloEndpoint `json:"slo,omitempty"`
	Sso *SamlSsoEndpoint `json:"sso,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlEndpoints SamlEndpoints

// NewSamlEndpoints instantiates a new SamlEndpoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlEndpoints() *SamlEndpoints {
	this := SamlEndpoints{}
	return &this
}

// NewSamlEndpointsWithDefaults instantiates a new SamlEndpoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlEndpointsWithDefaults() *SamlEndpoints {
	this := SamlEndpoints{}
	return &this
}

// GetAcs returns the Acs field value if set, zero value otherwise.
func (o *SamlEndpoints) GetAcs() SamlAcsEndpoint {
	if o == nil || o.Acs == nil {
		var ret SamlAcsEndpoint
		return ret
	}
	return *o.Acs
}

// GetAcsOk returns a tuple with the Acs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlEndpoints) GetAcsOk() (*SamlAcsEndpoint, bool) {
	if o == nil || o.Acs == nil {
		return nil, false
	}
	return o.Acs, true
}

// HasAcs returns a boolean if a field has been set.
func (o *SamlEndpoints) HasAcs() bool {
	if o != nil && o.Acs != nil {
		return true
	}

	return false
}

// SetAcs gets a reference to the given SamlAcsEndpoint and assigns it to the Acs field.
func (o *SamlEndpoints) SetAcs(v SamlAcsEndpoint) {
	o.Acs = &v
}

// GetSlo returns the Slo field value if set, zero value otherwise.
func (o *SamlEndpoints) GetSlo() SamlSloEndpoint {
	if o == nil || o.Slo == nil {
		var ret SamlSloEndpoint
		return ret
	}
	return *o.Slo
}

// GetSloOk returns a tuple with the Slo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlEndpoints) GetSloOk() (*SamlSloEndpoint, bool) {
	if o == nil || o.Slo == nil {
		return nil, false
	}
	return o.Slo, true
}

// HasSlo returns a boolean if a field has been set.
func (o *SamlEndpoints) HasSlo() bool {
	if o != nil && o.Slo != nil {
		return true
	}

	return false
}

// SetSlo gets a reference to the given SamlSloEndpoint and assigns it to the Slo field.
func (o *SamlEndpoints) SetSlo(v SamlSloEndpoint) {
	o.Slo = &v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *SamlEndpoints) GetSso() SamlSsoEndpoint {
	if o == nil || o.Sso == nil {
		var ret SamlSsoEndpoint
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlEndpoints) GetSsoOk() (*SamlSsoEndpoint, bool) {
	if o == nil || o.Sso == nil {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *SamlEndpoints) HasSso() bool {
	if o != nil && o.Sso != nil {
		return true
	}

	return false
}

// SetSso gets a reference to the given SamlSsoEndpoint and assigns it to the Sso field.
func (o *SamlEndpoints) SetSso(v SamlSsoEndpoint) {
	o.Sso = &v
}

func (o SamlEndpoints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Acs != nil {
		toSerialize["acs"] = o.Acs
	}
	if o.Slo != nil {
		toSerialize["slo"] = o.Slo
	}
	if o.Sso != nil {
		toSerialize["sso"] = o.Sso
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlEndpoints) UnmarshalJSON(bytes []byte) (err error) {
	varSamlEndpoints := _SamlEndpoints{}

	err = json.Unmarshal(bytes, &varSamlEndpoints)
	if err == nil {
		*o = SamlEndpoints(varSamlEndpoints)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "acs")
		delete(additionalProperties, "slo")
		delete(additionalProperties, "sso")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlEndpoints struct {
	value *SamlEndpoints
	isSet bool
}

func (v NullableSamlEndpoints) Get() *SamlEndpoints {
	return v.value
}

func (v *NullableSamlEndpoints) Set(val *SamlEndpoints) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlEndpoints) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlEndpoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlEndpoints(val *SamlEndpoints) *NullableSamlEndpoints {
	return &NullableSamlEndpoints{value: val, isSet: true}
}

func (v NullableSamlEndpoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlEndpoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

