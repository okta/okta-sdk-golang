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

// SalesforceApplicationSettingsApplication Salesforce app instance properties
type SalesforceApplicationSettingsApplication struct {
	// Salesforce instance that you want to connect to
	InstanceType string `json:"instanceType"`
	// Salesforce integration type
	IntegrationType string `json:"integrationType"`
	// The Login URL specified in your Salesforce Single Sign-On settings
	LoginUrl *string `json:"loginUrl,omitempty"`
	// Salesforce Logout URL
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SalesforceApplicationSettingsApplication SalesforceApplicationSettingsApplication

// NewSalesforceApplicationSettingsApplication instantiates a new SalesforceApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalesforceApplicationSettingsApplication(instanceType string, integrationType string) *SalesforceApplicationSettingsApplication {
	this := SalesforceApplicationSettingsApplication{}
	this.InstanceType = instanceType
	this.IntegrationType = integrationType
	return &this
}

// NewSalesforceApplicationSettingsApplicationWithDefaults instantiates a new SalesforceApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalesforceApplicationSettingsApplicationWithDefaults() *SalesforceApplicationSettingsApplication {
	this := SalesforceApplicationSettingsApplication{}
	return &this
}

// GetInstanceType returns the InstanceType field value
func (o *SalesforceApplicationSettingsApplication) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *SalesforceApplicationSettingsApplication) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value
func (o *SalesforceApplicationSettingsApplication) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetIntegrationType returns the IntegrationType field value
func (o *SalesforceApplicationSettingsApplication) GetIntegrationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *SalesforceApplicationSettingsApplication) GetIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value
func (o *SalesforceApplicationSettingsApplication) SetIntegrationType(v string) {
	o.IntegrationType = v
}

// GetLoginUrl returns the LoginUrl field value if set, zero value otherwise.
func (o *SalesforceApplicationSettingsApplication) GetLoginUrl() string {
	if o == nil || o.LoginUrl == nil {
		var ret string
		return ret
	}
	return *o.LoginUrl
}

// GetLoginUrlOk returns a tuple with the LoginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceApplicationSettingsApplication) GetLoginUrlOk() (*string, bool) {
	if o == nil || o.LoginUrl == nil {
		return nil, false
	}
	return o.LoginUrl, true
}

// HasLoginUrl returns a boolean if a field has been set.
func (o *SalesforceApplicationSettingsApplication) HasLoginUrl() bool {
	if o != nil && o.LoginUrl != nil {
		return true
	}

	return false
}

// SetLoginUrl gets a reference to the given string and assigns it to the LoginUrl field.
func (o *SalesforceApplicationSettingsApplication) SetLoginUrl(v string) {
	o.LoginUrl = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *SalesforceApplicationSettingsApplication) GetLogoutUrl() string {
	if o == nil || o.LogoutUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceApplicationSettingsApplication) GetLogoutUrlOk() (*string, bool) {
	if o == nil || o.LogoutUrl == nil {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *SalesforceApplicationSettingsApplication) HasLogoutUrl() bool {
	if o != nil && o.LogoutUrl != nil {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *SalesforceApplicationSettingsApplication) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

func (o SalesforceApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["instanceType"] = o.InstanceType
	}
	if true {
		toSerialize["integrationType"] = o.IntegrationType
	}
	if o.LoginUrl != nil {
		toSerialize["loginUrl"] = o.LoginUrl
	}
	if o.LogoutUrl != nil {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SalesforceApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varSalesforceApplicationSettingsApplication := _SalesforceApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varSalesforceApplicationSettingsApplication)
	if err == nil {
		*o = SalesforceApplicationSettingsApplication(varSalesforceApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "instanceType")
		delete(additionalProperties, "integrationType")
		delete(additionalProperties, "loginUrl")
		delete(additionalProperties, "logoutUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSalesforceApplicationSettingsApplication struct {
	value *SalesforceApplicationSettingsApplication
	isSet bool
}

func (v NullableSalesforceApplicationSettingsApplication) Get() *SalesforceApplicationSettingsApplication {
	return v.value
}

func (v *NullableSalesforceApplicationSettingsApplication) Set(val *SalesforceApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableSalesforceApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableSalesforceApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalesforceApplicationSettingsApplication(val *SalesforceApplicationSettingsApplication) *NullableSalesforceApplicationSettingsApplication {
	return &NullableSalesforceApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableSalesforceApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalesforceApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

