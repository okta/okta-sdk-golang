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
)

// checks if the WebAuthnRpId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebAuthnRpId{}

// WebAuthnRpId <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>The [RP ID](https://www.w3.org/TR/webauthn/#relying-party-identifier) object for WebAuthn configuration
type WebAuthnRpId struct {
	Domain *WebAuthnRpIdDomain `json:"domain,omitempty"`
	// Indicates whether the RP ID is active and is used for WebAuthn operations. It can only be set to `true` once the `validationStatus` of the `domain` object is `VERIFIED`. `enabled` can only be `true` for this same `domain`. Its value must be `false` to be able to configure the `domain`.
	Enabled              *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebAuthnRpId WebAuthnRpId

// NewWebAuthnRpId instantiates a new WebAuthnRpId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnRpId() *WebAuthnRpId {
	this := WebAuthnRpId{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewWebAuthnRpIdWithDefaults instantiates a new WebAuthnRpId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnRpIdWithDefaults() *WebAuthnRpId {
	this := WebAuthnRpId{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *WebAuthnRpId) GetDomain() WebAuthnRpIdDomain {
	if o == nil || IsNil(o.Domain) {
		var ret WebAuthnRpIdDomain
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnRpId) GetDomainOk() (*WebAuthnRpIdDomain, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *WebAuthnRpId) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given WebAuthnRpIdDomain and assigns it to the Domain field.
func (o *WebAuthnRpId) SetDomain(v WebAuthnRpIdDomain) {
	o.Domain = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *WebAuthnRpId) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnRpId) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *WebAuthnRpId) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *WebAuthnRpId) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o WebAuthnRpId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebAuthnRpId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebAuthnRpId) UnmarshalJSON(data []byte) (err error) {
	varWebAuthnRpId := _WebAuthnRpId{}

	err = json.Unmarshal(data, &varWebAuthnRpId)

	if err != nil {
		return err
	}

	*o = WebAuthnRpId(varWebAuthnRpId)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebAuthnRpId struct {
	value *WebAuthnRpId
	isSet bool
}

func (v NullableWebAuthnRpId) Get() *WebAuthnRpId {
	return v.value
}

func (v *NullableWebAuthnRpId) Set(val *WebAuthnRpId) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnRpId) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnRpId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnRpId(val *WebAuthnRpId) *NullableWebAuthnRpId {
	return &NullableWebAuthnRpId{value: val, isSet: true}
}

func (v NullableWebAuthnRpId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnRpId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
