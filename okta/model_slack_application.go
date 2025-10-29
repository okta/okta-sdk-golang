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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the SlackApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlackApplication{}

// SlackApplication Schema for the Slack app (key name: `slack`)  To create a Slack app, use the [Create an Application](/openapi/okta-management/management/tag/Application/#tag/Application/operation/createApplication) request with the following parameters in the request body. > **Note:** The Slack app only supports `BROWSER_PLUGIN` and `SAML_2_0` sign-on modes.
type SlackApplication struct {
	Accessibility *ApplicationAccessibility     `json:"accessibility,omitempty"`
	Credentials   *SchemeApplicationCredentials `json:"credentials,omitempty"`
	// User-defined display name for app
	Label     string                `json:"label"`
	Licensing *ApplicationLicensing `json:"licensing,omitempty"`
	Name      string                `json:"name"`
	// Contains any valid JSON schema for specifying properties that can be referenced from a request (only available to OAuth 2.0 client apps)
	Profile    map[string]map[string]interface{} `json:"profile,omitempty"`
	SignOnMode *string                           `json:"signOnMode,omitempty"`
	// App instance status
	Status               *string                  `json:"status,omitempty"`
	Visibility           *ApplicationVisibility   `json:"visibility,omitempty"`
	Settings             SlackApplicationSettings `json:"settings"`
	AdditionalProperties map[string]interface{}
}

type _SlackApplication SlackApplication

// NewSlackApplication instantiates a new SlackApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackApplication(label string, name string, settings SlackApplicationSettings) *SlackApplication {
	this := SlackApplication{}
	this.Label = label
	this.Name = name
	this.Settings = settings
	return &this
}

// NewSlackApplicationWithDefaults instantiates a new SlackApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackApplicationWithDefaults() *SlackApplication {
	this := SlackApplication{}
	return &this
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *SlackApplication) GetAccessibility() ApplicationAccessibility {
	if o == nil || IsNil(o.Accessibility) {
		var ret ApplicationAccessibility
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetAccessibilityOk() (*ApplicationAccessibility, bool) {
	if o == nil || IsNil(o.Accessibility) {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *SlackApplication) HasAccessibility() bool {
	if o != nil && !IsNil(o.Accessibility) {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given ApplicationAccessibility and assigns it to the Accessibility field.
func (o *SlackApplication) SetAccessibility(v ApplicationAccessibility) {
	o.Accessibility = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *SlackApplication) GetCredentials() SchemeApplicationCredentials {
	if o == nil || IsNil(o.Credentials) {
		var ret SchemeApplicationCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetCredentialsOk() (*SchemeApplicationCredentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *SlackApplication) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given SchemeApplicationCredentials and assigns it to the Credentials field.
func (o *SlackApplication) SetCredentials(v SchemeApplicationCredentials) {
	o.Credentials = &v
}

// GetLabel returns the Label field value
func (o *SlackApplication) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *SlackApplication) SetLabel(v string) {
	o.Label = v
}

// GetLicensing returns the Licensing field value if set, zero value otherwise.
func (o *SlackApplication) GetLicensing() ApplicationLicensing {
	if o == nil || IsNil(o.Licensing) {
		var ret ApplicationLicensing
		return ret
	}
	return *o.Licensing
}

// GetLicensingOk returns a tuple with the Licensing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetLicensingOk() (*ApplicationLicensing, bool) {
	if o == nil || IsNil(o.Licensing) {
		return nil, false
	}
	return o.Licensing, true
}

// HasLicensing returns a boolean if a field has been set.
func (o *SlackApplication) HasLicensing() bool {
	if o != nil && !IsNil(o.Licensing) {
		return true
	}

	return false
}

// SetLicensing gets a reference to the given ApplicationLicensing and assigns it to the Licensing field.
func (o *SlackApplication) SetLicensing(v ApplicationLicensing) {
	o.Licensing = &v
}

// GetName returns the Name field value
func (o *SlackApplication) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SlackApplication) SetName(v string) {
	o.Name = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SlackApplication) GetProfile() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetProfileOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Profile) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SlackApplication) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]map[string]interface{} and assigns it to the Profile field.
func (o *SlackApplication) SetProfile(v map[string]map[string]interface{}) {
	o.Profile = v
}

// GetSignOnMode returns the SignOnMode field value if set, zero value otherwise.
func (o *SlackApplication) GetSignOnMode() string {
	if o == nil || IsNil(o.SignOnMode) {
		var ret string
		return ret
	}
	return *o.SignOnMode
}

// GetSignOnModeOk returns a tuple with the SignOnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetSignOnModeOk() (*string, bool) {
	if o == nil || IsNil(o.SignOnMode) {
		return nil, false
	}
	return o.SignOnMode, true
}

// HasSignOnMode returns a boolean if a field has been set.
func (o *SlackApplication) HasSignOnMode() bool {
	if o != nil && !IsNil(o.SignOnMode) {
		return true
	}

	return false
}

// SetSignOnMode gets a reference to the given string and assigns it to the SignOnMode field.
func (o *SlackApplication) SetSignOnMode(v string) {
	o.SignOnMode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SlackApplication) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SlackApplication) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SlackApplication) SetStatus(v string) {
	o.Status = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *SlackApplication) GetVisibility() ApplicationVisibility {
	if o == nil || IsNil(o.Visibility) {
		var ret ApplicationVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetVisibilityOk() (*ApplicationVisibility, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *SlackApplication) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given ApplicationVisibility and assigns it to the Visibility field.
func (o *SlackApplication) SetVisibility(v ApplicationVisibility) {
	o.Visibility = &v
}

// GetSettings returns the Settings field value
func (o *SlackApplication) GetSettings() SlackApplicationSettings {
	if o == nil {
		var ret SlackApplicationSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *SlackApplication) GetSettingsOk() (*SlackApplicationSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *SlackApplication) SetSettings(v SlackApplicationSettings) {
	o.Settings = v
}

func (o SlackApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlackApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accessibility) {
		toSerialize["accessibility"] = o.Accessibility
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.Licensing) {
		toSerialize["licensing"] = o.Licensing
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.SignOnMode) {
		toSerialize["signOnMode"] = o.SignOnMode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SlackApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"name",
		"settings",
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

	varSlackApplication := _SlackApplication{}

	err = json.Unmarshal(data, &varSlackApplication)

	if err != nil {
		return err
	}

	*o = SlackApplication(varSlackApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessibility")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "label")
		delete(additionalProperties, "licensing")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "signOnMode")
		delete(additionalProperties, "status")
		delete(additionalProperties, "visibility")
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSlackApplication struct {
	value *SlackApplication
	isSet bool
}

func (v NullableSlackApplication) Get() *SlackApplication {
	return v.value
}

func (v *NullableSlackApplication) Set(val *SlackApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackApplication(val *SlackApplication) *NullableSlackApplication {
	return &NullableSlackApplication{value: val, isSet: true}
}

func (v NullableSlackApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
