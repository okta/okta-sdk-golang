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

// checks if the ApplicationAccessibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationAccessibility{}

// ApplicationAccessibility Specifies access settings for the app
type ApplicationAccessibility struct {
	// Custom error page URL for the app
	ErrorRedirectUrl *string `json:"errorRedirectUrl,omitempty"`
	// Custom login page URL for the app > **Note:** The `loginRedirectUrl` property is deprecated in Identity Engine. This property is used with the custom app login feature. Orgs that actively use this feature can continue to do so. See [Okta-hosted sign-in (redirect authentication)](https://developer.okta.com/docs/guides/redirect-authentication/) or [configure IdP routing rules](https://help.okta.com/okta_help.htm?type=oie&id=ext-cfg-routing-rules) to redirect users to the appropriate sign-in app for orgs that don't use the custom app login feature.
	LoginRedirectUrl *string `json:"loginRedirectUrl,omitempty"`
	// Represents whether the app can be self-assignable by users
	SelfService          *bool `json:"selfService,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationAccessibility ApplicationAccessibility

// NewApplicationAccessibility instantiates a new ApplicationAccessibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationAccessibility() *ApplicationAccessibility {
	this := ApplicationAccessibility{}
	return &this
}

// NewApplicationAccessibilityWithDefaults instantiates a new ApplicationAccessibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationAccessibilityWithDefaults() *ApplicationAccessibility {
	this := ApplicationAccessibility{}
	return &this
}

// GetErrorRedirectUrl returns the ErrorRedirectUrl field value if set, zero value otherwise.
func (o *ApplicationAccessibility) GetErrorRedirectUrl() string {
	if o == nil || IsNil(o.ErrorRedirectUrl) {
		var ret string
		return ret
	}
	return *o.ErrorRedirectUrl
}

// GetErrorRedirectUrlOk returns a tuple with the ErrorRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAccessibility) GetErrorRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorRedirectUrl) {
		return nil, false
	}
	return o.ErrorRedirectUrl, true
}

// HasErrorRedirectUrl returns a boolean if a field has been set.
func (o *ApplicationAccessibility) HasErrorRedirectUrl() bool {
	if o != nil && !IsNil(o.ErrorRedirectUrl) {
		return true
	}

	return false
}

// SetErrorRedirectUrl gets a reference to the given string and assigns it to the ErrorRedirectUrl field.
func (o *ApplicationAccessibility) SetErrorRedirectUrl(v string) {
	o.ErrorRedirectUrl = &v
}

// GetLoginRedirectUrl returns the LoginRedirectUrl field value if set, zero value otherwise.
func (o *ApplicationAccessibility) GetLoginRedirectUrl() string {
	if o == nil || IsNil(o.LoginRedirectUrl) {
		var ret string
		return ret
	}
	return *o.LoginRedirectUrl
}

// GetLoginRedirectUrlOk returns a tuple with the LoginRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAccessibility) GetLoginRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LoginRedirectUrl) {
		return nil, false
	}
	return o.LoginRedirectUrl, true
}

// HasLoginRedirectUrl returns a boolean if a field has been set.
func (o *ApplicationAccessibility) HasLoginRedirectUrl() bool {
	if o != nil && !IsNil(o.LoginRedirectUrl) {
		return true
	}

	return false
}

// SetLoginRedirectUrl gets a reference to the given string and assigns it to the LoginRedirectUrl field.
func (o *ApplicationAccessibility) SetLoginRedirectUrl(v string) {
	o.LoginRedirectUrl = &v
}

// GetSelfService returns the SelfService field value if set, zero value otherwise.
func (o *ApplicationAccessibility) GetSelfService() bool {
	if o == nil || IsNil(o.SelfService) {
		var ret bool
		return ret
	}
	return *o.SelfService
}

// GetSelfServiceOk returns a tuple with the SelfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAccessibility) GetSelfServiceOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfService) {
		return nil, false
	}
	return o.SelfService, true
}

// HasSelfService returns a boolean if a field has been set.
func (o *ApplicationAccessibility) HasSelfService() bool {
	if o != nil && !IsNil(o.SelfService) {
		return true
	}

	return false
}

// SetSelfService gets a reference to the given bool and assigns it to the SelfService field.
func (o *ApplicationAccessibility) SetSelfService(v bool) {
	o.SelfService = &v
}

func (o ApplicationAccessibility) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationAccessibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorRedirectUrl) {
		toSerialize["errorRedirectUrl"] = o.ErrorRedirectUrl
	}
	if !IsNil(o.LoginRedirectUrl) {
		toSerialize["loginRedirectUrl"] = o.LoginRedirectUrl
	}
	if !IsNil(o.SelfService) {
		toSerialize["selfService"] = o.SelfService
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationAccessibility) UnmarshalJSON(data []byte) (err error) {
	varApplicationAccessibility := _ApplicationAccessibility{}

	err = json.Unmarshal(data, &varApplicationAccessibility)

	if err != nil {
		return err
	}

	*o = ApplicationAccessibility(varApplicationAccessibility)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorRedirectUrl")
		delete(additionalProperties, "loginRedirectUrl")
		delete(additionalProperties, "selfService")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationAccessibility struct {
	value *ApplicationAccessibility
	isSet bool
}

func (v NullableApplicationAccessibility) Get() *ApplicationAccessibility {
	return v.value
}

func (v *NullableApplicationAccessibility) Set(val *ApplicationAccessibility) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationAccessibility) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationAccessibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationAccessibility(val *ApplicationAccessibility) *NullableApplicationAccessibility {
	return &NullableApplicationAccessibility{value: val, isSet: true}
}

func (v NullableApplicationAccessibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationAccessibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
