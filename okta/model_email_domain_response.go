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

// EmailDomainResponse struct for EmailDomainResponse
type EmailDomainResponse struct {
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

type _EmailDomainResponse EmailDomainResponse

// NewEmailDomainResponse instantiates a new EmailDomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomainResponse(displayName string, userName string) *EmailDomainResponse {
	this := EmailDomainResponse{}
	this.DisplayName = displayName
	this.UserName = userName
	return &this
}

// NewEmailDomainResponseWithDefaults instantiates a new EmailDomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainResponseWithDefaults() *EmailDomainResponse {
	this := EmailDomainResponse{}
	var validationSubdomain string = "mail"
	this.ValidationSubdomain = &validationSubdomain
	return &this
}

// GetDnsValidationRecords returns the DnsValidationRecords field value if set, zero value otherwise.
func (o *EmailDomainResponse) GetDnsValidationRecords() []EmailDomainDNSRecord {
	if o == nil || o.DnsValidationRecords == nil {
		var ret []EmailDomainDNSRecord
		return ret
	}
	return o.DnsValidationRecords
}

// GetDnsValidationRecordsOk returns a tuple with the DnsValidationRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponse) GetDnsValidationRecordsOk() ([]EmailDomainDNSRecord, bool) {
	if o == nil || o.DnsValidationRecords == nil {
		return nil, false
	}
	return o.DnsValidationRecords, true
}

// HasDnsValidationRecords returns a boolean if a field has been set.
func (o *EmailDomainResponse) HasDnsValidationRecords() bool {
	if o != nil && o.DnsValidationRecords != nil {
		return true
	}

	return false
}

// SetDnsValidationRecords gets a reference to the given []EmailDomainDNSRecord and assigns it to the DnsValidationRecords field.
func (o *EmailDomainResponse) SetDnsValidationRecords(v []EmailDomainDNSRecord) {
	o.DnsValidationRecords = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *EmailDomainResponse) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponse) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *EmailDomainResponse) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *EmailDomainResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailDomainResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailDomainResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailDomainResponse) SetId(v string) {
	o.Id = &v
}

// GetValidationStatus returns the ValidationStatus field value if set, zero value otherwise.
func (o *EmailDomainResponse) GetValidationStatus() string {
	if o == nil || o.ValidationStatus == nil {
		var ret string
		return ret
	}
	return *o.ValidationStatus
}

// GetValidationStatusOk returns a tuple with the ValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponse) GetValidationStatusOk() (*string, bool) {
	if o == nil || o.ValidationStatus == nil {
		return nil, false
	}
	return o.ValidationStatus, true
}

// HasValidationStatus returns a boolean if a field has been set.
func (o *EmailDomainResponse) HasValidationStatus() bool {
	if o != nil && o.ValidationStatus != nil {
		return true
	}

	return false
}

// SetValidationStatus gets a reference to the given string and assigns it to the ValidationStatus field.
func (o *EmailDomainResponse) SetValidationStatus(v string) {
	o.ValidationStatus = &v
}

// GetValidationSubdomain returns the ValidationSubdomain field value if set, zero value otherwise.
func (o *EmailDomainResponse) GetValidationSubdomain() string {
	if o == nil || o.ValidationSubdomain == nil {
		var ret string
		return ret
	}
	return *o.ValidationSubdomain
}

// GetValidationSubdomainOk returns a tuple with the ValidationSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainResponse) GetValidationSubdomainOk() (*string, bool) {
	if o == nil || o.ValidationSubdomain == nil {
		return nil, false
	}
	return o.ValidationSubdomain, true
}

// HasValidationSubdomain returns a boolean if a field has been set.
func (o *EmailDomainResponse) HasValidationSubdomain() bool {
	if o != nil && o.ValidationSubdomain != nil {
		return true
	}

	return false
}

// SetValidationSubdomain gets a reference to the given string and assigns it to the ValidationSubdomain field.
func (o *EmailDomainResponse) SetValidationSubdomain(v string) {
	o.ValidationSubdomain = &v
}

// GetDisplayName returns the DisplayName field value
func (o *EmailDomainResponse) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EmailDomainResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EmailDomainResponse) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetUserName returns the UserName field value
func (o *EmailDomainResponse) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *EmailDomainResponse) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *EmailDomainResponse) SetUserName(v string) {
	o.UserName = v
}

func (o EmailDomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

func (o *EmailDomainResponse) UnmarshalJSON(bytes []byte) (err error) {
	varEmailDomainResponse := _EmailDomainResponse{}

	err = json.Unmarshal(bytes, &varEmailDomainResponse)
	if err == nil {
		*o = EmailDomainResponse(varEmailDomainResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
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

type NullableEmailDomainResponse struct {
	value *EmailDomainResponse
	isSet bool
}

func (v NullableEmailDomainResponse) Get() *EmailDomainResponse {
	return v.value
}

func (v *NullableEmailDomainResponse) Set(val *EmailDomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomainResponse(val *EmailDomainResponse) *NullableEmailDomainResponse {
	return &NullableEmailDomainResponse{value: val, isSet: true}
}

func (v NullableEmailDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

