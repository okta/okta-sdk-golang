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

// checks if the LinksAerialConsentRevoked type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinksAerialConsentRevoked{}

// LinksAerialConsentRevoked struct for LinksAerialConsentRevoked
type LinksAerialConsentRevoked struct {
	Grant                *HrefObjectGrantAerialConsent `json:"grant,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksAerialConsentRevoked LinksAerialConsentRevoked

// NewLinksAerialConsentRevoked instantiates a new LinksAerialConsentRevoked object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksAerialConsentRevoked() *LinksAerialConsentRevoked {
	this := LinksAerialConsentRevoked{}
	return &this
}

// NewLinksAerialConsentRevokedWithDefaults instantiates a new LinksAerialConsentRevoked object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksAerialConsentRevokedWithDefaults() *LinksAerialConsentRevoked {
	this := LinksAerialConsentRevoked{}
	return &this
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *LinksAerialConsentRevoked) GetGrant() HrefObjectGrantAerialConsent {
	if o == nil || IsNil(o.Grant) {
		var ret HrefObjectGrantAerialConsent
		return ret
	}
	return *o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksAerialConsentRevoked) GetGrantOk() (*HrefObjectGrantAerialConsent, bool) {
	if o == nil || IsNil(o.Grant) {
		return nil, false
	}
	return o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *LinksAerialConsentRevoked) HasGrant() bool {
	if o != nil && !IsNil(o.Grant) {
		return true
	}

	return false
}

// SetGrant gets a reference to the given HrefObjectGrantAerialConsent and assigns it to the Grant field.
func (o *LinksAerialConsentRevoked) SetGrant(v HrefObjectGrantAerialConsent) {
	o.Grant = &v
}

func (o LinksAerialConsentRevoked) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinksAerialConsentRevoked) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Grant) {
		toSerialize["grant"] = o.Grant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinksAerialConsentRevoked) UnmarshalJSON(data []byte) (err error) {
	varLinksAerialConsentRevoked := _LinksAerialConsentRevoked{}

	err = json.Unmarshal(data, &varLinksAerialConsentRevoked)

	if err != nil {
		return err
	}

	*o = LinksAerialConsentRevoked(varLinksAerialConsentRevoked)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grant")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinksAerialConsentRevoked struct {
	value *LinksAerialConsentRevoked
	isSet bool
}

func (v NullableLinksAerialConsentRevoked) Get() *LinksAerialConsentRevoked {
	return v.value
}

func (v *NullableLinksAerialConsentRevoked) Set(val *LinksAerialConsentRevoked) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksAerialConsentRevoked) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksAerialConsentRevoked) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksAerialConsentRevoked(val *LinksAerialConsentRevoked) *NullableLinksAerialConsentRevoked {
	return &NullableLinksAerialConsentRevoked{value: val, isSet: true}
}

func (v NullableLinksAerialConsentRevoked) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksAerialConsentRevoked) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
