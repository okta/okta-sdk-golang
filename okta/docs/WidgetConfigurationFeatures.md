# WidgetConfigurationFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Router** | Pointer to **bool** |  | [optional] 
**SecurityImage** | Pointer to **bool** |  | [optional] 
**RememberMe** | Pointer to **bool** |  | [optional] 
**AutoPush** | Pointer to **bool** |  | [optional] 
**SmsRecovery** | Pointer to **bool** |  | [optional] 
**CallRecovery** | Pointer to **bool** |  | [optional] 
**EmailRecovery** | Pointer to **bool** |  | [optional] 
**Webauthn** | Pointer to **bool** |  | [optional] 
**SelfServiceUnlock** | Pointer to **bool** |  | [optional] 
**MultiOptionalFactorEnroll** | Pointer to **bool** |  | [optional] 
**DeviceFingerprinting** | Pointer to **bool** |  | [optional] 
**HideSignOutLinkInMFA** | Pointer to **bool** |  | [optional] 
**HideBackToSignInForReset** | Pointer to **bool** |  | [optional] 
**CustomExpiredPassword** | Pointer to **bool** |  | [optional] 
**Registration** | Pointer to **bool** |  | [optional] 
**IdpDiscovery** | Pointer to **bool** |  | [optional] 
**PasswordlessAuth** | Pointer to **bool** |  | [optional] 
**ShowPasswordToggleOnSignInPage** | Pointer to **bool** |  | [optional] 
**TrackTypingPattern** | Pointer to **bool** |  | [optional] 
**RedirectByFormSubmit** | Pointer to **bool** |  | [optional] 
**SameDeviceOVEnrollmentEnabled** | Pointer to **bool** |  | [optional] 
**UseDeviceFingerprintForSecurityImage** | Pointer to **bool** |  | [optional] 
**ShowPasswordRequirementsAsHtmlList** | Pointer to **bool** |  | [optional] 
**MfaOnlyFlow** | Pointer to **bool** |  | [optional] 

## Methods

### NewWidgetConfigurationFeatures

`func NewWidgetConfigurationFeatures() *WidgetConfigurationFeatures`

NewWidgetConfigurationFeatures instantiates a new WidgetConfigurationFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetConfigurationFeaturesWithDefaults

`func NewWidgetConfigurationFeaturesWithDefaults() *WidgetConfigurationFeatures`

NewWidgetConfigurationFeaturesWithDefaults instantiates a new WidgetConfigurationFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouter

`func (o *WidgetConfigurationFeatures) GetRouter() bool`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *WidgetConfigurationFeatures) GetRouterOk() (*bool, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *WidgetConfigurationFeatures) SetRouter(v bool)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *WidgetConfigurationFeatures) HasRouter() bool`

HasRouter returns a boolean if a field has been set.

### GetSecurityImage

`func (o *WidgetConfigurationFeatures) GetSecurityImage() bool`

GetSecurityImage returns the SecurityImage field if non-nil, zero value otherwise.

### GetSecurityImageOk

`func (o *WidgetConfigurationFeatures) GetSecurityImageOk() (*bool, bool)`

GetSecurityImageOk returns a tuple with the SecurityImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityImage

`func (o *WidgetConfigurationFeatures) SetSecurityImage(v bool)`

SetSecurityImage sets SecurityImage field to given value.

### HasSecurityImage

`func (o *WidgetConfigurationFeatures) HasSecurityImage() bool`

HasSecurityImage returns a boolean if a field has been set.

### GetRememberMe

`func (o *WidgetConfigurationFeatures) GetRememberMe() bool`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *WidgetConfigurationFeatures) GetRememberMeOk() (*bool, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *WidgetConfigurationFeatures) SetRememberMe(v bool)`

SetRememberMe sets RememberMe field to given value.

### HasRememberMe

`func (o *WidgetConfigurationFeatures) HasRememberMe() bool`

HasRememberMe returns a boolean if a field has been set.

### GetAutoPush

`func (o *WidgetConfigurationFeatures) GetAutoPush() bool`

GetAutoPush returns the AutoPush field if non-nil, zero value otherwise.

### GetAutoPushOk

`func (o *WidgetConfigurationFeatures) GetAutoPushOk() (*bool, bool)`

GetAutoPushOk returns a tuple with the AutoPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPush

