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

// SwaApplicationSettingsApplication struct for SwaApplicationSettingsApplication
type SwaApplicationSettingsApplication struct {
	ButtonField *string `json:"buttonField,omitempty"`
	ButtonSelector *string `json:"buttonSelector,omitempty"`
	Checkbox *string `json:"checkbox,omitempty"`
	ExtraFieldSelector *string `json:"extraFieldSelector,omitempty"`
	ExtraFieldValue *string `json:"extraFieldValue,omitempty"`
	LoginUrlRegex *string `json:"loginUrlRegex,omitempty"`
	PasswordField *string `json:"passwordField,omitempty"`
	PasswordSelector *string `json:"passwordSelector,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	TargetURL *string `json:"targetURL,omitempty"`
	Url *string `json:"url,omitempty"`
	UsernameField *string `json:"usernameField,omitempty"`
	UserNameSelector *string `json:"userNameSelector,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SwaApplicationSettingsApplication SwaApplicationSettingsApplication

// NewSwaApplicationSettingsApplication instantiates a new SwaApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwaApplicationSettingsApplication() *SwaApplicationSettingsApplication {
	this := SwaApplicationSettingsApplication{}
	return &this
}

// NewSwaApplicationSettingsApplicationWithDefaults instantiates a new SwaApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwaApplicationSettingsApplicationWithDefaults() *SwaApplicationSettingsApplication {
	this := SwaApplicationSettingsApplication{}
	return &this
}

// GetButtonField returns the ButtonField field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetButtonField() string {
	if o == nil || o.ButtonField == nil {
		var ret string
		return ret
	}
	return *o.ButtonField
}

// GetButtonFieldOk returns a tuple with the ButtonField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetButtonFieldOk() (*string, bool) {
	if o == nil || o.ButtonField == nil {
		return nil, false
	}
	return o.ButtonField, true
}

// HasButtonField returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasButtonField() bool {
	if o != nil && o.ButtonField != nil {
		return true
	}

	return false
}

// SetButtonField gets a reference to the given string and assigns it to the ButtonField field.
func (o *SwaApplicationSettingsApplication) SetButtonField(v string) {
	o.ButtonField = &v
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

// GetCheckbox returns the Checkbox field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetCheckbox() string {
	if o == nil || o.Checkbox == nil {
		var ret string
		return ret
	}
	return *o.Checkbox
}

// GetCheckboxOk returns a tuple with the Checkbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetCheckboxOk() (*string, bool) {
	if o == nil || o.Checkbox == nil {
		return nil, false
	}
	return o.Checkbox, true
}

// HasCheckbox returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasCheckbox() bool {
	if o != nil && o.Checkbox != nil {
		return true
	}

	return false
}

// SetCheckbox gets a reference to the given string and assigns it to the Checkbox field.
func (o *SwaApplicationSettingsApplication) SetCheckbox(v string) {
	o.Checkbox = &v
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

// GetPasswordField returns the PasswordField field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetPasswordField() string {
	if o == nil || o.PasswordField == nil {
		var ret string
		return ret
	}
	return *o.PasswordField
}

// GetPasswordFieldOk returns a tuple with the PasswordField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetPasswordFieldOk() (*string, bool) {
	if o == nil || o.PasswordField == nil {
		return nil, false
	}
	return o.PasswordField, true
}

// HasPasswordField returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasPasswordField() bool {
	if o != nil && o.PasswordField != nil {
		return true
	}

	return false
}

// SetPasswordField gets a reference to the given string and assigns it to the PasswordField field.
func (o *SwaApplicationSettingsApplication) SetPasswordField(v string) {
	o.PasswordField = &v
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

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *SwaApplicationSettingsApplication) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
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

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SwaApplicationSettingsApplication) SetUrl(v string) {
	o.Url = &v
}

// GetUsernameField returns the UsernameField field value if set, zero value otherwise.
func (o *SwaApplicationSettingsApplication) GetUsernameField() string {
	if o == nil || o.UsernameField == nil {
		var ret string
		return ret
	}
	return *o.UsernameField
}

// GetUsernameFieldOk returns a tuple with the UsernameField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwaApplicationSettingsApplication) GetUsernameFieldOk() (*string, bool) {
	if o == nil || o.UsernameField == nil {
		return nil, false
	}
	return o.UsernameField, true
}

// HasUsernameField returns a boolean if a field has been set.
func (o *SwaApplicationSettingsApplication) HasUsernameField() bool {
	if o != nil && o.UsernameField != nil {
		return true
	}

	return false
}

// SetUsernameField gets a reference to the given string and assigns it to the UsernameField field.
func (o *SwaApplicationSettingsApplication) SetUsernameField(v string) {
	o.UsernameField = &v
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
	if o.ButtonField != nil {
		toSerialize["buttonField"] = o.ButtonField
	}
	if o.ButtonSelector != nil {
		toSerialize["buttonSelector"] = o.ButtonSelector
	}
	if o.Checkbox != nil {
		toSerialize["checkbox"] = o.Checkbox
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
	if o.PasswordField != nil {
		toSerialize["passwordField"] = o.PasswordField
	}
	if o.PasswordSelector != nil {
		toSerialize["passwordSelector"] = o.PasswordSelector
	}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if o.TargetURL != nil {
		toSerialize["targetURL"] = o.TargetURL
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.UsernameField != nil {
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
		delete(additionalProperties, "checkbox")
		delete(additionalProperties, "extraFieldSelector")
		delete(additionalProperties, "extraFieldValue")
		delete(additionalProperties, "loginUrlRegex")
		delete(additionalProperties, "passwordField")
		delete(additionalProperties, "passwordSelector")
		delete(additionalProperties, "redirectUrl")
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

