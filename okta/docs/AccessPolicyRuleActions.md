# AccessPolicyRuleActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enroll** | Pointer to [**PolicyRuleActionsEnroll**](PolicyRuleActionsEnroll.md) |  | [optional] 
**Idp** | Pointer to [**IdpPolicyRuleAction**](IdpPolicyRuleAction.md) |  | [optional] 
**PasswordChange** | Pointer to [**PasswordPolicyRuleAction**](PasswordPolicyRuleAction.md) |  | [optional] 
**SelfServicePasswordReset** | Pointer to [**PasswordPolicyRuleAction**](PasswordPolicyRuleAction.md) |  | [optional] 
**SelfServiceUnlock** | Pointer to [**PasswordPolicyRuleAction**](PasswordPolicyRuleAction.md) |  | [optional] 
**Signon** | Pointer to [**OktaSignOnPolicyRuleSignonActions**](OktaSignOnPolicyRuleSignonActions.md) |  | [optional] 
**AppSignOn** | Pointer to [**AccessPolicyRuleApplicationSignOn**](AccessPolicyRuleApplicationSignOn.md) |  | [optional] 

## Methods

### NewAccessPolicyRuleActions

`func NewAccessPolicyRuleActions() *AccessPolicyRuleActions`

NewAccessPolicyRuleActions instantiates a new AccessPolicyRuleActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyRuleActionsWithDefaults

`func NewAccessPolicyRuleActionsWithDefaults() *AccessPolicyRuleActions`

NewAccessPolicyRuleActionsWithDefaults instantiates a new AccessPolicyRuleActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnroll

`func (o *AccessPolicyRuleActions) GetEnroll() PolicyRuleActionsEnroll`

GetEnroll returns the Enroll field if non-nil, zero value otherwise.

### GetEnrollOk

`func (o *AccessPolicyRuleActions) GetEnrollOk() (*PolicyRuleActionsEnroll, bool)`

GetEnrollOk returns a tuple with the Enroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnroll

`func (o *AccessPolicyRuleActions) SetEnroll(v PolicyRuleActionsEnroll)`

SetEnroll sets Enroll field to given value.

### HasEnroll

`func (o *AccessPolicyRuleActions) HasEnroll() bool`

HasEnroll returns a boolean if a field has been set.

### GetIdp

`func (o *AccessPolicyRuleActions) GetIdp() IdpPolicyRuleAction`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *AccessPolicyRuleActions) GetIdpOk() (*IdpPolicyRuleAction, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *AccessPolicyRuleActions) SetIdp(v IdpPolicyRuleAction)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *AccessPolicyRuleActions) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### GetPasswordChange

`func (o *AccessPolicyRuleActions) GetPasswordChange() PasswordPolicyRuleAction`

GetPasswordChange returns the PasswordChange field if non-nil, zero value otherwise.

### GetPasswordChangeOk

`func (o *AccessPolicyRuleActions) GetPasswordChangeOk() (*PasswordPolicyRuleAction, bool)`

GetPasswordChangeOk returns a tuple with the PasswordChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChange

`func (o *AccessPolicyRuleActions) SetPasswordChange(v PasswordPolicyRuleAction)`

SetPasswordChange sets PasswordChange field to given value.

### HasPasswordChange

`func (o *AccessPolicyRuleActions) HasPasswordChange() bool`

HasPasswordChange returns a boolean if a field has been set.

### GetSelfServicePasswordReset

`func (o *AccessPolicyRuleActions) GetSelfServicePasswordReset() PasswordPolicyRuleAction`

GetSelfServicePasswordReset returns the SelfServicePasswordReset field if non-nil, zero value otherwise.

### GetSelfServicePasswordResetOk

`func (o *AccessPolicyRuleActions) GetSelfServicePasswordResetOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServicePasswordResetOk returns a tuple with the SelfServicePasswordReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServicePasswordReset

`func (o *AccessPolicyRuleActions) SetSelfServicePasswordReset(v PasswordPolicyRuleAction)`

SetSelfServicePasswordReset sets SelfServicePasswordReset field to given value.

### HasSelfServicePasswordReset

`func (o *AccessPolicyRuleActions) HasSelfServicePasswordReset() bool`

HasSelfServicePasswordReset returns a boolean if a field has been set.

### GetSelfServiceUnlock

`func (o *AccessPolicyRuleActions) GetSelfServiceUnlock() PasswordPolicyRuleAction`

GetSelfServiceUnlock returns the SelfServiceUnlock field if non-nil, zero value otherwise.

### GetSelfServiceUnlockOk

`func (o *AccessPolicyRuleActions) GetSelfServiceUnlockOk() (*PasswordPolicyRuleAction, bool)`

GetSelfServiceUnlockOk returns a tuple with the SelfServiceUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceUnlock

`func (o *AccessPolicyRuleActions) SetSelfServiceUnlock(v PasswordPolicyRuleAction)`

SetSelfServiceUnlock sets SelfServiceUnlock field to given value.

### HasSelfServiceUnlock

`func (o *AccessPolicyRuleActions) HasSelfServiceUnlock() bool`

HasSelfServiceUnlock returns a boolean if a field has been set.

### GetSignon

`func (o *AccessPolicyRuleActions) GetSignon() OktaSignOnPolicyRuleSignonActions`

GetSignon returns the Signon field if non-nil, zero value otherwise.

### GetSignonOk

`func (o *AccessPolicyRuleActions) GetSignonOk() (*OktaSignOnPolicyRuleSignonActions, bool)`

GetSignonOk returns a tuple with the Signon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignon

`func (o *AccessPolicyRuleActions) SetSignon(v OktaSignOnPolicyRuleSignonActions)`

SetSignon sets Signon field to given value.

### HasSignon

`func (o *AccessPolicyRuleActions) HasSignon() bool`

HasSignon returns a boolean if a field has been set.

### GetAppSignOn

`func (o *AccessPolicyRuleActions) GetAppSignOn() AccessPolicyRuleApplicationSignOn`

GetAppSignOn returns the AppSignOn field if non-nil, zero value otherwise.

### GetAppSignOnOk

`func (o *AccessPolicyRuleActions) GetAppSignOnOk() (*AccessPolicyRuleApplicationSignOn, bool)`

GetAppSignOnOk returns a tuple with the AppSignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSignOn

`func (o *AccessPolicyRuleActions) SetAppSignOn(v AccessPolicyRuleApplicationSignOn)`

SetAppSignOn sets AppSignOn field to given value.

### HasAppSignOn

`func (o *AccessPolicyRuleActions) HasAppSignOn() bool`

HasAppSignOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


