# OktaSignOnPolicyRuleActions

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

### NewOktaSignOnPolicyRuleActions

`func NewOktaSignOnPolicyRuleActions() *OktaSignOnPolicyRuleActions`

NewOktaSignOnPolicyRuleActions instantiates a new OktaSignOnPolicyRuleActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaSignOnPolicyRuleActionsWithDefaults

`func NewOktaSignOnPolicyRuleActionsWithDefaults() *OktaSignOnPolicyRuleActions`

NewOktaSignOnPolicyRuleActionsWithDefaults instantiates a new OktaSignOnPolicyRuleActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnroll

`func (o *OktaSignOnPolicyRuleActions) GetEnroll() PolicyRuleActionsEnroll`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *OktaSignOnPolicyRuleActions) GetEnrollOk() (*PolicyRuleActionsEnroll, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *OktaSignOnPolicyRuleActions) SetEnroll(v PolicyRuleActionsEnroll)`

SetEnroll sets Enroll field to given value.

### HasEnroll

`func (o *OktaSignOnPolicyRuleActions) HasEnroll() bool`

HasEnroll returns a boolean if a field has been set.

### GetIdp

`func (o *OktaSignOnPolicyRuleActions) GetIdp() IdpPolicyRuleAction`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *OktaSignOnPolicyRuleActions) GetIdpOk() (*IdpPolicyRuleAction, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *OktaSignOnPolicyRuleActions) SetIdp(v IdpPolicyRuleAction)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *OktaSignOnPolicyRuleActions) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetPasswordChange

`func (o *OktaSignOnPolicyRuleActions) GetPasswordChange() PasswordPolicyRuleAction`

GetPasswordChange returns the PasswordChange field if non-nil, zero value otherwise.

### GetPasswordChangeOk

`func (o *OktaSignOnPolicyRuleActions) GetPasswordChangeOk() (*PasswordPolicyRuleAction, bool)`

GetPasswordChangeOk returns a tuple with the PasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChange

`func (o *OktaSignOnPolicyRuleActions) SetPasswordChange(v PasswordPolicyRuleAction)`

SetPasswordChange sets PasswordChange field to given value.

### HasPasswordChange

`func (o *OktaSignOnPolicyRuleActions) HasPasswordChange() bool`

HasPasswordChange returns a boolean if a field has been set.

### GetSelfServicePasswordReset

`func (o *OktaSignOnPolicyRuleActions) GetSelfServicePasswordReset() PasswordPolicyRuleAction`

GetSelfServicePasswordReset returns the SelfServicePasswordReset field if non-nil, zero value otherwise.

### GetSelfServicePasswordResetOk

`func (o *OktaSignOnPolicyRuleActions) GetSelfServicePasswordResetOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServicePasswordResetOk returns a tuple with the SelfServicePasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServicePasswordReset

`func (o *OktaSignOnPolicyRuleActions) SetSelfServicePasswordReset(v PasswordPolicyRuleAction)`

SetSelfServicePasswordReset sets SelfServicePasswordReset field to given value.

### HasSelfServicePasswordReset

`func (o *OktaSignOnPolicyRuleActions) HasSelfServicePasswordReset() bool`

HasSelfServicePasswordReset returns a boolean if a field has been set.

### GetSelfServiceUnlock

`func (o *OktaSignOnPolicyRuleActions) GetSelfServiceUnlock() PasswordPolicyRuleAction`

GetSelfServiceUnlock returns the SelfServiceUnlock field if non-nil, zero value otherwise.

### GetSelfServiceUnlockOk

`func (o *OktaSignOnPolicyRuleActions) GetSelfServiceUnlockOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServiceUnlockOk returns a tuple with the SelfServiceUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceUnlock

`func (o *OktaSignOnPolicyRuleActions) SetSelfServiceUnlock(v PasswordPolicyRuleAction)`

SetSelfServiceUnlock sets SelfServiceUnlock field to given value.

### HasSelfServiceUnlock

`func (o *OktaSignOnPolicyRuleActions) HasSelfServiceUnlock() bool`

HasSelfServiceUnlock returns a boolean if a field has been set.

### GetSignon

`func (o *OktaSignOnPolicyRuleActions) GetSignon() OktaSignOnPolicyRuleSignonActions`

GetSignon returns the Signon field if non-nil, zero value otherwise.

### GetSignonOk

`func (o *OktaSignOnPolicyRuleActions) GetSignonOk() (*OktaSignOnPolicyRuleSignonActions, bool)`

GetSignonOk returns a tuple with the Signon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignon

`func (o *OktaSignOnPolicyRuleActions) SetSignon(v OktaSignOnPolicyRuleSignonActions)`

SetSignon sets Signon field to given value.

### HasSignon

`func (o *OktaSignOnPolicyRuleActions) HasSignon() bool`

HasSignon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


