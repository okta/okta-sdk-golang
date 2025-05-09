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

// EmailDomainDNSRecord struct for EmailDomainDNSRecord
type EmailDomainDNSRecord struct {
	Fqdn *string `json:"fqdn,omitempty"`
	RecordType *string `json:"recordType,omitempty"`
	VerificationValue *string `json:"verificationValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmailDomainDNSRecord EmailDomainDNSRecord

// NewEmailDomainDNSRecord instantiates a new EmailDomainDNSRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomainDNSRecord() *EmailDomainDNSRecord {
	this := EmailDomainDNSRecord{}
	return &this
}

// NewEmailDomainDNSRecordWithDefaults instantiates a new EmailDomainDNSRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainDNSRecordWithDefaults() *EmailDomainDNSRecord {
	this := EmailDomainDNSRecord{}
	return &this
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *EmailDomainDNSRecord) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainDNSRecord) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *EmailDomainDNSRecord) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *EmailDomainDNSRecord) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *EmailDomainDNSRecord) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainDNSRecord) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *EmailDomainDNSRecord) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *EmailDomainDNSRecord) SetRecordType(v string) {
	o.RecordType = &v
}

// GetVerificationValue returns the VerificationValue field value if set, zero value otherwise.
func (o *EmailDomainDNSRecord) GetVerificationValue() string {
	if o == nil || o.VerificationValue == nil {
		var ret string
		return ret
	}
	return *o.VerificationValue
}

// GetVerificationValueOk returns a tuple with the VerificationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainDNSRecord) GetVerificationValueOk() (*string, bool) {
	if o == nil || o.VerificationValue == nil {
		return nil, false
	}
	return o.VerificationValue, true
}

// HasVerificationValue returns a boolean if a field has been set.
func (o *EmailDomainDNSRecord) HasVerificationValue() bool {
	if o != nil && o.VerificationValue != nil {
		return true
	}

	return false
}

// SetVerificationValue gets a reference to the given string and assigns it to the VerificationValue field.
func (o *EmailDomainDNSRecord) SetVerificationValue(v string) {
	o.VerificationValue = &v
}

func (o EmailDomainDNSRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.RecordType != nil {
		toSerialize["recordType"] = o.RecordType
	}
	if o.VerificationValue != nil {
		toSerialize["verificationValue"] = o.VerificationValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailDomainDNSRecord) UnmarshalJSON(bytes []byte) (err error) {
	varEmailDomainDNSRecord := _EmailDomainDNSRecord{}

	err = json.Unmarshal(bytes, &varEmailDomainDNSRecord)
	if err == nil {
		*o = EmailDomainDNSRecord(varEmailDomainDNSRecord)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "recordType")
		delete(additionalProperties, "verificationValue")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailDomainDNSRecord struct {
	value *EmailDomainDNSRecord
	isSet bool
}

func (v NullableEmailDomainDNSRecord) Get() *EmailDomainDNSRecord {
	return v.value
}

func (v *NullableEmailDomainDNSRecord) Set(val *EmailDomainDNSRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomainDNSRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomainDNSRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomainDNSRecord(val *EmailDomainDNSRecord) *NullableEmailDomainDNSRecord {
	return &NullableEmailDomainDNSRecord{value: val, isSet: true}
}

func (v NullableEmailDomainDNSRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomainDNSRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

