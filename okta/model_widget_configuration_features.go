/*
Okta Admin Management API

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
)

// checks if the WidgetConfigurationFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetConfigurationFeatures{}

// WidgetConfigurationFeatures struct for WidgetConfigurationFeatures
type WidgetConfigurationFeatures struct {
	Router                               *bool `json:"router,omitempty"`
	SecurityImage                        *bool `json:"securityImage,omitempty"`
	RememberMe                           *bool `json:"rememberMe,omitempty"`
	AutoPush                             *bool `json:"autoPush,omitempty"`
	SmsRecovery                          *bool `json:"smsRecovery,omitempty"`
	CallRecovery                         *bool `json:"callRecovery,omitempty"`
	EmailRecovery                        *bool `json:"emailRecovery,omitempty"`
	Webauthn                             *bool `json:"webauthn,omitempty"`
	SelfServiceUnlock                    *bool `json:"selfServiceUnlock,omitempty"`
	MultiOptionalFactorEnroll            *bool `json:"multiOptionalFactorEnroll,omitempty"`
	DeviceFingerprinting                 *bool `json:"deviceFingerprinting,omitempty"`
	HideSignOutLinkInMFA                 *bool `json:"hideSignOutLinkInMFA,omitempty"`
	HideBackToSignInForReset             *bool `json:"hideBackToSignInForReset,omitempty"`
	CustomExpiredPassword                *bool `json:"customExpiredPassword,omitempty"`
	Registration                         *bool `json:"registration,omitempty"`
	IdpDiscovery                         *bool `json:"idpDiscovery,omitempty"`
	PasswordlessAuth                     *bool `json:"passwordlessAuth,omitempty"`
	ShowPasswordToggleOnSignInPage       *bool `json:"showPasswordToggleOnSignInPage,omitempty"`
	TrackTypingPattern                   *bool `json:"trackTypingPattern,omitempty"`
	RedirectByFormSubmit                 *bool `json:"redirectByFormSubmit,omitempty"`
	SameDeviceOVEnrollmentEnabled        *bool `json:"sameDeviceOVEnrollmentEnabled,omitempty"`
	UseDeviceFingerprintForSecurityImage *bool `json:"useDeviceFingerprintForSecurityImage,omitempty"`
	ShowPasswordRequirementsAsHtmlList   *bool `json:"showPasswordRequirementsAsHtmlList,omitempty"`
	MfaOnlyFlow                          *bool `json:"mfaOnlyFlow,omitempty"`
}

// NewWidgetConfigurationFeatures instantiates a new WidgetConfigurationFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetConfigurationFeatures() *WidgetConfigurationFeatures {
	this := WidgetConfigurationFeatures{}
	return &this
}

// NewWidgetConfigurationFeaturesWithDefaults instantiates a new WidgetConfigurationFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetConfigurationFeaturesWithDefaults() *WidgetConfigurationFeatures {
	this := WidgetConfigurationFeatures{}
	return &this
}

// GetRouter returns the Router field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetRouter() bool {
	if o == nil || IsNil(o.Router) {
		var ret bool
		return ret
	}
	return *o.Router
}

// GetRouterOk returns a tuple with the Router field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetRouterOk() (*bool, bool) {
	if o == nil || IsNil(o.Router) {
		return nil, false
	}
	return o.Router, true
}

// HasRouter returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasRouter() bool {
	if o != nil && !IsNil(o.Router) {
		return true
	}

	return false
}

// SetRouter gets a reference to the given bool and assigns it to the Router field.
func (o *WidgetConfigurationFeatures) SetRouter(v bool) {
	o.Router = &v
}

// GetSecurityImage returns the SecurityImage field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetSecurityImage() bool {
	if o == nil || IsNil(o.SecurityImage) {
		var ret bool
		return ret
	}
	return *o.SecurityImage
}

// GetSecurityImageOk returns a tuple with the SecurityImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetSecurityImageOk() (*bool, bool) {
	if o == nil || IsNil(o.SecurityImage) {
		return nil, false
	}
	return o.SecurityImage, true
}

// HasSecurityImage returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasSecurityImage() bool {
	if o != nil && !IsNil(o.SecurityImage) {
		return true
	}

	return false
}

// SetSecurityImage gets a reference to the given bool and assigns it to the SecurityImage field.
func (o *WidgetConfigurationFeatures) SetSecurityImage(v bool) {
	o.SecurityImage = &v
}

// GetRememberMe returns the RememberMe field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetRememberMe() bool {
	if o == nil || IsNil(o.RememberMe) {
		var ret bool
		return ret
	}
	return *o.RememberMe
}

// GetRememberMeOk returns a tuple with the RememberMe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetRememberMeOk() (*bool, bool) {
	if o == nil || IsNil(o.RememberMe) {
		return nil, false
	}
	return o.RememberMe, true
}

// HasRememberMe returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasRememberMe() bool {
	if o != nil && !IsNil(o.RememberMe) {
		return true
	}

	return false
}

// SetRememberMe gets a reference to the given bool and assigns it to the RememberMe field.
func (o *WidgetConfigurationFeatures) SetRememberMe(v bool) {
	o.RememberMe = &v
}

// GetAutoPush returns the AutoPush field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetAutoPush() bool {
	if o == nil || IsNil(o.AutoPush) {
		var ret bool
		return ret
	}
	return *o.AutoPush
}

// GetAutoPushOk returns a tuple with the AutoPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetAutoPushOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPush) {
		return nil, false
	}
	return o.AutoPush, true
}

// HasAutoPush returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasAutoPush() bool {
	if o != nil && !IsNil(o.AutoPush) {
		return true
	}

	return false
}

// SetAutoPush gets a reference to the given bool and assigns it to the AutoPush field.
func (o *WidgetConfigurationFeatures) SetAutoPush(v bool) {
	o.AutoPush = &v
}

// GetSmsRecovery returns the SmsRecovery field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetSmsRecovery() bool {
	if o == nil || IsNil(o.SmsRecovery) {
		var ret bool
		return ret
	}
	return *o.SmsRecovery
}

// GetSmsRecoveryOk returns a tuple with the SmsRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetSmsRecoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.SmsRecovery) {
		return nil, false
	}
	return o.SmsRecovery, true
}

// HasSmsRecovery returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasSmsRecovery() bool {
	if o != nil && !IsNil(o.SmsRecovery) {
		return true
	}

	return false
}

// SetSmsRecovery gets a reference to the given bool and assigns it to the SmsRecovery field.
func (o *WidgetConfigurationFeatures) SetSmsRecovery(v bool) {
	o.SmsRecovery = &v
}

// GetCallRecovery returns the CallRecovery field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetCallRecovery() bool {
	if o == nil || IsNil(o.CallRecovery) {
		var ret bool
		return ret
	}
	return *o.CallRecovery
}

// GetCallRecoveryOk returns a tuple with the CallRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetCallRecoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.CallRecovery) {
		return nil, false
	}
	return o.CallRecovery, true
}

// HasCallRecovery returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasCallRecovery() bool {
	if o != nil && !IsNil(o.CallRecovery) {
		return true
	}

	return false
}

// SetCallRecovery gets a reference to the given bool and assigns it to the CallRecovery field.
func (o *WidgetConfigurationFeatures) SetCallRecovery(v bool) {
	o.CallRecovery = &v
}

// GetEmailRecovery returns the EmailRecovery field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetEmailRecovery() bool {
	if o == nil || IsNil(o.EmailRecovery) {
		var ret bool
		return ret
	}
	return *o.EmailRecovery
}

// GetEmailRecoveryOk returns a tuple with the EmailRecovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetEmailRecoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailRecovery) {
		return nil, false
	}
	return o.EmailRecovery, true
}

// HasEmailRecovery returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasEmailRecovery() bool {
	if o != nil && !IsNil(o.EmailRecovery) {
		return true
	}

	return false
}

// SetEmailRecovery gets a reference to the given bool and assigns it to the EmailRecovery field.
func (o *WidgetConfigurationFeatures) SetEmailRecovery(v bool) {
	o.EmailRecovery = &v
}

// GetWebauthn returns the Webauthn field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetWebauthn() bool {
	if o == nil || IsNil(o.Webauthn) {
		var ret bool
		return ret
	}
	return *o.Webauthn
}

// GetWebauthnOk returns a tuple with the Webauthn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetWebauthnOk() (*bool, bool) {
	if o == nil || IsNil(o.Webauthn) {
		return nil, false
	}
	return o.Webauthn, true
}

// HasWebauthn returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasWebauthn() bool {
	if o != nil && !IsNil(o.Webauthn) {
		return true
	}

	return false
}

// SetWebauthn gets a reference to the given bool and assigns it to the Webauthn field.
func (o *WidgetConfigurationFeatures) SetWebauthn(v bool) {
	o.Webauthn = &v
}

// GetSelfServiceUnlock returns the SelfServiceUnlock field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetSelfServiceUnlock() bool {
	if o == nil || IsNil(o.SelfServiceUnlock) {
		var ret bool
		return ret
	}
	return *o.SelfServiceUnlock
}

// GetSelfServiceUnlockOk returns a tuple with the SelfServiceUnlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetSelfServiceUnlockOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfServiceUnlock) {
		return nil, false
	}
	return o.SelfServiceUnlock, true
}

// HasSelfServiceUnlock returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasSelfServiceUnlock() bool {
	if o != nil && !IsNil(o.SelfServiceUnlock) {
		return true
	}

	return false
}

// SetSelfServiceUnlock gets a reference to the given bool and assigns it to the SelfServiceUnlock field.
func (o *WidgetConfigurationFeatures) SetSelfServiceUnlock(v bool) {
	o.SelfServiceUnlock = &v
}

// GetMultiOptionalFactorEnroll returns the MultiOptionalFactorEnroll field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetMultiOptionalFactorEnroll() bool {
	if o == nil || IsNil(o.MultiOptionalFactorEnroll) {
		var ret bool
		return ret
	}
	return *o.MultiOptionalFactorEnroll
}

// GetMultiOptionalFactorEnrollOk returns a tuple with the MultiOptionalFactorEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetMultiOptionalFactorEnrollOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiOptionalFactorEnroll) {
		return nil, false
	}
	return o.MultiOptionalFactorEnroll, true
}

// HasMultiOptionalFactorEnroll returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasMultiOptionalFactorEnroll() bool {
	if o != nil && !IsNil(o.MultiOptionalFactorEnroll) {
		return true
	}

	return false
}

// SetMultiOptionalFactorEnroll gets a reference to the given bool and assigns it to the MultiOptionalFactorEnroll field.
func (o *WidgetConfigurationFeatures) SetMultiOptionalFactorEnroll(v bool) {
	o.MultiOptionalFactorEnroll = &v
}

// GetDeviceFingerprinting returns the DeviceFingerprinting field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetDeviceFingerprinting() bool {
	if o == nil || IsNil(o.DeviceFingerprinting) {
		var ret bool
		return ret
	}
	return *o.DeviceFingerprinting
}

// GetDeviceFingerprintingOk returns a tuple with the DeviceFingerprinting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetDeviceFingerprintingOk() (*bool, bool) {
	if o == nil || IsNil(o.DeviceFingerprinting) {
		return nil, false
	}
	return o.DeviceFingerprinting, true
}

// HasDeviceFingerprinting returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasDeviceFingerprinting() bool {
	if o != nil && !IsNil(o.DeviceFingerprinting) {
		return true
	}

	return false
}

// SetDeviceFingerprinting gets a reference to the given bool and assigns it to the DeviceFingerprinting field.
func (o *WidgetConfigurationFeatures) SetDeviceFingerprinting(v bool) {
	o.DeviceFingerprinting = &v
}

// GetHideSignOutLinkInMFA returns the HideSignOutLinkInMFA field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetHideSignOutLinkInMFA() bool {
	if o == nil || IsNil(o.HideSignOutLinkInMFA) {
		var ret bool
		return ret
	}
	return *o.HideSignOutLinkInMFA
}

// GetHideSignOutLinkInMFAOk returns a tuple with the HideSignOutLinkInMFA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetHideSignOutLinkInMFAOk() (*bool, bool) {
	if o == nil || IsNil(o.HideSignOutLinkInMFA) {
		return nil, false
	}
	return o.HideSignOutLinkInMFA, true
}

// HasHideSignOutLinkInMFA returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasHideSignOutLinkInMFA() bool {
	if o != nil && !IsNil(o.HideSignOutLinkInMFA) {
		return true
	}

	return false
}

// SetHideSignOutLinkInMFA gets a reference to the given bool and assigns it to the HideSignOutLinkInMFA field.
func (o *WidgetConfigurationFeatures) SetHideSignOutLinkInMFA(v bool) {
	o.HideSignOutLinkInMFA = &v
}

// GetHideBackToSignInForReset returns the HideBackToSignInForReset field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetHideBackToSignInForReset() bool {
	if o == nil || IsNil(o.HideBackToSignInForReset) {
		var ret bool
		return ret
	}
	return *o.HideBackToSignInForReset
}

// GetHideBackToSignInForResetOk returns a tuple with the HideBackToSignInForReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetHideBackToSignInForResetOk() (*bool, bool) {
	if o == nil || IsNil(o.HideBackToSignInForReset) {
		return nil, false
	}
	return o.HideBackToSignInForReset, true
}

// HasHideBackToSignInForReset returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasHideBackToSignInForReset() bool {
	if o != nil && !IsNil(o.HideBackToSignInForReset) {
		return true
	}

	return false
}

// SetHideBackToSignInForReset gets a reference to the given bool and assigns it to the HideBackToSignInForReset field.
func (o *WidgetConfigurationFeatures) SetHideBackToSignInForReset(v bool) {
	o.HideBackToSignInForReset = &v
}

// GetCustomExpiredPassword returns the CustomExpiredPassword field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetCustomExpiredPassword() bool {
	if o == nil || IsNil(o.CustomExpiredPassword) {
		var ret bool
		return ret
	}
	return *o.CustomExpiredPassword
}

// GetCustomExpiredPasswordOk returns a tuple with the CustomExpiredPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetCustomExpiredPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.CustomExpiredPassword) {
		return nil, false
	}
	return o.CustomExpiredPassword, true
}

// HasCustomExpiredPassword returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasCustomExpiredPassword() bool {
	if o != nil && !IsNil(o.CustomExpiredPassword) {
		return true
	}

	return false
}

// SetCustomExpiredPassword gets a reference to the given bool and assigns it to the CustomExpiredPassword field.
func (o *WidgetConfigurationFeatures) SetCustomExpiredPassword(v bool) {
	o.CustomExpiredPassword = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetRegistration() bool {
	if o == nil || IsNil(o.Registration) {
		var ret bool
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetRegistrationOk() (*bool, bool) {
	if o == nil || IsNil(o.Registration) {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasRegistration() bool {
	if o != nil && !IsNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given bool and assigns it to the Registration field.
func (o *WidgetConfigurationFeatures) SetRegistration(v bool) {
	o.Registration = &v
}

// GetIdpDiscovery returns the IdpDiscovery field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetIdpDiscovery() bool {
	if o == nil || IsNil(o.IdpDiscovery) {
		var ret bool
		return ret
	}
	return *o.IdpDiscovery
}

// GetIdpDiscoveryOk returns a tuple with the IdpDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetIdpDiscoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.IdpDiscovery) {
		return nil, false
	}
	return o.IdpDiscovery, true
}

// HasIdpDiscovery returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasIdpDiscovery() bool {
	if o != nil && !IsNil(o.IdpDiscovery) {
		return true
	}

	return false
}

// SetIdpDiscovery gets a reference to the given bool and assigns it to the IdpDiscovery field.
func (o *WidgetConfigurationFeatures) SetIdpDiscovery(v bool) {
	o.IdpDiscovery = &v
}

// GetPasswordlessAuth returns the PasswordlessAuth field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetPasswordlessAuth() bool {
	if o == nil || IsNil(o.PasswordlessAuth) {
		var ret bool
		return ret
	}
	return *o.PasswordlessAuth
}

// GetPasswordlessAuthOk returns a tuple with the PasswordlessAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetPasswordlessAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.PasswordlessAuth) {
		return nil, false
	}
	return o.PasswordlessAuth, true
}

// HasPasswordlessAuth returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasPasswordlessAuth() bool {
	if o != nil && !IsNil(o.PasswordlessAuth) {
		return true
	}

	return false
}

// SetPasswordlessAuth gets a reference to the given bool and assigns it to the PasswordlessAuth field.
func (o *WidgetConfigurationFeatures) SetPasswordlessAuth(v bool) {
	o.PasswordlessAuth = &v
}

// GetShowPasswordToggleOnSignInPage returns the ShowPasswordToggleOnSignInPage field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetShowPasswordToggleOnSignInPage() bool {
	if o == nil || IsNil(o.ShowPasswordToggleOnSignInPage) {
		var ret bool
		return ret
	}
	return *o.ShowPasswordToggleOnSignInPage
}

// GetShowPasswordToggleOnSignInPageOk returns a tuple with the ShowPasswordToggleOnSignInPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetShowPasswordToggleOnSignInPageOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowPasswordToggleOnSignInPage) {
		return nil, false
	}
	return o.ShowPasswordToggleOnSignInPage, true
}

// HasShowPasswordToggleOnSignInPage returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasShowPasswordToggleOnSignInPage() bool {
	if o != nil && !IsNil(o.ShowPasswordToggleOnSignInPage) {
		return true
	}

	return false
}

// SetShowPasswordToggleOnSignInPage gets a reference to the given bool and assigns it to the ShowPasswordToggleOnSignInPage field.
func (o *WidgetConfigurationFeatures) SetShowPasswordToggleOnSignInPage(v bool) {
	o.ShowPasswordToggleOnSignInPage = &v
}

// GetTrackTypingPattern returns the TrackTypingPattern field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetTrackTypingPattern() bool {
	if o == nil || IsNil(o.TrackTypingPattern) {
		var ret bool
		return ret
	}
	return *o.TrackTypingPattern
}

// GetTrackTypingPatternOk returns a tuple with the TrackTypingPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetTrackTypingPatternOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackTypingPattern) {
		return nil, false
	}
	return o.TrackTypingPattern, true
}

// HasTrackTypingPattern returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasTrackTypingPattern() bool {
	if o != nil && !IsNil(o.TrackTypingPattern) {
		return true
	}

	return false
}

// SetTrackTypingPattern gets a reference to the given bool and assigns it to the TrackTypingPattern field.
func (o *WidgetConfigurationFeatures) SetTrackTypingPattern(v bool) {
	o.TrackTypingPattern = &v
}

// GetRedirectByFormSubmit returns the RedirectByFormSubmit field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetRedirectByFormSubmit() bool {
	if o == nil || IsNil(o.RedirectByFormSubmit) {
		var ret bool
		return ret
	}
	return *o.RedirectByFormSubmit
}

// GetRedirectByFormSubmitOk returns a tuple with the RedirectByFormSubmit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetRedirectByFormSubmitOk() (*bool, bool) {
	if o == nil || IsNil(o.RedirectByFormSubmit) {
		return nil, false
	}
	return o.RedirectByFormSubmit, true
}

// HasRedirectByFormSubmit returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasRedirectByFormSubmit() bool {
	if o != nil && !IsNil(o.RedirectByFormSubmit) {
		return true
	}

	return false
}

// SetRedirectByFormSubmit gets a reference to the given bool and assigns it to the RedirectByFormSubmit field.
func (o *WidgetConfigurationFeatures) SetRedirectByFormSubmit(v bool) {
	o.RedirectByFormSubmit = &v
}

// GetSameDeviceOVEnrollmentEnabled returns the SameDeviceOVEnrollmentEnabled field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetSameDeviceOVEnrollmentEnabled() bool {
	if o == nil || IsNil(o.SameDeviceOVEnrollmentEnabled) {
		var ret bool
		return ret
	}
	return *o.SameDeviceOVEnrollmentEnabled
}

// GetSameDeviceOVEnrollmentEnabledOk returns a tuple with the SameDeviceOVEnrollmentEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetSameDeviceOVEnrollmentEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SameDeviceOVEnrollmentEnabled) {
		return nil, false
	}
	return o.SameDeviceOVEnrollmentEnabled, true
}

// HasSameDeviceOVEnrollmentEnabled returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasSameDeviceOVEnrollmentEnabled() bool {
	if o != nil && !IsNil(o.SameDeviceOVEnrollmentEnabled) {
		return true
	}

	return false
}

// SetSameDeviceOVEnrollmentEnabled gets a reference to the given bool and assigns it to the SameDeviceOVEnrollmentEnabled field.
func (o *WidgetConfigurationFeatures) SetSameDeviceOVEnrollmentEnabled(v bool) {
	o.SameDeviceOVEnrollmentEnabled = &v
}

// GetUseDeviceFingerprintForSecurityImage returns the UseDeviceFingerprintForSecurityImage field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetUseDeviceFingerprintForSecurityImage() bool {
	if o == nil || IsNil(o.UseDeviceFingerprintForSecurityImage) {
		var ret bool
		return ret
	}
	return *o.UseDeviceFingerprintForSecurityImage
}

// GetUseDeviceFingerprintForSecurityImageOk returns a tuple with the UseDeviceFingerprintForSecurityImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetUseDeviceFingerprintForSecurityImageOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDeviceFingerprintForSecurityImage) {
		return nil, false
	}
	return o.UseDeviceFingerprintForSecurityImage, true
}

// HasUseDeviceFingerprintForSecurityImage returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasUseDeviceFingerprintForSecurityImage() bool {
	if o != nil && !IsNil(o.UseDeviceFingerprintForSecurityImage) {
		return true
	}

	return false
}

// SetUseDeviceFingerprintForSecurityImage gets a reference to the given bool and assigns it to the UseDeviceFingerprintForSecurityImage field.
func (o *WidgetConfigurationFeatures) SetUseDeviceFingerprintForSecurityImage(v bool) {
	o.UseDeviceFingerprintForSecurityImage = &v
}

// GetShowPasswordRequirementsAsHtmlList returns the ShowPasswordRequirementsAsHtmlList field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetShowPasswordRequirementsAsHtmlList() bool {
	if o == nil || IsNil(o.ShowPasswordRequirementsAsHtmlList) {
		var ret bool
		return ret
	}
	return *o.ShowPasswordRequirementsAsHtmlList
}

// GetShowPasswordRequirementsAsHtmlListOk returns a tuple with the ShowPasswordRequirementsAsHtmlList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetShowPasswordRequirementsAsHtmlListOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowPasswordRequirementsAsHtmlList) {
		return nil, false
	}
	return o.ShowPasswordRequirementsAsHtmlList, true
}

// HasShowPasswordRequirementsAsHtmlList returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasShowPasswordRequirementsAsHtmlList() bool {
	if o != nil && !IsNil(o.ShowPasswordRequirementsAsHtmlList) {
		return true
	}

	return false
}

// SetShowPasswordRequirementsAsHtmlList gets a reference to the given bool and assigns it to the ShowPasswordRequirementsAsHtmlList field.
func (o *WidgetConfigurationFeatures) SetShowPasswordRequirementsAsHtmlList(v bool) {
	o.ShowPasswordRequirementsAsHtmlList = &v
}

// GetMfaOnlyFlow returns the MfaOnlyFlow field value if set, zero value otherwise.
func (o *WidgetConfigurationFeatures) GetMfaOnlyFlow() bool {
	if o == nil || IsNil(o.MfaOnlyFlow) {
		var ret bool
		return ret
	}
	return *o.MfaOnlyFlow
}

// GetMfaOnlyFlowOk returns a tuple with the MfaOnlyFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetConfigurationFeatures) GetMfaOnlyFlowOk() (*bool, bool) {
	if o == nil || IsNil(o.MfaOnlyFlow) {
		return nil, false
	}
	return o.MfaOnlyFlow, true
}

// HasMfaOnlyFlow returns a boolean if a field has been set.
func (o *WidgetConfigurationFeatures) HasMfaOnlyFlow() bool {
	if o != nil && !IsNil(o.MfaOnlyFlow) {
		return true
	}

	return false
}

// SetMfaOnlyFlow gets a reference to the given bool and assigns it to the MfaOnlyFlow field.
func (o *WidgetConfigurationFeatures) SetMfaOnlyFlow(v bool) {
	o.MfaOnlyFlow = &v
}

func (o WidgetConfigurationFeatures) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetConfigurationFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Router) {
		toSerialize["router"] = o.Router
	}
	if !IsNil(o.SecurityImage) {
		toSerialize["securityImage"] = o.SecurityImage
	}
	if !IsNil(o.RememberMe) {
		toSerialize["rememberMe"] = o.RememberMe
	}
	if !IsNil(o.AutoPush) {
		toSerialize["autoPush"] = o.AutoPush
	}
	if !IsNil(o.SmsRecovery) {
		toSerialize["smsRecovery"] = o.SmsRecovery
	}
	if !IsNil(o.CallRecovery) {
		toSerialize["callRecovery"] = o.CallRecovery
	}
	if !IsNil(o.EmailRecovery) {
		toSerialize["emailRecovery"] = o.EmailRecovery
	}
	if !IsNil(o.Webauthn) {
		toSerialize["webauthn"] = o.Webauthn
	}
	if !IsNil(o.SelfServiceUnlock) {
		toSerialize["selfServiceUnlock"] = o.SelfServiceUnlock
	}
	if !IsNil(o.MultiOptionalFactorEnroll) {
		toSerialize["multiOptionalFactorEnroll"] = o.MultiOptionalFactorEnroll
	}
	if !IsNil(o.DeviceFingerprinting) {
		toSerialize["deviceFingerprinting"] = o.DeviceFingerprinting
	}
	if !IsNil(o.HideSignOutLinkInMFA) {
		toSerialize["hideSignOutLinkInMFA"] = o.HideSignOutLinkInMFA
	}
	if !IsNil(o.HideBackToSignInForReset) {
		toSerialize["hideBackToSignInForReset"] = o.HideBackToSignInForReset
	}
	if !IsNil(o.CustomExpiredPassword) {
		toSerialize["customExpiredPassword"] = o.CustomExpiredPassword
	}
	if !IsNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	if !IsNil(o.IdpDiscovery) {
		toSerialize["idpDiscovery"] = o.IdpDiscovery
	}
	if !IsNil(o.PasswordlessAuth) {
		toSerialize["passwordlessAuth"] = o.PasswordlessAuth
	}
	if !IsNil(o.ShowPasswordToggleOnSignInPage) {
		toSerialize["showPasswordToggleOnSignInPage"] = o.ShowPasswordToggleOnSignInPage
	}
	if !IsNil(o.TrackTypingPattern) {
		toSerialize["trackTypingPattern"] = o.TrackTypingPattern
	}
	if !IsNil(o.RedirectByFormSubmit) {
		toSerialize["redirectByFormSubmit"] = o.RedirectByFormSubmit
	}
	if !IsNil(o.SameDeviceOVEnrollmentEnabled) {
		toSerialize["sameDeviceOVEnrollmentEnabled"] = o.SameDeviceOVEnrollmentEnabled
	}
	if !IsNil(o.UseDeviceFingerprintForSecurityImage) {
		toSerialize["useDeviceFingerprintForSecurityImage"] = o.UseDeviceFingerprintForSecurityImage
	}
	if !IsNil(o.ShowPasswordRequirementsAsHtmlList) {
		toSerialize["showPasswordRequirementsAsHtmlList"] = o.ShowPasswordRequirementsAsHtmlList
	}
	if !IsNil(o.MfaOnlyFlow) {
		toSerialize["mfaOnlyFlow"] = o.MfaOnlyFlow
	}
	return toSerialize, nil
}

type NullableWidgetConfigurationFeatures struct {
	value *WidgetConfigurationFeatures
	isSet bool
}

func (v NullableWidgetConfigurationFeatures) Get() *WidgetConfigurationFeatures {
	return v.value
}

func (v *NullableWidgetConfigurationFeatures) Set(val *WidgetConfigurationFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetConfigurationFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetConfigurationFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetConfigurationFeatures(val *WidgetConfigurationFeatures) *NullableWidgetConfigurationFeatures {
	return &NullableWidgetConfigurationFeatures{value: val, isSet: true}
}

func (v NullableWidgetConfigurationFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetConfigurationFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
