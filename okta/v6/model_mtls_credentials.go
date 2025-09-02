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

// MtlsCredentials Certificate chain description for verifying assertions from the Smart Card
type MtlsCredentials struct {
	Trust *MtlsTrustCredentials `json:"trust,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MtlsCredentials MtlsCredentials

// NewMtlsCredentials instantiates a new MtlsCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtlsCredentials() *MtlsCredentials {
	this := MtlsCredentials{}
	return &this
}

// NewMtlsCredentialsWithDefaults instantiates a new MtlsCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtlsCredentialsWithDefaults() *MtlsCredentials {
	this := MtlsCredentials{}
	return &this
}

// GetTrust returns the Trust field value if set, zero value otherwise.
func (o *MtlsCredentials) GetTrust() MtlsTrustCredentials {
	if o == nil || o.Trust == nil {
		var ret MtlsTrustCredentials
		return ret
	}
	return *o.Trust
}

// GetTrustOk returns a tuple with the Trust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtlsCredentials) GetTrustOk() (*MtlsTrustCredentials, bool) {
	if o == nil || o.Trust == nil {
		return nil, false
	}
	return o.Trust, true
}

// HasTrust returns a boolean if a field has been set.
func (o *MtlsCredentials) HasTrust() bool {
	if o != nil && o.Trust != nil {
		return true
	}

	return false
}

// SetTrust gets a reference to the given MtlsTrustCredentials and assigns it to the Trust field.
func (o *MtlsCredentials) SetTrust(v MtlsTrustCredentials) {
	o.Trust = &v
}

func (o MtlsCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Trust != nil {
		toSerialize["trust"] = o.Trust
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MtlsCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varMtlsCredentials := _MtlsCredentials{}

	err = json.Unmarshal(bytes, &varMtlsCredentials)
	if err == nil {
		*o = MtlsCredentials(varMtlsCredentials)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "trust")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableMtlsCredentials struct {
	value *MtlsCredentials
	isSet bool
}

func (v NullableMtlsCredentials) Get() *MtlsCredentials {
	return v.value
}

func (v *NullableMtlsCredentials) Set(val *MtlsCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableMtlsCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableMtlsCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtlsCredentials(val *MtlsCredentials) *NullableMtlsCredentials {
	return &NullableMtlsCredentials{value: val, isSet: true}
}

func (v NullableMtlsCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtlsCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

