/*
Okta Admin Management API

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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the CustomTelephonyProviderCredentialSendTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomTelephonyProviderCredentialSendTestRequest{}

// CustomTelephonyProviderCredentialSendTestRequest struct for CustomTelephonyProviderCredentialSendTestRequest
type CustomTelephonyProviderCredentialSendTestRequest struct {
	// The country code for the phone number. Use the [Alpha-2 code from ISO 3166-1](https://www.iso.org/obp/ui/#search) for country codes.
	CountryCodeIso2 *string `json:"countryCodeIso2,omitempty"`
	// The type of test message to send
	Factor *string `json:"factor,omitempty"`
	// The phone number to which the test message or call is sent
	PhoneNumber          *string `json:"phoneNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomTelephonyProviderCredentialSendTestRequest CustomTelephonyProviderCredentialSendTestRequest

// NewCustomTelephonyProviderCredentialSendTestRequest instantiates a new CustomTelephonyProviderCredentialSendTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomTelephonyProviderCredentialSendTestRequest() *CustomTelephonyProviderCredentialSendTestRequest {
	this := CustomTelephonyProviderCredentialSendTestRequest{}
	return &this
}

// NewCustomTelephonyProviderCredentialSendTestRequestWithDefaults instantiates a new CustomTelephonyProviderCredentialSendTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomTelephonyProviderCredentialSendTestRequestWithDefaults() *CustomTelephonyProviderCredentialSendTestRequest {
	this := CustomTelephonyProviderCredentialSendTestRequest{}
	return &this
}

// GetCountryCodeIso2 returns the CountryCodeIso2 field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialSendTestRequest) GetCountryCodeIso2() string {
	if o == nil || IsNil(o.CountryCodeIso2) {
		var ret string
		return ret
	}
	return *o.CountryCodeIso2
}

// GetCountryCodeIso2Ok returns a tuple with the CountryCodeIso2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialSendTestRequest) GetCountryCodeIso2Ok() (*string, bool) {
	if o == nil || IsNil(o.CountryCodeIso2) {
		return nil, false
	}
	return o.CountryCodeIso2, true
}

// HasCountryCodeIso2 returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialSendTestRequest) HasCountryCodeIso2() bool {
	if o != nil && !IsNil(o.CountryCodeIso2) {
		return true
	}

	return false
}

// SetCountryCodeIso2 gets a reference to the given string and assigns it to the CountryCodeIso2 field.
func (o *CustomTelephonyProviderCredentialSendTestRequest) SetCountryCodeIso2(v string) {
	o.CountryCodeIso2 = &v
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialSendTestRequest) GetFactor() string {
	if o == nil || IsNil(o.Factor) {
		var ret string
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialSendTestRequest) GetFactorOk() (*string, bool) {
	if o == nil || IsNil(o.Factor) {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialSendTestRequest) HasFactor() bool {
	if o != nil && !IsNil(o.Factor) {
		return true
	}

	return false
}

// SetFactor gets a reference to the given string and assigns it to the Factor field.
func (o *CustomTelephonyProviderCredentialSendTestRequest) SetFactor(v string) {
	o.Factor = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CustomTelephonyProviderCredentialSendTestRequest) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomTelephonyProviderCredentialSendTestRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CustomTelephonyProviderCredentialSendTestRequest) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CustomTelephonyProviderCredentialSendTestRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o CustomTelephonyProviderCredentialSendTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomTelephonyProviderCredentialSendTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCodeIso2) {
		toSerialize["countryCodeIso2"] = o.CountryCodeIso2
	}
	if !IsNil(o.Factor) {
		toSerialize["factor"] = o.Factor
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomTelephonyProviderCredentialSendTestRequest) UnmarshalJSON(data []byte) (err error) {
	varCustomTelephonyProviderCredentialSendTestRequest := _CustomTelephonyProviderCredentialSendTestRequest{}

	err = json.Unmarshal(data, &varCustomTelephonyProviderCredentialSendTestRequest)

	if err != nil {
		return err
	}

	*o = CustomTelephonyProviderCredentialSendTestRequest(varCustomTelephonyProviderCredentialSendTestRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "countryCodeIso2")
		delete(additionalProperties, "factor")
		delete(additionalProperties, "phoneNumber")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomTelephonyProviderCredentialSendTestRequest struct {
	value *CustomTelephonyProviderCredentialSendTestRequest
	isSet bool
}

func (v NullableCustomTelephonyProviderCredentialSendTestRequest) Get() *CustomTelephonyProviderCredentialSendTestRequest {
	return v.value
}

func (v *NullableCustomTelephonyProviderCredentialSendTestRequest) Set(val *CustomTelephonyProviderCredentialSendTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomTelephonyProviderCredentialSendTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomTelephonyProviderCredentialSendTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomTelephonyProviderCredentialSendTestRequest(val *CustomTelephonyProviderCredentialSendTestRequest) *NullableCustomTelephonyProviderCredentialSendTestRequest {
	return &NullableCustomTelephonyProviderCredentialSendTestRequest{value: val, isSet: true}
}

func (v NullableCustomTelephonyProviderCredentialSendTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomTelephonyProviderCredentialSendTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
