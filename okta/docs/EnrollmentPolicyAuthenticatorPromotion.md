# EnrollmentPolicyAuthenticatorPromotion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cooldown** | [**AuthenticatorPromotionCooldown**](AuthenticatorPromotionCooldown.md) |  | 
**SkipCount** | Pointer to **int32** | The number of times the user can skip the nudge before prompting stops permanently. A value of &#x60;0&#x60; means the nudge is never limited by skip count. | [optional] 
**Type** | **string** | The promotion period type. Defines when prompting stops permanently for a user.  * &#x60;BY_SKIP_COUNT&#x60;: Prompting stops after the user skips the nudge &#x60;skipCount&#x60; times. | 

## Methods

### NewEnrollmentPolicyAuthenticatorPromotion

`func NewEnrollmentPolicyAuthenticatorPromotion(cooldown AuthenticatorPromotionCooldown, type_ string, ) *EnrollmentPolicyAuthenticatorPromotion`

NewEnrollmentPolicyAuthenticatorPromotion instantiates a new EnrollmentPolicyAuthenticatorPromotion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentPolicyAuthenticatorPromotionWithDefaults

`func NewEnrollmentPolicyAuthenticatorPromotionWithDefaults() *EnrollmentPolicyAuthenticatorPromotion`

NewEnrollmentPolicyAuthenticatorPromotionWithDefaults instantiates a new EnrollmentPolicyAuthenticatorPromotion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCooldown

`func (o *EnrollmentPolicyAuthenticatorPromotion) GetCooldown() AuthenticatorPromotionCooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *EnrollmentPolicyAuthenticatorPromotion) GetCooldownOk() (*AuthenticatorPromotionCooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *EnrollmentPolicyAuthenticatorPromotion) SetCooldown(v AuthenticatorPromotionCooldown)`

SetCooldown sets Cooldown field to given value.


### GetSkipCount

`func (o *EnrollmentPolicyAuthenticatorPromotion) GetSkipCount() int32`

GetSkipCount returns the SkipCount field if non-nil, zero value otherwise.

### GetSkipCountOk

`func (o *EnrollmentPolicyAuthenticatorPromotion) GetSkipCountOk() (*int32, bool)`

GetSkipCountOk returns a tuple with the SkipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCount

`func (o *EnrollmentPolicyAuthenticatorPromotion) SetSkipCount(v int32)`

SetSkipCount sets SkipCount field to given value.

### HasSkipCount

`func (o *EnrollmentPolicyAuthenticatorPromotion) HasSkipCount() bool`

HasSkipCount returns a boolean if a field has been set.

### GetType

`func (o *EnrollmentPolicyAuthenticatorPromotion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnrollmentPolicyAuthenticatorPromotion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnrollmentPolicyAuthenticatorPromotion) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