`func (o *WidgetConfigurationFeatures) SetAutoPush(v bool)`

SetAutoPush sets AutoPush field to given value.

### HasAutoPush

`func (o *WidgetConfigurationFeatures) HasAutoPush() bool`

HasAutoPush returns a boolean if a field has been set.

### GetSmsRecovery

`func (o *WidgetConfigurationFeatures) GetSmsRecovery() bool`

GetSmsRecovery returns the SmsRecovery field if non-nil, zero value otherwise.

### GetSmsRecoveryOk

`func (o *WidgetConfigurationFeatures) GetSmsRecoveryOk() (*bool, bool)`

GetSmsRecoveryOk returns a tuple with the SmsRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsRecovery

`func (o *WidgetConfigurationFeatures) SetSmsRecovery(v bool)`

SetSmsRecovery sets SmsRecovery field to given value.

### HasSmsRecovery

`func (o *WidgetConfigurationFeatures) HasSmsRecovery() bool`

HasSmsRecovery returns a boolean if a field has been set.

### GetCallRecovery

`func (o *WidgetConfigurationFeatures) GetCallRecovery() bool`

GetCallRecovery returns the CallRecovery field if non-nil, zero value otherwise.

### GetCallRecoveryOk

`func (o *WidgetConfigurationFeatures) GetCallRecoveryOk() (*bool, bool)`

GetCallRecoveryOk returns a tuple with the CallRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecovery

`func (o *WidgetConfigurationFeatures) SetCallRecovery(v bool)`

SetCallRecovery sets CallRecovery field to given value.

### HasCallRecovery

`func (o *WidgetConfigurationFeatures) HasCallRecovery() bool`

HasCallRecovery returns a boolean if a field has been set.

### GetEmailRecovery

`func (o *WidgetConfigurationFeatures) GetEmailRecovery() bool`

GetEmailRecovery returns the EmailRecovery field if non-nil, zero value otherwise.

### GetEmailRecoveryOk

`func (o *WidgetConfigurationFeatures) GetEmailRecoveryOk() (*bool, bool)`

GetEmailRecoveryOk returns a tuple with the EmailRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecovery

`func (o *WidgetConfigurationFeatures) SetEmailRecovery(v bool)`

SetEmailRecovery sets EmailRecovery field to given value.

### HasEmailRecovery

`func (o *WidgetConfigurationFeatures) HasEmailRecovery() bool`

HasEmailRecovery returns a boolean if a field has been set.

### GetWebauthn

`func (o *WidgetConfigurationFeatures) GetWebauthn() bool`

GetWebauthn returns the Webauthn field if non-nil, zero value otherwise.

### GetWebauthnOk

`func (o *WidgetConfigurationFeatures) GetWebauthnOk() (*bool, bool)`

GetWebauthnOk returns a tuple with the Webauthn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebauthn

`func (o *WidgetConfigurationFeatures) SetWebauthn(v bool)`

SetWebauthn sets Webauthn field to given value.

### HasWebauthn

`func (o *WidgetConfigurationFeatures) HasWebauthn() bool`

HasWebauthn returns a boolean if a field has been set.

### GetSelfServiceUnlock

`func (o *WidgetConfigurationFeatures) GetSelfServiceUnlock() bool`

GetSelfServiceUnlock returns the SelfServiceUnlock field if non-nil, zero value otherwise.

### GetSelfServiceUnlockOk

`func (o *WidgetConfigurationFeatures) GetSelfServiceUnlockOk() (*bool, bool)`

GetSelfServiceUnlockOk returns a tuple with the SelfServiceUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceUnlock

`func (o *WidgetConfigurationFeatures) SetSelfServiceUnlock(v bool)`

SetSelfServiceUnlock sets SelfServiceUnlock field to given value.

### HasSelfServiceUnlock

`func (o *WidgetConfigurationFeatures) HasSelfServiceUnlock() bool`

HasSelfServiceUnlock returns a boolean if a field has been set.

### GetMultiOptionalFactorEnroll

`func (o *WidgetConfigurationFeatures) GetMultiOptionalFactorEnroll() bool`

GetMultiOptionalFactorEnroll returns the MultiOptionalFactorEnroll field if non-nil, zero value otherwise.

### GetMultiOptionalFactorEnrollOk

`func (o *WidgetConfigurationFeatures) GetMultiOptionalFactorEnrollOk() (*bool, bool)`

