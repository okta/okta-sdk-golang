# PolicyRuleActions

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

### NewPolicyRuleActions

`func NewPolicyRuleActions() *PolicyRuleActions`

NewPolicyRuleActions instantiates a new PolicyRuleActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRuleActionsWithDefaults

`func NewPolicyRuleActionsWithDefaults() *PolicyRuleActions`

NewPolicyRuleActionsWithDefaults instantiates a new PolicyRuleActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnroll

`func (o *PolicyRuleActions) GetEnroll() PolicyRuleActionsEnroll`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *PolicyRuleActions) GetEnrollOk() (*PolicyRuleActionsEnroll, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *PolicyRuleActions) SetEnroll(v PolicyRuleActionsEnroll)`

SetEnroll sets Enroll field to given value.

### HasEnroll

`func (o *PolicyRuleActions) HasEnroll() bool`

HasEnroll returns a boolean if a field has been set.

### GetIdp

`func (o *PolicyRuleActions) GetIdp() IdpPolicyRuleAction`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *PolicyRuleActions) GetIdpOk() (*IdpPolicyRuleAction, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *PolicyRuleActions) SetIdp(v IdpPolicyRuleAction)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *PolicyRuleActions) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetPasswordChange

`func (o *PolicyRuleActions) GetPasswordChange() PasswordPolicyRuleAction`

GetPasswordChange returns the PasswordChange field if non-nil, zero value otherwise.

### GetPasswordChangeOk

`func (o *PolicyRuleActions) GetPasswordChangeOk() (*PasswordPolicyRuleAction, bool)`

GetPasswordChangeOk returns a tuple with the PasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChange

`func (o *PolicyRuleActions) SetPasswordChange(v PasswordPolicyRuleAction)`

SetPasswordChange sets PasswordChange field to given value.

### HasPasswordChange

`func (o *PolicyRuleActions) HasPasswordChange() bool`

HasPasswordChange returns a boolean if a field has been set.

### GetSelfServicePasswordReset

`func (o *PolicyRuleActions) GetSelfServicePasswordReset() PasswordPolicyRuleAction`

GetSelfServicePasswordReset returns the SelfServicePasswordReset field if non-nil, zero value otherwise.

### GetSelfServicePasswordResetOk

`func (o *PolicyRuleActions) GetSelfServicePasswordResetOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServicePasswordResetOk returns a tuple with the SelfServicePasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServicePasswordReset

`func (o *PolicyRuleActions) SetSelfServicePasswordReset(v PasswordPolicyRuleAction)`

SetSelfServicePasswordReset sets SelfServicePasswordReset field to given value.

### HasSelfServicePasswordReset

`func (o *PolicyRuleActions) HasSelfServicePasswordReset() bool`

HasSelfServicePasswordReset returns a boolean if a field has been set.

### GetSelfServiceUnlock

`func (o *PolicyRuleActions) GetSelfServiceUnlock() PasswordPolicyRuleAction`

GetSelfServiceUnlock returns the SelfServiceUnlock field if non-nil, zero value otherwise.

### GetSelfServiceUnlockOk

`func (o *PolicyRuleActions) GetSelfServiceUnlockOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServiceUnlockOk returns a tuple with the SelfServiceUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceUnlock

`func (o *PolicyRuleActions) SetSelfServiceUnlock(v PasswordPolicyRuleAction)`

SetSelfServiceUnlock sets SelfServiceUnlock field to given value.

### HasSelfServiceUnlock

`func (o *PolicyRuleActions) HasSelfServiceUnlock() bool`

HasSelfServiceUnlock returns a boolean if a field has been set.

### GetSignon

`func (o *PolicyRuleActions) GetSignon() OktaSignOnPolicyRuleSignonActions`

GetSignon returns the Signon field if non-nil, zero value otherwise.

### GetSignonOk

`func (o *PolicyRuleActions) GetSignonOk() (*OktaSignOnPolicyRuleSignonActions, bool)`

GetSignonOk returns a tuple with the Signon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignon

`func (o *PolicyRuleActions) SetSignon(v OktaSignOnPolicyRuleSignonActions)`

SetSignon sets Signon field to given value.

### HasSignon

`func (o *PolicyRuleActions) HasSignon() bool`

HasSignon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


