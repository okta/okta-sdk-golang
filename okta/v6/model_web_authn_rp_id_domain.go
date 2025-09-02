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

// WebAuthnRpIdDomain The RP domain object for the WebAuthn configuration
type WebAuthnRpIdDomain struct {
	DnsRecord *DNSRecord `json:"dnsRecord,omitempty"`
	// The [RP ID](https://www.w3.org/TR/webauthn/#relying-party-identifier) domain value to be used for all WebAuthn operations.  If it isn't specified, the `domain` object isn't included in the request, and the domain value defaults to the domain of the current page (the domain of your org or a custom domain, for example).  > **Note:** If you don't use a custom RP ID (the default behavior), the domain value defaults to the end user's current page. The domain value defaults to the full domain name of the page that the end user is on when they're attempting the WebAuthn credential operation (enrollment or verification).
	Name *string `json:"name,omitempty"`
	// Indicates the validation status of the domain
	ValidationStatus *string `json:"validationStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebAuthnRpIdDomain WebAuthnRpIdDomain

// NewWebAuthnRpIdDomain instantiates a new WebAuthnRpIdDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnRpIdDomain() *WebAuthnRpIdDomain {
	this := WebAuthnRpIdDomain{}
	return &this
}

// NewWebAuthnRpIdDomainWithDefaults instantiates a new WebAuthnRpIdDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnRpIdDomainWithDefaults() *WebAuthnRpIdDomain {
	this := WebAuthnRpIdDomain{}
	return &this
}

// GetDnsRecord returns the DnsRecord field value if set, zero value otherwise.
func (o *WebAuthnRpIdDomain) GetDnsRecord() DNSRecord {
	if o == nil || o.DnsRecord == nil {
		var ret DNSRecord
		return ret
	}
	return *o.DnsRecord
}

// GetDnsRecordOk returns a tuple with the DnsRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnRpIdDomain) GetDnsRecordOk() (*DNSRecord, bool) {
	if o == nil || o.DnsRecord == nil {
		return nil, false
	}
	return o.DnsRecord, true
}

// HasDnsRecord returns a boolean if a field has been set.
func (o *WebAuthnRpIdDomain) HasDnsRecord() bool {
	if o != nil && o.DnsRecord != nil {
		return true
	}

	return false
}

// SetDnsRecord gets a reference to the given DNSRecord and assigns it to the DnsRecord field.
func (o *WebAuthnRpIdDomain) SetDnsRecord(v DNSRecord) {
	o.DnsRecord = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebAuthnRpIdDomain) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnRpIdDomain) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebAuthnRpIdDomain) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebAuthnRpIdDomain) SetName(v string) {
	o.Name = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *WebAuthnRpIdDomain) GetValidationStatus() string {
	if o == nil || o.ValidationStatus == nil {
		var ret string
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnRpIdDomain) GetValidationStatusOk() (*string, bool) {
	if o == nil || o.ValidationStatus == nil {
		return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *WebAuthnRpIdDomain) HasValidationStatus() bool {
	if o != nil && o.ValidationStatus != nil {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given string and assigns it to the ValidationStatus field.
func (o *WebAuthnRpIdDomain) SetValidationStatus(v string) {
	o.ValidationStatus = &v
}

func (o WebAuthnRpIdDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DnsRecord != nil {
		toSerialize["dnsRecord"] = o.DnsRecord
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ValidationStatus != nil {
		toSerialize["validationStatus"] = o.ValidationStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WebAuthnRpIdDomain) UnmarshalJSON(bytes []byte) (err error) {
	varWebAuthnRpIdDomain := _WebAuthnRpIdDomain{}

	err = json.Unmarshal(bytes, &varWebAuthnRpIdDomain)
	if err == nil {
		*o = WebAuthnRpIdDomain(varWebAuthnRpIdDomain)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "dnsRecord")
		delete(additionalProperties, "name")
		delete(additionalProperties, "validationStatus")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWebAuthnRpIdDomain struct {
	value *WebAuthnRpIdDomain
	isSet bool
}

func (v NullableWebAuthnRpIdDomain) Get() *WebAuthnRpIdDomain {
	return v.value
}

func (v *NullableWebAuthnRpIdDomain) Set(val *WebAuthnRpIdDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnRpIdDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnRpIdDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnRpIdDomain(val *WebAuthnRpIdDomain) *NullableWebAuthnRpIdDomain {
	return &NullableWebAuthnRpIdDomain{value: val, isSet: true}
}

func (v NullableWebAuthnRpIdDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnRpIdDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

