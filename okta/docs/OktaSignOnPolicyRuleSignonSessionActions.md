# OktaSignOnPolicyRuleSignonSessionActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxSessionIdleMinutes** | Pointer to **int32** | Maximum number of minutes that a user session can be idle before the session is ended | [optional] [default to 120]
**MaxSessionLifetimeMinutes** | Pointer to **int32** | Maximum number of minutes (from when the user signs in) that a user&#39;s session is active. Set this to force users to sign in again after the number of specified minutes. Disable by setting to &#x60;0&#x60;. | [optional] [default to 0]
**UsePersistentCookie** | Pointer to **bool** | If set to &#x60;false&#x60;, user session cookies only last the length of a browser session. If set to &#x60;true&#x60;, user session cookies last across browser sessions. This setting doesn&#39;t impact administrators who can never have persistent session cookies. This property is read-only for the default rule of the default global session policy. | [optional] [default to false]

## Methods

### NewOktaSignOnPolicyRuleSignonSessionActions

`func NewOktaSignOnPolicyRuleSignonSessionActions() *OktaSignOnPolicyRuleSignonSessionActions`

NewOktaSignOnPolicyRuleSignonSessionActions instantiates a new OktaSignOnPolicyRuleSignonSessionActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaSignOnPolicyRuleSignonSessionActionsWithDefaults

`func NewOktaSignOnPolicyRuleSignonSessionActionsWithDefaults() *OktaSignOnPolicyRuleSignonSessionActions`

NewOktaSignOnPolicyRuleSignonSessionActionsWithDefaults instantiates a new OktaSignOnPolicyRuleSignonSessionActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxSessionIdleMinutes

`func (o *OktaSignOnPolicyRuleSignonSessionActions) GetMaxSessionIdleMinutes() int32`

GetMaxSessionIdleMinutes returns the MaxSessionIdleMinutes field if non-nil, zero value otherwise.

### GetMaxSessionIdleMinutesOk

`func (o *OktaSignOnPolicyRuleSignonSessionActions) GetMaxSessionIdleMinutesOk() (*int32, bool)`

GetMaxSessionIdleMinutesOk returns a tuple with the MaxSessionIdleMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionIdleMinutes

`func (o *OktaSignOnPolicyRuleSignonSessionActions) SetMaxSessionIdleMinutes(v int32)`

SetMaxSessionIdleMinutes sets MaxSessionIdleMinutes field to given value.

### HasMaxSessionIdleMinutes

`func (o *OktaSignOnPolicyRuleSignonSessionActions) HasMaxSessionIdleMinutes() bool`

HasMaxSessionIdleMinutes returns a boolean if a field has been set.

### GetMaxSessionLifetimeMinutes

`func (o *OktaSignOnPolicyRuleSignonSessionActions) GetMaxSessionLifetimeMinutes() int32`

GetMaxSessionLifetimeMinutes returns the MaxSessionLifetimeMinutes field if non-nil, zero value otherwise.

### GetMaxSessionLifetimeMinutesOk

`func (o *OktaSignOnPolicyRuleSignonSessionActions) GetMaxSessionLifetimeMinutesOk() (*int32, bool)`

GetMaxSessionLifetimeMinutesOk returns a tuple with the MaxSessionLifetimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionLifetimeMinutes

`func (o *OktaSignOnPolicyRuleSignonSessionActions) SetMaxSessionLifetimeMinutes(v int32)`

SetMaxSessionLifetimeMinutes sets MaxSessionLifetimeMinutes field to given value.

### HasMaxSessionLifetimeMinutes

`func (o *OktaSignOnPolicyRuleSignonSessionActions) HasMaxSessionLifetimeMinutes() bool`

HasMaxSessionLifetimeMinutes returns a boolean if a field has been set.

### GetUsePersistentCookie

`func (o *OktaSignOnPolicyRuleSignonSessionActions) GetUsePersistentCookie() bool`

GetUsePersistentCookie returns the UsePersistentCookie field if non-nil, zero value otherwise.

### GetUsePersistentCookieOk

`func (o *OktaSignOnPolicyRuleSignonSessionActions) GetUsePersistentCookieOk() (*bool, bool)`

GetUsePersistentCookieOk returns a tuple with the UsePersistentCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePersistentCookie

`func (o *OktaSignOnPolicyRuleSignonSessionActions) SetUsePersistentCookie(v bool)`

SetUsePersistentCookie sets UsePersistentCookie field to given value.

### HasUsePersistentCookie

`func (o *OktaSignOnPolicyRuleSignonSessionActions) HasUsePersistentCookie() bool`

HasUsePersistentCookie returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


