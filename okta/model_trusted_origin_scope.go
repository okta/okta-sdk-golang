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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// TrustedOriginScope struct for TrustedOriginScope
type TrustedOriginScope struct {
	AllowedOktaApps      []IframeEmbedScopeAllowedApps `json:"allowedOktaApps,omitempty"`
	Type                 *TrustedOriginScopeType       `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TrustedOriginScope TrustedOriginScope

// NewTrustedOriginScope instantiates a new TrustedOriginScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustedOriginScope() *TrustedOriginScope {
	this := TrustedOriginScope{}
	return &this
}

// NewTrustedOriginScopeWithDefaults instantiates a new TrustedOriginScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustedOriginScopeWithDefaults() *TrustedOriginScope {
	this := TrustedOriginScope{}
	return &this
}

// GetAllowedOktaApps returns the AllowedOktaApps field value if set, zero value otherwise.
func (o *TrustedOriginScope) GetAllowedOktaApps() []IframeEmbedScopeAllowedApps {
	if o == nil || o.AllowedOktaApps == nil {
		var ret []IframeEmbedScopeAllowedApps
		return ret
	}
	return o.AllowedOktaApps
}

// GetAllowedOktaAppsOk returns a tuple with the AllowedOktaApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOriginScope) GetAllowedOktaAppsOk() ([]IframeEmbedScopeAllowedApps, bool) {
	if o == nil || o.AllowedOktaApps == nil {
		return nil, false
	}
	return o.AllowedOktaApps, true
}

// HasAllowedOktaApps returns a boolean if a field has been set.
func (o *TrustedOriginScope) HasAllowedOktaApps() bool {
	if o != nil && o.AllowedOktaApps != nil {
		return true
	}

	return false
}

// SetAllowedOktaApps gets a reference to the given []IframeEmbedScopeAllowedApps and assigns it to the AllowedOktaApps field.
func (o *TrustedOriginScope) SetAllowedOktaApps(v []IframeEmbedScopeAllowedApps) {
	o.AllowedOktaApps = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TrustedOriginScope) GetType() TrustedOriginScopeType {
	if o == nil || o.Type == nil {
		var ret TrustedOriginScopeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOriginScope) GetTypeOk() (*TrustedOriginScopeType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TrustedOriginScope) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TrustedOriginScopeType and assigns it to the Type field.
func (o *TrustedOriginScope) SetType(v TrustedOriginScopeType) {
	o.Type = &v
}

func (o TrustedOriginScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedOktaApps != nil {
		toSerialize["allowedOktaApps"] = o.AllowedOktaApps
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TrustedOriginScope) UnmarshalJSON(bytes []byte) (err error) {
	varTrustedOriginScope := _TrustedOriginScope{}

	err = json.Unmarshal(bytes, &varTrustedOriginScope)
	if err == nil {
		*o = TrustedOriginScope(varTrustedOriginScope)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allowedOktaApps")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTrustedOriginScope struct {
	value *TrustedOriginScope
	isSet bool
}

func (v NullableTrustedOriginScope) Get() *TrustedOriginScope {
	return v.value
}

func (v *NullableTrustedOriginScope) Set(val *TrustedOriginScope) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustedOriginScope) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustedOriginScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustedOriginScope(val *TrustedOriginScope) *NullableTrustedOriginScope {
	return &NullableTrustedOriginScope{value: val, isSet: true}
}

func (v NullableTrustedOriginScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustedOriginScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
