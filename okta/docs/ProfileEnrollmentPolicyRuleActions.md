# ProfileEnrollmentPolicyRuleActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enroll** | Pointer to [**PolicyRuleActionsEnroll**](PolicyRuleActionsEnroll.md) |  | [optional] 
**Idp** | Pointer to [**IdpPolicyRuleAction**](IdpPolicyRuleAction.md) |  | [optional] 
**PasswordChange** | Pointer to [**PasswordPolicyRuleAction**](PasswordPolicyRuleAction.md) |  | [optional] 
**SelfServicePasswordReset** | Pointer to [**PasswordPolicyRuleAction**](PasswordPolicyRuleAction.md) |  | [optional] 
**SelfServiceUnlock** | Pointer to [**PasswordPolicyRuleAction**](PasswordPolicyRuleAction.md) |  | [optional] 
**Signon** | Pointer to [**OktaSignOnPolicyRuleSignonActions**](OktaSignOnPolicyRuleSignonActions.md) |  | [optional] 
**ProfileEnrollment** | Pointer to [**ProfileEnrollmentPolicyRuleAction**](ProfileEnrollmentPolicyRuleAction.md) |  | [optional] 

## Methods

### NewProfileEnrollmentPolicyRuleActions

`func NewProfileEnrollmentPolicyRuleActions() *ProfileEnrollmentPolicyRuleActions`

NewProfileEnrollmentPolicyRuleActions instantiates a new ProfileEnrollmentPolicyRuleActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileEnrollmentPolicyRuleActionsWithDefaults

`func NewProfileEnrollmentPolicyRuleActionsWithDefaults() *ProfileEnrollmentPolicyRuleActions`

NewProfileEnrollmentPolicyRuleActionsWithDefaults instantiates a new ProfileEnrollmentPolicyRuleActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnroll

`func (o *ProfileEnrollmentPolicyRuleActions) GetEnroll() PolicyRuleActionsEnroll`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *ProfileEnrollmentPolicyRuleActions) GetEnrollOk() (*PolicyRuleActionsEnroll, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *ProfileEnrollmentPolicyRuleActions) SetEnroll(v PolicyRuleActionsEnroll)`

SetEnroll sets Enroll field to given value.

### HasEnroll

`func (o *ProfileEnrollmentPolicyRuleActions) HasEnroll() bool`

HasEnroll returns a boolean if a field has been set.

### GetIdp

`func (o *ProfileEnrollmentPolicyRuleActions) GetIdp() IdpPolicyRuleAction`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *ProfileEnrollmentPolicyRuleActions) GetIdpOk() (*IdpPolicyRuleAction, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *ProfileEnrollmentPolicyRuleActions) SetIdp(v IdpPolicyRuleAction)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *ProfileEnrollmentPolicyRuleActions) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetPasswordChange

`func (o *ProfileEnrollmentPolicyRuleActions) GetPasswordChange() PasswordPolicyRuleAction`

GetPasswordChange returns the PasswordChange field if non-nil, zero value otherwise.

### GetPasswordChangeOk

`func (o *ProfileEnrollmentPolicyRuleActions) GetPasswordChangeOk() (*PasswordPolicyRuleAction, bool)`

GetPasswordChangeOk returns a tuple with the PasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChange

`func (o *ProfileEnrollmentPolicyRuleActions) SetPasswordChange(v PasswordPolicyRuleAction)`

SetPasswordChange sets PasswordChange field to given value.

### HasPasswordChange

`func (o *ProfileEnrollmentPolicyRuleActions) HasPasswordChange() bool`

HasPasswordChange returns a boolean if a field has been set.

### GetSelfServicePasswordReset

`func (o *ProfileEnrollmentPolicyRuleActions) GetSelfServicePasswordReset() PasswordPolicyRuleAction`

GetSelfServicePasswordReset returns the SelfServicePasswordReset field if non-nil, zero value otherwise.

### GetSelfServicePasswordResetOk

`func (o *ProfileEnrollmentPolicyRuleActions) GetSelfServicePasswordResetOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServicePasswordResetOk returns a tuple with the SelfServicePasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServicePasswordReset

`func (o *ProfileEnrollmentPolicyRuleActions) SetSelfServicePasswordReset(v PasswordPolicyRuleAction)`

SetSelfServicePasswordReset sets SelfServicePasswordReset field to given value.

### HasSelfServicePasswordReset

`func (o *ProfileEnrollmentPolicyRuleActions) HasSelfServicePasswordReset() bool`

HasSelfServicePasswordReset returns a boolean if a field has been set.

### GetSelfServiceUnlock

`func (o *ProfileEnrollmentPolicyRuleActions) GetSelfServiceUnlock() PasswordPolicyRuleAction`

GetSelfServiceUnlock returns the SelfServiceUnlock field if non-nil, zero value otherwise.

### GetSelfServiceUnlockOk

`func (o *ProfileEnrollmentPolicyRuleActions) GetSelfServiceUnlockOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServiceUnlockOk returns a tuple with the SelfServiceUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceUnlock

`func (o *ProfileEnrollmentPolicyRuleActions) SetSelfServiceUnlock(v PasswordPolicyRuleAction)`

SetSelfServiceUnlock sets SelfServiceUnlock field to given value.

### HasSelfServiceUnlock

`func (o *ProfileEnrollmentPolicyRuleActions) HasSelfServiceUnlock() bool`

HasSelfServiceUnlock returns a boolean if a field has been set.

### GetSignon

`func (o *ProfileEnrollmentPolicyRuleActions) GetSignon() OktaSignOnPolicyRuleSignonActions`

GetSignon returns the Signon field if non-nil, zero value otherwise.

### GetSignonOk

`func (o *ProfileEnrollmentPolicyRuleActions) GetSignonOk() (*OktaSignOnPolicyRuleSignonActions, bool)`

GetSignonOk returns a tuple with the Signon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignon

`func (o *ProfileEnrollmentPolicyRuleActions) SetSignon(v OktaSignOnPolicyRuleSignonActions)`

SetSignon sets Signon field to given value.

### HasSignon

`func (o *ProfileEnrollmentPolicyRuleActions) HasSignon() bool`

HasSignon returns a boolean if a field has been set.

### GetProfileEnrollment

`func (o *ProfileEnrollmentPolicyRuleActions) GetProfileEnrollment() ProfileEnrollmentPolicyRuleAction`

GetProfileEnrollment returns the ProfileEnrollment field if non-nil, zero value otherwise.

### GetProfileEnrollmentOk

`func (o *ProfileEnrollmentPolicyRuleActions) GetProfileEnrollmentOk() (*ProfileEnrollmentPolicyRuleAction, bool)`

GetProfileEnrollmentOk returns a tuple with the ProfileEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileEnrollment

`func (o *ProfileEnrollmentPolicyRuleActions) SetProfileEnrollment(v ProfileEnrollmentPolicyRuleAction)`

SetProfileEnrollment sets ProfileEnrollment field to given value.

### HasProfileEnrollment

`func (o *ProfileEnrollmentPolicyRuleActions) HasProfileEnrollment() bool`

HasProfileEnrollment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


