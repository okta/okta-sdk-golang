/*
Okta Admin Management

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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the IdentityProviderPropertiesIdvMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityProviderPropertiesIdvMetadata{}

// IdentityProviderPropertiesIdvMetadata Metadata about the IDV vendor. Available only for `IDV_STANDARD` IdPs.
type IdentityProviderPropertiesIdvMetadata struct {
	// The display name of the IDV vendor
	VendorDisplayName *string `json:"vendorDisplayName,omitempty"`
	// A URL that links to the terms of use for the IDV vendor
	TermsOfUse *string `json:"termsOfUse,omitempty"`
	// A URL that links to the privacy policy for the IDV vendor
	PrivacyPolicy        *string `json:"privacyPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderPropertiesIdvMetadata IdentityProviderPropertiesIdvMetadata

// NewIdentityProviderPropertiesIdvMetadata instantiates a new IdentityProviderPropertiesIdvMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderPropertiesIdvMetadata() *IdentityProviderPropertiesIdvMetadata {
	this := IdentityProviderPropertiesIdvMetadata{}
	return &this
}

// NewIdentityProviderPropertiesIdvMetadataWithDefaults instantiates a new IdentityProviderPropertiesIdvMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderPropertiesIdvMetadataWithDefaults() *IdentityProviderPropertiesIdvMetadata {
	this := IdentityProviderPropertiesIdvMetadata{}
	return &this
}

// GetVendorDisplayName returns the VendorDisplayName field value if set, zero value otherwise.
func (o *IdentityProviderPropertiesIdvMetadata) GetVendorDisplayName() string {
	if o == nil || IsNil(o.VendorDisplayName) {
		var ret string
		return ret
	}
	return *o.VendorDisplayName
}

// GetVendorDisplayNameOk returns a tuple with the VendorDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPropertiesIdvMetadata) GetVendorDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorDisplayName) {
		return nil, false
	}
	return o.VendorDisplayName, true
}

// HasVendorDisplayName returns a boolean if a field has been set.
func (o *IdentityProviderPropertiesIdvMetadata) HasVendorDisplayName() bool {
	if o != nil && !IsNil(o.VendorDisplayName) {
		return true
	}

	return false
}

// SetVendorDisplayName gets a reference to the given string and assigns it to the VendorDisplayName field.
func (o *IdentityProviderPropertiesIdvMetadata) SetVendorDisplayName(v string) {
	o.VendorDisplayName = &v
}

// GetTermsOfUse returns the TermsOfUse field value if set, zero value otherwise.
func (o *IdentityProviderPropertiesIdvMetadata) GetTermsOfUse() string {
	if o == nil || IsNil(o.TermsOfUse) {
		var ret string
		return ret
	}
	return *o.TermsOfUse
}

// GetTermsOfUseOk returns a tuple with the TermsOfUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPropertiesIdvMetadata) GetTermsOfUseOk() (*string, bool) {
	if o == nil || IsNil(o.TermsOfUse) {
		return nil, false
	}
	return o.TermsOfUse, true
}

// HasTermsOfUse returns a boolean if a field has been set.
func (o *IdentityProviderPropertiesIdvMetadata) HasTermsOfUse() bool {
	if o != nil && !IsNil(o.TermsOfUse) {
		return true
	}

	return false
}

// SetTermsOfUse gets a reference to the given string and assigns it to the TermsOfUse field.
func (o *IdentityProviderPropertiesIdvMetadata) SetTermsOfUse(v string) {
	o.TermsOfUse = &v
}

// GetPrivacyPolicy returns the PrivacyPolicy field value if set, zero value otherwise.
func (o *IdentityProviderPropertiesIdvMetadata) GetPrivacyPolicy() string {
	if o == nil || IsNil(o.PrivacyPolicy) {
		var ret string
		return ret
	}
	return *o.PrivacyPolicy
}

// GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPropertiesIdvMetadata) GetPrivacyPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPolicy) {
		return nil, false
	}
	return o.PrivacyPolicy, true
}

// HasPrivacyPolicy returns a boolean if a field has been set.
func (o *IdentityProviderPropertiesIdvMetadata) HasPrivacyPolicy() bool {
	if o != nil && !IsNil(o.PrivacyPolicy) {
		return true
	}

	return false
}

// SetPrivacyPolicy gets a reference to the given string and assigns it to the PrivacyPolicy field.
func (o *IdentityProviderPropertiesIdvMetadata) SetPrivacyPolicy(v string) {
	o.PrivacyPolicy = &v
}

func (o IdentityProviderPropertiesIdvMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityProviderPropertiesIdvMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VendorDisplayName) {
		toSerialize["vendorDisplayName"] = o.VendorDisplayName
	}
	if !IsNil(o.TermsOfUse) {
		toSerialize["termsOfUse"] = o.TermsOfUse
	}
	if !IsNil(o.PrivacyPolicy) {
		toSerialize["privacyPolicy"] = o.PrivacyPolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IdentityProviderPropertiesIdvMetadata) UnmarshalJSON(data []byte) (err error) {
	varIdentityProviderPropertiesIdvMetadata := _IdentityProviderPropertiesIdvMetadata{}

	err = json.Unmarshal(data, &varIdentityProviderPropertiesIdvMetadata)

	if err != nil {
		return err
	}

	*o = IdentityProviderPropertiesIdvMetadata(varIdentityProviderPropertiesIdvMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vendorDisplayName")
		delete(additionalProperties, "termsOfUse")
		delete(additionalProperties, "privacyPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityProviderPropertiesIdvMetadata struct {
	value *IdentityProviderPropertiesIdvMetadata
	isSet bool
}

func (v NullableIdentityProviderPropertiesIdvMetadata) Get() *IdentityProviderPropertiesIdvMetadata {
	return v.value
}

func (v *NullableIdentityProviderPropertiesIdvMetadata) Set(val *IdentityProviderPropertiesIdvMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderPropertiesIdvMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderPropertiesIdvMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderPropertiesIdvMetadata(val *IdentityProviderPropertiesIdvMetadata) *NullableIdentityProviderPropertiesIdvMetadata {
	return &NullableIdentityProviderPropertiesIdvMetadata{value: val, isSet: true}
}

func (v NullableIdentityProviderPropertiesIdvMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderPropertiesIdvMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
