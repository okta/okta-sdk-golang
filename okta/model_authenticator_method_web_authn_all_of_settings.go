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

// checks if the AuthenticatorMethodWebAuthnAllOfSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorMethodWebAuthnAllOfSettings{}

// AuthenticatorMethodWebAuthnAllOfSettings The settings for the WebAuthn authenticator method
type AuthenticatorMethodWebAuthnAllOfSettings struct {
	// The FIDO2 Authenticator Attestation Global Unique Identifiers (AAGUID) groups available to the WebAuthn authenticator
	AaguidGroups []AAGUIDGroupObject `json:"aaguidGroups,omitempty"`
	// User verification settings for enrollment.  This setting controls the user verification requirement during the enrollment of a new credential. It determines whether the authenticator requires verification when a user is registering their device or credential.
	UserVerification *string `json:"userVerification,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>User verification settings for verification. This setting controls the user verification requirement during authentication (verification). It determines whether the authenticator requires user verification when a user signs in with an already-registered credential.  For verification, the value defaults to `PREFERRED`, unless the enrollment setting is `REQUIRED`. If the enrollment setting is `REQUIRED` for the authenticator, then the verification setting is also implicitly `REQUIRED`.  > **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id=ext_Manage_Early_Access_features).
	UserVerificationForVerify *string `json:"userVerificationForVerify,omitempty"`
	// Method attachment
	Attachment *string       `json:"attachment,omitempty"`
	RpId       *WebAuthnRpId `json:"rpId,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Enables the passkeys autofill UI to display available WebAuthn discoverable credentials (\"resident key\") from the Sign-In Widget username field
	EnableAutofillUI *bool `json:"enableAutofillUI,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Resident key requirement setting. Okta recommends using only `REQUIRED` or `DISCOURAGED` to make the requirement preference explicit. Using `PREFERRED` can sometimes lead to unpredictable behavior depending on the client platform and authenticator capabilities.  > **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id=ext_Manage_Early_Access_features).
	ResidentKeyRequirement *string `json:"residentKeyRequirement,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Indicates if the **Sign in with a Passkey** button on the Sign-In Widget is shown.   > **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id=ext_Manage_Early_Access_features).
	ShowSignInWithAPasskeyButton *bool `json:"showSignInWithAPasskeyButton,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Indicates whether certificate-based attestation validation is enabled. When enabled, the authenticator's attestation certificate is validated against known root certificates (custom AAGUIDs with associated certificates or the [FIDO Metadata Service](https://fidoalliance.org/metadata/)) to ensure its validity.  > **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id=ext_Manage_Early_Access_features).
	CertBasedAttestationValidation *bool `json:"certBasedAttestationValidation,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Indicates whether the authenticator is required to store the private key on a hardware component  > **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id=ext_Manage_Early_Access_features).
	HardwareProtected *bool `json:"hardwareProtected,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Indicates whether the authenticator is required to be [Federal Information Processing Standards (FIPS)](https://csrc.nist.gov/glossary/term/federal_information_processing_standard) compliant  > **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id=ext_Manage_Early_Access_features).
	FipsCompliant *bool `json:"fipsCompliant,omitempty"`
	// <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Indicates whether syncable passkeys are allowed. When enabled, users can register passkeys that are synchronized across their devices by using platform-specific mechanisms (such as iCloud Keychain for Apple devices or Google Password Manager for Android devices).  > **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id=ext_Manage_Early_Access_features).
	AllowSyncablePasskeys *bool `json:"allowSyncablePasskeys,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _AuthenticatorMethodWebAuthnAllOfSettings AuthenticatorMethodWebAuthnAllOfSettings

// NewAuthenticatorMethodWebAuthnAllOfSettings instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorMethodWebAuthnAllOfSettings() *AuthenticatorMethodWebAuthnAllOfSettings {
	this := AuthenticatorMethodWebAuthnAllOfSettings{}
	var enableAutofillUI bool = false
	this.EnableAutofillUI = &enableAutofillUI
	var showSignInWithAPasskeyButton bool = false
	this.ShowSignInWithAPasskeyButton = &showSignInWithAPasskeyButton
	var certBasedAttestationValidation bool = false
	this.CertBasedAttestationValidation = &certBasedAttestationValidation
	var hardwareProtected bool = false
	this.HardwareProtected = &hardwareProtected
	var fipsCompliant bool = false
	this.FipsCompliant = &fipsCompliant
	var allowSyncablePasskeys bool = true
	this.AllowSyncablePasskeys = &allowSyncablePasskeys
	return &this
}

// NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults() *AuthenticatorMethodWebAuthnAllOfSettings {
	this := AuthenticatorMethodWebAuthnAllOfSettings{}
	var enableAutofillUI bool = false
	this.EnableAutofillUI = &enableAutofillUI
	var showSignInWithAPasskeyButton bool = false
	this.ShowSignInWithAPasskeyButton = &showSignInWithAPasskeyButton
	var certBasedAttestationValidation bool = false
	this.CertBasedAttestationValidation = &certBasedAttestationValidation
	var hardwareProtected bool = false
	this.HardwareProtected = &hardwareProtected
	var fipsCompliant bool = false
	this.FipsCompliant = &fipsCompliant
	var allowSyncablePasskeys bool = true
	this.AllowSyncablePasskeys = &allowSyncablePasskeys
	return &this
}

// GetAaguidGroups returns the AaguidGroups field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAaguidGroups() []AAGUIDGroupObject {
	if o == nil || IsNil(o.AaguidGroups) {
		var ret []AAGUIDGroupObject
		return ret
	}
	return o.AaguidGroups
}

// GetAaguidGroupsOk returns a tuple with the AaguidGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAaguidGroupsOk() ([]AAGUIDGroupObject, bool) {
	if o == nil || IsNil(o.AaguidGroups) {
		return nil, false
	}
	return o.AaguidGroups, true
}

// HasAaguidGroups returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAaguidGroups() bool {
	if o != nil && !IsNil(o.AaguidGroups) {
		return true
	}

	return false
}

// SetAaguidGroups gets a reference to the given []AAGUIDGroupObject and assigns it to the AaguidGroups field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAaguidGroups(v []AAGUIDGroupObject) {
	o.AaguidGroups = v
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerification() string {
	if o == nil || IsNil(o.UserVerification) {
		var ret string
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerificationOk() (*string, bool) {
	if o == nil || IsNil(o.UserVerification) {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasUserVerification() bool {
	if o != nil && !IsNil(o.UserVerification) {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given string and assigns it to the UserVerification field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetUserVerification(v string) {
	o.UserVerification = &v
}

// GetUserVerificationForVerify returns the UserVerificationForVerify field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerificationForVerify() string {
	if o == nil || IsNil(o.UserVerificationForVerify) {
		var ret string
		return ret
	}
	return *o.UserVerificationForVerify
}

// GetUserVerificationForVerifyOk returns a tuple with the UserVerificationForVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerificationForVerifyOk() (*string, bool) {
	if o == nil || IsNil(o.UserVerificationForVerify) {
		return nil, false
	}
	return o.UserVerificationForVerify, true
}

// HasUserVerificationForVerify returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasUserVerificationForVerify() bool {
	if o != nil && !IsNil(o.UserVerificationForVerify) {
		return true
	}

	return false
}

// SetUserVerificationForVerify gets a reference to the given string and assigns it to the UserVerificationForVerify field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetUserVerificationForVerify(v string) {
	o.UserVerificationForVerify = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAttachment() string {
	if o == nil || IsNil(o.Attachment) {
		var ret string
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAttachmentOk() (*string, bool) {
	if o == nil || IsNil(o.Attachment) {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAttachment() bool {
	if o != nil && !IsNil(o.Attachment) {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given string and assigns it to the Attachment field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAttachment(v string) {
	o.Attachment = &v
}

// GetRpId returns the RpId field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetRpId() WebAuthnRpId {
	if o == nil || IsNil(o.RpId) {
		var ret WebAuthnRpId
		return ret
	}
	return *o.RpId
}

// GetRpIdOk returns a tuple with the RpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetRpIdOk() (*WebAuthnRpId, bool) {
	if o == nil || IsNil(o.RpId) {
		return nil, false
	}
	return o.RpId, true
}

// HasRpId returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasRpId() bool {
	if o != nil && !IsNil(o.RpId) {
		return true
	}

	return false
}

// SetRpId gets a reference to the given WebAuthnRpId and assigns it to the RpId field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetRpId(v WebAuthnRpId) {
	o.RpId = &v
}

// GetEnableAutofillUI returns the EnableAutofillUI field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetEnableAutofillUI() bool {
	if o == nil || IsNil(o.EnableAutofillUI) {
		var ret bool
		return ret
	}
	return *o.EnableAutofillUI
}

// GetEnableAutofillUIOk returns a tuple with the EnableAutofillUI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetEnableAutofillUIOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutofillUI) {
		return nil, false
	}
	return o.EnableAutofillUI, true
}

// HasEnableAutofillUI returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasEnableAutofillUI() bool {
	if o != nil && !IsNil(o.EnableAutofillUI) {
		return true
	}

	return false
}

// SetEnableAutofillUI gets a reference to the given bool and assigns it to the EnableAutofillUI field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetEnableAutofillUI(v bool) {
	o.EnableAutofillUI = &v
}

