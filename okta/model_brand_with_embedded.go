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

// BrandWithEmbedded struct for BrandWithEmbedded
type BrandWithEmbedded struct {
	Embedded *map[string]interface{} `json:"_embedded,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
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

type _BrandWithEmbedded BrandWithEmbedded

// NewBrandWithEmbedded instantiates a new BrandWithEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandWithEmbedded() *BrandWithEmbedded {
	this := BrandWithEmbedded{}
	var removePoweredByOkta bool = false
	this.RemovePoweredByOkta = &removePoweredByOkta
	return &this
}

// NewBrandWithEmbeddedWithDefaults instantiates a new BrandWithEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandWithEmbeddedWithDefaults() *BrandWithEmbedded {
	this := BrandWithEmbedded{}
	var removePoweredByOkta bool = false
	this.RemovePoweredByOkta = &removePoweredByOkta
	return &this
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetEmbedded() map[string]interface{} {
	if o == nil || o.Embedded == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetEmbeddedOk() (*map[string]interface{}, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]interface{} and assigns it to the Embedded field.
func (o *BrandWithEmbedded) SetEmbedded(v map[string]interface{}) {
	o.Embedded = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *BrandWithEmbedded) SetLinks(v LinksSelf) {
	o.Links = &v
}

// GetAgreeToCustomPrivacyPolicy returns the AgreeToCustomPrivacyPolicy field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetAgreeToCustomPrivacyPolicy() bool {
	if o == nil || o.AgreeToCustomPrivacyPolicy == nil {
		var ret bool
		return ret
	}
	return *o.AgreeToCustomPrivacyPolicy
}

// GetAgreeToCustomPrivacyPolicyOk returns a tuple with the AgreeToCustomPrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetAgreeToCustomPrivacyPolicyOk() (*bool, bool) {
	if o == nil || o.AgreeToCustomPrivacyPolicy == nil {
		return nil, false
	}
	return o.AgreeToCustomPrivacyPolicy, true
}

// HasAgreeToCustomPrivacyPolicy returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasAgreeToCustomPrivacyPolicy() bool {
	if o != nil && o.AgreeToCustomPrivacyPolicy != nil {
		return true
	}

	return false
}

// SetAgreeToCustomPrivacyPolicy gets a reference to the given bool and assigns it to the AgreeToCustomPrivacyPolicy field.
func (o *BrandWithEmbedded) SetAgreeToCustomPrivacyPolicy(v bool) {
	o.AgreeToCustomPrivacyPolicy = &v
}

// GetCustomPrivacyPolicyUrl returns the CustomPrivacyPolicyUrl field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetCustomPrivacyPolicyUrl() string {
	if o == nil || o.CustomPrivacyPolicyUrl == nil {
		var ret string
		return ret
	}
	return *o.CustomPrivacyPolicyUrl
}

// GetCustomPrivacyPolicyUrlOk returns a tuple with the CustomPrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetCustomPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil || o.CustomPrivacyPolicyUrl == nil {
		return nil, false
	}
	return o.CustomPrivacyPolicyUrl, true
}

// HasCustomPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasCustomPrivacyPolicyUrl() bool {
	if o != nil && o.CustomPrivacyPolicyUrl != nil {
		return true
	}

	return false
}

// SetCustomPrivacyPolicyUrl gets a reference to the given string and assigns it to the CustomPrivacyPolicyUrl field.
func (o *BrandWithEmbedded) SetCustomPrivacyPolicyUrl(v string) {
	o.CustomPrivacyPolicyUrl = &v
}

// GetDefaultApp returns the DefaultApp field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetDefaultApp() DefaultApp {
	if o == nil || o.DefaultApp == nil {
		var ret DefaultApp
		return ret
	}
	return *o.DefaultApp
}

// GetDefaultAppOk returns a tuple with the DefaultApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetDefaultAppOk() (*DefaultApp, bool) {
	if o == nil || o.DefaultApp == nil {
		return nil, false
	}
	return o.DefaultApp, true
}

// HasDefaultApp returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasDefaultApp() bool {
	if o != nil && o.DefaultApp != nil {
		return true
	}

	return false
}

// SetDefaultApp gets a reference to the given DefaultApp and assigns it to the DefaultApp field.
func (o *BrandWithEmbedded) SetDefaultApp(v DefaultApp) {
	o.DefaultApp = &v
}

// GetEmailDomainId returns the EmailDomainId field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetEmailDomainId() string {
	if o == nil || o.EmailDomainId == nil {
		var ret string
		return ret
	}
	return *o.EmailDomainId
}

// GetEmailDomainIdOk returns a tuple with the EmailDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetEmailDomainIdOk() (*string, bool) {
	if o == nil || o.EmailDomainId == nil {
		return nil, false
	}
	return o.EmailDomainId, true
}

// HasEmailDomainId returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasEmailDomainId() bool {
	if o != nil && o.EmailDomainId != nil {
		return true
	}

	return false
}

// SetEmailDomainId gets a reference to the given string and assigns it to the EmailDomainId field.
func (o *BrandWithEmbedded) SetEmailDomainId(v string) {
	o.EmailDomainId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BrandWithEmbedded) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *BrandWithEmbedded) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *BrandWithEmbedded) SetLocale(v string) {
	o.Locale = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BrandWithEmbedded) SetName(v string) {
	o.Name = &v
}

// GetRemovePoweredByOkta returns the RemovePoweredByOkta field value if set, zero value otherwise.
func (o *BrandWithEmbedded) GetRemovePoweredByOkta() bool {
	if o == nil || o.RemovePoweredByOkta == nil {
		var ret bool
		return ret
	}
	return *o.RemovePoweredByOkta
}

// GetRemovePoweredByOktaOk returns a tuple with the RemovePoweredByOkta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandWithEmbedded) GetRemovePoweredByOktaOk() (*bool, bool) {
	if o == nil || o.RemovePoweredByOkta == nil {
		return nil, false
	}
	return o.RemovePoweredByOkta, true
}

// HasRemovePoweredByOkta returns a boolean if a field has been set.
func (o *BrandWithEmbedded) HasRemovePoweredByOkta() bool {
	if o != nil && o.RemovePoweredByOkta != nil {
		return true
	}

	return false
}

// SetRemovePoweredByOkta gets a reference to the given bool and assigns it to the RemovePoweredByOkta field.
func (o *BrandWithEmbedded) SetRemovePoweredByOkta(v bool) {
	o.RemovePoweredByOkta = &v
}

func (o BrandWithEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
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

func (o *BrandWithEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	varBrandWithEmbedded := _BrandWithEmbedded{}

	err = json.Unmarshal(bytes, &varBrandWithEmbedded)
	if err == nil {
		*o = BrandWithEmbedded(varBrandWithEmbedded)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
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

type NullableBrandWithEmbedded struct {
	value *BrandWithEmbedded
	isSet bool
}

func (v NullableBrandWithEmbedded) Get() *BrandWithEmbedded {
	return v.value
}

func (v *NullableBrandWithEmbedded) Set(val *BrandWithEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandWithEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandWithEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandWithEmbedded(val *BrandWithEmbedded) *NullableBrandWithEmbedded {
	return &NullableBrandWithEmbedded{value: val, isSet: true}
}

func (v NullableBrandWithEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandWithEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

