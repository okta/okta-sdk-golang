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

// ApplicationVisibilityHide Hides the app for specific end-user apps
type ApplicationVisibilityHide struct {
	IOS *bool `json:"iOS,omitempty"`
	Web *bool `json:"web,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationVisibilityHide ApplicationVisibilityHide

// NewApplicationVisibilityHide instantiates a new ApplicationVisibilityHide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationVisibilityHide() *ApplicationVisibilityHide {
	this := ApplicationVisibilityHide{}
	return &this
}

// NewApplicationVisibilityHideWithDefaults instantiates a new ApplicationVisibilityHide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationVisibilityHideWithDefaults() *ApplicationVisibilityHide {
	this := ApplicationVisibilityHide{}
	return &this
}

// GetIOS returns the IOS field value if set, zero value otherwise.
func (o *ApplicationVisibilityHide) GetIOS() bool {
	if o == nil || o.IOS == nil {
		var ret bool
		return ret
	}
	return *o.IOS
}

// GetIOSOk returns a tuple with the IOS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVisibilityHide) GetIOSOk() (*bool, bool) {
	if o == nil || o.IOS == nil {
		return nil, false
	}
	return o.IOS, true
}

// HasIOS returns a boolean if a field has been set.
func (o *ApplicationVisibilityHide) HasIOS() bool {
	if o != nil && o.IOS != nil {
		return true
	}

	return false
}

// SetIOS gets a reference to the given bool and assigns it to the IOS field.
func (o *ApplicationVisibilityHide) SetIOS(v bool) {
	o.IOS = &v
}

// GetWeb returns the Web field value if set, zero value otherwise.
func (o *ApplicationVisibilityHide) GetWeb() bool {
	if o == nil || o.Web == nil {
		var ret bool
		return ret
	}
	return *o.Web
}

// GetWebOk returns a tuple with the Web field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationVisibilityHide) GetWebOk() (*bool, bool) {
	if o == nil || o.Web == nil {
		return nil, false
	}
	return o.Web, true
}

// HasWeb returns a boolean if a field has been set.
func (o *ApplicationVisibilityHide) HasWeb() bool {
	if o != nil && o.Web != nil {
		return true
	}

	return false
}

// SetWeb gets a reference to the given bool and assigns it to the Web field.
func (o *ApplicationVisibilityHide) SetWeb(v bool) {
	o.Web = &v
}

func (o ApplicationVisibilityHide) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IOS != nil {
		toSerialize["iOS"] = o.IOS
	}
	if o.Web != nil {
		toSerialize["web"] = o.Web
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationVisibilityHide) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationVisibilityHide := _ApplicationVisibilityHide{}

	err = json.Unmarshal(bytes, &varApplicationVisibilityHide)
	if err == nil {
		*o = ApplicationVisibilityHide(varApplicationVisibilityHide)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "iOS")
		delete(additionalProperties, "web")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationVisibilityHide struct {
	value *ApplicationVisibilityHide
	isSet bool
}

func (v NullableApplicationVisibilityHide) Get() *ApplicationVisibilityHide {
	return v.value
}

func (v *NullableApplicationVisibilityHide) Set(val *ApplicationVisibilityHide) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationVisibilityHide) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationVisibilityHide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationVisibilityHide(val *ApplicationVisibilityHide) *NullableApplicationVisibilityHide {
	return &NullableApplicationVisibilityHide{value: val, isSet: true}
}

func (v NullableApplicationVisibilityHide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationVisibilityHide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

