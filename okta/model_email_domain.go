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

// EmailDomain struct for EmailDomain
type EmailDomain struct {
	BrandId string `json:"brandId"`
	Domain string `json:"domain"`
	// Subdomain for the email sender's custom mail domain. Specify your subdomain when you configure a custom mail domain.
	ValidationSubdomain *string `json:"validationSubdomain,omitempty"`
	DisplayName string `json:"displayName"`
	UserName string `json:"userName"`
	AdditionalProperties map[string]interface{}
}

type _EmailDomain EmailDomain

// NewEmailDomain instantiates a new EmailDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomain(brandId string, domain string, displayName string, userName string) *EmailDomain {
	this := EmailDomain{}
	this.DisplayName = displayName
	this.UserName = userName
	return &this
}

// NewEmailDomainWithDefaults instantiates a new EmailDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainWithDefaults() *EmailDomain {
	this := EmailDomain{}
	var validationSubdomain string = "mail"
	this.ValidationSubdomain = &validationSubdomain
	return &this
}

// GetBrandId returns the BrandId field value
func (o *EmailDomain) GetBrandId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetBrandIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandId, true
}

// SetBrandId sets field value
func (o *EmailDomain) SetBrandId(v string) {
	o.BrandId = v
}

// GetDomain returns the Domain field value
func (o *EmailDomain) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *EmailDomain) SetDomain(v string) {
	o.Domain = v
}

// GetValidationSubdomain returns the ValidationSubdomain field value if set, zero value otherwise.
func (o *EmailDomain) GetValidationSubdomain() string {
	if o == nil || o.ValidationSubdomain == nil {
		var ret string
		return ret
	}
	return *o.ValidationSubdomain
}

// GetValidationSubdomainOk returns a tuple with the ValidationSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetValidationSubdomainOk() (*string, bool) {
	if o == nil || o.ValidationSubdomain == nil {
		return nil, false
	}
	return o.ValidationSubdomain, true
}

// HasValidationSubdomain returns a boolean if a field has been set.
func (o *EmailDomain) HasValidationSubdomain() bool {
	if o != nil && o.ValidationSubdomain != nil {
		return true
	}

	return false
}

// SetValidationSubdomain gets a reference to the given string and assigns it to the ValidationSubdomain field.
func (o *EmailDomain) SetValidationSubdomain(v string) {
	o.ValidationSubdomain = &v
}

// GetDisplayName returns the DisplayName field value
func (o *EmailDomain) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *EmailDomain) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetUserName returns the UserName field value
func (o *EmailDomain) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *EmailDomain) SetUserName(v string) {
	o.UserName = v
}

func (o EmailDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["brandId"] = o.BrandId
	}
	if true {
		toSerialize["domain"] = o.Domain
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

func (o *EmailDomain) UnmarshalJSON(bytes []byte) (err error) {
	varEmailDomain := _EmailDomain{}

	err = json.Unmarshal(bytes, &varEmailDomain)
	if err == nil {
		*o = EmailDomain(varEmailDomain)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "brandId")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "validationSubdomain")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "userName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEmailDomain struct {
	value *EmailDomain
	isSet bool
}

func (v NullableEmailDomain) Get() *EmailDomain {
	return v.value
}

func (v *NullableEmailDomain) Set(val *EmailDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomain(val *EmailDomain) *NullableEmailDomain {
	return &NullableEmailDomain{value: val, isSet: true}
}

func (v NullableEmailDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