// GetResidentKeyRequirement returns the ResidentKeyRequirement field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetResidentKeyRequirement() string {
	if o == nil || IsNil(o.ResidentKeyRequirement) {
		var ret string
		return ret
	}
	return *o.ResidentKeyRequirement
}

// GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetResidentKeyRequirementOk() (*string, bool) {
	if o == nil || IsNil(o.ResidentKeyRequirement) {
		return nil, false
	}
	return o.ResidentKeyRequirement, true
}

// HasResidentKeyRequirement returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasResidentKeyRequirement() bool {
	if o != nil && !IsNil(o.ResidentKeyRequirement) {
		return true
	}

	return false
}

// SetResidentKeyRequirement gets a reference to the given string and assigns it to the ResidentKeyRequirement field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetResidentKeyRequirement(v string) {
	o.ResidentKeyRequirement = &v
}

// GetShowSignInWithAPasskeyButton returns the ShowSignInWithAPasskeyButton field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetShowSignInWithAPasskeyButton() bool {
	if o == nil || IsNil(o.ShowSignInWithAPasskeyButton) {
		var ret bool
		return ret
	}
	return *o.ShowSignInWithAPasskeyButton
}

// GetShowSignInWithAPasskeyButtonOk returns a tuple with the ShowSignInWithAPasskeyButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetShowSignInWithAPasskeyButtonOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowSignInWithAPasskeyButton) {
		return nil, false
	}
	return o.ShowSignInWithAPasskeyButton, true
}

// HasShowSignInWithAPasskeyButton returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasShowSignInWithAPasskeyButton() bool {
	if o != nil && !IsNil(o.ShowSignInWithAPasskeyButton) {
		return true
	}

	return false
}

// SetShowSignInWithAPasskeyButton gets a reference to the given bool and assigns it to the ShowSignInWithAPasskeyButton field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetShowSignInWithAPasskeyButton(v bool) {
	o.ShowSignInWithAPasskeyButton = &v
}

// GetCertBasedAttestationValidation returns the CertBasedAttestationValidation field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetCertBasedAttestationValidation() bool {
	if o == nil || IsNil(o.CertBasedAttestationValidation) {
		var ret bool
		return ret
	}
	return *o.CertBasedAttestationValidation
}

// GetCertBasedAttestationValidationOk returns a tuple with the CertBasedAttestationValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetCertBasedAttestationValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.CertBasedAttestationValidation) {
		return nil, false
	}
	return o.CertBasedAttestationValidation, true
}

// HasCertBasedAttestationValidation returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasCertBasedAttestationValidation() bool {
	if o != nil && !IsNil(o.CertBasedAttestationValidation) {
		return true
	}

	return false
}

// SetCertBasedAttestationValidation gets a reference to the given bool and assigns it to the CertBasedAttestationValidation field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetCertBasedAttestationValidation(v bool) {
	o.CertBasedAttestationValidation = &v
}

// GetHardwareProtected returns the HardwareProtected field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetHardwareProtected() bool {
	if o == nil || IsNil(o.HardwareProtected) {
		var ret bool
		return ret
	}
	return *o.HardwareProtected
}

// GetHardwareProtectedOk returns a tuple with the HardwareProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetHardwareProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.HardwareProtected) {
		return nil, false
	}
	return o.HardwareProtected, true
}

// HasHardwareProtected returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasHardwareProtected() bool {
	if o != nil && !IsNil(o.HardwareProtected) {
		return true
	}

	return false
}

// SetHardwareProtected gets a reference to the given bool and assigns it to the HardwareProtected field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetHardwareProtected(v bool) {
	o.HardwareProtected = &v
}

// GetFipsCompliant returns the FipsCompliant field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetFipsCompliant() bool {
	if o == nil || IsNil(o.FipsCompliant) {
		var ret bool
		return ret
	}
	return *o.FipsCompliant
}

// GetFipsCompliantOk returns a tuple with the FipsCompliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetFipsCompliantOk() (*bool, bool) {
	if o == nil || IsNil(o.FipsCompliant) {
		return nil, false
	}
	return o.FipsCompliant, true
}

// HasFipsCompliant returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasFipsCompliant() bool {
	if o != nil && !IsNil(o.FipsCompliant) {
		return true
	}

	return false
}

// SetFipsCompliant gets a reference to the given bool and assigns it to the FipsCompliant field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetFipsCompliant(v bool) {
	o.FipsCompliant = &v
}

