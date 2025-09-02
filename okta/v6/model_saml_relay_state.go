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

// SamlRelayState Relay state settings for IdP
type SamlRelayState struct {
	// The format used to generate the `relayState` in the SAML request. The `FROM_URL` format is used if this value is null.
	Format *string `json:"format,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlRelayState SamlRelayState

// NewSamlRelayState instantiates a new SamlRelayState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlRelayState() *SamlRelayState {
	this := SamlRelayState{}
	return &this
}

// NewSamlRelayStateWithDefaults instantiates a new SamlRelayState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlRelayStateWithDefaults() *SamlRelayState {
	this := SamlRelayState{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *SamlRelayState) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlRelayState) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *SamlRelayState) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *SamlRelayState) SetFormat(v string) {
	o.Format = &v
}

func (o SamlRelayState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SamlRelayState) UnmarshalJSON(bytes []byte) (err error) {
	varSamlRelayState := _SamlRelayState{}

	err = json.Unmarshal(bytes, &varSamlRelayState)
	if err == nil {
		*o = SamlRelayState(varSamlRelayState)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "format")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSamlRelayState struct {
	value *SamlRelayState
	isSet bool
}

func (v NullableSamlRelayState) Get() *SamlRelayState {
	return v.value
}

func (v *NullableSamlRelayState) Set(val *SamlRelayState) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlRelayState) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlRelayState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlRelayState(val *SamlRelayState) *NullableSamlRelayState {
	return &NullableSamlRelayState{value: val, isSet: true}
}

func (v NullableSamlRelayState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlRelayState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

