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

// OINApplication struct for OINApplication
type OINApplication struct {
	Accessibility *ApplicationAccessibility `json:"accessibility,omitempty"`
	Credentials *SchemeApplicationCredentials `json:"credentials,omitempty"`
	// User-defined display name for app
	Label *string `json:"label,omitempty"`
	Licensing *ApplicationLicensing `json:"licensing,omitempty"`
	// The key name for the OIN app definition
	Name *string `json:"name,omitempty"`
	// Contains any valid JSON schema for specifying properties that can be referenced from a request (only available to OAuth 2.0 client apps)
	Profile map[string]map[string]interface{} `json:"profile,omitempty"`
	// Authentication mode for the app
	SignOnMode *string `json:"signOnMode,omitempty"`
	// App instance status
	Status *string `json:"status,omitempty"`
	Visibility *ApplicationVisibility `json:"visibility,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OINApplication OINApplication

// NewOINApplication instantiates a new OINApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOINApplication() *OINApplication {
	this := OINApplication{}
	return &this
}

// NewOINApplicationWithDefaults instantiates a new OINApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOINApplicationWithDefaults() *OINApplication {
	this := OINApplication{}
	return &this
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *OINApplication) GetAccessibility() ApplicationAccessibility {
	if o == nil || o.Accessibility == nil {
		var ret ApplicationAccessibility
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINApplication) GetAccessibilityOk() (*ApplicationAccessibility, bool) {
	if o == nil || o.Accessibility == nil {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *OINApplication) HasAccessibility() bool {
	if o != nil && o.Accessibility != nil {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given ApplicationAccessibility and assigns it to the Accessibility field.
func (o *OINApplication) SetAccessibility(v ApplicationAccessibility) {
	o.Accessibility = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *OINApplication) GetCredentials() SchemeApplicationCredentials {
	if o == nil || o.Credentials == nil {
		var ret SchemeApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *OINApplication) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given SchemeApplicationCredentials and assigns it to the Credentials field.
func (o *OINApplication) SetCredentials(v SchemeApplicationCredentials) {
	o.Credentials = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *OINApplication) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINApplication) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *OINApplication) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *OINApplication) SetLabel(v string) {
	o.Label = &v
}

// GetLicensing returns the Licensing field value if set, zero value otherwise.
func (o *OINApplication) GetLicensing() ApplicationLicensing {
	if o == nil || o.Licensing == nil {
		var ret ApplicationLicensing
		return ret
	}
	return *o.Licensing
}

// GetLicensingOk returns a tuple with the Licensing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINApplication) GetLicensingOk() (*ApplicationLicensing, bool) {
	if o == nil || o.Licensing == nil {
		return nil, false
	}
	return o.Licensing, true
}

// HasLicensing returns a boolean if a field has been set.
func (o *OINApplication) HasLicensing() bool {
	if o != nil && o.Licensing != nil {
		return true
	}

	return false
}

// SetLicensing gets a reference to the given ApplicationLicensing and assigns it to the Licensing field.
func (o *OINApplication) SetLicensing(v ApplicationLicensing) {
	o.Licensing = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OINApplication) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINApplication) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OINApplication) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OINApplication) SetName(v string) {
	o.Name = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *OINApplication) GetProfile() map[string]map[string]interface{} {
	if o == nil || o.Profile == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINApplication) GetProfileOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *OINApplication) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]map[string]interface{} and assigns it to the Profile field.
func (o *OINApplication) SetProfile(v map[string]map[string]interface{}) {
	o.Profile = v
}

// GetSignOnMode returns the SignOnMode field value if set, zero value otherwise.
func (o *OINApplication) GetSignOnMode() string {
	if o == nil || o.SignOnMode == nil {
		var ret string
		return ret
	}
	return *o.SignOnMode
}

// GetSignOnModeOk returns a tuple with the SignOnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINApplication) GetSignOnModeOk() (*string, bool) {
	if o == nil || o.SignOnMode == nil {
		return nil, false
	}
	return o.SignOnMode, true
}

// HasSignOnMode returns a boolean if a field has been set.
func (o *OINApplication) HasSignOnMode() bool {
	if o != nil && o.SignOnMode != nil {
		return true
	}

	return false
}

// SetSignOnMode gets a reference to the given string and assigns it to the SignOnMode field.
func (o *OINApplication) SetSignOnMode(v string) {
	o.SignOnMode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OINApplication) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINApplication) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OINApplication) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OINApplication) SetStatus(v string) {
	o.Status = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *OINApplication) GetVisibility() ApplicationVisibility {
	if o == nil || o.Visibility == nil {
		var ret ApplicationVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OINApplication) GetVisibilityOk() (*ApplicationVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *OINApplication) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given ApplicationVisibility and assigns it to the Visibility field.
func (o *OINApplication) SetVisibility(v ApplicationVisibility) {
	o.Visibility = &v
}

func (o OINApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessibility != nil {
		toSerialize["accessibility"] = o.Accessibility
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Licensing != nil {
		toSerialize["licensing"] = o.Licensing
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.SignOnMode != nil {
		toSerialize["signOnMode"] = o.SignOnMode
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OINApplication) UnmarshalJSON(bytes []byte) (err error) {
	varOINApplication := _OINApplication{}

	err = json.Unmarshal(bytes, &varOINApplication)
	if err == nil {
		*o = OINApplication(varOINApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accessibility")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "label")
		delete(additionalProperties, "licensing")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "signOnMode")
		delete(additionalProperties, "status")
		delete(additionalProperties, "visibility")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOINApplication struct {
	value *OINApplication
	isSet bool
}

func (v NullableOINApplication) Get() *OINApplication {
	return v.value
}

func (v *NullableOINApplication) Set(val *OINApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableOINApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableOINApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOINApplication(val *OINApplication) *NullableOINApplication {
	return &NullableOINApplication{value: val, isSet: true}
}

func (v NullableOINApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOINApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

