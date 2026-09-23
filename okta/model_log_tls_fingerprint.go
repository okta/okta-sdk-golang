/*
Okta Admin Management API

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

// checks if the LogTlsFingerprint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogTlsFingerprint{}

// LogTlsFingerprint The TLS client fingerprint associated with the event's request
type LogTlsFingerprint struct {
	// The JA4 TLS client fingerprint hash associated with the event's request
	Ja4                  NullableString `json:"ja4,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogTlsFingerprint LogTlsFingerprint

// NewLogTlsFingerprint instantiates a new LogTlsFingerprint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogTlsFingerprint() *LogTlsFingerprint {
	this := LogTlsFingerprint{}
	return &this
}

// NewLogTlsFingerprintWithDefaults instantiates a new LogTlsFingerprint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogTlsFingerprintWithDefaults() *LogTlsFingerprint {
	this := LogTlsFingerprint{}
	return &this
}

// GetJa4 returns the Ja4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogTlsFingerprint) GetJa4() string {
	if o == nil || IsNil(o.Ja4.Get()) {
		var ret string
		return ret
	}
	return *o.Ja4.Get()
}

// GetJa4Ok returns a tuple with the Ja4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogTlsFingerprint) GetJa4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ja4.Get(), o.Ja4.IsSet()
}

// HasJa4 returns a boolean if a field has been set.
func (o *LogTlsFingerprint) HasJa4() bool {
	if o != nil && o.Ja4.IsSet() {
		return true
	}

	return false
}

// SetJa4 gets a reference to the given NullableString and assigns it to the Ja4 field.
func (o *LogTlsFingerprint) SetJa4(v string) {
	o.Ja4.Set(&v)
}

// SetJa4Nil sets the value for Ja4 to be an explicit nil
func (o *LogTlsFingerprint) SetJa4Nil() {
	o.Ja4.Set(nil)
}

// UnsetJa4 ensures that no value is present for Ja4, not even an explicit nil
func (o *LogTlsFingerprint) UnsetJa4() {
	o.Ja4.Unset()
}

func (o LogTlsFingerprint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogTlsFingerprint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ja4.IsSet() {
		toSerialize["ja4"] = o.Ja4.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogTlsFingerprint) UnmarshalJSON(data []byte) (err error) {
	varLogTlsFingerprint := _LogTlsFingerprint{}

	err = json.Unmarshal(data, &varLogTlsFingerprint)

	if err != nil {
		return err
	}

	*o = LogTlsFingerprint(varLogTlsFingerprint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ja4")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogTlsFingerprint struct {
	value *LogTlsFingerprint
	isSet bool
}

func (v NullableLogTlsFingerprint) Get() *LogTlsFingerprint {
	return v.value
}

func (v *NullableLogTlsFingerprint) Set(val *LogTlsFingerprint) {
	v.value = val
	v.isSet = true
}

func (v NullableLogTlsFingerprint) IsSet() bool {
	return v.isSet
}

func (v *NullableLogTlsFingerprint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogTlsFingerprint(val *LogTlsFingerprint) *NullableLogTlsFingerprint {
	return &NullableLogTlsFingerprint{value: val, isSet: true}
}

func (v NullableLogTlsFingerprint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogTlsFingerprint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
