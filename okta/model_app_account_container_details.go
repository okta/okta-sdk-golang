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
	"fmt"
)

// checks if the AppAccountContainerDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAccountContainerDetails{}

// AppAccountContainerDetails Container details for resource type APP_ACCOUNT
type AppAccountContainerDetails struct {
	// The application name
	AppName *string `json:"appName,omitempty"`
	// The app ID associated with the privileged resource
	ContainerId string `json:"containerId"`
	// Human-readable name of the container that owns the privileged resource
	DisplayName *string `json:"displayName,omitempty"`
	// The application global ID
	GlobalAppId *string `json:"globalAppId,omitempty"`
	// Indicates if the application supports password push
	PasswordPushSupported *bool `json:"passwordPushSupported,omitempty"`
	// Indicates if provisioning is enabled for this application
	ProvisioningEnabled *bool `json:"provisioningEnabled,omitempty"`
	// Current status of the application instance
	Status               *string                  `json:"status,omitempty"`
	Links                *AppAccountContainerLink `json:"_links,omitempty"`
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
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
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
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
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
	if o == nil || IsNil(o.GlobalAppId) {
		var ret string
		return ret
	}
	return *o.GlobalAppId
}

// GetGlobalAppIdOk returns a tuple with the GlobalAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetGlobalAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalAppId) {
		return nil, false
	}
	return o.GlobalAppId, true
}

// HasGlobalAppId returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasGlobalAppId() bool {
	if o != nil && !IsNil(o.GlobalAppId) {
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
	if o == nil || IsNil(o.PasswordPushSupported) {
		var ret bool
		return ret
	}
	return *o.PasswordPushSupported
}

// GetPasswordPushSupportedOk returns a tuple with the PasswordPushSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetPasswordPushSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordPushSupported) {
		return nil, false
	}
	return o.PasswordPushSupported, true
}

// HasPasswordPushSupported returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasPasswordPushSupported() bool {
	if o != nil && !IsNil(o.PasswordPushSupported) {
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
	if o == nil || IsNil(o.ProvisioningEnabled) {
		var ret bool
		return ret
	}
	return *o.ProvisioningEnabled
}

// GetProvisioningEnabledOk returns a tuple with the ProvisioningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetProvisioningEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProvisioningEnabled) {
		return nil, false
	}
	return o.ProvisioningEnabled, true
}

// HasProvisioningEnabled returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasProvisioningEnabled() bool {
	if o != nil && !IsNil(o.ProvisioningEnabled) {
		return true
	}

	return false
}

// SetProvisioningEnabled gets a reference to the given bool and assigns it to the ProvisioningEnabled field.
func (o *AppAccountContainerDetails) SetProvisioningEnabled(v bool) {
	o.ProvisioningEnabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppAccountContainerDetails) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppAccountContainerDetails) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppAccountContainerDetails) GetLinks() AppAccountContainerLink {
	if o == nil || IsNil(o.Links) {
		var ret AppAccountContainerLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAccountContainerDetails) GetLinksOk() (*AppAccountContainerLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppAccountContainerDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAccountContainerLink and assigns it to the Links field.
func (o *AppAccountContainerDetails) SetLinks(v AppAccountContainerLink) {
	o.Links = &v
}

func (o AppAccountContainerDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAccountContainerDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppName) {
		toSerialize["appName"] = o.AppName
	}
	toSerialize["containerId"] = o.ContainerId
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.GlobalAppId) {
		toSerialize["globalAppId"] = o.GlobalAppId
	}
	if !IsNil(o.PasswordPushSupported) {
		toSerialize["passwordPushSupported"] = o.PasswordPushSupported
	}
	if !IsNil(o.ProvisioningEnabled) {
		toSerialize["provisioningEnabled"] = o.ProvisioningEnabled
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppAccountContainerDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"containerId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAppAccountContainerDetails := _AppAccountContainerDetails{}

	err = json.Unmarshal(data, &varAppAccountContainerDetails)

	if err != nil {
		return err
	}

	*o = AppAccountContainerDetails(varAppAccountContainerDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appName")
		delete(additionalProperties, "containerId")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "globalAppId")
		delete(additionalProperties, "passwordPushSupported")
		delete(additionalProperties, "provisioningEnabled")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
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
