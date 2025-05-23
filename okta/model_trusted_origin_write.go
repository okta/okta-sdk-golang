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

// TrustedOriginWrite struct for TrustedOriginWrite
type TrustedOriginWrite struct {
	// Unique name for the Trusted Origin
	Name *string `json:"name,omitempty"`
	// Unique origin URL for the Trusted Origin. The supported schemes for this attribute are HTTP, HTTPS, FTP, Ionic 2, and Capacitor.
	Origin *string `json:"origin,omitempty"`
	// Array of Scope types that this Trusted Origin is used for
	Scopes []TrustedOriginScope `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TrustedOriginWrite TrustedOriginWrite

// NewTrustedOriginWrite instantiates a new TrustedOriginWrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustedOriginWrite() *TrustedOriginWrite {
	this := TrustedOriginWrite{}
	return &this
}

// NewTrustedOriginWriteWithDefaults instantiates a new TrustedOriginWrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustedOriginWriteWithDefaults() *TrustedOriginWrite {
	this := TrustedOriginWrite{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TrustedOriginWrite) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOriginWrite) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TrustedOriginWrite) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TrustedOriginWrite) SetName(v string) {
	o.Name = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *TrustedOriginWrite) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOriginWrite) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *TrustedOriginWrite) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *TrustedOriginWrite) SetOrigin(v string) {
	o.Origin = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TrustedOriginWrite) GetScopes() []TrustedOriginScope {
	if o == nil || o.Scopes == nil {
		var ret []TrustedOriginScope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOriginWrite) GetScopesOk() ([]TrustedOriginScope, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TrustedOriginWrite) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []TrustedOriginScope and assigns it to the Scopes field.
func (o *TrustedOriginWrite) SetScopes(v []TrustedOriginScope) {
	o.Scopes = v
}

func (o TrustedOriginWrite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TrustedOriginWrite) UnmarshalJSON(bytes []byte) (err error) {
	varTrustedOriginWrite := _TrustedOriginWrite{}

	err = json.Unmarshal(bytes, &varTrustedOriginWrite)
	if err == nil {
		*o = TrustedOriginWrite(varTrustedOriginWrite)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "origin")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTrustedOriginWrite struct {
	value *TrustedOriginWrite
	isSet bool
}

func (v NullableTrustedOriginWrite) Get() *TrustedOriginWrite {
	return v.value
}

func (v *NullableTrustedOriginWrite) Set(val *TrustedOriginWrite) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustedOriginWrite) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustedOriginWrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustedOriginWrite(val *TrustedOriginWrite) *NullableTrustedOriginWrite {
	return &NullableTrustedOriginWrite{value: val, isSet: true}
}

func (v NullableTrustedOriginWrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustedOriginWrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

