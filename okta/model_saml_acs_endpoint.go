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

// checks if the SamlAcsEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlAcsEndpoint{}

// SamlAcsEndpoint Okta's `SPSSODescriptor` endpoint where the IdP sends a `<SAMLResponse>` message
type SamlAcsEndpoint struct {
	Binding *string `json:"binding,omitempty"`
	// Determines whether to publish an instance-specific (trust) or organization (shared) ACS endpoint in the SAML metadata
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlAcsEndpoint SamlAcsEndpoint

// NewSamlAcsEndpoint instantiates a new SamlAcsEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlAcsEndpoint() *SamlAcsEndpoint {
	this := SamlAcsEndpoint{}
	var type_ string = "INSTANCE"
	this.Type = &type_
	return &this
}

// NewSamlAcsEndpointWithDefaults instantiates a new SamlAcsEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAcsEndpointWithDefaults() *SamlAcsEndpoint {
	this := SamlAcsEndpoint{}
	var type_ string = "INSTANCE"
	this.Type = &type_
	return &this
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *SamlAcsEndpoint) GetBinding() string {
	if o == nil || IsNil(o.Binding) {
		var ret string
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAcsEndpoint) GetBindingOk() (*string, bool) {
	if o == nil || IsNil(o.Binding) {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *SamlAcsEndpoint) HasBinding() bool {
	if o != nil && !IsNil(o.Binding) {
		return true
	}

	return false
}

// SetBinding gets a reference to the given string and assigns it to the Binding field.
func (o *SamlAcsEndpoint) SetBinding(v string) {
	o.Binding = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SamlAcsEndpoint) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlAcsEndpoint) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SamlAcsEndpoint) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SamlAcsEndpoint) SetType(v string) {
	o.Type = &v
}

func (o SamlAcsEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlAcsEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Binding) {
		toSerialize["binding"] = o.Binding
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlAcsEndpoint) UnmarshalJSON(data []byte) (err error) {
	varSamlAcsEndpoint := _SamlAcsEndpoint{}

	err = json.Unmarshal(data, &varSamlAcsEndpoint)

	if err != nil {
		return err
	}

	*o = SamlAcsEndpoint(varSamlAcsEndpoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "binding")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlAcsEndpoint struct {
	value *SamlAcsEndpoint
	isSet bool
}

func (v NullableSamlAcsEndpoint) Get() *SamlAcsEndpoint {
	return v.value
}

func (v *NullableSamlAcsEndpoint) Set(val *SamlAcsEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlAcsEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlAcsEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlAcsEndpoint(val *SamlAcsEndpoint) *NullableSamlAcsEndpoint {
	return &NullableSamlAcsEndpoint{value: val, isSet: true}
}

func (v NullableSamlAcsEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlAcsEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
