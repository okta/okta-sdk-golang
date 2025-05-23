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
	"time"
)

// WellKnownAppAuthenticatorConfiguration struct for WellKnownAppAuthenticatorConfiguration
type WellKnownAppAuthenticatorConfiguration struct {
	// The authenticator enrollment endpoint
	AppAuthenticatorEnrollEndpoint *string `json:"appAuthenticatorEnrollEndpoint,omitempty"`
	// The unique identifier of the app authenticator
	AuthenticatorId *string `json:"authenticatorId,omitempty"`
	// Timestamp when the Authenticator was created
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// A human-readable string that identifies the Authenticator
	Key *string `json:"key,omitempty"`
	// Timestamp when the Authenticator was last modified
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The authenticator display name
	Name *string `json:"name,omitempty"`
	// The `id` of the Okta Org
	OrgId *string `json:"orgId,omitempty"`
	Settings *WellKnownAppAuthenticatorConfigurationSettings `json:"settings,omitempty"`
	SupportedMethods []SupportedMethods `json:"supportedMethods,omitempty"`
	// The type of Authenticator
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownAppAuthenticatorConfiguration WellKnownAppAuthenticatorConfiguration

// NewWellKnownAppAuthenticatorConfiguration instantiates a new WellKnownAppAuthenticatorConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownAppAuthenticatorConfiguration() *WellKnownAppAuthenticatorConfiguration {
	this := WellKnownAppAuthenticatorConfiguration{}
	return &this
}

// NewWellKnownAppAuthenticatorConfigurationWithDefaults instantiates a new WellKnownAppAuthenticatorConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownAppAuthenticatorConfigurationWithDefaults() *WellKnownAppAuthenticatorConfiguration {
	this := WellKnownAppAuthenticatorConfiguration{}
	return &this
}

// GetAppAuthenticatorEnrollEndpoint returns the AppAuthenticatorEnrollEndpoint field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetAppAuthenticatorEnrollEndpoint() string {
	if o == nil || o.AppAuthenticatorEnrollEndpoint == nil {
		var ret string
		return ret
	}
	return *o.AppAuthenticatorEnrollEndpoint
}

// GetAppAuthenticatorEnrollEndpointOk returns a tuple with the AppAuthenticatorEnrollEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetAppAuthenticatorEnrollEndpointOk() (*string, bool) {
	if o == nil || o.AppAuthenticatorEnrollEndpoint == nil {
		return nil, false
	}
	return o.AppAuthenticatorEnrollEndpoint, true
}

// HasAppAuthenticatorEnrollEndpoint returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasAppAuthenticatorEnrollEndpoint() bool {
	if o != nil && o.AppAuthenticatorEnrollEndpoint != nil {
		return true
	}

	return false
}

// SetAppAuthenticatorEnrollEndpoint gets a reference to the given string and assigns it to the AppAuthenticatorEnrollEndpoint field.
func (o *WellKnownAppAuthenticatorConfiguration) SetAppAuthenticatorEnrollEndpoint(v string) {
	o.AppAuthenticatorEnrollEndpoint = &v
}

// GetAuthenticatorId returns the AuthenticatorId field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetAuthenticatorId() string {
	if o == nil || o.AuthenticatorId == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatorId
}

// GetAuthenticatorIdOk returns a tuple with the AuthenticatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetAuthenticatorIdOk() (*string, bool) {
	if o == nil || o.AuthenticatorId == nil {
		return nil, false
	}
	return o.AuthenticatorId, true
}

// HasAuthenticatorId returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasAuthenticatorId() bool {
	if o != nil && o.AuthenticatorId != nil {
		return true
	}

	return false
}

