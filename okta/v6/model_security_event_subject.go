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

// SecurityEventSubject The event subject
type SecurityEventSubject struct {
	// The format of the subject
	Format *string `json:"format,omitempty"`
	// An identifier of the actor
	Iss *string `json:"iss,omitempty"`
	// An identifier for the subject that was acted on
	Sub *string `json:"sub,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventSubject SecurityEventSubject

// NewSecurityEventSubject instantiates a new SecurityEventSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventSubject() *SecurityEventSubject {
	this := SecurityEventSubject{}
	return &this
}

// NewSecurityEventSubjectWithDefaults instantiates a new SecurityEventSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventSubjectWithDefaults() *SecurityEventSubject {
	this := SecurityEventSubject{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *SecurityEventSubject) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventSubject) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *SecurityEventSubject) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *SecurityEventSubject) SetFormat(v string) {
	o.Format = &v
}

// GetIss returns the Iss field value if set, zero value otherwise.
func (o *SecurityEventSubject) GetIss() string {
	if o == nil || o.Iss == nil {
		var ret string
		return ret
	}
	return *o.Iss
}

// GetIssOk returns a tuple with the Iss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventSubject) GetIssOk() (*string, bool) {
	if o == nil || o.Iss == nil {
		return nil, false
	}
	return o.Iss, true
}

// HasIss returns a boolean if a field has been set.
func (o *SecurityEventSubject) HasIss() bool {
	if o != nil && o.Iss != nil {
		return true
	}

	return false
}

// SetIss gets a reference to the given string and assigns it to the Iss field.
func (o *SecurityEventSubject) SetIss(v string) {
	o.Iss = &v
}

// GetSub returns the Sub field value if set, zero value otherwise.
func (o *SecurityEventSubject) GetSub() string {
	if o == nil || o.Sub == nil {
		var ret string
		return ret
	}
	return *o.Sub
}

// GetSubOk returns a tuple with the Sub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventSubject) GetSubOk() (*string, bool) {
	if o == nil || o.Sub == nil {
		return nil, false
	}
	return o.Sub, true
}

// HasSub returns a boolean if a field has been set.
func (o *SecurityEventSubject) HasSub() bool {
	if o != nil && o.Sub != nil {
		return true
	}

	return false
}

// SetSub gets a reference to the given string and assigns it to the Sub field.
func (o *SecurityEventSubject) SetSub(v string) {
	o.Sub = &v
}

func (o SecurityEventSubject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Iss != nil {
		toSerialize["iss"] = o.Iss
	}
	if o.Sub != nil {
		toSerialize["sub"] = o.Sub
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventSubject) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventSubject := _SecurityEventSubject{}

	err = json.Unmarshal(bytes, &varSecurityEventSubject)
	if err == nil {
		*o = SecurityEventSubject(varSecurityEventSubject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "iss")
		delete(additionalProperties, "sub")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventSubject struct {
	value *SecurityEventSubject
	isSet bool
}

func (v NullableSecurityEventSubject) Get() *SecurityEventSubject {
	return v.value
}

func (v *NullableSecurityEventSubject) Set(val *SecurityEventSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventSubject(val *SecurityEventSubject) *NullableSecurityEventSubject {
	return &NullableSecurityEventSubject{value: val, isSet: true}
}

func (v NullableSecurityEventSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