GetMultiOptionalFactorEnrollOk returns a tuple with the MultiOptionalFactorEnroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiOptionalFactorEnroll

`func (o *WidgetConfigurationFeatures) SetMultiOptionalFactorEnroll(v bool)`

SetMultiOptionalFactorEnroll sets MultiOptionalFactorEnroll field to given value.

### HasMultiOptionalFactorEnroll

`func (o *WidgetConfigurationFeatures) HasMultiOptionalFactorEnroll() bool`

HasMultiOptionalFactorEnroll returns a boolean if a field has been set.

### GetDeviceFingerprinting

`func (o *WidgetConfigurationFeatures) GetDeviceFingerprinting() bool`

GetDeviceFingerprinting returns the DeviceFingerprinting field if non-nil, zero value otherwise.

### GetDeviceFingerprintingOk

`func (o *WidgetConfigurationFeatures) GetDeviceFingerprintingOk() (*bool, bool)`

GetDeviceFingerprintingOk returns a tuple with the DeviceFingerprinting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceFingerprinting

`func (o *WidgetConfigurationFeatures) SetDeviceFingerprinting(v bool)`

SetDeviceFingerprinting sets DeviceFingerprinting field to given value.

### HasDeviceFingerprinting

`func (o *WidgetConfigurationFeatures) HasDeviceFingerprinting() bool`

HasDeviceFingerprinting returns a boolean if a field has been set.

### GetHideSignOutLinkInMFA

`func (o *WidgetConfigurationFeatures) GetHideSignOutLinkInMFA() bool`

GetHideSignOutLinkInMFA returns the HideSignOutLinkInMFA field if non-nil, zero value otherwise.

### GetHideSignOutLinkInMFAOk

`func (o *WidgetConfigurationFeatures) GetHideSignOutLinkInMFAOk() (*bool, bool)`

GetHideSignOutLinkInMFAOk returns a tuple with the HideSignOutLinkInMFA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideSignOutLinkInMFA

`func (o *WidgetConfigurationFeatures) SetHideSignOutLinkInMFA(v bool)`

SetHideSignOutLinkInMFA sets HideSignOutLinkInMFA field to given value.

### HasHideSignOutLinkInMFA

`func (o *WidgetConfigurationFeatures) HasHideSignOutLinkInMFA() bool`

HasHideSignOutLinkInMFA returns a boolean if a field has been set.

### GetHideBackToSignInForReset

`func (o *WidgetConfigurationFeatures) GetHideBackToSignInForReset() bool`

GetHideBackToSignInForReset returns the HideBackToSignInForReset field if non-nil, zero value otherwise.

### GetHideBackToSignInForResetOk

`func (o *WidgetConfigurationFeatures) GetHideBackToSignInForResetOk() (*bool, bool)`

GetHideBackToSignInForResetOk returns a tuple with the HideBackToSignInForReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideBackToSignInForReset

`func (o *WidgetConfigurationFeatures) SetHideBackToSignInForReset(v bool)`

SetHideBackToSignInForReset sets HideBackToSignInForReset field to given value.

### HasHideBackToSignInForReset

`func (o *WidgetConfigurationFeatures) HasHideBackToSignInForReset() bool`

HasHideBackToSignInForReset returns a boolean if a field has been set.

### GetCustomExpiredPassword

`func (o *WidgetConfigurationFeatures) GetCustomExpiredPassword() bool`

GetCustomExpiredPassword returns the CustomExpiredPassword field if non-nil, zero value otherwise.

### GetCustomExpiredPasswordOk

`func (o *WidgetConfigurationFeatures) GetCustomExpiredPasswordOk() (*bool, bool)`

GetCustomExpiredPasswordOk returns a tuple with the CustomExpiredPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomExpiredPassword

`func (o *WidgetConfigurationFeatures) SetCustomExpiredPassword(v bool)`

SetCustomExpiredPassword sets CustomExpiredPassword field to given value.

### HasCustomExpiredPassword

`func (o *WidgetConfigurationFeatures) HasCustomExpiredPassword() bool`

HasCustomExpiredPassword returns a boolean if a field has been set.

### GetRegistration

`func (o *WidgetConfigurationFeatures) GetRegistration() bool`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *WidgetConfigurationFeatures) GetRegistrationOk() (*bool, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *WidgetConfigurationFeatures) SetRegistration(v bool)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *WidgetConfigurationFeatures) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetIdpDiscovery

