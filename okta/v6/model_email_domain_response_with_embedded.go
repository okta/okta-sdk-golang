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
	"fmt"
)

// EmailDomainResponseWithEmbedded struct for EmailDomainResponseWithEmbedded
type EmailDomainResponseWithEmbedded struct {
	Embedded *map[string]interface{} `json:"_embedded,omitempty"`
	DnsValidationRecords []EmailDomainDNSRecord `json:"dnsValidationRecords,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Id *string `json:"id,omitempty"`
	ValidationStatus *string `json:"validationStatus,omitempty"`
	// The subdomain for the email sender's custom mail domain
	ValidationSubdomain *string `json:"validationSubdomain,omitempty"`
	DisplayName string `json:"displayName"`
	UserName string `json:"userName"`
	AdditionalProperties map[string]interface{}
}

type _EmailDomainResponseWithEmbedded EmailDomainResponseWithEmbedded

// NewEmailDomainResponseWithEmbedded instantiates a new EmailDomainResponseWithEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomainResponseWithEmbedded(displayName string, userName string) *EmailDomainResponseWithEmbedded {
	this := EmailDomainResponseWithEmbedded{}
	var validationSubdomain string = "mail"
	this.ValidationSubdomain = &validationSubdomain
	this.DisplayName = displayName
	this.UserName = userName
	return &this
}

// NewEmailDomainResponseWithEmbeddedWithDefaults instantiates a new EmailDomainResponseWithEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainResponseWithEmbeddedWithDefaults() *EmailDomainResponseWithEmbedded {
	this := EmailDomainResponseWithEmbedded{}
	var validationSubdomain string = "mail"
	this.ValidationSubdomain = &validationSubdomain
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *EmailDomainResponseWithEmbedded) GetEmbedded() map[string]interface{} {
	if o == nil || o.Embedded == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponseWithEmbedded) GetEmbeddedOk() (*map[string]interface{}, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *EmailDomainResponseWithEmbedded) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]interface{} and assigns it to the Embedded field.
func (o *EmailDomainResponseWithEmbedded) SetEmbedded(v map[string]interface{}) {
	o.Embedded = &v
}

// GetDnsValidationRecords returns the DnsValidationRecords field value if set, zero value otherwise.
func (o *EmailDomainResponseWithEmbedded) GetDnsValidationRecords() []EmailDomainDNSRecord {
	if o == nil || o.DnsValidationRecords == nil {
		var ret []EmailDomainDNSRecord
		return ret
	}
	return o.DnsValidationRecords
}

// GetDnsValidationRecordsOk returns a tuple with the DnsValidationRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponseWithEmbedded) GetDnsValidationRecordsOk() ([]EmailDomainDNSRecord, bool) {
	if o == nil || o.DnsValidationRecords == nil {
		return nil, false
	}
	return o.DnsValidationRecords, true
}

// HasDnsValidationRecords returns a boolean if a field has been set.
func (o *EmailDomainResponseWithEmbedded) HasDnsValidationRecords() bool {
	if o != nil && o.DnsValidationRecords != nil {
		return true
	}

	return false
}

// SetDnsValidationRecords gets a reference to the given []EmailDomainDNSRecord and assigns it to the DnsValidationRecords field.
func (o *EmailDomainResponseWithEmbedded) SetDnsValidationRecords(v []EmailDomainDNSRecord) {
	o.DnsValidationRecords = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *EmailDomainResponseWithEmbedded) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponseWithEmbedded) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *EmailDomainResponseWithEmbedded) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *EmailDomainResponseWithEmbedded) SetDomain(v string) {
	o.Domain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailDomainResponseWithEmbedded) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponseWithEmbedded) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailDomainResponseWithEmbedded) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailDomainResponseWithEmbedded) SetId(v string) {
	o.Id = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *EmailDomainResponseWithEmbedded) GetValidationStatus() string {
	if o == nil || o.ValidationStatus == nil {
		var ret string
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponseWithEmbedded) GetValidationStatusOk() (*string, bool) {
	if o == nil || o.ValidationStatus == nil {
		return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *EmailDomainResponseWithEmbedded) HasValidationStatus() bool {
	if o != nil && o.ValidationStatus != nil {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given string and assigns it to the ValidationStatus field.
func (o *EmailDomainResponseWithEmbedded) SetValidationStatus(v string) {
	o.ValidationStatus = &v
}

// GetValidationSubdomain returns the ValidationSubdomain field value if set, zero value otherwise.
func (o *EmailDomainResponseWithEmbedded) GetValidationSubdomain() string {
	if o == nil || o.ValidationSubdomain == nil {
		var ret string
		return ret
	}
	return *o.ValidationSubdomain
}

// GetValidationSubdomainOk returns a tuple with the ValidationSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponseWithEmbedded) GetValidationSubdomainOk() (*string, bool) {
	if o == nil || o.ValidationSubdomain == nil {
		return nil, false
	}
	return o.ValidationSubdomain, true
}

// HasValidationSubdomain returns a boolean if a field has been set.
func (o *EmailDomainResponseWithEmbedded) HasValidationSubdomain() bool {
	if o != nil && o.ValidationSubdomain != nil {
		return true
	}

	return false
}

// SetValidationSubdomain gets a reference to the given string and assigns it to the ValidationSubdomain field.
func (o *EmailDomainResponseWithEmbedded) SetValidationSubdomain(v string) {
	o.ValidationSubdomain = &v
}

// GetDisplayName returns the DisplayName field value
func (o *EmailDomainResponseWithEmbedded) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EmailDomainResponseWithEmbedded) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EmailDomainResponseWithEmbedded) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetUserName returns the UserName field value
func (o *EmailDomainResponseWithEmbedded) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *EmailDomainResponseWithEmbedded) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *EmailDomainResponseWithEmbedded) SetUserName(v string) {
	o.UserName = v
}

func (o EmailDomainResponseWithEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.DnsValidationRecords != nil {
		toSerialize["dnsValidationRecords"] = o.DnsValidationRecords
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ValidationStatus != nil {
		toSerialize["validationStatus"] = o.ValidationStatus
	}
	if o.ValidationSubdomain != nil {
		toSerialize["validationSubdomain"] = o.ValidationSubdomain
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["userName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmailDomainResponseWithEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	varEmailDomainResponseWithEmbedded := _EmailDomainResponseWithEmbedded{}

	err = json.Unmarshal(bytes, &varEmailDomainResponseWithEmbedded)
	if err == nil {
		*o = EmailDomainResponseWithEmbedded(varEmailDomainResponseWithEmbedded)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "dnsValidationRecords")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "id")
		delete(additionalProperties, "validationStatus")
		delete(additionalProperties, "validationSubdomain")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailDomainResponseWithEmbedded struct {
	value *EmailDomainResponseWithEmbedded
	isSet bool
}

func (v NullableEmailDomainResponseWithEmbedded) Get() *EmailDomainResponseWithEmbedded {
	return v.value
}

func (v *NullableEmailDomainResponseWithEmbedded) Set(val *EmailDomainResponseWithEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomainResponseWithEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomainResponseWithEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomainResponseWithEmbedded(val *EmailDomainResponseWithEmbedded) *NullableEmailDomainResponseWithEmbedded {
	return &NullableEmailDomainResponseWithEmbedded{value: val, isSet: true}
}

func (v NullableEmailDomainResponseWithEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomainResponseWithEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

