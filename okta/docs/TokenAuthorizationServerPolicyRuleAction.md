# TokenAuthorizationServerPolicyRuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenLifetimeMinutes** | Pointer to **int32** |  | [optional] 
**InlineHook** | Pointer to [**TokenAuthorizationServerPolicyRuleActionInlineHook**](TokenAuthorizationServerPolicyRuleActionInlineHook.md) |  | [optional] 
**RefreshTokenLifetimeMinutes** | Pointer to **int32** |  | [optional] 
**RefreshTokenWindowMinutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewTokenAuthorizationServerPolicyRuleAction

`func NewTokenAuthorizationServerPolicyRuleAction() *TokenAuthorizationServerPolicyRuleAction`

NewTokenAuthorizationServerPolicyRuleAction instantiates a new TokenAuthorizationServerPolicyRuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAuthorizationServerPolicyRuleActionWithDefaults

`func NewTokenAuthorizationServerPolicyRuleActionWithDefaults() *TokenAuthorizationServerPolicyRuleAction`

NewTokenAuthorizationServerPolicyRuleActionWithDefaults instantiates a new TokenAuthorizationServerPolicyRuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenLifetimeMinutes

`func (o *TokenAuthorizationServerPolicyRuleAction) GetAccessTokenLifetimeMinutes() int32`

GetAccessTokenLifetimeMinutes returns the AccessTokenLifetimeMinutes field if non-nil, zero value otherwise.

### GetAccessTokenLifetimeMinutesOk

`func (o *TokenAuthorizationServerPolicyRuleAction) GetAccessTokenLifetimeMinutesOk() (*int32, bool)`

GetAccessTokenLifetimeMinutesOk returns a tuple with the AccessTokenLifetimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenLifetimeMinutes

`func (o *TokenAuthorizationServerPolicyRuleAction) SetAccessTokenLifetimeMinutes(v int32)`

SetAccessTokenLifetimeMinutes sets AccessTokenLifetimeMinutes field to given value.

### HasAccessTokenLifetimeMinutes

`func (o *TokenAuthorizationServerPolicyRuleAction) HasAccessTokenLifetimeMinutes() bool`

HasAccessTokenLifetimeMinutes returns a boolean if a field has been set.

### GetInlineHook

`func (o *TokenAuthorizationServerPolicyRuleAction) GetInlineHook() TokenAuthorizationServerPolicyRuleActionInlineHook`

GetInlineHook returns the InlineHook field if non-nil, zero value otherwise.

### GetInlineHookOk

`func (o *TokenAuthorizationServerPolicyRuleAction) GetInlineHookOk() (*TokenAuthorizationServerPolicyRuleActionInlineHook, bool)`

GetInlineHookOk returns a tuple with the InlineHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHook

`func (o *TokenAuthorizationServerPolicyRuleAction) SetInlineHook(v TokenAuthorizationServerPolicyRuleActionInlineHook)`

SetInlineHook sets InlineHook field to given value.

### HasInlineHook

`func (o *TokenAuthorizationServerPolicyRuleAction) HasInlineHook() bool`

HasInlineHook returns a boolean if a field has been set.

### GetRefreshTokenLifetimeMinutes

`func (o *TokenAuthorizationServerPolicyRuleAction) GetRefreshTokenLifetimeMinutes() int32`

GetRefreshTokenLifetimeMinutes returns the RefreshTokenLifetimeMinutes field if non-nil, zero value otherwise.

### GetRefreshTokenLifetimeMinutesOk

`func (o *TokenAuthorizationServerPolicyRuleAction) GetRefreshTokenLifetimeMinutesOk() (*int32, bool)`

GetRefreshTokenLifetimeMinutesOk returns a tuple with the RefreshTokenLifetimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenLifetimeMinutes

`func (o *TokenAuthorizationServerPolicyRuleAction) SetRefreshTokenLifetimeMinutes(v int32)`

SetRefreshTokenLifetimeMinutes sets RefreshTokenLifetimeMinutes field to given value.

### HasRefreshTokenLifetimeMinutes

`func (o *TokenAuthorizationServerPolicyRuleAction) HasRefreshTokenLifetimeMinutes() bool`

HasRefreshTokenLifetimeMinutes returns a boolean if a field has been set.

### GetRefreshTokenWindowMinutes

`func (o *TokenAuthorizationServerPolicyRuleAction) GetRefreshTokenWindowMinutes() int32`

GetRefreshTokenWindowMinutes returns the RefreshTokenWindowMinutes field if non-nil, zero value otherwise.

### GetRefreshTokenWindowMinutesOk

`func (o *TokenAuthorizationServerPolicyRuleAction) GetRefreshTokenWindowMinutesOk() (*int32, bool)`

GetRefreshTokenWindowMinutesOk returns a tuple with the RefreshTokenWindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenWindowMinutes

`func (o *TokenAuthorizationServerPolicyRuleAction) SetRefreshTokenWindowMinutes(v int32)`

SetRefreshTokenWindowMinutes sets RefreshTokenWindowMinutes field to given value.

### HasRefreshTokenWindowMinutes

`func (o *TokenAuthorizationServerPolicyRuleAction) HasRefreshTokenWindowMinutes() bool`

HasRefreshTokenWindowMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


