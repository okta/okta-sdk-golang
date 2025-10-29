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

// checks if the DNSRecordAuthenticators type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSRecordAuthenticators{}

// DNSRecordAuthenticators DNS TXT record that must be registered for an RP ID domain that requires verification. This is used to verify the domain ownership for the WebAuthn RP ID configuration. After the domain ownership is verified, the `DNSRecord` isn't returned in the response.
type DNSRecordAuthenticators struct {
	// The DNS record name
	Fqdn       *string `json:"fqdn,omitempty"`
	RecordType *string `json:"recordType,omitempty"`
	// The DNS record verification value
	VerificationValue    *string `json:"verificationValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DNSRecordAuthenticators DNSRecordAuthenticators

// NewDNSRecordAuthenticators instantiates a new DNSRecordAuthenticators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSRecordAuthenticators() *DNSRecordAuthenticators {
	this := DNSRecordAuthenticators{}
	return &this
}

// NewDNSRecordAuthenticatorsWithDefaults instantiates a new DNSRecordAuthenticators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSRecordAuthenticatorsWithDefaults() *DNSRecordAuthenticators {
	this := DNSRecordAuthenticators{}
	return &this
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *DNSRecordAuthenticators) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordAuthenticators) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *DNSRecordAuthenticators) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *DNSRecordAuthenticators) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *DNSRecordAuthenticators) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordAuthenticators) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *DNSRecordAuthenticators) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *DNSRecordAuthenticators) SetRecordType(v string) {
	o.RecordType = &v
}

// GetVerificationValue returns the VerificationValue field value if set, zero value otherwise.
func (o *DNSRecordAuthenticators) GetVerificationValue() string {
	if o == nil || IsNil(o.VerificationValue) {
		var ret string
		return ret
	}
	return *o.VerificationValue
}

// GetVerificationValueOk returns a tuple with the VerificationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordAuthenticators) GetVerificationValueOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationValue) {
		return nil, false
	}
	return o.VerificationValue, true
}

// HasVerificationValue returns a boolean if a field has been set.
func (o *DNSRecordAuthenticators) HasVerificationValue() bool {
	if o != nil && !IsNil(o.VerificationValue) {
		return true
	}

	return false
}

// SetVerificationValue gets a reference to the given string and assigns it to the VerificationValue field.
func (o *DNSRecordAuthenticators) SetVerificationValue(v string) {
	o.VerificationValue = &v
}

func (o DNSRecordAuthenticators) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSRecordAuthenticators) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.RecordType) {
		toSerialize["recordType"] = o.RecordType
	}
	if !IsNil(o.VerificationValue) {
		toSerialize["verificationValue"] = o.VerificationValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DNSRecordAuthenticators) UnmarshalJSON(data []byte) (err error) {
	varDNSRecordAuthenticators := _DNSRecordAuthenticators{}

	err = json.Unmarshal(data, &varDNSRecordAuthenticators)

	if err != nil {
		return err
	}

	*o = DNSRecordAuthenticators(varDNSRecordAuthenticators)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "recordType")
		delete(additionalProperties, "verificationValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDNSRecordAuthenticators struct {
	value *DNSRecordAuthenticators
	isSet bool
}

func (v NullableDNSRecordAuthenticators) Get() *DNSRecordAuthenticators {
	return v.value
}

func (v *NullableDNSRecordAuthenticators) Set(val *DNSRecordAuthenticators) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSRecordAuthenticators) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSRecordAuthenticators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSRecordAuthenticators(val *DNSRecordAuthenticators) *NullableDNSRecordAuthenticators {
	return &NullableDNSRecordAuthenticators{value: val, isSet: true}
}

func (v NullableDNSRecordAuthenticators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSRecordAuthenticators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
