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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// SwaApplicationSettingsApplication struct for SwaApplicationSettingsApplication
type SwaApplicationSettingsApplication struct {
	// CSS selector for the **Sign-In** button in the sign-in form (for SWA apps with the `template_swa` app name definition)
	ButtonField string `json:"buttonField"`
	// CSS selector for the **Sign-In**  button in the sign-in form (for three-field SWA apps with the `template_swa3field` app name definition)
	ButtonSelector *string `json:"buttonSelector,omitempty"`
	// Enter the CSS selector for the extra field (for three-field SWA apps with the `template_swa3field` app name definition).
	ExtraFieldSelector *string `json:"extraFieldSelector,omitempty"`
	// Enter the value for the extra field in the form (for three-field SWA apps with the `template_swa3field` app name definition).
	ExtraFieldValue *string `json:"extraFieldValue,omitempty"`
	// A regular expression that further restricts targetURL to the specified regular expression
	LoginUrlRegex *string `json:"loginUrlRegex,omitempty"`
	// CSS selector for the **Password** field in the sign-in form (for SWA apps with the `template_swa` app name definition)
	PasswordField string `json:"passwordField"`
	// CSS selector for the **Password** field in the sign-in form (for three-field SWA apps with the `template_swa3field` app name definition)
	PasswordSelector *string `json:"passwordSelector,omitempty"`
	// The URL of the sign-in page for this app (for three-field SWA apps with the `template_swa3field` app name definition)
	TargetURL *string `json:"targetURL,omitempty"`
	// The URL of the sign-in page for this app (for SWA apps with the `template_swa` app name definition)
	Url string `json:"url"`
	// CSS selector for the **Username** field in the sign-in form (for SWA apps with the `template_swa` app name definition)
	UsernameField string `json:"usernameField"`
	// CSS selector for the **Username** field in the sign-in form (for three-field SWA apps with the `template_swa3field` app name definition)
	UserNameSelector *string `json:"userNameSelector,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwaApplicationSettingsApplication SwaApplicationSettingsApplication

// NewSwaApplicationSettingsApplication instantiates a new SwaApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwaApplicationSettingsApplication(buttonField string, passwordField string, url string, usernameField string) *SwaApplicationSettingsApplication {
	this := SwaApplicationSettingsApplication{}
	this.ButtonField = buttonField
	this.PasswordField = passwordField
	this.Url = url
	this.UsernameField = usernameField
	return &this
}

// NewSwaApplicationSettingsApplicationWithDefaults instantiates a new SwaApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwaApplicationSettingsApplicationWithDefaults() *SwaApplicationSettingsApplication {
	this := SwaApplicationSettingsApplication{}
	return &this
}

// GetButtonField returns the ButtonField field value
func (o *SwaApplicationSettingsApplication) GetButtonField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ButtonField
}

// GetButtonFieldOk returns a tuple with the ButtonField field value
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetButtonFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ButtonField, true
}

// SetButtonField sets field value
func (o *SwaApplicationSettingsApplication) SetButtonField(v string) {
	o.ButtonField = v
}

// GetButtonSelector returns the ButtonSelector field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetButtonSelector() string {
	if o == nil || o.ButtonSelector == nil {
		var ret string
		return ret
	}
	return *o.ButtonSelector
}

// GetButtonSelectorOk returns a tuple with the ButtonSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetButtonSelectorOk() (*string, bool) {
	if o == nil || o.ButtonSelector == nil {
		return nil, false
	}
	return o.ButtonSelector, true
}

// HasButtonSelector returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasButtonSelector() bool {
	if o != nil && o.ButtonSelector != nil {
		return true
	}

	return false
}

// SetButtonSelector gets a reference to the given string and assigns it to the ButtonSelector field.
func (o *SwaApplicationSettingsApplication) SetButtonSelector(v string) {
	o.ButtonSelector = &v
}

// GetExtraFieldSelector returns the ExtraFieldSelector field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetExtraFieldSelector() string {
	if o == nil || o.ExtraFieldSelector == nil {
		var ret string
		return ret
	}
	return *o.ExtraFieldSelector
}

// GetExtraFieldSelectorOk returns a tuple with the ExtraFieldSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetExtraFieldSelectorOk() (*string, bool) {
	if o == nil || o.ExtraFieldSelector == nil {
		return nil, false
	}
	return o.ExtraFieldSelector, true
}

// HasExtraFieldSelector returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasExtraFieldSelector() bool {
	if o != nil && o.ExtraFieldSelector != nil {
		return true
	}

	return false
}

// SetExtraFieldSelector gets a reference to the given string and assigns it to the ExtraFieldSelector field.
func (o *SwaApplicationSettingsApplication) SetExtraFieldSelector(v string) {
	o.ExtraFieldSelector = &v
}

// GetExtraFieldValue returns the ExtraFieldValue field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetExtraFieldValue() string {
	if o == nil || o.ExtraFieldValue == nil {
		var ret string
		return ret
	}
	return *o.ExtraFieldValue
}

// GetExtraFieldValueOk returns a tuple with the ExtraFieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetExtraFieldValueOk() (*string, bool) {
	if o == nil || o.ExtraFieldValue == nil {
		return nil, false
	}
	return o.ExtraFieldValue, true
}

// HasExtraFieldValue returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasExtraFieldValue() bool {
	if o != nil && o.ExtraFieldValue != nil {
		return true
	}

	return false
}

// SetExtraFieldValue gets a reference to the given string and assigns it to the ExtraFieldValue field.
func (o *SwaApplicationSettingsApplication) SetExtraFieldValue(v string) {
	o.ExtraFieldValue = &v
}

// GetLoginUrlRegex returns the LoginUrlRegex field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetLoginUrlRegex() string {
	if o == nil || o.LoginUrlRegex == nil {
		var ret string
		return ret
	}
	return *o.LoginUrlRegex
}

// GetLoginUrlRegexOk returns a tuple with the LoginUrlRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetLoginUrlRegexOk() (*string, bool) {
	if o == nil || o.LoginUrlRegex == nil {
		return nil, false
	}
	return o.LoginUrlRegex, true
}

// HasLoginUrlRegex returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasLoginUrlRegex() bool {
	if o != nil && o.LoginUrlRegex != nil {
		return true
	}

	return false
}

// SetLoginUrlRegex gets a reference to the given string and assigns it to the LoginUrlRegex field.
func (o *SwaApplicationSettingsApplication) SetLoginUrlRegex(v string) {
	o.LoginUrlRegex = &v
}

// GetPasswordField returns the PasswordField field value
func (o *SwaApplicationSettingsApplication) GetPasswordField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetPasswordFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordField, true
}

// SetPasswordField sets field value
func (o *SwaApplicationSettingsApplication) SetPasswordField(v string) {
	o.PasswordField = v
}

// GetPasswordSelector returns the PasswordSelector field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetPasswordSelector() string {
	if o == nil || o.PasswordSelector == nil {
		var ret string
		return ret
	}
	return *o.PasswordSelector
}

// GetPasswordSelectorOk returns a tuple with the PasswordSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetPasswordSelectorOk() (*string, bool) {
	if o == nil || o.PasswordSelector == nil {
		return nil, false
	}
	return o.PasswordSelector, true
}

// HasPasswordSelector returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasPasswordSelector() bool {
	if o != nil && o.PasswordSelector != nil {
		return true
	}

	return false
}

// SetPasswordSelector gets a reference to the given string and assigns it to the PasswordSelector field.
func (o *SwaApplicationSettingsApplication) SetPasswordSelector(v string) {
	o.PasswordSelector = &v
}

// GetTargetURL returns the TargetURL field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetTargetURL() string {
	if o == nil || o.TargetURL == nil {
		var ret string
		return ret
	}
	return *o.TargetURL
}

// GetTargetURLOk returns a tuple with the TargetURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetTargetURLOk() (*string, bool) {
	if o == nil || o.TargetURL == nil {
		return nil, false
	}
	return o.TargetURL, true
}

// HasTargetURL returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasTargetURL() bool {
	if o != nil && o.TargetURL != nil {
		return true
	}

	return false
}

// SetTargetURL gets a reference to the given string and assigns it to the TargetURL field.
func (o *SwaApplicationSettingsApplication) SetTargetURL(v string) {
	o.TargetURL = &v
}

// GetUrl returns the Url field value
func (o *SwaApplicationSettingsApplication) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SwaApplicationSettingsApplication) SetUrl(v string) {
	o.Url = v
}

// GetUsernameField returns the UsernameField field value
func (o *SwaApplicationSettingsApplication) GetUsernameField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsernameField
}

// GetUsernameFieldOk returns a tuple with the UsernameField field value
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetUsernameFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsernameField, true
}

// SetUsernameField sets field value
func (o *SwaApplicationSettingsApplication) SetUsernameField(v string) {
	o.UsernameField = v
}

// GetUserNameSelector returns the UserNameSelector field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetUserNameSelector() string {
	if o == nil || o.UserNameSelector == nil {
		var ret string
		return ret
	}
	return *o.UserNameSelector
}

// GetUserNameSelectorOk returns a tuple with the UserNameSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetUserNameSelectorOk() (*string, bool) {
	if o == nil || o.UserNameSelector == nil {
		return nil, false
	}
	return o.UserNameSelector, true
}

// HasUserNameSelector returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasUserNameSelector() bool {
	if o != nil && o.UserNameSelector != nil {
		return true
	}

	return false
}

// SetUserNameSelector gets a reference to the given string and assigns it to the UserNameSelector field.
func (o *SwaApplicationSettingsApplication) SetUserNameSelector(v string) {
	o.UserNameSelector = &v
}

func (o SwaApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["buttonField"] = o.ButtonField
	}
	if o.ButtonSelector != nil {
		toSerialize["buttonSelector"] = o.ButtonSelector
	}
	if o.ExtraFieldSelector != nil {
		toSerialize["extraFieldSelector"] = o.ExtraFieldSelector
	}
	if o.ExtraFieldValue != nil {
		toSerialize["extraFieldValue"] = o.ExtraFieldValue
	}
	if o.LoginUrlRegex != nil {
		toSerialize["loginUrlRegex"] = o.LoginUrlRegex
	}
	if true {
		toSerialize["passwordField"] = o.PasswordField
	}
	if o.PasswordSelector != nil {
		toSerialize["passwordSelector"] = o.PasswordSelector
	}
	if o.TargetURL != nil {
		toSerialize["targetURL"] = o.TargetURL
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["usernameField"] = o.UsernameField
	}
	if o.UserNameSelector != nil {
		toSerialize["userNameSelector"] = o.UserNameSelector
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SwaApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varSwaApplicationSettingsApplication := _SwaApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varSwaApplicationSettingsApplication)
	if err == nil {
		*o = SwaApplicationSettingsApplication(varSwaApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "buttonField")
		delete(additionalProperties, "buttonSelector")
		delete(additionalProperties, "extraFieldSelector")
		delete(additionalProperties, "extraFieldValue")
		delete(additionalProperties, "loginUrlRegex")
		delete(additionalProperties, "passwordField")
		delete(additionalProperties, "passwordSelector")
		delete(additionalProperties, "targetURL")
		delete(additionalProperties, "url")
		delete(additionalProperties, "usernameField")
		delete(additionalProperties, "userNameSelector")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSwaApplicationSettingsApplication struct {
	value *SwaApplicationSettingsApplication
	isSet bool
}

func (v NullableSwaApplicationSettingsApplication) Get() *SwaApplicationSettingsApplication {
	return v.value
}

func (v *NullableSwaApplicationSettingsApplication) Set(val *SwaApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableSwaApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableSwaApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwaApplicationSettingsApplication(val *SwaApplicationSettingsApplication) *NullableSwaApplicationSettingsApplication {
	return &NullableSwaApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableSwaApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwaApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

