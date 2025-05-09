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

// AppAccountContainerDetails Container details for resource type APP_ACCOUNT
type AppAccountContainerDetails struct {
	// The application name
	AppName *string `json:"appName,omitempty"`
	// The application ID associated with the privileged account
	ContainerId string `json:"containerId"`
	// Human-readable name of the container that owns the privileged resource
	DisplayName *string `json:"displayName,omitempty"`
	// The application global ID
	GlobalAppId *string `json:"globalAppId,omitempty"`
	// Indicates if the application supports password push
	PasswordPushSupported *bool `json:"passwordPushSupported,omitempty"`
	// Indicates if provisioning is enabled for this application
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	Links *AppLink `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppAccountContainerDetails AppAccountContainerDetails

// NewAppAccountContainerDetails instantiates a new AppAccountContainerDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAccountContainerDetails(containerId string) *AppAccountContainerDetails {
	this := AppAccountContainerDetails{}
	this.ContainerId = containerId
	return &this
}

// NewAppAccountContainerDetailsWithDefaults instantiates a new AppAccountContainerDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAccountContainerDetailsWithDefaults() *AppAccountContainerDetails {
	this := AppAccountContainerDetails{}
	return &this
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *AppAccountContainerDetails) GetAppName() string {
	if o == nil || o.AppName == nil {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetAppNameOk() (*string, bool) {
	if o == nil || o.AppName == nil {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasAppName() bool {
	if o != nil && o.AppName != nil {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *AppAccountContainerDetails) SetAppName(v string) {
	o.AppName = &v
}

// GetContainerId returns the ContainerId field value
func (o *AppAccountContainerDetails) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *AppAccountContainerDetails) SetContainerId(v string) {
	o.ContainerId = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AppAccountContainerDetails) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AppAccountContainerDetails) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetGlobalAppId returns the GlobalAppId field value if set, zero value otherwise.
func (o *AppAccountContainerDetails) GetGlobalAppId() string {
	if o == nil || o.GlobalAppId == nil {
		var ret string
		return ret
	}
	return *o.GlobalAppId
}

// GetGlobalAppIdOk returns a tuple with the GlobalAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetGlobalAppIdOk() (*string, bool) {
	if o == nil || o.GlobalAppId == nil {
		return nil, false
	}
	return o.GlobalAppId, true
}

// HasGlobalAppId returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasGlobalAppId() bool {
	if o != nil && o.GlobalAppId != nil {
		return true
	}

	return false
}

// SetGlobalAppId gets a reference to the given string and assigns it to the GlobalAppId field.
func (o *AppAccountContainerDetails) SetGlobalAppId(v string) {
	o.GlobalAppId = &v
}

// GetPasswordPushSupported returns the PasswordPushSupported field value if set, zero value otherwise.
func (o *AppAccountContainerDetails) GetPasswordPushSupported() bool {
	if o == nil || o.PasswordPushSupported == nil {
		var ret bool
		return ret
	}
	return *o.PasswordPushSupported
}

// GetPasswordPushSupportedOk returns a tuple with the PasswordPushSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetPasswordPushSupportedOk() (*bool, bool) {
	if o == nil || o.PasswordPushSupported == nil {
		return nil, false
	}
	return o.PasswordPushSupported, true
}

// HasPasswordPushSupported returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasPasswordPushSupported() bool {
	if o != nil && o.PasswordPushSupported != nil {
		return true
	}

	return false
}

// SetPasswordPushSupported gets a reference to the given bool and assigns it to the PasswordPushSupported field.
func (o *AppAccountContainerDetails) SetPasswordPushSupported(v bool) {
	o.PasswordPushSupported = &v
}

// GetProvisioningEnabled returns the ProvisioningEnabled field value if set, zero value otherwise.
func (o *AppAccountContainerDetails) GetProvisioningEnabled() bool {
	if o == nil || o.ProvisioningEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ProvisioningEnabled
}

// GetProvisioningEnabledOk returns a tuple with the ProvisioningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetProvisioningEnabledOk() (*bool, bool) {
	if o == nil || o.ProvisioningEnabled == nil {
		return nil, false
	}
	return o.ProvisioningEnabled, true
}

// HasProvisioningEnabled returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasProvisioningEnabled() bool {
	if o != nil && o.ProvisioningEnabled != nil {
		return true
	}

	return false
}

// SetProvisioningEnabled gets a reference to the given bool and assigns it to the ProvisioningEnabled field.
func (o *AppAccountContainerDetails) SetProvisioningEnabled(v bool) {
	o.ProvisioningEnabled = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppAccountContainerDetails) GetLinks() AppLink {
	if o == nil || o.Links == nil {
		var ret AppLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetLinksOk() (*AppLink, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppLink and assigns it to the Links field.
func (o *AppAccountContainerDetails) SetLinks(v AppLink) {
	o.Links = &v
}

func (o AppAccountContainerDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppName != nil {
		toSerialize["appName"] = o.AppName
	}
	if true {
		toSerialize["containerId"] = o.ContainerId
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.GlobalAppId != nil {
		toSerialize["globalAppId"] = o.GlobalAppId
	}
	if o.PasswordPushSupported != nil {
		toSerialize["passwordPushSupported"] = o.PasswordPushSupported
	}
	if o.ProvisioningEnabled != nil {
		toSerialize["provisioningEnabled"] = o.ProvisioningEnabled
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppAccountContainerDetails) UnmarshalJSON(bytes []byte) (err error) {
	varAppAccountContainerDetails := _AppAccountContainerDetails{}

	err = json.Unmarshal(bytes, &varAppAccountContainerDetails)
	if err == nil {
		*o = AppAccountContainerDetails(varAppAccountContainerDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appName")
		delete(additionalProperties, "containerId")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "globalAppId")
		delete(additionalProperties, "passwordPushSupported")
		delete(additionalProperties, "provisioningEnabled")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppAccountContainerDetails struct {
	value *AppAccountContainerDetails
	isSet bool
}

func (v NullableAppAccountContainerDetails) Get() *AppAccountContainerDetails {
	return v.value
}

func (v *NullableAppAccountContainerDetails) Set(val *AppAccountContainerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAccountContainerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAccountContainerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAccountContainerDetails(val *AppAccountContainerDetails) *NullableAppAccountContainerDetails {
	return &NullableAppAccountContainerDetails{value: val, isSet: true}
}

func (v NullableAppAccountContainerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAccountContainerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

