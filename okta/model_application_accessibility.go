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

// ApplicationAccessibility Specifies access settings for the app
type ApplicationAccessibility struct {
	// Custom error page URL for the app
	ErrorRedirectUrl *string `json:"errorRedirectUrl,omitempty"`
	// Custom login page URL for the app
	LoginRedirectUrl *string `json:"loginRedirectUrl,omitempty"`
	// Represents whether the app can be self-assignable by users
	SelfService *bool `json:"selfService,omitempty"`
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
	if o == nil || o.ErrorRedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.ErrorRedirectUrl
}

// GetErrorRedirectUrlOk returns a tuple with the ErrorRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAccessibility) GetErrorRedirectUrlOk() (*string, bool) {
	if o == nil || o.ErrorRedirectUrl == nil {
		return nil, false
	}
	return o.ErrorRedirectUrl, true
}

// HasErrorRedirectUrl returns a boolean if a field has been set.
func (o *ApplicationAccessibility) HasErrorRedirectUrl() bool {
	if o != nil && o.ErrorRedirectUrl != nil {
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
	if o == nil || o.LoginRedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.LoginRedirectUrl
}

// GetLoginRedirectUrlOk returns a tuple with the LoginRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAccessibility) GetLoginRedirectUrlOk() (*string, bool) {
	if o == nil || o.LoginRedirectUrl == nil {
		return nil, false
	}
	return o.LoginRedirectUrl, true
}

// HasLoginRedirectUrl returns a boolean if a field has been set.
func (o *ApplicationAccessibility) HasLoginRedirectUrl() bool {
	if o != nil && o.LoginRedirectUrl != nil {
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
	if o == nil || o.SelfService == nil {
		var ret bool
		return ret
	}
	return *o.SelfService
}

// GetSelfServiceOk returns a tuple with the SelfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAccessibility) GetSelfServiceOk() (*bool, bool) {
	if o == nil || o.SelfService == nil {
		return nil, false
	}
	return o.SelfService, true
}

// HasSelfService returns a boolean if a field has been set.
func (o *ApplicationAccessibility) HasSelfService() bool {
	if o != nil && o.SelfService != nil {
		return true
	}

	return false
}

// SetSelfService gets a reference to the given bool and assigns it to the SelfService field.
func (o *ApplicationAccessibility) SetSelfService(v bool) {
	o.SelfService = &v
}

func (o ApplicationAccessibility) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorRedirectUrl != nil {
		toSerialize["errorRedirectUrl"] = o.ErrorRedirectUrl
	}
	if o.LoginRedirectUrl != nil {
		toSerialize["loginRedirectUrl"] = o.LoginRedirectUrl
	}
	if o.SelfService != nil {
		toSerialize["selfService"] = o.SelfService
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationAccessibility) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationAccessibility := _ApplicationAccessibility{}

	err = json.Unmarshal(bytes, &varApplicationAccessibility)
	if err == nil {
		*o = ApplicationAccessibility(varApplicationAccessibility)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "errorRedirectUrl")
		delete(additionalProperties, "loginRedirectUrl")
		delete(additionalProperties, "selfService")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

