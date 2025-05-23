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

// AuthenticatorKeyCustomAppAllOfProviderConfigurationApns struct for AuthenticatorKeyCustomAppAllOfProviderConfigurationApns
type AuthenticatorKeyCustomAppAllOfProviderConfigurationApns struct {
	// ID of the APNs (Apple Push Notification Service) [configurations](https://developer.okta.com/docs/reference/api/push-providers/)
	Id *string `json:"id,omitempty"`
	// AppBundleId of the APNs (Apple Push Notification Service) [configurations](https://developer.okta.com/docs/reference/api/push-providers/)
	AppBundleId *string `json:"appBundleId,omitempty"`
	// DebugAppBundleId of the APNs (Apple Push Notification Service) [configurations](https://developer.okta.com/docs/reference/api/push-providers/)
	DebugAppBundleId *string `json:"debugAppBundleId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorKeyCustomAppAllOfProviderConfigurationApns AuthenticatorKeyCustomAppAllOfProviderConfigurationApns

// NewAuthenticatorKeyCustomAppAllOfProviderConfigurationApns instantiates a new AuthenticatorKeyCustomAppAllOfProviderConfigurationApns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorKeyCustomAppAllOfProviderConfigurationApns() *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns {
	this := AuthenticatorKeyCustomAppAllOfProviderConfigurationApns{}
	return &this
}

// NewAuthenticatorKeyCustomAppAllOfProviderConfigurationApnsWithDefaults instantiates a new AuthenticatorKeyCustomAppAllOfProviderConfigurationApns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorKeyCustomAppAllOfProviderConfigurationApnsWithDefaults() *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns {
	this := AuthenticatorKeyCustomAppAllOfProviderConfigurationApns{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) SetId(v string) {
	o.Id = &v
}

// GetAppBundleId returns the AppBundleId field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) GetAppBundleId() string {
	if o == nil || o.AppBundleId == nil {
		var ret string
		return ret
	}
	return *o.AppBundleId
}

// GetAppBundleIdOk returns a tuple with the AppBundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) GetAppBundleIdOk() (*string, bool) {
	if o == nil || o.AppBundleId == nil {
		return nil, false
	}
	return o.AppBundleId, true
}

// HasAppBundleId returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) HasAppBundleId() bool {
	if o != nil && o.AppBundleId != nil {
		return true
	}

	return false
}

// SetAppBundleId gets a reference to the given string and assigns it to the AppBundleId field.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) SetAppBundleId(v string) {
	o.AppBundleId = &v
}

// GetDebugAppBundleId returns the DebugAppBundleId field value if set, zero value otherwise.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) GetDebugAppBundleId() string {
	if o == nil || o.DebugAppBundleId == nil {
		var ret string
		return ret
	}
	return *o.DebugAppBundleId
}

// GetDebugAppBundleIdOk returns a tuple with the DebugAppBundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) GetDebugAppBundleIdOk() (*string, bool) {
	if o == nil || o.DebugAppBundleId == nil {
		return nil, false
	}
	return o.DebugAppBundleId, true
}

// HasDebugAppBundleId returns a boolean if a field has been set.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) HasDebugAppBundleId() bool {
	if o != nil && o.DebugAppBundleId != nil {
		return true
	}

	return false
}

// SetDebugAppBundleId gets a reference to the given string and assigns it to the DebugAppBundleId field.
func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) SetDebugAppBundleId(v string) {
	o.DebugAppBundleId = &v
}

func (o AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AppBundleId != nil {
		toSerialize["appBundleId"] = o.AppBundleId
	}
	if o.DebugAppBundleId != nil {
		toSerialize["debugAppBundleId"] = o.DebugAppBundleId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) UnmarshalJSON(bytes []byte) (err error) {
	varAuthenticatorKeyCustomAppAllOfProviderConfigurationApns := _AuthenticatorKeyCustomAppAllOfProviderConfigurationApns{}

	err = json.Unmarshal(bytes, &varAuthenticatorKeyCustomAppAllOfProviderConfigurationApns)
	if err == nil {
		*o = AuthenticatorKeyCustomAppAllOfProviderConfigurationApns(varAuthenticatorKeyCustomAppAllOfProviderConfigurationApns)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "appBundleId")
		delete(additionalProperties, "debugAppBundleId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns struct {
	value *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns
	isSet bool
}

func (v NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns) Get() *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns {
	return v.value
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns) Set(val *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns(val *AuthenticatorKeyCustomAppAllOfProviderConfigurationApns) *NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns {
	return &NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns{value: val, isSet: true}
}

func (v NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorKeyCustomAppAllOfProviderConfigurationApns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

