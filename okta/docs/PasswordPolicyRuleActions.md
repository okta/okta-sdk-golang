# PasswordPolicyRuleActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enroll** | Pointer to [**PolicyRuleActionsEnroll**](PolicyRuleActionsEnroll.md) |  | [optional] 
**Idp** | Pointer to [**IdpPolicyRuleAction**](IdpPolicyRuleAction.md) |  | [optional] 
**PasswordChange** | Pointer to [**PasswordPolicyRuleAction**](PasswordPolicyRuleAction.md) |  | [optional] 
**SelfServicePasswordReset** | Pointer to [**PasswordPolicyRuleAction**](PasswordPolicyRuleAction.md) |  | [optional] 
**SelfServiceUnlock** | Pointer to [**PasswordPolicyRuleAction**](PasswordPolicyRuleAction.md) |  | [optional] 
**Signon** | Pointer to [**OktaSignOnPolicyRuleSignonActions**](OktaSignOnPolicyRuleSignonActions.md) |  | [optional] 

## Methods

### NewPasswordPolicyRuleActions

`func NewPasswordPolicyRuleActions() *PasswordPolicyRuleActions`

NewPasswordPolicyRuleActions instantiates a new PasswordPolicyRuleActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyRuleActionsWithDefaults

`func NewPasswordPolicyRuleActionsWithDefaults() *PasswordPolicyRuleActions`

NewPasswordPolicyRuleActionsWithDefaults instantiates a new PasswordPolicyRuleActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnroll

`func (o *PasswordPolicyRuleActions) GetEnroll() PolicyRuleActionsEnroll`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *PasswordPolicyRuleActions) GetEnrollOk() (*PolicyRuleActionsEnroll, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *PasswordPolicyRuleActions) SetEnroll(v PolicyRuleActionsEnroll)`

SetEnroll sets Enroll field to given value.

### HasEnroll

`func (o *PasswordPolicyRuleActions) HasEnroll() bool`

HasEnroll returns a boolean if a field has been set.

### GetIdp

`func (o *PasswordPolicyRuleActions) GetIdp() IdpPolicyRuleAction`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *PasswordPolicyRuleActions) GetIdpOk() (*IdpPolicyRuleAction, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *PasswordPolicyRuleActions) SetIdp(v IdpPolicyRuleAction)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *PasswordPolicyRuleActions) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetPasswordChange

`func (o *PasswordPolicyRuleActions) GetPasswordChange() PasswordPolicyRuleAction`

GetPasswordChange returns the PasswordChange field if non-nil, zero value otherwise.

### GetPasswordChangeOk

`func (o *PasswordPolicyRuleActions) GetPasswordChangeOk() (*PasswordPolicyRuleAction, bool)`

GetPasswordChangeOk returns a tuple with the PasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChange

`func (o *PasswordPolicyRuleActions) SetPasswordChange(v PasswordPolicyRuleAction)`

SetPasswordChange sets PasswordChange field to given value.

### HasPasswordChange

`func (o *PasswordPolicyRuleActions) HasPasswordChange() bool`

HasPasswordChange returns a boolean if a field has been set.

### GetSelfServicePasswordReset

`func (o *PasswordPolicyRuleActions) GetSelfServicePasswordReset() PasswordPolicyRuleAction`

GetSelfServicePasswordReset returns the SelfServicePasswordReset field if non-nil, zero value otherwise.

### GetSelfServicePasswordResetOk

`func (o *PasswordPolicyRuleActions) GetSelfServicePasswordResetOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServicePasswordResetOk returns a tuple with the SelfServicePasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServicePasswordReset

`func (o *PasswordPolicyRuleActions) SetSelfServicePasswordReset(v PasswordPolicyRuleAction)`

SetSelfServicePasswordReset sets SelfServicePasswordReset field to given value.

### HasSelfServicePasswordReset

`func (o *PasswordPolicyRuleActions) HasSelfServicePasswordReset() bool`

HasSelfServicePasswordReset returns a boolean if a field has been set.

### GetSelfServiceUnlock

`func (o *PasswordPolicyRuleActions) GetSelfServiceUnlock() PasswordPolicyRuleAction`

GetSelfServiceUnlock returns the SelfServiceUnlock field if non-nil, zero value otherwise.

### GetSelfServiceUnlockOk

`func (o *PasswordPolicyRuleActions) GetSelfServiceUnlockOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServiceUnlockOk returns a tuple with the SelfServiceUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceUnlock

`func (o *PasswordPolicyRuleActions) SetSelfServiceUnlock(v PasswordPolicyRuleAction)`

SetSelfServiceUnlock sets SelfServiceUnlock field to given value.

### HasSelfServiceUnlock

`func (o *PasswordPolicyRuleActions) HasSelfServiceUnlock() bool`

HasSelfServiceUnlock returns a boolean if a field has been set.

### GetSignon

`func (o *PasswordPolicyRuleActions) GetSignon() OktaSignOnPolicyRuleSignonActions`

GetSignon returns the Signon field if non-nil, zero value otherwise.

### GetSignonOk

`func (o *PasswordPolicyRuleActions) GetSignonOk() (*OktaSignOnPolicyRuleSignonActions, bool)`

GetSignonOk returns a tuple with the Signon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignon

`func (o *PasswordPolicyRuleActions) SetSignon(v OktaSignOnPolicyRuleSignonActions)`

SetSignon sets Signon field to given value.

### HasSignon

`func (o *PasswordPolicyRuleActions) HasSignon() bool`

HasSignon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


