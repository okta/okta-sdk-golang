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

// checks if the SsfTransmitterSecurityEventSubject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SsfTransmitterSecurityEventSubject{}

// SsfTransmitterSecurityEventSubject The event subject
type SsfTransmitterSecurityEventSubject struct {
	// The format of the subject
	Format *string `json:"format,omitempty"`
	// An identifier of the actor
	Iss *string `json:"iss,omitempty"`
	// An identifier for the subject that was acted on
	Sub                  *string `json:"sub,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SsfTransmitterSecurityEventSubject SsfTransmitterSecurityEventSubject

// NewSsfTransmitterSecurityEventSubject instantiates a new SsfTransmitterSecurityEventSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsfTransmitterSecurityEventSubject() *SsfTransmitterSecurityEventSubject {
	this := SsfTransmitterSecurityEventSubject{}
	return &this
}

// NewSsfTransmitterSecurityEventSubjectWithDefaults instantiates a new SsfTransmitterSecurityEventSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsfTransmitterSecurityEventSubjectWithDefaults() *SsfTransmitterSecurityEventSubject {
	this := SsfTransmitterSecurityEventSubject{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *SsfTransmitterSecurityEventSubject) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsfTransmitterSecurityEventSubject) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *SsfTransmitterSecurityEventSubject) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *SsfTransmitterSecurityEventSubject) SetFormat(v string) {
	o.Format = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *SsfTransmitterSecurityEventSubject) GetIss() string {
	if o == nil || IsNil(o.Iss) {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsfTransmitterSecurityEventSubject) GetIssOk() (*string, bool) {
	if o == nil || IsNil(o.Iss) {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *SsfTransmitterSecurityEventSubject) HasIss() bool {
	if o != nil && !IsNil(o.Iss) {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *SsfTransmitterSecurityEventSubject) SetIss(v string) {
	o.Iss = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *SsfTransmitterSecurityEventSubject) GetSub() string {
	if o == nil || IsNil(o.Sub) {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsfTransmitterSecurityEventSubject) GetSubOk() (*string, bool) {
	if o == nil || IsNil(o.Sub) {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *SsfTransmitterSecurityEventSubject) HasSub() bool {
	if o != nil && !IsNil(o.Sub) {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *SsfTransmitterSecurityEventSubject) SetSub(v string) {
	o.Sub = &v
}

func (o SsfTransmitterSecurityEventSubject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SsfTransmitterSecurityEventSubject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Iss) {
		toSerialize["iss"] = o.Iss
	}
	if !IsNil(o.Sub) {
		toSerialize["sub"] = o.Sub
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SsfTransmitterSecurityEventSubject) UnmarshalJSON(data []byte) (err error) {
	varSsfTransmitterSecurityEventSubject := _SsfTransmitterSecurityEventSubject{}

	err = json.Unmarshal(data, &varSsfTransmitterSecurityEventSubject)

	if err != nil {
		return err
	}

	*o = SsfTransmitterSecurityEventSubject(varSsfTransmitterSecurityEventSubject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "iss")
		delete(additionalProperties, "sub")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSsfTransmitterSecurityEventSubject struct {
	value *SsfTransmitterSecurityEventSubject
	isSet bool
}

func (v NullableSsfTransmitterSecurityEventSubject) Get() *SsfTransmitterSecurityEventSubject {
	return v.value
}

func (v *NullableSsfTransmitterSecurityEventSubject) Set(val *SsfTransmitterSecurityEventSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableSsfTransmitterSecurityEventSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableSsfTransmitterSecurityEventSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsfTransmitterSecurityEventSubject(val *SsfTransmitterSecurityEventSubject) *NullableSsfTransmitterSecurityEventSubject {
	return &NullableSsfTransmitterSecurityEventSubject{value: val, isSet: true}
}

func (v NullableSsfTransmitterSecurityEventSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsfTransmitterSecurityEventSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
