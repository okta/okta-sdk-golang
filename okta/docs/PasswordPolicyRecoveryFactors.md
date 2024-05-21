# PasswordPolicyRecoveryFactors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OktaCall** | Pointer to [**PasswordPolicyRecoveryFactorSettings**](PasswordPolicyRecoveryFactorSettings.md) |  | [optional] 
**OktaEmail** | Pointer to [**PasswordPolicyRecoveryEmail**](PasswordPolicyRecoveryEmail.md) |  | [optional] 
**OktaSms** | Pointer to [**PasswordPolicyRecoveryFactorSettings**](PasswordPolicyRecoveryFactorSettings.md) |  | [optional] 
**RecoveryQuestion** | Pointer to [**PasswordPolicyRecoveryQuestion**](PasswordPolicyRecoveryQuestion.md) |  | [optional] 

## Methods

### NewPasswordPolicyRecoveryFactors

`func NewPasswordPolicyRecoveryFactors() *PasswordPolicyRecoveryFactors`

NewPasswordPolicyRecoveryFactors instantiates a new PasswordPolicyRecoveryFactors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyRecoveryFactorsWithDefaults

`func NewPasswordPolicyRecoveryFactorsWithDefaults() *PasswordPolicyRecoveryFactors`

NewPasswordPolicyRecoveryFactorsWithDefaults instantiates a new PasswordPolicyRecoveryFactors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOktaCall

`func (o *PasswordPolicyRecoveryFactors) GetOktaCall() PasswordPolicyRecoveryFactorSettings`

GetOktaCall returns the OktaCall field if non-nil, zero value otherwise.

### GetOktaCallOk

`func (o *PasswordPolicyRecoveryFactors) GetOktaCallOk() (*PasswordPolicyRecoveryFactorSettings, bool)`

GetOktaCallOk returns a tuple with the OktaCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaCall

`func (o *PasswordPolicyRecoveryFactors) SetOktaCall(v PasswordPolicyRecoveryFactorSettings)`

SetOktaCall sets OktaCall field to given value.

### HasOktaCall

`func (o *PasswordPolicyRecoveryFactors) HasOktaCall() bool`

HasOktaCall returns a boolean if a field has been set.

### GetOktaEmail

`func (o *PasswordPolicyRecoveryFactors) GetOktaEmail() PasswordPolicyRecoveryEmail`

GetOktaEmail returns the OktaEmail field if non-nil, zero value otherwise.

### GetOktaEmailOk

`func (o *PasswordPolicyRecoveryFactors) GetOktaEmailOk() (*PasswordPolicyRecoveryEmail, bool)`

GetOktaEmailOk returns a tuple with the OktaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaEmail

`func (o *PasswordPolicyRecoveryFactors) SetOktaEmail(v PasswordPolicyRecoveryEmail)`

SetOktaEmail sets OktaEmail field to given value.

### HasOktaEmail

`func (o *PasswordPolicyRecoveryFactors) HasOktaEmail() bool`

HasOktaEmail returns a boolean if a field has been set.

### GetOktaSms

`func (o *PasswordPolicyRecoveryFactors) GetOktaSms() PasswordPolicyRecoveryFactorSettings`

GetOktaSms returns the OktaSms field if non-nil, zero value otherwise.

### GetOktaSmsOk

`func (o *PasswordPolicyRecoveryFactors) GetOktaSmsOk() (*PasswordPolicyRecoveryFactorSettings, bool)`

GetOktaSmsOk returns a tuple with the OktaSms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaSms

`func (o *PasswordPolicyRecoveryFactors) SetOktaSms(v PasswordPolicyRecoveryFactorSettings)`

SetOktaSms sets OktaSms field to given value.

### HasOktaSms

`func (o *PasswordPolicyRecoveryFactors) HasOktaSms() bool`

HasOktaSms returns a boolean if a field has been set.

### GetRecoveryQuestion

`func (o *PasswordPolicyRecoveryFactors) GetRecoveryQuestion() PasswordPolicyRecoveryQuestion`

GetRecoveryQuestion returns the RecoveryQuestion field if non-nil, zero value otherwise.

### GetRecoveryQuestionOk

`func (o *PasswordPolicyRecoveryFactors) GetRecoveryQuestionOk() (*PasswordPolicyRecoveryQuestion, bool)`

GetRecoveryQuestionOk returns a tuple with the RecoveryQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryQuestion

`func (o *PasswordPolicyRecoveryFactors) SetRecoveryQuestion(v PasswordPolicyRecoveryQuestion)`

SetRecoveryQuestion sets RecoveryQuestion field to given value.

### HasRecoveryQuestion

`func (o *PasswordPolicyRecoveryFactors) HasRecoveryQuestion() bool`

HasRecoveryQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


