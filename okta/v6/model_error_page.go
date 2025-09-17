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

// checks if the ErrorPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorPage{}

// ErrorPage struct for ErrorPage
type ErrorPage struct {
	// The HTML for the page
	PageContent                  *string                       `json:"pageContent,omitempty"`
	ContentSecurityPolicySetting *ContentSecurityPolicySetting `json:"contentSecurityPolicySetting,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _ErrorPage ErrorPage

// NewErrorPage instantiates a new ErrorPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorPage() *ErrorPage {
	this := ErrorPage{}
	return &this
}

// NewErrorPageWithDefaults instantiates a new ErrorPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorPageWithDefaults() *ErrorPage {
	this := ErrorPage{}
	return &this
}

// GetPageContent returns the PageContent field value if set, zero value otherwise.
func (o *ErrorPage) GetPageContent() string {
	if o == nil || IsNil(o.PageContent) {
		var ret string
		return ret
	}
	return *o.PageContent
}

// GetPageContentOk returns a tuple with the PageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorPage) GetPageContentOk() (*string, bool) {
	if o == nil || IsNil(o.PageContent) {
		return nil, false
	}
	return o.PageContent, true
}

// HasPageContent returns a boolean if a field has been set.
func (o *ErrorPage) HasPageContent() bool {
	if o != nil && !IsNil(o.PageContent) {
		return true
	}

	return false
}

// SetPageContent gets a reference to the given string and assigns it to the PageContent field.
func (o *ErrorPage) SetPageContent(v string) {
	o.PageContent = &v
}

// GetContentSecurityPolicySetting returns the ContentSecurityPolicySetting field value if set, zero value otherwise.
func (o *ErrorPage) GetContentSecurityPolicySetting() ContentSecurityPolicySetting {
	if o == nil || IsNil(o.ContentSecurityPolicySetting) {
		var ret ContentSecurityPolicySetting
		return ret
	}
	return *o.ContentSecurityPolicySetting
}

// GetContentSecurityPolicySettingOk returns a tuple with the ContentSecurityPolicySetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorPage) GetContentSecurityPolicySettingOk() (*ContentSecurityPolicySetting, bool) {
	if o == nil || IsNil(o.ContentSecurityPolicySetting) {
		return nil, false
	}
	return o.ContentSecurityPolicySetting, true
}

// HasContentSecurityPolicySetting returns a boolean if a field has been set.
func (o *ErrorPage) HasContentSecurityPolicySetting() bool {
	if o != nil && !IsNil(o.ContentSecurityPolicySetting) {
		return true
	}

	return false
}

// SetContentSecurityPolicySetting gets a reference to the given ContentSecurityPolicySetting and assigns it to the ContentSecurityPolicySetting field.
func (o *ErrorPage) SetContentSecurityPolicySetting(v ContentSecurityPolicySetting) {
	o.ContentSecurityPolicySetting = &v
}

func (o ErrorPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageContent) {
		toSerialize["pageContent"] = o.PageContent
	}
	if !IsNil(o.ContentSecurityPolicySetting) {
		toSerialize["contentSecurityPolicySetting"] = o.ContentSecurityPolicySetting
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ErrorPage) UnmarshalJSON(data []byte) (err error) {
	varErrorPage := _ErrorPage{}

	err = json.Unmarshal(data, &varErrorPage)

	if err != nil {
		return err
	}

	*o = ErrorPage(varErrorPage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pageContent")
		delete(additionalProperties, "contentSecurityPolicySetting")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableErrorPage struct {
	value *ErrorPage
	isSet bool
}

func (v NullableErrorPage) Get() *ErrorPage {
	return v.value
}

func (v *NullableErrorPage) Set(val *ErrorPage) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorPage) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorPage(val *ErrorPage) *NullableErrorPage {
	return &NullableErrorPage{value: val, isSet: true}
}

func (v NullableErrorPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
