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

// Office365ApplicationSettingsApplication Office365 app instance properties
type Office365ApplicationSettingsApplication struct {
	// The domain for your Office 365 account
	Domain string `json:"domain"`
	// Microsoft tenant name
	MsftTenant string `json:"msftTenant"`
	AdditionalProperties map[string]interface{}
}

type _Office365ApplicationSettingsApplication Office365ApplicationSettingsApplication

// NewOffice365ApplicationSettingsApplication instantiates a new Office365ApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffice365ApplicationSettingsApplication(domain string, msftTenant string) *Office365ApplicationSettingsApplication {
	this := Office365ApplicationSettingsApplication{}
	this.Domain = domain
	this.MsftTenant = msftTenant
	return &this
}

// NewOffice365ApplicationSettingsApplicationWithDefaults instantiates a new Office365ApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffice365ApplicationSettingsApplicationWithDefaults() *Office365ApplicationSettingsApplication {
	this := Office365ApplicationSettingsApplication{}
	return &this
}

// GetDomain returns the Domain field value
func (o *Office365ApplicationSettingsApplication) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *Office365ApplicationSettingsApplication) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *Office365ApplicationSettingsApplication) SetDomain(v string) {
	o.Domain = v
}

// GetMsftTenant returns the MsftTenant field value
func (o *Office365ApplicationSettingsApplication) GetMsftTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsftTenant
}

// GetMsftTenantOk returns a tuple with the MsftTenant field value
// and a boolean to check if the value has been set.
func (o *Office365ApplicationSettingsApplication) GetMsftTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsftTenant, true
}

// SetMsftTenant sets field value
func (o *Office365ApplicationSettingsApplication) SetMsftTenant(v string) {
	o.MsftTenant = v
}

func (o Office365ApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["msftTenant"] = o.MsftTenant
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Office365ApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varOffice365ApplicationSettingsApplication := _Office365ApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varOffice365ApplicationSettingsApplication)
	if err == nil {
		*o = Office365ApplicationSettingsApplication(varOffice365ApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "msftTenant")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOffice365ApplicationSettingsApplication struct {
	value *Office365ApplicationSettingsApplication
	isSet bool
}

func (v NullableOffice365ApplicationSettingsApplication) Get() *Office365ApplicationSettingsApplication {
	return v.value
}

func (v *NullableOffice365ApplicationSettingsApplication) Set(val *Office365ApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableOffice365ApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableOffice365ApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffice365ApplicationSettingsApplication(val *Office365ApplicationSettingsApplication) *NullableOffice365ApplicationSettingsApplication {
	return &NullableOffice365ApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableOffice365ApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffice365ApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