// GetAllowSyncablePasskeys returns the AllowSyncablePasskeys field value if set, zero value otherwise.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAllowSyncablePasskeys() bool {
	if o == nil || IsNil(o.AllowSyncablePasskeys) {
		var ret bool
		return ret
	}
	return *o.AllowSyncablePasskeys
}

// GetAllowSyncablePasskeysOk returns a tuple with the AllowSyncablePasskeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAllowSyncablePasskeysOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowSyncablePasskeys) {
		return nil, false
	}
	return o.AllowSyncablePasskeys, true
}

// HasAllowSyncablePasskeys returns a boolean if a field has been set.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAllowSyncablePasskeys() bool {
	if o != nil && !IsNil(o.AllowSyncablePasskeys) {
		return true
	}

	return false
}

// SetAllowSyncablePasskeys gets a reference to the given bool and assigns it to the AllowSyncablePasskeys field.
func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAllowSyncablePasskeys(v bool) {
	o.AllowSyncablePasskeys = &v
}

func (o AuthenticatorMethodWebAuthnAllOfSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorMethodWebAuthnAllOfSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AaguidGroups) {
		toSerialize["aaguidGroups"] = o.AaguidGroups
	}
	if !IsNil(o.UserVerification) {
		toSerialize["userVerification"] = o.UserVerification
	}
	if !IsNil(o.UserVerificationForVerify) {
		toSerialize["userVerificationForVerify"] = o.UserVerificationForVerify
	}
	if !IsNil(o.Attachment) {
		toSerialize["attachment"] = o.Attachment
	}
	if !IsNil(o.RpId) {
		toSerialize["rpId"] = o.RpId
	}
	if !IsNil(o.EnableAutofillUI) {
		toSerialize["enableAutofillUI"] = o.EnableAutofillUI
	}
	if !IsNil(o.ResidentKeyRequirement) {
		toSerialize["residentKeyRequirement"] = o.ResidentKeyRequirement
	}
	if !IsNil(o.ShowSignInWithAPasskeyButton) {
		toSerialize["showSignInWithAPasskeyButton"] = o.ShowSignInWithAPasskeyButton
	}
	if !IsNil(o.CertBasedAttestationValidation) {
		toSerialize["certBasedAttestationValidation"] = o.CertBasedAttestationValidation
	}
	if !IsNil(o.HardwareProtected) {
		toSerialize["hardwareProtected"] = o.HardwareProtected
	}
	if !IsNil(o.FipsCompliant) {
		toSerialize["fipsCompliant"] = o.FipsCompliant
	}
	if !IsNil(o.AllowSyncablePasskeys) {
		toSerialize["allowSyncablePasskeys"] = o.AllowSyncablePasskeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorMethodWebAuthnAllOfSettings) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorMethodWebAuthnAllOfSettings := _AuthenticatorMethodWebAuthnAllOfSettings{}

	err = json.Unmarshal(data, &varAuthenticatorMethodWebAuthnAllOfSettings)

	if err != nil {
		return err
	}

	*o = AuthenticatorMethodWebAuthnAllOfSettings(varAuthenticatorMethodWebAuthnAllOfSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aaguidGroups")
		delete(additionalProperties, "userVerification")
		delete(additionalProperties, "userVerificationForVerify")
		delete(additionalProperties, "attachment")
		delete(additionalProperties, "rpId")
		delete(additionalProperties, "enableAutofillUI")
		delete(additionalProperties, "residentKeyRequirement")
		delete(additionalProperties, "showSignInWithAPasskeyButton")
		delete(additionalProperties, "certBasedAttestationValidation")
		delete(additionalProperties, "hardwareProtected")
		delete(additionalProperties, "fipsCompliant")
		delete(additionalProperties, "allowSyncablePasskeys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorMethodWebAuthnAllOfSettings struct {
	value *AuthenticatorMethodWebAuthnAllOfSettings
	isSet bool
}

func (v NullableAuthenticatorMethodWebAuthnAllOfSettings) Get() *AuthenticatorMethodWebAuthnAllOfSettings {
	return v.value
}

func (v *NullableAuthenticatorMethodWebAuthnAllOfSettings) Set(val *AuthenticatorMethodWebAuthnAllOfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorMethodWebAuthnAllOfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorMethodWebAuthnAllOfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorMethodWebAuthnAllOfSettings(val *AuthenticatorMethodWebAuthnAllOfSettings) *NullableAuthenticatorMethodWebAuthnAllOfSettings {
	return &NullableAuthenticatorMethodWebAuthnAllOfSettings{value: val, isSet: true}
}

func (v NullableAuthenticatorMethodWebAuthnAllOfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorMethodWebAuthnAllOfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
