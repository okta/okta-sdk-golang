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

// checks if the DNSRecordDomains type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DNSRecordDomains{}

// DNSRecordDomains DNS TXT and CNAME records to be registered for the Domain
type DNSRecordDomains struct {
	// DNS TXT record expiration
	Expiration *string `json:"expiration,omitempty"`
	// DNS record name
	Fqdn       *string `json:"fqdn,omitempty"`
	RecordType *string `json:"recordType,omitempty"`
	// DNS record value
	Values               []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DNSRecordDomains DNSRecordDomains

// NewDNSRecordDomains instantiates a new DNSRecordDomains object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSRecordDomains() *DNSRecordDomains {
	this := DNSRecordDomains{}
	return &this
}

// NewDNSRecordDomainsWithDefaults instantiates a new DNSRecordDomains object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSRecordDomainsWithDefaults() *DNSRecordDomains {
	this := DNSRecordDomains{}
	return &this
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *DNSRecordDomains) GetExpiration() string {
	if o == nil || IsNil(o.Expiration) {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordDomains) GetExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *DNSRecordDomains) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *DNSRecordDomains) SetExpiration(v string) {
	o.Expiration = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *DNSRecordDomains) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordDomains) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *DNSRecordDomains) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *DNSRecordDomains) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *DNSRecordDomains) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordDomains) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *DNSRecordDomains) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *DNSRecordDomains) SetRecordType(v string) {
	o.RecordType = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *DNSRecordDomains) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSRecordDomains) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *DNSRecordDomains) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *DNSRecordDomains) SetValues(v []string) {
	o.Values = v
}

func (o DNSRecordDomains) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSRecordDomains) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.RecordType) {
		toSerialize["recordType"] = o.RecordType
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DNSRecordDomains) UnmarshalJSON(data []byte) (err error) {
	varDNSRecordDomains := _DNSRecordDomains{}

	err = json.Unmarshal(data, &varDNSRecordDomains)

	if err != nil {
		return err
	}

	*o = DNSRecordDomains(varDNSRecordDomains)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "recordType")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDNSRecordDomains struct {
	value *DNSRecordDomains
	isSet bool
}

func (v NullableDNSRecordDomains) Get() *DNSRecordDomains {
	return v.value
}

func (v *NullableDNSRecordDomains) Set(val *DNSRecordDomains) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSRecordDomains) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSRecordDomains) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSRecordDomains(val *DNSRecordDomains) *NullableDNSRecordDomains {
	return &NullableDNSRecordDomains{value: val, isSet: true}
}

func (v NullableDNSRecordDomains) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSRecordDomains) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