`func (o *WidgetConfigurationFeatures) GetIdpDiscovery() bool`

GetIdpDiscovery returns the IdpDiscovery field if non-nil, zero value otherwise.

### GetIdpDiscoveryOk

`func (o *WidgetConfigurationFeatures) GetIdpDiscoveryOk() (*bool, bool)`

GetIdpDiscoveryOk returns a tuple with the IdpDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpDiscovery

`func (o *WidgetConfigurationFeatures) SetIdpDiscovery(v bool)`

SetIdpDiscovery sets IdpDiscovery field to given value.

### HasIdpDiscovery

`func (o *WidgetConfigurationFeatures) HasIdpDiscovery() bool`

HasIdpDiscovery returns a boolean if a field has been set.

### GetPasswordlessAuth

`func (o *WidgetConfigurationFeatures) GetPasswordlessAuth() bool`

GetPasswordlessAuth returns the PasswordlessAuth field if non-nil, zero value otherwise.

### GetPasswordlessAuthOk

`func (o *WidgetConfigurationFeatures) GetPasswordlessAuthOk() (*bool, bool)`

GetPasswordlessAuthOk returns a tuple with the PasswordlessAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordlessAuth

`func (o *WidgetConfigurationFeatures) SetPasswordlessAuth(v bool)`

SetPasswordlessAuth sets PasswordlessAuth field to given value.

### HasPasswordlessAuth

`func (o *WidgetConfigurationFeatures) HasPasswordlessAuth() bool`

HasPasswordlessAuth returns a boolean if a field has been set.

### GetShowPasswordToggleOnSignInPage

`func (o *WidgetConfigurationFeatures) GetShowPasswordToggleOnSignInPage() bool`

GetShowPasswordToggleOnSignInPage returns the ShowPasswordToggleOnSignInPage field if non-nil, zero value otherwise.

### GetShowPasswordToggleOnSignInPageOk

`func (o *WidgetConfigurationFeatures) GetShowPasswordToggleOnSignInPageOk() (*bool, bool)`

GetShowPasswordToggleOnSignInPageOk returns a tuple with the ShowPasswordToggleOnSignInPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPasswordToggleOnSignInPage

`func (o *WidgetConfigurationFeatures) SetShowPasswordToggleOnSignInPage(v bool)`

SetShowPasswordToggleOnSignInPage sets ShowPasswordToggleOnSignInPage field to given value.

### HasShowPasswordToggleOnSignInPage

`func (o *WidgetConfigurationFeatures) HasShowPasswordToggleOnSignInPage() bool`

HasShowPasswordToggleOnSignInPage returns a boolean if a field has been set.

### GetTrackTypingPattern

`func (o *WidgetConfigurationFeatures) GetTrackTypingPattern() bool`

GetTrackTypingPattern returns the TrackTypingPattern field if non-nil, zero value otherwise.

### GetTrackTypingPatternOk

`func (o *WidgetConfigurationFeatures) GetTrackTypingPatternOk() (*bool, bool)`

GetTrackTypingPatternOk returns a tuple with the TrackTypingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackTypingPattern

`func (o *WidgetConfigurationFeatures) SetTrackTypingPattern(v bool)`

SetTrackTypingPattern sets TrackTypingPattern field to given value.

### HasTrackTypingPattern

`func (o *WidgetConfigurationFeatures) HasTrackTypingPattern() bool`

HasTrackTypingPattern returns a boolean if a field has been set.

### GetRedirectByFormSubmit

`func (o *WidgetConfigurationFeatures) GetRedirectByFormSubmit() bool`

GetRedirectByFormSubmit returns the RedirectByFormSubmit field if non-nil, zero value otherwise.

### GetRedirectByFormSubmitOk

`func (o *WidgetConfigurationFeatures) GetRedirectByFormSubmitOk() (*bool, bool)`

GetRedirectByFormSubmitOk returns a tuple with the RedirectByFormSubmit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectByFormSubmit

`func (o *WidgetConfigurationFeatures) SetRedirectByFormSubmit(v bool)`

SetRedirectByFormSubmit sets RedirectByFormSubmit field to given value.

### HasRedirectByFormSubmit

`func (o *WidgetConfigurationFeatures) HasRedirectByFormSubmit() bool`

