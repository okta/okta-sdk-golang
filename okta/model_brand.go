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

// Brand struct for Brand
type Brand struct {
	// Consent for updating the custom privacy URL. Not required when resetting the URL.
	AgreeToCustomPrivacyPolicy *bool `json:"agreeToCustomPrivacyPolicy,omitempty"`
	// Custom privacy policy URL
	CustomPrivacyPolicyUrl *string `json:"customPrivacyPolicyUrl,omitempty"`
	DefaultApp *DefaultApp `json:"defaultApp,omitempty"`
	// The ID of the email domain
	EmailDomainId *string `json:"emailDomainId,omitempty"`
	// The Brand ID
	Id *string `json:"id,omitempty"`
	// If `true`, the Brand is used for the Okta subdomain
	IsDefault *bool `json:"isDefault,omitempty"`
	// The language specified as an [IETF BCP 47 language tag](https://datatracker.ietf.org/doc/html/rfc5646)
	Locale *string `json:"locale,omitempty"`
	// The name of the Brand
	Name *string `json:"name,omitempty"`
	// Removes \"Powered by Okta\" from the sign-in page in redirect authentication deployments, and \"© [current year] Okta, Inc.\" from the Okta End-User Dashboard
	RemovePoweredByOkta *bool `json:"removePoweredByOkta,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Brand Brand

// NewBrand instantiates a new Brand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrand() *Brand {
	this := Brand{}
	var removePoweredByOkta bool = false
	this.RemovePoweredByOkta = &removePoweredByOkta
	return &this
}

// NewBrandWithDefaults instantiates a new Brand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandWithDefaults() *Brand {
	this := Brand{}
	var removePoweredByOkta bool = false
	this.RemovePoweredByOkta = &removePoweredByOkta
	return &this
}

// GetAgreeToCustomPrivacyPolicy returns the AgreeToCustomPrivacyPolicy field value if set, zero value otherwise.
func (o *Brand) GetAgreeToCustomPrivacyPolicy() bool {
	if o == nil || o.AgreeToCustomPrivacyPolicy == nil {
		var ret bool
		return ret
	}
	return *o.AgreeToCustomPrivacyPolicy
}

// GetAgreeToCustomPrivacyPolicyOk returns a tuple with the AgreeToCustomPrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetAgreeToCustomPrivacyPolicyOk() (*bool, bool) {
	if o == nil || o.AgreeToCustomPrivacyPolicy == nil {
		return nil, false
	}
	return o.AgreeToCustomPrivacyPolicy, true
}

// HasAgreeToCustomPrivacyPolicy returns a boolean if a field has been set.
func (o *Brand) HasAgreeToCustomPrivacyPolicy() bool {
	if o != nil && o.AgreeToCustomPrivacyPolicy != nil {
		return true
	}

	return false
}

// SetAgreeToCustomPrivacyPolicy gets a reference to the given bool and assigns it to the AgreeToCustomPrivacyPolicy field.
func (o *Brand) SetAgreeToCustomPrivacyPolicy(v bool) {
	o.AgreeToCustomPrivacyPolicy = &v
}

// GetCustomPrivacyPolicyUrl returns the CustomPrivacyPolicyUrl field value if set, zero value otherwise.
func (o *Brand) GetCustomPrivacyPolicyUrl() string {
	if o == nil || o.CustomPrivacyPolicyUrl == nil {
		var ret string
		return ret
	}
	return *o.CustomPrivacyPolicyUrl
}

// GetCustomPrivacyPolicyUrlOk returns a tuple with the CustomPrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetCustomPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil || o.CustomPrivacyPolicyUrl == nil {
		return nil, false
	}
	return o.CustomPrivacyPolicyUrl, true
}

// HasCustomPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *Brand) HasCustomPrivacyPolicyUrl() bool {
	if o != nil && o.CustomPrivacyPolicyUrl != nil {
		return true
	}

	return false
}

// SetCustomPrivacyPolicyUrl gets a reference to the given string and assigns it to the CustomPrivacyPolicyUrl field.
func (o *Brand) SetCustomPrivacyPolicyUrl(v string) {
	o.CustomPrivacyPolicyUrl = &v
}

// GetDefaultApp returns the DefaultApp field value if set, zero value otherwise.
func (o *Brand) GetDefaultApp() DefaultApp {
	if o == nil || o.DefaultApp == nil {
		var ret DefaultApp
		return ret
	}
	return *o.DefaultApp
}

// GetDefaultAppOk returns a tuple with the DefaultApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetDefaultAppOk() (*DefaultApp, bool) {
	if o == nil || o.DefaultApp == nil {
		return nil, false
	}
	return o.DefaultApp, true
}

// HasDefaultApp returns a boolean if a field has been set.
func (o *Brand) HasDefaultApp() bool {
	if o != nil && o.DefaultApp != nil {
		return true
	}

	return false
}

// SetDefaultApp gets a reference to the given DefaultApp and assigns it to the DefaultApp field.
func (o *Brand) SetDefaultApp(v DefaultApp) {
	o.DefaultApp = &v
}

// GetEmailDomainId returns the EmailDomainId field value if set, zero value otherwise.
func (o *Brand) GetEmailDomainId() string {
	if o == nil || o.EmailDomainId == nil {
		var ret string
		return ret
	}
	return *o.EmailDomainId
}

// GetEmailDomainIdOk returns a tuple with the EmailDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetEmailDomainIdOk() (*string, bool) {
	if o == nil || o.EmailDomainId == nil {
		return nil, false
	}
	return o.EmailDomainId, true
}

// HasEmailDomainId returns a boolean if a field has been set.
func (o *Brand) HasEmailDomainId() bool {
	if o != nil && o.EmailDomainId != nil {
		return true
	}

	return false
}

// SetEmailDomainId gets a reference to the given string and assigns it to the EmailDomainId field.
func (o *Brand) SetEmailDomainId(v string) {
	o.EmailDomainId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Brand) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Brand) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Brand) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *Brand) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *Brand) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *Brand) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *Brand) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *Brand) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *Brand) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Brand) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Brand) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Brand) SetName(v string) {
	o.Name = &v
}

// GetRemovePoweredByOkta returns the RemovePoweredByOkta field value if set, zero value otherwise.
func (o *Brand) GetRemovePoweredByOkta() bool {
	if o == nil || o.RemovePoweredByOkta == nil {
		var ret bool
		return ret
	}
	return *o.RemovePoweredByOkta
}

// GetRemovePoweredByOktaOk returns a tuple with the RemovePoweredByOkta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetRemovePoweredByOktaOk() (*bool, bool) {
	if o == nil || o.RemovePoweredByOkta == nil {
		return nil, false
	}
	return o.RemovePoweredByOkta, true
}

// HasRemovePoweredByOkta returns a boolean if a field has been set.
func (o *Brand) HasRemovePoweredByOkta() bool {
	if o != nil && o.RemovePoweredByOkta != nil {
		return true
	}

	return false
}

// SetRemovePoweredByOkta gets a reference to the given bool and assigns it to the RemovePoweredByOkta field.
func (o *Brand) SetRemovePoweredByOkta(v bool) {
	o.RemovePoweredByOkta = &v
}

func (o Brand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AgreeToCustomPrivacyPolicy != nil {
		toSerialize["agreeToCustomPrivacyPolicy"] = o.AgreeToCustomPrivacyPolicy
	}
	if o.CustomPrivacyPolicyUrl != nil {
		toSerialize["customPrivacyPolicyUrl"] = o.CustomPrivacyPolicyUrl
	}
	if o.DefaultApp != nil {
		toSerialize["defaultApp"] = o.DefaultApp
	}
	if o.EmailDomainId != nil {
		toSerialize["emailDomainId"] = o.EmailDomainId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RemovePoweredByOkta != nil {
		toSerialize["removePoweredByOkta"] = o.RemovePoweredByOkta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Brand) UnmarshalJSON(bytes []byte) (err error) {
	varBrand := _Brand{}

	err = json.Unmarshal(bytes, &varBrand)
	if err == nil {
		*o = Brand(varBrand)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "agreeToCustomPrivacyPolicy")
		delete(additionalProperties, "customPrivacyPolicyUrl")
		delete(additionalProperties, "defaultApp")
		delete(additionalProperties, "emailDomainId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "isDefault")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "name")
		delete(additionalProperties, "removePoweredByOkta")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBrand struct {
	value *Brand
	isSet bool
}

func (v NullableBrand) Get() *Brand {
	return v.value
}

func (v *NullableBrand) Set(val *Brand) {
	v.value = val
	v.isSet = true
}

func (v NullableBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrand(val *Brand) *NullableBrand {
	return &NullableBrand{value: val, isSet: true}
}

func (v NullableBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

