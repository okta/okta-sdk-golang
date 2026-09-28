# AuthenticatorPromotionCooldown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **string** | The ISO 8601 period between prompts. Only month (&#x60;M&#x60;), week (&#x60;W&#x60;), and day (&#x60;D&#x60;) components are allowed (for example, &#x60;P30D&#x60;, &#x60;P2W&#x60;, or &#x60;P1M&#x60;). Year components and time components (hours, minutes, seconds) aren&#39;t permitted: a period that includes a year or time section is rejected even when that section is zero (for example, &#x60;P0Y&#x60;, &#x60;PT0H&#x60;, and &#x60;P1DT0H&#x60; are all invalid). The period must be positive and non-zero (&#x60;P0D&#x60; is invalid); to prompt on every eligible sign-in, use &#x60;type: BY_SIGN_IN&#x60; instead. Required when &#x60;type&#x60; is &#x60;BY_DURATION&#x60; and must be omitted when &#x60;type&#x60; is &#x60;BY_SIGN_IN&#x60;. | [optional] 
**Type** | **string** | The cadence driver between prompts.  * &#x60;BY_DURATION&#x60;: Re-prompts after &#x60;duration&#x60; has elapsed since the user&#39;s last skip. * &#x60;BY_SIGN_IN&#x60;: Re-prompts on every eligible sign-in. No &#x60;duration&#x60; is used. | 

## Methods

### NewAuthenticatorPromotionCooldown

`func NewAuthenticatorPromotionCooldown(type_ string, ) *AuthenticatorPromotionCooldown`

NewAuthenticatorPromotionCooldown instantiates a new AuthenticatorPromotionCooldown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorPromotionCooldownWithDefaults

`func NewAuthenticatorPromotionCooldownWithDefaults() *AuthenticatorPromotionCooldown`

NewAuthenticatorPromotionCooldownWithDefaults instantiates a new AuthenticatorPromotionCooldown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *AuthenticatorPromotionCooldown) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AuthenticatorPromotionCooldown) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AuthenticatorPromotionCooldown) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AuthenticatorPromotionCooldown) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetType

`func (o *AuthenticatorPromotionCooldown) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorPromotionCooldown) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorPromotionCooldown) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