HasRedirectByFormSubmit returns a boolean if a field has been set.

### GetSameDeviceOVEnrollmentEnabled

`func (o *WidgetConfigurationFeatures) GetSameDeviceOVEnrollmentEnabled() bool`

GetSameDeviceOVEnrollmentEnabled returns the SameDeviceOVEnrollmentEnabled field if non-nil, zero value otherwise.

### GetSameDeviceOVEnrollmentEnabledOk

`func (o *WidgetConfigurationFeatures) GetSameDeviceOVEnrollmentEnabledOk() (*bool, bool)`

GetSameDeviceOVEnrollmentEnabledOk returns a tuple with the SameDeviceOVEnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameDeviceOVEnrollmentEnabled

`func (o *WidgetConfigurationFeatures) SetSameDeviceOVEnrollmentEnabled(v bool)`

SetSameDeviceOVEnrollmentEnabled sets SameDeviceOVEnrollmentEnabled field to given value.

### HasSameDeviceOVEnrollmentEnabled

`func (o *WidgetConfigurationFeatures) HasSameDeviceOVEnrollmentEnabled() bool`

HasSameDeviceOVEnrollmentEnabled returns a boolean if a field has been set.

### GetUseDeviceFingerprintForSecurityImage

`func (o *WidgetConfigurationFeatures) GetUseDeviceFingerprintForSecurityImage() bool`

GetUseDeviceFingerprintForSecurityImage returns the UseDeviceFingerprintForSecurityImage field if non-nil, zero value otherwise.

### GetUseDeviceFingerprintForSecurityImageOk

`func (o *WidgetConfigurationFeatures) GetUseDeviceFingerprintForSecurityImageOk() (*bool, bool)`

GetUseDeviceFingerprintForSecurityImageOk returns a tuple with the UseDeviceFingerprintForSecurityImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDeviceFingerprintForSecurityImage

`func (o *WidgetConfigurationFeatures) SetUseDeviceFingerprintForSecurityImage(v bool)`

SetUseDeviceFingerprintForSecurityImage sets UseDeviceFingerprintForSecurityImage field to given value.

### HasUseDeviceFingerprintForSecurityImage

`func (o *WidgetConfigurationFeatures) HasUseDeviceFingerprintForSecurityImage() bool`

HasUseDeviceFingerprintForSecurityImage returns a boolean if a field has been set.

### GetShowPasswordRequirementsAsHtmlList

`func (o *WidgetConfigurationFeatures) GetShowPasswordRequirementsAsHtmlList() bool`

GetShowPasswordRequirementsAsHtmlList returns the ShowPasswordRequirementsAsHtmlList field if non-nil, zero value otherwise.

### GetShowPasswordRequirementsAsHtmlListOk

`func (o *WidgetConfigurationFeatures) GetShowPasswordRequirementsAsHtmlListOk() (*bool, bool)`

GetShowPasswordRequirementsAsHtmlListOk returns a tuple with the ShowPasswordRequirementsAsHtmlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPasswordRequirementsAsHtmlList

`func (o *WidgetConfigurationFeatures) SetShowPasswordRequirementsAsHtmlList(v bool)`

SetShowPasswordRequirementsAsHtmlList sets ShowPasswordRequirementsAsHtmlList field to given value.

### HasShowPasswordRequirementsAsHtmlList

`func (o *WidgetConfigurationFeatures) HasShowPasswordRequirementsAsHtmlList() bool`

HasShowPasswordRequirementsAsHtmlList returns a boolean if a field has been set.

### GetMfaOnlyFlow

`func (o *WidgetConfigurationFeatures) GetMfaOnlyFlow() bool`

GetMfaOnlyFlow returns the MfaOnlyFlow field if non-nil, zero value otherwise.

### GetMfaOnlyFlowOk

`func (o *WidgetConfigurationFeatures) GetMfaOnlyFlowOk() (*bool, bool)`

GetMfaOnlyFlowOk returns a tuple with the MfaOnlyFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaOnlyFlow

`func (o *WidgetConfigurationFeatures) SetMfaOnlyFlow(v bool)`

SetMfaOnlyFlow sets MfaOnlyFlow field to given value.

### HasMfaOnlyFlow

`func (o *WidgetConfigurationFeatures) HasMfaOnlyFlow() bool`

HasMfaOnlyFlow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


