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

// TestInfoOidcTestConfiguration OIDC test details
type TestInfoOidcTestConfiguration struct {
	// Read only.<br>Indicates if your integration supports IdP-initiated sign-in flows. If [`sso.oidc.initiateLoginUri`](/openapi/okta-management/management/tag/YourOinIntegrations/#tag/YourOinIntegrations/operation/createSubmission!path=sso/oidc/initiateLoginUri&t=request) is specified, this property is set to `true`. If [`sso.oidc.initiateLoginUri`](/openapi/okta-management/management/tag/YourOinIntegrations/#tag/YourOinIntegrations/operation/createSubmission!path=sso/oidc/initiateLoginUri&t=request) isn't set for the integration submission, this property is set to `false`
	Idp *bool `json:"idp,omitempty"`
	// Read only.<br>Indicates if your integration supports SP-initiated sign-in flows and is always set to `true` for OIDC SSO
	Sp *bool `json:"sp,omitempty"`
	// Indicates if your integration supports Just-In-Time (JIT) provisioning
	Jit *bool `json:"jit,omitempty"`
	// URL for SP-initiated sign-in flows (required if `sp = true`)
	SpInitiateUrl string `json:"spInitiateUrl"`
	AdditionalProperties map[string]interface{}
}

type _TestInfoOidcTestConfiguration TestInfoOidcTestConfiguration

// NewTestInfoOidcTestConfiguration instantiates a new TestInfoOidcTestConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestInfoOidcTestConfiguration(spInitiateUrl string) *TestInfoOidcTestConfiguration {
	this := TestInfoOidcTestConfiguration{}
	this.SpInitiateUrl = spInitiateUrl
	return &this
}

// NewTestInfoOidcTestConfigurationWithDefaults instantiates a new TestInfoOidcTestConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestInfoOidcTestConfigurationWithDefaults() *TestInfoOidcTestConfiguration {
	this := TestInfoOidcTestConfiguration{}
	return &this
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *TestInfoOidcTestConfiguration) GetIdp() bool {
	if o == nil || o.Idp == nil {
		var ret bool
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfoOidcTestConfiguration) GetIdpOk() (*bool, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *TestInfoOidcTestConfiguration) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given bool and assigns it to the Idp field.
func (o *TestInfoOidcTestConfiguration) SetIdp(v bool) {
	o.Idp = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *TestInfoOidcTestConfiguration) GetSp() bool {
	if o == nil || o.Sp == nil {
		var ret bool
		return ret
	}
	return *o.Sp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfoOidcTestConfiguration) GetSpOk() (*bool, bool) {
	if o == nil || o.Sp == nil {
		return nil, false
	}
	return o.Sp, true
}

// HasSp returns a boolean if a field has been set.
func (o *TestInfoOidcTestConfiguration) HasSp() bool {
	if o != nil && o.Sp != nil {
		return true
	}

	return false
}

// SetSp gets a reference to the given bool and assigns it to the Sp field.
func (o *TestInfoOidcTestConfiguration) SetSp(v bool) {
	o.Sp = &v
}

// GetJit returns the Jit field value if set, zero value otherwise.
func (o *TestInfoOidcTestConfiguration) GetJit() bool {
	if o == nil || o.Jit == nil {
		var ret bool
		return ret
	}
	return *o.Jit
}

// GetJitOk returns a tuple with the Jit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestInfoOidcTestConfiguration) GetJitOk() (*bool, bool) {
	if o == nil || o.Jit == nil {
		return nil, false
	}
	return o.Jit, true
}

// HasJit returns a boolean if a field has been set.
func (o *TestInfoOidcTestConfiguration) HasJit() bool {
	if o != nil && o.Jit != nil {
		return true
	}

	return false
}

// SetJit gets a reference to the given bool and assigns it to the Jit field.
func (o *TestInfoOidcTestConfiguration) SetJit(v bool) {
	o.Jit = &v
}

// GetSpInitiateUrl returns the SpInitiateUrl field value
func (o *TestInfoOidcTestConfiguration) GetSpInitiateUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpInitiateUrl
}

// GetSpInitiateUrlOk returns a tuple with the SpInitiateUrl field value
// and a boolean to check if the value has been set.
func (o *TestInfoOidcTestConfiguration) GetSpInitiateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpInitiateUrl, true
}

// SetSpInitiateUrl sets field value
func (o *TestInfoOidcTestConfiguration) SetSpInitiateUrl(v string) {
	o.SpInitiateUrl = v
}

func (o TestInfoOidcTestConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Idp != nil {
		toSerialize["idp"] = o.Idp
	}
	if o.Sp != nil {
		toSerialize["sp"] = o.Sp
	}
	if o.Jit != nil {
		toSerialize["jit"] = o.Jit
	}
	if true {
		toSerialize["spInitiateUrl"] = o.SpInitiateUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TestInfoOidcTestConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varTestInfoOidcTestConfiguration := _TestInfoOidcTestConfiguration{}

	err = json.Unmarshal(bytes, &varTestInfoOidcTestConfiguration)
	if err == nil {
		*o = TestInfoOidcTestConfiguration(varTestInfoOidcTestConfiguration)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "idp")
		delete(additionalProperties, "sp")
		delete(additionalProperties, "jit")
		delete(additionalProperties, "spInitiateUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTestInfoOidcTestConfiguration struct {
	value *TestInfoOidcTestConfiguration
	isSet bool
}

func (v NullableTestInfoOidcTestConfiguration) Get() *TestInfoOidcTestConfiguration {
	return v.value
}

func (v *NullableTestInfoOidcTestConfiguration) Set(val *TestInfoOidcTestConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInfoOidcTestConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInfoOidcTestConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInfoOidcTestConfiguration(val *TestInfoOidcTestConfiguration) *NullableTestInfoOidcTestConfiguration {
	return &NullableTestInfoOidcTestConfiguration{value: val, isSet: true}
}

func (v NullableTestInfoOidcTestConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInfoOidcTestConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

