/*
Okta Admin Management API

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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the WidgetConfigurationHelpLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetConfigurationHelpLinks{}

// WidgetConfigurationHelpLinks struct for WidgetConfigurationHelpLinks
type WidgetConfigurationHelpLinks struct {
	Help           *string                                  `json:"help,omitempty" validate:"regexp=^https?:\\/\\/"`
	ForgotPassword *string                                  `json:"forgotPassword,omitempty" validate:"regexp=^https?:\\/\\/"`
	FactorPage     *WidgetConfigurationHelpLinksFactorPage  `json:"factorPage,omitempty"`
	Unlock         *string                                  `json:"unlock,omitempty" validate:"regexp=^https?:\\/\\/"`
	Custom         []WidgetConfigurationHelpLinksFactorPage `json:"custom,omitempty"`
}

// NewWidgetConfigurationHelpLinks instantiates a new WidgetConfigurationHelpLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetConfigurationHelpLinks() *WidgetConfigurationHelpLinks {
	this := WidgetConfigurationHelpLinks{}
	return &this
}

// NewWidgetConfigurationHelpLinksWithDefaults instantiates a new WidgetConfigurationHelpLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetConfigurationHelpLinksWithDefaults() *WidgetConfigurationHelpLinks {
	this := WidgetConfigurationHelpLinks{}
	return &this
}

// GetHelp returns the Help field value if set, zero value otherwise.
func (o *WidgetConfigurationHelpLinks) GetHelp() string {
	if o == nil || IsNil(o.Help) {
		var ret string
		return ret
	}
	return *o.Help
}

// GetHelpOk returns a tuple with the Help field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationHelpLinks) GetHelpOk() (*string, bool) {
	if o == nil || IsNil(o.Help) {
		return nil, false
	}
	return o.Help, true
}

// HasHelp returns a boolean if a field has been set.
func (o *WidgetConfigurationHelpLinks) HasHelp() bool {
	if o != nil && !IsNil(o.Help) {
		return true
	}

	return false
}

// SetHelp gets a reference to the given string and assigns it to the Help field.
func (o *WidgetConfigurationHelpLinks) SetHelp(v string) {
	o.Help = &v
}

// GetForgotPassword returns the ForgotPassword field value if set, zero value otherwise.
func (o *WidgetConfigurationHelpLinks) GetForgotPassword() string {
	if o == nil || IsNil(o.ForgotPassword) {
		var ret string
		return ret
	}
	return *o.ForgotPassword
}

// GetForgotPasswordOk returns a tuple with the ForgotPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationHelpLinks) GetForgotPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ForgotPassword) {
		return nil, false
	}
	return o.ForgotPassword, true
}

// HasForgotPassword returns a boolean if a field has been set.
func (o *WidgetConfigurationHelpLinks) HasForgotPassword() bool {
	if o != nil && !IsNil(o.ForgotPassword) {
		return true
	}

	return false
}

// SetForgotPassword gets a reference to the given string and assigns it to the ForgotPassword field.
func (o *WidgetConfigurationHelpLinks) SetForgotPassword(v string) {
	o.ForgotPassword = &v
}

// GetFactorPage returns the FactorPage field value if set, zero value otherwise.
func (o *WidgetConfigurationHelpLinks) GetFactorPage() WidgetConfigurationHelpLinksFactorPage {
	if o == nil || IsNil(o.FactorPage) {
		var ret WidgetConfigurationHelpLinksFactorPage
		return ret
	}
	return *o.FactorPage
}

// GetFactorPageOk returns a tuple with the FactorPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationHelpLinks) GetFactorPageOk() (*WidgetConfigurationHelpLinksFactorPage, bool) {
	if o == nil || IsNil(o.FactorPage) {
		return nil, false
	}
	return o.FactorPage, true
}

// HasFactorPage returns a boolean if a field has been set.
func (o *WidgetConfigurationHelpLinks) HasFactorPage() bool {
	if o != nil && !IsNil(o.FactorPage) {
		return true
	}

	return false
}

// SetFactorPage gets a reference to the given WidgetConfigurationHelpLinksFactorPage and assigns it to the FactorPage field.
func (o *WidgetConfigurationHelpLinks) SetFactorPage(v WidgetConfigurationHelpLinksFactorPage) {
	o.FactorPage = &v
}

// GetUnlock returns the Unlock field value if set, zero value otherwise.
func (o *WidgetConfigurationHelpLinks) GetUnlock() string {
	if o == nil || IsNil(o.Unlock) {
		var ret string
		return ret
	}
	return *o.Unlock
}

// GetUnlockOk returns a tuple with the Unlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationHelpLinks) GetUnlockOk() (*string, bool) {
	if o == nil || IsNil(o.Unlock) {
		return nil, false
	}
	return o.Unlock, true
}

// HasUnlock returns a boolean if a field has been set.
func (o *WidgetConfigurationHelpLinks) HasUnlock() bool {
	if o != nil && !IsNil(o.Unlock) {
		return true
	}

	return false
}

// SetUnlock gets a reference to the given string and assigns it to the Unlock field.
func (o *WidgetConfigurationHelpLinks) SetUnlock(v string) {
	o.Unlock = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *WidgetConfigurationHelpLinks) GetCustom() []WidgetConfigurationHelpLinksFactorPage {
	if o == nil || IsNil(o.Custom) {
		var ret []WidgetConfigurationHelpLinksFactorPage
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationHelpLinks) GetCustomOk() ([]WidgetConfigurationHelpLinksFactorPage, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *WidgetConfigurationHelpLinks) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given []WidgetConfigurationHelpLinksFactorPage and assigns it to the Custom field.
func (o *WidgetConfigurationHelpLinks) SetCustom(v []WidgetConfigurationHelpLinksFactorPage) {
	o.Custom = v
}

func (o WidgetConfigurationHelpLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetConfigurationHelpLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Help) {
		toSerialize["help"] = o.Help
	}
	if !IsNil(o.ForgotPassword) {
		toSerialize["forgotPassword"] = o.ForgotPassword
	}
	if !IsNil(o.FactorPage) {
		toSerialize["factorPage"] = o.FactorPage
	}
	if !IsNil(o.Unlock) {
		toSerialize["unlock"] = o.Unlock
	}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	return toSerialize, nil
}

type NullableWidgetConfigurationHelpLinks struct {
	value *WidgetConfigurationHelpLinks
	isSet bool
}

func (v NullableWidgetConfigurationHelpLinks) Get() *WidgetConfigurationHelpLinks {
	return v.value
}

func (v *NullableWidgetConfigurationHelpLinks) Set(val *WidgetConfigurationHelpLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetConfigurationHelpLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetConfigurationHelpLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetConfigurationHelpLinks(val *WidgetConfigurationHelpLinks) *NullableWidgetConfigurationHelpLinks {
	return &NullableWidgetConfigurationHelpLinks{value: val, isSet: true}
}

func (v NullableWidgetConfigurationHelpLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetConfigurationHelpLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
