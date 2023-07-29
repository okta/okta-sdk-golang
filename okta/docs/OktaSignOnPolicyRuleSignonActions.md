# OktaSignOnPolicyRuleSignonActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**PolicyAccess**](PolicyAccess.md) |  | [optional] 
**FactorLifetime** | Pointer to **int32** |  | [optional] 
**FactorPromptMode** | Pointer to [**OktaSignOnPolicyFactorPromptMode**](OktaSignOnPolicyFactorPromptMode.md) |  | [optional] 
**RememberDeviceByDefault** | Pointer to **bool** |  | [optional] [default to false]
**RequireFactor** | Pointer to **bool** |  | [optional] [default to false]
**Session** | Pointer to [**OktaSignOnPolicyRuleSignonSessionActions**](OktaSignOnPolicyRuleSignonSessionActions.md) |  | [optional] 

## Methods

### NewOktaSignOnPolicyRuleSignonActions

`func NewOktaSignOnPolicyRuleSignonActions() *OktaSignOnPolicyRuleSignonActions`

NewOktaSignOnPolicyRuleSignonActions instantiates a new OktaSignOnPolicyRuleSignonActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaSignOnPolicyRuleSignonActionsWithDefaults

`func NewOktaSignOnPolicyRuleSignonActionsWithDefaults() *OktaSignOnPolicyRuleSignonActions`

NewOktaSignOnPolicyRuleSignonActionsWithDefaults instantiates a new OktaSignOnPolicyRuleSignonActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *OktaSignOnPolicyRuleSignonActions) GetAccess() PolicyAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *OktaSignOnPolicyRuleSignonActions) GetAccessOk() (*PolicyAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *OktaSignOnPolicyRuleSignonActions) SetAccess(v PolicyAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *OktaSignOnPolicyRuleSignonActions) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetFactorLifetime

`func (o *OktaSignOnPolicyRuleSignonActions) GetFactorLifetime() int32`

GetFactorLifetime returns the FactorLifetime field if non-nil, zero value otherwise.

### GetFactorLifetimeOk

`func (o *OktaSignOnPolicyRuleSignonActions) GetFactorLifetimeOk() (*int32, bool)`

GetFactorLifetimeOk returns a tuple with the FactorLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorLifetime

`func (o *OktaSignOnPolicyRuleSignonActions) SetFactorLifetime(v int32)`

SetFactorLifetime sets FactorLifetime field to given value.

### HasFactorLifetime

`func (o *OktaSignOnPolicyRuleSignonActions) HasFactorLifetime() bool`

HasFactorLifetime returns a boolean if a field has been set.

### GetFactorPromptMode

`func (o *OktaSignOnPolicyRuleSignonActions) GetFactorPromptMode() OktaSignOnPolicyFactorPromptMode`

GetFactorPromptMode returns the FactorPromptMode field if non-nil, zero value otherwise.

### GetFactorPromptModeOk

`func (o *OktaSignOnPolicyRuleSignonActions) GetFactorPromptModeOk() (*OktaSignOnPolicyFactorPromptMode, bool)`

GetFactorPromptModeOk returns a tuple with the FactorPromptMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorPromptMode

`func (o *OktaSignOnPolicyRuleSignonActions) SetFactorPromptMode(v OktaSignOnPolicyFactorPromptMode)`

SetFactorPromptMode sets FactorPromptMode field to given value.

### HasFactorPromptMode

`func (o *OktaSignOnPolicyRuleSignonActions) HasFactorPromptMode() bool`

HasFactorPromptMode returns a boolean if a field has been set.

### GetRememberDeviceByDefault

`func (o *OktaSignOnPolicyRuleSignonActions) GetRememberDeviceByDefault() bool`

GetRememberDeviceByDefault returns the RememberDeviceByDefault field if non-nil, zero value otherwise.

### GetRememberDeviceByDefaultOk

`func (o *OktaSignOnPolicyRuleSignonActions) GetRememberDeviceByDefaultOk() (*bool, bool)`

GetRememberDeviceByDefaultOk returns a tuple with the RememberDeviceByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDeviceByDefault

`func (o *OktaSignOnPolicyRuleSignonActions) SetRememberDeviceByDefault(v bool)`

SetRememberDeviceByDefault sets RememberDeviceByDefault field to given value.

### HasRememberDeviceByDefault

`func (o *OktaSignOnPolicyRuleSignonActions) HasRememberDeviceByDefault() bool`

HasRememberDeviceByDefault returns a boolean if a field has been set.

### GetRequireFactor

`func (o *OktaSignOnPolicyRuleSignonActions) GetRequireFactor() bool`

GetRequireFactor returns the RequireFactor field if non-nil, zero value otherwise.

### GetRequireFactorOk

`func (o *OktaSignOnPolicyRuleSignonActions) GetRequireFactorOk() (*bool, bool)`

GetRequireFactorOk returns a tuple with the RequireFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFactor

`func (o *OktaSignOnPolicyRuleSignonActions) SetRequireFactor(v bool)`

SetRequireFactor sets RequireFactor field to given value.

### HasRequireFactor

`func (o *OktaSignOnPolicyRuleSignonActions) HasRequireFactor() bool`

HasRequireFactor returns a boolean if a field has been set.

### GetSession

`func (o *OktaSignOnPolicyRuleSignonActions) GetSession() OktaSignOnPolicyRuleSignonSessionActions`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *OktaSignOnPolicyRuleSignonActions) GetSessionOk() (*OktaSignOnPolicyRuleSignonSessionActions, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *OktaSignOnPolicyRuleSignonActions) SetSession(v OktaSignOnPolicyRuleSignonSessionActions)`

SetSession sets Session field to given value.

### HasSession

`func (o *OktaSignOnPolicyRuleSignonActions) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


