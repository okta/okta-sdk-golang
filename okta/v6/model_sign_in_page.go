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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SignInPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignInPage{}

// SignInPage struct for SignInPage
type SignInPage struct {
	// The HTML for the page
	PageContent                  *string                              `json:"pageContent,omitempty"`
	ContentSecurityPolicySetting *ContentSecurityPolicySetting        `json:"contentSecurityPolicySetting,omitempty"`
	WidgetCustomizations         *SignInPageAllOfWidgetCustomizations `json:"widgetCustomizations,omitempty"`
	// The version specified as a [Semantic Version](https://semver.org/).
	WidgetVersion        *string `json:"widgetVersion,omitempty" validate:"regexp=^(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)\\\\.(0|[1-9]\\\\d*)(?:-((?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\\\.(?:0|[1-9]\\\\d*|\\\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\\\+([0-9a-zA-Z-]+(?:\\\\.[0-9a-zA-Z-]+)*))?$"`
	AdditionalProperties map[string]interface{}
}

type _SignInPage SignInPage

// NewSignInPage instantiates a new SignInPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignInPage() *SignInPage {
	this := SignInPage{}
	return &this
}

// NewSignInPageWithDefaults instantiates a new SignInPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignInPageWithDefaults() *SignInPage {
	this := SignInPage{}
	return &this
}

// GetPageContent returns the PageContent field value if set, zero value otherwise.
func (o *SignInPage) GetPageContent() string {
	if o == nil || IsNil(o.PageContent) {
		var ret string
		return ret
	}
	return *o.PageContent
}

// GetPageContentOk returns a tuple with the PageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPage) GetPageContentOk() (*string, bool) {
	if o == nil || IsNil(o.PageContent) {
		return nil, false
	}
	return o.PageContent, true
}

// HasPageContent returns a boolean if a field has been set.
func (o *SignInPage) HasPageContent() bool {
	if o != nil && !IsNil(o.PageContent) {
		return true
	}

	return false
}

// SetPageContent gets a reference to the given string and assigns it to the PageContent field.
func (o *SignInPage) SetPageContent(v string) {
	o.PageContent = &v
}

// GetContentSecurityPolicySetting returns the ContentSecurityPolicySetting field value if set, zero value otherwise.
func (o *SignInPage) GetContentSecurityPolicySetting() ContentSecurityPolicySetting {
	if o == nil || IsNil(o.ContentSecurityPolicySetting) {
		var ret ContentSecurityPolicySetting
		return ret
	}
	return *o.ContentSecurityPolicySetting
}

// GetContentSecurityPolicySettingOk returns a tuple with the ContentSecurityPolicySetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPage) GetContentSecurityPolicySettingOk() (*ContentSecurityPolicySetting, bool) {
	if o == nil || IsNil(o.ContentSecurityPolicySetting) {
		return nil, false
	}
	return o.ContentSecurityPolicySetting, true
}

// HasContentSecurityPolicySetting returns a boolean if a field has been set.
func (o *SignInPage) HasContentSecurityPolicySetting() bool {
	if o != nil && !IsNil(o.ContentSecurityPolicySetting) {
		return true
	}

	return false
}

// SetContentSecurityPolicySetting gets a reference to the given ContentSecurityPolicySetting and assigns it to the ContentSecurityPolicySetting field.
func (o *SignInPage) SetContentSecurityPolicySetting(v ContentSecurityPolicySetting) {
	o.ContentSecurityPolicySetting = &v
}

// GetWidgetCustomizations returns the WidgetCustomizations field value if set, zero value otherwise.
func (o *SignInPage) GetWidgetCustomizations() SignInPageAllOfWidgetCustomizations {
	if o == nil || IsNil(o.WidgetCustomizations) {
		var ret SignInPageAllOfWidgetCustomizations
		return ret
	}
	return *o.WidgetCustomizations
}

// GetWidgetCustomizationsOk returns a tuple with the WidgetCustomizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPage) GetWidgetCustomizationsOk() (*SignInPageAllOfWidgetCustomizations, bool) {
	if o == nil || IsNil(o.WidgetCustomizations) {
		return nil, false
	}
	return o.WidgetCustomizations, true
}

// HasWidgetCustomizations returns a boolean if a field has been set.
func (o *SignInPage) HasWidgetCustomizations() bool {
	if o != nil && !IsNil(o.WidgetCustomizations) {
		return true
	}

	return false
}

// SetWidgetCustomizations gets a reference to the given SignInPageAllOfWidgetCustomizations and assigns it to the WidgetCustomizations field.
func (o *SignInPage) SetWidgetCustomizations(v SignInPageAllOfWidgetCustomizations) {
	o.WidgetCustomizations = &v
}

// GetWidgetVersion returns the WidgetVersion field value if set, zero value otherwise.
func (o *SignInPage) GetWidgetVersion() string {
	if o == nil || IsNil(o.WidgetVersion) {
		var ret string
		return ret
	}
	return *o.WidgetVersion
}

// GetWidgetVersionOk returns a tuple with the WidgetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPage) GetWidgetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.WidgetVersion) {
		return nil, false
	}
	return o.WidgetVersion, true
}

// HasWidgetVersion returns a boolean if a field has been set.
func (o *SignInPage) HasWidgetVersion() bool {
	if o != nil && !IsNil(o.WidgetVersion) {
		return true
	}

	return false
}

// SetWidgetVersion gets a reference to the given string and assigns it to the WidgetVersion field.
func (o *SignInPage) SetWidgetVersion(v string) {
	o.WidgetVersion = &v
}

func (o SignInPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignInPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageContent) {
		toSerialize["pageContent"] = o.PageContent
	}
	if !IsNil(o.ContentSecurityPolicySetting) {
		toSerialize["contentSecurityPolicySetting"] = o.ContentSecurityPolicySetting
	}
	if !IsNil(o.WidgetCustomizations) {
		toSerialize["widgetCustomizations"] = o.WidgetCustomizations
	}
	if !IsNil(o.WidgetVersion) {
		toSerialize["widgetVersion"] = o.WidgetVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignInPage) UnmarshalJSON(data []byte) (err error) {
	varSignInPage := _SignInPage{}

	err = json.Unmarshal(data, &varSignInPage)

	if err != nil {
		return err
	}

	*o = SignInPage(varSignInPage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pageContent")
		delete(additionalProperties, "contentSecurityPolicySetting")
		delete(additionalProperties, "widgetCustomizations")
		delete(additionalProperties, "widgetVersion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignInPage struct {
	value *SignInPage
	isSet bool
}

func (v NullableSignInPage) Get() *SignInPage {
	return v.value
}

func (v *NullableSignInPage) Set(val *SignInPage) {
	v.value = val
	v.isSet = true
}

func (v NullableSignInPage) IsSet() bool {
	return v.isSet
}

func (v *NullableSignInPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignInPage(val *SignInPage) *NullableSignInPage {
	return &NullableSignInPage{value: val, isSet: true}
}

func (v NullableSignInPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignInPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
