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

// IdentityProviderProperties The properties in the IdP `properties` object vary depending on the IdP type
type IdentityProviderProperties struct {
	// The [authentication assurance level](https://developers.login.gov/oidc/#aal-values) (AAL) value for the Login.gov IdP. See [Add a Login.gov IdP](https://developer.okta.com/docs/guides/add-logingov-idp/). Applies to `LOGINGOV` and `LOGINGOV_SANDBOX` IdP types.
	AalValue NullableString `json:"aalValue,omitempty"`
	// The additional Assurance Methods References (AMR) values for Smart Card IdPs. Applies to `X509` IdP type.
	AdditionalAmr []string `json:"additionalAmr,omitempty"`
	// The [type of identity verification](https://developers.login.gov/oidc/#ial-values) (IAL) value for the Login.gov IdP. See [Add a Login.gov IdP](https://developer.okta.com/docs/guides/add-logingov-idp/). Applies to `LOGINGOV` and `LOGINGOV_SANDBOX` IdP types.
	IalValue NullableString `json:"ialValue,omitempty"`
	// The ID of the inquiry template from your Persona dashboard. The inquiry template always starts with `itmpl`. Applies to the `IDV_PERSONA` IdP type.
	InquiryTemplateId string `json:"inquiryTemplateId"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderProperties IdentityProviderProperties

// NewIdentityProviderProperties instantiates a new IdentityProviderProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderProperties(inquiryTemplateId string) *IdentityProviderProperties {
	this := IdentityProviderProperties{}
	this.InquiryTemplateId = inquiryTemplateId
	return &this
}

// NewIdentityProviderPropertiesWithDefaults instantiates a new IdentityProviderProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderPropertiesWithDefaults() *IdentityProviderProperties {
	this := IdentityProviderProperties{}
	return &this
}

// GetAalValue returns the AalValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProviderProperties) GetAalValue() string {
	if o == nil || o.AalValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.AalValue.Get()
}

// GetAalValueOk returns a tuple with the AalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProviderProperties) GetAalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AalValue.Get(), o.AalValue.IsSet()
}

// HasAalValue returns a boolean if a field has been set.
func (o *IdentityProviderProperties) HasAalValue() bool {
	if o != nil && o.AalValue.IsSet() {
		return true
	}

	return false
}

// SetAalValue gets a reference to the given NullableString and assigns it to the AalValue field.
func (o *IdentityProviderProperties) SetAalValue(v string) {
	o.AalValue.Set(&v)
}
// SetAalValueNil sets the value for AalValue to be an explicit nil
func (o *IdentityProviderProperties) SetAalValueNil() {
	o.AalValue.Set(nil)
}

// UnsetAalValue ensures that no value is present for AalValue, not even an explicit nil
func (o *IdentityProviderProperties) UnsetAalValue() {
	o.AalValue.Unset()
}

// GetAdditionalAmr returns the AdditionalAmr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProviderProperties) GetAdditionalAmr() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AdditionalAmr
}

// GetAdditionalAmrOk returns a tuple with the AdditionalAmr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProviderProperties) GetAdditionalAmrOk() ([]string, bool) {
	if o == nil || o.AdditionalAmr == nil {
		return nil, false
	}
	return o.AdditionalAmr, true
}

// HasAdditionalAmr returns a boolean if a field has been set.
func (o *IdentityProviderProperties) HasAdditionalAmr() bool {
	if o != nil && o.AdditionalAmr != nil {
		return true
	}

	return false
}

// SetAdditionalAmr gets a reference to the given []string and assigns it to the AdditionalAmr field.
func (o *IdentityProviderProperties) SetAdditionalAmr(v []string) {
	o.AdditionalAmr = v
}

// GetIalValue returns the IalValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProviderProperties) GetIalValue() string {
	if o == nil || o.IalValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.IalValue.Get()
}

// GetIalValueOk returns a tuple with the IalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProviderProperties) GetIalValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IalValue.Get(), o.IalValue.IsSet()
}

// HasIalValue returns a boolean if a field has been set.
func (o *IdentityProviderProperties) HasIalValue() bool {
	if o != nil && o.IalValue.IsSet() {
		return true
	}

	return false
}

// SetIalValue gets a reference to the given NullableString and assigns it to the IalValue field.
func (o *IdentityProviderProperties) SetIalValue(v string) {
	o.IalValue.Set(&v)
}
// SetIalValueNil sets the value for IalValue to be an explicit nil
func (o *IdentityProviderProperties) SetIalValueNil() {
	o.IalValue.Set(nil)
}

// UnsetIalValue ensures that no value is present for IalValue, not even an explicit nil
func (o *IdentityProviderProperties) UnsetIalValue() {
	o.IalValue.Unset()
}

// GetInquiryTemplateId returns the InquiryTemplateId field value
func (o *IdentityProviderProperties) GetInquiryTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InquiryTemplateId
}

// GetInquiryTemplateIdOk returns a tuple with the InquiryTemplateId field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderProperties) GetInquiryTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InquiryTemplateId, true
}

// SetInquiryTemplateId sets field value
func (o *IdentityProviderProperties) SetInquiryTemplateId(v string) {
	o.InquiryTemplateId = v
}

func (o IdentityProviderProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AalValue.IsSet() {
		toSerialize["aalValue"] = o.AalValue.Get()
	}
	if o.AdditionalAmr != nil {
		toSerialize["additionalAmr"] = o.AdditionalAmr
	}
	if o.IalValue.IsSet() {
		toSerialize["ialValue"] = o.IalValue.Get()
	}
	if true {
		toSerialize["inquiryTemplateId"] = o.InquiryTemplateId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderProperties) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderProperties := _IdentityProviderProperties{}

	err = json.Unmarshal(bytes, &varIdentityProviderProperties)
	if err == nil {
		*o = IdentityProviderProperties(varIdentityProviderProperties)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "aalValue")
		delete(additionalProperties, "additionalAmr")
		delete(additionalProperties, "ialValue")
		delete(additionalProperties, "inquiryTemplateId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderProperties struct {
	value *IdentityProviderProperties
	isSet bool
}

func (v NullableIdentityProviderProperties) Get() *IdentityProviderProperties {
	return v.value
}

func (v *NullableIdentityProviderProperties) Set(val *IdentityProviderProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderProperties(val *IdentityProviderProperties) *NullableIdentityProviderProperties {
	return &NullableIdentityProviderProperties{value: val, isSet: true}
}

func (v NullableIdentityProviderProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

