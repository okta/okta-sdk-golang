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

// ApplicationVisibility struct for ApplicationVisibility
type ApplicationVisibility struct {
	// Links or icons that appear on the End-User Dashboard when they're assigned to the app
	AppLinks *map[string]bool `json:"appLinks,omitempty"`
	// Automatically signs in to the app when user signs into Okta
	AutoLaunch *bool `json:"autoLaunch,omitempty"`
	// Automatically sign in when user lands on the sign-in page
	AutoSubmitToolbar *bool `json:"autoSubmitToolbar,omitempty"`
	Hide *ApplicationVisibilityHide `json:"hide,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationVisibility ApplicationVisibility

// NewApplicationVisibility instantiates a new ApplicationVisibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationVisibility() *ApplicationVisibility {
	this := ApplicationVisibility{}
	return &this
}

// NewApplicationVisibilityWithDefaults instantiates a new ApplicationVisibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationVisibilityWithDefaults() *ApplicationVisibility {
	this := ApplicationVisibility{}
	return &this
}

// GetAppLinks returns the AppLinks field value if set, zero value otherwise.
func (o *ApplicationVisibility) GetAppLinks() map[string]bool {
	if o == nil || o.AppLinks == nil {
		var ret map[string]bool
		return ret
	}
	return *o.AppLinks
}

// GetAppLinksOk returns a tuple with the AppLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVisibility) GetAppLinksOk() (*map[string]bool, bool) {
	if o == nil || o.AppLinks == nil {
		return nil, false
	}
	return o.AppLinks, true
}

// HasAppLinks returns a boolean if a field has been set.
func (o *ApplicationVisibility) HasAppLinks() bool {
	if o != nil && o.AppLinks != nil {
		return true
	}

	return false
}

// SetAppLinks gets a reference to the given map[string]bool and assigns it to the AppLinks field.
func (o *ApplicationVisibility) SetAppLinks(v map[string]bool) {
	o.AppLinks = &v
}

// GetAutoLaunch returns the AutoLaunch field value if set, zero value otherwise.
func (o *ApplicationVisibility) GetAutoLaunch() bool {
	if o == nil || o.AutoLaunch == nil {
		var ret bool
		return ret
	}
	return *o.AutoLaunch
}

// GetAutoLaunchOk returns a tuple with the AutoLaunch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVisibility) GetAutoLaunchOk() (*bool, bool) {
	if o == nil || o.AutoLaunch == nil {
		return nil, false
	}
	return o.AutoLaunch, true
}

// HasAutoLaunch returns a boolean if a field has been set.
func (o *ApplicationVisibility) HasAutoLaunch() bool {
	if o != nil && o.AutoLaunch != nil {
		return true
	}

	return false
}

// SetAutoLaunch gets a reference to the given bool and assigns it to the AutoLaunch field.
func (o *ApplicationVisibility) SetAutoLaunch(v bool) {
	o.AutoLaunch = &v
}

// GetAutoSubmitToolbar returns the AutoSubmitToolbar field value if set, zero value otherwise.
func (o *ApplicationVisibility) GetAutoSubmitToolbar() bool {
	if o == nil || o.AutoSubmitToolbar == nil {
		var ret bool
		return ret
	}
	return *o.AutoSubmitToolbar
}

// GetAutoSubmitToolbarOk returns a tuple with the AutoSubmitToolbar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVisibility) GetAutoSubmitToolbarOk() (*bool, bool) {
	if o == nil || o.AutoSubmitToolbar == nil {
		return nil, false
	}
	return o.AutoSubmitToolbar, true
}

// HasAutoSubmitToolbar returns a boolean if a field has been set.
func (o *ApplicationVisibility) HasAutoSubmitToolbar() bool {
	if o != nil && o.AutoSubmitToolbar != nil {
		return true
	}

	return false
}

// SetAutoSubmitToolbar gets a reference to the given bool and assigns it to the AutoSubmitToolbar field.
func (o *ApplicationVisibility) SetAutoSubmitToolbar(v bool) {
	o.AutoSubmitToolbar = &v
}

// GetHide returns the Hide field value if set, zero value otherwise.
func (o *ApplicationVisibility) GetHide() ApplicationVisibilityHide {
	if o == nil || o.Hide == nil {
		var ret ApplicationVisibilityHide
		return ret
	}
	return *o.Hide
}

// GetHideOk returns a tuple with the Hide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVisibility) GetHideOk() (*ApplicationVisibilityHide, bool) {
	if o == nil || o.Hide == nil {
		return nil, false
	}
	return o.Hide, true
}

// HasHide returns a boolean if a field has been set.
func (o *ApplicationVisibility) HasHide() bool {
	if o != nil && o.Hide != nil {
		return true
	}

	return false
}

// SetHide gets a reference to the given ApplicationVisibilityHide and assigns it to the Hide field.
func (o *ApplicationVisibility) SetHide(v ApplicationVisibilityHide) {
	o.Hide = &v
}

func (o ApplicationVisibility) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppLinks != nil {
		toSerialize["appLinks"] = o.AppLinks
	}
	if o.AutoLaunch != nil {
		toSerialize["autoLaunch"] = o.AutoLaunch
	}
	if o.AutoSubmitToolbar != nil {
		toSerialize["autoSubmitToolbar"] = o.AutoSubmitToolbar
	}
	if o.Hide != nil {
		toSerialize["hide"] = o.Hide
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationVisibility) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationVisibility := _ApplicationVisibility{}

	err = json.Unmarshal(bytes, &varApplicationVisibility)
	if err == nil {
		*o = ApplicationVisibility(varApplicationVisibility)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appLinks")
		delete(additionalProperties, "autoLaunch")
		delete(additionalProperties, "autoSubmitToolbar")
		delete(additionalProperties, "hide")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationVisibility struct {
	value *ApplicationVisibility
	isSet bool
}

func (v NullableApplicationVisibility) Get() *ApplicationVisibility {
	return v.value
}

func (v *NullableApplicationVisibility) Set(val *ApplicationVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationVisibility(val *ApplicationVisibility) *NullableApplicationVisibility {
	return &NullableApplicationVisibility{value: val, isSet: true}
}

func (v NullableApplicationVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

