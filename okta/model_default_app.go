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

// DefaultApp struct for DefaultApp
type DefaultApp struct {
	// ID for the App instance
	AppInstanceId *string `json:"appInstanceId,omitempty"`
	// Name for the app instance
	AppLinkName *string `json:"appLinkName,omitempty"`
	// Application URI for classic Orgs
	ClassicApplicationUri *string `json:"classicApplicationUri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DefaultApp DefaultApp

// NewDefaultApp instantiates a new DefaultApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultApp() *DefaultApp {
	this := DefaultApp{}
	return &this
}

// NewDefaultAppWithDefaults instantiates a new DefaultApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultAppWithDefaults() *DefaultApp {
	this := DefaultApp{}
	return &this
}

// GetAppInstanceId returns the AppInstanceId field value if set, zero value otherwise.
func (o *DefaultApp) GetAppInstanceId() string {
	if o == nil || o.AppInstanceId == nil {
		var ret string
		return ret
	}
	return *o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultApp) GetAppInstanceIdOk() (*string, bool) {
	if o == nil || o.AppInstanceId == nil {
		return nil, false
	}
	return o.AppInstanceId, true
}

// HasAppInstanceId returns a boolean if a field has been set.
func (o *DefaultApp) HasAppInstanceId() bool {
	if o != nil && o.AppInstanceId != nil {
		return true
	}

	return false
}

// SetAppInstanceId gets a reference to the given string and assigns it to the AppInstanceId field.
func (o *DefaultApp) SetAppInstanceId(v string) {
	o.AppInstanceId = &v
}

// GetAppLinkName returns the AppLinkName field value if set, zero value otherwise.
func (o *DefaultApp) GetAppLinkName() string {
	if o == nil || o.AppLinkName == nil {
		var ret string
		return ret
	}
	return *o.AppLinkName
}

// GetAppLinkNameOk returns a tuple with the AppLinkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultApp) GetAppLinkNameOk() (*string, bool) {
	if o == nil || o.AppLinkName == nil {
		return nil, false
	}
	return o.AppLinkName, true
}

// HasAppLinkName returns a boolean if a field has been set.
func (o *DefaultApp) HasAppLinkName() bool {
	if o != nil && o.AppLinkName != nil {
		return true
	}

	return false
}

// SetAppLinkName gets a reference to the given string and assigns it to the AppLinkName field.
func (o *DefaultApp) SetAppLinkName(v string) {
	o.AppLinkName = &v
}

// GetClassicApplicationUri returns the ClassicApplicationUri field value if set, zero value otherwise.
func (o *DefaultApp) GetClassicApplicationUri() string {
	if o == nil || o.ClassicApplicationUri == nil {
		var ret string
		return ret
	}
	return *o.ClassicApplicationUri
}

// GetClassicApplicationUriOk returns a tuple with the ClassicApplicationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultApp) GetClassicApplicationUriOk() (*string, bool) {
	if o == nil || o.ClassicApplicationUri == nil {
		return nil, false
	}
	return o.ClassicApplicationUri, true
}

// HasClassicApplicationUri returns a boolean if a field has been set.
func (o *DefaultApp) HasClassicApplicationUri() bool {
	if o != nil && o.ClassicApplicationUri != nil {
		return true
	}

	return false
}

// SetClassicApplicationUri gets a reference to the given string and assigns it to the ClassicApplicationUri field.
func (o *DefaultApp) SetClassicApplicationUri(v string) {
	o.ClassicApplicationUri = &v
}

func (o DefaultApp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppInstanceId != nil {
		toSerialize["appInstanceId"] = o.AppInstanceId
	}
	if o.AppLinkName != nil {
		toSerialize["appLinkName"] = o.AppLinkName
	}
	if o.ClassicApplicationUri != nil {
		toSerialize["classicApplicationUri"] = o.ClassicApplicationUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DefaultApp) UnmarshalJSON(bytes []byte) (err error) {
	varDefaultApp := _DefaultApp{}

	err = json.Unmarshal(bytes, &varDefaultApp)
	if err == nil {
		*o = DefaultApp(varDefaultApp)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appInstanceId")
		delete(additionalProperties, "appLinkName")
		delete(additionalProperties, "classicApplicationUri")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDefaultApp struct {
	value *DefaultApp
	isSet bool
}

func (v NullableDefaultApp) Get() *DefaultApp {
	return v.value
}

func (v *NullableDefaultApp) Set(val *DefaultApp) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultApp) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultApp(val *DefaultApp) *NullableDefaultApp {
	return &NullableDefaultApp{value: val, isSet: true}
}

func (v NullableDefaultApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

