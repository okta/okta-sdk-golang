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

// checks if the OktaSupportCase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaSupportCase{}

// OktaSupportCase struct for OktaSupportCase
type OktaSupportCase struct {
	// Okta Support case number
	CaseNumber    *string                       `json:"caseNumber,omitempty"`
	Impersonation *OktaSupportCaseImpersonation `json:"impersonation,omitempty"`
	SelfAssigned  *OktaSupportCaseSelfAssigned  `json:"selfAssigned,omitempty"`
	// Subject of the support case
	Subject              *string `json:"subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSupportCase OktaSupportCase

// NewOktaSupportCase instantiates a new OktaSupportCase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSupportCase() *OktaSupportCase {
	this := OktaSupportCase{}
	return &this
}

// NewOktaSupportCaseWithDefaults instantiates a new OktaSupportCase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSupportCaseWithDefaults() *OktaSupportCase {
	this := OktaSupportCase{}
	return &this
}

// GetCaseNumber returns the CaseNumber field value if set, zero value otherwise.
func (o *OktaSupportCase) GetCaseNumber() string {
	if o == nil || IsNil(o.CaseNumber) {
		var ret string
		return ret
	}
	return *o.CaseNumber
}

// GetCaseNumberOk returns a tuple with the CaseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSupportCase) GetCaseNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CaseNumber) {
		return nil, false
	}
	return o.CaseNumber, true
}

// HasCaseNumber returns a boolean if a field has been set.
func (o *OktaSupportCase) HasCaseNumber() bool {
	if o != nil && !IsNil(o.CaseNumber) {
		return true
	}

	return false
}

// SetCaseNumber gets a reference to the given string and assigns it to the CaseNumber field.
func (o *OktaSupportCase) SetCaseNumber(v string) {
	o.CaseNumber = &v
}

// GetImpersonation returns the Impersonation field value if set, zero value otherwise.
func (o *OktaSupportCase) GetImpersonation() OktaSupportCaseImpersonation {
	if o == nil || IsNil(o.Impersonation) {
		var ret OktaSupportCaseImpersonation
		return ret
	}
	return *o.Impersonation
}

// GetImpersonationOk returns a tuple with the Impersonation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSupportCase) GetImpersonationOk() (*OktaSupportCaseImpersonation, bool) {
	if o == nil || IsNil(o.Impersonation) {
		return nil, false
	}
	return o.Impersonation, true
}

// HasImpersonation returns a boolean if a field has been set.
func (o *OktaSupportCase) HasImpersonation() bool {
	if o != nil && !IsNil(o.Impersonation) {
		return true
	}

	return false
}

// SetImpersonation gets a reference to the given OktaSupportCaseImpersonation and assigns it to the Impersonation field.
func (o *OktaSupportCase) SetImpersonation(v OktaSupportCaseImpersonation) {
	o.Impersonation = &v
}

// GetSelfAssigned returns the SelfAssigned field value if set, zero value otherwise.
func (o *OktaSupportCase) GetSelfAssigned() OktaSupportCaseSelfAssigned {
	if o == nil || IsNil(o.SelfAssigned) {
		var ret OktaSupportCaseSelfAssigned
		return ret
	}
	return *o.SelfAssigned
}

// GetSelfAssignedOk returns a tuple with the SelfAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSupportCase) GetSelfAssignedOk() (*OktaSupportCaseSelfAssigned, bool) {
	if o == nil || IsNil(o.SelfAssigned) {
		return nil, false
	}
	return o.SelfAssigned, true
}

// HasSelfAssigned returns a boolean if a field has been set.
func (o *OktaSupportCase) HasSelfAssigned() bool {
	if o != nil && !IsNil(o.SelfAssigned) {
		return true
	}

	return false
}

// SetSelfAssigned gets a reference to the given OktaSupportCaseSelfAssigned and assigns it to the SelfAssigned field.
func (o *OktaSupportCase) SetSelfAssigned(v OktaSupportCaseSelfAssigned) {
	o.SelfAssigned = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *OktaSupportCase) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSupportCase) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *OktaSupportCase) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *OktaSupportCase) SetSubject(v string) {
	o.Subject = &v
}

func (o OktaSupportCase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaSupportCase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaseNumber) {
		toSerialize["caseNumber"] = o.CaseNumber
	}
	if !IsNil(o.Impersonation) {
		toSerialize["impersonation"] = o.Impersonation
	}
	if !IsNil(o.SelfAssigned) {
		toSerialize["selfAssigned"] = o.SelfAssigned
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaSupportCase) UnmarshalJSON(data []byte) (err error) {
	varOktaSupportCase := _OktaSupportCase{}

	err = json.Unmarshal(data, &varOktaSupportCase)

	if err != nil {
		return err
	}

	*o = OktaSupportCase(varOktaSupportCase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "caseNumber")
		delete(additionalProperties, "impersonation")
		delete(additionalProperties, "selfAssigned")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaSupportCase struct {
	value *OktaSupportCase
	isSet bool
}

func (v NullableOktaSupportCase) Get() *OktaSupportCase {
	return v.value
}

func (v *NullableOktaSupportCase) Set(val *OktaSupportCase) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSupportCase) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSupportCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSupportCase(val *OktaSupportCase) *NullableOktaSupportCase {
	return &NullableOktaSupportCase{value: val, isSet: true}
}

func (v NullableOktaSupportCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSupportCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
