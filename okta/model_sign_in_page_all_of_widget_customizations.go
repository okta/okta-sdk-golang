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

// SignInPageAllOfWidgetCustomizations struct for SignInPageAllOfWidgetCustomizations
type SignInPageAllOfWidgetCustomizations struct {
	SignInLabel *string `json:"signInLabel,omitempty"`
	UsernameLabel *string `json:"usernameLabel,omitempty"`
	UsernameInfoTip *string `json:"usernameInfoTip,omitempty"`
	PasswordLabel *string `json:"passwordLabel,omitempty"`
	PasswordInfoTip *string `json:"passwordInfoTip,omitempty"`
	ShowPasswordVisibilityToggle *bool `json:"showPasswordVisibilityToggle,omitempty"`
	ShowUserIdentifier *bool `json:"showUserIdentifier,omitempty"`
	ForgotPasswordLabel *string `json:"forgotPasswordLabel,omitempty"`
	ForgotPasswordUrl *string `json:"forgotPasswordUrl,omitempty"`
	UnlockAccountLabel *string `json:"unlockAccountLabel,omitempty"`
	UnlockAccountUrl *string `json:"unlockAccountUrl,omitempty"`
	HelpLabel *string `json:"helpLabel,omitempty"`
	HelpUrl *string `json:"helpUrl,omitempty"`
	CustomLink1Label *string `json:"customLink1Label,omitempty"`
	CustomLink1Url *string `json:"customLink1Url,omitempty"`
	CustomLink2Label *string `json:"customLink2Label,omitempty"`
	CustomLink2Url *string `json:"customLink2Url,omitempty"`
	AuthenticatorPageCustomLinkLabel *string `json:"authenticatorPageCustomLinkLabel,omitempty"`
	AuthenticatorPageCustomLinkUrl *string `json:"authenticatorPageCustomLinkUrl,omitempty"`
	ClassicRecoveryFlowEmailOrUsernameLabel *string `json:"classicRecoveryFlowEmailOrUsernameLabel,omitempty"`
	WidgetGeneration *string `json:"widgetGeneration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SignInPageAllOfWidgetCustomizations SignInPageAllOfWidgetCustomizations

// NewSignInPageAllOfWidgetCustomizations instantiates a new SignInPageAllOfWidgetCustomizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignInPageAllOfWidgetCustomizations() *SignInPageAllOfWidgetCustomizations {
	this := SignInPageAllOfWidgetCustomizations{}
	return &this
}

// NewSignInPageAllOfWidgetCustomizationsWithDefaults instantiates a new SignInPageAllOfWidgetCustomizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignInPageAllOfWidgetCustomizationsWithDefaults() *SignInPageAllOfWidgetCustomizations {
	this := SignInPageAllOfWidgetCustomizations{}
	return &this
}

// GetSignInLabel returns the SignInLabel field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetSignInLabel() string {
	if o == nil || o.SignInLabel == nil {
		var ret string
		return ret
	}
	return *o.SignInLabel
}

// GetSignInLabelOk returns a tuple with the SignInLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetSignInLabelOk() (*string, bool) {
	if o == nil || o.SignInLabel == nil {
		return nil, false
	}
	return o.SignInLabel, true
}

// HasSignInLabel returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasSignInLabel() bool {
	if o != nil && o.SignInLabel != nil {
		return true
	}

	return false
}

// SetSignInLabel gets a reference to the given string and assigns it to the SignInLabel field.
func (o *SignInPageAllOfWidgetCustomizations) SetSignInLabel(v string) {
	o.SignInLabel = &v
}

// GetUsernameLabel returns the UsernameLabel field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetUsernameLabel() string {
	if o == nil || o.UsernameLabel == nil {
		var ret string
		return ret
	}
	return *o.UsernameLabel
}

// GetUsernameLabelOk returns a tuple with the UsernameLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetUsernameLabelOk() (*string, bool) {
	if o == nil || o.UsernameLabel == nil {
		return nil, false
	}
	return o.UsernameLabel, true
}

// HasUsernameLabel returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasUsernameLabel() bool {
	if o != nil && o.UsernameLabel != nil {
		return true
	}

	return false
}

// SetUsernameLabel gets a reference to the given string and assigns it to the UsernameLabel field.
func (o *SignInPageAllOfWidgetCustomizations) SetUsernameLabel(v string) {
	o.UsernameLabel = &v
}

// GetUsernameInfoTip returns the UsernameInfoTip field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetUsernameInfoTip() string {
	if o == nil || o.UsernameInfoTip == nil {
		var ret string
		return ret
	}
	return *o.UsernameInfoTip
}

// GetUsernameInfoTipOk returns a tuple with the UsernameInfoTip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetUsernameInfoTipOk() (*string, bool) {
	if o == nil || o.UsernameInfoTip == nil {
		return nil, false
	}
	return o.UsernameInfoTip, true
}

// HasUsernameInfoTip returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasUsernameInfoTip() bool {
	if o != nil && o.UsernameInfoTip != nil {
		return true
	}

	return false
}

// SetUsernameInfoTip gets a reference to the given string and assigns it to the UsernameInfoTip field.
func (o *SignInPageAllOfWidgetCustomizations) SetUsernameInfoTip(v string) {
	o.UsernameInfoTip = &v
}

// GetPasswordLabel returns the PasswordLabel field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetPasswordLabel() string {
	if o == nil || o.PasswordLabel == nil {
		var ret string
		return ret
	}
	return *o.PasswordLabel
}

// GetPasswordLabelOk returns a tuple with the PasswordLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetPasswordLabelOk() (*string, bool) {
	if o == nil || o.PasswordLabel == nil {
		return nil, false
	}
	return o.PasswordLabel, true
}

// HasPasswordLabel returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasPasswordLabel() bool {
	if o != nil && o.PasswordLabel != nil {
		return true
	}

	return false
}

// SetPasswordLabel gets a reference to the given string and assigns it to the PasswordLabel field.
func (o *SignInPageAllOfWidgetCustomizations) SetPasswordLabel(v string) {
	o.PasswordLabel = &v
}

// GetPasswordInfoTip returns the PasswordInfoTip field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetPasswordInfoTip() string {
	if o == nil || o.PasswordInfoTip == nil {
		var ret string
		return ret
	}
	return *o.PasswordInfoTip
}

// GetPasswordInfoTipOk returns a tuple with the PasswordInfoTip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetPasswordInfoTipOk() (*string, bool) {
	if o == nil || o.PasswordInfoTip == nil {
		return nil, false
	}
	return o.PasswordInfoTip, true
}

// HasPasswordInfoTip returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasPasswordInfoTip() bool {
	if o != nil && o.PasswordInfoTip != nil {
		return true
	}

	return false
}

// SetPasswordInfoTip gets a reference to the given string and assigns it to the PasswordInfoTip field.
func (o *SignInPageAllOfWidgetCustomizations) SetPasswordInfoTip(v string) {
	o.PasswordInfoTip = &v
}

// GetShowPasswordVisibilityToggle returns the ShowPasswordVisibilityToggle field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetShowPasswordVisibilityToggle() bool {
	if o == nil || o.ShowPasswordVisibilityToggle == nil {
		var ret bool
		return ret
	}
	return *o.ShowPasswordVisibilityToggle
}

// GetShowPasswordVisibilityToggleOk returns a tuple with the ShowPasswordVisibilityToggle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetShowPasswordVisibilityToggleOk() (*bool, bool) {
	if o == nil || o.ShowPasswordVisibilityToggle == nil {
		return nil, false
	}
	return o.ShowPasswordVisibilityToggle, true
}

// HasShowPasswordVisibilityToggle returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasShowPasswordVisibilityToggle() bool {
	if o != nil && o.ShowPasswordVisibilityToggle != nil {
		return true
	}

	return false
}

// SetShowPasswordVisibilityToggle gets a reference to the given bool and assigns it to the ShowPasswordVisibilityToggle field.
func (o *SignInPageAllOfWidgetCustomizations) SetShowPasswordVisibilityToggle(v bool) {
	o.ShowPasswordVisibilityToggle = &v
}

// GetShowUserIdentifier returns the ShowUserIdentifier field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetShowUserIdentifier() bool {
	if o == nil || o.ShowUserIdentifier == nil {
		var ret bool
		return ret
	}
	return *o.ShowUserIdentifier
}

// GetShowUserIdentifierOk returns a tuple with the ShowUserIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetShowUserIdentifierOk() (*bool, bool) {
	if o == nil || o.ShowUserIdentifier == nil {
		return nil, false
	}
	return o.ShowUserIdentifier, true
}

// HasShowUserIdentifier returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasShowUserIdentifier() bool {
	if o != nil && o.ShowUserIdentifier != nil {
		return true
	}

	return false
}

// SetShowUserIdentifier gets a reference to the given bool and assigns it to the ShowUserIdentifier field.
func (o *SignInPageAllOfWidgetCustomizations) SetShowUserIdentifier(v bool) {
	o.ShowUserIdentifier = &v
}

// GetForgotPasswordLabel returns the ForgotPasswordLabel field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetForgotPasswordLabel() string {
	if o == nil || o.ForgotPasswordLabel == nil {
		var ret string
		return ret
	}
	return *o.ForgotPasswordLabel
}

// GetForgotPasswordLabelOk returns a tuple with the ForgotPasswordLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetForgotPasswordLabelOk() (*string, bool) {
	if o == nil || o.ForgotPasswordLabel == nil {
		return nil, false
	}
	return o.ForgotPasswordLabel, true
}

// HasForgotPasswordLabel returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasForgotPasswordLabel() bool {
	if o != nil && o.ForgotPasswordLabel != nil {
		return true
	}

	return false
}

// SetForgotPasswordLabel gets a reference to the given string and assigns it to the ForgotPasswordLabel field.
func (o *SignInPageAllOfWidgetCustomizations) SetForgotPasswordLabel(v string) {
	o.ForgotPasswordLabel = &v
}

// GetForgotPasswordUrl returns the ForgotPasswordUrl field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetForgotPasswordUrl() string {
	if o == nil || o.ForgotPasswordUrl == nil {
		var ret string
		return ret
	}
	return *o.ForgotPasswordUrl
}

// GetForgotPasswordUrlOk returns a tuple with the ForgotPasswordUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetForgotPasswordUrlOk() (*string, bool) {
	if o == nil || o.ForgotPasswordUrl == nil {
		return nil, false
	}
	return o.ForgotPasswordUrl, true
}

// HasForgotPasswordUrl returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasForgotPasswordUrl() bool {
	if o != nil && o.ForgotPasswordUrl != nil {
		return true
	}

	return false
}

// SetForgotPasswordUrl gets a reference to the given string and assigns it to the ForgotPasswordUrl field.
func (o *SignInPageAllOfWidgetCustomizations) SetForgotPasswordUrl(v string) {
	o.ForgotPasswordUrl = &v
}

// GetUnlockAccountLabel returns the UnlockAccountLabel field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetUnlockAccountLabel() string {
	if o == nil || o.UnlockAccountLabel == nil {
		var ret string
		return ret
	}
	return *o.UnlockAccountLabel
}

// GetUnlockAccountLabelOk returns a tuple with the UnlockAccountLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetUnlockAccountLabelOk() (*string, bool) {
	if o == nil || o.UnlockAccountLabel == nil {
		return nil, false
	}
	return o.UnlockAccountLabel, true
}

// HasUnlockAccountLabel returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasUnlockAccountLabel() bool {
	if o != nil && o.UnlockAccountLabel != nil {
		return true
	}

	return false
}

// SetUnlockAccountLabel gets a reference to the given string and assigns it to the UnlockAccountLabel field.
func (o *SignInPageAllOfWidgetCustomizations) SetUnlockAccountLabel(v string) {
	o.UnlockAccountLabel = &v
}

// GetUnlockAccountUrl returns the UnlockAccountUrl field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetUnlockAccountUrl() string {
	if o == nil || o.UnlockAccountUrl == nil {
		var ret string
		return ret
	}
	return *o.UnlockAccountUrl
}

// GetUnlockAccountUrlOk returns a tuple with the UnlockAccountUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetUnlockAccountUrlOk() (*string, bool) {
	if o == nil || o.UnlockAccountUrl == nil {
		return nil, false
	}
	return o.UnlockAccountUrl, true
}

// HasUnlockAccountUrl returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasUnlockAccountUrl() bool {
	if o != nil && o.UnlockAccountUrl != nil {
		return true
	}

	return false
}

// SetUnlockAccountUrl gets a reference to the given string and assigns it to the UnlockAccountUrl field.
func (o *SignInPageAllOfWidgetCustomizations) SetUnlockAccountUrl(v string) {
	o.UnlockAccountUrl = &v
}

// GetHelpLabel returns the HelpLabel field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetHelpLabel() string {
	if o == nil || o.HelpLabel == nil {
		var ret string
		return ret
	}
	return *o.HelpLabel
}

// GetHelpLabelOk returns a tuple with the HelpLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetHelpLabelOk() (*string, bool) {
	if o == nil || o.HelpLabel == nil {
		return nil, false
	}
	return o.HelpLabel, true
}

// HasHelpLabel returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasHelpLabel() bool {
	if o != nil && o.HelpLabel != nil {
		return true
	}

	return false
}

// SetHelpLabel gets a reference to the given string and assigns it to the HelpLabel field.
func (o *SignInPageAllOfWidgetCustomizations) SetHelpLabel(v string) {
	o.HelpLabel = &v
}

// GetHelpUrl returns the HelpUrl field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetHelpUrl() string {
	if o == nil || o.HelpUrl == nil {
		var ret string
		return ret
	}
	return *o.HelpUrl
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetHelpUrlOk() (*string, bool) {
	if o == nil || o.HelpUrl == nil {
		return nil, false
	}
	return o.HelpUrl, true
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasHelpUrl() bool {
	if o != nil && o.HelpUrl != nil {
		return true
	}

	return false
}

// SetHelpUrl gets a reference to the given string and assigns it to the HelpUrl field.
func (o *SignInPageAllOfWidgetCustomizations) SetHelpUrl(v string) {
	o.HelpUrl = &v
}

// GetCustomLink1Label returns the CustomLink1Label field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetCustomLink1Label() string {
	if o == nil || o.CustomLink1Label == nil {
		var ret string
		return ret
	}
	return *o.CustomLink1Label
}

// GetCustomLink1LabelOk returns a tuple with the CustomLink1Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetCustomLink1LabelOk() (*string, bool) {
	if o == nil || o.CustomLink1Label == nil {
		return nil, false
	}
	return o.CustomLink1Label, true
}

// HasCustomLink1Label returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasCustomLink1Label() bool {
	if o != nil && o.CustomLink1Label != nil {
		return true
	}

	return false
}

// SetCustomLink1Label gets a reference to the given string and assigns it to the CustomLink1Label field.
func (o *SignInPageAllOfWidgetCustomizations) SetCustomLink1Label(v string) {
	o.CustomLink1Label = &v
}

// GetCustomLink1Url returns the CustomLink1Url field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetCustomLink1Url() string {
	if o == nil || o.CustomLink1Url == nil {
		var ret string
		return ret
	}
	return *o.CustomLink1Url
}

// GetCustomLink1UrlOk returns a tuple with the CustomLink1Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetCustomLink1UrlOk() (*string, bool) {
	if o == nil || o.CustomLink1Url == nil {
		return nil, false
	}
	return o.CustomLink1Url, true
}

// HasCustomLink1Url returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasCustomLink1Url() bool {
	if o != nil && o.CustomLink1Url != nil {
		return true
	}

	return false
}

// SetCustomLink1Url gets a reference to the given string and assigns it to the CustomLink1Url field.
func (o *SignInPageAllOfWidgetCustomizations) SetCustomLink1Url(v string) {
	o.CustomLink1Url = &v
}

// GetCustomLink2Label returns the CustomLink2Label field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetCustomLink2Label() string {
	if o == nil || o.CustomLink2Label == nil {
		var ret string
		return ret
	}
	return *o.CustomLink2Label
}

// GetCustomLink2LabelOk returns a tuple with the CustomLink2Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetCustomLink2LabelOk() (*string, bool) {
	if o == nil || o.CustomLink2Label == nil {
		return nil, false
	}
	return o.CustomLink2Label, true
}

// HasCustomLink2Label returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasCustomLink2Label() bool {
	if o != nil && o.CustomLink2Label != nil {
		return true
	}

	return false
}

// SetCustomLink2Label gets a reference to the given string and assigns it to the CustomLink2Label field.
func (o *SignInPageAllOfWidgetCustomizations) SetCustomLink2Label(v string) {
	o.CustomLink2Label = &v
}

// GetCustomLink2Url returns the CustomLink2Url field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetCustomLink2Url() string {
	if o == nil || o.CustomLink2Url == nil {
		var ret string
		return ret
	}
	return *o.CustomLink2Url
}

// GetCustomLink2UrlOk returns a tuple with the CustomLink2Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetCustomLink2UrlOk() (*string, bool) {
	if o == nil || o.CustomLink2Url == nil {
		return nil, false
	}
	return o.CustomLink2Url, true
}

// HasCustomLink2Url returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasCustomLink2Url() bool {
	if o != nil && o.CustomLink2Url != nil {
		return true
	}

	return false
}

// SetCustomLink2Url gets a reference to the given string and assigns it to the CustomLink2Url field.
func (o *SignInPageAllOfWidgetCustomizations) SetCustomLink2Url(v string) {
	o.CustomLink2Url = &v
}

// GetAuthenticatorPageCustomLinkLabel returns the AuthenticatorPageCustomLinkLabel field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetAuthenticatorPageCustomLinkLabel() string {
	if o == nil || o.AuthenticatorPageCustomLinkLabel == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatorPageCustomLinkLabel
}

// GetAuthenticatorPageCustomLinkLabelOk returns a tuple with the AuthenticatorPageCustomLinkLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetAuthenticatorPageCustomLinkLabelOk() (*string, bool) {
	if o == nil || o.AuthenticatorPageCustomLinkLabel == nil {
		return nil, false
	}
	return o.AuthenticatorPageCustomLinkLabel, true
}

// HasAuthenticatorPageCustomLinkLabel returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasAuthenticatorPageCustomLinkLabel() bool {
	if o != nil && o.AuthenticatorPageCustomLinkLabel != nil {
		return true
	}

	return false
}

// SetAuthenticatorPageCustomLinkLabel gets a reference to the given string and assigns it to the AuthenticatorPageCustomLinkLabel field.
func (o *SignInPageAllOfWidgetCustomizations) SetAuthenticatorPageCustomLinkLabel(v string) {
	o.AuthenticatorPageCustomLinkLabel = &v
}

// GetAuthenticatorPageCustomLinkUrl returns the AuthenticatorPageCustomLinkUrl field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetAuthenticatorPageCustomLinkUrl() string {
	if o == nil || o.AuthenticatorPageCustomLinkUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthenticatorPageCustomLinkUrl
}

// GetAuthenticatorPageCustomLinkUrlOk returns a tuple with the AuthenticatorPageCustomLinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetAuthenticatorPageCustomLinkUrlOk() (*string, bool) {
	if o == nil || o.AuthenticatorPageCustomLinkUrl == nil {
		return nil, false
	}
	return o.AuthenticatorPageCustomLinkUrl, true
}

// HasAuthenticatorPageCustomLinkUrl returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasAuthenticatorPageCustomLinkUrl() bool {
	if o != nil && o.AuthenticatorPageCustomLinkUrl != nil {
		return true
	}

	return false
}

// SetAuthenticatorPageCustomLinkUrl gets a reference to the given string and assigns it to the AuthenticatorPageCustomLinkUrl field.
func (o *SignInPageAllOfWidgetCustomizations) SetAuthenticatorPageCustomLinkUrl(v string) {
	o.AuthenticatorPageCustomLinkUrl = &v
}

// GetClassicRecoveryFlowEmailOrUsernameLabel returns the ClassicRecoveryFlowEmailOrUsernameLabel field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetClassicRecoveryFlowEmailOrUsernameLabel() string {
	if o == nil || o.ClassicRecoveryFlowEmailOrUsernameLabel == nil {
		var ret string
		return ret
	}
	return *o.ClassicRecoveryFlowEmailOrUsernameLabel
}

// GetClassicRecoveryFlowEmailOrUsernameLabelOk returns a tuple with the ClassicRecoveryFlowEmailOrUsernameLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetClassicRecoveryFlowEmailOrUsernameLabelOk() (*string, bool) {
	if o == nil || o.ClassicRecoveryFlowEmailOrUsernameLabel == nil {
		return nil, false
	}
	return o.ClassicRecoveryFlowEmailOrUsernameLabel, true
}

// HasClassicRecoveryFlowEmailOrUsernameLabel returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasClassicRecoveryFlowEmailOrUsernameLabel() bool {
	if o != nil && o.ClassicRecoveryFlowEmailOrUsernameLabel != nil {
		return true
	}

	return false
}

// SetClassicRecoveryFlowEmailOrUsernameLabel gets a reference to the given string and assigns it to the ClassicRecoveryFlowEmailOrUsernameLabel field.
func (o *SignInPageAllOfWidgetCustomizations) SetClassicRecoveryFlowEmailOrUsernameLabel(v string) {
	o.ClassicRecoveryFlowEmailOrUsernameLabel = &v
}

// GetWidgetGeneration returns the WidgetGeneration field value if set, zero value otherwise.
func (o *SignInPageAllOfWidgetCustomizations) GetWidgetGeneration() string {
	if o == nil || o.WidgetGeneration == nil {
		var ret string
		return ret
	}
	return *o.WidgetGeneration
}

// GetWidgetGenerationOk returns a tuple with the WidgetGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageAllOfWidgetCustomizations) GetWidgetGenerationOk() (*string, bool) {
	if o == nil || o.WidgetGeneration == nil {
		return nil, false
	}
	return o.WidgetGeneration, true
}

// HasWidgetGeneration returns a boolean if a field has been set.
func (o *SignInPageAllOfWidgetCustomizations) HasWidgetGeneration() bool {
	if o != nil && o.WidgetGeneration != nil {
		return true
	}

	return false
}

// SetWidgetGeneration gets a reference to the given string and assigns it to the WidgetGeneration field.
func (o *SignInPageAllOfWidgetCustomizations) SetWidgetGeneration(v string) {
	o.WidgetGeneration = &v
}

func (o SignInPageAllOfWidgetCustomizations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SignInLabel != nil {
		toSerialize["signInLabel"] = o.SignInLabel
	}
	if o.UsernameLabel != nil {
		toSerialize["usernameLabel"] = o.UsernameLabel
	}
	if o.UsernameInfoTip != nil {
		toSerialize["usernameInfoTip"] = o.UsernameInfoTip
	}
	if o.PasswordLabel != nil {
		toSerialize["passwordLabel"] = o.PasswordLabel
	}
	if o.PasswordInfoTip != nil {
		toSerialize["passwordInfoTip"] = o.PasswordInfoTip
	}
	if o.ShowPasswordVisibilityToggle != nil {
		toSerialize["showPasswordVisibilityToggle"] = o.ShowPasswordVisibilityToggle
	}
	if o.ShowUserIdentifier != nil {
		toSerialize["showUserIdentifier"] = o.ShowUserIdentifier
	}
	if o.ForgotPasswordLabel != nil {
		toSerialize["forgotPasswordLabel"] = o.ForgotPasswordLabel
	}
	if o.ForgotPasswordUrl != nil {
		toSerialize["forgotPasswordUrl"] = o.ForgotPasswordUrl
	}
	if o.UnlockAccountLabel != nil {
		toSerialize["unlockAccountLabel"] = o.UnlockAccountLabel
	}
	if o.UnlockAccountUrl != nil {
		toSerialize["unlockAccountUrl"] = o.UnlockAccountUrl
	}
	if o.HelpLabel != nil {
		toSerialize["helpLabel"] = o.HelpLabel
	}
	if o.HelpUrl != nil {
		toSerialize["helpUrl"] = o.HelpUrl
	}
	if o.CustomLink1Label != nil {
		toSerialize["customLink1Label"] = o.CustomLink1Label
	}
	if o.CustomLink1Url != nil {
		toSerialize["customLink1Url"] = o.CustomLink1Url
	}
	if o.CustomLink2Label != nil {
		toSerialize["customLink2Label"] = o.CustomLink2Label
	}
	if o.CustomLink2Url != nil {
		toSerialize["customLink2Url"] = o.CustomLink2Url
	}
	if o.AuthenticatorPageCustomLinkLabel != nil {
		toSerialize["authenticatorPageCustomLinkLabel"] = o.AuthenticatorPageCustomLinkLabel
	}
	if o.AuthenticatorPageCustomLinkUrl != nil {
		toSerialize["authenticatorPageCustomLinkUrl"] = o.AuthenticatorPageCustomLinkUrl
	}
	if o.ClassicRecoveryFlowEmailOrUsernameLabel != nil {
		toSerialize["classicRecoveryFlowEmailOrUsernameLabel"] = o.ClassicRecoveryFlowEmailOrUsernameLabel
	}
	if o.WidgetGeneration != nil {
		toSerialize["widgetGeneration"] = o.WidgetGeneration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SignInPageAllOfWidgetCustomizations) UnmarshalJSON(bytes []byte) (err error) {
	varSignInPageAllOfWidgetCustomizations := _SignInPageAllOfWidgetCustomizations{}

	err = json.Unmarshal(bytes, &varSignInPageAllOfWidgetCustomizations)
	if err == nil {
		*o = SignInPageAllOfWidgetCustomizations(varSignInPageAllOfWidgetCustomizations)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "signInLabel")
		delete(additionalProperties, "usernameLabel")
		delete(additionalProperties, "usernameInfoTip")
		delete(additionalProperties, "passwordLabel")
		delete(additionalProperties, "passwordInfoTip")
		delete(additionalProperties, "showPasswordVisibilityToggle")
		delete(additionalProperties, "showUserIdentifier")
		delete(additionalProperties, "forgotPasswordLabel")
		delete(additionalProperties, "forgotPasswordUrl")
		delete(additionalProperties, "unlockAccountLabel")
		delete(additionalProperties, "unlockAccountUrl")
		delete(additionalProperties, "helpLabel")
		delete(additionalProperties, "helpUrl")
		delete(additionalProperties, "customLink1Label")
		delete(additionalProperties, "customLink1Url")
		delete(additionalProperties, "customLink2Label")
		delete(additionalProperties, "customLink2Url")
		delete(additionalProperties, "authenticatorPageCustomLinkLabel")
		delete(additionalProperties, "authenticatorPageCustomLinkUrl")
		delete(additionalProperties, "classicRecoveryFlowEmailOrUsernameLabel")
		delete(additionalProperties, "widgetGeneration")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSignInPageAllOfWidgetCustomizations struct {
	value *SignInPageAllOfWidgetCustomizations
	isSet bool
}

func (v NullableSignInPageAllOfWidgetCustomizations) Get() *SignInPageAllOfWidgetCustomizations {
	return v.value
}

func (v *NullableSignInPageAllOfWidgetCustomizations) Set(val *SignInPageAllOfWidgetCustomizations) {
	v.value = val
	v.isSet = true
}

func (v NullableSignInPageAllOfWidgetCustomizations) IsSet() bool {
	return v.isSet
}

func (v *NullableSignInPageAllOfWidgetCustomizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignInPageAllOfWidgetCustomizations(val *SignInPageAllOfWidgetCustomizations) *NullableSignInPageAllOfWidgetCustomizations {
	return &NullableSignInPageAllOfWidgetCustomizations{value: val, isSet: true}
}

func (v NullableSignInPageAllOfWidgetCustomizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignInPageAllOfWidgetCustomizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

