# AuthenticatorMethodWebAuthnAllOfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AaguidGroups** | Pointer to [**[]AAGUIDGroupObject**](AAGUIDGroupObject.md) | The FIDO2 Authenticator Attestation Global Unique Identifiers (AAGUID) groups available to the WebAuthn authenticator | [optional] 
**UserVerification** | Pointer to **string** | User verification settings for enrollment.  This setting controls the user verification requirement during the enrollment of a new credential. It determines whether the authenticator requires verification when a user is registering their device or credential. | [optional] 
**UserVerificationForVerify** | Pointer to **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;User verification settings for verification. This setting controls the user verification requirement during authentication (verification). It determines whether the authenticator requires user verification when a user signs in with an already-registered credential.  For verification, the value defaults to &#x60;PREFERRED&#x60;, unless the enrollment setting is &#x60;REQUIRED&#x60;. If the enrollment setting is &#x60;REQUIRED&#x60; for the authenticator, then the verification setting is also implicitly &#x60;REQUIRED&#x60;.  &gt; **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id&#x3D;ext_Manage_Early_Access_features). | [optional] 
**Attachment** | Pointer to **string** | Method attachment | [optional] 
**RpId** | Pointer to [**WebAuthnRpId**](WebAuthnRpId.md) |  | [optional] 
**EnableAutofillUI** | Pointer to **bool** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Enables the passkeys autofill UI to display available WebAuthn discoverable credentials (\&quot;resident key\&quot;) from the Sign-In Widget username field | [optional] [default to false]
**ResidentKeyRequirement** | Pointer to **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Resident key requirement setting. Okta recommends using only &#x60;REQUIRED&#x60; or &#x60;DISCOURAGED&#x60; to make the requirement preference explicit. Using &#x60;PREFERRED&#x60; can sometimes lead to unpredictable behavior depending on the client platform and authenticator capabilities.  &gt; **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id&#x3D;ext_Manage_Early_Access_features). | [optional] 
**ShowSignInWithAPasskeyButton** | Pointer to **bool** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Indicates if the **Sign in with a Passkey** button on the Sign-In Widget is shown.   &gt; **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id&#x3D;ext_Manage_Early_Access_features). | [optional] [default to false]
**CertBasedAttestationValidation** | Pointer to **bool** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Indicates whether certificate-based attestation validation is enabled. When enabled, the authenticator&#39;s attestation certificate is validated against known root certificates (custom AAGUIDs with associated certificates or the [FIDO Metadata Service](https://fidoalliance.org/metadata/)) to ensure its validity.  &gt; **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id&#x3D;ext_Manage_Early_Access_features). | [optional] [default to false]
**HardwareProtected** | Pointer to **bool** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Indicates whether the authenticator is required to store the private key on a hardware component  &gt; **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id&#x3D;ext_Manage_Early_Access_features). | [optional] [default to false]
**FipsCompliant** | Pointer to **bool** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Indicates whether the authenticator is required to be [Federal Information Processing Standards (FIPS)](https://csrc.nist.gov/glossary/term/federal_information_processing_standard) compliant  &gt; **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id&#x3D;ext_Manage_Early_Access_features). | [optional] [default to false]
**AllowSyncablePasskeys** | Pointer to **bool** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Indicates whether syncable passkeys are allowed. When enabled, users can register passkeys that are synchronized across their devices by using platform-specific mechanisms (such as iCloud Keychain for Apple devices or Google Password Manager for Android devices).  &gt; **Note:** This setting is only available when you have enabled the **Passkeys Rebrand** feature. See [Enable self-service features](https://help.okta.com/okta_help.htm?id&#x3D;ext_Manage_Early_Access_features). | [optional] [default to true]

## Methods

### NewAuthenticatorMethodWebAuthnAllOfSettings

`func NewAuthenticatorMethodWebAuthnAllOfSettings() *AuthenticatorMethodWebAuthnAllOfSettings`

NewAuthenticatorMethodWebAuthnAllOfSettings instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults

`func NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults() *AuthenticatorMethodWebAuthnAllOfSettings`

NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAaguidGroups

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAaguidGroups() []AAGUIDGroupObject`

GetAaguidGroups returns the AaguidGroups field if non-nil, zero value otherwise.

### GetAaguidGroupsOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAaguidGroupsOk() (*[]AAGUIDGroupObject, bool)`

GetAaguidGroupsOk returns a tuple with the AaguidGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaguidGroups

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAaguidGroups(v []AAGUIDGroupObject)`

SetAaguidGroups sets AaguidGroups field to given value.

### HasAaguidGroups

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAaguidGroups() bool`

HasAaguidGroups returns a boolean if a field has been set.

### GetUserVerification

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerification() string`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerificationOk() (*string, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetUserVerification(v string)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetUserVerificationForVerify

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerificationForVerify() string`

GetUserVerificationForVerify returns the UserVerificationForVerify field if non-nil, zero value otherwise.

### GetUserVerificationForVerifyOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerificationForVerifyOk() (*string, bool)`

GetUserVerificationForVerifyOk returns a tuple with the UserVerificationForVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerificationForVerify

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetUserVerificationForVerify(v string)`

SetUserVerificationForVerify sets UserVerificationForVerify field to given value.

### HasUserVerificationForVerify

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasUserVerificationForVerify() bool`

HasUserVerificationForVerify returns a boolean if a field has been set.

### GetAttachment

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAttachment() string`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAttachmentOk() (*string, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAttachment(v string)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.

### GetRpId

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetRpId() WebAuthnRpId`

GetRpId returns the RpId field if non-nil, zero value otherwise.

### GetRpIdOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetRpIdOk() (*WebAuthnRpId, bool)`

GetRpIdOk returns a tuple with the RpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpId

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetRpId(v WebAuthnRpId)`

SetRpId sets RpId field to given value.

### HasRpId

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasRpId() bool`

HasRpId returns a boolean if a field has been set.

### GetEnableAutofillUI

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetEnableAutofillUI() bool`

GetEnableAutofillUI returns the EnableAutofillUI field if non-nil, zero value otherwise.

### GetEnableAutofillUIOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetEnableAutofillUIOk() (*bool, bool)`

GetEnableAutofillUIOk returns a tuple with the EnableAutofillUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutofillUI

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetEnableAutofillUI(v bool)`

SetEnableAutofillUI sets EnableAutofillUI field to given value.

### HasEnableAutofillUI

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasEnableAutofillUI() bool`

HasEnableAutofillUI returns a boolean if a field has been set.

### GetResidentKeyRequirement

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetResidentKeyRequirement() string`

GetResidentKeyRequirement returns the ResidentKeyRequirement field if non-nil, zero value otherwise.

### GetResidentKeyRequirementOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetResidentKeyRequirementOk() (*string, bool)`

GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentKeyRequirement

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetResidentKeyRequirement(v string)`

SetResidentKeyRequirement sets ResidentKeyRequirement field to given value.

### HasResidentKeyRequirement

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasResidentKeyRequirement() bool`

HasResidentKeyRequirement returns a boolean if a field has been set.

### GetShowSignInWithAPasskeyButton

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetShowSignInWithAPasskeyButton() bool`

GetShowSignInWithAPasskeyButton returns the ShowSignInWithAPasskeyButton field if non-nil, zero value otherwise.

### GetShowSignInWithAPasskeyButtonOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetShowSignInWithAPasskeyButtonOk() (*bool, bool)`

GetShowSignInWithAPasskeyButtonOk returns a tuple with the ShowSignInWithAPasskeyButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSignInWithAPasskeyButton

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetShowSignInWithAPasskeyButton(v bool)`

SetShowSignInWithAPasskeyButton sets ShowSignInWithAPasskeyButton field to given value.

### HasShowSignInWithAPasskeyButton

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasShowSignInWithAPasskeyButton() bool`

HasShowSignInWithAPasskeyButton returns a boolean if a field has been set.

### GetCertBasedAttestationValidation

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetCertBasedAttestationValidation() bool`

GetCertBasedAttestationValidation returns the CertBasedAttestationValidation field if non-nil, zero value otherwise.

### GetCertBasedAttestationValidationOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetCertBasedAttestationValidationOk() (*bool, bool)`

GetCertBasedAttestationValidationOk returns a tuple with the CertBasedAttestationValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertBasedAttestationValidation

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetCertBasedAttestationValidation(v bool)`

SetCertBasedAttestationValidation sets CertBasedAttestationValidation field to given value.

### HasCertBasedAttestationValidation

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasCertBasedAttestationValidation() bool`

HasCertBasedAttestationValidation returns a boolean if a field has been set.

### GetHardwareProtected

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetHardwareProtected() bool`

GetHardwareProtected returns the HardwareProtected field if non-nil, zero value otherwise.

### GetHardwareProtectedOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetHardwareProtectedOk() (*bool, bool)`

GetHardwareProtectedOk returns a tuple with the HardwareProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareProtected

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetHardwareProtected(v bool)`

SetHardwareProtected sets HardwareProtected field to given value.

### HasHardwareProtected

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasHardwareProtected() bool`

HasHardwareProtected returns a boolean if a field has been set.

### GetFipsCompliant

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetFipsCompliant() bool`

GetFipsCompliant returns the FipsCompliant field if non-nil, zero value otherwise.

### GetFipsCompliantOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetFipsCompliantOk() (*bool, bool)`

GetFipsCompliantOk returns a tuple with the FipsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsCompliant

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetFipsCompliant(v bool)`

SetFipsCompliant sets FipsCompliant field to given value.

### HasFipsCompliant

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasFipsCompliant() bool`

HasFipsCompliant returns a boolean if a field has been set.

### GetAllowSyncablePasskeys

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAllowSyncablePasskeys() bool`

GetAllowSyncablePasskeys returns the AllowSyncablePasskeys field if non-nil, zero value otherwise.

### GetAllowSyncablePasskeysOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAllowSyncablePasskeysOk() (*bool, bool)`

GetAllowSyncablePasskeysOk returns a tuple with the AllowSyncablePasskeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSyncablePasskeys

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAllowSyncablePasskeys(v bool)`

SetAllowSyncablePasskeys sets AllowSyncablePasskeys field to given value.

### HasAllowSyncablePasskeys

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAllowSyncablePasskeys() bool`

HasAllowSyncablePasskeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


