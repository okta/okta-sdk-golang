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

// DNSRecord DNS TXT and CNAME records to be registered for the Domain
type DNSRecord struct {
	// DNS TXT record expiration
	Expiration *string `json:"expiration,omitempty"`
	// DNS record name
	Fqdn *string `json:"fqdn,omitempty"`
	RecordType *string `json:"recordType,omitempty"`
	// DNS record value
	Values []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DNSRecord DNSRecord

// NewDNSRecord instantiates a new DNSRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSRecord() *DNSRecord {
	this := DNSRecord{}
	return &this
}

// NewDNSRecordWithDefaults instantiates a new DNSRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSRecordWithDefaults() *DNSRecord {
	this := DNSRecord{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *DNSRecord) GetExpiration() string {
	if o == nil || o.Expiration == nil {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecord) GetExpirationOk() (*string, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *DNSRecord) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *DNSRecord) SetExpiration(v string) {
	o.Expiration = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *DNSRecord) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecord) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *DNSRecord) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *DNSRecord) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *DNSRecord) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecord) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *DNSRecord) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *DNSRecord) SetRecordType(v string) {
	o.RecordType = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *DNSRecord) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecord) GetValuesOk() ([]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *DNSRecord) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *DNSRecord) SetValues(v []string) {
	o.Values = v
}

func (o DNSRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if o.Fqdn != nil {
		toSerialize["fqdn"] = o.Fqdn
	}
	if o.RecordType != nil {
		toSerialize["recordType"] = o.RecordType
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DNSRecord) UnmarshalJSON(bytes []byte) (err error) {
	varDNSRecord := _DNSRecord{}

	err = json.Unmarshal(bytes, &varDNSRecord)
	if err == nil {
		*o = DNSRecord(varDNSRecord)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "recordType")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDNSRecord struct {
	value *DNSRecord
	isSet bool
}

func (v NullableDNSRecord) Get() *DNSRecord {
	return v.value
}

func (v *NullableDNSRecord) Set(val *DNSRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSRecord(val *DNSRecord) *NullableDNSRecord {
	return &NullableDNSRecord{value: val, isSet: true}
}

func (v NullableDNSRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