// SetAuthenticatorId gets a reference to the given string and assigns it to the AuthenticatorId field.
func (o *WellKnownAppAuthenticatorConfiguration) SetAuthenticatorId(v string) {
	o.AuthenticatorId = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetCreatedDate() time.Time {
	if o == nil || o.CreatedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *WellKnownAppAuthenticatorConfiguration) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *WellKnownAppAuthenticatorConfiguration) SetKey(v string) {
	o.Key = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WellKnownAppAuthenticatorConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WellKnownAppAuthenticatorConfiguration) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *WellKnownAppAuthenticatorConfiguration) SetOrgId(v string) {
	o.OrgId = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetSettings() WellKnownAppAuthenticatorConfigurationSettings {
	if o == nil || o.Settings == nil {
		var ret WellKnownAppAuthenticatorConfigurationSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetSettingsOk() (*WellKnownAppAuthenticatorConfigurationSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given WellKnownAppAuthenticatorConfigurationSettings and assigns it to the Settings field.
func (o *WellKnownAppAuthenticatorConfiguration) SetSettings(v WellKnownAppAuthenticatorConfigurationSettings) {
	o.Settings = &v
}

// GetSupportedMethods returns the SupportedMethods field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetSupportedMethods() []SupportedMethods {
	if o == nil || o.SupportedMethods == nil {
		var ret []SupportedMethods
		return ret
	}
	return o.SupportedMethods
}

// GetSupportedMethodsOk returns a tuple with the SupportedMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetSupportedMethodsOk() ([]SupportedMethods, bool) {
	if o == nil || o.SupportedMethods == nil {
		return nil, false
	}
	return o.SupportedMethods, true
}

// HasSupportedMethods returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasSupportedMethods() bool {
	if o != nil && o.SupportedMethods != nil {
		return true
	}

	return false
}

// SetSupportedMethods gets a reference to the given []SupportedMethods and assigns it to the SupportedMethods field.
func (o *WellKnownAppAuthenticatorConfiguration) SetSupportedMethods(v []SupportedMethods) {
	o.SupportedMethods = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WellKnownAppAuthenticatorConfiguration) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownAppAuthenticatorConfiguration) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WellKnownAppAuthenticatorConfiguration) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WellKnownAppAuthenticatorConfiguration) SetType(v string) {
	o.Type = &v
}

func (o WellKnownAppAuthenticatorConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppAuthenticatorEnrollEndpoint != nil {
		toSerialize["appAuthenticatorEnrollEndpoint"] = o.AppAuthenticatorEnrollEndpoint
	}
	if o.AuthenticatorId != nil {
		toSerialize["authenticatorId"] = o.AuthenticatorId
	}
	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrgId != nil {
		toSerialize["orgId"] = o.OrgId
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.SupportedMethods != nil {
		toSerialize["supportedMethods"] = o.SupportedMethods
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WellKnownAppAuthenticatorConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varWellKnownAppAuthenticatorConfiguration := _WellKnownAppAuthenticatorConfiguration{}

	err = json.Unmarshal(bytes, &varWellKnownAppAuthenticatorConfiguration)
	if err == nil {
		*o = WellKnownAppAuthenticatorConfiguration(varWellKnownAppAuthenticatorConfiguration)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "appAuthenticatorEnrollEndpoint")
		delete(additionalProperties, "authenticatorId")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "key")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "supportedMethods")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWellKnownAppAuthenticatorConfiguration struct {
	value *WellKnownAppAuthenticatorConfiguration
	isSet bool
}

func (v NullableWellKnownAppAuthenticatorConfiguration) Get() *WellKnownAppAuthenticatorConfiguration {
	return v.value
}

func (v *NullableWellKnownAppAuthenticatorConfiguration) Set(val *WellKnownAppAuthenticatorConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownAppAuthenticatorConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownAppAuthenticatorConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownAppAuthenticatorConfiguration(val *WellKnownAppAuthenticatorConfiguration) *NullableWellKnownAppAuthenticatorConfiguration {
	return &NullableWellKnownAppAuthenticatorConfiguration{value: val, isSet: true}
}

func (v NullableWellKnownAppAuthenticatorConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownAppAuthenticatorConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

